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

// DeployPackageReply struct for DeployPackageReply
type DeployPackageReply struct {
	State *PackageState `json:"state,omitempty"`
}

// NewDeployPackageReply instantiates a new DeployPackageReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployPackageReply() *DeployPackageReply {
	this := DeployPackageReply{}
	var state PackageState = UPLOADING
	this.State = &state
	return &this
}

// NewDeployPackageReplyWithDefaults instantiates a new DeployPackageReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployPackageReplyWithDefaults() *DeployPackageReply {
	this := DeployPackageReply{}
	var state PackageState = UPLOADING
	this.State = &state
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DeployPackageReply) GetState() PackageState {
	if o == nil || o.State == nil {
		var ret PackageState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployPackageReply) GetStateOk() (*PackageState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DeployPackageReply) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given PackageState and assigns it to the State field.
func (o *DeployPackageReply) SetState(v PackageState) {
	o.State = &v
}

func (o DeployPackageReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableDeployPackageReply struct {
	value *DeployPackageReply
	isSet bool
}

func (v NullableDeployPackageReply) Get() *DeployPackageReply {
	return v.value
}

func (v *NullableDeployPackageReply) Set(val *DeployPackageReply) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployPackageReply) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployPackageReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployPackageReply(val *DeployPackageReply) *NullableDeployPackageReply {
	return &NullableDeployPackageReply{value: val, isSet: true}
}

func (v NullableDeployPackageReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployPackageReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


