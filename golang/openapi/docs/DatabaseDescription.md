# DatabaseDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseHeader** | Pointer to [**DatabaseHeader**](DatabaseHeader.md) |  | [optional] 
**Tables** | Pointer to [**[]DatabaseTableDescription**](DatabaseTableDescription.md) |  | [optional] 
**ServiceNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDatabaseDescription

`func NewDatabaseDescription() *DatabaseDescription`

NewDatabaseDescription instantiates a new DatabaseDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseDescriptionWithDefaults

`func NewDatabaseDescriptionWithDefaults() *DatabaseDescription`

NewDatabaseDescriptionWithDefaults instantiates a new DatabaseDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseHeader

`func (o *DatabaseDescription) GetDatabaseHeader() DatabaseHeader`

GetDatabaseHeader returns the DatabaseHeader field if non-nil, zero value otherwise.

### GetDatabaseHeaderOk

`func (o *DatabaseDescription) GetDatabaseHeaderOk() (*DatabaseHeader, bool)`

GetDatabaseHeaderOk returns a tuple with the DatabaseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseHeader

`func (o *DatabaseDescription) SetDatabaseHeader(v DatabaseHeader)`

SetDatabaseHeader sets DatabaseHeader field to given value.

### HasDatabaseHeader

`func (o *DatabaseDescription) HasDatabaseHeader() bool`

HasDatabaseHeader returns a boolean if a field has been set.

### GetTables

`func (o *DatabaseDescription) GetTables() []DatabaseTableDescription`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DatabaseDescription) GetTablesOk() (*[]DatabaseTableDescription, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DatabaseDescription) SetTables(v []DatabaseTableDescription)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DatabaseDescription) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetServiceNames

`func (o *DatabaseDescription) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *DatabaseDescription) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *DatabaseDescription) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *DatabaseDescription) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


