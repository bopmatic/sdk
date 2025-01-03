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

// checks if the DescribePackageReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePackageReply{}

// DescribePackageReply struct for DescribePackageReply
type DescribePackageReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Desc *PackageDescription `json:"desc,omitempty"`
}

// NewDescribePackageReply instantiates a new DescribePackageReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePackageReply() *DescribePackageReply {
	this := DescribePackageReply{}
	return &this
}

// NewDescribePackageReplyWithDefaults instantiates a new DescribePackageReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePackageReplyWithDefaults() *DescribePackageReply {
	this := DescribePackageReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DescribePackageReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePackageReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DescribePackageReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *DescribePackageReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DescribePackageReply) GetDesc() PackageDescription {
	if o == nil || IsNil(o.Desc) {
		var ret PackageDescription
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePackageReply) GetDescOk() (*PackageDescription, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DescribePackageReply) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given PackageDescription and assigns it to the Desc field.
func (o *DescribePackageReply) SetDesc(v PackageDescription) {
	o.Desc = &v
}

func (o DescribePackageReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePackageReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	return toSerialize, nil
}

type NullableDescribePackageReply struct {
	value *DescribePackageReply
	isSet bool
}

func (v NullableDescribePackageReply) Get() *DescribePackageReply {
	return v.value
}

func (v *NullableDescribePackageReply) Set(val *DescribePackageReply) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePackageReply) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePackageReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePackageReply(val *DescribePackageReply) *NullableDescribePackageReply {
	return &NullableDescribePackageReply{value: val, isSet: true}
}

func (v NullableDescribePackageReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePackageReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


