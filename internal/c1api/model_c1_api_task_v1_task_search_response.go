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

// checks if the C1ApiTaskV1TaskSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1TaskSearchResponse{}

// C1ApiTaskV1TaskSearchResponse The TaskSearchResponse message.
type C1ApiTaskV1TaskSearchResponse struct {
	// The expanded field.
	Expanded []C1ApiAppV1AppResourceServiceGetResponseExpandedInner `json:"expanded,omitempty"`
	// The list field.
	List []C1ApiTaskV1TaskView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1TaskSearchResponse C1ApiTaskV1TaskSearchResponse

// NewC1ApiTaskV1TaskSearchResponse instantiates a new C1ApiTaskV1TaskSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1TaskSearchResponse() *C1ApiTaskV1TaskSearchResponse {
	this := C1ApiTaskV1TaskSearchResponse{}
	return &this
}

// NewC1ApiTaskV1TaskSearchResponseWithDefaults instantiates a new C1ApiTaskV1TaskSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1TaskSearchResponseWithDefaults() *C1ApiTaskV1TaskSearchResponse {
	this := C1ApiTaskV1TaskSearchResponse{}
	return &this
}

// GetExpanded returns the Expanded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiTaskV1TaskSearchResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner {
	if o == nil {
		var ret []C1ApiAppV1AppResourceServiceGetResponseExpandedInner
		return ret
	}
	return o.Expanded
}

// GetExpandedOk returns a tuple with the Expanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiTaskV1TaskSearchResponse) GetExpandedOk() ([]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool) {
	if o == nil || IsNil(o.Expanded) {
		return nil, false
	}
	return o.Expanded, true
}

// HasExpanded returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskSearchResponse) HasExpanded() bool {
	if o != nil && IsNil(o.Expanded) {
		return true
	}

	return false
}

// SetExpanded gets a reference to the given []C1ApiAppV1AppResourceServiceGetResponseExpandedInner and assigns it to the Expanded field.
func (o *C1ApiTaskV1TaskSearchResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner) {
	o.Expanded = v
}

// GetList returns the List field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiTaskV1TaskSearchResponse) GetList() []C1ApiTaskV1TaskView {
	if o == nil {
		var ret []C1ApiTaskV1TaskView
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiTaskV1TaskSearchResponse) GetListOk() ([]C1ApiTaskV1TaskView, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskSearchResponse) HasList() bool {
	if o != nil && IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []C1ApiTaskV1TaskView and assigns it to the List field.
func (o *C1ApiTaskV1TaskSearchResponse) SetList(v []C1ApiTaskV1TaskView) {
	o.List = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *C1ApiTaskV1TaskSearchResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1TaskSearchResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *C1ApiTaskV1TaskSearchResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *C1ApiTaskV1TaskSearchResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o C1ApiTaskV1TaskSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1TaskSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Expanded != nil {
		toSerialize["expanded"] = o.Expanded
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1TaskSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1TaskSearchResponse := _C1ApiTaskV1TaskSearchResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1TaskSearchResponse); err == nil {
		*o = C1ApiTaskV1TaskSearchResponse(varC1ApiTaskV1TaskSearchResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expanded")
		delete(additionalProperties, "list")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1TaskSearchResponse struct {
	value *C1ApiTaskV1TaskSearchResponse
	isSet bool
}

func (v NullableC1ApiTaskV1TaskSearchResponse) Get() *C1ApiTaskV1TaskSearchResponse {
	return v.value
}

func (v *NullableC1ApiTaskV1TaskSearchResponse) Set(val *C1ApiTaskV1TaskSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1TaskSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1TaskSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1TaskSearchResponse(val *C1ApiTaskV1TaskSearchResponse) *NullableC1ApiTaskV1TaskSearchResponse {
	return &NullableC1ApiTaskV1TaskSearchResponse{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1TaskSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1TaskSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

