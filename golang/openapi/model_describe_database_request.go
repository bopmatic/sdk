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

// checks if the DescribeDatabaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeDatabaseRequest{}

// DescribeDatabaseRequest struct for DescribeDatabaseRequest
type DescribeDatabaseRequest struct {
	DatabaseHeader *DatabaseHeader `json:"databaseHeader,omitempty"`
}

// NewDescribeDatabaseRequest instantiates a new DescribeDatabaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDatabaseRequest() *DescribeDatabaseRequest {
	this := DescribeDatabaseRequest{}
	return &this
}

// NewDescribeDatabaseRequestWithDefaults instantiates a new DescribeDatabaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDatabaseRequestWithDefaults() *DescribeDatabaseRequest {
	this := DescribeDatabaseRequest{}
	return &this
}

// GetDatabaseHeader returns the DatabaseHeader field value if set, zero value otherwise.
func (o *DescribeDatabaseRequest) GetDatabaseHeader() DatabaseHeader {
	if o == nil || IsNil(o.DatabaseHeader) {
		var ret DatabaseHeader
		return ret
	}
	return *o.DatabaseHeader
}

// GetDatabaseHeaderOk returns a tuple with the DatabaseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDatabaseRequest) GetDatabaseHeaderOk() (*DatabaseHeader, bool) {
	if o == nil || IsNil(o.DatabaseHeader) {
		return nil, false
	}
	return o.DatabaseHeader, true
}

// HasDatabaseHeader returns a boolean if a field has been set.
func (o *DescribeDatabaseRequest) HasDatabaseHeader() bool {
	if o != nil && !IsNil(o.DatabaseHeader) {
		return true
	}

	return false
}

// SetDatabaseHeader gets a reference to the given DatabaseHeader and assigns it to the DatabaseHeader field.
func (o *DescribeDatabaseRequest) SetDatabaseHeader(v DatabaseHeader) {
	o.DatabaseHeader = &v
}

func (o DescribeDatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDatabaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseHeader) {
		toSerialize["databaseHeader"] = o.DatabaseHeader
	}
	return toSerialize, nil
}

type NullableDescribeDatabaseRequest struct {
	value *DescribeDatabaseRequest
	isSet bool
}

func (v NullableDescribeDatabaseRequest) Get() *DescribeDatabaseRequest {
	return v.value
}

func (v *NullableDescribeDatabaseRequest) Set(val *DescribeDatabaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDatabaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDatabaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDatabaseRequest(val *DescribeDatabaseRequest) *NullableDescribeDatabaseRequest {
	return &NullableDescribeDatabaseRequest{value: val, isSet: true}
}

func (v NullableDescribeDatabaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDatabaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

