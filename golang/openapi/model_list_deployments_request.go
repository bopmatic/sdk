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

// checks if the ListDeploymentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeploymentsRequest{}

// ListDeploymentsRequest struct for ListDeploymentsRequest
type ListDeploymentsRequest struct {
	ProjEnvHeader *ProjEnvHeader `json:"projEnvHeader,omitempty"`
}

// NewListDeploymentsRequest instantiates a new ListDeploymentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeploymentsRequest() *ListDeploymentsRequest {
	this := ListDeploymentsRequest{}
	return &this
}

// NewListDeploymentsRequestWithDefaults instantiates a new ListDeploymentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeploymentsRequestWithDefaults() *ListDeploymentsRequest {
	this := ListDeploymentsRequest{}
	return &this
}

// GetProjEnvHeader returns the ProjEnvHeader field value if set, zero value otherwise.
func (o *ListDeploymentsRequest) GetProjEnvHeader() ProjEnvHeader {
	if o == nil || IsNil(o.ProjEnvHeader) {
		var ret ProjEnvHeader
		return ret
	}
	return *o.ProjEnvHeader
}

// GetProjEnvHeaderOk returns a tuple with the ProjEnvHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeploymentsRequest) GetProjEnvHeaderOk() (*ProjEnvHeader, bool) {
	if o == nil || IsNil(o.ProjEnvHeader) {
		return nil, false
	}
	return o.ProjEnvHeader, true
}

// HasProjEnvHeader returns a boolean if a field has been set.
func (o *ListDeploymentsRequest) HasProjEnvHeader() bool {
	if o != nil && !IsNil(o.ProjEnvHeader) {
		return true
	}

	return false
}

// SetProjEnvHeader gets a reference to the given ProjEnvHeader and assigns it to the ProjEnvHeader field.
func (o *ListDeploymentsRequest) SetProjEnvHeader(v ProjEnvHeader) {
	o.ProjEnvHeader = &v
}

func (o ListDeploymentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeploymentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjEnvHeader) {
		toSerialize["projEnvHeader"] = o.ProjEnvHeader
	}
	return toSerialize, nil
}

type NullableListDeploymentsRequest struct {
	value *ListDeploymentsRequest
	isSet bool
}

func (v NullableListDeploymentsRequest) Get() *ListDeploymentsRequest {
	return v.value
}

func (v *NullableListDeploymentsRequest) Set(val *ListDeploymentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeploymentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeploymentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeploymentsRequest(val *ListDeploymentsRequest) *NullableListDeploymentsRequest {
	return &NullableListDeploymentsRequest{value: val, isSet: true}
}

func (v NullableListDeploymentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeploymentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


