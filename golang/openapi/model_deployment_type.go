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

// DeploymentType the deployment was created in order to remove the project from  - SECURITY_UPDATE: the environment (customer initiated)  the deployment was created in order to perform a security update  - INF_UPDATE: on the backend infrastructure (service initiated)  the deployment was created in order to perform a non-security related  - UNKNOWN_DEPLOY_TYPE: MAX_INT
type DeploymentType string

// List of DeploymentType
const (
	INVALID_DEPLOY_TYPE DeploymentType = "INVALID_DEPLOY_TYPE"
	NEW_PACKAGE DeploymentType = "NEW_PACKAGE"
	ENV_TEARDOWN DeploymentType = "ENV_TEARDOWN"
	SECURITY_UPDATE DeploymentType = "SECURITY_UPDATE"
	INF_UPDATE DeploymentType = "INF_UPDATE"
	UNKNOWN_DEPLOY_TYPE DeploymentType = "UNKNOWN_DEPLOY_TYPE"
)

// All allowed values of DeploymentType enum
var AllowedDeploymentTypeEnumValues = []DeploymentType{
	"INVALID_DEPLOY_TYPE",
	"NEW_PACKAGE",
	"ENV_TEARDOWN",
	"SECURITY_UPDATE",
	"INF_UPDATE",
	"UNKNOWN_DEPLOY_TYPE",
}

func (v *DeploymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentType(value)
	for _, existing := range AllowedDeploymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentType", value)
}

// NewDeploymentTypeFromValue returns a pointer to a valid DeploymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentTypeFromValue(v string) (*DeploymentType, error) {
	ev := DeploymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentType: valid values are %v", v, AllowedDeploymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentType) IsValid() bool {
	for _, existing := range AllowedDeploymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentType value
func (v DeploymentType) Ptr() *DeploymentType {
	return &v
}

type NullableDeploymentType struct {
	value *DeploymentType
	isSet bool
}

func (v NullableDeploymentType) Get() *DeploymentType {
	return v.value
}

func (v *NullableDeploymentType) Set(val *DeploymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentType(val *DeploymentType) *NullableDeploymentType {
	return &NullableDeploymentType{value: val, isSet: true}
}

func (v NullableDeploymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

