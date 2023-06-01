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

// checks if the C1ApiTaskV1GrantEntitlementInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiTaskV1GrantEntitlementInstance{}

// C1ApiTaskV1GrantEntitlementInstance The GrantEntitlementInstance message.
type C1ApiTaskV1GrantEntitlementInstance struct {
	// The entitlementId field.
	EntitlementId *string `json:"entitlementId,omitempty"`
	GrantDuration *string `json:"grantDuration,omitempty"`
	// The state field.
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiTaskV1GrantEntitlementInstance C1ApiTaskV1GrantEntitlementInstance

// NewC1ApiTaskV1GrantEntitlementInstance instantiates a new C1ApiTaskV1GrantEntitlementInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiTaskV1GrantEntitlementInstance() *C1ApiTaskV1GrantEntitlementInstance {
	this := C1ApiTaskV1GrantEntitlementInstance{}
	return &this
}

// NewC1ApiTaskV1GrantEntitlementInstanceWithDefaults instantiates a new C1ApiTaskV1GrantEntitlementInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiTaskV1GrantEntitlementInstanceWithDefaults() *C1ApiTaskV1GrantEntitlementInstance {
	this := C1ApiTaskV1GrantEntitlementInstance{}
	return &this
}

// GetEntitlementId returns the EntitlementId field value if set, zero value otherwise.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetEntitlementId() string {
	if o == nil || IsNil(o.EntitlementId) {
		var ret string
		return ret
	}
	return *o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetEntitlementIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementId) {
		return nil, false
	}
	return o.EntitlementId, true
}

// HasEntitlementId returns a boolean if a field has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) HasEntitlementId() bool {
	if o != nil && !IsNil(o.EntitlementId) {
		return true
	}

	return false
}

// SetEntitlementId gets a reference to the given string and assigns it to the EntitlementId field.
func (o *C1ApiTaskV1GrantEntitlementInstance) SetEntitlementId(v string) {
	o.EntitlementId = &v
}

// GetGrantDuration returns the GrantDuration field value if set, zero value otherwise.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetGrantDuration() string {
	if o == nil || IsNil(o.GrantDuration) {
		var ret string
		return ret
	}
	return *o.GrantDuration
}

// GetGrantDurationOk returns a tuple with the GrantDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetGrantDurationOk() (*string, bool) {
	if o == nil || IsNil(o.GrantDuration) {
		return nil, false
	}
	return o.GrantDuration, true
}

// HasGrantDuration returns a boolean if a field has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) HasGrantDuration() bool {
	if o != nil && !IsNil(o.GrantDuration) {
		return true
	}

	return false
}

// SetGrantDuration gets a reference to the given string and assigns it to the GrantDuration field.
func (o *C1ApiTaskV1GrantEntitlementInstance) SetGrantDuration(v string) {
	o.GrantDuration = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *C1ApiTaskV1GrantEntitlementInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *C1ApiTaskV1GrantEntitlementInstance) SetState(v string) {
	o.State = &v
}

func (o C1ApiTaskV1GrantEntitlementInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiTaskV1GrantEntitlementInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntitlementId) {
		toSerialize["entitlementId"] = o.EntitlementId
	}
	if !IsNil(o.GrantDuration) {
		toSerialize["grantDuration"] = o.GrantDuration
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiTaskV1GrantEntitlementInstance) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiTaskV1GrantEntitlementInstance := _C1ApiTaskV1GrantEntitlementInstance{}

	if err = json.Unmarshal(bytes, &varC1ApiTaskV1GrantEntitlementInstance); err == nil {
		*o = C1ApiTaskV1GrantEntitlementInstance(varC1ApiTaskV1GrantEntitlementInstance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "grantDuration")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiTaskV1GrantEntitlementInstance struct {
	value *C1ApiTaskV1GrantEntitlementInstance
	isSet bool
}

func (v NullableC1ApiTaskV1GrantEntitlementInstance) Get() *C1ApiTaskV1GrantEntitlementInstance {
	return v.value
}

func (v *NullableC1ApiTaskV1GrantEntitlementInstance) Set(val *C1ApiTaskV1GrantEntitlementInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiTaskV1GrantEntitlementInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiTaskV1GrantEntitlementInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiTaskV1GrantEntitlementInstance(val *C1ApiTaskV1GrantEntitlementInstance) *NullableC1ApiTaskV1GrantEntitlementInstance {
	return &NullableC1ApiTaskV1GrantEntitlementInstance{value: val, isSet: true}
}

func (v NullableC1ApiTaskV1GrantEntitlementInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiTaskV1GrantEntitlementInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

