# GetMetricSamplesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**MetricBuf** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMetricSamplesReply

`func NewGetMetricSamplesReply() *GetMetricSamplesReply`

NewGetMetricSamplesReply instantiates a new GetMetricSamplesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricSamplesReplyWithDefaults

`func NewGetMetricSamplesReplyWithDefaults() *GetMetricSamplesReply`

NewGetMetricSamplesReplyWithDefaults instantiates a new GetMetricSamplesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetMetricSamplesReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMetricSamplesReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMetricSamplesReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMetricSamplesReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetMetricBuf

`func (o *GetMetricSamplesReply) GetMetricBuf() string`

GetMetricBuf returns the MetricBuf field if non-nil, zero value otherwise.

### GetMetricBufOk

`func (o *GetMetricSamplesReply) GetMetricBufOk() (*string, bool)`

GetMetricBufOk returns a tuple with the MetricBuf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricBuf

`func (o *GetMetricSamplesReply) SetMetricBuf(v string)`

SetMetricBuf sets MetricBuf field to given value.

### HasMetricBuf

`func (o *GetMetricSamplesReply) HasMetricBuf() bool`

HasMetricBuf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


