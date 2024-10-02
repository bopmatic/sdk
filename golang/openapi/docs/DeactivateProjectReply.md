# DeactivateProjectReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**DeployId** | Pointer to **string** |  | [optional] 

## Methods

### NewDeactivateProjectReply

`func NewDeactivateProjectReply() *DeactivateProjectReply`

NewDeactivateProjectReply instantiates a new DeactivateProjectReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateProjectReplyWithDefaults

`func NewDeactivateProjectReplyWithDefaults() *DeactivateProjectReply`

NewDeactivateProjectReplyWithDefaults instantiates a new DeactivateProjectReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeactivateProjectReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeactivateProjectReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeactivateProjectReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DeactivateProjectReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDeployId

`func (o *DeactivateProjectReply) GetDeployId() string`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *DeactivateProjectReply) GetDeployIdOk() (*string, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *DeactivateProjectReply) SetDeployId(v string)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *DeactivateProjectReply) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


