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

// checks if the ListMetricsReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricsReply{}

// ListMetricsReply struct for ListMetricsReply
type ListMetricsReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Entries []ListMetricsEntry `json:"entries,omitempty"`
}

// NewListMetricsReply instantiates a new ListMetricsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricsReply() *ListMetricsReply {
	this := ListMetricsReply{}
	return &this
}

// NewListMetricsReplyWithDefaults instantiates a new ListMetricsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricsReplyWithDefaults() *ListMetricsReply {
	this := ListMetricsReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListMetricsReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListMetricsReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *ListMetricsReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *ListMetricsReply) GetEntries() []ListMetricsEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []ListMetricsEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsReply) GetEntriesOk() ([]ListMetricsEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ListMetricsReply) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []ListMetricsEntry and assigns it to the Entries field.
func (o *ListMetricsReply) SetEntries(v []ListMetricsEntry) {
	o.Entries = v
}

func (o ListMetricsReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetricsReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableListMetricsReply struct {
	value *ListMetricsReply
	isSet bool
}

func (v NullableListMetricsReply) Get() *ListMetricsReply {
	return v.value
}

func (v *NullableListMetricsReply) Set(val *ListMetricsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricsReply(val *ListMetricsReply) *NullableListMetricsReply {
	return &NullableListMetricsReply{value: val, isSet: true}
}

func (v NullableListMetricsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


