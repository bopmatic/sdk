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

// ListPackagesRequest struct for ListPackagesRequest
type ListPackagesRequest struct {
	ProjectName *string `json:"projectName,omitempty"`
}

// NewListPackagesRequest instantiates a new ListPackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackagesRequest() *ListPackagesRequest {
	this := ListPackagesRequest{}
	return &this
}

// NewListPackagesRequestWithDefaults instantiates a new ListPackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackagesRequestWithDefaults() *ListPackagesRequest {
	this := ListPackagesRequest{}
	return &this
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *ListPackagesRequest) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackagesRequest) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *ListPackagesRequest) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *ListPackagesRequest) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o ListPackagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectName != nil {
		toSerialize["projectName"] = o.ProjectName
	}
	return json.Marshal(toSerialize)
}

type NullableListPackagesRequest struct {
	value *ListPackagesRequest
	isSet bool
}

func (v NullableListPackagesRequest) Get() *ListPackagesRequest {
	return v.value
}

func (v *NullableListPackagesRequest) Set(val *ListPackagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackagesRequest(val *ListPackagesRequest) *NullableListPackagesRequest {
	return &NullableListPackagesRequest{value: val, isSet: true}
}

func (v NullableListPackagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPackagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

