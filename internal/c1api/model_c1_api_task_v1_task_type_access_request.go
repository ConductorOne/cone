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

// checks if the C1ApiTaskV1TaskTypeAccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1TaskTypeAccessRequest{}

// C1ApiTaskV1TaskTypeAccessRequest The TaskTypeAccessRequest message.
type C1ApiTaskV1TaskTypeAccessRequest struct {
	AppProfile *C1ApiAppV1AppProfile `json:"appProfile,omitempty"`
	// The appUserId field.
	AppUserId *string `json:"appUserId,omitempty"`
	// The identityUserId field.
	IdentityUserId *string `json:"identityUserId,omitempty"`
	// The outcome field.
	Outcome *string `json:"outcome,omitempty"`
	OutcomeTime *time.Time `json:"outcomeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1TaskTypeAccessRequest C1ApiTaskV1TaskTypeAccessRequest

// NewC1ApiTaskV1TaskTypeAccessRequest instantiates a new C1ApiTaskV1TaskTypeAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1TaskTypeAccessRequest() *C1ApiTaskV1TaskTypeAccessRequest {
	this := C1ApiTaskV1TaskTypeAccessRequest{}
	return &this
}

// NewC1ApiTaskV1TaskTypeAccessRequestWithDefaults instantiates a new C1ApiTaskV1TaskTypeAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1TaskTypeAccessRequestWithDefaults() *C1ApiTaskV1TaskTypeAccessRequest {
	this := C1ApiTaskV1TaskTypeAccessRequest{}
	return &this
}

// GetAppProfile returns the AppProfile field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppProfile() C1ApiAppV1AppProfile {
	if o == nil || IsNil(o.AppProfile) {
		var ret C1ApiAppV1AppProfile
		return ret
	}
	return *o.AppProfile
}

// GetAppProfileOk returns a tuple with the AppProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppProfileOk() (*C1ApiAppV1AppProfile, bool) {
	if o == nil || IsNil(o.AppProfile) {
		return nil, false
	}
	return o.AppProfile, true
}

// HasAppProfile returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) HasAppProfile() bool {
	if o != nil && !IsNil(o.AppProfile) {
		return true
	}

	return false
}

// SetAppProfile gets a reference to the given C1ApiAppV1AppProfile and assigns it to the AppProfile field.
func (o *C1ApiTaskV1TaskTypeAccessRequest) SetAppProfile(v C1ApiAppV1AppProfile) {
	o.AppProfile = &v
}

// GetAppUserId returns the AppUserId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppUserId() string {
	if o == nil || IsNil(o.AppUserId) {
		var ret string
		return ret
	}
	return *o.AppUserId
}

// GetAppUserIdOk returns a tuple with the AppUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppUserId) {
		return nil, false
	}
	return o.AppUserId, true
}

// HasAppUserId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) HasAppUserId() bool {
	if o != nil && !IsNil(o.AppUserId) {
		return true
	}

	return false
}

// SetAppUserId gets a reference to the given string and assigns it to the AppUserId field.
func (o *C1ApiTaskV1TaskTypeAccessRequest) SetAppUserId(v string) {
	o.AppUserId = &v
}

// GetIdentityUserId returns the IdentityUserId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetIdentityUserId() string {
	if o == nil || IsNil(o.IdentityUserId) {
		var ret string
		return ret
	}
	return *o.IdentityUserId
}

// GetIdentityUserIdOk returns a tuple with the IdentityUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetIdentityUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityUserId) {
		return nil, false
	}
	return o.IdentityUserId, true
}

// HasIdentityUserId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) HasIdentityUserId() bool {
	if o != nil && !IsNil(o.IdentityUserId) {
		return true
	}

	return false
}

// SetIdentityUserId gets a reference to the given string and assigns it to the IdentityUserId field.
func (o *C1ApiTaskV1TaskTypeAccessRequest) SetIdentityUserId(v string) {
	o.IdentityUserId = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *C1ApiTaskV1TaskTypeAccessRequest) SetOutcome(v string) {
	o.Outcome = &v
}

// GetOutcomeTime returns the OutcomeTime field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeTime() time.Time {
	if o == nil || IsNil(o.OutcomeTime) {
		var ret time.Time
		return ret
	}
	return *o.OutcomeTime
}

// GetOutcomeTimeOk returns a tuple with the OutcomeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OutcomeTime) {
		return nil, false
	}
	return o.OutcomeTime, true
}

// HasOutcomeTime returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskTypeAccessRequest) HasOutcomeTime() bool {
	if o != nil && !IsNil(o.OutcomeTime) {
		return true
	}

	return false
}

// SetOutcomeTime gets a reference to the given time.Time and assigns it to the OutcomeTime field.
func (o *C1ApiTaskV1TaskTypeAccessRequest) SetOutcomeTime(v time.Time) {
	o.OutcomeTime = &v
}

func (o C1ApiTaskV1TaskTypeAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1TaskTypeAccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppProfile) {
		toSerialize["appProfile"] = o.AppProfile
	}
	if !IsNil(o.AppUserId) {
		toSerialize["appUserId"] = o.AppUserId
	}
	if !IsNil(o.IdentityUserId) {
		toSerialize["identityUserId"] = o.IdentityUserId
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.OutcomeTime) {
		toSerialize["outcomeTime"] = o.OutcomeTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1TaskTypeAccessRequest) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1TaskTypeAccessRequest := _C1ApiTaskV1TaskTypeAccessRequest{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1TaskTypeAccessRequest); err == nil {
		*o = C1ApiTaskV1TaskTypeAccessRequest(varC1ApiTaskV1TaskTypeAccessRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appProfile")
		delete(additionalProperties, "appUserId")
		delete(additionalProperties, "identityUserId")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "outcomeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1TaskTypeAccessRequest struct {
	value *C1ApiTaskV1TaskTypeAccessRequest
	isSet bool
}

func (v NullableC1ApiTaskV1TaskTypeAccessRequest) Get() *C1ApiTaskV1TaskTypeAccessRequest {
	return v.value
}

func (v *NullableC1ApiTaskV1TaskTypeAccessRequest) Set(val *C1ApiTaskV1TaskTypeAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1TaskTypeAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1TaskTypeAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1TaskTypeAccessRequest(val *C1ApiTaskV1TaskTypeAccessRequest) *NullableC1ApiTaskV1TaskTypeAccessRequest {
	return &NullableC1ApiTaskV1TaskTypeAccessRequest{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1TaskTypeAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1TaskTypeAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


