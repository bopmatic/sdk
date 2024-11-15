/*
pb/sr.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiKeyDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyDescription{}

// ApiKeyDescription struct for ApiKeyDescription
type ApiKeyDescription struct {
	KeyId *string `json:"keyId,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// time the key was created expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC.
	CreateTime *string `json:"createTime,omitempty"`
	// time the key should expire expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC. A value of 0 indicates the key should never expire.
	ExpireTime *string `json:"expireTime,omitempty"`
	// time the key was last uzed expressed as the number of seconds since Jan 1, 1970 00:00:00 UTC. A value of 0 indicates the key was never used.
	LastUsed *string `json:"lastUsed,omitempty"`
}

// NewApiKeyDescription instantiates a new ApiKeyDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyDescription() *ApiKeyDescription {
	this := ApiKeyDescription{}
	return &this
}

// NewApiKeyDescriptionWithDefaults instantiates a new ApiKeyDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyDescriptionWithDefaults() *ApiKeyDescription {
	this := ApiKeyDescription{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ApiKeyDescription) SetKeyId(v string) {
	o.KeyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiKeyDescription) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyDescription) SetDescription(v string) {
	o.Description = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ApiKeyDescription) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetExpireTime() string {
	if o == nil || IsNil(o.ExpireTime) {
		var ret string
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given string and assigns it to the ExpireTime field.
func (o *ApiKeyDescription) SetExpireTime(v string) {
	o.ExpireTime = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *ApiKeyDescription) GetLastUsed() string {
	if o == nil || IsNil(o.LastUsed) {
		var ret string
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyDescription) GetLastUsedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsed) {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *ApiKeyDescription) HasLastUsed() bool {
	if o != nil && !IsNil(o.LastUsed) {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given string and assigns it to the LastUsed field.
func (o *ApiKeyDescription) SetLastUsed(v string) {
	o.LastUsed = &v
}

func (o ApiKeyDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.LastUsed) {
		toSerialize["lastUsed"] = o.LastUsed
	}
	return toSerialize, nil
}

type NullableApiKeyDescription struct {
	value *ApiKeyDescription
	isSet bool
}

func (v NullableApiKeyDescription) Get() *ApiKeyDescription {
	return v.value
}

func (v *NullableApiKeyDescription) Set(val *ApiKeyDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyDescription(val *ApiKeyDescription) *NullableApiKeyDescription {
	return &NullableApiKeyDescription{value: val, isSet: true}
}

func (v NullableApiKeyDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

