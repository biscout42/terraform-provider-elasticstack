/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
)

// checks if the FieldmapPropertiesPropertiesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldmapPropertiesPropertiesValue{}

// FieldmapPropertiesPropertiesValue struct for FieldmapPropertiesPropertiesValue
type FieldmapPropertiesPropertiesValue struct {
	// The data type for each object property.
	Type *string `json:"type,omitempty"`
}

// NewFieldmapPropertiesPropertiesValue instantiates a new FieldmapPropertiesPropertiesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldmapPropertiesPropertiesValue() *FieldmapPropertiesPropertiesValue {
	this := FieldmapPropertiesPropertiesValue{}
	return &this
}

// NewFieldmapPropertiesPropertiesValueWithDefaults instantiates a new FieldmapPropertiesPropertiesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldmapPropertiesPropertiesValueWithDefaults() *FieldmapPropertiesPropertiesValue {
	this := FieldmapPropertiesPropertiesValue{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldmapPropertiesPropertiesValue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldmapPropertiesPropertiesValue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldmapPropertiesPropertiesValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldmapPropertiesPropertiesValue) SetType(v string) {
	o.Type = &v
}

func (o FieldmapPropertiesPropertiesValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldmapPropertiesPropertiesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFieldmapPropertiesPropertiesValue struct {
	value *FieldmapPropertiesPropertiesValue
	isSet bool
}

func (v NullableFieldmapPropertiesPropertiesValue) Get() *FieldmapPropertiesPropertiesValue {
	return v.value
}

func (v *NullableFieldmapPropertiesPropertiesValue) Set(val *FieldmapPropertiesPropertiesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldmapPropertiesPropertiesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldmapPropertiesPropertiesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldmapPropertiesPropertiesValue(val *FieldmapPropertiesPropertiesValue) *NullableFieldmapPropertiesPropertiesValue {
	return &NullableFieldmapPropertiesPropertiesValue{value: val, isSet: true}
}

func (v NullableFieldmapPropertiesPropertiesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldmapPropertiesPropertiesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
