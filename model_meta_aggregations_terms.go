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

// checks if the MetaAggregationsTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaAggregationsTerms{}

// MetaAggregationsTerms struct for MetaAggregationsTerms
type MetaAggregationsTerms struct {
	Field *string `json:"field,omitempty"`
	// { \"_count\": \"asc\" }
	Order *map[string]string `json:"order,omitempty"`
	Size  *int32             `json:"size,omitempty"`
}

// NewMetaAggregationsTerms instantiates a new MetaAggregationsTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAggregationsTerms() *MetaAggregationsTerms {
	this := MetaAggregationsTerms{}
	return &this
}

// NewMetaAggregationsTermsWithDefaults instantiates a new MetaAggregationsTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAggregationsTermsWithDefaults() *MetaAggregationsTerms {
	this := MetaAggregationsTerms{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *MetaAggregationsTerms) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationsTerms) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *MetaAggregationsTerms) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *MetaAggregationsTerms) SetField(v string) {
	o.Field = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *MetaAggregationsTerms) GetOrder() map[string]string {
	if o == nil || IsNil(o.Order) {
		var ret map[string]string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationsTerms) GetOrderOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *MetaAggregationsTerms) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given map[string]string and assigns it to the Order field.
func (o *MetaAggregationsTerms) SetOrder(v map[string]string) {
	o.Order = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MetaAggregationsTerms) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationsTerms) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MetaAggregationsTerms) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MetaAggregationsTerms) SetSize(v int32) {
	o.Size = &v
}

func (o MetaAggregationsTerms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaAggregationsTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableMetaAggregationsTerms struct {
	value *MetaAggregationsTerms
	isSet bool
}

func (v NullableMetaAggregationsTerms) Get() *MetaAggregationsTerms {
	return v.value
}

func (v *NullableMetaAggregationsTerms) Set(val *MetaAggregationsTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAggregationsTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAggregationsTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAggregationsTerms(val *MetaAggregationsTerms) *NullableMetaAggregationsTerms {
	return &NullableMetaAggregationsTerms{value: val, isSet: true}
}

func (v NullableMetaAggregationsTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAggregationsTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
