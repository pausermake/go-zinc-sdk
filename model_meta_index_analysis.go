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

// checks if the MetaIndexAnalysis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaIndexAnalysis{}

// MetaIndexAnalysis struct for MetaIndexAnalysis
type MetaIndexAnalysis struct {
	Analyzer   *map[string]MetaAnalyzer `json:"analyzer,omitempty"`
	CharFilter map[string]interface{}   `json:"char_filter,omitempty"`
	// compatibility with es, alias for TokenFilter
	Filter      map[string]interface{} `json:"filter,omitempty"`
	TokenFilter map[string]interface{} `json:"token_filter,omitempty"`
	Tokenizer   map[string]interface{} `json:"tokenizer,omitempty"`
}

// NewMetaIndexAnalysis instantiates a new MetaIndexAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaIndexAnalysis() *MetaIndexAnalysis {
	this := MetaIndexAnalysis{}
	return &this
}

// NewMetaIndexAnalysisWithDefaults instantiates a new MetaIndexAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaIndexAnalysisWithDefaults() *MetaIndexAnalysis {
	this := MetaIndexAnalysis{}
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *MetaIndexAnalysis) GetAnalyzer() map[string]MetaAnalyzer {
	if o == nil || IsNil(o.Analyzer) {
		var ret map[string]MetaAnalyzer
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexAnalysis) GetAnalyzerOk() (*map[string]MetaAnalyzer, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *MetaIndexAnalysis) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given map[string]MetaAnalyzer and assigns it to the Analyzer field.
func (o *MetaIndexAnalysis) SetAnalyzer(v map[string]MetaAnalyzer) {
	o.Analyzer = &v
}

// GetCharFilter returns the CharFilter field value if set, zero value otherwise.
func (o *MetaIndexAnalysis) GetCharFilter() map[string]interface{} {
	if o == nil || IsNil(o.CharFilter) {
		var ret map[string]interface{}
		return ret
	}
	return o.CharFilter
}

// GetCharFilterOk returns a tuple with the CharFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexAnalysis) GetCharFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CharFilter) {
		return map[string]interface{}{}, false
	}
	return o.CharFilter, true
}

// HasCharFilter returns a boolean if a field has been set.
func (o *MetaIndexAnalysis) HasCharFilter() bool {
	if o != nil && !IsNil(o.CharFilter) {
		return true
	}

	return false
}

// SetCharFilter gets a reference to the given map[string]interface{} and assigns it to the CharFilter field.
func (o *MetaIndexAnalysis) SetCharFilter(v map[string]interface{}) {
	o.CharFilter = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MetaIndexAnalysis) GetFilter() map[string]interface{} {
	if o == nil || IsNil(o.Filter) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexAnalysis) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MetaIndexAnalysis) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *MetaIndexAnalysis) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetTokenFilter returns the TokenFilter field value if set, zero value otherwise.
func (o *MetaIndexAnalysis) GetTokenFilter() map[string]interface{} {
	if o == nil || IsNil(o.TokenFilter) {
		var ret map[string]interface{}
		return ret
	}
	return o.TokenFilter
}

// GetTokenFilterOk returns a tuple with the TokenFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexAnalysis) GetTokenFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TokenFilter) {
		return map[string]interface{}{}, false
	}
	return o.TokenFilter, true
}

// HasTokenFilter returns a boolean if a field has been set.
func (o *MetaIndexAnalysis) HasTokenFilter() bool {
	if o != nil && !IsNil(o.TokenFilter) {
		return true
	}

	return false
}

// SetTokenFilter gets a reference to the given map[string]interface{} and assigns it to the TokenFilter field.
func (o *MetaIndexAnalysis) SetTokenFilter(v map[string]interface{}) {
	o.TokenFilter = v
}

// GetTokenizer returns the Tokenizer field value if set, zero value otherwise.
func (o *MetaIndexAnalysis) GetTokenizer() map[string]interface{} {
	if o == nil || IsNil(o.Tokenizer) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexAnalysis) GetTokenizerOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tokenizer) {
		return map[string]interface{}{}, false
	}
	return o.Tokenizer, true
}

// HasTokenizer returns a boolean if a field has been set.
func (o *MetaIndexAnalysis) HasTokenizer() bool {
	if o != nil && !IsNil(o.Tokenizer) {
		return true
	}

	return false
}

// SetTokenizer gets a reference to the given map[string]interface{} and assigns it to the Tokenizer field.
func (o *MetaIndexAnalysis) SetTokenizer(v map[string]interface{}) {
	o.Tokenizer = v
}

func (o MetaIndexAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaIndexAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.CharFilter) {
		toSerialize["char_filter"] = o.CharFilter
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.TokenFilter) {
		toSerialize["token_filter"] = o.TokenFilter
	}
	if !IsNil(o.Tokenizer) {
		toSerialize["tokenizer"] = o.Tokenizer
	}
	return toSerialize, nil
}

type NullableMetaIndexAnalysis struct {
	value *MetaIndexAnalysis
	isSet bool
}

func (v NullableMetaIndexAnalysis) Get() *MetaIndexAnalysis {
	return v.value
}

func (v *NullableMetaIndexAnalysis) Set(val *MetaIndexAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaIndexAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaIndexAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaIndexAnalysis(val *MetaIndexAnalysis) *NullableMetaIndexAnalysis {
	return &NullableMetaIndexAnalysis{value: val, isSet: true}
}

func (v NullableMetaIndexAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaIndexAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
