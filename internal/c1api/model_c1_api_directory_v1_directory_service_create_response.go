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

// checks if the C1ApiDirectoryV1DirectoryServiceCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiDirectoryV1DirectoryServiceCreateResponse{}

// C1ApiDirectoryV1DirectoryServiceCreateResponse The DirectoryServiceCreateResponse message.
type C1ApiDirectoryV1DirectoryServiceCreateResponse struct {
	DirectoryView *C1ApiDirectoryV1DirectoryView `json:"directoryView,omitempty"`
	// The expanded field.
	Expanded []C1ApiAppV1AppResourceServiceGetResponseExpandedInner `json:"expanded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiDirectoryV1DirectoryServiceCreateResponse C1ApiDirectoryV1DirectoryServiceCreateResponse

// NewC1ApiDirectoryV1DirectoryServiceCreateResponse instantiates a new C1ApiDirectoryV1DirectoryServiceCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiDirectoryV1DirectoryServiceCreateResponse() *C1ApiDirectoryV1DirectoryServiceCreateResponse {
	this := C1ApiDirectoryV1DirectoryServiceCreateResponse{}
	return &this
}

// NewC1ApiDirectoryV1DirectoryServiceCreateResponseWithDefaults instantiates a new C1ApiDirectoryV1DirectoryServiceCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiDirectoryV1DirectoryServiceCreateResponseWithDefaults() *C1ApiDirectoryV1DirectoryServiceCreateResponse {
	this := C1ApiDirectoryV1DirectoryServiceCreateResponse{}
	return &this
}

// GetDirectoryView returns the DirectoryView field value if set, zero value otherwise.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) GetDirectoryView() C1ApiDirectoryV1DirectoryView {
	if o == nil || IsNil(o.DirectoryView) {
		var ret C1ApiDirectoryV1DirectoryView
		return ret
	}
	return *o.DirectoryView
}

// GetDirectoryViewOk returns a tuple with the DirectoryView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) GetDirectoryViewOk() (*C1ApiDirectoryV1DirectoryView, bool) {
	if o == nil || IsNil(o.DirectoryView) {
		return nil, false
	}
	return o.DirectoryView, true
}

// HasDirectoryView returns a boolean if a field has been set.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) HasDirectoryView() bool {
	if o != nil && !IsNil(o.DirectoryView) {
		return true
	}

	return false
}

// SetDirectoryView gets a reference to the given C1ApiDirectoryV1DirectoryView and assigns it to the DirectoryView field.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) SetDirectoryView(v C1ApiDirectoryV1DirectoryView) {
	o.DirectoryView = &v
}

// GetExpanded returns the Expanded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner {
	if o == nil {
		var ret []C1ApiAppV1AppResourceServiceGetResponseExpandedInner
		return ret
	}
	return o.Expanded
}

// GetExpandedOk returns a tuple with the Expanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) GetExpandedOk() ([]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool) {
	if o == nil || IsNil(o.Expanded) {
		return nil, false
	}
	return o.Expanded, true
}

// HasExpanded returns a boolean if a field has been set.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) HasExpanded() bool {
	if o != nil && IsNil(o.Expanded) {
		return true
	}

	return false
}

// SetExpanded gets a reference to the given []C1ApiAppV1AppResourceServiceGetResponseExpandedInner and assigns it to the Expanded field.
func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner) {
	o.Expanded = v
}

func (o C1ApiDirectoryV1DirectoryServiceCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiDirectoryV1DirectoryServiceCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectoryView) {
		toSerialize["directoryView"] = o.DirectoryView
	}
	if o.Expanded != nil {
		toSerialize["expanded"] = o.Expanded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiDirectoryV1DirectoryServiceCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiDirectoryV1DirectoryServiceCreateResponse := _C1ApiDirectoryV1DirectoryServiceCreateResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiDirectoryV1DirectoryServiceCreateResponse); err == nil {
		*o = C1ApiDirectoryV1DirectoryServiceCreateResponse(varC1ApiDirectoryV1DirectoryServiceCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "directoryView")
		delete(additionalProperties, "expanded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiDirectoryV1DirectoryServiceCreateResponse struct {
	value *C1ApiDirectoryV1DirectoryServiceCreateResponse
	isSet bool
}

func (v NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) Get() *C1ApiDirectoryV1DirectoryServiceCreateResponse {
	return v.value
}

func (v *NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) Set(val *C1ApiDirectoryV1DirectoryServiceCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiDirectoryV1DirectoryServiceCreateResponse(val *C1ApiDirectoryV1DirectoryServiceCreateResponse) *NullableC1ApiDirectoryV1DirectoryServiceCreateResponse {
	return &NullableC1ApiDirectoryV1DirectoryServiceCreateResponse{value: val, isSet: true}
}

func (v NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiDirectoryV1DirectoryServiceCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


