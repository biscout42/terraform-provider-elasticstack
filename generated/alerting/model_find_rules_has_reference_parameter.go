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

// checks if the FindRulesHasReferenceParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindRulesHasReferenceParameter{}

// FindRulesHasReferenceParameter struct for FindRulesHasReferenceParameter
type FindRulesHasReferenceParameter struct {
	Id   NullableString `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
}

// NewFindRulesHasReferenceParameter instantiates a new FindRulesHasReferenceParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindRulesHasReferenceParameter() *FindRulesHasReferenceParameter {
	this := FindRulesHasReferenceParameter{}
	return &this
}

// NewFindRulesHasReferenceParameterWithDefaults instantiates a new FindRulesHasReferenceParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindRulesHasReferenceParameterWithDefaults() *FindRulesHasReferenceParameter {
	this := FindRulesHasReferenceParameter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FindRulesHasReferenceParameter) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FindRulesHasReferenceParameter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FindRulesHasReferenceParameter) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *FindRulesHasReferenceParameter) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *FindRulesHasReferenceParameter) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FindRulesHasReferenceParameter) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FindRulesHasReferenceParameter) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FindRulesHasReferenceParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *FindRulesHasReferenceParameter) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *FindRulesHasReferenceParameter) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *FindRulesHasReferenceParameter) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *FindRulesHasReferenceParameter) UnsetType() {
	o.Type.Unset()
}

func (o FindRulesHasReferenceParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindRulesHasReferenceParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableFindRulesHasReferenceParameter struct {
	value *FindRulesHasReferenceParameter
	isSet bool
}

func (v NullableFindRulesHasReferenceParameter) Get() *FindRulesHasReferenceParameter {
	return v.value
}

func (v *NullableFindRulesHasReferenceParameter) Set(val *FindRulesHasReferenceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFindRulesHasReferenceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFindRulesHasReferenceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindRulesHasReferenceParameter(val *FindRulesHasReferenceParameter) *NullableFindRulesHasReferenceParameter {
	return &NullableFindRulesHasReferenceParameter{value: val, isSet: true}
}

func (v NullableFindRulesHasReferenceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindRulesHasReferenceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
