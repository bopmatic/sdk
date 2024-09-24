# ServiceHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjEnvHeader** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceHeader

`func NewServiceHeader() *ServiceHeader`

NewServiceHeader instantiates a new ServiceHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHeaderWithDefaults

`func NewServiceHeaderWithDefaults() *ServiceHeader`

NewServiceHeaderWithDefaults instantiates a new ServiceHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjEnvHeader

`func (o *ServiceHeader) GetProjEnvHeader() ProjEnvHeader`

GetProjEnvHeader returns the ProjEnvHeader field if non-nil, zero value otherwise.

### GetProjEnvHeaderOk

`func (o *ServiceHeader) GetProjEnvHeaderOk() (*ProjEnvHeader, bool)`

GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjEnvHeader

`func (o *ServiceHeader) SetProjEnvHeader(v ProjEnvHeader)`

SetProjEnvHeader sets ProjEnvHeader field to given value.

### HasProjEnvHeader

`func (o *ServiceHeader) HasProjEnvHeader() bool`

HasProjEnvHeader returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceHeader) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceHeader) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceHeader) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceHeader) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


