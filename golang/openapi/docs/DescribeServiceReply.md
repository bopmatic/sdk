# DescribeServiceReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Desc** | Pointer to [**ServiceDescription**](ServiceDescription.md) |  | [optional] 

## Methods

### NewDescribeServiceReply

`func NewDescribeServiceReply() *DescribeServiceReply`

NewDescribeServiceReply instantiates a new DescribeServiceReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceReplyWithDefaults

`func NewDescribeServiceReplyWithDefaults() *DescribeServiceReply`

NewDescribeServiceReplyWithDefaults instantiates a new DescribeServiceReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DescribeServiceReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DescribeServiceReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DescribeServiceReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DescribeServiceReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDesc

`func (o *DescribeServiceReply) GetDesc() ServiceDescription`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DescribeServiceReply) GetDescOk() (*ServiceDescription, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DescribeServiceReply) SetDesc(v ServiceDescription)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DescribeServiceReply) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


