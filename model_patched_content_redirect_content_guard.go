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
)

// PatchedContentRedirectContentGuard A serializer for ContentRedirectContentGuard.
type PatchedContentRedirectContentGuard struct {
	// The unique name.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewPatchedContentRedirectContentGuard instantiates a new PatchedContentRedirectContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedContentRedirectContentGuard() *PatchedContentRedirectContentGuard {
	this := PatchedContentRedirectContentGuard{}
	return &this
}

// NewPatchedContentRedirectContentGuardWithDefaults instantiates a new PatchedContentRedirectContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedContentRedirectContentGuardWithDefaults() *PatchedContentRedirectContentGuard {
	this := PatchedContentRedirectContentGuard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedContentRedirectContentGuard) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedContentRedirectContentGuard) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedContentRedirectContentGuard) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedContentRedirectContentGuard) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedContentRedirectContentGuard) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedContentRedirectContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedContentRedirectContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedContentRedirectContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedContentRedirectContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedContentRedirectContentGuard) UnsetDescription() {
	o.Description.Unset()
}

func (o PatchedContentRedirectContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedContentRedirectContentGuard struct {
	value *PatchedContentRedirectContentGuard
	isSet bool
}

func (v NullablePatchedContentRedirectContentGuard) Get() *PatchedContentRedirectContentGuard {
	return v.value
}

func (v *NullablePatchedContentRedirectContentGuard) Set(val *PatchedContentRedirectContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedContentRedirectContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedContentRedirectContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedContentRedirectContentGuard(val *PatchedContentRedirectContentGuard) *NullablePatchedContentRedirectContentGuard {
	return &NullablePatchedContentRedirectContentGuard{value: val, isSet: true}
}

func (v NullablePatchedContentRedirectContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedContentRedirectContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


