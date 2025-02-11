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

// checks if the MetaUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaUser{}

// MetaUser struct for MetaUser
type MetaUser struct {
	Id        *string `json:"_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name      *string `json:"name,omitempty"`
	Password  *string `json:"password,omitempty"`
	Role      *string `json:"role,omitempty"`
	Salt      *string `json:"salt,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewMetaUser instantiates a new MetaUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaUser() *MetaUser {
	this := MetaUser{}
	return &this
}

// NewMetaUserWithDefaults instantiates a new MetaUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaUserWithDefaults() *MetaUser {
	this := MetaUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetaUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetaUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetaUser) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetaUser) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetaUser) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MetaUser) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaUser) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MetaUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MetaUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MetaUser) SetPassword(v string) {
	o.Password = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MetaUser) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MetaUser) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MetaUser) SetRole(v string) {
	o.Role = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *MetaUser) GetSalt() string {
	if o == nil || IsNil(o.Salt) {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetSaltOk() (*string, bool) {
	if o == nil || IsNil(o.Salt) {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *MetaUser) HasSalt() bool {
	if o != nil && !IsNil(o.Salt) {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *MetaUser) SetSalt(v string) {
	o.Salt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MetaUser) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MetaUser) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MetaUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o MetaUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Salt) {
		toSerialize["salt"] = o.Salt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableMetaUser struct {
	value *MetaUser
	isSet bool
}

func (v NullableMetaUser) Get() *MetaUser {
	return v.value
}

func (v *NullableMetaUser) Set(val *MetaUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaUser(val *MetaUser) *NullableMetaUser {
	return &NullableMetaUser{value: val, isSet: true}
}

func (v NullableMetaUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
