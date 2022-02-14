package golang

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
	"github.com/bopmatic/sdk/golang/openapi"
	"github.com/bopmatic/sdk/golang/pb"
	"github.com/bopmatic/sdk/golang/util"
)

const tarRootPath = "pkg"

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
	stdErr io.Writer) (*Package, error) {

	curWd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	err = os.Chdir(proj.Desc.Root)
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
	defer os.RemoveAll(filepath.Join(proj.Desc.Root, workPath))

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
		dstExeDir := filepath.Join(pkgWorkPath, path.Dir(svc.Executable))
		err = os.MkdirAll(dstExeDir, 0755)
		if err != nil {
			return nil, err
		}
		err = util.CopyFileToDir(svc.Executable, dstExeDir)
		if err != nil {
			return nil, err
		}

		dstExePath := filepath.Join(dstExeDir, path.Base(svc.Executable))
		err = util.RunContainerCommand(context.Background(),
			[]string{"strip", dstExePath}, stdOut, stdErr)
		if err != nil {
			return nil, err
		}

		err = os.Chmod(dstExePath, 0755)
		if err != nil {
			return nil, err
		}

		apiDefDstDir := filepath.Join(pkgWorkPath, path.Dir(svc.ApiDefinition))
		err = os.MkdirAll(apiDefDstDir, 0755)
		if err != nil {
			return nil, err
		}
		err = util.CopyFileToDir(svc.ApiDefinition, apiDefDstDir)
		if err != nil {
			return nil, err
		}
	}

	tarFileName := filepath.Join(workPath, "pkg.tar.xz")
	err = util.RunContainerCommand(context.Background(),
		[]string{"tar", "-Jcvf", tarFileName, "-C", workPath, tarRootPath},
		stdOut, stdErr)
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

