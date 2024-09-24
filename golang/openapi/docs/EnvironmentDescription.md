# EnvironmentDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Header** | Pointer to [**EnvironmentHeader**](EnvironmentHeader.md) |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**ActiveDeployIds** | Pointer to **[]string** |  | [optional] 
**PendingDeployIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnvironmentDescription

`func NewEnvironmentDescription() *EnvironmentDescription`

NewEnvironmentDescription instantiates a new EnvironmentDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDescriptionWithDefaults

`func NewEnvironmentDescriptionWithDefaults() *EnvironmentDescription`

NewEnvironmentDescriptionWithDefaults instantiates a new EnvironmentDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHeader

`func (o *EnvironmentDescription) GetHeader() EnvironmentHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *EnvironmentDescription) GetHeaderOk() (*EnvironmentHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *EnvironmentDescription) SetHeader(v EnvironmentHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *EnvironmentDescription) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetCreateTime

`func (o *EnvironmentDescription) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EnvironmentDescription) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EnvironmentDescription) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EnvironmentDescription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetActiveDeployIds

`func (o *EnvironmentDescription) GetActiveDeployIds() []string`

GetActiveDeployIds returns the ActiveDeployIds field if non-nil, zero value otherwise.

### GetActiveDeployIdsOk

`func (o *EnvironmentDescription) GetActiveDeployIdsOk() (*[]string, bool)`

GetActiveDeployIdsOk returns a tuple with the ActiveDeployIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeployIds

`func (o *EnvironmentDescription) SetActiveDeployIds(v []string)`

SetActiveDeployIds sets ActiveDeployIds field to given value.

### HasActiveDeployIds

`func (o *EnvironmentDescription) HasActiveDeployIds() bool`

HasActiveDeployIds returns a boolean if a field has been set.

### GetPendingDeployIds

`func (o *EnvironmentDescription) GetPendingDeployIds() []string`

GetPendingDeployIds returns the PendingDeployIds field if non-nil, zero value otherwise.

### GetPendingDeployIdsOk

`func (o *EnvironmentDescription) GetPendingDeployIdsOk() (*[]string, bool)`

GetPendingDeployIdsOk returns a tuple with the PendingDeployIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeployIds

`func (o *EnvironmentDescription) SetPendingDeployIds(v []string)`

SetPendingDeployIds sets PendingDeployIds field to given value.

### HasPendingDeployIds

`func (o *EnvironmentDescription) HasPendingDeployIds() bool`

HasPendingDeployIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


