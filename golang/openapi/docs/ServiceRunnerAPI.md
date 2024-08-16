# \ServiceRunnerAPI

All URIs are relative to *https://api.bopmatic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePackage**](ServiceRunnerAPI.md#DeletePackage) | **Post** /ServiceRunner/DeletePackage | 
[**DeployPackage**](ServiceRunnerAPI.md#DeployPackage) | **Post** /ServiceRunner/DeployPackage | 
[**DescribePackage**](ServiceRunnerAPI.md#DescribePackage) | **Post** /ServiceRunner/DescribePackage | 
[**GetLogs**](ServiceRunnerAPI.md#GetLogs) | **Post** /ServiceRunner/GetLogs | 
[**GetUploadURL**](ServiceRunnerAPI.md#GetUploadURL) | **Post** /ServiceRunner/GetUploadURL | 
[**ListPackages**](ServiceRunnerAPI.md#ListPackages) | **Post** /ServiceRunner/ListPackages | 



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


## DeployPackage

> DeployPackageReply DeployPackage(ctx).Body(body).Execute()



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
	body := *openapiclient.NewDeployPackageRequest() // DeployPackageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRunnerAPI.DeployPackage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerAPI.DeployPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployPackage`: DeployPackageReply
	fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerAPI.DeployPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeployPackageRequest**](DeployPackageRequest.md) |  | 

### Return type

[**DeployPackageReply**](DeployPackageReply.md)

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

