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

// checks if the C1ApiPolicyV1Approval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1Approval{}

// C1ApiPolicyV1Approval The Approval message.  This message contains a oneof named typ. Only a single field of the following list may be set at a time:   - users   - manager   - appOwners   - group   - self   - entitlementOwners 
type C1ApiPolicyV1Approval struct {
	// The allowReassignment field.
	AllowReassignment *bool `json:"allowReassignment,omitempty"`
	AppOwners NullableC1ApiPolicyV1AppOwnerApproval `json:"appOwners,omitempty"`
	// The assigned field.
	Assigned *bool `json:"assigned,omitempty"`
	EntitlementOwners NullableC1ApiPolicyV1EntitlementOwnerApproval `json:"entitlementOwners,omitempty"`
	Group NullableC1ApiPolicyV1AppGroupApproval `json:"group,omitempty"`
	Manager NullableC1ApiPolicyV1ManagerApproval `json:"manager,omitempty"`
	// The requireApprovalReason field.
	RequireApprovalReason *bool `json:"requireApprovalReason,omitempty"`
	// The requireReassignmentReason field.
	RequireReassignmentReason *bool `json:"requireReassignmentReason,omitempty"`
	Self NullableC1ApiPolicyV1SelfApproval `json:"self,omitempty"`
	Users NullableC1ApiPolicyV1UserApproval `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1Approval C1ApiPolicyV1Approval

// NewC1ApiPolicyV1Approval instantiates a new C1ApiPolicyV1Approval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1Approval() *C1ApiPolicyV1Approval {
	this := C1ApiPolicyV1Approval{}
	return &this
}

// NewC1ApiPolicyV1ApprovalWithDefaults instantiates a new C1ApiPolicyV1Approval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1ApprovalWithDefaults() *C1ApiPolicyV1Approval {
	this := C1ApiPolicyV1Approval{}
	return &this
}

// GetAllowReassignment returns the AllowReassignment field value if set, zero value otherwise.
func (o *C1ApiPolicyV1Approval) GetAllowReassignment() bool {
	if o == nil || IsNil(o.AllowReassignment) {
		var ret bool
		return ret
	}
	return *o.AllowReassignment
}

// GetAllowReassignmentOk returns a tuple with the AllowReassignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1Approval) GetAllowReassignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowReassignment) {
		return nil, false
	}
	return o.AllowReassignment, true
}

// HasAllowReassignment returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasAllowReassignment() bool {
	if o != nil && !IsNil(o.AllowReassignment) {
		return true
	}

	return false
}

// SetAllowReassignment gets a reference to the given bool and assigns it to the AllowReassignment field.
func (o *C1ApiPolicyV1Approval) SetAllowReassignment(v bool) {
	o.AllowReassignment = &v
}

// GetAppOwners returns the AppOwners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetAppOwners() C1ApiPolicyV1AppOwnerApproval {
	if o == nil || IsNil(o.AppOwners.Get()) {
		var ret C1ApiPolicyV1AppOwnerApproval
		return ret
	}
	return *o.AppOwners.Get()
}

// GetAppOwnersOk returns a tuple with the AppOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetAppOwnersOk() (*C1ApiPolicyV1AppOwnerApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppOwners.Get(), o.AppOwners.IsSet()
}

// HasAppOwners returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasAppOwners() bool {
	if o != nil && o.AppOwners.IsSet() {
		return true
	}

	return false
}

// SetAppOwners gets a reference to the given NullableC1ApiPolicyV1AppOwnerApproval and assigns it to the AppOwners field.
func (o *C1ApiPolicyV1Approval) SetAppOwners(v C1ApiPolicyV1AppOwnerApproval) {
	o.AppOwners.Set(&v)
}
// SetAppOwnersNil sets the value for AppOwners to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetAppOwnersNil() {
	o.AppOwners.Set(nil)
}

// UnsetAppOwners ensures that no value is present for AppOwners, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetAppOwners() {
	o.AppOwners.Unset()
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *C1ApiPolicyV1Approval) GetAssigned() bool {
	if o == nil || IsNil(o.Assigned) {
		var ret bool
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1Approval) GetAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given bool and assigns it to the Assigned field.
func (o *C1ApiPolicyV1Approval) SetAssigned(v bool) {
	o.Assigned = &v
}

// GetEntitlementOwners returns the EntitlementOwners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetEntitlementOwners() C1ApiPolicyV1EntitlementOwnerApproval {
	if o == nil || IsNil(o.EntitlementOwners.Get()) {
		var ret C1ApiPolicyV1EntitlementOwnerApproval
		return ret
	}
	return *o.EntitlementOwners.Get()
}

// GetEntitlementOwnersOk returns a tuple with the EntitlementOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetEntitlementOwnersOk() (*C1ApiPolicyV1EntitlementOwnerApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementOwners.Get(), o.EntitlementOwners.IsSet()
}

// HasEntitlementOwners returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasEntitlementOwners() bool {
	if o != nil && o.EntitlementOwners.IsSet() {
		return true
	}

	return false
}

// SetEntitlementOwners gets a reference to the given NullableC1ApiPolicyV1EntitlementOwnerApproval and assigns it to the EntitlementOwners field.
func (o *C1ApiPolicyV1Approval) SetEntitlementOwners(v C1ApiPolicyV1EntitlementOwnerApproval) {
	o.EntitlementOwners.Set(&v)
}
// SetEntitlementOwnersNil sets the value for EntitlementOwners to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetEntitlementOwnersNil() {
	o.EntitlementOwners.Set(nil)
}

// UnsetEntitlementOwners ensures that no value is present for EntitlementOwners, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetEntitlementOwners() {
	o.EntitlementOwners.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetGroup() C1ApiPolicyV1AppGroupApproval {
	if o == nil || IsNil(o.Group.Get()) {
		var ret C1ApiPolicyV1AppGroupApproval
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetGroupOk() (*C1ApiPolicyV1AppGroupApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableC1ApiPolicyV1AppGroupApproval and assigns it to the Group field.
func (o *C1ApiPolicyV1Approval) SetGroup(v C1ApiPolicyV1AppGroupApproval) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetGroup() {
	o.Group.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetManager() C1ApiPolicyV1ManagerApproval {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret C1ApiPolicyV1ManagerApproval
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetManagerOk() (*C1ApiPolicyV1ManagerApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableC1ApiPolicyV1ManagerApproval and assigns it to the Manager field.
func (o *C1ApiPolicyV1Approval) SetManager(v C1ApiPolicyV1ManagerApproval) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetManager() {
	o.Manager.Unset()
}

// GetRequireApprovalReason returns the RequireApprovalReason field value if set, zero value otherwise.
func (o *C1ApiPolicyV1Approval) GetRequireApprovalReason() bool {
	if o == nil || IsNil(o.RequireApprovalReason) {
		var ret bool
		return ret
	}
	return *o.RequireApprovalReason
}

// GetRequireApprovalReasonOk returns a tuple with the RequireApprovalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1Approval) GetRequireApprovalReasonOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireApprovalReason) {
		return nil, false
	}
	return o.RequireApprovalReason, true
}

// HasRequireApprovalReason returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasRequireApprovalReason() bool {
	if o != nil && !IsNil(o.RequireApprovalReason) {
		return true
	}

	return false
}

// SetRequireApprovalReason gets a reference to the given bool and assigns it to the RequireApprovalReason field.
func (o *C1ApiPolicyV1Approval) SetRequireApprovalReason(v bool) {
	o.RequireApprovalReason = &v
}

// GetRequireReassignmentReason returns the RequireReassignmentReason field value if set, zero value otherwise.
func (o *C1ApiPolicyV1Approval) GetRequireReassignmentReason() bool {
	if o == nil || IsNil(o.RequireReassignmentReason) {
		var ret bool
		return ret
	}
	return *o.RequireReassignmentReason
}

// GetRequireReassignmentReasonOk returns a tuple with the RequireReassignmentReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1Approval) GetRequireReassignmentReasonOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireReassignmentReason) {
		return nil, false
	}
	return o.RequireReassignmentReason, true
}

// HasRequireReassignmentReason returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasRequireReassignmentReason() bool {
	if o != nil && !IsNil(o.RequireReassignmentReason) {
		return true
	}

	return false
}

// SetRequireReassignmentReason gets a reference to the given bool and assigns it to the RequireReassignmentReason field.
func (o *C1ApiPolicyV1Approval) SetRequireReassignmentReason(v bool) {
	o.RequireReassignmentReason = &v
}

// GetSelf returns the Self field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetSelf() C1ApiPolicyV1SelfApproval {
	if o == nil || IsNil(o.Self.Get()) {
		var ret C1ApiPolicyV1SelfApproval
		return ret
	}
	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetSelfOk() (*C1ApiPolicyV1SelfApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// HasSelf returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasSelf() bool {
	if o != nil && o.Self.IsSet() {
		return true
	}

	return false
}

// SetSelf gets a reference to the given NullableC1ApiPolicyV1SelfApproval and assigns it to the Self field.
func (o *C1ApiPolicyV1Approval) SetSelf(v C1ApiPolicyV1SelfApproval) {
	o.Self.Set(&v)
}
// SetSelfNil sets the value for Self to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetSelfNil() {
	o.Self.Set(nil)
}

// UnsetSelf ensures that no value is present for Self, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetSelf() {
	o.Self.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1Approval) GetUsers() C1ApiPolicyV1UserApproval {
	if o == nil || IsNil(o.Users.Get()) {
		var ret C1ApiPolicyV1UserApproval
		return ret
	}
	return *o.Users.Get()
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1Approval) GetUsersOk() (*C1ApiPolicyV1UserApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users.Get(), o.Users.IsSet()
}

// HasUsers returns a boolean if a field has been set.
func (o *C1ApiPolicyV1Approval) HasUsers() bool {
	if o != nil && o.Users.IsSet() {
		return true
	}

	return false
}

// SetUsers gets a reference to the given NullableC1ApiPolicyV1UserApproval and assigns it to the Users field.
func (o *C1ApiPolicyV1Approval) SetUsers(v C1ApiPolicyV1UserApproval) {
	o.Users.Set(&v)
}
// SetUsersNil sets the value for Users to be an explicit nil
func (o *C1ApiPolicyV1Approval) SetUsersNil() {
	o.Users.Set(nil)
}

// UnsetUsers ensures that no value is present for Users, not even an explicit nil
func (o *C1ApiPolicyV1Approval) UnsetUsers() {
	o.Users.Unset()
}

func (o C1ApiPolicyV1Approval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1Approval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowReassignment) {
		toSerialize["allowReassignment"] = o.AllowReassignment
	}
	if o.AppOwners.IsSet() {
		toSerialize["appOwners"] = o.AppOwners.Get()
	}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if o.EntitlementOwners.IsSet() {
		toSerialize["entitlementOwners"] = o.EntitlementOwners.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if !IsNil(o.RequireApprovalReason) {
		toSerialize["requireApprovalReason"] = o.RequireApprovalReason
	}
	if !IsNil(o.RequireReassignmentReason) {
		toSerialize["requireReassignmentReason"] = o.RequireReassignmentReason
	}
	if o.Self.IsSet() {
		toSerialize["self"] = o.Self.Get()
	}
	if o.Users.IsSet() {
		toSerialize["users"] = o.Users.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1Approval) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1Approval := _C1ApiPolicyV1Approval{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1Approval); err == nil {
		*o = C1ApiPolicyV1Approval(varC1ApiPolicyV1Approval)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowReassignment")
		delete(additionalProperties, "appOwners")
		delete(additionalProperties, "assigned")
		delete(additionalProperties, "entitlementOwners")
		delete(additionalProperties, "group")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "requireApprovalReason")
		delete(additionalProperties, "requireReassignmentReason")
		delete(additionalProperties, "self")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1Approval struct {
	value *C1ApiPolicyV1Approval
	isSet bool
}

func (v NullableC1ApiPolicyV1Approval) Get() *C1ApiPolicyV1Approval {
	return v.value
}

func (v *NullableC1ApiPolicyV1Approval) Set(val *C1ApiPolicyV1Approval) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1Approval) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1Approval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1Approval(val *C1ApiPolicyV1Approval) *NullableC1ApiPolicyV1Approval {
	return &NullableC1ApiPolicyV1Approval{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1Approval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1Approval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


