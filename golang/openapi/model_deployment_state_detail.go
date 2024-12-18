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

// DeploymentStateDetail the package associated with the deployment failed to compile  - DPLY_SUPPORT_NEEDED: the deployment failed for an unknown reason; Bopmatic support  - UNKNOWN_DEPLOY_STATE_DET: MAX_INT
type DeploymentStateDetail string

// List of DeploymentStateDetail
const (
	INVALID_DEPLOY_STATE_DET DeploymentStateDetail = "INVALID_DEPLOY_STATE_DET"
	NONE DeploymentStateDetail = "NONE"
	PKG_INVALID DeploymentStateDetail = "PKG_INVALID"
	BLD_INVALID DeploymentStateDetail = "BLD_INVALID"
	DPLY_SUPPORT_NEEDED DeploymentStateDetail = "DPLY_SUPPORT_NEEDED"
	UNKNOWN_DEPLOY_STATE_DET DeploymentStateDetail = "UNKNOWN_DEPLOY_STATE_DET"
)

// All allowed values of DeploymentStateDetail enum
var AllowedDeploymentStateDetailEnumValues = []DeploymentStateDetail{
	"INVALID_DEPLOY_STATE_DET",
	"NONE",
	"PKG_INVALID",
	"BLD_INVALID",
	"DPLY_SUPPORT_NEEDED",
	"UNKNOWN_DEPLOY_STATE_DET",
}

func (v *DeploymentStateDetail) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentStateDetail(value)
	for _, existing := range AllowedDeploymentStateDetailEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentStateDetail", value)
}

// NewDeploymentStateDetailFromValue returns a pointer to a valid DeploymentStateDetail
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentStateDetailFromValue(v string) (*DeploymentStateDetail, error) {
	ev := DeploymentStateDetail(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentStateDetail: valid values are %v", v, AllowedDeploymentStateDetailEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentStateDetail) IsValid() bool {
	for _, existing := range AllowedDeploymentStateDetailEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentStateDetail value
func (v DeploymentStateDetail) Ptr() *DeploymentStateDetail {
	return &v
}

type NullableDeploymentStateDetail struct {
	value *DeploymentStateDetail
	isSet bool
}

func (v NullableDeploymentStateDetail) Get() *DeploymentStateDetail {
	return v.value
}

func (v *NullableDeploymentStateDetail) Set(val *DeploymentStateDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateDetail(val *DeploymentStateDetail) *NullableDeploymentStateDetail {
	return &NullableDeploymentStateDetail{value: val, isSet: true}
}

func (v NullableDeploymentStateDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

