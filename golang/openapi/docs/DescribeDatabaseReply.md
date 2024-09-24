# DescribeDatabaseReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Desc** | Pointer to [**DatabaseDescription**](DatabaseDescription.md) |  | [optional] 

## Methods

### NewDescribeDatabaseReply

`func NewDescribeDatabaseReply() *DescribeDatabaseReply`

NewDescribeDatabaseReply instantiates a new DescribeDatabaseReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDatabaseReplyWithDefaults

`func NewDescribeDatabaseReplyWithDefaults() *DescribeDatabaseReply`

NewDescribeDatabaseReplyWithDefaults instantiates a new DescribeDatabaseReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DescribeDatabaseReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DescribeDatabaseReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DescribeDatabaseReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DescribeDatabaseReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDesc

`func (o *DescribeDatabaseReply) GetDesc() DatabaseDescription`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DescribeDatabaseReply) GetDescOk() (*DatabaseDescription, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DescribeDatabaseReply) SetDesc(v DatabaseDescription)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DescribeDatabaseReply) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


