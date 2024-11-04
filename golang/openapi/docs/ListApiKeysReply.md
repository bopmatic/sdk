# ListApiKeysReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**KeyIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListApiKeysReply

`func NewListApiKeysReply() *ListApiKeysReply`

NewListApiKeysReply instantiates a new ListApiKeysReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysReplyWithDefaults

`func NewListApiKeysReplyWithDefaults() *ListApiKeysReply`

NewListApiKeysReplyWithDefaults instantiates a new ListApiKeysReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListApiKeysReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListApiKeysReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListApiKeysReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListApiKeysReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetKeyIds

`func (o *ListApiKeysReply) GetKeyIds() []string`

GetKeyIds returns the KeyIds field if non-nil, zero value otherwise.

### GetKeyIdsOk

`func (o *ListApiKeysReply) GetKeyIdsOk() (*[]string, bool)`

GetKeyIdsOk returns a tuple with the KeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIds

`func (o *ListApiKeysReply) SetKeyIds(v []string)`

SetKeyIds sets KeyIds field to given value.

### HasKeyIds

`func (o *ListApiKeysReply) HasKeyIds() bool`

HasKeyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


