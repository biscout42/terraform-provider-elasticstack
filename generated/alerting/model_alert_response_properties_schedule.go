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

// checks if the AlertResponsePropertiesSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertResponsePropertiesSchedule{}

// AlertResponsePropertiesSchedule struct for AlertResponsePropertiesSchedule
type AlertResponsePropertiesSchedule struct {
	Interval *string `json:"interval,omitempty"`
}

// NewAlertResponsePropertiesSchedule instantiates a new AlertResponsePropertiesSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertResponsePropertiesSchedule() *AlertResponsePropertiesSchedule {
	this := AlertResponsePropertiesSchedule{}
	return &this
}

// NewAlertResponsePropertiesScheduleWithDefaults instantiates a new AlertResponsePropertiesSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertResponsePropertiesScheduleWithDefaults() *AlertResponsePropertiesSchedule {
	this := AlertResponsePropertiesSchedule{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AlertResponsePropertiesSchedule) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponsePropertiesSchedule) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AlertResponsePropertiesSchedule) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *AlertResponsePropertiesSchedule) SetInterval(v string) {
	o.Interval = &v
}

func (o AlertResponsePropertiesSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertResponsePropertiesSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableAlertResponsePropertiesSchedule struct {
	value *AlertResponsePropertiesSchedule
	isSet bool
}

func (v NullableAlertResponsePropertiesSchedule) Get() *AlertResponsePropertiesSchedule {
	return v.value
}

func (v *NullableAlertResponsePropertiesSchedule) Set(val *AlertResponsePropertiesSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertResponsePropertiesSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertResponsePropertiesSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertResponsePropertiesSchedule(val *AlertResponsePropertiesSchedule) *NullableAlertResponsePropertiesSchedule {
	return &NullableAlertResponsePropertiesSchedule{value: val, isSet: true}
}

func (v NullableAlertResponsePropertiesSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertResponsePropertiesSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
