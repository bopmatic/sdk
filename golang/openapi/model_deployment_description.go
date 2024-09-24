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

// checks if the DeploymentDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDescription{}

// DeploymentDescription struct for DeploymentDescription
type DeploymentDescription struct {
	Id *string `json:"id,omitempty"`
	Header *DeploymentHeader `json:"header,omitempty"`
	State *DeploymentState `json:"state,omitempty"`
	StateDetail *DeploymentStateDetail `json:"stateDetail,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ValidationStartTime *string `json:"validationStartTime,omitempty"`
	BuildStartTime *string `json:"buildStartTime,omitempty"`
	DeployStartTime *string `json:"deployStartTime,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
}

// NewDeploymentDescription instantiates a new DeploymentDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDescription() *DeploymentDescription {
	this := DeploymentDescription{}
	var state DeploymentState = CREATED
	this.State = &state
	var stateDetail DeploymentStateDetail = NONE
	this.StateDetail = &stateDetail
	return &this
}

// NewDeploymentDescriptionWithDefaults instantiates a new DeploymentDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDescriptionWithDefaults() *DeploymentDescription {
	this := DeploymentDescription{}
	var state DeploymentState = CREATED
	this.State = &state
	var stateDetail DeploymentStateDetail = NONE
	this.StateDetail = &stateDetail
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentDescription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentDescription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeploymentDescription) SetId(v string) {
	o.Id = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *DeploymentDescription) GetHeader() DeploymentHeader {
	if o == nil || IsNil(o.Header) {
		var ret DeploymentHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetHeaderOk() (*DeploymentHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *DeploymentDescription) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given DeploymentHeader and assigns it to the Header field.
func (o *DeploymentDescription) SetHeader(v DeploymentHeader) {
	o.Header = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DeploymentDescription) GetState() DeploymentState {
	if o == nil || IsNil(o.State) {
		var ret DeploymentState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetStateOk() (*DeploymentState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DeploymentDescription) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given DeploymentState and assigns it to the State field.
func (o *DeploymentDescription) SetState(v DeploymentState) {
	o.State = &v
}

// GetStateDetail returns the StateDetail field value if set, zero value otherwise.
func (o *DeploymentDescription) GetStateDetail() DeploymentStateDetail {
	if o == nil || IsNil(o.StateDetail) {
		var ret DeploymentStateDetail
		return ret
	}
	return *o.StateDetail
}

// GetStateDetailOk returns a tuple with the StateDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetStateDetailOk() (*DeploymentStateDetail, bool) {
	if o == nil || IsNil(o.StateDetail) {
		return nil, false
	}
	return o.StateDetail, true
}

// HasStateDetail returns a boolean if a field has been set.
func (o *DeploymentDescription) HasStateDetail() bool {
	if o != nil && !IsNil(o.StateDetail) {
		return true
	}

	return false
}

// SetStateDetail gets a reference to the given DeploymentStateDetail and assigns it to the StateDetail field.
func (o *DeploymentDescription) SetStateDetail(v DeploymentStateDetail) {
	o.StateDetail = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *DeploymentDescription) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *DeploymentDescription) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *DeploymentDescription) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetValidationStartTime returns the ValidationStartTime field value if set, zero value otherwise.
func (o *DeploymentDescription) GetValidationStartTime() string {
	if o == nil || IsNil(o.ValidationStartTime) {
		var ret string
		return ret
	}
	return *o.ValidationStartTime
}

// GetValidationStartTimeOk returns a tuple with the ValidationStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetValidationStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationStartTime) {
		return nil, false
	}
	return o.ValidationStartTime, true
}

// HasValidationStartTime returns a boolean if a field has been set.
func (o *DeploymentDescription) HasValidationStartTime() bool {
	if o != nil && !IsNil(o.ValidationStartTime) {
		return true
	}

	return false
}

// SetValidationStartTime gets a reference to the given string and assigns it to the ValidationStartTime field.
func (o *DeploymentDescription) SetValidationStartTime(v string) {
	o.ValidationStartTime = &v
}

// GetBuildStartTime returns the BuildStartTime field value if set, zero value otherwise.
func (o *DeploymentDescription) GetBuildStartTime() string {
	if o == nil || IsNil(o.BuildStartTime) {
		var ret string
		return ret
	}
	return *o.BuildStartTime
}

// GetBuildStartTimeOk returns a tuple with the BuildStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetBuildStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BuildStartTime) {
		return nil, false
	}
	return o.BuildStartTime, true
}

// HasBuildStartTime returns a boolean if a field has been set.
func (o *DeploymentDescription) HasBuildStartTime() bool {
	if o != nil && !IsNil(o.BuildStartTime) {
		return true
	}

	return false
}

// SetBuildStartTime gets a reference to the given string and assigns it to the BuildStartTime field.
func (o *DeploymentDescription) SetBuildStartTime(v string) {
	o.BuildStartTime = &v
}

// GetDeployStartTime returns the DeployStartTime field value if set, zero value otherwise.
func (o *DeploymentDescription) GetDeployStartTime() string {
	if o == nil || IsNil(o.DeployStartTime) {
		var ret string
		return ret
	}
	return *o.DeployStartTime
}

// GetDeployStartTimeOk returns a tuple with the DeployStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetDeployStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DeployStartTime) {
		return nil, false
	}
	return o.DeployStartTime, true
}

// HasDeployStartTime returns a boolean if a field has been set.
func (o *DeploymentDescription) HasDeployStartTime() bool {
	if o != nil && !IsNil(o.DeployStartTime) {
		return true
	}

	return false
}

// SetDeployStartTime gets a reference to the given string and assigns it to the DeployStartTime field.
func (o *DeploymentDescription) SetDeployStartTime(v string) {
	o.DeployStartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DeploymentDescription) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDescription) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DeploymentDescription) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *DeploymentDescription) SetEndTime(v string) {
	o.EndTime = &v
}

func (o DeploymentDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDescription) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StateDetail) {
		toSerialize["stateDetail"] = o.StateDetail
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ValidationStartTime) {
		toSerialize["validationStartTime"] = o.ValidationStartTime
	}
	if !IsNil(o.BuildStartTime) {
		toSerialize["buildStartTime"] = o.BuildStartTime
	}
	if !IsNil(o.DeployStartTime) {
		toSerialize["deployStartTime"] = o.DeployStartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableDeploymentDescription struct {
	value *DeploymentDescription
	isSet bool
}

func (v NullableDeploymentDescription) Get() *DeploymentDescription {
	return v.value
}

func (v *NullableDeploymentDescription) Set(val *DeploymentDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDescription(val *DeploymentDescription) *NullableDeploymentDescription {
	return &NullableDeploymentDescription{value: val, isSet: true}
}

func (v NullableDeploymentDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


