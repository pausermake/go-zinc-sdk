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

// checks if the V1ZincQueryForSDK type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ZincQueryForSDK{}

// V1ZincQueryForSDK struct for V1ZincQueryForSDK
type V1ZincQueryForSDK struct {
	Source     []string                        `json:"_source,omitempty"`
	Aggs       *map[string]V1AggregationParams `json:"aggs,omitempty"`
	Explain    *bool                           `json:"explain,omitempty"`
	From       *int32                          `json:"from,omitempty"`
	Highlight  *MetaHighlight                  `json:"highlight,omitempty"`
	MaxResults *int32                          `json:"max_results,omitempty"`
	Query      *V1QueryParams                  `json:"query,omitempty"`
	// SearchType is the type of search to perform. Can be match, match_phrase, query_string, etc
	SearchType *string  `json:"search_type,omitempty"`
	SortFields []string `json:"sort_fields,omitempty"`
}

// NewV1ZincQueryForSDK instantiates a new V1ZincQueryForSDK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ZincQueryForSDK() *V1ZincQueryForSDK {
	this := V1ZincQueryForSDK{}
	return &this
}

// NewV1ZincQueryForSDKWithDefaults instantiates a new V1ZincQueryForSDK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ZincQueryForSDKWithDefaults() *V1ZincQueryForSDK {
	this := V1ZincQueryForSDK{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetSource() []string {
	if o == nil || IsNil(o.Source) {
		var ret []string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetSourceOk() ([]string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given []string and assigns it to the Source field.
func (o *V1ZincQueryForSDK) SetSource(v []string) {
	o.Source = v
}

// GetAggs returns the Aggs field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetAggs() map[string]V1AggregationParams {
	if o == nil || IsNil(o.Aggs) {
		var ret map[string]V1AggregationParams
		return ret
	}
	return *o.Aggs
}

// GetAggsOk returns a tuple with the Aggs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetAggsOk() (*map[string]V1AggregationParams, bool) {
	if o == nil || IsNil(o.Aggs) {
		return nil, false
	}
	return o.Aggs, true
}

// HasAggs returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasAggs() bool {
	if o != nil && !IsNil(o.Aggs) {
		return true
	}

	return false
}

// SetAggs gets a reference to the given map[string]V1AggregationParams and assigns it to the Aggs field.
func (o *V1ZincQueryForSDK) SetAggs(v map[string]V1AggregationParams) {
	o.Aggs = &v
}

// GetExplain returns the Explain field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetExplain() bool {
	if o == nil || IsNil(o.Explain) {
		var ret bool
		return ret
	}
	return *o.Explain
}

// GetExplainOk returns a tuple with the Explain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetExplainOk() (*bool, bool) {
	if o == nil || IsNil(o.Explain) {
		return nil, false
	}
	return o.Explain, true
}

// HasExplain returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasExplain() bool {
	if o != nil && !IsNil(o.Explain) {
		return true
	}

	return false
}

// SetExplain gets a reference to the given bool and assigns it to the Explain field.
func (o *V1ZincQueryForSDK) SetExplain(v bool) {
	o.Explain = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *V1ZincQueryForSDK) SetFrom(v int32) {
	o.From = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetHighlight() MetaHighlight {
	if o == nil || IsNil(o.Highlight) {
		var ret MetaHighlight
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetHighlightOk() (*MetaHighlight, bool) {
	if o == nil || IsNil(o.Highlight) {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given MetaHighlight and assigns it to the Highlight field.
func (o *V1ZincQueryForSDK) SetHighlight(v MetaHighlight) {
	o.Highlight = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *V1ZincQueryForSDK) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetQuery() V1QueryParams {
	if o == nil || IsNil(o.Query) {
		var ret V1QueryParams
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetQueryOk() (*V1QueryParams, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given V1QueryParams and assigns it to the Query field.
func (o *V1ZincQueryForSDK) SetQuery(v V1QueryParams) {
	o.Query = &v
}

// GetSearchType returns the SearchType field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetSearchType() string {
	if o == nil || IsNil(o.SearchType) {
		var ret string
		return ret
	}
	return *o.SearchType
}

// GetSearchTypeOk returns a tuple with the SearchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetSearchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SearchType) {
		return nil, false
	}
	return o.SearchType, true
}

// HasSearchType returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasSearchType() bool {
	if o != nil && !IsNil(o.SearchType) {
		return true
	}

	return false
}

// SetSearchType gets a reference to the given string and assigns it to the SearchType field.
func (o *V1ZincQueryForSDK) SetSearchType(v string) {
	o.SearchType = &v
}

// GetSortFields returns the SortFields field value if set, zero value otherwise.
func (o *V1ZincQueryForSDK) GetSortFields() []string {
	if o == nil || IsNil(o.SortFields) {
		var ret []string
		return ret
	}
	return o.SortFields
}

// GetSortFieldsOk returns a tuple with the SortFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ZincQueryForSDK) GetSortFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.SortFields) {
		return nil, false
	}
	return o.SortFields, true
}

// HasSortFields returns a boolean if a field has been set.
func (o *V1ZincQueryForSDK) HasSortFields() bool {
	if o != nil && !IsNil(o.SortFields) {
		return true
	}

	return false
}

// SetSortFields gets a reference to the given []string and assigns it to the SortFields field.
func (o *V1ZincQueryForSDK) SetSortFields(v []string) {
	o.SortFields = v
}

func (o V1ZincQueryForSDK) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ZincQueryForSDK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["_source"] = o.Source
	}
	if !IsNil(o.Aggs) {
		toSerialize["aggs"] = o.Aggs
	}
	if !IsNil(o.Explain) {
		toSerialize["explain"] = o.Explain
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	if !IsNil(o.MaxResults) {
		toSerialize["max_results"] = o.MaxResults
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.SearchType) {
		toSerialize["search_type"] = o.SearchType
	}
	if !IsNil(o.SortFields) {
		toSerialize["sort_fields"] = o.SortFields
	}
	return toSerialize, nil
}

type NullableV1ZincQueryForSDK struct {
	value *V1ZincQueryForSDK
	isSet bool
}

func (v NullableV1ZincQueryForSDK) Get() *V1ZincQueryForSDK {
	return v.value
}

func (v *NullableV1ZincQueryForSDK) Set(val *V1ZincQueryForSDK) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ZincQueryForSDK) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ZincQueryForSDK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ZincQueryForSDK(val *V1ZincQueryForSDK) *NullableV1ZincQueryForSDK {
	return &NullableV1ZincQueryForSDK{value: val, isSet: true}
}

func (v NullableV1ZincQueryForSDK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ZincQueryForSDK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
