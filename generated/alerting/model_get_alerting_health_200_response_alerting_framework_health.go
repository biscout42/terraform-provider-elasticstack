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

// checks if the GetAlertingHealth200ResponseAlertingFrameworkHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlertingHealth200ResponseAlertingFrameworkHealth{}

// GetAlertingHealth200ResponseAlertingFrameworkHealth Three substates identify the health of the alerting framework: `decryption_health`, `execution_health`, and `read_health`.
type GetAlertingHealth200ResponseAlertingFrameworkHealth struct {
	DecryptionHealth *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth `json:"decryption_health,omitempty"`
	ExecutionHealth  *GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth  `json:"execution_health,omitempty"`
	ReadHealth       *GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth       `json:"read_health,omitempty"`
}

// NewGetAlertingHealth200ResponseAlertingFrameworkHealth instantiates a new GetAlertingHealth200ResponseAlertingFrameworkHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlertingHealth200ResponseAlertingFrameworkHealth() *GetAlertingHealth200ResponseAlertingFrameworkHealth {
	this := GetAlertingHealth200ResponseAlertingFrameworkHealth{}
	return &this
}

// NewGetAlertingHealth200ResponseAlertingFrameworkHealthWithDefaults instantiates a new GetAlertingHealth200ResponseAlertingFrameworkHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlertingHealth200ResponseAlertingFrameworkHealthWithDefaults() *GetAlertingHealth200ResponseAlertingFrameworkHealth {
	this := GetAlertingHealth200ResponseAlertingFrameworkHealth{}
	return &this
}

// GetDecryptionHealth returns the DecryptionHealth field value if set, zero value otherwise.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetDecryptionHealth() GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	if o == nil || IsNil(o.DecryptionHealth) {
		var ret GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth
		return ret
	}
	return *o.DecryptionHealth
}

// GetDecryptionHealthOk returns a tuple with the DecryptionHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetDecryptionHealthOk() (*GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth, bool) {
	if o == nil || IsNil(o.DecryptionHealth) {
		return nil, false
	}
	return o.DecryptionHealth, true
}

// HasDecryptionHealth returns a boolean if a field has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) HasDecryptionHealth() bool {
	if o != nil && !IsNil(o.DecryptionHealth) {
		return true
	}

	return false
}

// SetDecryptionHealth gets a reference to the given GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth and assigns it to the DecryptionHealth field.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) SetDecryptionHealth(v GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) {
	o.DecryptionHealth = &v
}

// GetExecutionHealth returns the ExecutionHealth field value if set, zero value otherwise.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetExecutionHealth() GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth {
	if o == nil || IsNil(o.ExecutionHealth) {
		var ret GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth
		return ret
	}
	return *o.ExecutionHealth
}

// GetExecutionHealthOk returns a tuple with the ExecutionHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetExecutionHealthOk() (*GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth, bool) {
	if o == nil || IsNil(o.ExecutionHealth) {
		return nil, false
	}
	return o.ExecutionHealth, true
}

// HasExecutionHealth returns a boolean if a field has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) HasExecutionHealth() bool {
	if o != nil && !IsNil(o.ExecutionHealth) {
		return true
	}

	return false
}

// SetExecutionHealth gets a reference to the given GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth and assigns it to the ExecutionHealth field.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) SetExecutionHealth(v GetAlertingHealth200ResponseAlertingFrameworkHealthExecutionHealth) {
	o.ExecutionHealth = &v
}

// GetReadHealth returns the ReadHealth field value if set, zero value otherwise.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetReadHealth() GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth {
	if o == nil || IsNil(o.ReadHealth) {
		var ret GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth
		return ret
	}
	return *o.ReadHealth
}

// GetReadHealthOk returns a tuple with the ReadHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) GetReadHealthOk() (*GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth, bool) {
	if o == nil || IsNil(o.ReadHealth) {
		return nil, false
	}
	return o.ReadHealth, true
}

// HasReadHealth returns a boolean if a field has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) HasReadHealth() bool {
	if o != nil && !IsNil(o.ReadHealth) {
		return true
	}

	return false
}

// SetReadHealth gets a reference to the given GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth and assigns it to the ReadHealth field.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealth) SetReadHealth(v GetAlertingHealth200ResponseAlertingFrameworkHealthReadHealth) {
	o.ReadHealth = &v
}

func (o GetAlertingHealth200ResponseAlertingFrameworkHealth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlertingHealth200ResponseAlertingFrameworkHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DecryptionHealth) {
		toSerialize["decryption_health"] = o.DecryptionHealth
	}
	if !IsNil(o.ExecutionHealth) {
		toSerialize["execution_health"] = o.ExecutionHealth
	}
	if !IsNil(o.ReadHealth) {
		toSerialize["read_health"] = o.ReadHealth
	}
	return toSerialize, nil
}

type NullableGetAlertingHealth200ResponseAlertingFrameworkHealth struct {
	value *GetAlertingHealth200ResponseAlertingFrameworkHealth
	isSet bool
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) Get() *GetAlertingHealth200ResponseAlertingFrameworkHealth {
	return v.value
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) Set(val *GetAlertingHealth200ResponseAlertingFrameworkHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlertingHealth200ResponseAlertingFrameworkHealth(val *GetAlertingHealth200ResponseAlertingFrameworkHealth) *NullableGetAlertingHealth200ResponseAlertingFrameworkHealth {
	return &NullableGetAlertingHealth200ResponseAlertingFrameworkHealth{value: val, isSet: true}
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
