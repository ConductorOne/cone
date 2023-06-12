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

// checks if the C1ApiAppV1AppReportServiceListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1AppReportServiceListResponse{}

// C1ApiAppV1AppReportServiceListResponse The AppReportServiceListResponse message.
type C1ApiAppV1AppReportServiceListResponse struct {
	// The list field.
	List []C1ApiAppV1AppPopulationReport `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1AppReportServiceListResponse C1ApiAppV1AppReportServiceListResponse

// NewC1ApiAppV1AppReportServiceListResponse instantiates a new C1ApiAppV1AppReportServiceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1AppReportServiceListResponse() *C1ApiAppV1AppReportServiceListResponse {
	this := C1ApiAppV1AppReportServiceListResponse{}
	return &this
}

// NewC1ApiAppV1AppReportServiceListResponseWithDefaults instantiates a new C1ApiAppV1AppReportServiceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1AppReportServiceListResponseWithDefaults() *C1ApiAppV1AppReportServiceListResponse {
	this := C1ApiAppV1AppReportServiceListResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiAppV1AppReportServiceListResponse) GetList() []C1ApiAppV1AppPopulationReport {
	if o == nil {
		var ret []C1ApiAppV1AppPopulationReport
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiAppV1AppReportServiceListResponse) GetListOk() ([]C1ApiAppV1AppPopulationReport, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *C1ApiAppV1AppReportServiceListResponse) HasList() bool {
	if o != nil && IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []C1ApiAppV1AppPopulationReport and assigns it to the List field.
func (o *C1ApiAppV1AppReportServiceListResponse) SetList(v []C1ApiAppV1AppPopulationReport) {
	o.List = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *C1ApiAppV1AppReportServiceListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppReportServiceListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *C1ApiAppV1AppReportServiceListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *C1ApiAppV1AppReportServiceListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o C1ApiAppV1AppReportServiceListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1AppReportServiceListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *C1ApiAppV1AppReportServiceListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1AppReportServiceListResponse := _C1ApiAppV1AppReportServiceListResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1AppReportServiceListResponse); err == nil {
		*o = C1ApiAppV1AppReportServiceListResponse(varC1ApiAppV1AppReportServiceListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1AppReportServiceListResponse struct {
	value *C1ApiAppV1AppReportServiceListResponse
	isSet bool
}

func (v NullableC1ApiAppV1AppReportServiceListResponse) Get() *C1ApiAppV1AppReportServiceListResponse {
	return v.value
}

func (v *NullableC1ApiAppV1AppReportServiceListResponse) Set(val *C1ApiAppV1AppReportServiceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1AppReportServiceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1AppReportServiceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1AppReportServiceListResponse(val *C1ApiAppV1AppReportServiceListResponse) *NullableC1ApiAppV1AppReportServiceListResponse {
	return &NullableC1ApiAppV1AppReportServiceListResponse{value: val, isSet: true}
}

func (v NullableC1ApiAppV1AppReportServiceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1AppReportServiceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


