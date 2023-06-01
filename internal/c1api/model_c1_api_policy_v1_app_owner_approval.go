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

// checks if the C1ApiPolicyV1AppOwnerApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1AppOwnerApproval{}

// C1ApiPolicyV1AppOwnerApproval The AppOwnerApproval message.
type C1ApiPolicyV1AppOwnerApproval struct {
	//  App owner is based on the app id and doesn't need to have self-contained data 
	AllowSelfApproval *bool `json:"allowSelfApproval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1AppOwnerApproval C1ApiPolicyV1AppOwnerApproval

// NewC1ApiPolicyV1AppOwnerApproval instantiates a new C1ApiPolicyV1AppOwnerApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1AppOwnerApproval() *C1ApiPolicyV1AppOwnerApproval {
	this := C1ApiPolicyV1AppOwnerApproval{}
	return &this
}

// NewC1ApiPolicyV1AppOwnerApprovalWithDefaults instantiates a new C1ApiPolicyV1AppOwnerApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1AppOwnerApprovalWithDefaults() *C1ApiPolicyV1AppOwnerApproval {
	this := C1ApiPolicyV1AppOwnerApproval{}
	return &this
}

// GetAllowSelfApproval returns the AllowSelfApproval field value if set, zero value otherwise.
func (o *C1ApiPolicyV1AppOwnerApproval) GetAllowSelfApproval() bool {
	if o == nil || IsNil(o.AllowSelfApproval) {
		var ret bool
		return ret
	}
	return *o.AllowSelfApproval
}

// GetAllowSelfApprovalOk returns a tuple with the AllowSelfApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1AppOwnerApproval) GetAllowSelfApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSelfApproval) {
		return nil, false
	}
	return o.AllowSelfApproval, true
}

// HasAllowSelfApproval returns a boolean if a field has been set.
func (o *C1ApiPolicyV1AppOwnerApproval) HasAllowSelfApproval() bool {
	if o != nil && !IsNil(o.AllowSelfApproval) {
		return true
	}

	return false
}

// SetAllowSelfApproval gets a reference to the given bool and assigns it to the AllowSelfApproval field.
func (o *C1ApiPolicyV1AppOwnerApproval) SetAllowSelfApproval(v bool) {
	o.AllowSelfApproval = &v
}

func (o C1ApiPolicyV1AppOwnerApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1AppOwnerApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowSelfApproval) {
		toSerialize["allowSelfApproval"] = o.AllowSelfApproval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1AppOwnerApproval) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1AppOwnerApproval := _C1ApiPolicyV1AppOwnerApproval{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1AppOwnerApproval); err == nil {
		*o = C1ApiPolicyV1AppOwnerApproval(varC1ApiPolicyV1AppOwnerApproval)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowSelfApproval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1AppOwnerApproval struct {
	value *C1ApiPolicyV1AppOwnerApproval
	isSet bool
}

func (v NullableC1ApiPolicyV1AppOwnerApproval) Get() *C1ApiPolicyV1AppOwnerApproval {
	return v.value
}

func (v *NullableC1ApiPolicyV1AppOwnerApproval) Set(val *C1ApiPolicyV1AppOwnerApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1AppOwnerApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1AppOwnerApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1AppOwnerApproval(val *C1ApiPolicyV1AppOwnerApproval) *NullableC1ApiPolicyV1AppOwnerApproval {
	return &NullableC1ApiPolicyV1AppOwnerApproval{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1AppOwnerApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1AppOwnerApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


