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

// RepositoryVersionResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type RepositoryVersionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Number *int32 `json:"number,omitempty"`
	Repository *string `json:"repository,omitempty"`
	// A repository version whose content was used as the initial set of content for this repository version
	BaseVersion *string `json:"base_version,omitempty"`
	// Various count summaries of the content in the version and the HREF to view them.
	ContentSummary *ContentSummaryResponse `json:"content_summary,omitempty"`
}

// NewRepositoryVersionResponse instantiates a new RepositoryVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryVersionResponse() *RepositoryVersionResponse {
	this := RepositoryVersionResponse{}
	return &this
}

// NewRepositoryVersionResponseWithDefaults instantiates a new RepositoryVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryVersionResponseWithDefaults() *RepositoryVersionResponse {
	this := RepositoryVersionResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RepositoryVersionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RepositoryVersionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *RepositoryVersionResponse) SetNumber(v int32) {
	o.Number = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RepositoryVersionResponse) SetRepository(v string) {
	o.Repository = &v
}

// GetBaseVersion returns the BaseVersion field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetBaseVersion() string {
	if o == nil || o.BaseVersion == nil {
		var ret string
		return ret
	}
	return *o.BaseVersion
}

// GetBaseVersionOk returns a tuple with the BaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetBaseVersionOk() (*string, bool) {
	if o == nil || o.BaseVersion == nil {
		return nil, false
	}
	return o.BaseVersion, true
}

// HasBaseVersion returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasBaseVersion() bool {
	if o != nil && o.BaseVersion != nil {
		return true
	}

	return false
}

// SetBaseVersion gets a reference to the given string and assigns it to the BaseVersion field.
func (o *RepositoryVersionResponse) SetBaseVersion(v string) {
	o.BaseVersion = &v
}

// GetContentSummary returns the ContentSummary field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetContentSummary() ContentSummaryResponse {
	if o == nil || o.ContentSummary == nil {
		var ret ContentSummaryResponse
		return ret
	}
	return *o.ContentSummary
}

// GetContentSummaryOk returns a tuple with the ContentSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetContentSummaryOk() (*ContentSummaryResponse, bool) {
	if o == nil || o.ContentSummary == nil {
		return nil, false
	}
	return o.ContentSummary, true
}

// HasContentSummary returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasContentSummary() bool {
	if o != nil && o.ContentSummary != nil {
		return true
	}

	return false
}

// SetContentSummary gets a reference to the given ContentSummaryResponse and assigns it to the ContentSummary field.
func (o *RepositoryVersionResponse) SetContentSummary(v ContentSummaryResponse) {
	o.ContentSummary = &v
}

func (o RepositoryVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.BaseVersion != nil {
		toSerialize["base_version"] = o.BaseVersion
	}
	if o.ContentSummary != nil {
		toSerialize["content_summary"] = o.ContentSummary
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryVersionResponse struct {
	value *RepositoryVersionResponse
	isSet bool
}

func (v NullableRepositoryVersionResponse) Get() *RepositoryVersionResponse {
	return v.value
}

func (v *NullableRepositoryVersionResponse) Set(val *RepositoryVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryVersionResponse(val *RepositoryVersionResponse) *NullableRepositoryVersionResponse {
	return &NullableRepositoryVersionResponse{value: val, isSet: true}
}

func (v NullableRepositoryVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


