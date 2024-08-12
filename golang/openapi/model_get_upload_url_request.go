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

// checks if the GetUploadURLRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUploadURLRequest{}

// GetUploadURLRequest struct for GetUploadURLRequest
type GetUploadURLRequest struct {
	Key *string `json:"key,omitempty"`
}

// NewGetUploadURLRequest instantiates a new GetUploadURLRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUploadURLRequest() *GetUploadURLRequest {
	this := GetUploadURLRequest{}
	return &this
}

// NewGetUploadURLRequestWithDefaults instantiates a new GetUploadURLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUploadURLRequestWithDefaults() *GetUploadURLRequest {
	this := GetUploadURLRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GetUploadURLRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUploadURLRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GetUploadURLRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GetUploadURLRequest) SetKey(v string) {
	o.Key = &v
}

func (o GetUploadURLRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUploadURLRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableGetUploadURLRequest struct {
	value *GetUploadURLRequest
	isSet bool
}

func (v NullableGetUploadURLRequest) Get() *GetUploadURLRequest {
	return v.value
}

func (v *NullableGetUploadURLRequest) Set(val *GetUploadURLRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUploadURLRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUploadURLRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUploadURLRequest(val *GetUploadURLRequest) *NullableGetUploadURLRequest {
	return &NullableGetUploadURLRequest{value: val, isSet: true}
}

func (v NullableGetUploadURLRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUploadURLRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


