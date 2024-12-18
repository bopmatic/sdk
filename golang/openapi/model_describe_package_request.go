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

// checks if the DescribePackageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePackageRequest{}

// DescribePackageRequest struct for DescribePackageRequest
type DescribePackageRequest struct {
	PackageId *string `json:"packageId,omitempty"`
}

// NewDescribePackageRequest instantiates a new DescribePackageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePackageRequest() *DescribePackageRequest {
	this := DescribePackageRequest{}
	return &this
}

// NewDescribePackageRequestWithDefaults instantiates a new DescribePackageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePackageRequestWithDefaults() *DescribePackageRequest {
	this := DescribePackageRequest{}
	return &this
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *DescribePackageRequest) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePackageRequest) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *DescribePackageRequest) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *DescribePackageRequest) SetPackageId(v string) {
	o.PackageId = &v
}

func (o DescribePackageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePackageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	return toSerialize, nil
}

type NullableDescribePackageRequest struct {
	value *DescribePackageRequest
	isSet bool
}

func (v NullableDescribePackageRequest) Get() *DescribePackageRequest {
	return v.value
}

func (v *NullableDescribePackageRequest) Set(val *DescribePackageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePackageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePackageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePackageRequest(val *DescribePackageRequest) *NullableDescribePackageRequest {
	return &NullableDescribePackageRequest{value: val, isSet: true}
}

func (v NullableDescribePackageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePackageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


