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

// Repair struct for Repair
type Repair struct {
	// Will verify that the checksum of all stored files matches what saved in the database. Otherwise only the existence of the files will be checked. Enabled by default
	VerifyChecksums *bool `json:"verify_checksums,omitempty"`
}

// NewRepair instantiates a new Repair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepair() *Repair {
	this := Repair{}
	var verifyChecksums bool = true
	this.VerifyChecksums = &verifyChecksums
	return &this
}

// NewRepairWithDefaults instantiates a new Repair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairWithDefaults() *Repair {
	this := Repair{}
	var verifyChecksums bool = true
	this.VerifyChecksums = &verifyChecksums
	return &this
}

// GetVerifyChecksums returns the VerifyChecksums field value if set, zero value otherwise.
func (o *Repair) GetVerifyChecksums() bool {
	if o == nil || o.VerifyChecksums == nil {
		var ret bool
		return ret
	}
	return *o.VerifyChecksums
}

// GetVerifyChecksumsOk returns a tuple with the VerifyChecksums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repair) GetVerifyChecksumsOk() (*bool, bool) {
	if o == nil || o.VerifyChecksums == nil {
		return nil, false
	}
	return o.VerifyChecksums, true
}

// HasVerifyChecksums returns a boolean if a field has been set.
func (o *Repair) HasVerifyChecksums() bool {
	if o != nil && o.VerifyChecksums != nil {
		return true
	}

	return false
}

// SetVerifyChecksums gets a reference to the given bool and assigns it to the VerifyChecksums field.
func (o *Repair) SetVerifyChecksums(v bool) {
	o.VerifyChecksums = &v
}

func (o Repair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VerifyChecksums != nil {
		toSerialize["verify_checksums"] = o.VerifyChecksums
	}
	return json.Marshal(toSerialize)
}

type NullableRepair struct {
	value *Repair
	isSet bool
}

func (v NullableRepair) Get() *Repair {
	return v.value
}

func (v *NullableRepair) Set(val *Repair) {
	v.value = val
	v.isSet = true
}

func (v NullableRepair) IsSet() bool {
	return v.isSet
}

func (v *NullableRepair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepair(val *Repair) *NullableRepair {
	return &NullableRepair{value: val, isSet: true}
}

func (v NullableRepair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


