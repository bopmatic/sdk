# GetLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjId** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** | latest log message to retrieve expressed as the number of | [optional] 

## Methods

### NewGetLogsRequest

`func NewGetLogsRequest() *GetLogsRequest`

NewGetLogsRequest instantiates a new GetLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogsRequestWithDefaults

`func NewGetLogsRequestWithDefaults() *GetLogsRequest`

NewGetLogsRequestWithDefaults instantiates a new GetLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjId

`func (o *GetLogsRequest) GetProjId() string`

GetProjId returns the ProjId field if non-nil, zero value otherwise.

### GetProjIdOk

`func (o *GetLogsRequest) GetProjIdOk() (*string, bool)`

GetProjIdOk returns a tuple with the ProjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjId

`func (o *GetLogsRequest) SetProjId(v string)`

SetProjId sets ProjId field to given value.

### HasProjId

`func (o *GetLogsRequest) HasProjId() bool`

HasProjId returns a boolean if a field has been set.

### GetEnvId

`func (o *GetLogsRequest) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *GetLogsRequest) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *GetLogsRequest) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *GetLogsRequest) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetServiceName

`func (o *GetLogsRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GetLogsRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GetLogsRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *GetLogsRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStartTime

`func (o *GetLogsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetLogsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetLogsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetLogsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GetLogsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetLogsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetLogsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetLogsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


