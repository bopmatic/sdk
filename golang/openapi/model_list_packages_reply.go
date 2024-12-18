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

// checks if the ListPackagesReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackagesReply{}

// ListPackagesReply struct for ListPackagesReply
type ListPackagesReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Items []ListPackagesReplyListPackagesItem `json:"items,omitempty"`
}

// NewListPackagesReply instantiates a new ListPackagesReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackagesReply() *ListPackagesReply {
	this := ListPackagesReply{}
	return &this
}

// NewListPackagesReplyWithDefaults instantiates a new ListPackagesReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackagesReplyWithDefaults() *ListPackagesReply {
	this := ListPackagesReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListPackagesReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackagesReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListPackagesReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *ListPackagesReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListPackagesReply) GetItems() []ListPackagesReplyListPackagesItem {
	if o == nil || IsNil(o.Items) {
		var ret []ListPackagesReplyListPackagesItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackagesReply) GetItemsOk() ([]ListPackagesReplyListPackagesItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListPackagesReply) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ListPackagesReplyListPackagesItem and assigns it to the Items field.
func (o *ListPackagesReply) SetItems(v []ListPackagesReplyListPackagesItem) {
	o.Items = v
}

func (o ListPackagesReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPackagesReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListPackagesReply struct {
	value *ListPackagesReply
	isSet bool
}

func (v NullableListPackagesReply) Get() *ListPackagesReply {
	return v.value
}

func (v *NullableListPackagesReply) Set(val *ListPackagesReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackagesReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackagesReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackagesReply(val *ListPackagesReply) *NullableListPackagesReply {
	return &NullableListPackagesReply{value: val, isSet: true}
}

func (v NullableListPackagesReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPackagesReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


