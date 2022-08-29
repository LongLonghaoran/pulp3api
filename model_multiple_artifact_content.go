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

// MultipleArtifactContent Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type MultipleArtifactContent struct {
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
}

// NewMultipleArtifactContent instantiates a new MultipleArtifactContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleArtifactContent(artifacts map[string]interface{}, ) *MultipleArtifactContent {
	this := MultipleArtifactContent{}
	this.Artifacts = artifacts
	return &this
}

// NewMultipleArtifactContentWithDefaults instantiates a new MultipleArtifactContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleArtifactContentWithDefaults() *MultipleArtifactContent {
	this := MultipleArtifactContent{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *MultipleArtifactContent) GetArtifacts() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *MultipleArtifactContent) GetArtifactsOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacts, true
}

// SetArtifacts sets field value
func (o *MultipleArtifactContent) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

func (o MultipleArtifactContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleArtifactContent struct {
	value *MultipleArtifactContent
	isSet bool
}

func (v NullableMultipleArtifactContent) Get() *MultipleArtifactContent {
	return v.value
}

func (v *NullableMultipleArtifactContent) Set(val *MultipleArtifactContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleArtifactContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleArtifactContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleArtifactContent(val *MultipleArtifactContent) *NullableMultipleArtifactContent {
	return &NullableMultipleArtifactContent{value: val, isSet: true}
}

func (v NullableMultipleArtifactContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleArtifactContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


