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

// checks if the ProjectDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDescription{}

// ProjectDescription struct for ProjectDescription
type ProjectDescription struct {
	Id *string `json:"id,omitempty"`
	Header *ProjectHeader `json:"header,omitempty"`
	State *ProjectState `json:"state,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ActiveDeployIds []string `json:"activeDeployIds,omitempty"`
	PendingDeployIds []string `json:"pendingDeployIds,omitempty"`
}

// NewProjectDescription instantiates a new ProjectDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDescription() *ProjectDescription {
	this := ProjectDescription{}
	var state ProjectState = INACTIVE
	this.State = &state
	return &this
}

// NewProjectDescriptionWithDefaults instantiates a new ProjectDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDescriptionWithDefaults() *ProjectDescription {
	this := ProjectDescription{}
	var state ProjectState = INACTIVE
	this.State = &state
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectDescription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectDescription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectDescription) SetId(v string) {
	o.Id = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ProjectDescription) GetHeader() ProjectHeader {
	if o == nil || IsNil(o.Header) {
		var ret ProjectHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetHeaderOk() (*ProjectHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ProjectDescription) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given ProjectHeader and assigns it to the Header field.
func (o *ProjectDescription) SetHeader(v ProjectHeader) {
	o.Header = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProjectDescription) GetState() ProjectState {
	if o == nil || IsNil(o.State) {
		var ret ProjectState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetStateOk() (*ProjectState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProjectDescription) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ProjectState and assigns it to the State field.
func (o *ProjectDescription) SetState(v ProjectState) {
	o.State = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ProjectDescription) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ProjectDescription) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ProjectDescription) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetActiveDeployIds returns the ActiveDeployIds field value if set, zero value otherwise.
func (o *ProjectDescription) GetActiveDeployIds() []string {
	if o == nil || IsNil(o.ActiveDeployIds) {
		var ret []string
		return ret
	}
	return o.ActiveDeployIds
}

// GetActiveDeployIdsOk returns a tuple with the ActiveDeployIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetActiveDeployIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ActiveDeployIds) {
		return nil, false
	}
	return o.ActiveDeployIds, true
}

// HasActiveDeployIds returns a boolean if a field has been set.
func (o *ProjectDescription) HasActiveDeployIds() bool {
	if o != nil && !IsNil(o.ActiveDeployIds) {
		return true
	}

	return false
}

// SetActiveDeployIds gets a reference to the given []string and assigns it to the ActiveDeployIds field.
func (o *ProjectDescription) SetActiveDeployIds(v []string) {
	o.ActiveDeployIds = v
}

// GetPendingDeployIds returns the PendingDeployIds field value if set, zero value otherwise.
func (o *ProjectDescription) GetPendingDeployIds() []string {
	if o == nil || IsNil(o.PendingDeployIds) {
		var ret []string
		return ret
	}
	return o.PendingDeployIds
}

// GetPendingDeployIdsOk returns a tuple with the PendingDeployIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDescription) GetPendingDeployIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PendingDeployIds) {
		return nil, false
	}
	return o.PendingDeployIds, true
}

// HasPendingDeployIds returns a boolean if a field has been set.
func (o *ProjectDescription) HasPendingDeployIds() bool {
	if o != nil && !IsNil(o.PendingDeployIds) {
		return true
	}

	return false
}

// SetPendingDeployIds gets a reference to the given []string and assigns it to the PendingDeployIds field.
func (o *ProjectDescription) SetPendingDeployIds(v []string) {
	o.PendingDeployIds = v
}

func (o ProjectDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ActiveDeployIds) {
		toSerialize["activeDeployIds"] = o.ActiveDeployIds
	}
	if !IsNil(o.PendingDeployIds) {
		toSerialize["pendingDeployIds"] = o.PendingDeployIds
	}
	return toSerialize, nil
}

type NullableProjectDescription struct {
	value *ProjectDescription
	isSet bool
}

func (v NullableProjectDescription) Get() *ProjectDescription {
	return v.value
}

func (v *NullableProjectDescription) Set(val *ProjectDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDescription(val *ProjectDescription) *NullableProjectDescription {
	return &NullableProjectDescription{value: val, isSet: true}
}

func (v NullableProjectDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


