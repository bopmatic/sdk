# ListServicesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Header** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**ServiceNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListServicesReply

`func NewListServicesReply() *ListServicesReply`

NewListServicesReply instantiates a new ListServicesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicesReplyWithDefaults

`func NewListServicesReplyWithDefaults() *ListServicesReply`

NewListServicesReplyWithDefaults instantiates a new ListServicesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListServicesReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListServicesReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListServicesReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListServicesReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetHeader

`func (o *ListServicesReply) GetHeader() ProjEnvHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ListServicesReply) GetHeaderOk() (*ProjEnvHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ListServicesReply) SetHeader(v ProjEnvHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ListServicesReply) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetServiceNames

`func (o *ListServicesReply) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *ListServicesReply) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *ListServicesReply) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *ListServicesReply) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


