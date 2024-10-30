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

// checks if the DeactivateProjectReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeactivateProjectReply{}

// DeactivateProjectReply struct for DeactivateProjectReply
type DeactivateProjectReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	DeployId *string `json:"deployId,omitempty"`
}

// NewDeactivateProjectReply instantiates a new DeactivateProjectReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeactivateProjectReply() *DeactivateProjectReply {
	this := DeactivateProjectReply{}
	return &this
}

// NewDeactivateProjectReplyWithDefaults instantiates a new DeactivateProjectReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeactivateProjectReplyWithDefaults() *DeactivateProjectReply {
	this := DeactivateProjectReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DeactivateProjectReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeactivateProjectReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DeactivateProjectReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *DeactivateProjectReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetDeployId returns the DeployId field value if set, zero value otherwise.
func (o *DeactivateProjectReply) GetDeployId() string {
	if o == nil || IsNil(o.DeployId) {
		var ret string
		return ret
	}
	return *o.DeployId
}

// GetDeployIdOk returns a tuple with the DeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeactivateProjectReply) GetDeployIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeployId) {
		return nil, false
	}
	return o.DeployId, true
}

// HasDeployId returns a boolean if a field has been set.
func (o *DeactivateProjectReply) HasDeployId() bool {
	if o != nil && !IsNil(o.DeployId) {
		return true
	}

	return false
}

// SetDeployId gets a reference to the given string and assigns it to the DeployId field.
func (o *DeactivateProjectReply) SetDeployId(v string) {
	o.DeployId = &v
}

func (o DeactivateProjectReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeactivateProjectReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.DeployId) {
		toSerialize["deployId"] = o.DeployId
	}
	return toSerialize, nil
}

type NullableDeactivateProjectReply struct {
	value *DeactivateProjectReply
	isSet bool
}

func (v NullableDeactivateProjectReply) Get() *DeactivateProjectReply {
	return v.value
}

func (v *NullableDeactivateProjectReply) Set(val *DeactivateProjectReply) {
	v.value = val
	v.isSet = true
}

func (v NullableDeactivateProjectReply) IsSet() bool {
	return v.isSet
}

func (v *NullableDeactivateProjectReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeactivateProjectReply(val *DeactivateProjectReply) *NullableDeactivateProjectReply {
	return &NullableDeactivateProjectReply{value: val, isSet: true}
}

func (v NullableDeactivateProjectReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeactivateProjectReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

