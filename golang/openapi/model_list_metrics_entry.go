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

// checks if the ListMetricsEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricsEntry{}

// ListMetricsEntry struct for ListMetricsEntry
type ListMetricsEntry struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SamplePeriod *string `json:"samplePeriod,omitempty"`
}

// NewListMetricsEntry instantiates a new ListMetricsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricsEntry() *ListMetricsEntry {
	this := ListMetricsEntry{}
	return &this
}

// NewListMetricsEntryWithDefaults instantiates a new ListMetricsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricsEntryWithDefaults() *ListMetricsEntry {
	this := ListMetricsEntry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListMetricsEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListMetricsEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListMetricsEntry) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListMetricsEntry) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsEntry) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListMetricsEntry) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListMetricsEntry) SetDescription(v string) {
	o.Description = &v
}

// GetSamplePeriod returns the SamplePeriod field value if set, zero value otherwise.
func (o *ListMetricsEntry) GetSamplePeriod() string {
	if o == nil || IsNil(o.SamplePeriod) {
		var ret string
		return ret
	}
	return *o.SamplePeriod
}

// GetSamplePeriodOk returns a tuple with the SamplePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsEntry) GetSamplePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.SamplePeriod) {
		return nil, false
	}
	return o.SamplePeriod, true
}

// HasSamplePeriod returns a boolean if a field has been set.
func (o *ListMetricsEntry) HasSamplePeriod() bool {
	if o != nil && !IsNil(o.SamplePeriod) {
		return true
	}

	return false
}

// SetSamplePeriod gets a reference to the given string and assigns it to the SamplePeriod field.
func (o *ListMetricsEntry) SetSamplePeriod(v string) {
	o.SamplePeriod = &v
}

func (o ListMetricsEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetricsEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SamplePeriod) {
		toSerialize["samplePeriod"] = o.SamplePeriod
	}
	return toSerialize, nil
}

type NullableListMetricsEntry struct {
	value *ListMetricsEntry
	isSet bool
}

func (v NullableListMetricsEntry) Get() *ListMetricsEntry {
	return v.value
}

func (v *NullableListMetricsEntry) Set(val *ListMetricsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricsEntry(val *ListMetricsEntry) *NullableListMetricsEntry {
	return &NullableListMetricsEntry{value: val, isSet: true}
}

func (v NullableListMetricsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

