package golang

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

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
	Description string `yaml:"desc"`
}

// Database is a group of DatabaseTables accessible by one or more services.
// Running services may access databases by employing the DAPR state store APIs
// via the SDK of the user's preferred programming language. See
// https://docs.dapr.io/developing-applications/sdks/ for further detail.
type Database struct {
	Name        string          `yaml:"name"`
	Description string          `yaml:"desc"`
	Tables      []DatabaseTable `yaml:"tables"`
	Services    []string        `yaml:"services_access"`
}

// Service is a representation of an individual service defined within
// a Bopmatic project. This includes its name, a short description, a link
// to the set of APIs it defines, and the port it should run on. Fields marked
// with yaml are parsed from Bopmatic.yaml while the remaining fields are
// derived once the project is parsed (e.g. when NewProject() returns
// successfully).
// Currently the only supported ApiDefinition format is protobuf/GRPC
type Service struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"desc"`
	ApiDefinition string `yaml:"apidef"`
	Port          uint64 `yaml:"port"`
	Executable    string `yaml:"executable"`
	ExecAssets    string `yaml:"executable_assets"`

	Rpcs []string
}

// ProjectDesc see Project for a complete description
type ProjectDesc struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"desc"`
	Services    []Service  `yaml:"services"`
	Databases   []Database `yaml:"databases"`
	SiteAssets  string     `yaml:"sitedir"`
	BuildCmd    string     `yaml:"buildcmd"`

	Root string
}

// Project is a full representation of a Bopmatic project, which defines
// a set of services to publish, their RPCs, their TCP ports to run on,
// and a collection of static website assets to publish. FormatVersion
// describes the set of supported and expected fields within a Bopmatic.yaml
// project description. Currently this is expected to be "1.0" and is unused;
// however it is intended to be used in the future for backward compatibility
// purposes. Fields marked with yaml are parsed from Bopmatic.yaml while the
// remaining fields are derived once the project is parsed (e.g. when
// NewProject() returns successfully).
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
	sb.WriteString(fmt.Sprintf("\tSiteAssets: %v\n", proj.Desc.SiteAssets))
	if proj.Desc.BuildCmd != "" {
		sb.WriteString(fmt.Sprintf("\tBuildCmd: %v\n", proj.Desc.BuildCmd))
	}
	sb.WriteString(fmt.Sprintf("\tRoot: %v\n", proj.Desc.Root))
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
		sb.WriteString(fmt.Sprintf("\t\tPort: %v\n", svc.Port))
		sb.WriteString(fmt.Sprintf("\t\tExecutable: %v\n", svc.Executable))
		if svc.ExecAssets != "" {
			sb.WriteString(fmt.Sprintf("\t\tExecAssets: %v\n", svc.ExecAssets))
		}
		sb.WriteString(fmt.Sprintf("\t\tRpcs: %v\n", len(svc.Rpcs)))
		for idx2, funcName := range svc.Rpcs {
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
	return sb.String()
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
		proj.Desc.Root = path.Dir(projFile)
	} else {
		proj.Desc.Root = path.Dir(projFile)
	}

	return &proj, nil
}

func (svc *Service) populateRpcs(svcName string,
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
				svc.Rpcs = append(svc.Rpcs, rpc.RPCName)
			}
		}
	}

	if !foundSvc {
		return fmt.Errorf("Failed to find service %v in %v", svcName,
			protoFileName)
	}

	if len(svc.Rpcs) == 0 {
		return fmt.Errorf("Service %v in %v must define at least 1 RPC", svcName,
			protoFileName)
	}

	return nil
}

func (proj *Project) validateProject(projFile string, validateSiteAssets bool) error {
	const missingFieldFmt = "%v %v definition is missing required field %v"

	if proj.FormatVersion == "" {
		return fmt.Errorf(missingFieldFmt, "Project", projFile, "formatversion")
	}
	if proj.Desc.Name == "" {
		return fmt.Errorf(missingFieldFmt, "Project", projFile, "name")
	}

	if validateSiteAssets && proj.Desc.SiteAssets != "" {
		siteAssetsPath := filepath.Join(proj.Desc.Root, proj.Desc.SiteAssets)
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

		usedPorts = append(usedPorts, svc.Port)
		apiDefPath := filepath.Join(proj.Desc.Root, svc.ApiDefinition)
		file, err := os.Open(apiDefPath)
		if err != nil {
			return fmt.Errorf("Failed to open API definition for Service %v: %w", svc.Name, err)
		}
		defer file.Close()

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
	err = os.Chdir(proj.Desc.Root)
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

// NewProject instantiates a new Project instance from the specified project
// filename. It will parse the project file & derivative API definitions for
// each service defined in the project file and validate that it is well
// formed. Otherwise, an error is returned. Here is an example Bopmatic.yaml
// project file:
//
// formatversion: "1.0"
// project:
//   name: "Foo"
//   desc: "Foo Project"
//   sitedir: "site"
//   services:
//   - name: "Greeter"
//     desc: "Service for greeting customers"
//     apidef: "pb/greeter.proto"
//     port: 26001
//     executable: "greeter_server"
//   - name: "Orders"
//     desc: "Service for taking customer orders"
//     apidef: "pb/orders.proto"
//     port: 26002
//     executable: "orders_server"
//   databases:
//   - name: "Customers"
//     desc: "Customer database"
//     tables:
//     - name: "ContactDetails"
//       desc: "Customer names, shipping address, phone, etc."
//     - name: "Orders"
//       desc: "Customer orders"
//     services_access: [ "Greeter" ]
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
