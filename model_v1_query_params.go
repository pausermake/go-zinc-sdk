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

// checks if the V1QueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1QueryParams{}

// V1QueryParams struct for V1QueryParams
type V1QueryParams struct {
	Boost     *int32  `json:"boost,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	Field     *string `json:"field,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
	Term      *string `json:"term,omitempty"`
	// For multi phrase query
	Terms [][]string `json:"terms,omitempty"`
}

// NewV1QueryParams instantiates a new V1QueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1QueryParams() *V1QueryParams {
	this := V1QueryParams{}
	return &this
}

// NewV1QueryParamsWithDefaults instantiates a new V1QueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1QueryParamsWithDefaults() *V1QueryParams {
	this := V1QueryParams{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *V1QueryParams) GetBoost() int32 {
	if o == nil || IsNil(o.Boost) {
		var ret int32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetBoostOk() (*int32, bool) {
	if o == nil || IsNil(o.Boost) {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *V1QueryParams) HasBoost() bool {
	if o != nil && !IsNil(o.Boost) {
		return true
	}

	return false
}

// SetBoost gets a reference to the given int32 and assigns it to the Boost field.
func (o *V1QueryParams) SetBoost(v int32) {
	o.Boost = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V1QueryParams) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V1QueryParams) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *V1QueryParams) SetEndTime(v string) {
	o.EndTime = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *V1QueryParams) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *V1QueryParams) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *V1QueryParams) SetField(v string) {
	o.Field = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V1QueryParams) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V1QueryParams) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *V1QueryParams) SetStartTime(v string) {
	o.StartTime = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *V1QueryParams) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *V1QueryParams) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *V1QueryParams) SetTerm(v string) {
	o.Term = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *V1QueryParams) GetTerms() [][]string {
	if o == nil || IsNil(o.Terms) {
		var ret [][]string
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1QueryParams) GetTermsOk() ([][]string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *V1QueryParams) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given [][]string and assigns it to the Terms field.
func (o *V1QueryParams) SetTerms(v [][]string) {
	o.Terms = v
}

func (o V1QueryParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1QueryParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Boost) {
		toSerialize["boost"] = o.Boost
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	return toSerialize, nil
}

type NullableV1QueryParams struct {
	value *V1QueryParams
	isSet bool
}

func (v NullableV1QueryParams) Get() *V1QueryParams {
	return v.value
}

func (v *NullableV1QueryParams) Set(val *V1QueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV1QueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV1QueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1QueryParams(val *V1QueryParams) *NullableV1QueryParams {
	return &NullableV1QueryParams{value: val, isSet: true}
}

func (v NullableV1QueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1QueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
