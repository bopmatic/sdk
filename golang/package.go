package golang

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

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
