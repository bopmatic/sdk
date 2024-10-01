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

// MetricsFormat the model 'MetricsFormat'
type MetricsFormat string

// List of MetricsFormat
const (
	METRIC_FORMAT_OPENMETRICS MetricsFormat = "METRIC_FORMAT_OPENMETRICS"
	UNKNOWN_METRIC_FORMAT MetricsFormat = "UNKNOWN_METRIC_FORMAT"
)

// All allowed values of MetricsFormat enum
var AllowedMetricsFormatEnumValues = []MetricsFormat{
	"METRIC_FORMAT_OPENMETRICS",
	"UNKNOWN_METRIC_FORMAT",
}

func (v *MetricsFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricsFormat(value)
	for _, existing := range AllowedMetricsFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricsFormat", value)
}

// NewMetricsFormatFromValue returns a pointer to a valid MetricsFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricsFormatFromValue(v string) (*MetricsFormat, error) {
	ev := MetricsFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricsFormat: valid values are %v", v, AllowedMetricsFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricsFormat) IsValid() bool {
	for _, existing := range AllowedMetricsFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsFormat value
func (v MetricsFormat) Ptr() *MetricsFormat {
	return &v
}

type NullableMetricsFormat struct {
	value *MetricsFormat
	isSet bool
}

func (v NullableMetricsFormat) Get() *MetricsFormat {
	return v.value
}

func (v *NullableMetricsFormat) Set(val *MetricsFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsFormat(val *MetricsFormat) *NullableMetricsFormat {
	return &NullableMetricsFormat{value: val, isSet: true}
}

func (v NullableMetricsFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
