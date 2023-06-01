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

// checks if the C1ApiTaskV1TaskServiceCreateGrantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1TaskServiceCreateGrantRequest{}

// C1ApiTaskV1TaskServiceCreateGrantRequest The TaskServiceCreateGrantRequest message.
type C1ApiTaskV1TaskServiceCreateGrantRequest struct {
	// The appEntitlementId field.
	AppEntitlementId *string `json:"appEntitlementId,omitempty"`
	// The appId field.
	AppId *string `json:"appId,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	ExpandMask *C1ApiTaskV1TaskExpandMask `json:"expandMask,omitempty"`
	GrantDuration *string `json:"grantDuration,omitempty"`
	// The identityUserId field.
	IdentityUserId *string `json:"identityUserId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1TaskServiceCreateGrantRequest C1ApiTaskV1TaskServiceCreateGrantRequest

// NewC1ApiTaskV1TaskServiceCreateGrantRequest instantiates a new C1ApiTaskV1TaskServiceCreateGrantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1TaskServiceCreateGrantRequest() *C1ApiTaskV1TaskServiceCreateGrantRequest {
	this := C1ApiTaskV1TaskServiceCreateGrantRequest{}
	return &this
}

// NewC1ApiTaskV1TaskServiceCreateGrantRequestWithDefaults instantiates a new C1ApiTaskV1TaskServiceCreateGrantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1TaskServiceCreateGrantRequestWithDefaults() *C1ApiTaskV1TaskServiceCreateGrantRequest {
	this := C1ApiTaskV1TaskServiceCreateGrantRequest{}
	return &this
}

// GetAppEntitlementId returns the AppEntitlementId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetAppEntitlementId() string {
	if o == nil || IsNil(o.AppEntitlementId) {
		var ret string
		return ret
	}
	return *o.AppEntitlementId
}

// GetAppEntitlementIdOk returns a tuple with the AppEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetAppEntitlementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppEntitlementId) {
		return nil, false
	}
	return o.AppEntitlementId, true
}

// HasAppEntitlementId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasAppEntitlementId() bool {
	if o != nil && !IsNil(o.AppEntitlementId) {
		return true
	}

	return false
}

// SetAppEntitlementId gets a reference to the given string and assigns it to the AppEntitlementId field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetAppEntitlementId(v string) {
	o.AppEntitlementId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetAppId(v string) {
	o.AppId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpandMask returns the ExpandMask field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetExpandMask() C1ApiTaskV1TaskExpandMask {
	if o == nil || IsNil(o.ExpandMask) {
		var ret C1ApiTaskV1TaskExpandMask
		return ret
	}
	return *o.ExpandMask
}

// GetExpandMaskOk returns a tuple with the ExpandMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetExpandMaskOk() (*C1ApiTaskV1TaskExpandMask, bool) {
	if o == nil || IsNil(o.ExpandMask) {
		return nil, false
	}
	return o.ExpandMask, true
}

// HasExpandMask returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasExpandMask() bool {
	if o != nil && !IsNil(o.ExpandMask) {
		return true
	}

	return false
}

// SetExpandMask gets a reference to the given C1ApiTaskV1TaskExpandMask and assigns it to the ExpandMask field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetExpandMask(v C1ApiTaskV1TaskExpandMask) {
	o.ExpandMask = &v
}

// GetGrantDuration returns the GrantDuration field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetGrantDuration() string {
	if o == nil || IsNil(o.GrantDuration) {
		var ret string
		return ret
	}
	return *o.GrantDuration
}

// GetGrantDurationOk returns a tuple with the GrantDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetGrantDurationOk() (*string, bool) {
	if o == nil || IsNil(o.GrantDuration) {
		return nil, false
	}
	return o.GrantDuration, true
}

// HasGrantDuration returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasGrantDuration() bool {
	if o != nil && !IsNil(o.GrantDuration) {
		return true
	}

	return false
}

// SetGrantDuration gets a reference to the given string and assigns it to the GrantDuration field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetGrantDuration(v string) {
	o.GrantDuration = &v
}

// GetIdentityUserId returns the IdentityUserId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetIdentityUserId() string {
	if o == nil || IsNil(o.IdentityUserId) {
		var ret string
		return ret
	}
	return *o.IdentityUserId
}

// GetIdentityUserIdOk returns a tuple with the IdentityUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) GetIdentityUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityUserId) {
		return nil, false
	}
	return o.IdentityUserId, true
}

// HasIdentityUserId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) HasIdentityUserId() bool {
	if o != nil && !IsNil(o.IdentityUserId) {
		return true
	}

	return false
}

// SetIdentityUserId gets a reference to the given string and assigns it to the IdentityUserId field.
func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) SetIdentityUserId(v string) {
	o.IdentityUserId = &v
}

func (o C1ApiTaskV1TaskServiceCreateGrantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1TaskServiceCreateGrantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppEntitlementId) {
		toSerialize["appEntitlementId"] = o.AppEntitlementId
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpandMask) {
		toSerialize["expandMask"] = o.ExpandMask
	}
	if !IsNil(o.GrantDuration) {
		toSerialize["grantDuration"] = o.GrantDuration
	}
	if !IsNil(o.IdentityUserId) {
		toSerialize["identityUserId"] = o.IdentityUserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1TaskServiceCreateGrantRequest) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1TaskServiceCreateGrantRequest := _C1ApiTaskV1TaskServiceCreateGrantRequest{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1TaskServiceCreateGrantRequest); err == nil {
		*o = C1ApiTaskV1TaskServiceCreateGrantRequest(varC1ApiTaskV1TaskServiceCreateGrantRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appEntitlementId")
		delete(additionalProperties, "appId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "expandMask")
		delete(additionalProperties, "grantDuration")
		delete(additionalProperties, "identityUserId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1TaskServiceCreateGrantRequest struct {
	value *C1ApiTaskV1TaskServiceCreateGrantRequest
	isSet bool
}

func (v NullableC1ApiTaskV1TaskServiceCreateGrantRequest) Get() *C1ApiTaskV1TaskServiceCreateGrantRequest {
	return v.value
}

func (v *NullableC1ApiTaskV1TaskServiceCreateGrantRequest) Set(val *C1ApiTaskV1TaskServiceCreateGrantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1TaskServiceCreateGrantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1TaskServiceCreateGrantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1TaskServiceCreateGrantRequest(val *C1ApiTaskV1TaskServiceCreateGrantRequest) *NullableC1ApiTaskV1TaskServiceCreateGrantRequest {
	return &NullableC1ApiTaskV1TaskServiceCreateGrantRequest{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1TaskServiceCreateGrantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1TaskServiceCreateGrantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


