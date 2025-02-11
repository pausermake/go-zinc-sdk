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

// checks if the MetaMatchBoolPrefixQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaMatchBoolPrefixQuery{}

// MetaMatchBoolPrefixQuery struct for MetaMatchBoolPrefixQuery
type MetaMatchBoolPrefixQuery struct {
	Analyzer *string  `json:"analyzer,omitempty"`
	Boost    *float32 `json:"boost,omitempty"`
	Query    *string  `json:"query,omitempty"`
}

// NewMetaMatchBoolPrefixQuery instantiates a new MetaMatchBoolPrefixQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaMatchBoolPrefixQuery() *MetaMatchBoolPrefixQuery {
	this := MetaMatchBoolPrefixQuery{}
	return &this
}

// NewMetaMatchBoolPrefixQueryWithDefaults instantiates a new MetaMatchBoolPrefixQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaMatchBoolPrefixQueryWithDefaults() *MetaMatchBoolPrefixQuery {
	this := MetaMatchBoolPrefixQuery{}
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *MetaMatchBoolPrefixQuery) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchBoolPrefixQuery) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *MetaMatchBoolPrefixQuery) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *MetaMatchBoolPrefixQuery) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaMatchBoolPrefixQuery) GetBoost() float32 {
	if o == nil || IsNil(o.Boost) {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchBoolPrefixQuery) GetBoostOk() (*float32, bool) {
	if o == nil || IsNil(o.Boost) {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaMatchBoolPrefixQuery) HasBoost() bool {
	if o != nil && !IsNil(o.Boost) {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaMatchBoolPrefixQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetaMatchBoolPrefixQuery) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchBoolPrefixQuery) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetaMatchBoolPrefixQuery) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MetaMatchBoolPrefixQuery) SetQuery(v string) {
	o.Query = &v
}

func (o MetaMatchBoolPrefixQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaMatchBoolPrefixQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.Boost) {
		toSerialize["boost"] = o.Boost
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableMetaMatchBoolPrefixQuery struct {
	value *MetaMatchBoolPrefixQuery
	isSet bool
}

func (v NullableMetaMatchBoolPrefixQuery) Get() *MetaMatchBoolPrefixQuery {
	return v.value
}

func (v *NullableMetaMatchBoolPrefixQuery) Set(val *MetaMatchBoolPrefixQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaMatchBoolPrefixQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaMatchBoolPrefixQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaMatchBoolPrefixQuery(val *MetaMatchBoolPrefixQuery) *NullableMetaMatchBoolPrefixQuery {
	return &NullableMetaMatchBoolPrefixQuery{value: val, isSet: true}
}

func (v NullableMetaMatchBoolPrefixQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaMatchBoolPrefixQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
