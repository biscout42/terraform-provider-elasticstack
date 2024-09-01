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

// checks if the GetRuleTypes200ResponseInnerAlertsMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRuleTypes200ResponseInnerAlertsMappings{}

// GetRuleTypes200ResponseInnerAlertsMappings struct for GetRuleTypes200ResponseInnerAlertsMappings
type GetRuleTypes200ResponseInnerAlertsMappings struct {
	// Mapping information for each field supported in alerts as data documents for this rule type. For more information about mapping parameters, refer to the Elasticsearch documentation.
	FieldMap map[string]FieldmapProperties `json:"fieldMap,omitempty"`
}

// NewGetRuleTypes200ResponseInnerAlertsMappings instantiates a new GetRuleTypes200ResponseInnerAlertsMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRuleTypes200ResponseInnerAlertsMappings() *GetRuleTypes200ResponseInnerAlertsMappings {
	this := GetRuleTypes200ResponseInnerAlertsMappings{}
	return &this
}

// NewGetRuleTypes200ResponseInnerAlertsMappingsWithDefaults instantiates a new GetRuleTypes200ResponseInnerAlertsMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRuleTypes200ResponseInnerAlertsMappingsWithDefaults() *GetRuleTypes200ResponseInnerAlertsMappings {
	this := GetRuleTypes200ResponseInnerAlertsMappings{}
	return &this
}

// GetFieldMap returns the FieldMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetRuleTypes200ResponseInnerAlertsMappings) GetFieldMap() map[string]FieldmapProperties {
	if o == nil {
		var ret map[string]FieldmapProperties
		return ret
	}
	return o.FieldMap
}

// GetFieldMapOk returns a tuple with the FieldMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetRuleTypes200ResponseInnerAlertsMappings) GetFieldMapOk() (*map[string]FieldmapProperties, bool) {
	if o == nil || IsNil(o.FieldMap) {
		return nil, false
	}
	return &o.FieldMap, true
}

// HasFieldMap returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerAlertsMappings) HasFieldMap() bool {
	if o != nil && IsNil(o.FieldMap) {
		return true
	}

	return false
}

// SetFieldMap gets a reference to the given map[string]FieldmapProperties and assigns it to the FieldMap field.
func (o *GetRuleTypes200ResponseInnerAlertsMappings) SetFieldMap(v map[string]FieldmapProperties) {
	o.FieldMap = v
}

func (o GetRuleTypes200ResponseInnerAlertsMappings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRuleTypes200ResponseInnerAlertsMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldMap != nil {
		toSerialize["fieldMap"] = o.FieldMap
	}
	return toSerialize, nil
}

type NullableGetRuleTypes200ResponseInnerAlertsMappings struct {
	value *GetRuleTypes200ResponseInnerAlertsMappings
	isSet bool
}

func (v NullableGetRuleTypes200ResponseInnerAlertsMappings) Get() *GetRuleTypes200ResponseInnerAlertsMappings {
	return v.value
}

func (v *NullableGetRuleTypes200ResponseInnerAlertsMappings) Set(val *GetRuleTypes200ResponseInnerAlertsMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRuleTypes200ResponseInnerAlertsMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRuleTypes200ResponseInnerAlertsMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRuleTypes200ResponseInnerAlertsMappings(val *GetRuleTypes200ResponseInnerAlertsMappings) *NullableGetRuleTypes200ResponseInnerAlertsMappings {
	return &NullableGetRuleTypes200ResponseInnerAlertsMappings{value: val, isSet: true}
}

func (v NullableGetRuleTypes200ResponseInnerAlertsMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRuleTypes200ResponseInnerAlertsMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}