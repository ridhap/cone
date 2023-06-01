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

// checks if the C1ApiPolicyV1EntitlementOwnerApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiPolicyV1EntitlementOwnerApproval{}

// C1ApiPolicyV1EntitlementOwnerApproval The EntitlementOwnerApproval message.
type C1ApiPolicyV1EntitlementOwnerApproval struct {
	//  Entitlement owner is based on the current entitlement's id and doesn't need to have self-contained data 
	AllowSelfApproval *bool `json:"allowSelfApproval,omitempty"`
	// The fallback field.
	Fallback *bool `json:"fallback,omitempty"`
	// The fallbackUserIds field.
	FallbackUserIds []string `json:"fallbackUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiPolicyV1EntitlementOwnerApproval C1ApiPolicyV1EntitlementOwnerApproval

// NewC1ApiPolicyV1EntitlementOwnerApproval instantiates a new C1ApiPolicyV1EntitlementOwnerApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiPolicyV1EntitlementOwnerApproval() *C1ApiPolicyV1EntitlementOwnerApproval {
	this := C1ApiPolicyV1EntitlementOwnerApproval{}
	return &this
}

// NewC1ApiPolicyV1EntitlementOwnerApprovalWithDefaults instantiates a new C1ApiPolicyV1EntitlementOwnerApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiPolicyV1EntitlementOwnerApprovalWithDefaults() *C1ApiPolicyV1EntitlementOwnerApproval {
	this := C1ApiPolicyV1EntitlementOwnerApproval{}
	return &this
}

// GetAllowSelfApproval returns the AllowSelfApproval field value if set, zero value otherwise.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetAllowSelfApproval() bool {
	if o == nil || IsNil(o.AllowSelfApproval) {
		var ret bool
		return ret
	}
	return *o.AllowSelfApproval
}

// GetAllowSelfApprovalOk returns a tuple with the AllowSelfApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetAllowSelfApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSelfApproval) {
		return nil, false
	}
	return o.AllowSelfApproval, true
}

// HasAllowSelfApproval returns a boolean if a field has been set.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasAllowSelfApproval() bool {
	if o != nil && !IsNil(o.AllowSelfApproval) {
		return true
	}

	return false
}

// SetAllowSelfApproval gets a reference to the given bool and assigns it to the AllowSelfApproval field.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetAllowSelfApproval(v bool) {
	o.AllowSelfApproval = &v
}

// GetFallback returns the Fallback field value if set, zero value otherwise.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallback() bool {
	if o == nil || IsNil(o.Fallback) {
		var ret bool
		return ret
	}
	return *o.Fallback
}

// GetFallbackOk returns a tuple with the Fallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackOk() (*bool, bool) {
	if o == nil || IsNil(o.Fallback) {
		return nil, false
	}
	return o.Fallback, true
}

// HasFallback returns a boolean if a field has been set.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasFallback() bool {
	if o != nil && !IsNil(o.Fallback) {
		return true
	}

	return false
}

// SetFallback gets a reference to the given bool and assigns it to the Fallback field.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetFallback(v bool) {
	o.Fallback = &v
}

// GetFallbackUserIds returns the FallbackUserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FallbackUserIds
}

// GetFallbackUserIdsOk returns a tuple with the FallbackUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FallbackUserIds) {
		return nil, false
	}
	return o.FallbackUserIds, true
}

// HasFallbackUserIds returns a boolean if a field has been set.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasFallbackUserIds() bool {
	if o != nil && IsNil(o.FallbackUserIds) {
		return true
	}

	return false
}

// SetFallbackUserIds gets a reference to the given []string and assigns it to the FallbackUserIds field.
func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetFallbackUserIds(v []string) {
	o.FallbackUserIds = v
}

func (o C1ApiPolicyV1EntitlementOwnerApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiPolicyV1EntitlementOwnerApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowSelfApproval) {
		toSerialize["allowSelfApproval"] = o.AllowSelfApproval
	}
	if !IsNil(o.Fallback) {
		toSerialize["fallback"] = o.Fallback
	}
	if o.FallbackUserIds != nil {
		toSerialize["fallbackUserIds"] = o.FallbackUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiPolicyV1EntitlementOwnerApproval) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiPolicyV1EntitlementOwnerApproval := _C1ApiPolicyV1EntitlementOwnerApproval{}

	if err = json.Unmarshal(bytes, &varC1ApiPolicyV1EntitlementOwnerApproval); err == nil {
		*o = C1ApiPolicyV1EntitlementOwnerApproval(varC1ApiPolicyV1EntitlementOwnerApproval)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowSelfApproval")
		delete(additionalProperties, "fallback")
		delete(additionalProperties, "fallbackUserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiPolicyV1EntitlementOwnerApproval struct {
	value *C1ApiPolicyV1EntitlementOwnerApproval
	isSet bool
}

func (v NullableC1ApiPolicyV1EntitlementOwnerApproval) Get() *C1ApiPolicyV1EntitlementOwnerApproval {
	return v.value
}

func (v *NullableC1ApiPolicyV1EntitlementOwnerApproval) Set(val *C1ApiPolicyV1EntitlementOwnerApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiPolicyV1EntitlementOwnerApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiPolicyV1EntitlementOwnerApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiPolicyV1EntitlementOwnerApproval(val *C1ApiPolicyV1EntitlementOwnerApproval) *NullableC1ApiPolicyV1EntitlementOwnerApproval {
	return &NullableC1ApiPolicyV1EntitlementOwnerApproval{value: val, isSet: true}
}

func (v NullableC1ApiPolicyV1EntitlementOwnerApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiPolicyV1EntitlementOwnerApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

