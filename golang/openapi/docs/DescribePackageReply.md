# DescribePackageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to [**PackageDescription**](PackageDescription.md) |  | [optional] 
**PackageState** | Pointer to [**PackageState**](PackageState.md) |  | [optional] [default to UPLOADING]
**SiteEndpoint** | Pointer to **string** |  | [optional] 
**RpcEndpoints** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDescribePackageReply

`func NewDescribePackageReply() *DescribePackageReply`

NewDescribePackageReply instantiates a new DescribePackageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePackageReplyWithDefaults

`func NewDescribePackageReplyWithDefaults() *DescribePackageReply`

NewDescribePackageReplyWithDefaults instantiates a new DescribePackageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *DescribePackageReply) GetDesc() PackageDescription`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DescribePackageReply) GetDescOk() (*PackageDescription, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DescribePackageReply) SetDesc(v PackageDescription)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DescribePackageReply) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetPackageState

`func (o *DescribePackageReply) GetPackageState() PackageState`

GetPackageState returns the PackageState field if non-nil, zero value otherwise.

### GetPackageStateOk

`func (o *DescribePackageReply) GetPackageStateOk() (*PackageState, bool)`

GetPackageStateOk returns a tuple with the PackageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageState

`func (o *DescribePackageReply) SetPackageState(v PackageState)`

SetPackageState sets PackageState field to given value.

### HasPackageState

`func (o *DescribePackageReply) HasPackageState() bool`

HasPackageState returns a boolean if a field has been set.

### GetSiteEndpoint

`func (o *DescribePackageReply) GetSiteEndpoint() string`

GetSiteEndpoint returns the SiteEndpoint field if non-nil, zero value otherwise.

### GetSiteEndpointOk

`func (o *DescribePackageReply) GetSiteEndpointOk() (*string, bool)`

GetSiteEndpointOk returns a tuple with the SiteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteEndpoint

`func (o *DescribePackageReply) SetSiteEndpoint(v string)`

SetSiteEndpoint sets SiteEndpoint field to given value.

### HasSiteEndpoint

`func (o *DescribePackageReply) HasSiteEndpoint() bool`

HasSiteEndpoint returns a boolean if a field has been set.

### GetRpcEndpoints

`func (o *DescribePackageReply) GetRpcEndpoints() []string`

GetRpcEndpoints returns the RpcEndpoints field if non-nil, zero value otherwise.

### GetRpcEndpointsOk

`func (o *DescribePackageReply) GetRpcEndpointsOk() (*[]string, bool)`

GetRpcEndpointsOk returns a tuple with the RpcEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcEndpoints

`func (o *DescribePackageReply) SetRpcEndpoints(v []string)`

SetRpcEndpoints sets RpcEndpoints field to given value.

### HasRpcEndpoints

`func (o *DescribePackageReply) HasRpcEndpoints() bool`

HasRpcEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


