# DeletePackageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**State** | Pointer to [**PackageState**](PackageState.md) |  | [optional] [default to INVALID_PKG_STATE]

## Methods

### NewDeletePackageReply

`func NewDeletePackageReply() *DeletePackageReply`

NewDeletePackageReply instantiates a new DeletePackageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePackageReplyWithDefaults

`func NewDeletePackageReplyWithDefaults() *DeletePackageReply`

NewDeletePackageReplyWithDefaults instantiates a new DeletePackageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeletePackageReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeletePackageReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeletePackageReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DeletePackageReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetState

`func (o *DeletePackageReply) GetState() PackageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeletePackageReply) GetStateOk() (*PackageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeletePackageReply) SetState(v PackageState)`

SetState sets State field to given value.

### HasState

`func (o *DeletePackageReply) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


