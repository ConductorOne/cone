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

// checks if the C1ApiPolicyV1ApprovalInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1ApprovalInstance{}

// C1ApiPolicyV1ApprovalInstance The ApprovalInstance message.  This message contains a oneof named outcome. Only a single field of the following list may be set at a time:   - approved   - denied   - reassigned   - restarted   - reassignedByError 
type C1ApiPolicyV1ApprovalInstance struct {
	Approval NullableC1ApiPolicyV1Approval `json:"approval,omitempty"`
	Approved NullableC1ApiPolicyV1ApprovedAction `json:"approved,omitempty"`
	Denied NullableC1ApiPolicyV1DeniedAction `json:"denied,omitempty"`
	Reassigned NullableC1ApiPolicyV1ReassignedAction `json:"reassigned,omitempty"`
	ReassignedByError NullableC1ApiPolicyV1ReassignedByErrorAction `json:"reassignedByError,omitempty"`
	Restarted NullableC1ApiPolicyV1RestartAction `json:"restarted,omitempty"`
	// The state field.
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1ApprovalInstance C1ApiPolicyV1ApprovalInstance

// NewC1ApiPolicyV1ApprovalInstance instantiates a new C1ApiPolicyV1ApprovalInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1ApprovalInstance() *C1ApiPolicyV1ApprovalInstance {
	this := C1ApiPolicyV1ApprovalInstance{}
	return &this
}

// NewC1ApiPolicyV1ApprovalInstanceWithDefaults instantiates a new C1ApiPolicyV1ApprovalInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1ApprovalInstanceWithDefaults() *C1ApiPolicyV1ApprovalInstance {
	this := C1ApiPolicyV1ApprovalInstance{}
	return &this
}

// GetApproval returns the Approval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetApproval() C1ApiPolicyV1Approval {
	if o == nil || IsNil(o.Approval.Get()) {
		var ret C1ApiPolicyV1Approval
		return ret
	}
	return *o.Approval.Get()
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetApprovalOk() (*C1ApiPolicyV1Approval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approval.Get(), o.Approval.IsSet()
}

// HasApproval returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasApproval() bool {
	if o != nil && o.Approval.IsSet() {
		return true
	}

	return false
}

// SetApproval gets a reference to the given NullableC1ApiPolicyV1Approval and assigns it to the Approval field.
func (o *C1ApiPolicyV1ApprovalInstance) SetApproval(v C1ApiPolicyV1Approval) {
	o.Approval.Set(&v)
}
// SetApprovalNil sets the value for Approval to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetApprovalNil() {
	o.Approval.Set(nil)
}

// UnsetApproval ensures that no value is present for Approval, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetApproval() {
	o.Approval.Unset()
}

// GetApproved returns the Approved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetApproved() C1ApiPolicyV1ApprovedAction {
	if o == nil || IsNil(o.Approved.Get()) {
		var ret C1ApiPolicyV1ApprovedAction
		return ret
	}
	return *o.Approved.Get()
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetApprovedOk() (*C1ApiPolicyV1ApprovedAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approved.Get(), o.Approved.IsSet()
}

// HasApproved returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasApproved() bool {
	if o != nil && o.Approved.IsSet() {
		return true
	}

	return false
}

// SetApproved gets a reference to the given NullableC1ApiPolicyV1ApprovedAction and assigns it to the Approved field.
func (o *C1ApiPolicyV1ApprovalInstance) SetApproved(v C1ApiPolicyV1ApprovedAction) {
	o.Approved.Set(&v)
}
// SetApprovedNil sets the value for Approved to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetApprovedNil() {
	o.Approved.Set(nil)
}

// UnsetApproved ensures that no value is present for Approved, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetApproved() {
	o.Approved.Unset()
}

// GetDenied returns the Denied field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetDenied() C1ApiPolicyV1DeniedAction {
	if o == nil || IsNil(o.Denied.Get()) {
		var ret C1ApiPolicyV1DeniedAction
		return ret
	}
	return *o.Denied.Get()
}

// GetDeniedOk returns a tuple with the Denied field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetDeniedOk() (*C1ApiPolicyV1DeniedAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Denied.Get(), o.Denied.IsSet()
}

// HasDenied returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasDenied() bool {
	if o != nil && o.Denied.IsSet() {
		return true
	}

	return false
}

// SetDenied gets a reference to the given NullableC1ApiPolicyV1DeniedAction and assigns it to the Denied field.
func (o *C1ApiPolicyV1ApprovalInstance) SetDenied(v C1ApiPolicyV1DeniedAction) {
	o.Denied.Set(&v)
}
// SetDeniedNil sets the value for Denied to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetDeniedNil() {
	o.Denied.Set(nil)
}

