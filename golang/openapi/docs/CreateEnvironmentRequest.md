# CreateEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**EnvironmentHeader**](EnvironmentHeader.md) |  | [optional] 

## Methods

### NewCreateEnvironmentRequest

`func NewCreateEnvironmentRequest() *CreateEnvironmentRequest`

NewCreateEnvironmentRequest instantiates a new CreateEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentRequestWithDefaults

`func NewCreateEnvironmentRequestWithDefaults() *CreateEnvironmentRequest`

NewCreateEnvironmentRequestWithDefaults instantiates a new CreateEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CreateEnvironmentRequest) GetHeader() EnvironmentHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CreateEnvironmentRequest) GetHeaderOk() (*EnvironmentHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CreateEnvironmentRequest) SetHeader(v EnvironmentHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CreateEnvironmentRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


