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

// PatchedPulpImporter Serializer for PulpImporters.
type PatchedPulpImporter struct {
	// Unique name of the Importer.
	Name *string `json:"name,omitempty"`
	// Mapping of repo names in an export file to the repo names in Pulp. For example, if the export has a repo named 'foo' and the repo to import content into was 'bar', the mapping would be \"{'foo': 'bar'}\".
	RepoMapping *map[string]string `json:"repo_mapping,omitempty"`
}

// NewPatchedPulpImporter instantiates a new PatchedPulpImporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPulpImporter() *PatchedPulpImporter {
	this := PatchedPulpImporter{}
	return &this
}

// NewPatchedPulpImporterWithDefaults instantiates a new PatchedPulpImporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPulpImporterWithDefaults() *PatchedPulpImporter {
	this := PatchedPulpImporter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPulpImporter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPulpImporter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPulpImporter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPulpImporter) SetName(v string) {
	o.Name = &v
}

// GetRepoMapping returns the RepoMapping field value if set, zero value otherwise.
func (o *PatchedPulpImporter) GetRepoMapping() map[string]string {
	if o == nil || o.RepoMapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.RepoMapping
}

// GetRepoMappingOk returns a tuple with the RepoMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPulpImporter) GetRepoMappingOk() (*map[string]string, bool) {
	if o == nil || o.RepoMapping == nil {
		return nil, false
	}
	return o.RepoMapping, true
}

// HasRepoMapping returns a boolean if a field has been set.
func (o *PatchedPulpImporter) HasRepoMapping() bool {
	if o != nil && o.RepoMapping != nil {
		return true
	}

	return false
}

// SetRepoMapping gets a reference to the given map[string]string and assigns it to the RepoMapping field.
func (o *PatchedPulpImporter) SetRepoMapping(v map[string]string) {
	o.RepoMapping = &v
}

func (o PatchedPulpImporter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RepoMapping != nil {
		toSerialize["repo_mapping"] = o.RepoMapping
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPulpImporter struct {
	value *PatchedPulpImporter
	isSet bool
}

func (v NullablePatchedPulpImporter) Get() *PatchedPulpImporter {
	return v.value
}

func (v *NullablePatchedPulpImporter) Set(val *PatchedPulpImporter) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPulpImporter) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPulpImporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPulpImporter(val *PatchedPulpImporter) *NullablePatchedPulpImporter {
	return &NullablePatchedPulpImporter{value: val, isSet: true}
}

func (v NullablePatchedPulpImporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPulpImporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


