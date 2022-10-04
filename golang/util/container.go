package util

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

const defaultRetries = 1000

const BopmaticBuildImageName = "bopmatic/build:latest"
const DockerInstallErrMsg = "Could not invoke docker; please double check that you have docker installed: %w"

func HasBopmaticBuildImage() (bool, error) {
	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv,
		dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return false, fmt.Errorf(DockerInstallErrMsg, err)
	}

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*5)
	defer cancelFunc()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return false, fmt.Errorf("Failed to list images: %w", err)
	}

	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == BopmaticBuildImageName {
				return true, nil
			}
		}
	}

	return false, nil
}

func RunContainerCommand(ctx context.Context, cmdAndArgs []string,
	stdOut io.Writer, stdErr io.Writer) error {

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
		resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil,
			nil, "")
		if err != nil {
			return fmt.Errorf("Failed to create container: %w", err)
		}

		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
			return fmt.Errorf("Failed to start container: %w", err)
		}

		logOutputOpts := types.ContainerLogsOptions{
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
