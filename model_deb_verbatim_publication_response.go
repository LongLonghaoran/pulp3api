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
	"time"
)

// DebVerbatimPublicationResponse A Serializer for VerbatimPublication.
type DebVerbatimPublicationResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
}

// NewDebVerbatimPublicationResponse instantiates a new DebVerbatimPublicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebVerbatimPublicationResponse() *DebVerbatimPublicationResponse {
	this := DebVerbatimPublicationResponse{}
	return &this
}

// NewDebVerbatimPublicationResponseWithDefaults instantiates a new DebVerbatimPublicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebVerbatimPublicationResponseWithDefaults() *DebVerbatimPublicationResponse {
	this := DebVerbatimPublicationResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebVerbatimPublicationResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublicationResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebVerbatimPublicationResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebVerbatimPublicationResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebVerbatimPublicationResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublicationResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebVerbatimPublicationResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebVerbatimPublicationResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *DebVerbatimPublicationResponse) GetRepositoryVersion() string {
	if o == nil || o.RepositoryVersion == nil {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublicationResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || o.RepositoryVersion == nil {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *DebVerbatimPublicationResponse) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion != nil {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *DebVerbatimPublicationResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebVerbatimPublicationResponse) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebVerbatimPublicationResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebVerbatimPublicationResponse) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebVerbatimPublicationResponse) SetRepository(v string) {
	o.Repository = &v
}

func (o DebVerbatimPublicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.RepositoryVersion != nil {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableDebVerbatimPublicationResponse struct {
	value *DebVerbatimPublicationResponse
	isSet bool
}

func (v NullableDebVerbatimPublicationResponse) Get() *DebVerbatimPublicationResponse {
	return v.value
}

func (v *NullableDebVerbatimPublicationResponse) Set(val *DebVerbatimPublicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebVerbatimPublicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebVerbatimPublicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebVerbatimPublicationResponse(val *DebVerbatimPublicationResponse) *NullableDebVerbatimPublicationResponse {
	return &NullableDebVerbatimPublicationResponse{value: val, isSet: true}
}

func (v NullableDebVerbatimPublicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebVerbatimPublicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


