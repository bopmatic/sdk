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

// checks if the ListDatabasesReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabasesReply{}

// ListDatabasesReply struct for ListDatabasesReply
type ListDatabasesReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Header *ProjEnvHeader `json:"header,omitempty"`
	DatabaseNames []string `json:"databaseNames,omitempty"`
}

// NewListDatabasesReply instantiates a new ListDatabasesReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabasesReply() *ListDatabasesReply {
	this := ListDatabasesReply{}
	return &this
}

// NewListDatabasesReplyWithDefaults instantiates a new ListDatabasesReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabasesReplyWithDefaults() *ListDatabasesReply {
	this := ListDatabasesReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListDatabasesReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabasesReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListDatabasesReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *ListDatabasesReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ListDatabasesReply) GetHeader() ProjEnvHeader {
	if o == nil || IsNil(o.Header) {
		var ret ProjEnvHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabasesReply) GetHeaderOk() (*ProjEnvHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ListDatabasesReply) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ProjEnvHeader and assigns it to the Header field.
func (o *ListDatabasesReply) SetHeader(v ProjEnvHeader) {
	o.Header = &v
}

// GetDatabaseNames returns the DatabaseNames field value if set, zero value otherwise.
func (o *ListDatabasesReply) GetDatabaseNames() []string {
	if o == nil || IsNil(o.DatabaseNames) {
		var ret []string
		return ret
	}
	return o.DatabaseNames
}

// GetDatabaseNamesOk returns a tuple with the DatabaseNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabasesReply) GetDatabaseNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DatabaseNames) {
		return nil, false
	}
	return o.DatabaseNames, true
}

// HasDatabaseNames returns a boolean if a field has been set.
func (o *ListDatabasesReply) HasDatabaseNames() bool {
	if o != nil && !IsNil(o.DatabaseNames) {
		return true
	}

	return false
}

// SetDatabaseNames gets a reference to the given []string and assigns it to the DatabaseNames field.
func (o *ListDatabasesReply) SetDatabaseNames(v []string) {
	o.DatabaseNames = v
}

func (o ListDatabasesReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabasesReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.DatabaseNames) {
		toSerialize["databaseNames"] = o.DatabaseNames
	}
	return toSerialize, nil
}

type NullableListDatabasesReply struct {
	value *ListDatabasesReply
	isSet bool
}

func (v NullableListDatabasesReply) Get() *ListDatabasesReply {
	return v.value
}

func (v *NullableListDatabasesReply) Set(val *ListDatabasesReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabasesReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabasesReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabasesReply(val *ListDatabasesReply) *NullableListDatabasesReply {
	return &NullableListDatabasesReply{value: val, isSet: true}
}

func (v NullableListDatabasesReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabasesReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


