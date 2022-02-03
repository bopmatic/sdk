package golang

import (
	"fmt"
	"io/ioutil"
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
		"missingApiDef":      "Failed to open API definition for Service Greeter: open test_assets/project/missingApiDef/pb/greeter.proto: no such file or directory",
		"missingProject":     "Failed to open project test_assets/project/missingProject/Bopmatic.yaml: open test_assets/project/missingProject/Bopmatic.yaml: no such file or directory",
		"missingImportPath":  "API def(pb/greeter.proto) for service Greeter missing go_package import",
		"corruptApiDef":      "Failed to parse pb/greeter.proto for service Greeter: found \"CORRUPT\" but expected [;]",
	}
	expectedSuccesses := map[string]string{
		"multiSvcMultiRpc": "jNx753jfnTZPRRaBHxiIXYqwI7PdFRD2O1m2GM1hgRk=",
		"multiSvcOneRpc":   "YPG_L5rFNTFfqcYr0yFgWc1upxu5K-mYP8aRpBL31Wc=",
		"oneSvcMultiRpc":   "3RBV3sSpR9V1yDF6ydLrq4A1vAL32O-y08Cn3-1XvGc=",
		"oneSvcOneRpc":     "6y-0JwSX204i2Qb6ejU0Ld0cjeBnOXtob6LhnfWm2qk=",
		"staticOnly":       "fhioOjBw0HMxpHYiRl8VIfvctOP_Vl6fE3iIlzwwCg4=",
		"svcsOnly":         "vNCrurr14X45c8ziRCA1eXfmC5tosuX-XrjJ8yD_m0Q=",
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
