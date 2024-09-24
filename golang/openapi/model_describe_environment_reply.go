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

// checks if the DescribeEnvironmentReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeEnvironmentReply{}

// DescribeEnvironmentReply struct for DescribeEnvironmentReply
type DescribeEnvironmentReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Desc *EnvironmentDescription `json:"desc,omitempty"`
}

// NewDescribeEnvironmentReply instantiates a new DescribeEnvironmentReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeEnvironmentReply() *DescribeEnvironmentReply {
	this := DescribeEnvironmentReply{}
	return &this
}

// NewDescribeEnvironmentReplyWithDefaults instantiates a new DescribeEnvironmentReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeEnvironmentReplyWithDefaults() *DescribeEnvironmentReply {
	this := DescribeEnvironmentReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DescribeEnvironmentReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeEnvironmentReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DescribeEnvironmentReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *DescribeEnvironmentReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DescribeEnvironmentReply) GetDesc() EnvironmentDescription {
	if o == nil || IsNil(o.Desc) {
		var ret EnvironmentDescription
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeEnvironmentReply) GetDescOk() (*EnvironmentDescription, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DescribeEnvironmentReply) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given EnvironmentDescription and assigns it to the Desc field.
func (o *DescribeEnvironmentReply) SetDesc(v EnvironmentDescription) {
	o.Desc = &v
}

func (o DescribeEnvironmentReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeEnvironmentReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	return toSerialize, nil
}

type NullableDescribeEnvironmentReply struct {
	value *DescribeEnvironmentReply
	isSet bool
}

func (v NullableDescribeEnvironmentReply) Get() *DescribeEnvironmentReply {
	return v.value
}

func (v *NullableDescribeEnvironmentReply) Set(val *DescribeEnvironmentReply) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeEnvironmentReply) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeEnvironmentReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeEnvironmentReply(val *DescribeEnvironmentReply) *NullableDescribeEnvironmentReply {
	return &NullableDescribeEnvironmentReply{value: val, isSet: true}
}

func (v NullableDescribeEnvironmentReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeEnvironmentReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


