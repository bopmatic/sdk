# DatastoreHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjEnvHeader** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**DatastoreName** | Pointer to **string** |  | [optional] 

## Methods

### NewDatastoreHeader

`func NewDatastoreHeader() *DatastoreHeader`

NewDatastoreHeader instantiates a new DatastoreHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreHeaderWithDefaults

`func NewDatastoreHeaderWithDefaults() *DatastoreHeader`

NewDatastoreHeaderWithDefaults instantiates a new DatastoreHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjEnvHeader

`func (o *DatastoreHeader) GetProjEnvHeader() ProjEnvHeader`

GetProjEnvHeader returns the ProjEnvHeader field if non-nil, zero value otherwise.

### GetProjEnvHeaderOk

`func (o *DatastoreHeader) GetProjEnvHeaderOk() (*ProjEnvHeader, bool)`

GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjEnvHeader

`func (o *DatastoreHeader) SetProjEnvHeader(v ProjEnvHeader)`

SetProjEnvHeader sets ProjEnvHeader field to given value.

### HasProjEnvHeader

`func (o *DatastoreHeader) HasProjEnvHeader() bool`

HasProjEnvHeader returns a boolean if a field has been set.

### GetDatastoreName

`func (o *DatastoreHeader) GetDatastoreName() string`

GetDatastoreName returns the DatastoreName field if non-nil, zero value otherwise.

### GetDatastoreNameOk

`func (o *DatastoreHeader) GetDatastoreNameOk() (*string, bool)`

GetDatastoreNameOk returns a tuple with the DatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreName

`func (o *DatastoreHeader) SetDatastoreName(v string)`

SetDatastoreName sets DatastoreName field to given value.

### HasDatastoreName

`func (o *DatastoreHeader) HasDatastoreName() bool`

HasDatastoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


