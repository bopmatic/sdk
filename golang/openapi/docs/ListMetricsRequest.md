# ListMetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**MetricsScope**](MetricsScope.md) |  | [optional] [default to INVALID_METRIC_SCOPE]

## Methods

### NewListMetricsRequest

`func NewListMetricsRequest() *ListMetricsRequest`

NewListMetricsRequest instantiates a new ListMetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricsRequestWithDefaults

`func NewListMetricsRequestWithDefaults() *ListMetricsRequest`

NewListMetricsRequestWithDefaults instantiates a new ListMetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ListMetricsRequest) GetScope() MetricsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListMetricsRequest) GetScopeOk() (*MetricsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListMetricsRequest) SetScope(v MetricsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListMetricsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


