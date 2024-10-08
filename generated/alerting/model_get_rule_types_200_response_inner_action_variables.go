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

// checks if the GetRuleTypes200ResponseInnerActionVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRuleTypes200ResponseInnerActionVariables{}

// GetRuleTypes200ResponseInnerActionVariables A list of action variables that the rule type makes available via context and state in action parameter templates, and a short human readable description. When you create a rule in Kibana, it uses this information to prompt you for these variables in action parameter editors.
type GetRuleTypes200ResponseInnerActionVariables struct {
	Context []GetRuleTypes200ResponseInnerActionVariablesContextInner `json:"context,omitempty"`
	Params  []GetRuleTypes200ResponseInnerActionVariablesParamsInner  `json:"params,omitempty"`
	State   []GetRuleTypes200ResponseInnerActionVariablesStateInner   `json:"state,omitempty"`
}

// NewGetRuleTypes200ResponseInnerActionVariables instantiates a new GetRuleTypes200ResponseInnerActionVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRuleTypes200ResponseInnerActionVariables() *GetRuleTypes200ResponseInnerActionVariables {
	this := GetRuleTypes200ResponseInnerActionVariables{}
	return &this
}

// NewGetRuleTypes200ResponseInnerActionVariablesWithDefaults instantiates a new GetRuleTypes200ResponseInnerActionVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRuleTypes200ResponseInnerActionVariablesWithDefaults() *GetRuleTypes200ResponseInnerActionVariables {
	this := GetRuleTypes200ResponseInnerActionVariables{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetContext() []GetRuleTypes200ResponseInnerActionVariablesContextInner {
	if o == nil || IsNil(o.Context) {
		var ret []GetRuleTypes200ResponseInnerActionVariablesContextInner
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetContextOk() ([]GetRuleTypes200ResponseInnerActionVariablesContextInner, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []GetRuleTypes200ResponseInnerActionVariablesContextInner and assigns it to the Context field.
func (o *GetRuleTypes200ResponseInnerActionVariables) SetContext(v []GetRuleTypes200ResponseInnerActionVariablesContextInner) {
	o.Context = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetParams() []GetRuleTypes200ResponseInnerActionVariablesParamsInner {
	if o == nil || IsNil(o.Params) {
		var ret []GetRuleTypes200ResponseInnerActionVariablesParamsInner
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetParamsOk() ([]GetRuleTypes200ResponseInnerActionVariablesParamsInner, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []GetRuleTypes200ResponseInnerActionVariablesParamsInner and assigns it to the Params field.
func (o *GetRuleTypes200ResponseInnerActionVariables) SetParams(v []GetRuleTypes200ResponseInnerActionVariablesParamsInner) {
	o.Params = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetState() []GetRuleTypes200ResponseInnerActionVariablesStateInner {
	if o == nil || IsNil(o.State) {
		var ret []GetRuleTypes200ResponseInnerActionVariablesStateInner
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) GetStateOk() ([]GetRuleTypes200ResponseInnerActionVariablesStateInner, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerActionVariables) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given []GetRuleTypes200ResponseInnerActionVariablesStateInner and assigns it to the State field.
func (o *GetRuleTypes200ResponseInnerActionVariables) SetState(v []GetRuleTypes200ResponseInnerActionVariablesStateInner) {
	o.State = v
}

func (o GetRuleTypes200ResponseInnerActionVariables) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRuleTypes200ResponseInnerActionVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableGetRuleTypes200ResponseInnerActionVariables struct {
	value *GetRuleTypes200ResponseInnerActionVariables
	isSet bool
}

func (v NullableGetRuleTypes200ResponseInnerActionVariables) Get() *GetRuleTypes200ResponseInnerActionVariables {
	return v.value
}

func (v *NullableGetRuleTypes200ResponseInnerActionVariables) Set(val *GetRuleTypes200ResponseInnerActionVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRuleTypes200ResponseInnerActionVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRuleTypes200ResponseInnerActionVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRuleTypes200ResponseInnerActionVariables(val *GetRuleTypes200ResponseInnerActionVariables) *NullableGetRuleTypes200ResponseInnerActionVariables {
	return &NullableGetRuleTypes200ResponseInnerActionVariables{value: val, isSet: true}
}

func (v NullableGetRuleTypes200ResponseInnerActionVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRuleTypes200ResponseInnerActionVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
