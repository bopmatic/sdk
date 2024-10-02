# GetMetricSamplesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjId** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**MetricsScope**](MetricsScope.md) |  | [optional] [default to INVALID_METRIC_SCOPE]
**ScopeQualifier** | Pointer to **string** |  | [optional] 
**MetricNames** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** | latest metric to retrieve expressed as the number of | [optional] 
**Format** | Pointer to [**MetricsFormat**](MetricsFormat.md) |  | [optional] [default to INVALID_METRIC_FORMAT]

## Methods

### NewGetMetricSamplesRequest

`func NewGetMetricSamplesRequest() *GetMetricSamplesRequest`

NewGetMetricSamplesRequest instantiates a new GetMetricSamplesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricSamplesRequestWithDefaults

`func NewGetMetricSamplesRequestWithDefaults() *GetMetricSamplesRequest`

NewGetMetricSamplesRequestWithDefaults instantiates a new GetMetricSamplesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjId

`func (o *GetMetricSamplesRequest) GetProjId() string`

GetProjId returns the ProjId field if non-nil, zero value otherwise.

### GetProjIdOk

`func (o *GetMetricSamplesRequest) GetProjIdOk() (*string, bool)`

GetProjIdOk returns a tuple with the ProjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjId

`func (o *GetMetricSamplesRequest) SetProjId(v string)`

SetProjId sets ProjId field to given value.

### HasProjId

`func (o *GetMetricSamplesRequest) HasProjId() bool`

HasProjId returns a boolean if a field has been set.

### GetEnvId

`func (o *GetMetricSamplesRequest) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *GetMetricSamplesRequest) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *GetMetricSamplesRequest) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *GetMetricSamplesRequest) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetScope

`func (o *GetMetricSamplesRequest) GetScope() MetricsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetMetricSamplesRequest) GetScopeOk() (*MetricsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetMetricSamplesRequest) SetScope(v MetricsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetMetricSamplesRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeQualifier

`func (o *GetMetricSamplesRequest) GetScopeQualifier() string`

GetScopeQualifier returns the ScopeQualifier field if non-nil, zero value otherwise.

### GetScopeQualifierOk

`func (o *GetMetricSamplesRequest) GetScopeQualifierOk() (*string, bool)`

GetScopeQualifierOk returns a tuple with the ScopeQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeQualifier

`func (o *GetMetricSamplesRequest) SetScopeQualifier(v string)`

SetScopeQualifier sets ScopeQualifier field to given value.

### HasScopeQualifier

`func (o *GetMetricSamplesRequest) HasScopeQualifier() bool`

HasScopeQualifier returns a boolean if a field has been set.

### GetMetricNames

`func (o *GetMetricSamplesRequest) GetMetricNames() []string`

GetMetricNames returns the MetricNames field if non-nil, zero value otherwise.

### GetMetricNamesOk

`func (o *GetMetricSamplesRequest) GetMetricNamesOk() (*[]string, bool)`

GetMetricNamesOk returns a tuple with the MetricNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNames

`func (o *GetMetricSamplesRequest) SetMetricNames(v []string)`

SetMetricNames sets MetricNames field to given value.

### HasMetricNames

`func (o *GetMetricSamplesRequest) HasMetricNames() bool`

HasMetricNames returns a boolean if a field has been set.

### GetStartTime

`func (o *GetMetricSamplesRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetMetricSamplesRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetMetricSamplesRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetMetricSamplesRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GetMetricSamplesRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetMetricSamplesRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetMetricSamplesRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetMetricSamplesRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFormat

`func (o *GetMetricSamplesRequest) GetFormat() MetricsFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GetMetricSamplesRequest) GetFormatOk() (*MetricsFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GetMetricSamplesRequest) SetFormat(v MetricsFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GetMetricSamplesRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


