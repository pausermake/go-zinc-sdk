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

// checks if the AuthLoginUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthLoginUser{}

// AuthLoginUser struct for AuthLoginUser
type AuthLoginUser struct {
	Id   *string `json:"_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NewAuthLoginUser instantiates a new AuthLoginUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLoginUser() *AuthLoginUser {
	this := AuthLoginUser{}
	return &this
}

// NewAuthLoginUserWithDefaults instantiates a new AuthLoginUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLoginUserWithDefaults() *AuthLoginUser {
	this := AuthLoginUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthLoginUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLoginUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthLoginUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthLoginUser) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthLoginUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLoginUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthLoginUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthLoginUser) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AuthLoginUser) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLoginUser) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AuthLoginUser) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AuthLoginUser) SetRole(v string) {
	o.Role = &v
}

func (o AuthLoginUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthLoginUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableAuthLoginUser struct {
	value *AuthLoginUser
	isSet bool
}

func (v NullableAuthLoginUser) Get() *AuthLoginUser {
	return v.value
}

func (v *NullableAuthLoginUser) Set(val *AuthLoginUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLoginUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLoginUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLoginUser(val *AuthLoginUser) *NullableAuthLoginUser {
	return &NullableAuthLoginUser{value: val, isSet: true}
}

func (v NullableAuthLoginUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLoginUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
