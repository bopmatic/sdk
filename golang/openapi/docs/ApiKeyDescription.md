# ApiKeyDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** | time the key was created expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC. | [optional] 
**ExpireTime** | Pointer to **string** | time the key should expire expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC. A value of 0 indicates the key should never expire. | [optional] 
**LastUsed** | Pointer to **string** | time the key was last uzed expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC. A value of 0 indicates the key was never used. | [optional] 

## Methods

### NewApiKeyDescription

`func NewApiKeyDescription() *ApiKeyDescription`

NewApiKeyDescription instantiates a new ApiKeyDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyDescriptionWithDefaults

`func NewApiKeyDescriptionWithDefaults() *ApiKeyDescription`

NewApiKeyDescriptionWithDefaults instantiates a new ApiKeyDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ApiKeyDescription) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ApiKeyDescription) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ApiKeyDescription) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ApiKeyDescription) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetName

`func (o *ApiKeyDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiKeyDescription) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiKeyDescription) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiKeyDescription) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiKeyDescription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *ApiKeyDescription) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *ApiKeyDescription) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *ApiKeyDescription) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *ApiKeyDescription) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLastUsed

`func (o *ApiKeyDescription) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ApiKeyDescription) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ApiKeyDescription) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *ApiKeyDescription) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