func NewPackageFromExisting(proj *Project, pkgId string) (*Package, error) {
	packagesPath := filepath.Join(proj.Desc.Root, DefaultArtifactDir,
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
	return filepath.Join(pkg.Proj.Desc.Root, pkg.TarballPath)
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

func (pkg *Package) Deploy(opts ...DeployOption) error {
	// the go-swagger generated client seems to produce the cleanest most
	// concise client code between the 3 approaches so make this the default
	return pkg.DeployViaGoSwagger(opts...)
}

func fillOptions(opts ...DeployOption) *deployOptions {
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

// Deploy() implemented using protojson & http.Post()
func (pkg *Package) DeployViaProtoJson(opts ...DeployOption) error {
	deployOpts := fillOptions(opts...)

	tarballAbsPath := filepath.Join(pkg.Proj.Desc.Root, pkg.TarballPath)
	tarballData, err := ioutil.ReadFile(tarballAbsPath)
	if err != nil {
		return err
	}

	deployPackageReq := &pb.DeployPackageRequest{
		Desc: &pb.PackageDescription{
			ProjectName: pkg.Proj.Desc.Name,
			PackageId:   pkg.Id,
			PackageName: pkg.Name,
			// protojson.Marshal() will base64 encode byte fields so caller does
			// not have to
			PackageXsum:        pkg.Xsum,
			PackageTarballData: tarballData,
		},
	}

	marshalledReq, err := protojson.Marshal(deployPackageReq)
	if err != nil {
		return fmt.Errorf("Marshal failure: %v", err)
	}

	endpoint := "https://api.bopmatic.com/ServiceRunner/DeployPackage"
	resp, err := deployOpts.httpClient.Post(endpoint, "application/json", bytes.NewReader(marshalledReq))
	if err != nil {
		return fmt.Errorf("Client failure: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP failure: %v", resp)
	}

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	var deployReply pb.DeployPackageReply
	err = protojson.Unmarshal(buf.Bytes(), &deployReply)
	if err != nil {
		return fmt.Errorf("Unmarshal failure: %v", err)
	}
	if deployReply.State != pb.PackageState_UPLOADED {
		return fmt.Errorf("DeployPackage failure state: %v", deployReply.State)
	}

	return nil
}

// Deploy() implemented using a client generated with openapi-generator:
//   https://github.com/OpenAPITools/openapi-generator
func (pkg *Package) DeployViaOpenApiGenerator(opts ...DeployOption) error {
	deployOpts := fillOptions(opts...)

	tarballAbsPath := filepath.Join(pkg.Proj.Desc.Root, pkg.TarballPath)
	tarballData, err := ioutil.ReadFile(tarballAbsPath)
	if err != nil {
		return err
	}
	// openapi-generator appears to not base64 encode byte fields so caller
	// has to
	encodedTarData := base64.StdEncoding.EncodeToString(tarballData)
	encodedXsum := base64.StdEncoding.EncodeToString(pkg.Xsum)

	deployPackageReq := openapi.DeployPackageRequest{
		Desc: &openapi.PackageDescription{
			ProjectName:        openapi.PtrString(pkg.Proj.Desc.Name),
			PackageId:          openapi.PtrString(pkg.Id),
			PackageName:        openapi.PtrString(pkg.Name),
			PackageXsum:        openapi.PtrString(encodedXsum),
			PackageTarballData: openapi.PtrString(encodedTarData),
		},
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	config := openapi.NewConfiguration()
	config.HTTPClient = deployOpts.httpClient
	client := openapi.NewAPIClient(config)
	ctx := context.Background()
	deployReply, httpResp, err :=
		client.ServiceRunnerApi.DeployPackage(ctx).Body(deployPackageReq).Execute()
	if err != nil {
		return fmt.Errorf("Client failure: %v", err)
	}
	if httpResp.StatusCode != 200 {
		return fmt.Errorf("HTTP failure: %v", httpResp)
	}
	if *deployReply.State != openapi.UPLOADED {
		return fmt.Errorf("DeployPackage failure state: %v",
			*deployReply.State)
	}

	return nil
}

// Deploy() implemented using a client generated with go-swagger:
//  https://github.com/go-swagger/go-swagger
func (pkg *Package) DeployViaGoSwagger(opts ...DeployOption) error {
	deployOpts := fillOptions(opts...)

	tarballAbsPath := filepath.Join(pkg.Proj.Desc.Root, pkg.TarballPath)
	tarballData, err := ioutil.ReadFile(tarballAbsPath)
	if err != nil {
		return err
	}

	deployPackageReq := &models.DeployPackageRequest{
		Desc: &models.PackageDescription{
			ProjectName: pkg.Proj.Desc.Name,
			PackageID:   pkg.Id,
			PackageName: pkg.Name,
			// go-swagger will base64 encode byte fields so caller does
			// not have to
			PackageXsum:        pkg.Xsum,
			PackageTarballData: tarballData,
		},
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	deployPackageParams := service_runner.NewDeployPackageParams().
		WithBody(deployPackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)
	resp, err := client.ServiceRunner.DeployPackage(deployPackageParams)
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	deployReply := resp.GetPayload()

	if *deployReply.State != models.PackageStateUPLOADED {
		return fmt.Errorf("DeployPackage failure state: %v", deployReply.State)
	}

	return nil
}

// Delete() implemented using a client generated with go-swagger:
//  https://github.com/go-swagger/go-swagger
func Delete(projName string, packageId string, opts ...DeployOption) error {
	deployOpts := fillOptions(opts...)

	deletePackageReq := &models.DeletePackageRequest{
		ProjectName: projName,
		PackageID:   packageId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	deletePackageParams := service_runner.NewDeletePackageParams().
		WithBody(deletePackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)
	resp, err := client.ServiceRunner.DeletePackage(deletePackageParams)
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	deleteReply := resp.GetPayload()

	if *deleteReply.State != models.PackageStateDELETING &&
		*deleteReply.State != models.PackageStateDEACTIVATING {
		return fmt.Errorf("DeletePackage failure state: %v", deleteReply.State)
	}

	return nil
}

// Describe() implemented using a client generated with go-swagger:
//  https://github.com/go-swagger/go-swagger
func Describe(packageId string, opts ...DeployOption) (*pb.DescribePackageReply, error) {

	deployOpts := fillOptions(opts...)

	describePackageReq := &models.DescribePackageRequest{
		PackageID: packageId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	describePackageParams := service_runner.NewDescribePackageParams().
		WithBody(describePackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)
	resp, err := client.ServiceRunner.DescribePackage(describePackageParams)
	if err != nil {
		return nil, fmt.Errorf("Client/HTTP failure: %v", err)
	}
	describeReply := resp.GetPayload()

	pkgStateInt, ok :=
		pb.PackageState_value[string(*describeReply.PackageState)]
	pkgState := pb.PackageState(pkgStateInt)
	if !ok {
		pkgState = pb.PackageState_UNKNOWN_PKG_STATE
	}
	ret := &pb.DescribePackageReply{
		Desc: &pb.PackageDescription{
			ProjectName: describeReply.Desc.ProjectName,
			PackageId:   describeReply.Desc.PackageID,
			PackageName: describeReply.Desc.PackageName,
		},
		PackageState: pkgState,
	}
	if *describeReply.PackageState == models.PackageStatePRODUCTION {
		ret.SiteEndpoint = describeReply.SiteEndpoint
		ret.RpcEndpoints = describeReply.RPCEndpoints
	}

	return ret, nil
}

// List() implemented using a client generated with go-swagger:
//  https://github.com/go-swagger/go-swagger
func List(projName string,
	opts ...DeployOption) ([]pb.ListPackagesReply_ListPackagesItem, error) {

	deployOpts := fillOptions(opts...)

	listPackagesReq := &models.ListPackagesRequest{
		ProjectName: projName,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	listPackagesParams := service_runner.NewListPackagesParams().
		WithBody(listPackagesReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)
	resp, err := client.ServiceRunner.ListPackages(listPackagesParams)
	if err != nil {
		return []pb.ListPackagesReply_ListPackagesItem{},
			fmt.Errorf("Client/HTTP failure: %v", err)
	}
	listReply := resp.GetPayload()

	var itemList []pb.ListPackagesReply_ListPackagesItem
	for _, itemSwag := range listReply.Items {
		itemList = append(itemList, pb.ListPackagesReply_ListPackagesItem{
			ProjectName: itemSwag.ProjectName,
			PackageId:   itemSwag.PackageID,
		})
	}

	return itemList, nil
}

// NewProjectFromPackage instantiates a new Project instance from the specified
// package file
func NewProjectFromPackage(pkgFile string, projRoot string) (*Project, error) {

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
	err = util.RunContainerCommand(context.Background(),
		[]string{"tar", "-Jxvf", pkgFileBase, "-C", tmpDirBase}, os.Stdout,
		os.Stdout)
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

	return NewProject(filepath.Join(projRoot, DefaultProjectFilename))
}

func RemoveStalePackages(proj *Project) error {

	packagesPath := filepath.Join(proj.Desc.Root, DefaultArtifactDir,
		PackagesSubdir)

	return os.RemoveAll(packagesPath)
}