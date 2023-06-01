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

// checks if the C1ApiPolicyV1SelfApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1SelfApproval{}

// C1ApiPolicyV1SelfApproval The SelfApproval message.
type C1ApiPolicyV1SelfApproval struct {
	// The assignedUserIds field.
	AssignedUserIds []string `json:"assignedUserIds,omitempty"`
	// The fallback field.
	Fallback *bool `json:"fallback,omitempty"`
	//  Self approval is the target of the ticket 
	FallbackUserIds []string `json:"fallbackUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1SelfApproval C1ApiPolicyV1SelfApproval

// NewC1ApiPolicyV1SelfApproval instantiates a new C1ApiPolicyV1SelfApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1SelfApproval() *C1ApiPolicyV1SelfApproval {
	this := C1ApiPolicyV1SelfApproval{}
	return &this
}

// NewC1ApiPolicyV1SelfApprovalWithDefaults instantiates a new C1ApiPolicyV1SelfApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1SelfApprovalWithDefaults() *C1ApiPolicyV1SelfApproval {
	this := C1ApiPolicyV1SelfApproval{}
	return &this
}

// GetAssignedUserIds returns the AssignedUserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1SelfApproval) GetAssignedUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssignedUserIds
}

// GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1SelfApproval) GetAssignedUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUserIds) {
		return nil, false
	}
	return o.AssignedUserIds, true
}

// HasAssignedUserIds returns a boolean if a field has been set.
func (o *C1ApiPolicyV1SelfApproval) HasAssignedUserIds() bool {
	if o != nil && IsNil(o.AssignedUserIds) {
		return true
	}

	return false
}

// SetAssignedUserIds gets a reference to the given []string and assigns it to the AssignedUserIds field.
func (o *C1ApiPolicyV1SelfApproval) SetAssignedUserIds(v []string) {
	o.AssignedUserIds = v
}

// GetFallback returns the Fallback field value if set, zero value otherwise.
func (o *C1ApiPolicyV1SelfApproval) GetFallback() bool {
	if o == nil || IsNil(o.Fallback) {
		var ret bool
		return ret
	}
	return *o.Fallback
}

// GetFallbackOk returns a tuple with the Fallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1SelfApproval) GetFallbackOk() (*bool, bool) {
	if o == nil || IsNil(o.Fallback) {
		return nil, false
	}
	return o.Fallback, true
}

// HasFallback returns a boolean if a field has been set.
func (o *C1ApiPolicyV1SelfApproval) HasFallback() bool {
	if o != nil && !IsNil(o.Fallback) {
		return true
	}

	return false
}

// SetFallback gets a reference to the given bool and assigns it to the Fallback field.
func (o *C1ApiPolicyV1SelfApproval) SetFallback(v bool) {
	o.Fallback = &v
}

// GetFallbackUserIds returns the FallbackUserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1SelfApproval) GetFallbackUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FallbackUserIds
}

// GetFallbackUserIdsOk returns a tuple with the FallbackUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1SelfApproval) GetFallbackUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FallbackUserIds) {
		return nil, false
	}
	return o.FallbackUserIds, true
}

// HasFallbackUserIds returns a boolean if a field has been set.
func (o *C1ApiPolicyV1SelfApproval) HasFallbackUserIds() bool {
	if o != nil && IsNil(o.FallbackUserIds) {
		return true
	}

	return false
}

// SetFallbackUserIds gets a reference to the given []string and assigns it to the FallbackUserIds field.
func (o *C1ApiPolicyV1SelfApproval) SetFallbackUserIds(v []string) {
	o.FallbackUserIds = v
}

func (o C1ApiPolicyV1SelfApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1SelfApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedUserIds != nil {
		toSerialize["assignedUserIds"] = o.AssignedUserIds
	}
	if !IsNil(o.Fallback) {
		toSerialize["fallback"] = o.Fallback
	}
	if o.FallbackUserIds != nil {
		toSerialize["fallbackUserIds"] = o.FallbackUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1SelfApproval) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1SelfApproval := _C1ApiPolicyV1SelfApproval{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1SelfApproval); err == nil {
		*o = C1ApiPolicyV1SelfApproval(varC1ApiPolicyV1SelfApproval)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "assignedUserIds")
		delete(additionalProperties, "fallback")
		delete(additionalProperties, "fallbackUserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1SelfApproval struct {
	value *C1ApiPolicyV1SelfApproval
	isSet bool
}

func (v NullableC1ApiPolicyV1SelfApproval) Get() *C1ApiPolicyV1SelfApproval {
	return v.value
}

func (v *NullableC1ApiPolicyV1SelfApproval) Set(val *C1ApiPolicyV1SelfApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1SelfApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1SelfApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1SelfApproval(val *C1ApiPolicyV1SelfApproval) *NullableC1ApiPolicyV1SelfApproval {
	return &NullableC1ApiPolicyV1SelfApproval{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1SelfApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1SelfApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


