# DescribeDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseHeader** | Pointer to [**DatabaseHeader**](DatabaseHeader.md) |  | [optional] 

## Methods

### NewDescribeDatabaseRequest

`func NewDescribeDatabaseRequest() *DescribeDatabaseRequest`

NewDescribeDatabaseRequest instantiates a new DescribeDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDatabaseRequestWithDefaults

`func NewDescribeDatabaseRequestWithDefaults() *DescribeDatabaseRequest`

NewDescribeDatabaseRequestWithDefaults instantiates a new DescribeDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseHeader

`func (o *DescribeDatabaseRequest) GetDatabaseHeader() DatabaseHeader`

GetDatabaseHeader returns the DatabaseHeader field if non-nil, zero value otherwise.

### GetDatabaseHeaderOk

`func (o *DescribeDatabaseRequest) GetDatabaseHeaderOk() (*DatabaseHeader, bool)`

GetDatabaseHeaderOk returns a tuple with the DatabaseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseHeader

`func (o *DescribeDatabaseRequest) SetDatabaseHeader(v DatabaseHeader)`

SetDatabaseHeader sets DatabaseHeader field to given value.

### HasDatabaseHeader

`func (o *DescribeDatabaseRequest) HasDatabaseHeader() bool`

HasDatabaseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


