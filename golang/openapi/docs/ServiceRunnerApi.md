# \ServiceRunnerApi

All URIs are relative to *https://api.bopmatic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePackage**](ServiceRunnerApi.md#DeletePackage) | **Post** /ServiceRunner/DeletePackage | 
[**DeployPackage**](ServiceRunnerApi.md#DeployPackage) | **Post** /ServiceRunner/DeployPackage | 
[**DescribePackage**](ServiceRunnerApi.md#DescribePackage) | **Post** /ServiceRunner/DescribePackage | 
[**ListPackages**](ServiceRunnerApi.md#ListPackages) | **Post** /ServiceRunner/ListPackages | 



## DeletePackage

> DeletePackageReply DeletePackage(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDeletePackageRequest() // DeletePackageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRunnerApi.DeletePackage(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerApi.DeletePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePackage`: DeletePackageReply
    fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerApi.DeletePackage`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDeployPackageRequest() // DeployPackageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRunnerApi.DeployPackage(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerApi.DeployPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployPackage`: DeployPackageReply
    fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerApi.DeployPackage`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDescribePackageRequest() // DescribePackageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRunnerApi.DescribePackage(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerApi.DescribePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribePackage`: DescribePackageReply
    fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerApi.DescribePackage`: %v\n", resp)
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


## ListPackages

> ListPackagesReply ListPackages(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewListPackagesRequest() // ListPackagesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRunnerApi.ListPackages(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRunnerApi.ListPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPackages`: ListPackagesReply
    fmt.Fprintf(os.Stdout, "Response from `ServiceRunnerApi.ListPackages`: %v\n", resp)
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

