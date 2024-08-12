package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

const defaultRetries = 100

const (
	BopmaticImageRepo      = "bopmatic/build"
	BopmaticImageTag       = "latest"
	BopmaticBuildImageName = BopmaticImageRepo + ":" + BopmaticImageTag

	DockerInstallErrMsg = "Could not invoke docker; please double check that you have docker installed: %w"
)

var ErrImageNotFound = errors.New("image does not exist")

func HasBopmaticBuildImage() (bool, error) {
	return HasImage(BopmaticImageRepo, BopmaticImageTag)
}

func HasImage(repository string, tag string) (bool, error) {
	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv,
		dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return false, fmt.Errorf(DockerInstallErrMsg, err)
	}

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*5)
	defer cancelFunc()

	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("Failed to list images: %w", err)
	}

	for _, img := range images {
		for _, t := range img.RepoTags {
			if t == repository+":"+tag {
				return true, nil
			}
		}
	}

	return false, nil
}

func GetLocalImageDigest(repository string, tag string) (string, error) {
	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv,
		dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf(DockerInstallErrMsg, err)
	}

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*5)
	defer cancelFunc()

	imgInspect, _, err := cli.ImageInspectWithRaw(ctx, repository+":"+tag)
	if err != nil {
		if client.IsErrNotFound(err) {
			err = fmt.Errorf("Could not find local image digest for %v:%v: %w",
				repository, tag, ErrImageNotFound)
		}
		return "", err
	}

	for _, digest := range imgInspect.RepoDigests {
		// cannot find documentation on RepoDigests content; just guessing
		// here that there may be potentially more digest checksums than sha256
		if strings.Contains(digest, "sha256") &&
			strings.Contains(digest, repository) {
			// digest format here differs from that of GetRemoteImageDigest()
			// the local digest includes a repo w/ @ prefix so strip that here
			digestParts := strings.Split(digest, "@")
			if len(digestParts) != 2 {
				return "", fmt.Errorf("Could not parse local digest %v", digest)
			}
			return digestParts[1], nil
		}
	}

	err = fmt.Errorf("Could not find local image digest for %v:%v: %w",
		repository, tag, ErrImageNotFound)

	return "", err
}

// @todo investigate whether this is achievable with github.com/docker/docker
// rather than writing our own custom json parsing logic
func GetRemoteImageDigest(repository string, tag string) (string, error) {
	// note TagsResult&TagsResponse as defined here do not capture the complete
	// set of information returned from the tagsURL; we only define the minimum
	// subset of fields we need to extract the image digest
	type TagsResult struct {
		Name   string `json:"name"`
		Digest string `json:"digest"`
	}
	type TagsResponse struct {
		Count   int          `json:"count"`
		Results []TagsResult `json:"results"`
	}

	tagsURL := fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/%v/tags",
		repository)

	client := http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Get(tagsURL)
	if err != nil {
		return "", err
	}

	tagsJsonDoc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tagsResponse TagsResponse
	err = json.Unmarshal(tagsJsonDoc, &tagsResponse)
	if err != nil {
		return "", err
	}

	for _, tagResult := range tagsResponse.Results {
		if tagResult.Name == tag {
			return tagResult.Digest, nil
		}
	}

	return "", fmt.Errorf("Could not find tag %v for repo %v at %v: %w", tag,
		repository, tagsURL, ErrImageNotFound)
}

func DoesLocalImageNeedUpdate(repository string, tag string) (bool, error) {
	localDigest, err := GetLocalImageDigest(repository, tag)
	if errors.Is(err, ErrImageNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}
	remoteDigest, err := GetRemoteImageDigest(repository, tag)
	if err != nil {
		return false, err
	}

	return localDigest != remoteDigest, nil
}

func RunContainerCommand(ctx context.Context, cmdAndArgs []string,
	stdOut io.Writer, stdErr io.Writer) error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}

	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv,
		dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf(DockerInstallErrMsg, err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Could not get current working dir: %w", err)
	}

	hostConfig := &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: pwd,
				Target: pwd,
			},
		},
		Binds: []string{
			"/etc/passwd:/etc/passwd",
		},
	}

	if homeDir != "" {
		hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: homeDir,
			Target: homeDir,
		})
	}

	containerConfig := &container.Config{
		User:       fmt.Sprintf("%v:%v", os.Geteuid(), os.Getegid()),
		Cmd:        cmdAndArgs,
		Image:      BopmaticBuildImageName,
		WorkingDir: pwd,
	}

	var retErr error

	// retry due to occasional spurious 'container not found' & 'no such file
	// or directory'
	for retries := defaultRetries; retries > 0; retries-- {
		retErr = nil

		resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil,
			nil, "")
		if err != nil {
			return fmt.Errorf("Failed to create container: %w", err)
		}

		err = cli.ContainerStart(ctx, resp.ID, container.StartOptions{})
		if err != nil {
			return fmt.Errorf("Failed to start container: %w", err)
		}

		logOutputOpts := container.LogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
		}
		logOutput, err := cli.ContainerLogs(ctx, resp.ID, logOutputOpts)
		if err != nil {
			retErr = fmt.Errorf("Failed to get container output: %w", err)
			time.Sleep(10 * time.Millisecond)
			continue
		}
		defer logOutput.Close()

		// the container's stdout and stderr are muxed into a Docker specific
		// output format; so we demux them here
		stdcopy.StdCopy(stdOut, stdErr, logOutput)

		statusCh, errCh := cli.ContainerWait(ctx, resp.ID,
			container.WaitConditionRemoved)
		select {
		case err := <-errCh:
			if err != nil {
				if dockerClient.IsErrNotFound(err) {
					retErr = fmt.Errorf("Container run failed/ not found: %w\n", err)
					time.Sleep(10 * time.Millisecond)
					continue
				}

				return fmt.Errorf("Container run failed: %w\n", err)
			}
		case status := <-statusCh:
			if status.StatusCode != 0 {
				return fmt.Errorf("%v failed with status:%v", cmdAndArgs[0],
					status.StatusCode)
			}
		}

		break
	}

	return retErr
}
