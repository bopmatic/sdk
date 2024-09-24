# DatabaseHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjEnvHeader** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabaseHeader

`func NewDatabaseHeader() *DatabaseHeader`

NewDatabaseHeader instantiates a new DatabaseHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseHeaderWithDefaults

`func NewDatabaseHeaderWithDefaults() *DatabaseHeader`

NewDatabaseHeaderWithDefaults instantiates a new DatabaseHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjEnvHeader

`func (o *DatabaseHeader) GetProjEnvHeader() ProjEnvHeader`

GetProjEnvHeader returns the ProjEnvHeader field if non-nil, zero value otherwise.

### GetProjEnvHeaderOk

`func (o *DatabaseHeader) GetProjEnvHeaderOk() (*ProjEnvHeader, bool)`

GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjEnvHeader

`func (o *DatabaseHeader) SetProjEnvHeader(v ProjEnvHeader)`

SetProjEnvHeader sets ProjEnvHeader field to given value.

### HasProjEnvHeader

`func (o *DatabaseHeader) HasProjEnvHeader() bool`

HasProjEnvHeader returns a boolean if a field has been set.

### GetDatabaseName

`func (o *DatabaseHeader) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DatabaseHeader) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DatabaseHeader) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *DatabaseHeader) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


