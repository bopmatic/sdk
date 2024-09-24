# DeploymentDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Header** | Pointer to [**DeploymentHeader**](DeploymentHeader.md) |  | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] [default to CREATED]
**StateDetail** | Pointer to [**DeploymentStateDetail**](DeploymentStateDetail.md) |  | [optional] [default to NONE]
**CreateTime** | Pointer to **string** |  | [optional] 
**ValidationStartTime** | Pointer to **string** |  | [optional] 
**BuildStartTime** | Pointer to **string** |  | [optional] 
**DeployStartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentDescription

`func NewDeploymentDescription() *DeploymentDescription`

NewDeploymentDescription instantiates a new DeploymentDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDescriptionWithDefaults

`func NewDeploymentDescriptionWithDefaults() *DeploymentDescription`

NewDeploymentDescriptionWithDefaults instantiates a new DeploymentDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHeader

`func (o *DeploymentDescription) GetHeader() DeploymentHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *DeploymentDescription) GetHeaderOk() (*DeploymentHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *DeploymentDescription) SetHeader(v DeploymentHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *DeploymentDescription) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetState

`func (o *DeploymentDescription) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentDescription) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentDescription) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentDescription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDetail

`func (o *DeploymentDescription) GetStateDetail() DeploymentStateDetail`

GetStateDetail returns the StateDetail field if non-nil, zero value otherwise.

### GetStateDetailOk

`func (o *DeploymentDescription) GetStateDetailOk() (*DeploymentStateDetail, bool)`

GetStateDetailOk returns a tuple with the StateDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDetail

`func (o *DeploymentDescription) SetStateDetail(v DeploymentStateDetail)`

SetStateDetail sets StateDetail field to given value.

### HasStateDetail

`func (o *DeploymentDescription) HasStateDetail() bool`

HasStateDetail returns a boolean if a field has been set.

### GetCreateTime

`func (o *DeploymentDescription) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DeploymentDescription) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DeploymentDescription) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DeploymentDescription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetValidationStartTime

`func (o *DeploymentDescription) GetValidationStartTime() string`

GetValidationStartTime returns the ValidationStartTime field if non-nil, zero value otherwise.

### GetValidationStartTimeOk

`func (o *DeploymentDescription) GetValidationStartTimeOk() (*string, bool)`

GetValidationStartTimeOk returns a tuple with the ValidationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStartTime

`func (o *DeploymentDescription) SetValidationStartTime(v string)`

SetValidationStartTime sets ValidationStartTime field to given value.

### HasValidationStartTime

`func (o *DeploymentDescription) HasValidationStartTime() bool`

HasValidationStartTime returns a boolean if a field has been set.

### GetBuildStartTime

`func (o *DeploymentDescription) GetBuildStartTime() string`

GetBuildStartTime returns the BuildStartTime field if non-nil, zero value otherwise.

### GetBuildStartTimeOk

`func (o *DeploymentDescription) GetBuildStartTimeOk() (*string, bool)`

GetBuildStartTimeOk returns a tuple with the BuildStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartTime

`func (o *DeploymentDescription) SetBuildStartTime(v string)`

SetBuildStartTime sets BuildStartTime field to given value.

### HasBuildStartTime

`func (o *DeploymentDescription) HasBuildStartTime() bool`

HasBuildStartTime returns a boolean if a field has been set.

### GetDeployStartTime

`func (o *DeploymentDescription) GetDeployStartTime() string`

GetDeployStartTime returns the DeployStartTime field if non-nil, zero value otherwise.

### GetDeployStartTimeOk

`func (o *DeploymentDescription) GetDeployStartTimeOk() (*string, bool)`

GetDeployStartTimeOk returns a tuple with the DeployStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStartTime

`func (o *DeploymentDescription) SetDeployStartTime(v string)`

SetDeployStartTime sets DeployStartTime field to given value.

### HasDeployStartTime

`func (o *DeploymentDescription) HasDeployStartTime() bool`

HasDeployStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DeploymentDescription) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DeploymentDescription) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DeploymentDescription) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DeploymentDescription) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


