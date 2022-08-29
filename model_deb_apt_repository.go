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

// DebAptRepository A Serializer for AptRepository.
type DebAptRepository struct {
	PulpLabels *map[string]interface{} `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future.
	RetainRepoVersions NullableInt32 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
}

// NewDebAptRepository instantiates a new DebAptRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebAptRepository(name string, ) *DebAptRepository {
	this := DebAptRepository{}
	this.Name = name
	return &this
}

// NewDebAptRepositoryWithDefaults instantiates a new DebAptRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebAptRepositoryWithDefaults() *DebAptRepository {
	this := DebAptRepository{}
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *DebAptRepository) GetPulpLabels() map[string]interface{} {
	if o == nil || o.PulpLabels == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptRepository) GetPulpLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || o.PulpLabels == nil {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *DebAptRepository) HasPulpLabels() bool {
	if o != nil && o.PulpLabels != nil {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]interface{} and assigns it to the PulpLabels field.
func (o *DebAptRepository) SetPulpLabels(v map[string]interface{}) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *DebAptRepository) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DebAptRepository) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DebAptRepository) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptRepository) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptRepository) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DebAptRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DebAptRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DebAptRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DebAptRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptRepository) GetRetainRepoVersions() int32 {
	if o == nil || o.RetainRepoVersions.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptRepository) GetRetainRepoVersionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *DebAptRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt32 and assigns it to the RetainRepoVersions field.
func (o *DebAptRepository) SetRetainRepoVersions(v int32) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *DebAptRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *DebAptRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptRepository) GetRemote() string {
	if o == nil || o.Remote.Get() == nil {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptRepository) GetRemoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *DebAptRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *DebAptRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *DebAptRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *DebAptRepository) UnsetRemote() {
	o.Remote.Unset()
}

func (o DebAptRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpLabels != nil {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDebAptRepository struct {
	value *DebAptRepository
	isSet bool
}

func (v NullableDebAptRepository) Get() *DebAptRepository {
	return v.value
}

func (v *NullableDebAptRepository) Set(val *DebAptRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableDebAptRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableDebAptRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebAptRepository(val *DebAptRepository) *NullableDebAptRepository {
	return &NullableDebAptRepository{value: val, isSet: true}
}

func (v NullableDebAptRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebAptRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


