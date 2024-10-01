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

// checks if the DescribeDatastoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeDatastoreRequest{}

// DescribeDatastoreRequest struct for DescribeDatastoreRequest
type DescribeDatastoreRequest struct {
	DatastoreHeader *DatastoreHeader `json:"datastoreHeader,omitempty"`
}

// NewDescribeDatastoreRequest instantiates a new DescribeDatastoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDatastoreRequest() *DescribeDatastoreRequest {
	this := DescribeDatastoreRequest{}
	return &this
}

// NewDescribeDatastoreRequestWithDefaults instantiates a new DescribeDatastoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDatastoreRequestWithDefaults() *DescribeDatastoreRequest {
	this := DescribeDatastoreRequest{}
	return &this
}

// GetDatastoreHeader returns the DatastoreHeader field value if set, zero value otherwise.
func (o *DescribeDatastoreRequest) GetDatastoreHeader() DatastoreHeader {
	if o == nil || IsNil(o.DatastoreHeader) {
		var ret DatastoreHeader
		return ret
	}
	return *o.DatastoreHeader
}

// GetDatastoreHeaderOk returns a tuple with the DatastoreHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDatastoreRequest) GetDatastoreHeaderOk() (*DatastoreHeader, bool) {
	if o == nil || IsNil(o.DatastoreHeader) {
		return nil, false
	}
	return o.DatastoreHeader, true
}

// HasDatastoreHeader returns a boolean if a field has been set.
func (o *DescribeDatastoreRequest) HasDatastoreHeader() bool {
	if o != nil && !IsNil(o.DatastoreHeader) {
		return true
	}

	return false
}

// SetDatastoreHeader gets a reference to the given DatastoreHeader and assigns it to the DatastoreHeader field.
func (o *DescribeDatastoreRequest) SetDatastoreHeader(v DatastoreHeader) {
	o.DatastoreHeader = &v
}

func (o DescribeDatastoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDatastoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatastoreHeader) {
		toSerialize["datastoreHeader"] = o.DatastoreHeader
	}
	return toSerialize, nil
}

type NullableDescribeDatastoreRequest struct {
	value *DescribeDatastoreRequest
	isSet bool
}

func (v NullableDescribeDatastoreRequest) Get() *DescribeDatastoreRequest {
	return v.value
}

func (v *NullableDescribeDatastoreRequest) Set(val *DescribeDatastoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDatastoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDatastoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDatastoreRequest(val *DescribeDatastoreRequest) *NullableDescribeDatastoreRequest {
	return &NullableDescribeDatastoreRequest{value: val, isSet: true}
}

func (v NullableDescribeDatastoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDatastoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

