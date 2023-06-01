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

// checks if the C1ApiAppV1AppResourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1AppResourceType{}

// C1ApiAppV1AppResourceType The AppResourceType message.
type C1ApiAppV1AppResourceType struct {
	// The appId field.
	AppId *string `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The id field.
	Id *string `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1AppResourceType C1ApiAppV1AppResourceType

// NewC1ApiAppV1AppResourceType instantiates a new C1ApiAppV1AppResourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1AppResourceType() *C1ApiAppV1AppResourceType {
	this := C1ApiAppV1AppResourceType{}
	return &this
}

// NewC1ApiAppV1AppResourceTypeWithDefaults instantiates a new C1ApiAppV1AppResourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1AppResourceTypeWithDefaults() *C1ApiAppV1AppResourceType {
	this := C1ApiAppV1AppResourceType{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *C1ApiAppV1AppResourceType) SetAppId(v string) {
	o.AppId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *C1ApiAppV1AppResourceType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *C1ApiAppV1AppResourceType) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *C1ApiAppV1AppResourceType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *C1ApiAppV1AppResourceType) SetId(v string) {
	o.Id = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceType) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceType) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceType) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *C1ApiAppV1AppResourceType) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o C1ApiAppV1AppResourceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1AppResourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiAppV1AppResourceType) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1AppResourceType := _C1ApiAppV1AppResourceType{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1AppResourceType); err == nil {
		*o = C1ApiAppV1AppResourceType(varC1ApiAppV1AppResourceType)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appId")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "deletedAt")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1AppResourceType struct {
	value *C1ApiAppV1AppResourceType
	isSet bool
}

func (v NullableC1ApiAppV1AppResourceType) Get() *C1ApiAppV1AppResourceType {
	return v.value
}

func (v *NullableC1ApiAppV1AppResourceType) Set(val *C1ApiAppV1AppResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1AppResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1AppResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1AppResourceType(val *C1ApiAppV1AppResourceType) *NullableC1ApiAppV1AppResourceType {
	return &NullableC1ApiAppV1AppResourceType{value: val, isSet: true}
}

func (v NullableC1ApiAppV1AppResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1AppResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

