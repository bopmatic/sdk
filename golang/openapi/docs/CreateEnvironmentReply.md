# CreateEnvironmentReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateEnvironmentReply

`func NewCreateEnvironmentReply() *CreateEnvironmentReply`

NewCreateEnvironmentReply instantiates a new CreateEnvironmentReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentReplyWithDefaults

`func NewCreateEnvironmentReplyWithDefaults() *CreateEnvironmentReply`

NewCreateEnvironmentReplyWithDefaults instantiates a new CreateEnvironmentReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateEnvironmentReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateEnvironmentReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateEnvironmentReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateEnvironmentReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetId

`func (o *CreateEnvironmentReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEnvironmentReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEnvironmentReply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEnvironmentReply) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


