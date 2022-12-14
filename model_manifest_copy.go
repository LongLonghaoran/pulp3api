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

// ManifestCopy Serializer for copying manifests from a source repository to a destination repository.
type ManifestCopy struct {
	// A URI of the repository to copy content from.
	SourceRepository *string `json:"source_repository,omitempty"`
	// A URI of the repository version to copy content from.
	SourceRepositoryVersion *string `json:"source_repository_version,omitempty"`
	// A list of manifest digests to copy.
	Digests []interface{} `json:"digests,omitempty"`
	// A list of media_types to copy.
	MediaTypes []MediaTypesEnum `json:"media_types,omitempty"`
}

// NewManifestCopy instantiates a new ManifestCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManifestCopy() *ManifestCopy {
	this := ManifestCopy{}
	return &this
}

// NewManifestCopyWithDefaults instantiates a new ManifestCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManifestCopyWithDefaults() *ManifestCopy {
	this := ManifestCopy{}
	return &this
}

// GetSourceRepository returns the SourceRepository field value if set, zero value otherwise.
func (o *ManifestCopy) GetSourceRepository() string {
	if o == nil || o.SourceRepository == nil {
		var ret string
		return ret
	}
	return *o.SourceRepository
}

// GetSourceRepositoryOk returns a tuple with the SourceRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestCopy) GetSourceRepositoryOk() (*string, bool) {
	if o == nil || o.SourceRepository == nil {
		return nil, false
	}
	return o.SourceRepository, true
}

// HasSourceRepository returns a boolean if a field has been set.
func (o *ManifestCopy) HasSourceRepository() bool {
	if o != nil && o.SourceRepository != nil {
		return true
	}

	return false
}

// SetSourceRepository gets a reference to the given string and assigns it to the SourceRepository field.
func (o *ManifestCopy) SetSourceRepository(v string) {
	o.SourceRepository = &v
}

// GetSourceRepositoryVersion returns the SourceRepositoryVersion field value if set, zero value otherwise.
func (o *ManifestCopy) GetSourceRepositoryVersion() string {
	if o == nil || o.SourceRepositoryVersion == nil {
		var ret string
		return ret
	}
	return *o.SourceRepositoryVersion
}

// GetSourceRepositoryVersionOk returns a tuple with the SourceRepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestCopy) GetSourceRepositoryVersionOk() (*string, bool) {
	if o == nil || o.SourceRepositoryVersion == nil {
		return nil, false
	}
	return o.SourceRepositoryVersion, true
}

// HasSourceRepositoryVersion returns a boolean if a field has been set.
func (o *ManifestCopy) HasSourceRepositoryVersion() bool {
	if o != nil && o.SourceRepositoryVersion != nil {
		return true
	}

	return false
}

// SetSourceRepositoryVersion gets a reference to the given string and assigns it to the SourceRepositoryVersion field.
func (o *ManifestCopy) SetSourceRepositoryVersion(v string) {
	o.SourceRepositoryVersion = &v
}

// GetDigests returns the Digests field value if set, zero value otherwise.
func (o *ManifestCopy) GetDigests() []interface{} {
	if o == nil || o.Digests == nil {
		var ret []interface{}
		return ret
	}
	return o.Digests
}

// GetDigestsOk returns a tuple with the Digests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestCopy) GetDigestsOk() ([]interface{}, bool) {
	if o == nil || o.Digests == nil {
		return nil, false
	}
	return o.Digests, true
}

// HasDigests returns a boolean if a field has been set.
func (o *ManifestCopy) HasDigests() bool {
	if o != nil && o.Digests != nil {
		return true
	}

	return false
}

// SetDigests gets a reference to the given []interface{} and assigns it to the Digests field.
func (o *ManifestCopy) SetDigests(v []interface{}) {
	o.Digests = v
}

// GetMediaTypes returns the MediaTypes field value if set, zero value otherwise.
func (o *ManifestCopy) GetMediaTypes() []MediaTypesEnum {
	if o == nil || o.MediaTypes == nil {
		var ret []MediaTypesEnum
		return ret
	}
	return o.MediaTypes
}

// GetMediaTypesOk returns a tuple with the MediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestCopy) GetMediaTypesOk() ([]MediaTypesEnum, bool) {
	if o == nil || o.MediaTypes == nil {
		return nil, false
	}
	return o.MediaTypes, true
}

// HasMediaTypes returns a boolean if a field has been set.
func (o *ManifestCopy) HasMediaTypes() bool {
	if o != nil && o.MediaTypes != nil {
		return true
	}

	return false
}

// SetMediaTypes gets a reference to the given []MediaTypesEnum and assigns it to the MediaTypes field.
func (o *ManifestCopy) SetMediaTypes(v []MediaTypesEnum) {
	o.MediaTypes = v
}

func (o ManifestCopy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceRepository != nil {
		toSerialize["source_repository"] = o.SourceRepository
	}
	if o.SourceRepositoryVersion != nil {
		toSerialize["source_repository_version"] = o.SourceRepositoryVersion
	}
	if o.Digests != nil {
		toSerialize["digests"] = o.Digests
	}
	if o.MediaTypes != nil {
		toSerialize["media_types"] = o.MediaTypes
	}
	return json.Marshal(toSerialize)
}

type NullableManifestCopy struct {
	value *ManifestCopy
	isSet bool
}

func (v NullableManifestCopy) Get() *ManifestCopy {
	return v.value
}

func (v *NullableManifestCopy) Set(val *ManifestCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableManifestCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableManifestCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManifestCopy(val *ManifestCopy) *NullableManifestCopy {
	return &NullableManifestCopy{value: val, isSet: true}
}

func (v NullableManifestCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManifestCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


