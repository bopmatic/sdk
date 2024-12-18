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

// checks if the CreateEnvironmentReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEnvironmentReply{}

// CreateEnvironmentReply struct for CreateEnvironmentReply
type CreateEnvironmentReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCreateEnvironmentReply instantiates a new CreateEnvironmentReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentReply() *CreateEnvironmentReply {
	this := CreateEnvironmentReply{}
	return &this
}

// NewCreateEnvironmentReplyWithDefaults instantiates a new CreateEnvironmentReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentReplyWithDefaults() *CreateEnvironmentReply {
	this := CreateEnvironmentReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateEnvironmentReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateEnvironmentReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *CreateEnvironmentReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateEnvironmentReply) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentReply) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateEnvironmentReply) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateEnvironmentReply) SetId(v string) {
	o.Id = &v
}

func (o CreateEnvironmentReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEnvironmentReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateEnvironmentReply struct {
	value *CreateEnvironmentReply
	isSet bool
}

func (v NullableCreateEnvironmentReply) Get() *CreateEnvironmentReply {
	return v.value
}

func (v *NullableCreateEnvironmentReply) Set(val *CreateEnvironmentReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentReply(val *CreateEnvironmentReply) *NullableCreateEnvironmentReply {
	return &NullableCreateEnvironmentReply{value: val, isSet: true}
}

func (v NullableCreateEnvironmentReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


