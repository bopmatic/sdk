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

// checks if the CreateDeploymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeploymentRequest{}

// CreateDeploymentRequest struct for CreateDeploymentRequest
type CreateDeploymentRequest struct {
	Header *DeploymentHeader `json:"header,omitempty"`
}

// NewCreateDeploymentRequest instantiates a new CreateDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeploymentRequest() *CreateDeploymentRequest {
	this := CreateDeploymentRequest{}
	return &this
}

// NewCreateDeploymentRequestWithDefaults instantiates a new CreateDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeploymentRequestWithDefaults() *CreateDeploymentRequest {
	this := CreateDeploymentRequest{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CreateDeploymentRequest) GetHeader() DeploymentHeader {
	if o == nil || IsNil(o.Header) {
		var ret DeploymentHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeploymentRequest) GetHeaderOk() (*DeploymentHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CreateDeploymentRequest) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given DeploymentHeader and assigns it to the Header field.
func (o *CreateDeploymentRequest) SetHeader(v DeploymentHeader) {
	o.Header = &v
}

func (o CreateDeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeploymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	return toSerialize, nil
}

type NullableCreateDeploymentRequest struct {
	value *CreateDeploymentRequest
	isSet bool
}

func (v NullableCreateDeploymentRequest) Get() *CreateDeploymentRequest {
	return v.value
}

func (v *NullableCreateDeploymentRequest) Set(val *CreateDeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeploymentRequest(val *CreateDeploymentRequest) *NullableCreateDeploymentRequest {
	return &NullableCreateDeploymentRequest{value: val, isSet: true}
}

func (v NullableCreateDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


