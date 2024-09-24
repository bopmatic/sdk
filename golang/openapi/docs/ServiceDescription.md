# ServiceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcHeader** | Pointer to [**ServiceHeader**](ServiceHeader.md) |  | [optional] 
**RpcEndpoints** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceDescription

`func NewServiceDescription() *ServiceDescription`

NewServiceDescription instantiates a new ServiceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDescriptionWithDefaults

`func NewServiceDescriptionWithDefaults() *ServiceDescription`

NewServiceDescriptionWithDefaults instantiates a new ServiceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcHeader

`func (o *ServiceDescription) GetSvcHeader() ServiceHeader`

GetSvcHeader returns the SvcHeader field if non-nil, zero value otherwise.

### GetSvcHeaderOk

`func (o *ServiceDescription) GetSvcHeaderOk() (*ServiceHeader, bool)`

GetSvcHeaderOk returns a tuple with the SvcHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcHeader

`func (o *ServiceDescription) SetSvcHeader(v ServiceHeader)`

SetSvcHeader sets SvcHeader field to given value.

### HasSvcHeader

`func (o *ServiceDescription) HasSvcHeader() bool`

HasSvcHeader returns a boolean if a field has been set.

### GetRpcEndpoints

`func (o *ServiceDescription) GetRpcEndpoints() []string`

GetRpcEndpoints returns the RpcEndpoints field if non-nil, zero value otherwise.

### GetRpcEndpointsOk

`func (o *ServiceDescription) GetRpcEndpointsOk() (*[]string, bool)`

GetRpcEndpointsOk returns a tuple with the RpcEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcEndpoints

`func (o *ServiceDescription) SetRpcEndpoints(v []string)`

SetRpcEndpoints sets RpcEndpoints field to given value.

### HasRpcEndpoints

`func (o *ServiceDescription) HasRpcEndpoints() bool`

HasRpcEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


