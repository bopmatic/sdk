# ListDatabasesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Header** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**DatabaseNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListDatabasesReply

`func NewListDatabasesReply() *ListDatabasesReply`

NewListDatabasesReply instantiates a new ListDatabasesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatabasesReplyWithDefaults

`func NewListDatabasesReplyWithDefaults() *ListDatabasesReply`

NewListDatabasesReplyWithDefaults instantiates a new ListDatabasesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDatabasesReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDatabasesReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDatabasesReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDatabasesReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetHeader

`func (o *ListDatabasesReply) GetHeader() ProjEnvHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ListDatabasesReply) GetHeaderOk() (*ProjEnvHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ListDatabasesReply) SetHeader(v ProjEnvHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ListDatabasesReply) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetDatabaseNames

`func (o *ListDatabasesReply) GetDatabaseNames() []string`

GetDatabaseNames returns the DatabaseNames field if non-nil, zero value otherwise.

### GetDatabaseNamesOk

`func (o *ListDatabasesReply) GetDatabaseNamesOk() (*[]string, bool)`

GetDatabaseNamesOk returns a tuple with the DatabaseNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseNames

`func (o *ListDatabasesReply) SetDatabaseNames(v []string)`

SetDatabaseNames sets DatabaseNames field to given value.

### HasDatabaseNames

`func (o *ListDatabasesReply) HasDatabaseNames() bool`

HasDatabaseNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


