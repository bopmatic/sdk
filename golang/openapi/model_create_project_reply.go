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

// checks if the CreateProjectReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectReply{}

// CreateProjectReply struct for CreateProjectReply
type CreateProjectReply struct {
	Result *ServiceRunnerResult `json:"result,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCreateProjectReply instantiates a new CreateProjectReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectReply() *CreateProjectReply {
	this := CreateProjectReply{}
	return &this
}

// NewCreateProjectReplyWithDefaults instantiates a new CreateProjectReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectReplyWithDefaults() *CreateProjectReply {
	this := CreateProjectReply{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateProjectReply) GetResult() ServiceRunnerResult {
	if o == nil || IsNil(o.Result) {
		var ret ServiceRunnerResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectReply) GetResultOk() (*ServiceRunnerResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateProjectReply) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ServiceRunnerResult and assigns it to the Result field.
func (o *CreateProjectReply) SetResult(v ServiceRunnerResult) {
	o.Result = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateProjectReply) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectReply) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateProjectReply) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateProjectReply) SetId(v string) {
	o.Id = &v
}

func (o CreateProjectReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateProjectReply struct {
	value *CreateProjectReply
	isSet bool
}

func (v NullableCreateProjectReply) Get() *CreateProjectReply {
	return v.value
}

func (v *NullableCreateProjectReply) Set(val *CreateProjectReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectReply(val *CreateProjectReply) *NullableCreateProjectReply {
	return &NullableCreateProjectReply{value: val, isSet: true}
}

func (v NullableCreateProjectReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


