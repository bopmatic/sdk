# DescribeEnvironmentReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Desc** | Pointer to [**EnvironmentDescription**](EnvironmentDescription.md) |  | [optional] 

## Methods

### NewDescribeEnvironmentReply

`func NewDescribeEnvironmentReply() *DescribeEnvironmentReply`

NewDescribeEnvironmentReply instantiates a new DescribeEnvironmentReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeEnvironmentReplyWithDefaults

`func NewDescribeEnvironmentReplyWithDefaults() *DescribeEnvironmentReply`

NewDescribeEnvironmentReplyWithDefaults instantiates a new DescribeEnvironmentReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DescribeEnvironmentReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DescribeEnvironmentReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DescribeEnvironmentReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DescribeEnvironmentReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDesc

`func (o *DescribeEnvironmentReply) GetDesc() EnvironmentDescription`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DescribeEnvironmentReply) GetDescOk() (*EnvironmentDescription, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DescribeEnvironmentReply) SetDesc(v EnvironmentDescription)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DescribeEnvironmentReply) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


