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

// checks if the C1ApiAppV1GetAppResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1GetAppResponse{}

// C1ApiAppV1GetAppResponse The GetAppResponse message.
type C1ApiAppV1GetAppResponse struct {
	App *C1ApiAppV1App `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1GetAppResponse C1ApiAppV1GetAppResponse

// NewC1ApiAppV1GetAppResponse instantiates a new C1ApiAppV1GetAppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1GetAppResponse() *C1ApiAppV1GetAppResponse {
	this := C1ApiAppV1GetAppResponse{}
	return &this
}

// NewC1ApiAppV1GetAppResponseWithDefaults instantiates a new C1ApiAppV1GetAppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1GetAppResponseWithDefaults() *C1ApiAppV1GetAppResponse {
	this := C1ApiAppV1GetAppResponse{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *C1ApiAppV1GetAppResponse) GetApp() C1ApiAppV1App {
	if o == nil || IsNil(o.App) {
		var ret C1ApiAppV1App
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1GetAppResponse) GetAppOk() (*C1ApiAppV1App, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *C1ApiAppV1GetAppResponse) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given C1ApiAppV1App and assigns it to the App field.
func (o *C1ApiAppV1GetAppResponse) SetApp(v C1ApiAppV1App) {
	o.App = &v
}

func (o C1ApiAppV1GetAppResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1GetAppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiAppV1GetAppResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1GetAppResponse := _C1ApiAppV1GetAppResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1GetAppResponse); err == nil {
		*o = C1ApiAppV1GetAppResponse(varC1ApiAppV1GetAppResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1GetAppResponse struct {
	value *C1ApiAppV1GetAppResponse
	isSet bool
}

func (v NullableC1ApiAppV1GetAppResponse) Get() *C1ApiAppV1GetAppResponse {
	return v.value
}

func (v *NullableC1ApiAppV1GetAppResponse) Set(val *C1ApiAppV1GetAppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1GetAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1GetAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1GetAppResponse(val *C1ApiAppV1GetAppResponse) *NullableC1ApiAppV1GetAppResponse {
	return &NullableC1ApiAppV1GetAppResponse{value: val, isSet: true}
}

func (v NullableC1ApiAppV1GetAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1GetAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

