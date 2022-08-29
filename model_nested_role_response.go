/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NestedRoleResponse Serializer to add/remove object roles to/from users/groups.  This is used in conjunction with ``pulpcore.app.viewsets.base.RolesMixin`` and requires the underlying object to be passed as ``content_object`` in the context.
type NestedRoleResponse struct {
	Users []string `json:"users,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Role string `json:"role"`
}

// NewNestedRoleResponse instantiates a new NestedRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRoleResponse(role string) *NestedRoleResponse {
	this := NestedRoleResponse{}
	this.Role = role
	return &this
}

// NewNestedRoleResponseWithDefaults instantiates a new NestedRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRoleResponseWithDefaults() *NestedRoleResponse {
	this := NestedRoleResponse{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NestedRoleResponse) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NestedRoleResponse) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *NestedRoleResponse) SetUsers(v []string) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *NestedRoleResponse) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *NestedRoleResponse) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *NestedRoleResponse) SetGroups(v []string) {
	o.Groups = v
}

// GetRole returns the Role field value
func (o *NestedRoleResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *NestedRoleResponse) SetRole(v string) {
	o.Role = v
}

func (o NestedRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableNestedRoleResponse struct {
	value *NestedRoleResponse
	isSet bool
}

func (v NullableNestedRoleResponse) Get() *NestedRoleResponse {
	return v.value
}

func (v *NullableNestedRoleResponse) Set(val *NestedRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRoleResponse(val *NestedRoleResponse) *NullableNestedRoleResponse {
	return &NullableNestedRoleResponse{value: val, isSet: true}
}

func (v NullableNestedRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


