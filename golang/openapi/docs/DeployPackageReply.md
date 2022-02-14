# DeployPackageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**PackageState**](PackageState.md) |  | [optional] [default to UPLOADING]

## Methods

### NewDeployPackageReply

`func NewDeployPackageReply() *DeployPackageReply`

NewDeployPackageReply instantiates a new DeployPackageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployPackageReplyWithDefaults

`func NewDeployPackageReplyWithDefaults() *DeployPackageReply`

NewDeployPackageReplyWithDefaults instantiates a new DeployPackageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DeployPackageReply) GetState() PackageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeployPackageReply) GetStateOk() (*PackageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeployPackageReply) SetState(v PackageState)`

SetState sets State field to given value.

### HasState

`func (o *DeployPackageReply) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

