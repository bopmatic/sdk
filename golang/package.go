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

	const tarRootPath = "pkg"
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

func (pkg *Package) AbsTarballPath() string {
	return filepath.Join(pkg.Proj.Desc.Root, pkg.TarballPath)
}

func (pkg *Package) Deploy() error {
	// the go-swagger generated client seems to produce the cleanest most
	// concise client code between the 3 approaches so make this the default
	return pkg.DeployViaGoSwagger()
}

// Deploy() implemented using protojson & http.Post()
func (pkg *Package) DeployViaProtoJson() error {
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
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(marshalledReq))
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
func (pkg *Package) DeployViaOpenApiGenerator() error {
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
func (pkg *Package) DeployViaGoSwagger() error {
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
	deployPackageParams :=
		service_runner.NewDeployPackageParams().WithBody(deployPackageReq)
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
