# UploadPackageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**PkgId** | Pointer to **string** |  | [optional] 

## Methods

### NewUploadPackageReply

`func NewUploadPackageReply() *UploadPackageReply`

NewUploadPackageReply instantiates a new UploadPackageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPackageReplyWithDefaults

`func NewUploadPackageReplyWithDefaults() *UploadPackageReply`

NewUploadPackageReplyWithDefaults instantiates a new UploadPackageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UploadPackageReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UploadPackageReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UploadPackageReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *UploadPackageReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetPkgId

`func (o *UploadPackageReply) GetPkgId() string`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *UploadPackageReply) GetPkgIdOk() (*string, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *UploadPackageReply) SetPkgId(v string)`

SetPkgId sets PkgId field to given value.

### HasPkgId

`func (o *UploadPackageReply) HasPkgId() bool`

HasPkgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


