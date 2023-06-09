/*
ConductorOne API

The ConductorOne API is a HTTP API for managing ConductorOne resources.

API version: 0.1.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package c1api

import (
	"encoding/json"
)

// checks if the C1ApiTaskV1TaskActionsServiceDenyRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1TaskActionsServiceDenyRequestInput{}

// C1ApiTaskV1TaskActionsServiceDenyRequestInput The TaskActionsServiceDenyRequest message.
type C1ApiTaskV1TaskActionsServiceDenyRequestInput struct {
	// The comment field.
	Comment *string `json:"comment,omitempty"`
	ExpandMask *C1ApiTaskV1TaskExpandMask `json:"expandMask,omitempty"`
	// The policyStepId field.
	PolicyStepId *string `json:"policyStepId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1TaskActionsServiceDenyRequestInput C1ApiTaskV1TaskActionsServiceDenyRequestInput

// NewC1ApiTaskV1TaskActionsServiceDenyRequestInput instantiates a new C1ApiTaskV1TaskActionsServiceDenyRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1TaskActionsServiceDenyRequestInput() *C1ApiTaskV1TaskActionsServiceDenyRequestInput {
	this := C1ApiTaskV1TaskActionsServiceDenyRequestInput{}
	return &this
}

// NewC1ApiTaskV1TaskActionsServiceDenyRequestInputWithDefaults instantiates a new C1ApiTaskV1TaskActionsServiceDenyRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1TaskActionsServiceDenyRequestInputWithDefaults() *C1ApiTaskV1TaskActionsServiceDenyRequestInput {
	this := C1ApiTaskV1TaskActionsServiceDenyRequestInput{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) SetComment(v string) {
	o.Comment = &v
}

// GetExpandMask returns the ExpandMask field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetExpandMask() C1ApiTaskV1TaskExpandMask {
	if o == nil || IsNil(o.ExpandMask) {
		var ret C1ApiTaskV1TaskExpandMask
		return ret
	}
	return *o.ExpandMask
}

// GetExpandMaskOk returns a tuple with the ExpandMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetExpandMaskOk() (*C1ApiTaskV1TaskExpandMask, bool) {
	if o == nil || IsNil(o.ExpandMask) {
		return nil, false
	}
	return o.ExpandMask, true
}

// HasExpandMask returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) HasExpandMask() bool {
	if o != nil && !IsNil(o.ExpandMask) {
		return true
	}

	return false
}

// SetExpandMask gets a reference to the given C1ApiTaskV1TaskExpandMask and assigns it to the ExpandMask field.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) SetExpandMask(v C1ApiTaskV1TaskExpandMask) {
	o.ExpandMask = &v
}

// GetPolicyStepId returns the PolicyStepId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetPolicyStepId() string {
	if o == nil || IsNil(o.PolicyStepId) {
		var ret string
		return ret
	}
	return *o.PolicyStepId
}

// GetPolicyStepIdOk returns a tuple with the PolicyStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) GetPolicyStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyStepId) {
		return nil, false
	}
	return o.PolicyStepId, true
}

// HasPolicyStepId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) HasPolicyStepId() bool {
	if o != nil && !IsNil(o.PolicyStepId) {
		return true
	}

	return false
}

// SetPolicyStepId gets a reference to the given string and assigns it to the PolicyStepId field.
func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) SetPolicyStepId(v string) {
	o.PolicyStepId = &v
}

func (o C1ApiTaskV1TaskActionsServiceDenyRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1TaskActionsServiceDenyRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ExpandMask) {
		toSerialize["expandMask"] = o.ExpandMask
	}
	if !IsNil(o.PolicyStepId) {
		toSerialize["policyStepId"] = o.PolicyStepId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1TaskActionsServiceDenyRequestInput) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1TaskActionsServiceDenyRequestInput := _C1ApiTaskV1TaskActionsServiceDenyRequestInput{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1TaskActionsServiceDenyRequestInput); err == nil {
		*o = C1ApiTaskV1TaskActionsServiceDenyRequestInput(varC1ApiTaskV1TaskActionsServiceDenyRequestInput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "expandMask")
		delete(additionalProperties, "policyStepId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput struct {
	value *C1ApiTaskV1TaskActionsServiceDenyRequestInput
	isSet bool
}

func (v NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) Get() *C1ApiTaskV1TaskActionsServiceDenyRequestInput {
	return v.value
}

func (v *NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) Set(val *C1ApiTaskV1TaskActionsServiceDenyRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1TaskActionsServiceDenyRequestInput(val *C1ApiTaskV1TaskActionsServiceDenyRequestInput) *NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput {
	return &NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1TaskActionsServiceDenyRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

