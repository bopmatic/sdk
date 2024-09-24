# ListDeploymentsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListDeploymentsReply

`func NewListDeploymentsReply() *ListDeploymentsReply`

NewListDeploymentsReply instantiates a new ListDeploymentsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentsReplyWithDefaults

`func NewListDeploymentsReplyWithDefaults() *ListDeploymentsReply`

NewListDeploymentsReplyWithDefaults instantiates a new ListDeploymentsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDeploymentsReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDeploymentsReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDeploymentsReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDeploymentsReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetIds

`func (o *ListDeploymentsReply) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListDeploymentsReply) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListDeploymentsReply) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListDeploymentsReply) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


