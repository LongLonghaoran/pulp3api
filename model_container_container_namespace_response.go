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
	"time"
)

// ContainerContainerNamespaceResponse Serializer for ContainerNamespaces.
type ContainerContainerNamespaceResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Name string `json:"name"`
}

// NewContainerContainerNamespaceResponse instantiates a new ContainerContainerNamespaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerNamespaceResponse(name string) *ContainerContainerNamespaceResponse {
	this := ContainerContainerNamespaceResponse{}
	this.Name = name
	return &this
}

// NewContainerContainerNamespaceResponseWithDefaults instantiates a new ContainerContainerNamespaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerNamespaceResponseWithDefaults() *ContainerContainerNamespaceResponse {
	this := ContainerContainerNamespaceResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerNamespaceResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerNamespaceResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerNamespaceResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerNamespaceResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerNamespaceResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerNamespaceResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerNamespaceResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerNamespaceResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *ContainerContainerNamespaceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerNamespaceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerNamespaceResponse) SetName(v string) {
	o.Name = v
}

func (o ContainerContainerNamespaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContainerNamespaceResponse struct {
	value *ContainerContainerNamespaceResponse
	isSet bool
}

func (v NullableContainerContainerNamespaceResponse) Get() *ContainerContainerNamespaceResponse {
	return v.value
}

func (v *NullableContainerContainerNamespaceResponse) Set(val *ContainerContainerNamespaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerNamespaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerNamespaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerNamespaceResponse(val *ContainerContainerNamespaceResponse) *NullableContainerContainerNamespaceResponse {
	return &NullableContainerContainerNamespaceResponse{value: val, isSet: true}
}

func (v NullableContainerContainerNamespaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerNamespaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

