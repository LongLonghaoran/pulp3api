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

// DebInstallerFileIndex A serializer for InstallerFileIndex.
type DebInstallerFileIndex struct {
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
	// Release this index file belongs to.
	Release string `json:"release"`
	// Component of the component - architecture combination.
	Component string `json:"component"`
	// Architecture of the component - architecture combination.
	Architecture string `json:"architecture"`
	// Path of directory containing MD5SUMS and SHA256SUMS relative to url.
	RelativePath *string `json:"relative_path,omitempty"`
}

// NewDebInstallerFileIndex instantiates a new DebInstallerFileIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebInstallerFileIndex(artifacts map[string]interface{}, release string, component string, architecture string) *DebInstallerFileIndex {
	this := DebInstallerFileIndex{}
	this.Artifacts = artifacts
	this.Release = release
	this.Component = component
	this.Architecture = architecture
	return &this
}

// NewDebInstallerFileIndexWithDefaults instantiates a new DebInstallerFileIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebInstallerFileIndexWithDefaults() *DebInstallerFileIndex {
	this := DebInstallerFileIndex{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *DebInstallerFileIndex) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *DebInstallerFileIndex) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetRelease returns the Release field value
func (o *DebInstallerFileIndex) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *DebInstallerFileIndex) SetRelease(v string) {
	o.Release = v
}

// GetComponent returns the Component field value
func (o *DebInstallerFileIndex) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DebInstallerFileIndex) SetComponent(v string) {
	o.Component = v
}

// GetArchitecture returns the Architecture field value
func (o *DebInstallerFileIndex) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *DebInstallerFileIndex) SetArchitecture(v string) {
	o.Architecture = v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *DebInstallerFileIndex) GetRelativePath() string {
	if o == nil || o.RelativePath == nil {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebInstallerFileIndex) GetRelativePathOk() (*string, bool) {
	if o == nil || o.RelativePath == nil {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *DebInstallerFileIndex) HasRelativePath() bool {
	if o != nil && o.RelativePath != nil {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *DebInstallerFileIndex) SetRelativePath(v string) {
	o.RelativePath = &v
}

func (o DebInstallerFileIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
	if true {
		toSerialize["release"] = o.Release
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["architecture"] = o.Architecture
	}
	if o.RelativePath != nil {
		toSerialize["relative_path"] = o.RelativePath
	}
	return json.Marshal(toSerialize)
}

type NullableDebInstallerFileIndex struct {
	value *DebInstallerFileIndex
	isSet bool
}

func (v NullableDebInstallerFileIndex) Get() *DebInstallerFileIndex {
	return v.value
}

func (v *NullableDebInstallerFileIndex) Set(val *DebInstallerFileIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableDebInstallerFileIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableDebInstallerFileIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebInstallerFileIndex(val *DebInstallerFileIndex) *NullableDebInstallerFileIndex {
	return &NullableDebInstallerFileIndex{value: val, isSet: true}
}

func (v NullableDebInstallerFileIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebInstallerFileIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


