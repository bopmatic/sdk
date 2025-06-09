package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var expectedSuccesses = map[string]string{
	"multiSvcMultiDb":         "AoRGICrBWT3X1vEVCzZ_8BceOBdui8xnvoE4CakmxIo=",
	"multiSvcMultiRpc":        "zngIO3KMWHzBB296nwuAKL05iNIomZVCc-dpCGaK_7A=",
	"multiSvcOneRpc":          "2hSRbBxU5GIaVq5Nrk2_63wToYu675MrDd-GvRJHsJQ=",
	"oneSvcOneDb":             "l_oIVApKgcjugJ4NYEQ4-QDdNIrkEA3t7ylDxhfVWV4=",
	"oneSvcMultiRpc":          "vU8ae5ZIS51Lf2bGsL8p2pKHrAvwB8kVhYzZoBEvMHU=",
	"oneSvcOneRpc":            "cTEdYL572lNMeer_JU0K-j44hewJ-v708GE_5XYOVSU=",
	"oneSvcOneRpcOpenAPI":     "UqzkfbdwQtdNzVmf5WxF2Nv0nhCzkI8J4vYkpr8YEQw=",
	"staticOnly":              "cFFz0ff1CLvOmjGOVcL0z39SF_NTIeYIyID4Qj2wBm4=",
	"svcsOnly":                "UzRKDEQ4WFKW7LgrEgU57h1OSrOzB_V70zh5ranwubg=",
	"oneSvcOneUserGroup":      "uvQSroXeUn4WCw0sTdRNupuKHx4LqZaw9e4ZNZjw0bE=",
	"multiSvcMultiUserGroup":  "YHDepdKUmgAsT2PxSkaENJ5PjFOi9W9GItfv0or4uxY=",
	"oneSvcWithApiAssets":     "pqItjeJrKhxXYLbz8Hw2NSSUUJ7rZ5F4Y3ePyykC1pQ=",
	"oneSvcWithRuntimeConfig": "26suPRpTK8lssw3TEz-_mVKhSLGoJvZ6GfnVQT93LwM=",
	"oneSvcOneObjStore":       "3IrKn2JkDzJAj1CqwLOZsTV7J_FSqzAEomnbCHDV4_c=",
	"multiSvcMultiObjStore":   "_os6gjbOovAEQMM3fMgkSgI4QXsWE7aFF55H9pIE1EI=",
}

