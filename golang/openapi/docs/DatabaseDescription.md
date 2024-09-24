# DatabaseDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseHeader** | Pointer to [**DatabaseHeader**](DatabaseHeader.md) |  | [optional] 
**TableNames** | Pointer to **[]string** |  | [optional] 

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

### GetTableNames

`func (o *DatabaseDescription) GetTableNames() []string`

GetTableNames returns the TableNames field if non-nil, zero value otherwise.

### GetTableNamesOk

`func (o *DatabaseDescription) GetTableNamesOk() (*[]string, bool)`

GetTableNamesOk returns a tuple with the TableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNames

`func (o *DatabaseDescription) SetTableNames(v []string)`

SetTableNames sets TableNames field to given value.

### HasTableNames

`func (o *DatabaseDescription) HasTableNames() bool`

HasTableNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


