# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is the Bopmatic Golang SDK, which provides Go client libraries for interacting with the Bopmatic platform. The SDK enables developers to create, deploy, and manage cloud-native applications with services, databases, object stores, and user groups.

## Build and Test Commands

### Building
```bash
make build
```
This compiles the SDK and utility packages. Requires vendor dependencies to be present.

### Testing
```bash
make test
```
Runs all tests with verbose output, including both main package and util package tests.

### Running Specific Tests
```bash
go test -v -run TestName
go test -v github.com/bopmatic/sdk/golang/util -run TestName
```

### Generating Unit Test Reports
```bash
make unit-tests.xml
```
Creates JUnit XML format test results (requires `gotestsum` to be installed).

## Code Generation Workflow

This repository heavily relies on code generation from Protocol Buffers definitions. The workflow involves multiple stages:

### 1. Protocol Buffer to Go/gRPC/OpenAPI
The core service definitions are in `pb/sr.proto` (Service Runner) and `pb/signup.proto`. These are compiled to:
- Go protobuf structs (`pb/sr.pb.go`, `pb/signup.pb.go`)
- gRPC client/server code (`pb/sr_grpc.pb.go`, `pb/signup_grpc.pb.go`)
- OpenAPI/Swagger specs (`pb/sr.swagger.json`, `pb/signup.swagger.json`)

**Command:**
```bash
protoc -I ./ -I ./pb -I /usr/local/include --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --openapiv2_out=./ --openapiv2_opt=simple_operation_ids=true,generate_unbound_methods=true ./pb/sr.proto
```

### 2. Bopmatic-Specific OpenAPI Specs
The base swagger specs are enhanced with Bopmatic-specific configuration (host, schemes, security) to create `pb/sr.bopmatic.json` and `pb/signup.bopmatic.json`. This is done with `jq` transformations.

### 3. Go Swagger Client Generation
Go-swagger generates HTTP clients from the bopmatic.json specs into the `goswag/` and `goswag_signup/` directories:
```bash
swagger generate client -f pb/sr.bopmatic.json -c goswag --default-scheme='https' -A goswag
```

### 4. OpenAPI Generator for Alternative Clients
The OpenAPI Generator creates an additional Go client in the `openapi/` directory:
```bash
make openapi
```

### Regenerating All Generated Code
If you modify proto files, regenerate everything with:
```bash
make clean
make deps    # This regenerates pb/, goswag/, and openapi/ directories
make build
```

**Note:** The `make deps` target is destructive and will reinitialize go.mod/go.sum.

## Architecture

### Core Abstractions

1. **Project** (`project.go`): The top-level abstraction representing a Bopmatic application. A project is defined in a `Bopmatic.yaml` file and contains:
   - Services (microservices with API definitions)
   - Databases (DAPR state stores with tables)
   - Object Stores (S3-compatible storage)
   - User Groups (authentication/authorization)
   - Site Assets (static web content)
   - Build configuration

2. **Package** (`package.go`): A deployable artifact containing compiled code and assets. Created from a Project via tarball packaging.

3. **Deployment** (`deployment.go`): Represents an active deployment of a Package to a specific Environment within a Project.

4. **Environment**: Projects can have multiple environments (dev, staging, prod). Each environment has isolated deployments.

5. **API Key** (`apikey.go`): Credentials for authenticating API requests to the Bopmatic platform.

### Project Configuration (Bopmatic.yaml)

Projects are defined declaratively in YAML. Example structure:
```yaml
formatversion: "1.0"
project:
  name: "MyApp"
  desc: "My Application"
  buildcmd: "make"
  services:
  - name: "Greeter"
    desc: "Greeting service"
    apidef: "pb/greeter.proto"    # or OpenAPI spec
    port: 26001
    executable: "greeter_server"
    user_access: "anon_public"     # or "none"
  databases:
  - name: "UserDB"
    tables:
    - name: "users"
    services_access: ["Greeter"]
  object_stores:
  - name: "Assets"
    services_access: ["Greeter"]
```

### API Type Support

Services can use one of three API types:
- **gRPC/Protobuf**: `apidef` points to a `.proto` file
- **OpenAPI**: `apidef` points to an OpenAPI/Swagger spec JSON file
- **MCP (Model Context Protocol)**: `apidef` points to an MCP manifest JSON file

The SDK automatically detects the API type by parsing the definition file:
- `.proto` extension → gRPC
- `.json` with MCP fields (`tools`, `resources`, `prompts`) → MCP
- `.json` with OpenAPI fields → OpenAPI

#### MCP Support

MCP (Model Context Protocol) is Anthropic's open protocol for AI-LLM integration with external tools and data sources. An MCP manifest defines:

1. **Tools**: Functions that can be invoked (e.g., `create_task`, `list_tasks`)
2. **Resources**: Data sources that can be accessed (e.g., databases, config files)
3. **Prompts**: Templated prompts for specific workflows

When parsing MCP manifests, the SDK extracts:
- Tool names as direct RPCs
- Resource names prefixed with `resource_` as RPCs
- Prompt names prefixed with `prompt_` as RPCs

Example MCP service in `Bopmatic.yaml`:
```yaml
services:
  - name: "TaskService"
    desc: "MCP-based task management"
    apidef: "mcp-manifest.json"
    port: 26001
    executable: "task_server"
```

See `test_assets/project/mcpOnly/` for a complete MCP example.

### Container-Based Builds

The SDK uses Docker containers for building projects via the `util/container.go` utilities. The Bopmatic build image (`bopmatic/build:latest`) provides a consistent build environment.

### Generated Client Structure

The SDK provides multiple ways to interact with Bopmatic services:

1. **High-level SDK functions** (in package root): `CreateProject()`, `UploadPackage()`, `CreateDeployment()`, etc.
2. **goswag clients**: Auto-generated HTTP clients in `goswag/service_runner/` and `goswag_signup/sign_up/`
3. **openapi clients**: Alternative OpenAPI-generated clients in `openapi/`
4. **protobuf/gRPC**: Direct gRPC access via `pb/` (though SDK primarily uses HTTP)

### Key Files

- `project.go`: Project parsing, validation, and API extraction
- `project_details.go`: Extended project operations (upload, describe, list, etc.)
- `package.go`: Package creation, upload, and management
- `deployment.go`: Deployment operations
- `logs.go`: Log retrieval from deployments
- `util/container.go`: Docker interaction utilities
- `util/dapr.go`: DAPR runtime utilities
- `pb/sr.proto`: Service Runner gRPC service definition
- `models/`: Generated OpenAPI models

## Development Notes

### Vendor Dependencies
This project uses Go modules with vendoring (`GOFLAGS=-mod=vendor`). After modifying `go.mod`, run:
```bash
go mod download
go mod vendor
```

### Test Assets
The `test_assets/` directory contains sample projects and packages used for testing. These include both valid configurations and intentionally invalid ones for error handling tests.

### Adding New Service Runner APIs

1. Add the RPC definition to `pb/sr.proto`
2. Run `make clean && make deps && make build`
3. Implement high-level SDK functions in the appropriate file (e.g., `project.go`, `deployment.go`)
4. The generated clients in `goswag/`, `openapi/`, and `models/` will be automatically updated

### Authentication
SDK functions accept optional `DeployOption` parameters for authentication. The primary auth methods are:
- Bearer tokens (user sessions)
- API keys (programmatic access)

These are passed via HTTP headers to the Bopmatic API (api.bopmatic.com).
