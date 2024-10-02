# ProjectDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Header** | Pointer to [**ProjectHeader**](ProjectHeader.md) |  | [optional] 
**State** | Pointer to [**ProjectState**](ProjectState.md) |  | [optional] [default to INVALID_PROJ_STATE]
**CreateTime** | Pointer to **string** |  | [optional] 
**ActiveDeployIds** | Pointer to **[]string** |  | [optional] 
**PendingDeployIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectDescription

`func NewProjectDescription() *ProjectDescription`

NewProjectDescription instantiates a new ProjectDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDescriptionWithDefaults

`func NewProjectDescriptionWithDefaults() *ProjectDescription`

NewProjectDescriptionWithDefaults instantiates a new ProjectDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHeader

`func (o *ProjectDescription) GetHeader() ProjectHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ProjectDescription) GetHeaderOk() (*ProjectHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ProjectDescription) SetHeader(v ProjectHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ProjectDescription) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetState

`func (o *ProjectDescription) GetState() ProjectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectDescription) GetStateOk() (*ProjectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectDescription) SetState(v ProjectState)`

SetState sets State field to given value.

### HasState

`func (o *ProjectDescription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreateTime

`func (o *ProjectDescription) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ProjectDescription) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ProjectDescription) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ProjectDescription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetActiveDeployIds

`func (o *ProjectDescription) GetActiveDeployIds() []string`

GetActiveDeployIds returns the ActiveDeployIds field if non-nil, zero value otherwise.

### GetActiveDeployIdsOk

`func (o *ProjectDescription) GetActiveDeployIdsOk() (*[]string, bool)`

GetActiveDeployIdsOk returns a tuple with the ActiveDeployIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeployIds

`func (o *ProjectDescription) SetActiveDeployIds(v []string)`

SetActiveDeployIds sets ActiveDeployIds field to given value.

### HasActiveDeployIds

`func (o *ProjectDescription) HasActiveDeployIds() bool`

HasActiveDeployIds returns a boolean if a field has been set.

### GetPendingDeployIds

`func (o *ProjectDescription) GetPendingDeployIds() []string`

GetPendingDeployIds returns the PendingDeployIds field if non-nil, zero value otherwise.

### GetPendingDeployIdsOk

`func (o *ProjectDescription) GetPendingDeployIdsOk() (*[]string, bool)`

GetPendingDeployIdsOk returns a tuple with the PendingDeployIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeployIds

`func (o *ProjectDescription) SetPendingDeployIds(v []string)`

SetPendingDeployIds sets PendingDeployIds field to given value.

### HasPendingDeployIds

`func (o *ProjectDescription) HasPendingDeployIds() bool`

HasPendingDeployIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


