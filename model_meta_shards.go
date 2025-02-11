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

// checks if the MetaShards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaShards{}

// MetaShards struct for MetaShards
type MetaShards struct {
	Failed     *int32 `json:"failed,omitempty"`
	Skipped    *int32 `json:"skipped,omitempty"`
	Successful *int32 `json:"successful,omitempty"`
	Total      *int32 `json:"total,omitempty"`
}

// NewMetaShards instantiates a new MetaShards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaShards() *MetaShards {
	this := MetaShards{}
	return &this
}

// NewMetaShardsWithDefaults instantiates a new MetaShards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaShardsWithDefaults() *MetaShards {
	this := MetaShards{}
	return &this
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *MetaShards) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaShards) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *MetaShards) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *MetaShards) SetFailed(v int32) {
	o.Failed = &v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *MetaShards) GetSkipped() int32 {
	if o == nil || IsNil(o.Skipped) {
		var ret int32
		return ret
	}
	return *o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaShards) GetSkippedOk() (*int32, bool) {
	if o == nil || IsNil(o.Skipped) {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *MetaShards) HasSkipped() bool {
	if o != nil && !IsNil(o.Skipped) {
		return true
	}

	return false
}

// SetSkipped gets a reference to the given int32 and assigns it to the Skipped field.
func (o *MetaShards) SetSkipped(v int32) {
	o.Skipped = &v
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *MetaShards) GetSuccessful() int32 {
	if o == nil || IsNil(o.Successful) {
		var ret int32
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaShards) GetSuccessfulOk() (*int32, bool) {
	if o == nil || IsNil(o.Successful) {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *MetaShards) HasSuccessful() bool {
	if o != nil && !IsNil(o.Successful) {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given int32 and assigns it to the Successful field.
func (o *MetaShards) SetSuccessful(v int32) {
	o.Successful = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MetaShards) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaShards) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MetaShards) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MetaShards) SetTotal(v int32) {
	o.Total = &v
}

func (o MetaShards) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaShards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Skipped) {
		toSerialize["skipped"] = o.Skipped
	}
	if !IsNil(o.Successful) {
		toSerialize["successful"] = o.Successful
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMetaShards struct {
	value *MetaShards
	isSet bool
}

func (v NullableMetaShards) Get() *MetaShards {
	return v.value
}

func (v *NullableMetaShards) Set(val *MetaShards) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaShards) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaShards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaShards(val *MetaShards) *NullableMetaShards {
	return &NullableMetaShards{value: val, isSet: true}
}

func (v NullableMetaShards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaShards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
