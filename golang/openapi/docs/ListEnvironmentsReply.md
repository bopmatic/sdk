# ListEnvironmentsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListEnvironmentsReply

`func NewListEnvironmentsReply() *ListEnvironmentsReply`

NewListEnvironmentsReply instantiates a new ListEnvironmentsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentsReplyWithDefaults

`func NewListEnvironmentsReplyWithDefaults() *ListEnvironmentsReply`

NewListEnvironmentsReplyWithDefaults instantiates a new ListEnvironmentsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListEnvironmentsReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListEnvironmentsReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListEnvironmentsReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListEnvironmentsReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetIds

`func (o *ListEnvironmentsReply) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListEnvironmentsReply) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListEnvironmentsReply) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListEnvironmentsReply) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


