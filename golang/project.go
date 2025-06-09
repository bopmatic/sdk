/* Copyright © 2022-2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
	"github.com/bopmatic/sdk/golang/pb"
	"github.com/bopmatic/sdk/golang/util"

	"github.com/yoheimuta/go-protoparser/v4"

	"gopkg.in/yaml.v2"
)

const DefaultProjectFilename = "Bopmatic.yaml"
const DefaultArtifactDir = ".bopmatic"
const PackagesSubdir = "pkgs"

// DatabaseTable is a representation of persistent state organized by a set of
// related key/value pairs.
type DatabaseTable struct {
	Name        string `yaml:"name"`
	Description string `yaml:"desc,omitempty"`
}

// Database is a group of DatabaseTables accessible by one or more services.
// Running services may access databases by employing the DAPR state store APIs
// via the SDK of the user's preferred programming language. See
// https://docs.dapr.io/developing-applications/sdks/ for further detail.
type Database struct {
	Name        string          `yaml:"name"`
	Description string          `yaml:"desc,omitempty"`
	Tables      []DatabaseTable `yaml:"tables,flow"`
	Services    []string        `yaml:"services_access,flow"`
}

// ObjectStore is an object storage container accessible by one or more
// services. Running services may access object stores by employing the
// S3 APIs in the AWS SDK of the user's preferred programming language.
type ObjectStore struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"desc,omitempty"`
	Services    []string `yaml:"services_access,flow"`
}

// Service is a representation of an individual service defined within
// a Bopmatic project. This includes its name, a short description, a link
// to the set of APIs it defines, and the port it should run on. Fields marked
// with yaml are parsed from Bopmatic.yaml while the remaining fields are
// derived once the project is parsed (e.g. when NewProject() returns
// successfully).
// Currently the only supported ApiDefinition formats are protobuf/GRPC and
// OpenAPI(swagger)
const (
	UserAccessAnonPublic = "anon_public"
	UserAccessNone       = "none"
)

type ApiType int

const (
	ApiTypeUnknown ApiType = iota
	ApiTypeGRPC
	ApiTypeOpenAPI
)

type Service struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"desc,omitempty"`
	ApiDefinition string `yaml:"apidef"`
	ApiDefAssets  string `yaml:"apidef_assets,omitempty"`
	Port          uint64 `yaml:"port"`
	Executable    string `yaml:"executable"`
	ExecAssets    string `yaml:"executable_assets,omitempty"`
	// The user access refers to the name of user group which is allowed to
	// invoke a service's APIs. The user group should either also be defined
	// in the same Bopmatic.yaml or refer to one of the following two built-in
	// user accesses:
	//     1. anon_public - Access to the APIs is public and does not require
	//                      any authentication. e.g. a google search
	//     2. none        - No users are allowed to access the service directly.
	//                      e.g. an internal sales analytics service that is
	//                      only invoked by other services and not by end users
	//                      directly. This is the default.
	UserAccess string `yaml:"user_access,omitempty"`

	rpcs    []string
	apiType ApiType
}

func (svc *Service) GetRpcs() []string {
	return svc.rpcs
}

func (svc *Service) GetApiType() ApiType {
	return svc.apiType
}

// ProjectDesc see Project for a complete description
type ProjectDesc struct {
	Name          string        `yaml:"name"`
	Id            string        `yaml:"id"`
	Description   string        `yaml:"desc,omitempty"`
	Services      []Service     `yaml:"services,omitempty,flow"`
	Databases     []Database    `yaml:"databases,omitempty,flow"`
	ObjectStores  []ObjectStore `yaml:"object_stores,omitempty,flow"`
	UserGroups    []UserGroup   `yaml:"usergroups,omitempty,flow"`
	SiteAssets    string        `yaml:"sitedir,omitempty"`
	RuntimeConfig string        `yaml:"runtime_config,omitempty"`
	BuildCmd      string        `yaml:"buildcmd"`

	root     string
	projFile string
}

func (desc *ProjectDesc) GetRoot() string {
	return desc.root
}

const (
	UserGroupTypePublic  = "public"
	UserGroupTypePrivate = "private"
)

type UserGroup struct {
	Name        string `yaml:"name"`
	Description string `yaml:"desc,omitempty"`

	// 2 usergroup types are currently defined:
	//
	// 1. public - membership to the group is available to the public. users
	//             can sign themselves up without any interaction from service
	//             administrators. e.g. a new customer can signup
	//             at will.
	// 2. private - membership to the group is controlled by service
	//             administrators. Users may not sign themselves up. e.g.
	//             a new employee has to be onboarded by HR.
	Type string `yaml:"type"`
}

// Project is a full representation of a Bopmatic project, which defines
// a set of services to publish, their RPCs, their TCP ports to run on,
// and a collection of static website assets to publish. FormatVersion
// describes the set of supported and expected fields within a Bopmatic.yaml
// project description. Currently this is expected to be "1.1". Older "1.0"
// existing templates are also supported. This field is intended to be used
// for backward compatibility. Fields marked with yaml are parsed from
// Bopmatic.yaml while the remaining fields are derived once the project is
// parsed (e.g. when NewProject() returns successfully).
//
// Differences between formatversion 1.0 and 1.1:
//
//	1.1 introduces usergroups & relationships between usersgroups and
//	    services for access control purposes. In version 1.0 as usergroups
//	    had not yet been introduced, all defined services were implictly
//	    vended with anonymous public access by default. In version 1.1, the
//	    default access is instead none.
const (
	FormatVersion1_0     = "1.0"
	FormatVersion1_1     = "1.1"
	FormatVersionCurrent = FormatVersion1_1
)

type Project struct {
	FormatVersion string      `yaml:"formatversion"`
	Desc          ProjectDesc `yaml:"project"`
}

func IsGoodProjectName(projectName string) (bool, error) {
	if projectName == "" {
		return false, fmt.Errorf("Project name must be a non-empty string")
	}
	projectName = strings.ToLower(projectName)
	if strings.ContainsRune(projectName, '.') {
		return false,
			fmt.Errorf("Project names containing a '.' are not yet supported")
	}
	if strings.ContainsRune(projectName, '_') {
		return false,
			fmt.Errorf("Project names containing an '_' are not yet supported")
	}
	url, err := url.ParseRequestURI("https://" + projectName + ".bopmatic.com")
	if err != nil {
		return false, fmt.Errorf("%v.bopmatic.com is not a valid endpoint",
			projectName)
	}
	_, err = net.LookupIP(url.Host)
	if err == nil {
		return false, fmt.Errorf("%v is already taken by another Bopmatic customer", url.Host)
	}

	_, err = os.Stat(projectName)
	if err == nil {
		return false, fmt.Errorf("%v already exists in your current directory",
			projectName)
	} // else

	return true, nil
}

// String converts a Project into a human-friendly printable string
func (proj *Project) String() string {
	var sb strings.Builder

	sb.WriteString("Project:\n")
	sb.WriteString(fmt.Sprintf("\tFormat: %v\n", proj.FormatVersion))
	sb.WriteString(fmt.Sprintf("\tName: %v\n", proj.Desc.Name))
	if proj.Desc.Id != "" {
		sb.WriteString(fmt.Sprintf("\tId: %v\n", proj.Desc.Id))
	}
	sb.WriteString(fmt.Sprintf("\tSiteAssets: %v\n", proj.Desc.SiteAssets))
	if proj.Desc.RuntimeConfig != "" {
		sb.WriteString(fmt.Sprintf("\tRuntimeConfig: %v\n",
			proj.Desc.RuntimeConfig))
	}
	if proj.Desc.BuildCmd != "" {
		sb.WriteString(fmt.Sprintf("\tBuildCmd: %v\n", proj.Desc.BuildCmd))
	}
	sb.WriteString(fmt.Sprintf("\tRoot: %v\n", proj.Desc.root))
	if proj.Desc.Description != "" {
		sb.WriteString(fmt.Sprintf("\tDescription: %v\n", proj.Desc.Description))
	}
	sb.WriteString(fmt.Sprintf("\tServices: %v\n", len(proj.Desc.Services)))
	for idx, svc := range proj.Desc.Services {
		sb.WriteString(fmt.Sprintf("\tService[%v]:\n", idx))
		sb.WriteString(fmt.Sprintf("\t\tName: %v\n", svc.Name))
		if svc.Description != "" {
			sb.WriteString(fmt.Sprintf("\t\tDescription: %v\n", svc.Description))
		}
		sb.WriteString(fmt.Sprintf("\t\tApiDef: %v\n", svc.ApiDefinition))
		if svc.ApiDefAssets != "" {
			sb.WriteString(fmt.Sprintf("\t\tApiDefAssets: %v\n",
				svc.ApiDefAssets))
		}
		sb.WriteString(fmt.Sprintf("\t\tUserAccess: %v\n", svc.UserAccess))
		sb.WriteString(fmt.Sprintf("\t\tPort: %v\n", svc.Port))
		sb.WriteString(fmt.Sprintf("\t\tExecutable: %v\n", svc.Executable))
		if svc.ExecAssets != "" {
			sb.WriteString(fmt.Sprintf("\t\tExecAssets: %v\n", svc.ExecAssets))
		}
		sb.WriteString(fmt.Sprintf("\t\tRpcs: %v\n", len(svc.rpcs)))
		for idx2, funcName := range svc.rpcs {
			sb.WriteString(fmt.Sprintf("\t\tRpc[%v]: %v\n", idx2, funcName))
		}
	}
	sb.WriteString(fmt.Sprintf("\tDatabases: %v\n", len(proj.Desc.Databases)))
	for idx, db := range proj.Desc.Databases {
		sb.WriteString(fmt.Sprintf("\tDatabase[%v]:\n", idx))
		sb.WriteString(fmt.Sprintf("\t\tName: %v\n", db.Name))
		if db.Description != "" {
			sb.WriteString(fmt.Sprintf("\t\tDescription: %v\n", db.Description))
		}
		sb.WriteString(fmt.Sprintf("\t\tTables: %v\n", len(db.Tables)))
		for idx2, tbl := range db.Tables {
			sb.WriteString(fmt.Sprintf("\t\tTable[%v]:\n", idx2))
			sb.WriteString(fmt.Sprintf("\t\t\tName: %v\n", tbl.Name))
			if tbl.Description != "" {
				sb.WriteString(fmt.Sprintf("\t\t\tDescription: %v\n",
					tbl.Description))
			}
		}
		sb.WriteString(fmt.Sprintf("\t\tServicesAccess: %v\n", len(db.Services)))
		for idx2, svc := range db.Services {
			sb.WriteString(fmt.Sprintf("\t\tService[%v]: %v\n", idx2, svc))
		}
	}
	sb.WriteString(fmt.Sprintf("\tObject Stores: %v\n", len(proj.Desc.ObjectStores)))
	for idx, objStore := range proj.Desc.ObjectStores {
		sb.WriteString(fmt.Sprintf("\tObject Store[%v]:\n", idx))
		sb.WriteString(fmt.Sprintf("\t\tName: %v\n", objStore.Name))
		if objStore.Description != "" {
			sb.WriteString(fmt.Sprintf("\t\tDescription: %v\n",
				objStore.Description))
		}
		sb.WriteString(fmt.Sprintf("\t\tServicesAccess: %v\n",
			len(objStore.Services)))
		for idx2, svc := range objStore.Services {
			sb.WriteString(fmt.Sprintf("\t\tService[%v]: %v\n", idx2, svc))
		}
	}
	sb.WriteString(fmt.Sprintf("\tUserGroups: %v\n", len(proj.Desc.UserGroups)))
	for idx, ug := range proj.Desc.UserGroups {
		sb.WriteString(fmt.Sprintf("\tUserGroup[%v]:\n", idx))
		sb.WriteString(fmt.Sprintf("\t\tName: %v\n", ug.Name))
		if ug.Description != "" {
			sb.WriteString(fmt.Sprintf("\t\tDescription: %v\n", ug.Description))
		}
		sb.WriteString(fmt.Sprintf("\t\tType: %v\n", ug.Type))
	}

	return sb.String()
}

func isPrimitiveType(kind reflect.Kind) bool {
	ret := false

	switch kind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		ret = true
	}

	return ret
}

func yamlEncodeToString(input interface{}, indent uint, skipFirstIndent bool) string {
	v := reflect.ValueOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var sb strings.Builder

	if isPrimitiveType(v.Kind()) {
		if v.Kind() == reflect.String {
			sb.WriteString(fmt.Sprintf("\"%v\"\n", v.Interface()))
		} else {
			sb.WriteString(fmt.Sprintf("%v\n", v.Interface()))
		}

		return sb.String()
	}

	if v.Kind() == reflect.Slice {
		len := v.Len()
		for ii := 0; ii < len; ii++ {
			for ii := uint(0); ii < indent-2; ii++ {
				sb.WriteString(" ")
			}

			sb.WriteString("- ")
			element := v.Index(ii)
			sb.WriteString(yamlEncodeToString(element.Interface(), indent, true))
		}

		return sb.String()
	} else if v.Kind() != reflect.Struct {
		panic(fmt.Sprintf("yamlEncodeToString() unsupported type: %v",
			v.Kind().String()))
	}

	t := v.Type()

	for ii := 0; ii < v.NumField(); ii++ {
		field := t.Field(ii)
		if !field.IsExported() {
			continue
		}
		tag := field.Tag.Get("yaml")
		value := v.Field(ii)
		omitEmpty := false
		if tag == "" {
			tag = field.Name
		} else {
			omitEmpty = strings.Contains(tag, "omitempty")
			tag = strings.Split(tag, ",")[0]
		}

		if value.IsZero() && omitEmpty {
			continue
		}

		if !skipFirstIndent || ii > 0 {
			for jj := uint(0); jj < indent; jj++ {
				sb.WriteString(" ")
			}
		}

		if isPrimitiveType(value.Kind()) {
			sb.WriteString(fmt.Sprintf("%v: %v", tag,
				yamlEncodeToString(value.Interface(), 0, false)))
		} else {
			sb.WriteString(fmt.Sprintf("%v:\n%v", tag,
				yamlEncodeToString(value.Interface(), indent+2, false)))
		}
	}

	return sb.String()
}

// YAMLString converts a Project into a human-friendly printable YAML string
// This is a replacement for yaml.Encoder.Encode which produces a much more
// compact strict that isn't as	easy to read
func (proj *Project) YAMLString() string {
	return yamlEncodeToString(proj, 0, false)
}

// String converts a Project into shorthand string suitable as a GUID
func (proj *Project) HashString() string {
	hasher := sha256.New()
	hasher.Write([]byte(proj.String()))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func parseProject(projFile string) (*Project, error) {
	file, err := os.Open(projFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open project %v: %w", projFile, err)
	}
	defer file.Close()

	var proj Project
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&proj)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse project %v: %w", projFile, err)
	}

	if projFile[0] != os.PathSeparator {
		proj.Desc.root = path.Dir(projFile)
	} else {
		proj.Desc.root = path.Dir(projFile)
	}
	proj.Desc.projFile = projFile

	return &proj, nil
}

func (svc *Service) populateRpcsProtobuf(svcName string,
	protoFileName string,
	protoFileInput io.Reader) error {

	p, err := protoparser.Parse(protoFileInput)
	if err != nil {
		return fmt.Errorf("Failed to parse %v for service %v: %w",
			protoFileName, svcName, err)
	}

	up, err := protoparser.UnorderedInterpret(p)
	if err != nil {
		return fmt.Errorf("Failed to interpret %v for service %v: %w",
			protoFileName, svcName, err)
	}

	if len(up.ProtoBody.Services) == 0 {
		return fmt.Errorf("%v does not define any services; looking for %v",
			protoFileName, svcName)
	}

	foundSvc := false
	for _, upSvc := range up.ProtoBody.Services {
		if upSvc.ServiceName == svcName {
			foundSvc = true
			for _, rpc := range upSvc.ServiceBody.RPCs {
				svc.rpcs = append(svc.rpcs, rpc.RPCName)
			}
		}
	}

	if !foundSvc {
		return fmt.Errorf("Failed to find service %v in %v", svcName,
			protoFileName)
	}

	if len(svc.rpcs) == 0 {
		return fmt.Errorf("Service %v in %v must define at least 1 RPC", svcName,
			protoFileName)
	}

	return nil
}

func (svc *Service) populateRpcsOpenAPI(svcName, specFileName string,
	specFileInput io.Reader) error {

	data, err := io.ReadAll(specFileInput)
	if err != nil {
		return fmt.Errorf("Failed to read %v for service %v: %w", specFileName,
			svcName, err)
	}

	// Minimal representation of OpenAPI paths we care about
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string   `json:"operationId"`
			Tags        []string `json:"tags"`
		} `json:"paths"`
	}
	err = json.Unmarshal(data, &spec)
	if err != nil {
		return fmt.Errorf("Failed to parse %v for service %v: %w", specFileName,
			svcName, err)
	}

	if len(spec.Paths) == 0 {
		return fmt.Errorf("%v does not define any paths; looking for service %v",
			specFileName, svcName)
	}

	for pathStr, ops := range spec.Paths {
		for method, op := range ops {
			// Determine service grouping
			serviceName := svcName
			if len(op.Tags) > 0 && op.Tags[0] != "" {
				serviceName = op.Tags[0]
			} else {
				parts := strings.Split(strings.Trim(pathStr, "/"), "/")
				if len(parts) > 0 && parts[0] != "" {
					serviceName = parts[0]
				}
			}
			if serviceName != svcName {
				continue
			}

			// Determine RPC name
			rpcName := op.OperationID
			if rpcName == "" {
				base := path.Base(pathStr)
				if base == "" || base == "/" {
					base = svcName
				}
				rpcName = strings.Title(strings.ToLower(method)) +
					strings.Title(base)
			}

			svc.rpcs = append(svc.rpcs, rpcName)
		}
	}

	if len(svc.rpcs) == 0 {
		return fmt.Errorf("Service %v in %v must define at least 1 RPC",
			svcName, specFileName)
	}

	return nil
}

func (svc *Service) populateRpcs(svcName string,
	protoFileName string,
	protoFileInput io.Reader) error {

	if strings.HasSuffix(protoFileName, ".proto") {
		svc.apiType = ApiTypeGRPC
		return svc.populateRpcsProtobuf(svcName, protoFileName, protoFileInput)
	} else if strings.HasSuffix(protoFileName, ".json") {
		svc.apiType = ApiTypeOpenAPI
		return svc.populateRpcsOpenAPI(svcName, protoFileName, protoFileInput)
	} // else

	return fmt.Errorf("Unrecognized Api definition format in %v", protoFileName)

}

func (proj *Project) validateProject(projFile string, validateSiteAssets bool) error {
	const missingFieldFmt = "%v %v definition is missing required field %v"

	if proj.FormatVersion == "" {
		return fmt.Errorf(missingFieldFmt, "Project", projFile, "formatversion")
	}
	if proj.FormatVersion != FormatVersion1_0 &&
		proj.FormatVersion != FormatVersion1_1 {
		return fmt.Errorf("Project %v specifies an unsupported formatversion %v. The latest supported formatversion is %v.",
			projFile, proj.FormatVersion, FormatVersionCurrent)
	}
	if proj.Desc.Name == "" {
		return fmt.Errorf(missingFieldFmt, "Project", projFile, "name")
	}

	if validateSiteAssets && proj.Desc.SiteAssets != "" {
		siteAssetsPath := filepath.Join(proj.Desc.root, proj.Desc.SiteAssets)
		assetEntries, err := ioutil.ReadDir(siteAssetsPath)
		if err != nil {
			return fmt.Errorf("Could not open site assets %v: %w", proj.Desc.SiteAssets, err)
		}
		foundIndex := false
		for _, entry := range assetEntries {
			if entry.Name() == "index.html" {
				foundIndex = true
				break
			}
		}
		if !foundIndex {
			return fmt.Errorf("Site assets %v is missing index.html", proj.Desc.SiteAssets)
		}
		if proj.Desc.RuntimeConfig != "" {
			runtimeConfigPath := filepath.Join(siteAssetsPath,
				proj.Desc.RuntimeConfig)
			_, err := ioutil.ReadFile(runtimeConfigPath)
			if err != nil {
				return fmt.Errorf("Could not open site assets(%v)' runtime config(%v): %v",
					proj.Desc.SiteAssets, proj.Desc.RuntimeConfig, err)
			}
		}
	}

	if len(proj.Desc.Services) > 0 && proj.Desc.BuildCmd == "" {
		return fmt.Errorf(missingFieldFmt, "Project", proj.Desc.Name, "buildcmd")
	}

	var usedPorts []uint64
	for idx, _ := range proj.Desc.Services {
		svc := &proj.Desc.Services[idx]
		if svc.Name == "" {
			return fmt.Errorf(missingFieldFmt, "Service", idx, "name")
		}
		if svc.ApiDefinition == "" {
			return fmt.Errorf(missingFieldFmt, "Service", svc.Name, "apidef")
		}
		if svc.Port == 0 {
			return fmt.Errorf(missingFieldFmt, "Service", svc.Name, "port")
		}
		if svc.Executable == "" {
			return fmt.Errorf(missingFieldFmt, "Service", svc.Name, "executable")
		}
		for portIdx, port := range usedPorts {
			if svc.Port == port {
				return fmt.Errorf("Service %v port %v conflicts with service %v",
					svc.Name, port, proj.Desc.Services[portIdx].Name)
			}
		}

		if svc.UserAccess == "" {
			if proj.FormatVersion == FormatVersion1_0 {
				svc.UserAccess = UserAccessAnonPublic
			} else {
				svc.UserAccess = UserAccessNone
			}
		} else if svc.UserAccess != UserAccessAnonPublic &&
			svc.UserAccess != UserAccessNone {
			var found bool
			for _, ug := range proj.Desc.UserGroups {
				found = false
				if ug.Name == svc.UserAccess {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Service %v in Project %v defines access for user group %v but no user group named %v is defined",
					svc.Name, proj.Desc.Name, svc.UserAccess, svc.UserAccess)
			}
		}

		usedPorts = append(usedPorts, svc.Port)
		apiDefPath := filepath.Join(proj.Desc.root, svc.ApiDefinition)
		file, err := os.Open(apiDefPath)
		if err != nil {
			return fmt.Errorf("Failed to open API definition for Service %v: %w", svc.Name, err)
		}
		defer file.Close()

		apiDefAssetsPath := filepath.Join(proj.Desc.root, svc.ApiDefAssets)
		apiDefAssetsStat, err := os.Stat(apiDefAssetsPath)
		if err != nil {
			return fmt.Errorf("Failed to open API def assets directory for Service %v: %w",
				svc.Name, err)
		}
		if !apiDefAssetsStat.IsDir() {
			return fmt.Errorf("Failed to open API def assets directory for Service %v: %v is not a directory",
				svc.Name, apiDefAssetsPath)
		}

		err = svc.populateRpcs(svc.Name, svc.ApiDefinition, file)
		if err != nil {
			return err
		}
	}

	for idx, _ := range proj.Desc.Databases {
		db := &proj.Desc.Databases[idx]

		if db.Name == "" {
			return fmt.Errorf(missingFieldFmt, "Database in Project",
				proj.Desc.Name, "name")
		}
		if len(db.Tables) == 0 {
			return fmt.Errorf("Database %v in Project %v must define at least 1 table",
				db.Name, proj.Desc.Name)
		}

		for tidx, _ := range db.Tables {
			tbl := &db.Tables[tidx]

			if tbl.Name == "" {
				return fmt.Errorf(missingFieldFmt, "Table in Database",
					db.Name, "name")
			}
		}

		if len(db.Services) == 0 {
			return fmt.Errorf("Database %v in Project %v must define at least 1 service access",
				db.Name, proj.Desc.Name)
		}

		var found bool
		for _, svcAccess := range db.Services {
			found = false
			for sidx, _ := range proj.Desc.Services {
				svc := &proj.Desc.Services[sidx]
				if svcAccess == svc.Name {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Database %v in Project %v defines access for service %v but no service named %v is defined",
					db.Name, proj.Desc.Name, svcAccess, svcAccess)
			}
		}
	}

	for idx, _ := range proj.Desc.ObjectStores {
		objStore := &proj.Desc.ObjectStores[idx]

		if objStore.Name == "" {
			return fmt.Errorf(missingFieldFmt, "Object Store in Project",
				proj.Desc.Name, "name")
		}

		if len(objStore.Services) == 0 {
			return fmt.Errorf("Object Store %v in Project %v must define at least 1 service access",
				objStore.Name, proj.Desc.Name)
		}

		var found bool
		for _, svcAccess := range objStore.Services {
			found = false
			for sidx, _ := range proj.Desc.Services {
				svc := &proj.Desc.Services[sidx]
				if svcAccess == svc.Name {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Object Store %v in Project %v defines access for service %v but no service named %v is defined",
					objStore.Name, proj.Desc.Name, svcAccess, svcAccess)
			}
		}
	}

	for idx, _ := range proj.Desc.UserGroups {
		ug := &proj.Desc.UserGroups[idx]

		if ug.Name == "" {
			return fmt.Errorf(missingFieldFmt, "UserGroup in Project",
				proj.Desc.Name, "name")
		} else if ug.Name == UserAccessAnonPublic || ug.Name == UserAccessNone {
			return fmt.Errorf("UserGroup in Project %v is using reserved name %v. Please choose a different user group name.",
				proj.Desc.Name, ug.Name)
		}
		if ug.Type == "" {
			return fmt.Errorf(missingFieldFmt, "UserGroup", ug.Name, "type")
		}
		if ug.Type != UserGroupTypePublic {
			return fmt.Errorf("Unsupported user group type %v defined in UserGroup %v for Project %v",
				ug.Type, proj.Desc.Name, ug.Name)
		}

		var found bool
		for idx, _ := range proj.Desc.Services {
			svc := &proj.Desc.Services[idx]
			if svc.UserAccess == ug.Name {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Usergroup %v in Project %v is defined but no service references user access to it",
				ug.Name, proj.Desc.Name)
		}
	}

	return nil
}

// Compile the project
func (proj *Project) Build(stdOut io.Writer, stdErr io.Writer) error {
	if proj.Desc.BuildCmd == "" {
		return nil
	}

	curWd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(proj.Desc.root)
	if err != nil {
		return err
	}

	err = util.RunContainerCommand(context.Background(),
		[]string{proj.Desc.BuildCmd}, stdOut, stdErr)
	if err != nil {
		return err
	}

	_ = os.Chdir(curWd)

	return nil
}

// NewPackage collects all of the build & project artifacts required in order
// to submit to Bopmatic's ServiceRunner for deployment. Upon success a Package
// instance is returned.
func (proj *Project) NewPackageCreate(pkgName string, stdOut io.Writer,
	stdErr io.Writer, opts ...PkgOption) (*Package, error) {

	return NewPackage(pkgName, proj, stdOut, stdErr, opts...)
}

func (proj *Project) NewPackageExisting(pkgId string) (*Package, error) {
	return NewPackageFromExisting(proj, pkgId)
}

// RemoveStalePackages deletes any previously created project packages
func (proj *Project) RemoveStalePackages() error {
	return RemoveStalePackages(proj)
}

// ExportToFile export Project to a new Bopmatic.yaml file
func (proj *Project) ExportToFile(projFile string) error {
	tmpVer := proj.FormatVersion
	proj.FormatVersion = FormatVersionCurrent
	defer func() { proj.FormatVersion = tmpVer }()

	err := os.WriteFile(projFile, []byte(proj.YAMLString()), 0644)
	if err != nil {
		return fmt.Errorf("Failed to write %v: %w", projFile, err)
	}

	return nil
}

// IsEqual compares project with cmpProj to determine whether they are equivalent
func (proj *Project) IsEqual(cmpProj *Project) bool {
	if proj.Desc.Name != cmpProj.Desc.Name {
		return false
	}
	if proj.Desc.Id != cmpProj.Desc.Id {
		return false
	}
	if proj.Desc.Description != cmpProj.Desc.Description {
		return false
	}
	if proj.Desc.SiteAssets != cmpProj.Desc.SiteAssets {
		return false
	}
	if proj.Desc.RuntimeConfig != cmpProj.Desc.RuntimeConfig {
		return false
	}
	if proj.Desc.BuildCmd != cmpProj.Desc.BuildCmd {
		return false
	}
	if len(proj.Desc.Services) != len(cmpProj.Desc.Services) {
		return false
	}
	if len(proj.Desc.Databases) != len(cmpProj.Desc.Databases) {
		return false
	}
	if len(proj.Desc.ObjectStores) != len(cmpProj.Desc.ObjectStores) {
		return false
	}
	if len(proj.Desc.UserGroups) != len(cmpProj.Desc.UserGroups) {
		return false
	}
	var found bool
	for _, svc := range proj.Desc.Services {
		found = false
		for _, cmpSvc := range cmpProj.Desc.Services {
			if svc.Name != cmpSvc.Name {
				continue
			}

			found = true
			if svc.Description != cmpSvc.Description {
				return false
			}
			if svc.ApiDefinition != cmpSvc.ApiDefinition {
				return false
			}
			if svc.ApiDefAssets != cmpSvc.ApiDefAssets {
				return false
			}
			if svc.Port != cmpSvc.Port {
				return false
			}
			if svc.Executable != cmpSvc.Executable {
				return false
			}
			if svc.ExecAssets != cmpSvc.ExecAssets {
				return false
			}
			if svc.UserAccess != cmpSvc.UserAccess {
				return false
			}

			break
		}

		if !found {
			return false
		}
	}

	for _, db := range proj.Desc.Databases {
		found = false
		for _, cmpDb := range cmpProj.Desc.Databases {
			if db.Name != cmpDb.Name {
				continue
			}

			found = true
			if db.Description != cmpDb.Description {
				return false
			}
			if len(db.Tables) != len(cmpDb.Tables) {
				return false
			}
			if len(db.Services) != len(cmpDb.Services) {
				return false
			}

			var found2 bool
			for _, tbl := range db.Tables {
				found2 = false
				for _, cmpTbl := range cmpDb.Tables {
					if tbl.Name != cmpTbl.Name {
						continue
					}
					found2 = true
					if tbl.Description != cmpTbl.Description {
						return false
					}
					break
				}

				if !found2 {
					return false
				}
			}
			for _, svc := range db.Services {
				found2 = false
				for _, cmpSvc := range cmpDb.Services {
					if svc != cmpSvc {
						continue
					}
					found2 = true
					break
				}

				if !found2 {
					return false
				}
			}

			break
		}

		if !found {
			return false
		}
	}

	for _, objStore := range proj.Desc.ObjectStores {
		found = false
		for _, cmpObjStore := range cmpProj.Desc.ObjectStores {
			if objStore.Name != cmpObjStore.Name {
				continue
			}

			found = true
			if objStore.Description != cmpObjStore.Description {
				return false
			}
			if len(objStore.Services) != len(cmpObjStore.Services) {
				return false
			}

			var found2 bool
			for _, svc := range objStore.Services {
				found2 = false
				for _, cmpSvc := range cmpObjStore.Services {
					if svc != cmpSvc {
						continue
					}
					found2 = true
					break
				}

				if !found2 {
					return false
				}
			}

			break
		}

		if !found {
			return false
		}
	}

	for _, ug := range proj.Desc.UserGroups {
		found = false
		for _, cmpUg := range cmpProj.Desc.UserGroups {
			if ug.Name != cmpUg.Name {
				continue
			}

			found = true
			if ug.Description != cmpUg.Description {
				return false
			}
			if ug.Type != cmpUg.Type {
				return false
			}

			break
		}

		if !found {
			return false
		}
	}

	return true
}

type projectOptions struct {
	validateSiteAssets bool
}

type ProjectOption func(*projectOptions) *projectOptions

// ProjectOptValidateSiteAssets()instructs NewProject() to validate the existence and
// content (e.g. whether index.html is present) of a project's specified
// sitedir. The default is false because projects may create sitedir at build
// time so it may not be present prior to a build
func ProjectOptValidateSiteAssets() ProjectOption {
	return func(po *projectOptions) *projectOptions {
		po.validateSiteAssets = true

		return po
	}
}

func fillProjectOptions(opts ...ProjectOption) *projectOptions {
	options := &projectOptions{
		validateSiteAssets: false,
	}
	for _, optApplyFunc := range opts {
		if optApplyFunc != nil {
			options = optApplyFunc(options)
		}
	}

	return options
}

func ListProjects(opts ...DeployOption) ([]string, error) {

	deployOpts := fillDeployOptions(opts...)

	listProjectsReq := &models.ListDeploymentsRequest{}

	// default endpoint is inferred from host field in sr.bopmatic.json
	listProjectsParams := service_runner.NewListProjectsParams().
		WithBody(listProjectsReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListProjectsOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListProjects(listProjectsParams,
			deployOpts)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return make([]string, 0), fmt.Errorf("Client/HTTP failure: %v", err)
	}
	listReply := resp.GetPayload()
	if listReply.Result.Status != nil &&
		*listReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return make([]string, 0),
			fmt.Errorf("ListProjects failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	return listReply.Ids, nil
}

func (proj *Project) Unregister(opts ...DeployOption) error {
	return UnregisterProject(proj.Desc.Id, opts...)
}

func (proj *Project) Deactivate(envId string,
	opts ...DeployOption) (string, error) {

	return DeactivateProject(proj.Desc.Id, envId, opts...)
}

func UnregisterProject(projId string, opts ...DeployOption) error {
	deployOpts := fillDeployOptions(opts...)

	deleteProjectReq := &models.DeleteProjectRequest{
		ID: projId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	deleteProjectParams := service_runner.NewDeleteProjectParams().
		WithBody(deleteProjectReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DeleteProjectOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DeleteProject(deleteProjectParams,
			deployOpts)
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
	if deleteReply.Result.Status != nil &&
		*deleteReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("DeleteProject failure(%v): %v",
			*deleteReply.Result.Status, deleteReply.Result.StatusDetail)
	}

	return nil
}

func DeactivateProject(projId string, envId string,
	opts ...DeployOption) (string, error) {

	deployOpts := fillDeployOptions(opts...)

	deactivateProjectReq := &models.DeactivateProjectRequest{
		ProjEnvHeader: &models.ProjEnvHeader{
			ProjID: projId,
			EnvID:  envId,
		},
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	deactivateProjectParams := service_runner.NewDeactivateProjectParams().
		WithBody(deactivateProjectReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DeactivateProjectOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err =
			client.ServiceRunner.DeactivateProject(deactivateProjectParams,
				deployOpts)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return "", fmt.Errorf("Client/HTTP failure: %v", err)
	}
	deactivateReply := resp.GetPayload()
	if deactivateReply.Result.Status != nil &&
		*deactivateReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return "", fmt.Errorf("DeactivateProject failure(%v): %v",
			*deactivateReply.Result.Status, deactivateReply.Result.StatusDetail)
	}

	return deactivateReply.DeployID, nil
}

func (proj *Project) Register(opts ...DeployOption) error {
	deployOpts := fillDeployOptions(opts...)

	createProjectReq := &models.CreateProjectRequest{
		Header: &models.ProjectHeader{
			DNSDomain: "",
			DNSPrefix: "",
			Name:      proj.Desc.Name,
		},
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	createProjectParams := service_runner.NewCreateProjectParams().
		WithBody(createProjectReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.CreateProjectOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.CreateProject(createProjectParams,
			deployOpts)
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
	createReply := resp.GetPayload()
	if createReply.Result.Status != nil &&
		*createReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("Createproject failure(%v): %v",
			*createReply.Result.Status, createReply.Result.StatusDetail)
	}

	proj.Desc.Id = createReply.ID
	return proj.ExportToFile(proj.Desc.projFile)
}

func DescribeProject(projId string, opts ...DeployOption) (*pb.ProjectDescription, error) {

	deployOpts := fillDeployOptions(opts...)

	describeProjectReq := &models.DescribeProjectRequest{
		ID: projId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	describeProjectParams := service_runner.NewDescribeProjectParams().
		WithBody(describeProjectReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DescribeProjectOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeProject(describeProjectParams,
			deployOpts)
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
	if describeReply.Result.Status != nil &&
		*describeReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil,
			fmt.Errorf("DescribeProject failure(%v): %v",
				*describeReply.Result.Status, describeReply.Result.StatusDetail)
	}

	var ok bool
	projStateInt := int32(0)
	projState := pb.ProjectState(projStateInt)
	if describeReply.Desc.State != nil {
		projStateInt, ok = pb.ProjectState_value[string(*describeReply.Desc.State)]
		projState = pb.ProjectState(projStateInt)
		if !ok {
			projState = pb.ProjectState_UNKNOWN_PROJ_STATE
		}
	}

	ret := &pb.ProjectDescription{
		Id: describeReply.Desc.ID,
		Header: &pb.ProjectHeader{
			Name:      describeReply.Desc.Header.Name,
			DnsPrefix: describeReply.Desc.Header.DNSPrefix,
			DnsDomain: describeReply.Desc.Header.DNSDomain,
		},
		State:            projState,
		CreateTime:       convertRESTTimeToInt(describeReply.Desc.CreateTime),
		ActiveDeployIds:  describeReply.Desc.ActiveDeployIds,
		PendingDeployIds: describeReply.Desc.PendingDeployIds,
	}

	return ret, nil
}

// NewProject instantiates a new Project instance from the specified project
// filename. It will parse the project file & derivative API definitions for
// each service defined in the project file and validate that it is well
// formed. Otherwise, an error is returned. Here is an example Bopmatic.yaml
// project file:
//
// formatversion: "1.1"
// project:
//
//	name: "Foo"
//	desc: "Foo Project"
//	sitedir: "site"
//	services:
//	- name: "Greeter"
//	  desc: "Service for greeting customers"
//	  apidef: "pb/greeter.proto"
//	  port: 26001
//	  executable: "greeter_server"
//	  useraccess: "anon_public"
//	- name: "Orders"
//	  desc: "Service for taking customer orders"
//	  apidef: "pb/orders.proto"
//	  port: 26002
//	  executable: "orders_server"
//	  useraccess: "CustomerUserGroup"
//	databases:
//	- name: "Customers"
//	  desc: "Customer database"
//	  tables:
//	  - name: "ContactDetails"
//	    desc: "Customer names, shipping address, phone, etc."
//	  - name: "Orders"
//	    desc: "Customer orders"
//	  services_access: [ "Greeter" ]
//	object_stores:
//	- name: "UploadBucket"
//	  desc: "Bucket for uploading data"
//	  services_access: [ "Orders" ]
//	usergroups:
//	- name: "CustomerUserGroup"
//	  desc: "Customer user group"
//	  type: "public"
func NewProject(projFile string, opts ...ProjectOption) (*Project, error) {

	projectOpts := fillProjectOptions(opts...)

	proj, err := parseProject(projFile)
	if err != nil {
		return nil, err
	}

	err = proj.validateProject(projFile, projectOpts.validateSiteAssets)
	if err != nil {
		return nil, err
	}

	return proj, nil
}
