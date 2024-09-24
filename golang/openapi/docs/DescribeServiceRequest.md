# DescribeServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcHeader** | Pointer to [**ServiceHeader**](ServiceHeader.md) |  | [optional] 

## Methods

### NewDescribeServiceRequest

`func NewDescribeServiceRequest() *DescribeServiceRequest`

NewDescribeServiceRequest instantiates a new DescribeServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceRequestWithDefaults

`func NewDescribeServiceRequestWithDefaults() *DescribeServiceRequest`

NewDescribeServiceRequestWithDefaults instantiates a new DescribeServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcHeader

`func (o *DescribeServiceRequest) GetSvcHeader() ServiceHeader`

GetSvcHeader returns the SvcHeader field if non-nil, zero value otherwise.

### GetSvcHeaderOk

`func (o *DescribeServiceRequest) GetSvcHeaderOk() (*ServiceHeader, bool)`

GetSvcHeaderOk returns a tuple with the SvcHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcHeader

`func (o *DescribeServiceRequest) SetSvcHeader(v ServiceHeader)`

SetSvcHeader sets SvcHeader field to given value.

### HasSvcHeader

`func (o *DescribeServiceRequest) HasSvcHeader() bool`

HasSvcHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


