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

// RepositoryAddRemoveContent Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type RepositoryAddRemoveContent struct {
	// A list of content units to add to a new repository version. This content is added after remove_content_units are removed.
	AddContentUnits []interface{} `json:"add_content_units,omitempty"`
	// A list of content units to remove from the latest repository version. You may also specify '*' as an entry to remove all content. This content is removed before add_content_units are added.
	RemoveContentUnits []interface{} `json:"remove_content_units,omitempty"`
	// A repository version whose content will be used as the initial set of content for the new repository version
	BaseVersion *string `json:"base_version,omitempty"`
}

// NewRepositoryAddRemoveContent instantiates a new RepositoryAddRemoveContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryAddRemoveContent() *RepositoryAddRemoveContent {
	this := RepositoryAddRemoveContent{}
	return &this
}

// NewRepositoryAddRemoveContentWithDefaults instantiates a new RepositoryAddRemoveContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryAddRemoveContentWithDefaults() *RepositoryAddRemoveContent {
	this := RepositoryAddRemoveContent{}
	return &this
}

// GetAddContentUnits returns the AddContentUnits field value if set, zero value otherwise.
func (o *RepositoryAddRemoveContent) GetAddContentUnits() []interface{} {
	if o == nil || o.AddContentUnits == nil {
		var ret []interface{}
		return ret
	}
	return o.AddContentUnits
}

// GetAddContentUnitsOk returns a tuple with the AddContentUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAddRemoveContent) GetAddContentUnitsOk() ([]interface{}, bool) {
	if o == nil || o.AddContentUnits == nil {
		return nil, false
	}
	return o.AddContentUnits, true
}

// HasAddContentUnits returns a boolean if a field has been set.
func (o *RepositoryAddRemoveContent) HasAddContentUnits() bool {
	if o != nil && o.AddContentUnits != nil {
		return true
	}

	return false
}

// SetAddContentUnits gets a reference to the given []interface{} and assigns it to the AddContentUnits field.
func (o *RepositoryAddRemoveContent) SetAddContentUnits(v []interface{}) {
	o.AddContentUnits = v
}

// GetRemoveContentUnits returns the RemoveContentUnits field value if set, zero value otherwise.
func (o *RepositoryAddRemoveContent) GetRemoveContentUnits() []interface{} {
	if o == nil || o.RemoveContentUnits == nil {
		var ret []interface{}
		return ret
	}
	return o.RemoveContentUnits
}

// GetRemoveContentUnitsOk returns a tuple with the RemoveContentUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAddRemoveContent) GetRemoveContentUnitsOk() ([]interface{}, bool) {
	if o == nil || o.RemoveContentUnits == nil {
		return nil, false
	}
	return o.RemoveContentUnits, true
}

// HasRemoveContentUnits returns a boolean if a field has been set.
func (o *RepositoryAddRemoveContent) HasRemoveContentUnits() bool {
	if o != nil && o.RemoveContentUnits != nil {
		return true
	}

	return false
}

// SetRemoveContentUnits gets a reference to the given []interface{} and assigns it to the RemoveContentUnits field.
func (o *RepositoryAddRemoveContent) SetRemoveContentUnits(v []interface{}) {
	o.RemoveContentUnits = v
}

// GetBaseVersion returns the BaseVersion field value if set, zero value otherwise.
func (o *RepositoryAddRemoveContent) GetBaseVersion() string {
	if o == nil || o.BaseVersion == nil {
		var ret string
		return ret
	}
	return *o.BaseVersion
}

// GetBaseVersionOk returns a tuple with the BaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAddRemoveContent) GetBaseVersionOk() (*string, bool) {
	if o == nil || o.BaseVersion == nil {
		return nil, false
	}
	return o.BaseVersion, true
}

// HasBaseVersion returns a boolean if a field has been set.
func (o *RepositoryAddRemoveContent) HasBaseVersion() bool {
	if o != nil && o.BaseVersion != nil {
		return true
	}

	return false
}

// SetBaseVersion gets a reference to the given string and assigns it to the BaseVersion field.
func (o *RepositoryAddRemoveContent) SetBaseVersion(v string) {
	o.BaseVersion = &v
}

func (o RepositoryAddRemoveContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddContentUnits != nil {
		toSerialize["add_content_units"] = o.AddContentUnits
	}
	if o.RemoveContentUnits != nil {
		toSerialize["remove_content_units"] = o.RemoveContentUnits
	}
	if o.BaseVersion != nil {
		toSerialize["base_version"] = o.BaseVersion
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryAddRemoveContent struct {
	value *RepositoryAddRemoveContent
	isSet bool
}

func (v NullableRepositoryAddRemoveContent) Get() *RepositoryAddRemoveContent {
	return v.value
}

func (v *NullableRepositoryAddRemoveContent) Set(val *RepositoryAddRemoveContent) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryAddRemoveContent) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryAddRemoveContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryAddRemoveContent(val *RepositoryAddRemoveContent) *NullableRepositoryAddRemoveContent {
	return &NullableRepositoryAddRemoveContent{value: val, isSet: true}
}

func (v NullableRepositoryAddRemoveContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryAddRemoveContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


