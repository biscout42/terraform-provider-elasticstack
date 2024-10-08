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

// checks if the GetAlertingHealth200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlertingHealth200Response{}

// GetAlertingHealth200Response struct for GetAlertingHealth200Response
type GetAlertingHealth200Response struct {
	AlertingFrameworkHealth *GetAlertingHealth200ResponseAlertingFrameworkHealth `json:"alerting_framework_health,omitempty"`
	// If `false`, the encrypted saved object plugin does not have a permanent encryption key.
	HasPermanentEncryptionKey *bool `json:"has_permanent_encryption_key,omitempty"`
	// If `false`, security is enabled but TLS is not.
	IsSufficientlySecure *bool `json:"is_sufficiently_secure,omitempty"`
}

// NewGetAlertingHealth200Response instantiates a new GetAlertingHealth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlertingHealth200Response() *GetAlertingHealth200Response {
	this := GetAlertingHealth200Response{}
	return &this
}

// NewGetAlertingHealth200ResponseWithDefaults instantiates a new GetAlertingHealth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlertingHealth200ResponseWithDefaults() *GetAlertingHealth200Response {
	this := GetAlertingHealth200Response{}
	return &this
}

// GetAlertingFrameworkHealth returns the AlertingFrameworkHealth field value if set, zero value otherwise.
func (o *GetAlertingHealth200Response) GetAlertingFrameworkHealth() GetAlertingHealth200ResponseAlertingFrameworkHealth {
	if o == nil || IsNil(o.AlertingFrameworkHealth) {
		var ret GetAlertingHealth200ResponseAlertingFrameworkHealth
		return ret
	}
	return *o.AlertingFrameworkHealth
}

// GetAlertingFrameworkHealthOk returns a tuple with the AlertingFrameworkHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200Response) GetAlertingFrameworkHealthOk() (*GetAlertingHealth200ResponseAlertingFrameworkHealth, bool) {
	if o == nil || IsNil(o.AlertingFrameworkHealth) {
		return nil, false
	}
	return o.AlertingFrameworkHealth, true
}

// HasAlertingFrameworkHealth returns a boolean if a field has been set.
func (o *GetAlertingHealth200Response) HasAlertingFrameworkHealth() bool {
	if o != nil && !IsNil(o.AlertingFrameworkHealth) {
		return true
	}

	return false
}

// SetAlertingFrameworkHealth gets a reference to the given GetAlertingHealth200ResponseAlertingFrameworkHealth and assigns it to the AlertingFrameworkHealth field.
func (o *GetAlertingHealth200Response) SetAlertingFrameworkHealth(v GetAlertingHealth200ResponseAlertingFrameworkHealth) {
	o.AlertingFrameworkHealth = &v
}

// GetHasPermanentEncryptionKey returns the HasPermanentEncryptionKey field value if set, zero value otherwise.
func (o *GetAlertingHealth200Response) GetHasPermanentEncryptionKey() bool {
	if o == nil || IsNil(o.HasPermanentEncryptionKey) {
		var ret bool
		return ret
	}
	return *o.HasPermanentEncryptionKey
}

// GetHasPermanentEncryptionKeyOk returns a tuple with the HasPermanentEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200Response) GetHasPermanentEncryptionKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPermanentEncryptionKey) {
		return nil, false
	}
	return o.HasPermanentEncryptionKey, true
}

// HasHasPermanentEncryptionKey returns a boolean if a field has been set.
func (o *GetAlertingHealth200Response) HasHasPermanentEncryptionKey() bool {
	if o != nil && !IsNil(o.HasPermanentEncryptionKey) {
		return true
	}

	return false
}

// SetHasPermanentEncryptionKey gets a reference to the given bool and assigns it to the HasPermanentEncryptionKey field.
func (o *GetAlertingHealth200Response) SetHasPermanentEncryptionKey(v bool) {
	o.HasPermanentEncryptionKey = &v
}

// GetIsSufficientlySecure returns the IsSufficientlySecure field value if set, zero value otherwise.
func (o *GetAlertingHealth200Response) GetIsSufficientlySecure() bool {
	if o == nil || IsNil(o.IsSufficientlySecure) {
		var ret bool
		return ret
	}
	return *o.IsSufficientlySecure
}

// GetIsSufficientlySecureOk returns a tuple with the IsSufficientlySecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200Response) GetIsSufficientlySecureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSufficientlySecure) {
		return nil, false
	}
	return o.IsSufficientlySecure, true
}

// HasIsSufficientlySecure returns a boolean if a field has been set.
func (o *GetAlertingHealth200Response) HasIsSufficientlySecure() bool {
	if o != nil && !IsNil(o.IsSufficientlySecure) {
		return true
	}

	return false
}

// SetIsSufficientlySecure gets a reference to the given bool and assigns it to the IsSufficientlySecure field.
func (o *GetAlertingHealth200Response) SetIsSufficientlySecure(v bool) {
	o.IsSufficientlySecure = &v
}

func (o GetAlertingHealth200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlertingHealth200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertingFrameworkHealth) {
		toSerialize["alerting_framework_health"] = o.AlertingFrameworkHealth
	}
	if !IsNil(o.HasPermanentEncryptionKey) {
		toSerialize["has_permanent_encryption_key"] = o.HasPermanentEncryptionKey
	}
	if !IsNil(o.IsSufficientlySecure) {
		toSerialize["is_sufficiently_secure"] = o.IsSufficientlySecure
	}
	return toSerialize, nil
}

type NullableGetAlertingHealth200Response struct {
	value *GetAlertingHealth200Response
	isSet bool
}

func (v NullableGetAlertingHealth200Response) Get() *GetAlertingHealth200Response {
	return v.value
}

func (v *NullableGetAlertingHealth200Response) Set(val *GetAlertingHealth200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlertingHealth200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlertingHealth200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlertingHealth200Response(val *GetAlertingHealth200Response) *NullableGetAlertingHealth200Response {
	return &NullableGetAlertingHealth200Response{value: val, isSet: true}
}

func (v NullableGetAlertingHealth200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlertingHealth200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
