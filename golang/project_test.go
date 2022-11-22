package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var expectedSuccesses = map[string]string{
	"multiSvcMultiDb":         "X2konu6DisZiYulAskdwBYp0mcsna_18zm42iyo4_Rw=",
	"multiSvcMultiRpc":        "i9_h-2oZSBg8MternhT2i4CmCOg1feQ8dI0ovRJfmdM=",
	"multiSvcOneRpc":          "7MumxaAkGhjDGSp11bi40STguXXpLvlIzrJ7hGqJfdw=",
	"oneSvcOneDb":             "SsIAGB3C5Q8sP4RSiMlLfedW2zg70dVQKoyHA3qTq5k=",
	"oneSvcMultiRpc":          "SqkzRfV5o3Z82onxEXitPLAQ3-gaDy5WuW8RTHOU9DU=",
	"oneSvcOneRpc":            "FeI6TqOQSFwO87V5W74KumLcOIXV79ACUHeEeZkjWyI=",
	"staticOnly":              "eTqZYv3IzHa6hJIl8UrRWuHNxXQg54Z82xMh4yZYnyc=",
	"svcsOnly":                "8z9jNFbnaQnm57LouqpRhPxIgTE0UmH8H29rLlI0KPI=",
	"oneSvcOneUserGroup":      "NPTAlcTQ_CIb8gKhuYADXNcRph6Ni-lxzjhJUiUkMiI=",
	"multiSvcMultiUserGroup":  "vByOh78WqLq7YwSkuJ6ZC4v4B-CkFDqHYv3PD8ii-P8=",
	"oneSvcWithApiAssets":     "FtYzjgXGDGL7OqG9g8fUmyadFvNn-b_OCcJLUSbp7Jg=",
	"oneSvcWithRuntimeConfig": "hzh9-H_AQFurjZ3hut0P-RuhHQxTHRccOsZoe6_4r5g=",
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
		"missingUserGroup":      "Service Greeter in Project Foo defines access for user group NoSuchUserGroup but no user group named NoSuchUserGroup is defined",
		"missingUserGroupName":  "UserGroup in Project Foo definition is missing required field name",
		"reservedUserGroupName": "UserGroup in Project Foo is using reserved name anon_public. Please choose a different user group name.",
		"missingUserGroupType":  "UserGroup CustomerGroup definition is missing required field type",
		"missingUserGroupRef":   "Usergroup CustomerUserGroup in Project Foo is defined but no service references user access to it",
		"bogusUserGroupType":    "Unsupported user group type randomUnsupportedType defined in UserGroup Foo for Project CustomerUserGroup",
		"missingApiAssets":      "Failed to open API def assets directory for Service Greeter: stat test_assets/project/missingApiAssets/pb: no such file or directory",
		"apiAssetsNotDir":       "Failed to open API def assets directory for Service Greeter: test_assets/project/apiAssetsNotDir/pb is not a directory",
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
