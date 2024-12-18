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

// checks if the GetUploadURLReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUploadURLReply{}

// GetUploadURLReply struct for GetUploadURLReply
type GetUploadURLReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	URL *string `json:"URL,omitempty"`
}

// NewGetUploadURLReply instantiates a new GetUploadURLReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUploadURLReply() *GetUploadURLReply {
	this := GetUploadURLReply{}
	return &this
}

// NewGetUploadURLReplyWithDefaults instantiates a new GetUploadURLReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUploadURLReplyWithDefaults() *GetUploadURLReply {
	this := GetUploadURLReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetUploadURLReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUploadURLReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetUploadURLReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *GetUploadURLReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *GetUploadURLReply) GetURL() string {
	if o == nil || IsNil(o.URL) {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUploadURLReply) GetURLOk() (*string, bool) {
	if o == nil || IsNil(o.URL) {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *GetUploadURLReply) HasURL() bool {
	if o != nil && !IsNil(o.URL) {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *GetUploadURLReply) SetURL(v string) {
	o.URL = &v
}

func (o GetUploadURLReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUploadURLReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.URL) {
		toSerialize["URL"] = o.URL
	}
	return toSerialize, nil
}

type NullableGetUploadURLReply struct {
	value *GetUploadURLReply
	isSet bool
}

func (v NullableGetUploadURLReply) Get() *GetUploadURLReply {
	return v.value
}

func (v *NullableGetUploadURLReply) Set(val *GetUploadURLReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUploadURLReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUploadURLReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUploadURLReply(val *GetUploadURLReply) *NullableGetUploadURLReply {
	return &NullableGetUploadURLReply{value: val, isSet: true}
}

func (v NullableGetUploadURLReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUploadURLReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


