# DeploymentHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkgId** | Pointer to **string** |  | [optional] 
**ProjId** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DeploymentType**](DeploymentType.md) |  | [optional] [default to NEW_PACKAGE]
**Initiator** | Pointer to [**DeploymentInitiator**](DeploymentInitiator.md) |  | [optional] [default to CUSTOMER]

## Methods

### NewDeploymentHeader

`func NewDeploymentHeader() *DeploymentHeader`

NewDeploymentHeader instantiates a new DeploymentHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHeaderWithDefaults

`func NewDeploymentHeaderWithDefaults() *DeploymentHeader`

NewDeploymentHeaderWithDefaults instantiates a new DeploymentHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkgId

`func (o *DeploymentHeader) GetPkgId() string`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *DeploymentHeader) GetPkgIdOk() (*string, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *DeploymentHeader) SetPkgId(v string)`

SetPkgId sets PkgId field to given value.

### HasPkgId

`func (o *DeploymentHeader) HasPkgId() bool`

HasPkgId returns a boolean if a field has been set.

### GetProjId

`func (o *DeploymentHeader) GetProjId() string`

GetProjId returns the ProjId field if non-nil, zero value otherwise.

### GetProjIdOk

`func (o *DeploymentHeader) GetProjIdOk() (*string, bool)`

GetProjIdOk returns a tuple with the ProjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjId

`func (o *DeploymentHeader) SetProjId(v string)`

SetProjId sets ProjId field to given value.

### HasProjId

`func (o *DeploymentHeader) HasProjId() bool`

HasProjId returns a boolean if a field has been set.

### GetEnvId

`func (o *DeploymentHeader) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *DeploymentHeader) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *DeploymentHeader) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *DeploymentHeader) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetType

`func (o *DeploymentHeader) GetType() DeploymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentHeader) GetTypeOk() (*DeploymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentHeader) SetType(v DeploymentType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentHeader) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInitiator

`func (o *DeploymentHeader) GetInitiator() DeploymentInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *DeploymentHeader) GetInitiatorOk() (*DeploymentInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *DeploymentHeader) SetInitiator(v DeploymentInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *DeploymentHeader) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


