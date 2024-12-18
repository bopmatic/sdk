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

// checks if the GetLogsEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogsEntry{}

// GetLogsEntry struct for GetLogsEntry
type GetLogsEntry struct {
	Timestamp *string `json:"timestamp,omitempty"`
	// log message
	Message *string `json:"message,omitempty"`
}

// NewGetLogsEntry instantiates a new GetLogsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogsEntry() *GetLogsEntry {
	this := GetLogsEntry{}
	return &this
}

// NewGetLogsEntryWithDefaults instantiates a new GetLogsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogsEntryWithDefaults() *GetLogsEntry {
	this := GetLogsEntry{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetLogsEntry) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogsEntry) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetLogsEntry) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetLogsEntry) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetLogsEntry) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogsEntry) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetLogsEntry) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetLogsEntry) SetMessage(v string) {
	o.Message = &v
}

func (o GetLogsEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogsEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetLogsEntry struct {
	value *GetLogsEntry
	isSet bool
}

func (v NullableGetLogsEntry) Get() *GetLogsEntry {
	return v.value
}

func (v *NullableGetLogsEntry) Set(val *GetLogsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogsEntry(val *GetLogsEntry) *NullableGetLogsEntry {
	return &NullableGetLogsEntry{value: val, isSet: true}
}

func (v NullableGetLogsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


