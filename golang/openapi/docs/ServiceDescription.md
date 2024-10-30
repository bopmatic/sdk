# ServiceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcHeader** | Pointer to [**ServiceHeader**](ServiceHeader.md) |  | [optional] 
**RpcEndpoints** | Pointer to **[]string** |  | [optional] 
**DatabaseNames** | Pointer to **[]string** |  | [optional] 
**DatastoreNames** | Pointer to **[]string** |  | [optional] 
**ApiDef** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 

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

### GetDatabaseNames

`func (o *ServiceDescription) GetDatabaseNames() []string`

GetDatabaseNames returns the DatabaseNames field if non-nil, zero value otherwise.

### GetDatabaseNamesOk

`func (o *ServiceDescription) GetDatabaseNamesOk() (*[]string, bool)`

GetDatabaseNamesOk returns a tuple with the DatabaseNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseNames

`func (o *ServiceDescription) SetDatabaseNames(v []string)`

SetDatabaseNames sets DatabaseNames field to given value.

### HasDatabaseNames

`func (o *ServiceDescription) HasDatabaseNames() bool`

HasDatabaseNames returns a boolean if a field has been set.

### GetDatastoreNames

`func (o *ServiceDescription) GetDatastoreNames() []string`

GetDatastoreNames returns the DatastoreNames field if non-nil, zero value otherwise.

### GetDatastoreNamesOk

`func (o *ServiceDescription) GetDatastoreNamesOk() (*[]string, bool)`

GetDatastoreNamesOk returns a tuple with the DatastoreNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreNames

`func (o *ServiceDescription) SetDatastoreNames(v []string)`

SetDatastoreNames sets DatastoreNames field to given value.

### HasDatastoreNames

`func (o *ServiceDescription) HasDatastoreNames() bool`

HasDatastoreNames returns a boolean if a field has been set.

### GetApiDef

`func (o *ServiceDescription) GetApiDef() string`

GetApiDef returns the ApiDef field if non-nil, zero value otherwise.

### GetApiDefOk

`func (o *ServiceDescription) GetApiDefOk() (*string, bool)`

GetApiDefOk returns a tuple with the ApiDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDef

`func (o *ServiceDescription) SetApiDef(v string)`

SetApiDef sets ApiDef field to given value.

### HasApiDef

`func (o *ServiceDescription) HasApiDef() bool`

HasApiDef returns a boolean if a field has been set.

### GetPort

`func (o *ServiceDescription) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServiceDescription) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServiceDescription) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServiceDescription) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


