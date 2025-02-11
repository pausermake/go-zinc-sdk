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

// checks if the MetaExistsQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaExistsQuery{}

// MetaExistsQuery struct for MetaExistsQuery
type MetaExistsQuery struct {
	Field *string `json:"field,omitempty"`
}

// NewMetaExistsQuery instantiates a new MetaExistsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaExistsQuery() *MetaExistsQuery {
	this := MetaExistsQuery{}
	return &this
}

// NewMetaExistsQueryWithDefaults instantiates a new MetaExistsQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaExistsQueryWithDefaults() *MetaExistsQuery {
	this := MetaExistsQuery{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *MetaExistsQuery) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaExistsQuery) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *MetaExistsQuery) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *MetaExistsQuery) SetField(v string) {
	o.Field = &v
}

func (o MetaExistsQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaExistsQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}

type NullableMetaExistsQuery struct {
	value *MetaExistsQuery
	isSet bool
}

func (v NullableMetaExistsQuery) Get() *MetaExistsQuery {
	return v.value
}

func (v *NullableMetaExistsQuery) Set(val *MetaExistsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaExistsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaExistsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaExistsQuery(val *MetaExistsQuery) *NullableMetaExistsQuery {
	return &NullableMetaExistsQuery{value: val, isSet: true}
}

func (v NullableMetaExistsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaExistsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
