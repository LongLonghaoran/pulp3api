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

// GroupResponse Serializer for Group.
type GroupResponse struct {
	// Name
	Name string `json:"name"`
	PulpHref *string `json:"pulp_href,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewGroupResponse instantiates a new GroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResponse(name string) *GroupResponse {
	this := GroupResponse{}
	this.Name = name
	return &this
}

// NewGroupResponseWithDefaults instantiates a new GroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResponseWithDefaults() *GroupResponse {
	this := GroupResponse{}
	return &this
}

// GetName returns the Name field value
func (o *GroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupResponse) SetName(v string) {
	o.Name = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *GroupResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *GroupResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *GroupResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GroupResponse) SetId(v int32) {
	o.Id = &v
}

func (o GroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGroupResponse struct {
	value *GroupResponse
	isSet bool
}

func (v NullableGroupResponse) Get() *GroupResponse {
	return v.value
}

func (v *NullableGroupResponse) Set(val *GroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResponse(val *GroupResponse) *NullableGroupResponse {
	return &NullableGroupResponse{value: val, isSet: true}
}

func (v NullableGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


