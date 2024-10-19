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

// checks if the MetaHit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaHit{}

// MetaHit struct for MetaHit
type MetaHit struct {
	Timestamp *string                `json:"@timestamp,omitempty"`
	Id        *string                `json:"_id,omitempty"`
	Index     *string                `json:"_index,omitempty"`
	Score     *float32               `json:"_score,omitempty"`
	Source    map[string]interface{} `json:"_source,omitempty"`
	Type      *string                `json:"_type,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Highlight map[string]interface{} `json:"highlight,omitempty"`
}

// NewMetaHit instantiates a new MetaHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaHit() *MetaHit {
	this := MetaHit{}
	return &this
}

// NewMetaHitWithDefaults instantiates a new MetaHit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaHitWithDefaults() *MetaHit {
	this := MetaHit{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MetaHit) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetaHit) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *MetaHit) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetaHit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetaHit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetaHit) SetId(v string) {
	o.Id = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MetaHit) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MetaHit) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *MetaHit) SetIndex(v string) {
	o.Index = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *MetaHit) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *MetaHit) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *MetaHit) SetScore(v float32) {
	o.Score = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MetaHit) GetSource() map[string]interface{} {
	if o == nil || IsNil(o.Source) {
		var ret map[string]interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetSourceOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Source) {
		return map[string]interface{}{}, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MetaHit) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given map[string]interface{} and assigns it to the Source field.
func (o *MetaHit) SetSource(v map[string]interface{}) {
	o.Source = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetaHit) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetaHit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetaHit) SetType(v string) {
	o.Type = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaHit) GetFields() map[string]interface{} {
	if o == nil || IsNil(o.Fields) {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Fields) {
		return map[string]interface{}{}, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaHit) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *MetaHit) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *MetaHit) GetHighlight() map[string]interface{} {
	if o == nil || IsNil(o.Highlight) {
		var ret map[string]interface{}
		return ret
	}
	return o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHit) GetHighlightOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Highlight) {
		return map[string]interface{}{}, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *MetaHit) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given map[string]interface{} and assigns it to the Highlight field.
func (o *MetaHit) SetHighlight(v map[string]interface{}) {
	o.Highlight = v
}

func (o MetaHit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaHit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["@timestamp"] = o.Timestamp
	}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Index) {
		toSerialize["_index"] = o.Index
	}
	if !IsNil(o.Score) {
		toSerialize["_score"] = o.Score
	}
	if !IsNil(o.Source) {
		toSerialize["_source"] = o.Source
	}
	if !IsNil(o.Type) {
		toSerialize["_type"] = o.Type
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	return toSerialize, nil
}

type NullableMetaHit struct {
	value *MetaHit
	isSet bool
}

func (v NullableMetaHit) Get() *MetaHit {
	return v.value
}

func (v *NullableMetaHit) Set(val *MetaHit) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaHit) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaHit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaHit(val *MetaHit) *NullableMetaHit {
	return &NullableMetaHit{value: val, isSet: true}
}

func (v NullableMetaHit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaHit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
