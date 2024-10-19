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

// checks if the MetaIdsQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaIdsQuery{}

// MetaIdsQuery struct for MetaIdsQuery
type MetaIdsQuery struct {
	Values []string `json:"values,omitempty"`
}

// NewMetaIdsQuery instantiates a new MetaIdsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaIdsQuery() *MetaIdsQuery {
	this := MetaIdsQuery{}
	return &this
}

// NewMetaIdsQueryWithDefaults instantiates a new MetaIdsQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaIdsQueryWithDefaults() *MetaIdsQuery {
	this := MetaIdsQuery{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MetaIdsQuery) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIdsQuery) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MetaIdsQuery) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *MetaIdsQuery) SetValues(v []string) {
	o.Values = v
}

func (o MetaIdsQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaIdsQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableMetaIdsQuery struct {
	value *MetaIdsQuery
	isSet bool
}

func (v NullableMetaIdsQuery) Get() *MetaIdsQuery {
	return v.value
}

func (v *NullableMetaIdsQuery) Set(val *MetaIdsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaIdsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaIdsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaIdsQuery(val *MetaIdsQuery) *NullableMetaIdsQuery {
	return &NullableMetaIdsQuery{value: val, isSet: true}
}

func (v NullableMetaIdsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaIdsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
