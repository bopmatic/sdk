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

// PackageState the model 'PackageState'
type PackageState string

// List of PackageState
const (
	UPLOADING PackageState = "UPLOADING"
	UPLOADED PackageState = "UPLOADED"
	VALIDATING PackageState = "VALIDATING"
	INVALID PackageState = "INVALID"
	DEPLOYING PackageState = "DEPLOYING"
	PRODUCTION PackageState = "PRODUCTION"
	DEACTIVATING PackageState = "DEACTIVATING"
	DELETING PackageState = "DELETING"
	UNKNOWN_PKG_STATE PackageState = "UNKNOWN_PKG_STATE"
)

// All allowed values of PackageState enum
var AllowedPackageStateEnumValues = []PackageState{
	"UPLOADING",
	"UPLOADED",
	"VALIDATING",
	"INVALID",
	"DEPLOYING",
	"PRODUCTION",
	"DEACTIVATING",
	"DELETING",
	"UNKNOWN_PKG_STATE",
}

func (v *PackageState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageState(value)
	for _, existing := range AllowedPackageStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageState", value)
}

// NewPackageStateFromValue returns a pointer to a valid PackageState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageStateFromValue(v string) (*PackageState, error) {
	ev := PackageState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageState: valid values are %v", v, AllowedPackageStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageState) IsValid() bool {
	for _, existing := range AllowedPackageStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackageState value
func (v PackageState) Ptr() *PackageState {
	return &v
}

type NullablePackageState struct {
	value *PackageState
	isSet bool
}

func (v NullablePackageState) Get() *PackageState {
	return v.value
}

func (v *NullablePackageState) Set(val *PackageState) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageState) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageState(val *PackageState) *NullablePackageState {
	return &NullablePackageState{value: val, isSet: true}
}

func (v NullablePackageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

