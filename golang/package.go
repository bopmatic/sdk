/* Copyright © 2022-2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
	"github.com/bopmatic/sdk/golang/pb"
	"github.com/bopmatic/sdk/golang/util"
)

const tarRootPath = "pkg"
const defaultRetries = 10

type Package struct {
	Proj        *Project
	Id          string
	Name        string
	TarballPath string
	Xsum        []byte
}

func (pkg *Package) String() string {
	var sb strings.Builder

	sb.WriteString("Package:\n")
	sb.WriteString(fmt.Sprintf("\tProject: %v\n", pkg.Proj.Desc.Name))
	sb.WriteString(fmt.Sprintf("\tId: %v\n", pkg.Id))
	sb.WriteString(fmt.Sprintf("\tName: %v\n", pkg.Name))
	sb.WriteString(fmt.Sprintf("\tTarball: %v\n", pkg.TarballPath))
	sb.WriteString(fmt.Sprintf("\tXsum: %v\n", hex.EncodeToString(pkg.Xsum)))

	return sb.String()
}

// @todo add logic to avoid generating a new package when the underlying
// binaries haven't changed; currently this isn't happening due to the
// copied files in the tarball having different timestamps. Fixing this is
// non-trivial since golang's os.Stat() doesn't return st_atim and os.Chtimes()
// isn't nanosecond precision
func NewPackage(pkgName string, proj *Project, stdOut io.Writer,
	stdErr io.Writer, opts ...PkgOption) (*Package, error) {

	pkgOpts := fillPkgOptions(opts...)

	curWd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	err = os.Chdir(proj.Desc.root)
	if err != nil {
		return nil, err
	}

	packagesPath := filepath.Join(DefaultArtifactDir, PackagesSubdir)
	err = os.MkdirAll(packagesPath, 0755)
	if err != nil {
		return nil, err
	}

	workPath, err := ioutil.TempDir(packagesPath, "in_prgs")
	if err != nil {
		return nil, err
	}
	// subtle; Join() required here since workPath is relative and the defer
	// statement will execute subsequent to Chdir() back to curWd
	defer os.RemoveAll(filepath.Join(proj.Desc.root, workPath))

	pkgWorkPath := filepath.Join(workPath, tarRootPath)
	err = os.MkdirAll(pkgWorkPath, 0755)
	if err != nil {
		return nil, err
	}

	err = util.CopyFileToDir(DefaultProjectFilename, pkgWorkPath)
	if err != nil {
		return nil, err
	}

	if proj.Desc.SiteAssets != "" {
		dstSiteDir := filepath.Join(pkgWorkPath, path.Dir(proj.Desc.SiteAssets))
		err = os.MkdirAll(dstSiteDir, 0755)
		if err != nil {
			return nil, err
		}
		dstSitePath := filepath.Join(dstSiteDir, path.Base(proj.Desc.SiteAssets))
		err = util.CopyDir(proj.Desc.SiteAssets, dstSitePath)
		if err != nil {
			return nil, err
		}
	}

	for _, svc := range proj.Desc.Services {
		err = copyExecAssets(pkgWorkPath, &svc, stdOut, stdErr, pkgOpts)
		if err != nil {
			return nil, err
		}
		err = copyApiDefAssets(pkgWorkPath, &svc, stdOut, stdErr, pkgOpts)
		if err != nil {
			return nil, err
		}
	}

	// try to mitigate spurious "file changed as we read it" errors from tar
	syscall.Sync()

	tarFileName := filepath.Join(workPath, "pkg.tar.xz")
	if pkgOpts.useHostOS {
		err = util.RunHostCommand(context.Background(),
			[]string{"tar", "-Jcvf", tarFileName, "-C", workPath, tarRootPath},
			stdOut, stdErr)
	} else {
		err = util.RunContainerCommand(context.Background(),
			[]string{"tar", "-Jcvf", tarFileName, "-C", workPath, tarRootPath},
			stdOut, stdErr)
	}

	if err != nil {
		return nil, err
	}

	tarFileContent, err := ioutil.ReadFile(tarFileName)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	hasher.Write(tarFileContent)
	xsum := hasher.Sum(nil)

	xsumStr := hex.EncodeToString(xsum)
	tarFileBase := xsumStr + ".tar.xz"
	finalTarFileName := filepath.Join(packagesPath, tarFileBase)

	err = util.RenameFile(tarFileName, finalTarFileName)
	if err != nil {
		return nil, err
	}

	_ = os.Chdir(curWd)

	pkg := &Package{
		Proj:        proj,
		Id:          xsumStr[0:8],
		Name:        pkgName,
		TarballPath: finalTarFileName,
		Xsum:        xsum,
	}

	return pkg, nil
}

func copyExecAssets(pkgWorkPath string, svc *Service, stdOut io.Writer,
	stdErr io.Writer, pkgOpts *pkgOptions) error {

	if svc.ExecAssets == "" {
		dstExeDir := filepath.Join(pkgWorkPath, path.Dir(svc.Executable))
		err := os.MkdirAll(dstExeDir, 0755)
		if err != nil {
			return err
		}
		err = util.CopyFileToDir(svc.Executable, dstExeDir)
		if err != nil {
			return err
		}

		dstExePath := filepath.Join(dstExeDir, path.Base(svc.Executable))
		if pkgOpts.useHostOS {
			err = util.RunHostCommand(context.Background(),
				[]string{"strip", dstExePath}, stdOut, stdErr)
		} else {
			err = util.RunContainerCommand(context.Background(),
				[]string{"strip", dstExePath}, stdOut, stdErr)
		}
		if err != nil {
			return err
		}

		err = os.Chmod(dstExePath, 0755)
		if err != nil {
			return err
		}

		return nil
	} // else

	return util.CopyDir(svc.ExecAssets,
		filepath.Join(pkgWorkPath, path.Base(svc.ExecAssets)))
}

func copyApiDefAssets(pkgWorkPath string, svc *Service, stdOut io.Writer,
	stdErr io.Writer, pkgOpts *pkgOptions) error {

	if svc.ApiDefAssets == "" {
		apiDefDstDir := filepath.Join(pkgWorkPath, path.Dir(svc.ApiDefinition))
		err := os.MkdirAll(apiDefDstDir, 0755)
		if err != nil {
			return err
		}
		err = util.CopyFileToDir(svc.ApiDefinition, apiDefDstDir)
		if err != nil {
			return err
		}

		return nil
	} // else

	return util.CopyDir(svc.ApiDefAssets,
		filepath.Join(pkgWorkPath, path.Base(svc.ApiDefAssets)))
}

func NewPackageFromExisting(proj *Project, pkgId string) (*Package, error) {
	packagesPath := filepath.Join(proj.Desc.root, DefaultArtifactDir,
		PackagesSubdir)

	entries, err := ioutil.ReadDir(packagesPath)
	if err != nil {
		return nil, err
	}

	var pkgFile = ""
	if pkgId == "" {
		if len(entries) == 1 {
			pkgFile = filepath.Join(packagesPath, entries[0].Name())
		} else {
			return nil, fmt.Errorf("Multiple packages present; must specify pkgId")
		}
	} else {
		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), pkgId) {
				if pkgFile == "" {
					pkgFile = filepath.Join(packagesPath, entry.Name())
				} else {
					return nil, fmt.Errorf("Abiguous pkgId:%v; multiple matches", pkgId)
				}
			}
		}
	}

	if pkgFile == "" {
		return nil, fmt.Errorf("Cannot find pkgId:%v", pkgId)
	}

	tarballData, err := ioutil.ReadFile(pkgFile)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	hasher.Write(tarballData)
	xsumVal := hasher.Sum(nil)
	actualXsum := hex.EncodeToString(xsumVal)
	pkgFileBase := path.Base(pkgFile)
	expectedXsum := strings.TrimSuffix(pkgFileBase, ".tar.xz")
	if actualXsum != expectedXsum {
		return nil, fmt.Errorf("%v failed checksum verification", pkgFile)
	}
	return &Package{
		Proj:        proj,
		Id:          actualXsum[0:8],
		TarballPath: pkgFile,
		Xsum:        xsumVal,
	}, nil
}

func (pkg *Package) AbsTarballPath() string {
	return filepath.Join(pkg.Proj.Desc.root, pkg.TarballPath)
}

type deployOptions struct {
	httpClient *http.Client
}

type DeployOption func(*deployOptions) *deployOptions

func DeployOptHttpClient(httpClient *http.Client) DeployOption {
	return func(do *deployOptions) *deployOptions {
		do.httpClient = httpClient

		return do
	}
}

func fillDeployOptions(opts ...DeployOption) *deployOptions {
	options := &deployOptions{
		httpClient: http.DefaultClient,
	}
	for _, optApplyFunc := range opts {
		if optApplyFunc != nil {
			options = optApplyFunc(options)
		}
	}

	return options
}

func uploadToURL(url string, data []byte) error {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-gtar")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("HTTP Put failed status:%v headers:%v\n",
			resp.Status, resp.Header)

		return fmt.Errorf(errMsg)
	}

	return nil
}

// Upload() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func (pkg *Package) Upload(opts ...DeployOption) error {
	deployOpts := fillDeployOptions(opts...)

	tarballAbsPath := filepath.Join(pkg.Proj.Desc.root, pkg.TarballPath)
	tarballData, err := ioutil.ReadFile(tarballAbsPath)
	if err != nil {
		return err
	}

	uploadUrlReq := &models.GetUploadURLRequest{
		Key: path.Base(pkg.TarballPath),
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	uploadUrlParams := service_runner.NewGetUploadURLParams().
		WithBody(uploadUrlReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var uploadUrlResp *service_runner.GetUploadURLOK

	for retries := defaultRetries; retries > 0; retries-- {
		uploadUrlResp, err = client.ServiceRunner.GetUploadURL(uploadUrlParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	uploadUrlReply := uploadUrlResp.GetPayload()
	err = uploadToURL(uploadUrlReply.URL, tarballData)
	if err != nil {
		return err
	}

	uploadPackageReq := &models.UploadPackageRequest{
		ProjID:            pkg.Proj.Desc.Id,
		PackageXsum:       pkg.Xsum,
		PackageTarballURL: uploadUrlReply.URL,
	}

	uploadPackageParams := service_runner.NewUploadPackageParams().
		WithBody(uploadPackageReq).WithHTTPClient(deployOpts.httpClient)
	var resp *service_runner.UploadPackageOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.UploadPackage(uploadPackageParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	uploadReply := resp.GetPayload()

	if *uploadReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("UploadPackage failure(%v): %v",
			*uploadReply.Result.Status, uploadReply.Result.StatusDetail)
	}
	pkg.Id = uploadReply.PkgID

	return nil
}

// Deploy() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func (pkg *Package) Deploy(envId string, opts ...DeployOption) (string, error) {
	err := pkg.Upload(opts...)
	if err != nil {
		return "", err
	}

	deployment := NewDeployment(pkg.Id, pkg.Proj.Desc.Id, envId)
	err = deployment.Deploy(opts...)
	if err != nil {
		return "", err
	}

	return deployment.DeployId, nil
}

// Delete() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func Delete(packageId string, opts ...DeployOption) error {
	deployOpts := fillDeployOptions(opts...)

	deletePackageReq := &models.DeletePackageRequest{
		PackageID: packageId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	deletePackageParams := service_runner.NewDeletePackageParams().
		WithBody(deletePackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DeletePackageOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DeletePackage(deletePackageParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	deleteReply := resp.GetPayload()
	if *deleteReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("DeletePackage failure(%v): %v",
			*deleteReply.Result.Status, deleteReply.Result.StatusDetail)
	}

	return nil
}

// Describe() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func Describe(packageId string, opts ...DeployOption) (*pb.PackageDescription, error) {

	deployOpts := fillDeployOptions(opts...)

	describePackageReq := &models.DescribePackageRequest{
		PackageID: packageId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	describePackageParams := service_runner.NewDescribePackageParams().
		WithBody(describePackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DescribePackageOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribePackage(describePackageParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("Client/HTTP failure: %v", err)
	}
	describeReply := resp.GetPayload()
	if *describeReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribePackage failure(%v): %v",
			*describeReply.Result.Status, describeReply.Result.StatusDetail)
	}

	pkgStateInt, ok :=
		pb.PackageState_value[string(*describeReply.Desc.State)]
	pkgState := pb.PackageState(pkgStateInt)
	if !ok {
		pkgState = pb.PackageState_UNKNOWN_PKG_STATE
	}

	ret := &pb.PackageDescription{
		PackageId:   describeReply.Desc.PackageID,
		ProjId:      describeReply.Desc.ProjID,
		State:       pkgState,
		UploadTime:  convertRESTTimeToInt(describeReply.Desc.UploadTime),
		PackageSize: convertRESTIntToInt(describeReply.Desc.PackageSize),
	}

	return ret, nil
}

// List() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func List(projId string,
	opts ...DeployOption) ([]pb.ListPackagesReply_ListPackagesItem, error) {

	deployOpts := fillDeployOptions(opts...)

	listPackagesReq := &models.ListPackagesRequest{
		ProjID: projId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	listPackagesParams := service_runner.NewListPackagesParams().
		WithBody(listPackagesReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListPackagesOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListPackages(listPackagesParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return []pb.ListPackagesReply_ListPackagesItem{},
			fmt.Errorf("Client/HTTP failure: %v", err)
	}
	listReply := resp.GetPayload()
	if *listReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return []pb.ListPackagesReply_ListPackagesItem{},
			fmt.Errorf("ListPackages failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	var itemList []pb.ListPackagesReply_ListPackagesItem
	for _, itemSwag := range listReply.Items {
		itemList = append(itemList, pb.ListPackagesReply_ListPackagesItem{
			ProjId:    itemSwag.ProjID,
			PackageId: itemSwag.PackageID,
		})
	}

	return itemList, nil
}

type pkgOptions struct {
	useHostOS bool
}

type PkgOption func(*pkgOptions) *pkgOptions

// PkgOptUseHostOS() instructs NewProjectFromPackage() to utilize tar&xz
// from the host operating system directly rather than the default behavior
// of utilizing them via the Bopmatic Build container. It is the caller's
// responsibility to ensure tar&xz are installed on the host OS when utilizing
// this option.
func PkgOptUseHostOS() PkgOption {
	return func(po *pkgOptions) *pkgOptions {
		po.useHostOS = true

		return po
	}
}

func fillPkgOptions(opts ...PkgOption) *pkgOptions {
	options := &pkgOptions{
		useHostOS: false,
	}
	for _, optApplyFunc := range opts {
		if optApplyFunc != nil {
			options = optApplyFunc(options)
		}
	}

	return options
}

// NewProjectFromPackage instantiates a new Project instance from the specified
// package file
func NewProjectFromPackage(pkgFile string, projRoot string,
	opts ...PkgOption) (*Project, error) {

	pkgOpts := fillPkgOptions(opts...)

	curWd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	projRootParent := path.Dir(projRoot)
	projRootBase := path.Base(projRoot)
	pkgFileBase := path.Base(pkgFile)

	tmpDir, err := ioutil.TempDir(projRootParent, projRootBase)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	tarballData, err := ioutil.ReadFile(pkgFile)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	hasher.Write(tarballData)
	actualXsum := hex.EncodeToString(hasher.Sum(nil))
	expectedXsum := strings.TrimSuffix(pkgFileBase, ".tar.xz")
	if actualXsum != expectedXsum {
		return nil, fmt.Errorf("%v failed checksum verification", pkgFile)
	}

	err = util.CopyFileToDir(pkgFile, projRootParent)
	if err != nil {
		return nil, err
	}

	pkgFileCopy := filepath.Join(projRootParent, pkgFileBase)
	defer os.Remove(pkgFileCopy)

	err = os.Chdir(projRootParent)
	if err != nil {
		return nil, err
	}

	tmpDirBase := path.Base(tmpDir)
	if pkgOpts.useHostOS {
		err = util.RunHostCommand(context.Background(),
			[]string{"tar", "-Jxvf", pkgFileBase, "-C", tmpDirBase}, os.Stdout,
			os.Stdout)
	} else {
		err = util.RunContainerCommand(context.Background(),
			[]string{"tar", "-Jxvf", pkgFileBase, "-C", tmpDirBase}, os.Stdout,
			os.Stdout)
	}
	if err != nil {
		_ = os.Chdir(curWd)
		return nil, err
	}

	err = os.Rename(filepath.Join(tmpDirBase, tarRootPath), projRootBase)
	if err != nil {
		_ = os.Chdir(curWd)
		return nil, err
	}
	_ = os.Chdir(curWd)

	return NewProject(filepath.Join(projRoot, DefaultProjectFilename),
		ProjectOptValidateSiteAssets())
}

func RemoveStalePackages(proj *Project) error {

	packagesPath := filepath.Join(proj.Desc.root, DefaultArtifactDir,
		PackagesSubdir)

	return os.RemoveAll(packagesPath)
}
