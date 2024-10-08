/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
	"time"
)

// checks if the LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}

// LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth The timestamp and status of the alert decryption.
type LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth struct {
	Status    *string    `json:"status,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth instantiates a new LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth() *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	this := LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}
	return &this
}

// NewLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealthWithDefaults instantiates a new LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealthWithDefaults() *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	this := LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth struct {
	value *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth
	isSet bool
}

func (v NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Get() *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	return v.value
}

func (v *NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Set(val *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth(val *LegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) *NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	return &NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{value: val, isSet: true}
}

func (v NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