// UnsetDenied ensures that no value is present for Denied, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetDenied() {
	o.Denied.Unset()
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetReassigned() C1ApiPolicyV1ReassignedAction {
	if o == nil || IsNil(o.Reassigned.Get()) {
		var ret C1ApiPolicyV1ReassignedAction
		return ret
	}
	return *o.Reassigned.Get()
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedOk() (*C1ApiPolicyV1ReassignedAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reassigned.Get(), o.Reassigned.IsSet()
}

// HasReassigned returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasReassigned() bool {
	if o != nil && o.Reassigned.IsSet() {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given NullableC1ApiPolicyV1ReassignedAction and assigns it to the Reassigned field.
func (o *C1ApiPolicyV1ApprovalInstance) SetReassigned(v C1ApiPolicyV1ReassignedAction) {
	o.Reassigned.Set(&v)
}
// SetReassignedNil sets the value for Reassigned to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedNil() {
	o.Reassigned.Set(nil)
}

// UnsetReassigned ensures that no value is present for Reassigned, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetReassigned() {
	o.Reassigned.Unset()
}

// GetReassignedByError returns the ReassignedByError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedByError() C1ApiPolicyV1ReassignedByErrorAction {
	if o == nil || IsNil(o.ReassignedByError.Get()) {
		var ret C1ApiPolicyV1ReassignedByErrorAction
		return ret
	}
	return *o.ReassignedByError.Get()
}

// GetReassignedByErrorOk returns a tuple with the ReassignedByError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedByErrorOk() (*C1ApiPolicyV1ReassignedByErrorAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReassignedByError.Get(), o.ReassignedByError.IsSet()
}

// HasReassignedByError returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasReassignedByError() bool {
	if o != nil && o.ReassignedByError.IsSet() {
		return true
	}

	return false
}

// SetReassignedByError gets a reference to the given NullableC1ApiPolicyV1ReassignedByErrorAction and assigns it to the ReassignedByError field.
func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedByError(v C1ApiPolicyV1ReassignedByErrorAction) {
	o.ReassignedByError.Set(&v)
}
// SetReassignedByErrorNil sets the value for ReassignedByError to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedByErrorNil() {
	o.ReassignedByError.Set(nil)
}

// UnsetReassignedByError ensures that no value is present for ReassignedByError, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetReassignedByError() {
	o.ReassignedByError.Unset()
}

// GetRestarted returns the Restarted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1ApprovalInstance) GetRestarted() C1ApiPolicyV1RestartAction {
	if o == nil || IsNil(o.Restarted.Get()) {
		var ret C1ApiPolicyV1RestartAction
		return ret
	}
	return *o.Restarted.Get()
}

// GetRestartedOk returns a tuple with the Restarted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1ApprovalInstance) GetRestartedOk() (*C1ApiPolicyV1RestartAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Restarted.Get(), o.Restarted.IsSet()
}

// HasRestarted returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasRestarted() bool {
	if o != nil && o.Restarted.IsSet() {
		return true
	}

	return false
}

// SetRestarted gets a reference to the given NullableC1ApiPolicyV1RestartAction and assigns it to the Restarted field.
func (o *C1ApiPolicyV1ApprovalInstance) SetRestarted(v C1ApiPolicyV1RestartAction) {
	o.Restarted.Set(&v)
}
// SetRestartedNil sets the value for Restarted to be an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) SetRestartedNil() {
	o.Restarted.Set(nil)
}

// UnsetRestarted ensures that no value is present for Restarted, not even an explicit nil
func (o *C1ApiPolicyV1ApprovalInstance) UnsetRestarted() {
	o.Restarted.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *C1ApiPolicyV1ApprovalInstance) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1ApprovalInstance) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *C1ApiPolicyV1ApprovalInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *C1ApiPolicyV1ApprovalInstance) SetState(v string) {
	o.State = &v
}

func (o C1ApiPolicyV1ApprovalInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1ApprovalInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Approval.IsSet() {
		toSerialize["approval"] = o.Approval.Get()
	}
	if o.Approved.IsSet() {
		toSerialize["approved"] = o.Approved.Get()
	}
	if o.Denied.IsSet() {
		toSerialize["denied"] = o.Denied.Get()
	}
	if o.Reassigned.IsSet() {
		toSerialize["reassigned"] = o.Reassigned.Get()
	}
	if o.ReassignedByError.IsSet() {
		toSerialize["reassignedByError"] = o.ReassignedByError.Get()
	}
	if o.Restarted.IsSet() {
		toSerialize["restarted"] = o.Restarted.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1ApprovalInstance) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1ApprovalInstance := _C1ApiPolicyV1ApprovalInstance{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1ApprovalInstance); err == nil {
		*o = C1ApiPolicyV1ApprovalInstance(varC1ApiPolicyV1ApprovalInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approval")
		delete(additionalProperties, "approved")
		delete(additionalProperties, "denied")
		delete(additionalProperties, "reassigned")
		delete(additionalProperties, "reassignedByError")
		delete(additionalProperties, "restarted")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1ApprovalInstance struct {
	value *C1ApiPolicyV1ApprovalInstance
	isSet bool
}

func (v NullableC1ApiPolicyV1ApprovalInstance) Get() *C1ApiPolicyV1ApprovalInstance {
	return v.value
}

func (v *NullableC1ApiPolicyV1ApprovalInstance) Set(val *C1ApiPolicyV1ApprovalInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1ApprovalInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1ApprovalInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1ApprovalInstance(val *C1ApiPolicyV1ApprovalInstance) *NullableC1ApiPolicyV1ApprovalInstance {
	return &NullableC1ApiPolicyV1ApprovalInstance{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1ApprovalInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1ApprovalInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