func TestProjectParse(t *testing.T) {
	var testCasesDir = filepath.Join("test_assets", "project")

	testCases, err := ioutil.ReadDir(testCasesDir)
	if err != nil {
		t.Errorf("Failed to open project test case dir: %v", err)
	}

	expectedFailures := map[string]string{
		"emptySiteAssets":      "Site assets siteAssetsEmpty is missing index.html",
		"missingSiteAssets":    "Could not open site assets missing: open test_assets/project/missingSiteAssets/missing: no such file or directory",
		"missingRuntimeConfig": "Could not open site assets(site)' runtime config(bopmatic-config.js): open test_assets/project/missingRuntimeConfig/site/bopmatic-config.js: no such file or directory",
		"missingPort":          "Service Greeter definition is missing required field port",
		"missingProjectName":   "Project test_assets/project/missingProjectName/Bopmatic.yaml definition is missing required field name",
		"missingSvcInApiDef":   "Failed to find service Missing in pb/greeter.proto",
		"missingSvcName":       "Service 0 definition is missing required field name",
		"noRpcsInSvc":          "Service Greeter in pb/greeter.proto must define at least 1 RPC",
		"noSvcInApiDef":        "pb/greeter.proto does not define any services; looking for Greeter",
		"portConflict":         "Service Orders port 26001 conflicts with service Greeter",
		"missingApiDefField":   "Service Greeter definition is missing required field apidef",
		"missingApiDefFile":    "Failed to open API definition for Service Greeter: open test_assets/project/missingApiDefFile/pb/greeter.proto: no such file or directory",
		"missingDbName":        "Database in Project Foo definition is missing required field name",
		"missingDbTable":       "Database Customers in Project Foo must define at least 1 table",
		"missingDbTableName":   "Table in Database Customers definition is missing required field name",
		"missingDbSvcAccess":   "Database Customers in Project Foo must define at least 1 service access",
		"missingProject":       "Failed to open project test_assets/project/missingProject/Bopmatic.yaml: open test_assets/project/missingProject/Bopmatic.yaml: no such file or directory",
		"corruptApiDef":        "Failed to parse pb/greeter.proto for service Greeter: found \"CORRUPT\" but expected [;]",
		"missingBuildCmd":      "Project Foo definition is missing required field buildcmd",
		"missingExecutable":    "Service Greeter definition is missing required field executable",
		"noSuchDbSvcAccess":    "Database Customers in Project Foo defines access for service NoSuchSvc but no service named NoSuchSvc is defined",
		"badFormatVersion": "Project test_assets/project/badFormatVersion/Bopmatic.yaml specifies an unsupported formatversion 0.9. The latest supported formatversion is " +
			FormatVersionCurrent + ".",
		"missingUserGroup":         "Service Greeter in Project Foo defines access for user group NoSuchUserGroup but no user group named NoSuchUserGroup is defined",
		"missingUserGroupName":     "UserGroup in Project Foo definition is missing required field name",
		"reservedUserGroupName":    "UserGroup in Project Foo is using reserved name anon_public. Please choose a different user group name.",
		"missingUserGroupType":     "UserGroup CustomerGroup definition is missing required field type",
		"missingUserGroupRef":      "Usergroup CustomerUserGroup in Project Foo is defined but no service references user access to it",
		"bogusUserGroupType":       "Unsupported user group type randomUnsupportedType defined in UserGroup Foo for Project CustomerUserGroup",
		"missingApiAssets":         "Failed to open API def assets directory for Service Greeter: stat test_assets/project/missingApiAssets/pb: no such file or directory",
		"apiAssetsNotDir":          "Failed to open API def assets directory for Service Greeter: test_assets/project/apiAssetsNotDir/pb is not a directory",
		"missingObjStoreName":      "Object Store in Project Foo definition is missing required field name",
		"missingObjStoreSvcAccess": "Object Store Uploads in Project Foo must define at least 1 service access",
		"noSuchObjSToreSvcAccess":  "Object Store Uploads in Project Foo defines access for service NoSuchSvc but no service named NoSuchSvc is defined",
	}

	for idx, tCase := range testCases {
		if !tCase.IsDir() {
			continue
		}

		projFile := filepath.Join(testCasesDir, tCase.Name(), "Bopmatic.yaml")
		proj, err := NewProject(projFile, ProjectOptValidateSiteAssets())
		expectedErrStr, hasExpectedFailure := expectedFailures[tCase.Name()]
		expectedSuccessStr, hasExpectedSuccess := expectedSuccesses[tCase.Name()]

		if !hasExpectedSuccess && !hasExpectedFailure {
			t.Errorf("Test case %v(%v) expectation missing; please update TestProject()",
				idx, tCase.Name())
		} else if hasExpectedFailure {
			if err == nil {
				t.Errorf("Test case %v(%v) succeeded but should have failed with %v",
					idx, tCase.Name(), expectedErrStr)
			} else if fmt.Sprintf("%v", err) != expectedErrStr {
				t.Errorf("Test case %v(%v) failed with %v but should have failed with %v",
					idx, tCase.Name(), err, expectedErrStr)
			}
		} else if err != nil {
			t.Errorf("Test case %v(%v) failed with %v but should have succeeded",
				idx, tCase.Name(), err)
		} else if expectedSuccessStr != proj.HashString() {
			t.Errorf("Test case %v(%v) succeeded with %v but should have succeeded with %v: %v",
				idx, tCase.Name(), proj.HashString(), expectedSuccessStr, proj.String())
		}
	}
}

func TestProjectName(t *testing.T) {
	goodNames := []string{"foo", "barreaasdfasdfasdfasdfreallylongname", "SomeNameWithCaps"}
	for _, name := range goodNames {
		isGood, err := IsGoodProjectName(name)
		if err != nil || isGood == false {
			t.Errorf("IsGoodProjectName(%v) failed but should have succeeded: %v",
				name, err)
		}
	}

	workPath, err := ioutil.TempDir(".", "test")
	if err != nil {
		t.Errorf("Failed to create tempdir: %v", err)
	}

	badNames := []string{"", ",@#$", "name.withperiod", workPath,
		"name_withunderscore"}
	for _, name := range badNames {
		isGood, err := IsGoodProjectName(name)
		if err == nil || isGood == true {
			t.Errorf("IsGoodProjectName(%v) succeeded but should have failed: %v",
				name, err)
		}
	}

	os.RemoveAll(workPath)
}

func TestProjectExport(t *testing.T) {
	var testCasesDir = filepath.Join("test_assets", "project")

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
		_, hasExpectedSuccess := expectedSuccesses[tCase.Name()]
		if !hasExpectedSuccess {
			continue
		}
		if err != nil {
			t.Errorf("Failed to instantiate project for %v: %v", tCase.Name(), err)
		}

		tmpFile, err := os.CreateTemp(filepath.Join(testCasesDir, tCase.Name()),
			"bopsdk-golang-proj-export-test-*")
		if err != nil {
			t.Errorf("Failed to create temp file: %v", err)
		}
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		err = proj.ExportToFile(tmpFile.Name())
		if err != nil {
			t.Errorf("Failed to export for test case %v: %v", tCase.Name(), err)
		}

		newProj, err := NewProject(tmpFile.Name())
		if err != nil {
			t.Errorf("Failed to re-import exported file for test case %v: %v",
				tCase.Name(), err)
		}
		if !proj.IsEqual(newProj) {
			t.Logf("Original proj: %v", proj)
			t.Logf("Exported proj: %v", newProj)

			t.Errorf("Re-imported export does not match original for test case %v",
				tCase.Name())
		}
	}
}
