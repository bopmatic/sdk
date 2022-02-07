package golang

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPackage(t *testing.T) {
	var testCasesDir = filepath.Join("test_assets", "package")

	testCases, err := ioutil.ReadDir(testCasesDir)
	if err != nil {
		t.Errorf("Failed to open project test case dir: %v", err)
	}

	for _, tCase := range testCases {
		if !tCase.IsDir() {
			continue
		}

		projFile := filepath.Join(testCasesDir, tCase.Name(), "Bopmatic.yaml")
		proj, err := NewProject(projFile)
		if err != nil {
			t.Errorf("Failed to parse project: %v", err)
		}

		pkg, err := proj.NewPackage("somepkg", os.Stdout, os.Stderr)
		if err != nil {
			t.Errorf("Failed to create package: %v", err)
		}

		derivedPkgId := hex.EncodeToString(pkg.Xsum)[0:8]
		if derivedPkgId != pkg.Id {
			t.Errorf("Package Id is %v but expected %v", pkg.Id, derivedPkgId)
		}
		if pkg.Name != "somepkg" {
			t.Errorf("Package Name is %v but expected %v", pkg.Name, "somepkg")
		}
		expectedTarballPath := fmt.Sprintf(".bopmatic/pkgs/%v.tar.xz",
			hex.EncodeToString(pkg.Xsum))
		if pkg.TarballPath != expectedTarballPath {
			t.Errorf("Package Tarball is %v but expected %v", pkg.TarballPath,
				expectedTarballPath)
		}

		tarballData, err := ioutil.ReadFile(pkg.AbsTarballPath())
		if err != nil {
			t.Errorf("Failed to read package tarball: %v", err)
		}
		hasher := sha256.New()
		hasher.Write(tarballData)
		xsum := hasher.Sum(nil)
		if bytes.Compare(xsum, pkg.Xsum) != 0 {
			t.Errorf("Package %v checksum failed", pkg.Id)
		}

		fmt.Printf("Attempting upload via openapi-generator....\n")
		err = pkg.DeployViaOpenApiGenerator()
		if err != nil {
			t.Errorf("Failed to upload package via openapi-generator: %v", err)
		}

		fmt.Printf("Attempting upload via protojson....\n")

		err = pkg.DeployViaProtoJson()
		if err != nil {
			t.Errorf("Failed to upload package via protojson: %v", err)
		}

		fmt.Printf("Attempting upload via go-swagger....\n")

		err = pkg.DeployViaGoSwagger()
		if err != nil {
			t.Errorf("Failed to upload package via go-swagger: %v", err)
		}
	}
}
