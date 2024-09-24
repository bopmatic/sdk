# DescribeSiteReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**ProjEnvHeader** | Pointer to [**ProjEnvHeader**](ProjEnvHeader.md) |  | [optional] 
**SiteEndpoint** | Pointer to **string** |  | [optional] 

## Methods

### NewDescribeSiteReply

`func NewDescribeSiteReply() *DescribeSiteReply`

NewDescribeSiteReply instantiates a new DescribeSiteReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSiteReplyWithDefaults

`func NewDescribeSiteReplyWithDefaults() *DescribeSiteReply`

NewDescribeSiteReplyWithDefaults instantiates a new DescribeSiteReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DescribeSiteReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DescribeSiteReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DescribeSiteReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DescribeSiteReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetProjEnvHeader

`func (o *DescribeSiteReply) GetProjEnvHeader() ProjEnvHeader`

GetProjEnvHeader returns the ProjEnvHeader field if non-nil, zero value otherwise.

### GetProjEnvHeaderOk

`func (o *DescribeSiteReply) GetProjEnvHeaderOk() (*ProjEnvHeader, bool)`

GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjEnvHeader

`func (o *DescribeSiteReply) SetProjEnvHeader(v ProjEnvHeader)`

SetProjEnvHeader sets ProjEnvHeader field to given value.

### HasProjEnvHeader

`func (o *DescribeSiteReply) HasProjEnvHeader() bool`

HasProjEnvHeader returns a boolean if a field has been set.

### GetSiteEndpoint

`func (o *DescribeSiteReply) GetSiteEndpoint() string`

GetSiteEndpoint returns the SiteEndpoint field if non-nil, zero value otherwise.

### GetSiteEndpointOk

`func (o *DescribeSiteReply) GetSiteEndpointOk() (*string, bool)`

GetSiteEndpointOk returns a tuple with the SiteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteEndpoint

`func (o *DescribeSiteReply) SetSiteEndpoint(v string)`

SetSiteEndpoint sets SiteEndpoint field to given value.

### HasSiteEndpoint

`func (o *DescribeSiteReply) HasSiteEndpoint() bool`

HasSiteEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


