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

// checks if the C1ApiTaskV1TaskRevokeSourceReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1TaskRevokeSourceReview{}

// C1ApiTaskV1TaskRevokeSourceReview The TaskRevokeSourceReview message.
type C1ApiTaskV1TaskRevokeSourceReview struct {
	// The accessReviewId field.
	AccessReviewId *string `json:"accessReviewId,omitempty"`
	// The certTicketId field.
	CertTicketId *string `json:"certTicketId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1TaskRevokeSourceReview C1ApiTaskV1TaskRevokeSourceReview

// NewC1ApiTaskV1TaskRevokeSourceReview instantiates a new C1ApiTaskV1TaskRevokeSourceReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1TaskRevokeSourceReview() *C1ApiTaskV1TaskRevokeSourceReview {
	this := C1ApiTaskV1TaskRevokeSourceReview{}
	return &this
}

// NewC1ApiTaskV1TaskRevokeSourceReviewWithDefaults instantiates a new C1ApiTaskV1TaskRevokeSourceReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1TaskRevokeSourceReviewWithDefaults() *C1ApiTaskV1TaskRevokeSourceReview {
	this := C1ApiTaskV1TaskRevokeSourceReview{}
	return &this
}

// GetAccessReviewId returns the AccessReviewId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskRevokeSourceReview) GetAccessReviewId() string {
	if o == nil || IsNil(o.AccessReviewId) {
		var ret string
		return ret
	}
	return *o.AccessReviewId
}

// GetAccessReviewIdOk returns a tuple with the AccessReviewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskRevokeSourceReview) GetAccessReviewIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessReviewId) {
		return nil, false
	}
	return o.AccessReviewId, true
}

// HasAccessReviewId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskRevokeSourceReview) HasAccessReviewId() bool {
	if o != nil && !IsNil(o.AccessReviewId) {
		return true
	}

	return false
}

// SetAccessReviewId gets a reference to the given string and assigns it to the AccessReviewId field.
func (o *C1ApiTaskV1TaskRevokeSourceReview) SetAccessReviewId(v string) {
	o.AccessReviewId = &v
}

// GetCertTicketId returns the CertTicketId field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskRevokeSourceReview) GetCertTicketId() string {
	if o == nil || IsNil(o.CertTicketId) {
		var ret string
		return ret
	}
	return *o.CertTicketId
}

// GetCertTicketIdOk returns a tuple with the CertTicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskRevokeSourceReview) GetCertTicketIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertTicketId) {
		return nil, false
	}
	return o.CertTicketId, true
}

// HasCertTicketId returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskRevokeSourceReview) HasCertTicketId() bool {
	if o != nil && !IsNil(o.CertTicketId) {
		return true
	}

	return false
}

// SetCertTicketId gets a reference to the given string and assigns it to the CertTicketId field.
func (o *C1ApiTaskV1TaskRevokeSourceReview) SetCertTicketId(v string) {
	o.CertTicketId = &v
}

func (o C1ApiTaskV1TaskRevokeSourceReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1TaskRevokeSourceReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessReviewId) {
		toSerialize["accessReviewId"] = o.AccessReviewId
	}
	if !IsNil(o.CertTicketId) {
		toSerialize["certTicketId"] = o.CertTicketId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1TaskRevokeSourceReview) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1TaskRevokeSourceReview := _C1ApiTaskV1TaskRevokeSourceReview{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1TaskRevokeSourceReview); err == nil {
		*o = C1ApiTaskV1TaskRevokeSourceReview(varC1ApiTaskV1TaskRevokeSourceReview)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessReviewId")
		delete(additionalProperties, "certTicketId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1TaskRevokeSourceReview struct {
	value *C1ApiTaskV1TaskRevokeSourceReview
	isSet bool
}

func (v NullableC1ApiTaskV1TaskRevokeSourceReview) Get() *C1ApiTaskV1TaskRevokeSourceReview {
	return v.value
}

func (v *NullableC1ApiTaskV1TaskRevokeSourceReview) Set(val *C1ApiTaskV1TaskRevokeSourceReview) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1TaskRevokeSourceReview) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1TaskRevokeSourceReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1TaskRevokeSourceReview(val *C1ApiTaskV1TaskRevokeSourceReview) *NullableC1ApiTaskV1TaskRevokeSourceReview {
	return &NullableC1ApiTaskV1TaskRevokeSourceReview{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1TaskRevokeSourceReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1TaskRevokeSourceReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


