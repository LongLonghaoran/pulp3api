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

// ContentGuard Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type ContentGuard struct {
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewContentGuard instantiates a new ContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentGuard(name string) *ContentGuard {
	this := ContentGuard{}
	this.Name = name
	return &this
}

// NewContentGuardWithDefaults instantiates a new ContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentGuardWithDefaults() *ContentGuard {
	this := ContentGuard{}
	return &this
}

// GetName returns the Name field value
func (o *ContentGuard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentGuard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentGuard) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentGuard) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContentGuard) UnsetDescription() {
	o.Description.Unset()
}

func (o ContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContentGuard struct {
	value *ContentGuard
	isSet bool
}

func (v NullableContentGuard) Get() *ContentGuard {
	return v.value
}

func (v *NullableContentGuard) Set(val *ContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentGuard(val *ContentGuard) *NullableContentGuard {
	return &NullableContentGuard{value: val, isSet: true}
}

func (v NullableContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


