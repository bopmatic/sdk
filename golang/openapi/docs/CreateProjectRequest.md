# CreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**ProjectHeader**](ProjectHeader.md) |  | [optional] 

## Methods

### NewCreateProjectRequest

`func NewCreateProjectRequest() *CreateProjectRequest`

NewCreateProjectRequest instantiates a new CreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectRequestWithDefaults

`func NewCreateProjectRequestWithDefaults() *CreateProjectRequest`

NewCreateProjectRequestWithDefaults instantiates a new CreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *CreateProjectRequest) GetHeader() ProjectHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CreateProjectRequest) GetHeaderOk() (*ProjectHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CreateProjectRequest) SetHeader(v ProjectHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CreateProjectRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


