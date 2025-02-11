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

// checks if the MetaTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaTemplate{}

// MetaTemplate struct for MetaTemplate
type MetaTemplate struct {
	IndexTemplate *MetaIndexTemplate `json:"index_template,omitempty"`
	Name          *string            `json:"name,omitempty"`
}

// NewMetaTemplate instantiates a new MetaTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaTemplate() *MetaTemplate {
	this := MetaTemplate{}
	return &this
}

// NewMetaTemplateWithDefaults instantiates a new MetaTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaTemplateWithDefaults() *MetaTemplate {
	this := MetaTemplate{}
	return &this
}

// GetIndexTemplate returns the IndexTemplate field value if set, zero value otherwise.
func (o *MetaTemplate) GetIndexTemplate() MetaIndexTemplate {
	if o == nil || IsNil(o.IndexTemplate) {
		var ret MetaIndexTemplate
		return ret
	}
	return *o.IndexTemplate
}

// GetIndexTemplateOk returns a tuple with the IndexTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaTemplate) GetIndexTemplateOk() (*MetaIndexTemplate, bool) {
	if o == nil || IsNil(o.IndexTemplate) {
		return nil, false
	}
	return o.IndexTemplate, true
}

// HasIndexTemplate returns a boolean if a field has been set.
func (o *MetaTemplate) HasIndexTemplate() bool {
	if o != nil && !IsNil(o.IndexTemplate) {
		return true
	}

	return false
}

// SetIndexTemplate gets a reference to the given MetaIndexTemplate and assigns it to the IndexTemplate field.
func (o *MetaTemplate) SetIndexTemplate(v MetaIndexTemplate) {
	o.IndexTemplate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaTemplate) SetName(v string) {
	o.Name = &v
}

func (o MetaTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndexTemplate) {
		toSerialize["index_template"] = o.IndexTemplate
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMetaTemplate struct {
	value *MetaTemplate
	isSet bool
}

func (v NullableMetaTemplate) Get() *MetaTemplate {
	return v.value
}

func (v *NullableMetaTemplate) Set(val *MetaTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaTemplate(val *MetaTemplate) *NullableMetaTemplate {
	return &NullableMetaTemplate{value: val, isSet: true}
}

func (v NullableMetaTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
