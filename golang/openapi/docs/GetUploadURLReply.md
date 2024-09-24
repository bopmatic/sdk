# GetUploadURLReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUploadURLReply

`func NewGetUploadURLReply() *GetUploadURLReply`

NewGetUploadURLReply instantiates a new GetUploadURLReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUploadURLReplyWithDefaults

`func NewGetUploadURLReplyWithDefaults() *GetUploadURLReply`

NewGetUploadURLReplyWithDefaults instantiates a new GetUploadURLReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetUploadURLReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUploadURLReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUploadURLReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetUploadURLReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetURL

`func (o *GetUploadURLReply) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *GetUploadURLReply) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *GetUploadURLReply) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *GetUploadURLReply) HasURL() bool`

HasURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


