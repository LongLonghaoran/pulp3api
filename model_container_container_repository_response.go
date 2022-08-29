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
	"time"
)

// ContainerContainerRepositoryResponse Serializer for Container Repositories.
type ContainerContainerRepositoryResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	VersionsHref *string `json:"versions_href,omitempty"`
	PulpLabels map[string]interface{} `json:"pulp_labels,omitempty"`
	LatestVersionHref *string `json:"latest_version_href,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future.
	RetainRepoVersions NullableInt32 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// A reference to an associated signing service.
	ManifestSigningService NullableString `json:"manifest_signing_service,omitempty"`
}

// NewContainerContainerRepositoryResponse instantiates a new ContainerContainerRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerRepositoryResponse(name string) *ContainerContainerRepositoryResponse {
	this := ContainerContainerRepositoryResponse{}
	this.Name = name
	return &this
}

// NewContainerContainerRepositoryResponseWithDefaults instantiates a new ContainerContainerRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerRepositoryResponseWithDefaults() *ContainerContainerRepositoryResponse {
	this := ContainerContainerRepositoryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerRepositoryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerRepositoryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetVersionsHref returns the VersionsHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetVersionsHref() string {
	if o == nil || o.VersionsHref == nil {
		var ret string
		return ret
	}
	return *o.VersionsHref
}

// GetVersionsHrefOk returns a tuple with the VersionsHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetVersionsHrefOk() (*string, bool) {
	if o == nil || o.VersionsHref == nil {
		return nil, false
	}
	return o.VersionsHref, true
}

// HasVersionsHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasVersionsHref() bool {
	if o != nil && o.VersionsHref != nil {
		return true
	}

	return false
}

// SetVersionsHref gets a reference to the given string and assigns it to the VersionsHref field.
func (o *ContainerContainerRepositoryResponse) SetVersionsHref(v string) {
	o.VersionsHref = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpLabels() map[string]interface{} {
	if o == nil || o.PulpLabels == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpLabelsOk() (map[string]interface{}, bool) {
	if o == nil || o.PulpLabels == nil {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpLabels() bool {
	if o != nil && o.PulpLabels != nil {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]interface{} and assigns it to the PulpLabels field.
func (o *ContainerContainerRepositoryResponse) SetPulpLabels(v map[string]interface{}) {
	o.PulpLabels = v
}

// GetLatestVersionHref returns the LatestVersionHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetLatestVersionHref() string {
	if o == nil || o.LatestVersionHref == nil {
		var ret string
		return ret
	}
	return *o.LatestVersionHref
}

// GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetLatestVersionHrefOk() (*string, bool) {
	if o == nil || o.LatestVersionHref == nil {
		return nil, false
	}
	return o.LatestVersionHref, true
}

// HasLatestVersionHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasLatestVersionHref() bool {
	if o != nil && o.LatestVersionHref != nil {
		return true
	}

	return false
}

// SetLatestVersionHref gets a reference to the given string and assigns it to the LatestVersionHref field.
func (o *ContainerContainerRepositoryResponse) SetLatestVersionHref(v string) {
	o.LatestVersionHref = &v
}

// GetName returns the Name field value
func (o *ContainerContainerRepositoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerRepositoryResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerRepositoryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersions() int32 {
	if o == nil || o.RetainRepoVersions.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt32 and assigns it to the RetainRepoVersions field.
func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersions(v int32) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetRemote() string {
	if o == nil || o.Remote.Get() == nil {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *ContainerContainerRepositoryResponse) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetRemote() {
	o.Remote.Unset()
}

// GetManifestSigningService returns the ManifestSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetManifestSigningService() string {
	if o == nil || o.ManifestSigningService.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManifestSigningService.Get()
}

// GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetManifestSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManifestSigningService.Get(), o.ManifestSigningService.IsSet()
}

// HasManifestSigningService returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasManifestSigningService() bool {
	if o != nil && o.ManifestSigningService.IsSet() {
		return true
	}

	return false
}

// SetManifestSigningService gets a reference to the given NullableString and assigns it to the ManifestSigningService field.
func (o *ContainerContainerRepositoryResponse) SetManifestSigningService(v string) {
	o.ManifestSigningService.Set(&v)
}
// SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetManifestSigningServiceNil() {
	o.ManifestSigningService.Set(nil)
}

// UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetManifestSigningService() {
	o.ManifestSigningService.Unset()
}

func (o ContainerContainerRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.VersionsHref != nil {
		toSerialize["versions_href"] = o.VersionsHref
	}
	if o.PulpLabels != nil {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if o.LatestVersionHref != nil {
		toSerialize["latest_version_href"] = o.LatestVersionHref
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
	if o.ManifestSigningService.IsSet() {
		toSerialize["manifest_signing_service"] = o.ManifestSigningService.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContainerRepositoryResponse struct {
	value *ContainerContainerRepositoryResponse
	isSet bool
}

func (v NullableContainerContainerRepositoryResponse) Get() *ContainerContainerRepositoryResponse {
	return v.value
}

func (v *NullableContainerContainerRepositoryResponse) Set(val *ContainerContainerRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerRepositoryResponse(val *ContainerContainerRepositoryResponse) *NullableContainerContainerRepositoryResponse {
	return &NullableContainerContainerRepositoryResponse{value: val, isSet: true}
}

func (v NullableContainerContainerRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


