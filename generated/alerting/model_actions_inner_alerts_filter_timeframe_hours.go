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

// checks if the ActionsInnerAlertsFilterTimeframeHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionsInnerAlertsFilterTimeframeHours{}

// ActionsInnerAlertsFilterTimeframeHours Defines the range of time in a day that the action can run. If the `start` value is `00:00` and the `end` value is `24:00`, actions be generated all day.
type ActionsInnerAlertsFilterTimeframeHours struct {
	// The end of the time frame in 24-hour notation (`hh:mm`).
	End *string `json:"end,omitempty"`
	// The start of the time frame in 24-hour notation (`hh:mm`).
	Start *string `json:"start,omitempty"`
}

// NewActionsInnerAlertsFilterTimeframeHours instantiates a new ActionsInnerAlertsFilterTimeframeHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsInnerAlertsFilterTimeframeHours() *ActionsInnerAlertsFilterTimeframeHours {
	this := ActionsInnerAlertsFilterTimeframeHours{}
	return &this
}

// NewActionsInnerAlertsFilterTimeframeHoursWithDefaults instantiates a new ActionsInnerAlertsFilterTimeframeHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsInnerAlertsFilterTimeframeHoursWithDefaults() *ActionsInnerAlertsFilterTimeframeHours {
	this := ActionsInnerAlertsFilterTimeframeHours{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ActionsInnerAlertsFilterTimeframeHours) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsInnerAlertsFilterTimeframeHours) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ActionsInnerAlertsFilterTimeframeHours) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *ActionsInnerAlertsFilterTimeframeHours) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ActionsInnerAlertsFilterTimeframeHours) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsInnerAlertsFilterTimeframeHours) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ActionsInnerAlertsFilterTimeframeHours) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *ActionsInnerAlertsFilterTimeframeHours) SetStart(v string) {
	o.Start = &v
}

func (o ActionsInnerAlertsFilterTimeframeHours) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionsInnerAlertsFilterTimeframeHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableActionsInnerAlertsFilterTimeframeHours struct {
	value *ActionsInnerAlertsFilterTimeframeHours
	isSet bool
}

func (v NullableActionsInnerAlertsFilterTimeframeHours) Get() *ActionsInnerAlertsFilterTimeframeHours {
	return v.value
}

func (v *NullableActionsInnerAlertsFilterTimeframeHours) Set(val *ActionsInnerAlertsFilterTimeframeHours) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsInnerAlertsFilterTimeframeHours) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsInnerAlertsFilterTimeframeHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsInnerAlertsFilterTimeframeHours(val *ActionsInnerAlertsFilterTimeframeHours) *NullableActionsInnerAlertsFilterTimeframeHours {
	return &NullableActionsInnerAlertsFilterTimeframeHours{value: val, isSet: true}
}

func (v NullableActionsInnerAlertsFilterTimeframeHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsInnerAlertsFilterTimeframeHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}