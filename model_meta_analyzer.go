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

// checks if the MetaAnalyzer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaAnalyzer{}

// MetaAnalyzer struct for MetaAnalyzer
type MetaAnalyzer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	// compatibility with es, alias for TokenFilter
	Filter []string `json:"filter,omitempty"`
	// for type=pattern
	Lowercase *bool `json:"lowercase,omitempty"`
	// for type=pattern
	Pattern *string `json:"pattern,omitempty"`
	// for type=pattern,standard,stop
	Stopwords   []string `json:"stopwords,omitempty"`
	TokenFilter []string `json:"token_filter,omitempty"`
	Tokenizer   *string  `json:"tokenizer,omitempty"`
	// options for compatible
	Type *string `json:"type,omitempty"`
}

// NewMetaAnalyzer instantiates a new MetaAnalyzer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAnalyzer() *MetaAnalyzer {
	this := MetaAnalyzer{}
	return &this
}

// NewMetaAnalyzerWithDefaults instantiates a new MetaAnalyzer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAnalyzerWithDefaults() *MetaAnalyzer {
	this := MetaAnalyzer{}
	return &this
}

// GetCharFilter returns the CharFilter field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetCharFilter() []string {
	if o == nil || IsNil(o.CharFilter) {
		var ret []string
		return ret
	}
	return o.CharFilter
}

// GetCharFilterOk returns a tuple with the CharFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetCharFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.CharFilter) {
		return nil, false
	}
	return o.CharFilter, true
}

// HasCharFilter returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasCharFilter() bool {
	if o != nil && !IsNil(o.CharFilter) {
		return true
	}

	return false
}

// SetCharFilter gets a reference to the given []string and assigns it to the CharFilter field.
func (o *MetaAnalyzer) SetCharFilter(v []string) {
	o.CharFilter = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetFilter() []string {
	if o == nil || IsNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *MetaAnalyzer) SetFilter(v []string) {
	o.Filter = v
}

// GetLowercase returns the Lowercase field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetLowercase() bool {
	if o == nil || IsNil(o.Lowercase) {
		var ret bool
		return ret
	}
	return *o.Lowercase
}

// GetLowercaseOk returns a tuple with the Lowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetLowercaseOk() (*bool, bool) {
	if o == nil || IsNil(o.Lowercase) {
		return nil, false
	}
	return o.Lowercase, true
}

// HasLowercase returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasLowercase() bool {
	if o != nil && !IsNil(o.Lowercase) {
		return true
	}

	return false
}

// SetLowercase gets a reference to the given bool and assigns it to the Lowercase field.
func (o *MetaAnalyzer) SetLowercase(v bool) {
	o.Lowercase = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *MetaAnalyzer) SetPattern(v string) {
	o.Pattern = &v
}

// GetStopwords returns the Stopwords field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetStopwords() []string {
	if o == nil || IsNil(o.Stopwords) {
		var ret []string
		return ret
	}
	return o.Stopwords
}

// GetStopwordsOk returns a tuple with the Stopwords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetStopwordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Stopwords) {
		return nil, false
	}
	return o.Stopwords, true
}

// HasStopwords returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasStopwords() bool {
	if o != nil && !IsNil(o.Stopwords) {
		return true
	}

	return false
}

// SetStopwords gets a reference to the given []string and assigns it to the Stopwords field.
func (o *MetaAnalyzer) SetStopwords(v []string) {
	o.Stopwords = v
}

// GetTokenFilter returns the TokenFilter field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetTokenFilter() []string {
	if o == nil || IsNil(o.TokenFilter) {
		var ret []string
		return ret
	}
	return o.TokenFilter
}

// GetTokenFilterOk returns a tuple with the TokenFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetTokenFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.TokenFilter) {
		return nil, false
	}
	return o.TokenFilter, true
}

// HasTokenFilter returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasTokenFilter() bool {
	if o != nil && !IsNil(o.TokenFilter) {
		return true
	}

	return false
}

// SetTokenFilter gets a reference to the given []string and assigns it to the TokenFilter field.
func (o *MetaAnalyzer) SetTokenFilter(v []string) {
	o.TokenFilter = v
}

// GetTokenizer returns the Tokenizer field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetTokenizer() string {
	if o == nil || IsNil(o.Tokenizer) {
		var ret string
		return ret
	}
	return *o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetTokenizerOk() (*string, bool) {
	if o == nil || IsNil(o.Tokenizer) {
		return nil, false
	}
	return o.Tokenizer, true
}

// HasTokenizer returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasTokenizer() bool {
	if o != nil && !IsNil(o.Tokenizer) {
		return true
	}

	return false
}

// SetTokenizer gets a reference to the given string and assigns it to the Tokenizer field.
func (o *MetaAnalyzer) SetTokenizer(v string) {
	o.Tokenizer = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetaAnalyzer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAnalyzer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetaAnalyzer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetaAnalyzer) SetType(v string) {
	o.Type = &v
}

func (o MetaAnalyzer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaAnalyzer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CharFilter) {
		toSerialize["char_filter"] = o.CharFilter
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Lowercase) {
		toSerialize["lowercase"] = o.Lowercase
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Stopwords) {
		toSerialize["stopwords"] = o.Stopwords
	}
	if !IsNil(o.TokenFilter) {
		toSerialize["token_filter"] = o.TokenFilter
	}
	if !IsNil(o.Tokenizer) {
		toSerialize["tokenizer"] = o.Tokenizer
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMetaAnalyzer struct {
	value *MetaAnalyzer
	isSet bool
}

func (v NullableMetaAnalyzer) Get() *MetaAnalyzer {
	return v.value
}

func (v *NullableMetaAnalyzer) Set(val *MetaAnalyzer) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAnalyzer) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAnalyzer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAnalyzer(val *MetaAnalyzer) *NullableMetaAnalyzer {
	return &NullableMetaAnalyzer{value: val, isSet: true}
}

func (v NullableMetaAnalyzer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAnalyzer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
