package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestProject(t *testing.T) {
	var testCasesDir = filepath.Join("test_assets", "project")

	testCases, err := ioutil.ReadDir(testCasesDir)
	if err != nil {
		t.Errorf("Failed to open project test case dir: %v", err)
	}

	expectedFailures := map[string]string{
		"missingPort":        "Service Greeter definition is missing required field port",
		"missingProjectName": "Project test_assets/project/missingProjectName/Bopmatic.yaml definition is missing required field name",
		"missingSvcInApiDef": "Failed to find service Missing in pb/greeter.proto",
		"missingSvcName":     "Service 0 definition is missing required field name",
		"missingSiteAssets":  "Could not open site assets missing: open test_assets/project/missingSiteAssets/missing: no such file or directory",
		"emptySiteAssets":    "Site assets siteAssetsEmpty is missing index.html",
		"noRpcsInSvc":        "Service Greeter in pb/greeter.proto must define at least 1 RPC",
		"noSvcInApiDef":      "pb/greeter.proto does not define any services; looking for Greeter",
		"portConflict":       "Service Orders port 26001 conflicts with service Greeter",
		"missingApiDefField": "Service Greeter definition is missing required field apidef",
		"missingApiDefFile":  "Failed to open API definition for Service Greeter: open test_assets/project/missingApiDefFile/pb/greeter.proto: no such file or directory",
		"missingDbName":      "Database in Project Foo definition is missing required field name",
		"missingDbTable":     "Database Customers in Project Foo must define at least 1 table",
		"missingDbTableName": "Table in Database Customers definition is missing required field name",
		"missingDbSvcAccess": "Database Customers in Project Foo must define at least 1 service access",
		"missingProject":     "Failed to open project test_assets/project/missingProject/Bopmatic.yaml: open test_assets/project/missingProject/Bopmatic.yaml: no such file or directory",
		"corruptApiDef":      "Failed to parse pb/greeter.proto for service Greeter: found \"CORRUPT\" but expected [;]",
		"missingBuildCmd":    "Project Foo definition is missing required field buildcmd",
		"missingExecutable":  "Service Greeter definition is missing required field executable",
		"noSuchDbSvcAccess":  "Database Customers in Project Foo defines access for service NoSuchSvc but no service named NoSuchSvc is defined",
	}
	expectedSuccesses := map[string]string{
		"multiSvcMultiDb":  "qmw1EEvv6ElphVdobA1aowMC6OPAPe-NRCEc-Tz1ofE=",
		"multiSvcMultiRpc": "bqhsBRoQ8DkBq3J4PcTJ-yS0XlQnN4NTEt_ljAI84y0=",
		"multiSvcOneRpc":   "L3AMBkLJ9mcGOS_EdLjhh6j1eCbKGg7-Xgc2R-T3KGU=",
		"oneSvcOneDb":      "xue3VPpRy1iLEhgm-lHEFjUsGdvq82rtV2EurLKCTOE=",
		"oneSvcMultiRpc":   "grCSNtGtcIU4wgCzc7EWKE7chxlzvq0rlJzPQqu1lfo=",
		"oneSvcOneRpc":     "oAantAxe_EmpZa9FShPEiML968ODqs84mbFpH9ErhaQ=",
		"staticOnly":       "j-yzmjEWgOi4ybhNGY5j1DAKFoALFt62HFU20CjEfWI=",
		"svcsOnly":         "m1Wo0vRbEOE8KMg3kzvdxIQIHp6-81ugYYxdXrLwAbQ=",
	}

	for idx, tCase := range testCases {
		if !tCase.IsDir() {
			continue
		}

		projFile := filepath.Join(testCasesDir, tCase.Name(), "Bopmatic.yaml")
		proj, err := NewProject(projFile)
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
