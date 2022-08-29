/*
 * Pulp 3 API
 *
 * Fetch, Upload, Organize, and Distribute Software Packages
 *
 * API version: v3
 * Contact: pulp-list@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// GroupRoleResponse Serializer for GroupRole.
type GroupRoleResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Role string `json:"role"`
	// pulp_href of the object for which role permissions should be asserted. If set to 'null', permissions will act on the model-level.
	ContentObject NullableString `json:"content_object"`
	Description *string `json:"description,omitempty"`
	Permissions *[]string `json:"permissions,omitempty"`
}

// NewGroupRoleResponse instantiates a new GroupRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRoleResponse(role string, contentObject NullableString, ) *GroupRoleResponse {
	this := GroupRoleResponse{}
	this.Role = role
	this.ContentObject = contentObject
	return &this
}

// NewGroupRoleResponseWithDefaults instantiates a new GroupRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRoleResponseWithDefaults() *GroupRoleResponse {
	this := GroupRoleResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *GroupRoleResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *GroupRoleResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *GroupRoleResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *GroupRoleResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *GroupRoleResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *GroupRoleResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetRole returns the Role field value
func (o *GroupRoleResponse) GetRole() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *GroupRoleResponse) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *GroupRoleResponse) SetRole(v string) {
	o.Role = v
}

// GetContentObject returns the ContentObject field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupRoleResponse) GetContentObject() string {
	if o == nil || o.ContentObject.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContentObject.Get()
}

// GetContentObjectOk returns a tuple with the ContentObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRoleResponse) GetContentObjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentObject.Get(), o.ContentObject.IsSet()
}

// SetContentObject sets field value
func (o *GroupRoleResponse) SetContentObject(v string) {
	o.ContentObject.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupRoleResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupRoleResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupRoleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupRoleResponse) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRoleResponse) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupRoleResponse) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GroupRoleResponse) SetPermissions(v []string) {
	o.Permissions = &v
}

func (o GroupRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["content_object"] = o.ContentObject.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableGroupRoleResponse struct {
	value *GroupRoleResponse
	isSet bool
}

func (v NullableGroupRoleResponse) Get() *GroupRoleResponse {
	return v.value
}

func (v *NullableGroupRoleResponse) Set(val *GroupRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRoleResponse(val *GroupRoleResponse) *NullableGroupRoleResponse {
	return &NullableGroupRoleResponse{value: val, isSet: true}
}

func (v NullableGroupRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


