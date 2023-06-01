/*
ConductorOne API

The ConductorOne API is a HTTP API for managing ConductorOne resources.

API version: 0.1.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package c1api

import (
	"encoding/json"
	"time"
)

// checks if the C1ApiPolicyV1ReassignedByErrorAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1ReassignedByErrorAction{}

// C1ApiPolicyV1ReassignedByErrorAction The ReassignedByErrorAction message.
type C1ApiPolicyV1ReassignedByErrorAction struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The errorCode field.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The errorUserId field.
	ErrorUserId *string `json:"errorUserId,omitempty"`
	ErroredAt *time.Time `json:"erroredAt,omitempty"`
	// The newPolicyStepId field.
	NewPolicyStepId *string `json:"newPolicyStepId,omitempty"`
	ReassignedAt *time.Time `json:"reassignedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1ReassignedByErrorAction C1ApiPolicyV1ReassignedByErrorAction

// NewC1ApiPolicyV1ReassignedByErrorAction instantiates a new C1ApiPolicyV1ReassignedByErrorAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1ReassignedByErrorAction() *C1ApiPolicyV1ReassignedByErrorAction {
	this := C1ApiPolicyV1ReassignedByErrorAction{}
	return &this
}

// NewC1ApiPolicyV1ReassignedByErrorActionWithDefaults instantiates a new C1ApiPolicyV1ReassignedByErrorAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1ReassignedByErrorActionWithDefaults() *C1ApiPolicyV1ReassignedByErrorAction {
	this := C1ApiPolicyV1ReassignedByErrorAction{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetDescription(v string) {
	o.Description = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorUserId returns the ErrorUserId field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErrorUserId() string {
	if o == nil || IsNil(o.ErrorUserId) {
		var ret string
		return ret
	}
	return *o.ErrorUserId
}

// GetErrorUserIdOk returns a tuple with the ErrorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErrorUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorUserId) {
		return nil, false
	}
	return o.ErrorUserId, true
}

// HasErrorUserId returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasErrorUserId() bool {
	if o != nil && !IsNil(o.ErrorUserId) {
		return true
	}

	return false
}

// SetErrorUserId gets a reference to the given string and assigns it to the ErrorUserId field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetErrorUserId(v string) {
	o.ErrorUserId = &v
}

// GetErroredAt returns the ErroredAt field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErroredAt() time.Time {
	if o == nil || IsNil(o.ErroredAt) {
		var ret time.Time
		return ret
	}
	return *o.ErroredAt
}

// GetErroredAtOk returns a tuple with the ErroredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetErroredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ErroredAt) {
		return nil, false
	}
	return o.ErroredAt, true
}

// HasErroredAt returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasErroredAt() bool {
	if o != nil && !IsNil(o.ErroredAt) {
		return true
	}

	return false
}

// SetErroredAt gets a reference to the given time.Time and assigns it to the ErroredAt field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetErroredAt(v time.Time) {
	o.ErroredAt = &v
}

// GetNewPolicyStepId returns the NewPolicyStepId field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetNewPolicyStepId() string {
	if o == nil || IsNil(o.NewPolicyStepId) {
		var ret string
		return ret
	}
	return *o.NewPolicyStepId
}

// GetNewPolicyStepIdOk returns a tuple with the NewPolicyStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetNewPolicyStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewPolicyStepId) {
		return nil, false
	}
	return o.NewPolicyStepId, true
}

// HasNewPolicyStepId returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasNewPolicyStepId() bool {
	if o != nil && !IsNil(o.NewPolicyStepId) {
		return true
	}

	return false
}

// SetNewPolicyStepId gets a reference to the given string and assigns it to the NewPolicyStepId field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetNewPolicyStepId(v string) {
	o.NewPolicyStepId = &v
}

// GetReassignedAt returns the ReassignedAt field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetReassignedAt() time.Time {
	if o == nil || IsNil(o.ReassignedAt) {
		var ret time.Time
		return ret
	}
	return *o.ReassignedAt
}

// GetReassignedAtOk returns a tuple with the ReassignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) GetReassignedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReassignedAt) {
		return nil, false
	}
	return o.ReassignedAt, true
}

// HasReassignedAt returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ReassignedByErrorAction) HasReassignedAt() bool {
	if o != nil && !IsNil(o.ReassignedAt) {
		return true
	}

	return false
}

// SetReassignedAt gets a reference to the given time.Time and assigns it to the ReassignedAt field.
func (o *C1ApiPolicyV1ReassignedByErrorAction) SetReassignedAt(v time.Time) {
	o.ReassignedAt = &v
}

func (o C1ApiPolicyV1ReassignedByErrorAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1ReassignedByErrorAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorUserId) {
		toSerialize["errorUserId"] = o.ErrorUserId
	}
	if !IsNil(o.ErroredAt) {
		toSerialize["erroredAt"] = o.ErroredAt
	}
	if !IsNil(o.NewPolicyStepId) {
		toSerialize["newPolicyStepId"] = o.NewPolicyStepId
	}
	if !IsNil(o.ReassignedAt) {
		toSerialize["reassignedAt"] = o.ReassignedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1ReassignedByErrorAction) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1ReassignedByErrorAction := _C1ApiPolicyV1ReassignedByErrorAction{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1ReassignedByErrorAction); err == nil {
		*o = C1ApiPolicyV1ReassignedByErrorAction(varC1ApiPolicyV1ReassignedByErrorAction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorUserId")
		delete(additionalProperties, "erroredAt")
		delete(additionalProperties, "newPolicyStepId")
		delete(additionalProperties, "reassignedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1ReassignedByErrorAction struct {
	value *C1ApiPolicyV1ReassignedByErrorAction
	isSet bool
}

func (v NullableC1ApiPolicyV1ReassignedByErrorAction) Get() *C1ApiPolicyV1ReassignedByErrorAction {
	return v.value
}

func (v *NullableC1ApiPolicyV1ReassignedByErrorAction) Set(val *C1ApiPolicyV1ReassignedByErrorAction) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1ReassignedByErrorAction) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1ReassignedByErrorAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1ReassignedByErrorAction(val *C1ApiPolicyV1ReassignedByErrorAction) *NullableC1ApiPolicyV1ReassignedByErrorAction {
	return &NullableC1ApiPolicyV1ReassignedByErrorAction{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1ReassignedByErrorAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1ReassignedByErrorAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


