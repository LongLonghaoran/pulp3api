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

// ContainerTagResponse Serializer for Tags.
type ContainerTagResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Tag name
	Name string `json:"name"`
	// Manifest that is tagged
	TaggedManifest string `json:"tagged_manifest"`
}

// NewContainerTagResponse instantiates a new ContainerTagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTagResponse(name string, taggedManifest string) *ContainerTagResponse {
	this := ContainerTagResponse{}
	this.Name = name
	this.TaggedManifest = taggedManifest
	return &this
}

// NewContainerTagResponseWithDefaults instantiates a new ContainerTagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTagResponseWithDefaults() *ContainerTagResponse {
	this := ContainerTagResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerTagResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTagResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerTagResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerTagResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerTagResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTagResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerTagResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerTagResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *ContainerTagResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerTagResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerTagResponse) SetName(v string) {
	o.Name = v
}

// GetTaggedManifest returns the TaggedManifest field value
func (o *ContainerTagResponse) GetTaggedManifest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaggedManifest
}

// GetTaggedManifestOk returns a tuple with the TaggedManifest field value
// and a boolean to check if the value has been set.
func (o *ContainerTagResponse) GetTaggedManifestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaggedManifest, true
}

// SetTaggedManifest sets field value
func (o *ContainerTagResponse) SetTaggedManifest(v string) {
	o.TaggedManifest = v
}

func (o ContainerTagResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["tagged_manifest"] = o.TaggedManifest
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTagResponse struct {
	value *ContainerTagResponse
	isSet bool
}

func (v NullableContainerTagResponse) Get() *ContainerTagResponse {
	return v.value
}

func (v *NullableContainerTagResponse) Set(val *ContainerTagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTagResponse(val *ContainerTagResponse) *NullableContainerTagResponse {
	return &NullableContainerTagResponse{value: val, isSet: true}
}

func (v NullableContainerTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


