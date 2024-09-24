# ServiceRunnerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ServiceRunnerStatus**](ServiceRunnerStatus.md) |  | [optional] [default to OK]
**StatusDetail** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceRunnerResult

`func NewServiceRunnerResult() *ServiceRunnerResult`

NewServiceRunnerResult instantiates a new ServiceRunnerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRunnerResultWithDefaults

`func NewServiceRunnerResultWithDefaults() *ServiceRunnerResult`

NewServiceRunnerResultWithDefaults instantiates a new ServiceRunnerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ServiceRunnerResult) GetStatus() ServiceRunnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceRunnerResult) GetStatusOk() (*ServiceRunnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceRunnerResult) SetStatus(v ServiceRunnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceRunnerResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *ServiceRunnerResult) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *ServiceRunnerResult) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *ServiceRunnerResult) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *ServiceRunnerResult) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


