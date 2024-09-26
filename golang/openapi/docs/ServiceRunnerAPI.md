# \ServiceRunnerAPI

All URIs are relative to *https://api.bopmatic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeployment**](ServiceRunnerAPI.md#CreateDeployment) | **Post** /ServiceRunner/CreateDeployment | 
[**CreateEnvironment**](ServiceRunnerAPI.md#CreateEnvironment) | **Post** /ServiceRunner/CreateEnvironment | 
[**CreateProject**](ServiceRunnerAPI.md#CreateProject) | **Post** /ServiceRunner/CreateProject | 
[**DeleteEnvironment**](ServiceRunnerAPI.md#DeleteEnvironment) | **Post** /ServiceRunner/DeleteEnvironment | 
[**DeletePackage**](ServiceRunnerAPI.md#DeletePackage) | **Post** /ServiceRunner/DeletePackage | 
[**DeleteProject**](ServiceRunnerAPI.md#DeleteProject) | **Post** /ServiceRunner/DeleteProject | 
[**DescribeDatabase**](ServiceRunnerAPI.md#DescribeDatabase) | **Post** /ServiceRunner/DescribeDatabase | 
[**DescribeDatastore**](ServiceRunnerAPI.md#DescribeDatastore) | **Post** /ServiceRunner/DescribeDatastore | 
[**DescribeDeployment**](ServiceRunnerAPI.md#DescribeDeployment) | **Post** /ServiceRunner/DescribeDeployment | 
[**DescribeEnvironment**](ServiceRunnerAPI.md#DescribeEnvironment) | **Post** /ServiceRunner/DescribeEnvironment | 
[**DescribePackage**](ServiceRunnerAPI.md#DescribePackage) | **Post** /ServiceRunner/DescribePackage | 
[**DescribeProject**](ServiceRunnerAPI.md#DescribeProject) | **Post** /ServiceRunner/DescribeProject | 
[**DescribeService**](ServiceRunnerAPI.md#DescribeService) | **Post** /ServiceRunner/DescribeService | 
[**DescribeSite**](ServiceRunnerAPI.md#DescribeSite) | **Post** /ServiceRunner/DescribeSite | 
[**GetLogs**](ServiceRunnerAPI.md#GetLogs) | **Post** /ServiceRunner/GetLogs | 
[**GetMetricSamples**](ServiceRunnerAPI.md#GetMetricSamples) | **Post** /ServiceRunner/GetMetricSamples | 
[**GetUploadURL**](ServiceRunnerAPI.md#GetUploadURL) | **Post** /ServiceRunner/GetUploadURL | 
[**ListDatabases**](ServiceRunnerAPI.md#ListDatabases) | **Post** /ServiceRunner/ListDatabases | 
[**ListDatastores**](ServiceRunnerAPI.md#ListDatastores) | **Post** /ServiceRunner/ListDatastores | 
[**ListDeployments**](ServiceRunnerAPI.md#ListDeployments) | **Post** /ServiceRunner/ListDeployments | 
[**ListEnvironments**](ServiceRunnerAPI.md#ListEnvironments) | **Post** /ServiceRunner/ListEnvironments | 
[**ListMetrics**](ServiceRunnerAPI.md#ListMetrics) | **Post** /ServiceRunner/ListMetrics | 
[**ListPackages**](ServiceRunnerAPI.md#ListPackages) | **Post** /ServiceRunner/ListPackages | 
[**ListProjects**](ServiceRunnerAPI.md#ListProjects) | **Post** /ServiceRunner/ListProjects | 
[**ListServices**](ServiceRunnerAPI.md#ListServices) | **Post** /ServiceRunner/ListServices | 
[**UploadPackage**](ServiceRunnerAPI.md#UploadPackage) | **Post** /ServiceRunner/UploadPackage | 



## CreateDeployment

> CreateDeploymentReply CreateDeployment(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewCreateDeploymentRequest() // CreateDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.CreateDeployment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.CreateDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeployment`: CreateDeploymentReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.CreateDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDeploymentRequest**](CreateDeploymentRequest.md) |  | 

### Return type

[**CreateDeploymentReply**](CreateDeploymentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> CreateEnvironmentReply CreateEnvironment(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewCreateEnvironmentRequest() // CreateEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.CreateEnvironment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.CreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironment`: CreateEnvironmentReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateEnvironmentRequest**](CreateEnvironmentRequest.md) |  | 

### Return type

[**CreateEnvironmentReply**](CreateEnvironmentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> CreateProjectReply CreateProject(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewCreateProjectRequest() // CreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.CreateProject(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: CreateProjectReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateProjectRequest**](CreateProjectRequest.md) |  | 

### Return type

[**CreateProjectReply**](CreateProjectReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironmentReply DeleteEnvironment(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDeleteEnvironmentRequest() // DeleteEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DeleteEnvironment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnvironment`: DeleteEnvironmentReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteEnvironmentRequest**](DeleteEnvironmentRequest.md) |  | 

### Return type

[**DeleteEnvironmentReply**](DeleteEnvironmentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePackage

> DeletePackageReply DeletePackage(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDeletePackageRequest() // DeletePackageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DeletePackage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DeletePackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePackage`: DeletePackageReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DeletePackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeletePackageRequest**](DeletePackageRequest.md) |  | 

### Return type

[**DeletePackageReply**](DeletePackageReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProjectReply DeleteProject(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDeleteProjectRequest() // DeleteProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DeleteProject(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProject`: DeleteProjectReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DeleteProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteProjectRequest**](DeleteProjectRequest.md) |  | 

### Return type

[**DeleteProjectReply**](DeleteProjectReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDatabase

> DescribeDatabaseReply DescribeDatabase(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeDatabaseRequest() // DescribeDatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeDatabase(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDatabase`: DescribeDatabaseReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeDatabaseRequest**](DescribeDatabaseRequest.md) |  | 

### Return type

[**DescribeDatabaseReply**](DescribeDatabaseReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDatastore

> DescribeDatastoreReply DescribeDatastore(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeDatastoreRequest() // DescribeDatastoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeDatastore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDatastore`: DescribeDatastoreReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeDatastore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeDatastoreRequest**](DescribeDatastoreRequest.md) |  | 

### Return type

[**DescribeDatastoreReply**](DescribeDatastoreReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDeployment

> DescribeDeploymentReply DescribeDeployment(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeDeploymentRequest() // DescribeDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeDeployment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDeployment`: DescribeDeploymentReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeDeploymentRequest**](DescribeDeploymentRequest.md) |  | 

### Return type

[**DescribeDeploymentReply**](DescribeDeploymentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeEnvironment

> DescribeEnvironmentReply DescribeEnvironment(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeEnvironmentRequest() // DescribeEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeEnvironment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeEnvironment`: DescribeEnvironmentReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeEnvironmentRequest**](DescribeEnvironmentRequest.md) |  | 

### Return type

[**DescribeEnvironmentReply**](DescribeEnvironmentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePackage

> DescribePackageReply DescribePackage(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribePackageRequest() // DescribePackageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribePackage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribePackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePackage`: DescribePackageReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribePackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribePackageRequest**](DescribePackageRequest.md) |  | 

### Return type

[**DescribePackageReply**](DescribePackageReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProject

> DescribeProjectReply DescribeProject(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeProjectRequest() // DescribeProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeProject(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProject`: DescribeProjectReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeProjectRequest**](DescribeProjectRequest.md) |  | 

### Return type

[**DescribeProjectReply**](DescribeProjectReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeService

> DescribeServiceReply DescribeService(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeServiceRequest() // DescribeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeService(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeService`: DescribeServiceReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeServiceRequest**](DescribeServiceRequest.md) |  | 

### Return type

[**DescribeServiceReply**](DescribeServiceReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSite

> DescribeSiteReply DescribeSite(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewDescribeSiteRequest() // DescribeSiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DescribeSite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DescribeSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeSite`: DescribeSiteReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DescribeSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeSiteRequest**](DescribeSiteRequest.md) |  | 

### Return type

[**DescribeSiteReply**](DescribeSiteReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> GetLogsReply GetLogs(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewGetLogsRequest() // GetLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.GetLogs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: GetLogsReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetLogsRequest**](GetLogsRequest.md) |  | 

### Return type

[**GetLogsReply**](GetLogsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricSamples

> GetMetricSamplesReply GetMetricSamples(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewGetMetricSamplesRequest() // GetMetricSamplesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.GetMetricSamples(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.GetMetricSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricSamples`: GetMetricSamplesReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.GetMetricSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetMetricSamplesRequest**](GetMetricSamplesRequest.md) |  | 

### Return type

[**GetMetricSamplesReply**](GetMetricSamplesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadURL

> GetUploadURLReply GetUploadURL(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewGetUploadURLRequest() // GetUploadURLRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.GetUploadURL(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.GetUploadURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUploadURL`: GetUploadURLReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.GetUploadURL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetUploadURLRequest**](GetUploadURLRequest.md) |  | 

### Return type

[**GetUploadURLReply**](GetUploadURLReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> ListDatabasesReply ListDatabases(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListDatabasesRequest() // ListDatabasesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListDatabases(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabases`: ListDatabasesReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListDatabasesRequest**](ListDatabasesRequest.md) |  | 

### Return type

[**ListDatabasesReply**](ListDatabasesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatastores

> ListDatastoresReply ListDatastores(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListDatastoresRequest() // ListDatastoresRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListDatastores(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatastores`: ListDatastoresReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListDatastores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListDatastoresRequest**](ListDatastoresRequest.md) |  | 

### Return type

[**ListDatastoresReply**](ListDatastoresReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployments

> ListDeploymentsReply ListDeployments(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListDeploymentsRequest() // ListDeploymentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListDeployments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeployments`: ListDeploymentsReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListDeploymentsRequest**](ListDeploymentsRequest.md) |  | 

### Return type

[**ListDeploymentsReply**](ListDeploymentsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironments

> ListEnvironmentsReply ListEnvironments(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListEnvironments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironments`: ListEnvironmentsReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListEnvironmentsReply**](ListEnvironmentsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> ListMetricsReply ListMetrics(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListMetricsRequest() // ListMetricsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListMetrics(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetrics`: ListMetricsReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListMetricsRequest**](ListMetricsRequest.md) |  | 

### Return type

[**ListMetricsReply**](ListMetricsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPackages

> ListPackagesReply ListPackages(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListPackagesRequest() // ListPackagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListPackages(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPackages`: ListPackagesReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListPackagesRequest**](ListPackagesRequest.md) |  | 

### Return type

[**ListPackagesReply**](ListPackagesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> ListProjectsReply ListProjects(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListProjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjects`: ListProjectsReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListProjectsReply**](ListProjectsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ListServicesReply ListServices(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewListServicesRequest() // ListServicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.ListServices(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.ListServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServices`: ListServicesReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.ListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListServicesRequest**](ListServicesRequest.md) |  | 

### Return type

[**ListServicesReply**](ListServicesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPackage

> UploadPackageReply UploadPackage(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bopmatic/sdk/golang/openapi"
)

func main() {
	body := *openapiclient.NewUploadPackageRequest() // UploadPackageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.UploadPackage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.UploadPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPackage`: UploadPackageReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.UploadPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UploadPackageRequest**](UploadPackageRequest.md) |  | 

### Return type

[**UploadPackageReply**](UploadPackageReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

