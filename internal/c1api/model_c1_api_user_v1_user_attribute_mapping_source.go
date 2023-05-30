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

// checks if the C1ApiUserV1UserAttributeMappingSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiUserV1UserAttributeMappingSource{}

// C1ApiUserV1UserAttributeMappingSource The UserAttributeMappingSource message.
type C1ApiUserV1UserAttributeMappingSource struct {
	AppId interface{} `json:"appId,omitempty"`
	AppUserId interface{} `json:"appUserId,omitempty"`
	AppUserProfileAttributeKey interface{} `json:"appUserProfileAttributeKey,omitempty"`
	UserAttributeMappingId interface{} `json:"userAttributeMappingId,omitempty"`
	Value interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiUserV1UserAttributeMappingSource C1ApiUserV1UserAttributeMappingSource

// NewC1ApiUserV1UserAttributeMappingSource instantiates a new C1ApiUserV1UserAttributeMappingSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiUserV1UserAttributeMappingSource() *C1ApiUserV1UserAttributeMappingSource {
	this := C1ApiUserV1UserAttributeMappingSource{}
	return &this
}

// NewC1ApiUserV1UserAttributeMappingSourceWithDefaults instantiates a new C1ApiUserV1UserAttributeMappingSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiUserV1UserAttributeMappingSourceWithDefaults() *C1ApiUserV1UserAttributeMappingSource {
	this := C1ApiUserV1UserAttributeMappingSource{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return &o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *C1ApiUserV1UserAttributeMappingSource) HasAppId() bool {
	if o != nil && IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given interface{} and assigns it to the AppId field.
func (o *C1ApiUserV1UserAttributeMappingSource) SetAppId(v interface{}) {
	o.AppId = v
}

// GetAppUserId returns the AppUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppUserId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AppUserId
}

// GetAppUserIdOk returns a tuple with the AppUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppUserIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AppUserId) {
		return nil, false
	}
	return &o.AppUserId, true
}

// HasAppUserId returns a boolean if a field has been set.
func (o *C1ApiUserV1UserAttributeMappingSource) HasAppUserId() bool {
	if o != nil && IsNil(o.AppUserId) {
		return true
	}

	return false
}

// SetAppUserId gets a reference to the given interface{} and assigns it to the AppUserId field.
func (o *C1ApiUserV1UserAttributeMappingSource) SetAppUserId(v interface{}) {
	o.AppUserId = v
}

// GetAppUserProfileAttributeKey returns the AppUserProfileAttributeKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppUserProfileAttributeKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AppUserProfileAttributeKey
}

// GetAppUserProfileAttributeKeyOk returns a tuple with the AppUserProfileAttributeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserAttributeMappingSource) GetAppUserProfileAttributeKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AppUserProfileAttributeKey) {
		return nil, false
	}
	return &o.AppUserProfileAttributeKey, true
}

// HasAppUserProfileAttributeKey returns a boolean if a field has been set.
func (o *C1ApiUserV1UserAttributeMappingSource) HasAppUserProfileAttributeKey() bool {
	if o != nil && IsNil(o.AppUserProfileAttributeKey) {
		return true
	}

	return false
}

// SetAppUserProfileAttributeKey gets a reference to the given interface{} and assigns it to the AppUserProfileAttributeKey field.
func (o *C1ApiUserV1UserAttributeMappingSource) SetAppUserProfileAttributeKey(v interface{}) {
	o.AppUserProfileAttributeKey = v
}

// GetUserAttributeMappingId returns the UserAttributeMappingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserAttributeMappingSource) GetUserAttributeMappingId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UserAttributeMappingId
}

// GetUserAttributeMappingIdOk returns a tuple with the UserAttributeMappingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserAttributeMappingSource) GetUserAttributeMappingIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UserAttributeMappingId) {
		return nil, false
	}
	return &o.UserAttributeMappingId, true
}

// HasUserAttributeMappingId returns a boolean if a field has been set.
func (o *C1ApiUserV1UserAttributeMappingSource) HasUserAttributeMappingId() bool {
	if o != nil && IsNil(o.UserAttributeMappingId) {
		return true
	}

	return false
}

// SetUserAttributeMappingId gets a reference to the given interface{} and assigns it to the UserAttributeMappingId field.
func (o *C1ApiUserV1UserAttributeMappingSource) SetUserAttributeMappingId(v interface{}) {
	o.UserAttributeMappingId = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserAttributeMappingSource) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserAttributeMappingSource) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *C1ApiUserV1UserAttributeMappingSource) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *C1ApiUserV1UserAttributeMappingSource) SetValue(v interface{}) {
	o.Value = v
}

func (o C1ApiUserV1UserAttributeMappingSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiUserV1UserAttributeMappingSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUserId != nil {
		toSerialize["appUserId"] = o.AppUserId
	}
	if o.AppUserProfileAttributeKey != nil {
		toSerialize["appUserProfileAttributeKey"] = o.AppUserProfileAttributeKey
	}
	if o.UserAttributeMappingId != nil {
		toSerialize["userAttributeMappingId"] = o.UserAttributeMappingId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiUserV1UserAttributeMappingSource) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiUserV1UserAttributeMappingSource := _C1ApiUserV1UserAttributeMappingSource{}

	if err = json.Unmarshal(bytes, &varC1ApiUserV1UserAttributeMappingSource); err == nil {
		*o = C1ApiUserV1UserAttributeMappingSource(varC1ApiUserV1UserAttributeMappingSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appId")
		delete(additionalProperties, "appUserId")
		delete(additionalProperties, "appUserProfileAttributeKey")
		delete(additionalProperties, "userAttributeMappingId")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiUserV1UserAttributeMappingSource struct {
	value *C1ApiUserV1UserAttributeMappingSource
	isSet bool
}

func (v NullableC1ApiUserV1UserAttributeMappingSource) Get() *C1ApiUserV1UserAttributeMappingSource {
	return v.value
}

func (v *NullableC1ApiUserV1UserAttributeMappingSource) Set(val *C1ApiUserV1UserAttributeMappingSource) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiUserV1UserAttributeMappingSource) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiUserV1UserAttributeMappingSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiUserV1UserAttributeMappingSource(val *C1ApiUserV1UserAttributeMappingSource) *NullableC1ApiUserV1UserAttributeMappingSource {
	return &NullableC1ApiUserV1UserAttributeMappingSource{value: val, isSet: true}
}

func (v NullableC1ApiUserV1UserAttributeMappingSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiUserV1UserAttributeMappingSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

