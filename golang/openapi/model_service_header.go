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

// checks if the ServiceHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceHeader{}

// ServiceHeader struct for ServiceHeader
type ServiceHeader struct {
	ProjEnvHeader *ProjEnvHeader `json:"projEnvHeader,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
}

// NewServiceHeader instantiates a new ServiceHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHeader() *ServiceHeader {
	this := ServiceHeader{}
	return &this
}

// NewServiceHeaderWithDefaults instantiates a new ServiceHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHeaderWithDefaults() *ServiceHeader {
	this := ServiceHeader{}
	return &this
}

// GetProjEnvHeader returns the ProjEnvHeader field value if set, zero value otherwise.
func (o *ServiceHeader) GetProjEnvHeader() ProjEnvHeader {
	if o == nil || IsNil(o.ProjEnvHeader) {
		var ret ProjEnvHeader
		return ret
	}
	return *o.ProjEnvHeader
}

// GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHeader) GetProjEnvHeaderOk() (*ProjEnvHeader, bool) {
	if o == nil || IsNil(o.ProjEnvHeader) {
		return nil, false
	}
	return o.ProjEnvHeader, true
}

// HasProjEnvHeader returns a boolean if a field has been set.
func (o *ServiceHeader) HasProjEnvHeader() bool {
	if o != nil && !IsNil(o.ProjEnvHeader) {
		return true
	}

	return false
}

// SetProjEnvHeader gets a reference to the given ProjEnvHeader and assigns it to the ProjEnvHeader field.
func (o *ServiceHeader) SetProjEnvHeader(v ProjEnvHeader) {
	o.ProjEnvHeader = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceHeader) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHeader) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceHeader) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceHeader) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o ServiceHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjEnvHeader) {
		toSerialize["projEnvHeader"] = o.ProjEnvHeader
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	return toSerialize, nil
}

type NullableServiceHeader struct {
	value *ServiceHeader
	isSet bool
}

func (v NullableServiceHeader) Get() *ServiceHeader {
	return v.value
}

func (v *NullableServiceHeader) Set(val *ServiceHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHeader(val *ServiceHeader) *NullableServiceHeader {
	return &NullableServiceHeader{value: val, isSet: true}
}

func (v NullableServiceHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


