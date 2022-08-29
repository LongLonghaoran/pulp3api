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

// Importer Base serializer for Importers.
type Importer struct {
	// Unique name of the Importer.
	Name string `json:"name"`
}

// NewImporter instantiates a new Importer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImporter(name string, ) *Importer {
	this := Importer{}
	this.Name = name
	return &this
}

// NewImporterWithDefaults instantiates a new Importer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImporterWithDefaults() *Importer {
	this := Importer{}
	return &this
}

// GetName returns the Name field value
func (o *Importer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Importer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Importer) SetName(v string) {
	o.Name = v
}

func (o Importer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableImporter struct {
	value *Importer
	isSet bool
}

func (v NullableImporter) Get() *Importer {
	return v.value
}

func (v *NullableImporter) Set(val *Importer) {
	v.value = val
	v.isSet = true
}

func (v NullableImporter) IsSet() bool {
	return v.isSet
}

func (v *NullableImporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImporter(val *Importer) *NullableImporter {
	return &NullableImporter{value: val, isSet: true}
}

func (v NullableImporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


