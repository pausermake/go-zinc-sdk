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

// checks if the MetaMultiMatchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaMultiMatchQuery{}

// MetaMultiMatchQuery struct for MetaMultiMatchQuery
type MetaMultiMatchQuery struct {
	Analyzer           *string  `json:"analyzer,omitempty"`
	Boost              *float32 `json:"boost,omitempty"`
	Fields             []string `json:"fields,omitempty"`
	MinimumShouldMatch *float32 `json:"minimum_should_match,omitempty"`
	// or(default), and
	Operator *string `json:"operator,omitempty"`
	Query    *string `json:"query,omitempty"`
	// best_fields(default), most_fields, cross_fields, phrase, phrase_prefix, bool_prefix
	Type *string `json:"type,omitempty"`
}

// NewMetaMultiMatchQuery instantiates a new MetaMultiMatchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaMultiMatchQuery() *MetaMultiMatchQuery {
	this := MetaMultiMatchQuery{}
	return &this
}

// NewMetaMultiMatchQueryWithDefaults instantiates a new MetaMultiMatchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaMultiMatchQueryWithDefaults() *MetaMultiMatchQuery {
	this := MetaMultiMatchQuery{}
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *MetaMultiMatchQuery) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetBoost() float32 {
	if o == nil || IsNil(o.Boost) {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetBoostOk() (*float32, bool) {
	if o == nil || IsNil(o.Boost) {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasBoost() bool {
	if o != nil && !IsNil(o.Boost) {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaMultiMatchQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetFields() []string {
	if o == nil || IsNil(o.Fields) {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *MetaMultiMatchQuery) SetFields(v []string) {
	o.Fields = v
}

// GetMinimumShouldMatch returns the MinimumShouldMatch field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetMinimumShouldMatch() float32 {
	if o == nil || IsNil(o.MinimumShouldMatch) {
		var ret float32
		return ret
	}
	return *o.MinimumShouldMatch
}

// GetMinimumShouldMatchOk returns a tuple with the MinimumShouldMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetMinimumShouldMatchOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumShouldMatch) {
		return nil, false
	}
	return o.MinimumShouldMatch, true
}

// HasMinimumShouldMatch returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasMinimumShouldMatch() bool {
	if o != nil && !IsNil(o.MinimumShouldMatch) {
		return true
	}

	return false
}

// SetMinimumShouldMatch gets a reference to the given float32 and assigns it to the MinimumShouldMatch field.
func (o *MetaMultiMatchQuery) SetMinimumShouldMatch(v float32) {
	o.MinimumShouldMatch = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *MetaMultiMatchQuery) SetOperator(v string) {
	o.Operator = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MetaMultiMatchQuery) SetQuery(v string) {
	o.Query = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetaMultiMatchQuery) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMultiMatchQuery) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetaMultiMatchQuery) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetaMultiMatchQuery) SetType(v string) {
	o.Type = &v
}

func (o MetaMultiMatchQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaMultiMatchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.Boost) {
		toSerialize["boost"] = o.Boost
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.MinimumShouldMatch) {
		toSerialize["minimum_should_match"] = o.MinimumShouldMatch
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMetaMultiMatchQuery struct {
	value *MetaMultiMatchQuery
	isSet bool
}

func (v NullableMetaMultiMatchQuery) Get() *MetaMultiMatchQuery {
	return v.value
}

func (v *NullableMetaMultiMatchQuery) Set(val *MetaMultiMatchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaMultiMatchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaMultiMatchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaMultiMatchQuery(val *MetaMultiMatchQuery) *NullableMetaMultiMatchQuery {
	return &NullableMetaMultiMatchQuery{value: val, isSet: true}
}

func (v NullableMetaMultiMatchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaMultiMatchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
