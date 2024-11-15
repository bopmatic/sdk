# DescribeApiKeyReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Desc** | Pointer to [**ApiKeyDescription**](ApiKeyDescription.md) |  | [optional] 

## Methods

### NewDescribeApiKeyReply

`func NewDescribeApiKeyReply() *DescribeApiKeyReply`

NewDescribeApiKeyReply instantiates a new DescribeApiKeyReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeApiKeyReplyWithDefaults

`func NewDescribeApiKeyReplyWithDefaults() *DescribeApiKeyReply`

NewDescribeApiKeyReplyWithDefaults instantiates a new DescribeApiKeyReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DescribeApiKeyReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DescribeApiKeyReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DescribeApiKeyReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DescribeApiKeyReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDesc

`func (o *DescribeApiKeyReply) GetDesc() ApiKeyDescription`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DescribeApiKeyReply) GetDescOk() (*ApiKeyDescription, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DescribeApiKeyReply) SetDesc(v ApiKeyDescription)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DescribeApiKeyReply) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

