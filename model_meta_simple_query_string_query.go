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

// checks if the MetaSimpleQueryStringQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaSimpleQueryStringQuery{}

// MetaSimpleQueryStringQuery struct for MetaSimpleQueryStringQuery
type MetaSimpleQueryStringQuery struct {
	AllFields *bool    `json:"all_fields,omitempty"`
	Analyzer  *string  `json:"analyzer,omitempty"`
	Boost     *float32 `json:"boost,omitempty"`
	// or(default), and
	DefaultOperator *string  `json:"default_operator,omitempty"`
	Fields          []string `json:"fields,omitempty"`
	Query           *string  `json:"query,omitempty"`
}

// NewMetaSimpleQueryStringQuery instantiates a new MetaSimpleQueryStringQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaSimpleQueryStringQuery() *MetaSimpleQueryStringQuery {
	this := MetaSimpleQueryStringQuery{}
	return &this
}

// NewMetaSimpleQueryStringQueryWithDefaults instantiates a new MetaSimpleQueryStringQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaSimpleQueryStringQueryWithDefaults() *MetaSimpleQueryStringQuery {
	this := MetaSimpleQueryStringQuery{}
	return &this
}

// GetAllFields returns the AllFields field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetAllFields() bool {
	if o == nil || IsNil(o.AllFields) {
		var ret bool
		return ret
	}
	return *o.AllFields
}

// GetAllFieldsOk returns a tuple with the AllFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetAllFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllFields) {
		return nil, false
	}
	return o.AllFields, true
}

// HasAllFields returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasAllFields() bool {
	if o != nil && !IsNil(o.AllFields) {
		return true
	}

	return false
}

// SetAllFields gets a reference to the given bool and assigns it to the AllFields field.
func (o *MetaSimpleQueryStringQuery) SetAllFields(v bool) {
	o.AllFields = &v
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *MetaSimpleQueryStringQuery) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetBoost() float32 {
	if o == nil || IsNil(o.Boost) {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetBoostOk() (*float32, bool) {
	if o == nil || IsNil(o.Boost) {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasBoost() bool {
	if o != nil && !IsNil(o.Boost) {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaSimpleQueryStringQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetDefaultOperator returns the DefaultOperator field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetDefaultOperator() string {
	if o == nil || IsNil(o.DefaultOperator) {
		var ret string
		return ret
	}
	return *o.DefaultOperator
}

// GetDefaultOperatorOk returns a tuple with the DefaultOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetDefaultOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultOperator) {
		return nil, false
	}
	return o.DefaultOperator, true
}

// HasDefaultOperator returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasDefaultOperator() bool {
	if o != nil && !IsNil(o.DefaultOperator) {
		return true
	}

	return false
}

// SetDefaultOperator gets a reference to the given string and assigns it to the DefaultOperator field.
func (o *MetaSimpleQueryStringQuery) SetDefaultOperator(v string) {
	o.DefaultOperator = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetFields() []string {
	if o == nil || IsNil(o.Fields) {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *MetaSimpleQueryStringQuery) SetFields(v []string) {
	o.Fields = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetaSimpleQueryStringQuery) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaSimpleQueryStringQuery) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetaSimpleQueryStringQuery) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MetaSimpleQueryStringQuery) SetQuery(v string) {
	o.Query = &v
}

func (o MetaSimpleQueryStringQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaSimpleQueryStringQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllFields) {
		toSerialize["all_fields"] = o.AllFields
	}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.Boost) {
		toSerialize["boost"] = o.Boost
	}
	if !IsNil(o.DefaultOperator) {
		toSerialize["default_operator"] = o.DefaultOperator
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableMetaSimpleQueryStringQuery struct {
	value *MetaSimpleQueryStringQuery
	isSet bool
}

func (v NullableMetaSimpleQueryStringQuery) Get() *MetaSimpleQueryStringQuery {
	return v.value
}

func (v *NullableMetaSimpleQueryStringQuery) Set(val *MetaSimpleQueryStringQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaSimpleQueryStringQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaSimpleQueryStringQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaSimpleQueryStringQuery(val *MetaSimpleQueryStringQuery) *NullableMetaSimpleQueryStringQuery {
	return &NullableMetaSimpleQueryStringQuery{value: val, isSet: true}
}

func (v NullableMetaSimpleQueryStringQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaSimpleQueryStringQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
