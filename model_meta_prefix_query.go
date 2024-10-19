/*
Zinc Search engine API

Zinc Search engine API documents https://zincsearch-docs.zinc.dev

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MetaPrefixQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaPrefixQuery{}

// MetaPrefixQuery struct for MetaPrefixQuery
type MetaPrefixQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	// You can speed up prefix queries using the index_prefixes mapping parameter.
	Value *string `json:"value,omitempty"`
}

// NewMetaPrefixQuery instantiates a new MetaPrefixQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaPrefixQuery() *MetaPrefixQuery {
	this := MetaPrefixQuery{}
	return &this
}

// NewMetaPrefixQueryWithDefaults instantiates a new MetaPrefixQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaPrefixQueryWithDefaults() *MetaPrefixQuery {
	this := MetaPrefixQuery{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaPrefixQuery) GetBoost() float32 {
	if o == nil || IsNil(o.Boost) {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPrefixQuery) GetBoostOk() (*float32, bool) {
	if o == nil || IsNil(o.Boost) {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaPrefixQuery) HasBoost() bool {
	if o != nil && !IsNil(o.Boost) {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaPrefixQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetaPrefixQuery) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPrefixQuery) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetaPrefixQuery) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MetaPrefixQuery) SetValue(v string) {
	o.Value = &v
}

func (o MetaPrefixQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaPrefixQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Boost) {
		toSerialize["boost"] = o.Boost
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMetaPrefixQuery struct {
	value *MetaPrefixQuery
	isSet bool
}

func (v NullableMetaPrefixQuery) Get() *MetaPrefixQuery {
	return v.value
}

func (v *NullableMetaPrefixQuery) Set(val *MetaPrefixQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaPrefixQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaPrefixQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaPrefixQuery(val *MetaPrefixQuery) *NullableMetaPrefixQuery {
	return &NullableMetaPrefixQuery{value: val, isSet: true}
}

func (v NullableMetaPrefixQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaPrefixQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
