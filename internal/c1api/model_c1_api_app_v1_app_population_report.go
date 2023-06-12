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

// checks if the C1ApiAppV1AppPopulationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1AppPopulationReport{}

// C1ApiAppV1AppPopulationReport The AppPopulationReport message.
type C1ApiAppV1AppPopulationReport struct {
	// The appId field.
	AppId *string `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The downloadUrl field.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// The hashes field.
	Hashes *map[string]string `json:"hashes,omitempty"`
	// The id field.
	Id *string `json:"id,omitempty"`
	// The state field.
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1AppPopulationReport C1ApiAppV1AppPopulationReport

// NewC1ApiAppV1AppPopulationReport instantiates a new C1ApiAppV1AppPopulationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1AppPopulationReport() *C1ApiAppV1AppPopulationReport {
	this := C1ApiAppV1AppPopulationReport{}
	return &this
}

// NewC1ApiAppV1AppPopulationReportWithDefaults instantiates a new C1ApiAppV1AppPopulationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1AppPopulationReportWithDefaults() *C1ApiAppV1AppPopulationReport {
	this := C1ApiAppV1AppPopulationReport{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *C1ApiAppV1AppPopulationReport) SetAppId(v string) {
	o.AppId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *C1ApiAppV1AppPopulationReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *C1ApiAppV1AppPopulationReport) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetHashes returns the Hashes field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetHashes() map[string]string {
	if o == nil || IsNil(o.Hashes) {
		var ret map[string]string
		return ret
	}
	return *o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetHashesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Hashes) {
		return nil, false
	}
	return o.Hashes, true
}

// HasHashes returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasHashes() bool {
	if o != nil && !IsNil(o.Hashes) {
		return true
	}

	return false
}

// SetHashes gets a reference to the given map[string]string and assigns it to the Hashes field.
func (o *C1ApiAppV1AppPopulationReport) SetHashes(v map[string]string) {
	o.Hashes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *C1ApiAppV1AppPopulationReport) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *C1ApiAppV1AppPopulationReport) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppPopulationReport) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *C1ApiAppV1AppPopulationReport) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *C1ApiAppV1AppPopulationReport) SetState(v string) {
	o.State = &v
}

func (o C1ApiAppV1AppPopulationReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1AppPopulationReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.Hashes) {
		toSerialize["hashes"] = o.Hashes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiAppV1AppPopulationReport) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1AppPopulationReport := _C1ApiAppV1AppPopulationReport{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1AppPopulationReport); err == nil {
		*o = C1ApiAppV1AppPopulationReport(varC1ApiAppV1AppPopulationReport)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appId")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "downloadUrl")
		delete(additionalProperties, "hashes")
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1AppPopulationReport struct {
	value *C1ApiAppV1AppPopulationReport
	isSet bool
}

func (v NullableC1ApiAppV1AppPopulationReport) Get() *C1ApiAppV1AppPopulationReport {
	return v.value
}

func (v *NullableC1ApiAppV1AppPopulationReport) Set(val *C1ApiAppV1AppPopulationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1AppPopulationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1AppPopulationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1AppPopulationReport(val *C1ApiAppV1AppPopulationReport) *NullableC1ApiAppV1AppPopulationReport {
	return &NullableC1ApiAppV1AppPopulationReport{value: val, isSet: true}
}

func (v NullableC1ApiAppV1AppPopulationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1AppPopulationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


