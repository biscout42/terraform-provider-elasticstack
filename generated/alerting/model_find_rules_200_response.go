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

// checks if the FindRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindRules200Response{}

// FindRules200Response struct for FindRules200Response
type FindRules200Response struct {
	Data    []RuleResponseProperties `json:"data,omitempty"`
	Page    *int32                   `json:"page,omitempty"`
	PerPage *int32                   `json:"per_page,omitempty"`
	Total   *int32                   `json:"total,omitempty"`
}

// NewFindRules200Response instantiates a new FindRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindRules200Response() *FindRules200Response {
	this := FindRules200Response{}
	return &this
}

// NewFindRules200ResponseWithDefaults instantiates a new FindRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindRules200ResponseWithDefaults() *FindRules200Response {
	this := FindRules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FindRules200Response) GetData() []RuleResponseProperties {
	if o == nil || IsNil(o.Data) {
		var ret []RuleResponseProperties
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindRules200Response) GetDataOk() ([]RuleResponseProperties, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FindRules200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RuleResponseProperties and assigns it to the Data field.
func (o *FindRules200Response) SetData(v []RuleResponseProperties) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *FindRules200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindRules200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *FindRules200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *FindRules200Response) SetPage(v int32) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *FindRules200Response) GetPerPage() int32 {
	if o == nil || IsNil(o.PerPage) {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindRules200Response) GetPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *FindRules200Response) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *FindRules200Response) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FindRules200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindRules200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FindRules200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *FindRules200Response) SetTotal(v int32) {
	o.Total = &v
}

func (o FindRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindRules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableFindRules200Response struct {
	value *FindRules200Response
	isSet bool
}

func (v NullableFindRules200Response) Get() *FindRules200Response {
	return v.value
}

func (v *NullableFindRules200Response) Set(val *FindRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindRules200Response(val *FindRules200Response) *NullableFindRules200Response {
	return &NullableFindRules200Response{value: val, isSet: true}
}

func (v NullableFindRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
