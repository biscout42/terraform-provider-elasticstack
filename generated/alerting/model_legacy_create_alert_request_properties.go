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

// checks if the LegacyCreateAlertRequestProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyCreateAlertRequestProperties{}

// LegacyCreateAlertRequestProperties struct for LegacyCreateAlertRequestProperties
type LegacyCreateAlertRequestProperties struct {
	Actions []LegacyUpdateAlertRequestPropertiesActionsInner `json:"actions,omitempty"`
	// The ID of the alert type that you want to call when the alert is scheduled to run.
	AlertTypeId string `json:"alertTypeId"`
	// The name of the application that owns the alert. This name has to match the Kibana feature name, as that dictates the required role-based access control privileges.
	Consumer string `json:"consumer"`
	// Indicates if you want to run the alert on an interval basis after it is created.
	Enabled *bool `json:"enabled,omitempty"`
	// A name to reference and search.
	Name string `json:"name"`
	// The condition for throttling the notification.
	NotifyWhen string `json:"notifyWhen"`
	// The parameters to pass to the alert type executor `params` value. This will also validate against the alert type params validator, if defined.
	Params   map[string]interface{}                     `json:"params"`
	Schedule LegacyUpdateAlertRequestPropertiesSchedule `json:"schedule"`
	Tags     []string                                   `json:"tags,omitempty"`
	// How often this alert should fire the same actions. This will prevent the alert from sending out the same notification over and over. For example, if an alert with a schedule of 1 minute stays in a triggered state for 90 minutes, setting a throttle of `10m` or `1h` will prevent it from sending 90 notifications during this period.
	Throttle *string `json:"throttle,omitempty"`
}

// NewLegacyCreateAlertRequestProperties instantiates a new LegacyCreateAlertRequestProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyCreateAlertRequestProperties(alertTypeId string, consumer string, name string, notifyWhen string, params map[string]interface{}, schedule LegacyUpdateAlertRequestPropertiesSchedule) *LegacyCreateAlertRequestProperties {
	this := LegacyCreateAlertRequestProperties{}
	this.AlertTypeId = alertTypeId
	this.Consumer = consumer
	this.Name = name
	this.NotifyWhen = notifyWhen
	this.Params = params
	this.Schedule = schedule
	return &this
}

// NewLegacyCreateAlertRequestPropertiesWithDefaults instantiates a new LegacyCreateAlertRequestProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyCreateAlertRequestPropertiesWithDefaults() *LegacyCreateAlertRequestProperties {
	this := LegacyCreateAlertRequestProperties{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *LegacyCreateAlertRequestProperties) GetActions() []LegacyUpdateAlertRequestPropertiesActionsInner {
	if o == nil || IsNil(o.Actions) {
		var ret []LegacyUpdateAlertRequestPropertiesActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetActionsOk() ([]LegacyUpdateAlertRequestPropertiesActionsInner, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *LegacyCreateAlertRequestProperties) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []LegacyUpdateAlertRequestPropertiesActionsInner and assigns it to the Actions field.
func (o *LegacyCreateAlertRequestProperties) SetActions(v []LegacyUpdateAlertRequestPropertiesActionsInner) {
	o.Actions = v
}

// GetAlertTypeId returns the AlertTypeId field value
func (o *LegacyCreateAlertRequestProperties) GetAlertTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetAlertTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertTypeId, true
}

// SetAlertTypeId sets field value
func (o *LegacyCreateAlertRequestProperties) SetAlertTypeId(v string) {
	o.AlertTypeId = v
}

// GetConsumer returns the Consumer field value
func (o *LegacyCreateAlertRequestProperties) GetConsumer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetConsumerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consumer, true
}

// SetConsumer sets field value
func (o *LegacyCreateAlertRequestProperties) SetConsumer(v string) {
	o.Consumer = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LegacyCreateAlertRequestProperties) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LegacyCreateAlertRequestProperties) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LegacyCreateAlertRequestProperties) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value
func (o *LegacyCreateAlertRequestProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LegacyCreateAlertRequestProperties) SetName(v string) {
	o.Name = v
}

// GetNotifyWhen returns the NotifyWhen field value
func (o *LegacyCreateAlertRequestProperties) GetNotifyWhen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifyWhen
}

// GetNotifyWhenOk returns a tuple with the NotifyWhen field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetNotifyWhenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyWhen, true
}

// SetNotifyWhen sets field value
func (o *LegacyCreateAlertRequestProperties) SetNotifyWhen(v string) {
	o.NotifyWhen = v
}

// GetParams returns the Params field value
func (o *LegacyCreateAlertRequestProperties) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *LegacyCreateAlertRequestProperties) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetSchedule returns the Schedule field value
func (o *LegacyCreateAlertRequestProperties) GetSchedule() LegacyUpdateAlertRequestPropertiesSchedule {
	if o == nil {
		var ret LegacyUpdateAlertRequestPropertiesSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetScheduleOk() (*LegacyUpdateAlertRequestPropertiesSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *LegacyCreateAlertRequestProperties) SetSchedule(v LegacyUpdateAlertRequestPropertiesSchedule) {
	o.Schedule = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LegacyCreateAlertRequestProperties) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LegacyCreateAlertRequestProperties) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LegacyCreateAlertRequestProperties) SetTags(v []string) {
	o.Tags = v
}

// GetThrottle returns the Throttle field value if set, zero value otherwise.
func (o *LegacyCreateAlertRequestProperties) GetThrottle() string {
	if o == nil || IsNil(o.Throttle) {
		var ret string
		return ret
	}
	return *o.Throttle
}

// GetThrottleOk returns a tuple with the Throttle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyCreateAlertRequestProperties) GetThrottleOk() (*string, bool) {
	if o == nil || IsNil(o.Throttle) {
		return nil, false
	}
	return o.Throttle, true
}

// HasThrottle returns a boolean if a field has been set.
func (o *LegacyCreateAlertRequestProperties) HasThrottle() bool {
	if o != nil && !IsNil(o.Throttle) {
		return true
	}

	return false
}

// SetThrottle gets a reference to the given string and assigns it to the Throttle field.
func (o *LegacyCreateAlertRequestProperties) SetThrottle(v string) {
	o.Throttle = &v
}

func (o LegacyCreateAlertRequestProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyCreateAlertRequestProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	toSerialize["alertTypeId"] = o.AlertTypeId
	toSerialize["consumer"] = o.Consumer
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["name"] = o.Name
	toSerialize["notifyWhen"] = o.NotifyWhen
	toSerialize["params"] = o.Params
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Throttle) {
		toSerialize["throttle"] = o.Throttle
	}
	return toSerialize, nil
}

type NullableLegacyCreateAlertRequestProperties struct {
	value *LegacyCreateAlertRequestProperties
	isSet bool
}

func (v NullableLegacyCreateAlertRequestProperties) Get() *LegacyCreateAlertRequestProperties {
	return v.value
}

func (v *NullableLegacyCreateAlertRequestProperties) Set(val *LegacyCreateAlertRequestProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyCreateAlertRequestProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyCreateAlertRequestProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyCreateAlertRequestProperties(val *LegacyCreateAlertRequestProperties) *NullableLegacyCreateAlertRequestProperties {
	return &NullableLegacyCreateAlertRequestProperties{value: val, isSet: true}
}

func (v NullableLegacyCreateAlertRequestProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyCreateAlertRequestProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
