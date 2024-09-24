/*
pb/sr.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProjectState the model 'ProjectState'
type ProjectState string

// List of ProjectState
const (
	INACTIVE ProjectState = "INACTIVE"
	ACTIVE ProjectState = "ACTIVE"
)

// All allowed values of ProjectState enum
var AllowedProjectStateEnumValues = []ProjectState{
	"INACTIVE",
	"ACTIVE",
}

func (v *ProjectState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectState(value)
	for _, existing := range AllowedProjectStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectState", value)
}

// NewProjectStateFromValue returns a pointer to a valid ProjectState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectStateFromValue(v string) (*ProjectState, error) {
	ev := ProjectState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectState: valid values are %v", v, AllowedProjectStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectState) IsValid() bool {
	for _, existing := range AllowedProjectStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectState value
func (v ProjectState) Ptr() *ProjectState {
	return &v
}

type NullableProjectState struct {
	value *ProjectState
	isSet bool
}

func (v NullableProjectState) Get() *ProjectState {
	return v.value
}

func (v *NullableProjectState) Set(val *ProjectState) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectState) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectState(val *ProjectState) *NullableProjectState {
	return &NullableProjectState{value: val, isSet: true}
}

func (v NullableProjectState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

