# ListDatastoresReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Header** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**DatastoreNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListDatastoresReply

`func NewListDatastoresReply() *ListDatastoresReply`

NewListDatastoresReply instantiates a new ListDatastoresReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatastoresReplyWithDefaults

`func NewListDatastoresReplyWithDefaults() *ListDatastoresReply`

NewListDatastoresReplyWithDefaults instantiates a new ListDatastoresReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDatastoresReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDatastoresReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDatastoresReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDatastoresReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetHeader

`func (o *ListDatastoresReply) GetHeader() ProjEnvHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ListDatastoresReply) GetHeaderOk() (*ProjEnvHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ListDatastoresReply) SetHeader(v ProjEnvHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ListDatastoresReply) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetDatastoreNames

`func (o *ListDatastoresReply) GetDatastoreNames() []string`

GetDatastoreNames returns the DatastoreNames field if non-nil, zero value otherwise.

### GetDatastoreNamesOk

`func (o *ListDatastoresReply) GetDatastoreNamesOk() (*[]string, bool)`

GetDatastoreNamesOk returns a tuple with the DatastoreNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreNames

`func (o *ListDatastoresReply) SetDatastoreNames(v []string)`

SetDatastoreNames sets DatastoreNames field to given value.

### HasDatastoreNames

`func (o *ListDatastoresReply) HasDatastoreNames() bool`

HasDatastoreNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


