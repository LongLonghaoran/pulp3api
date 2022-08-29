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

// DebAptPublicationResponse A Serializer for AptPublication.
type DebAptPublicationResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	// Activate simple publishing mode (all packages in one release component).
	Simple *bool `json:"simple,omitempty"`
	// Activate structured publishing mode.
	Structured *bool `json:"structured,omitempty"`
	// Sign Release files with this signing key
	SigningService *string `json:"signing_service,omitempty"`
}

// NewDebAptPublicationResponse instantiates a new DebAptPublicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebAptPublicationResponse() *DebAptPublicationResponse {
	this := DebAptPublicationResponse{}
	var simple bool = false
	this.Simple = &simple
	var structured bool = false
	this.Structured = &structured
	return &this
}

// NewDebAptPublicationResponseWithDefaults instantiates a new DebAptPublicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebAptPublicationResponseWithDefaults() *DebAptPublicationResponse {
	this := DebAptPublicationResponse{}
	var simple bool = false
	this.Simple = &simple
	var structured bool = false
	this.Structured = &structured
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebAptPublicationResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebAptPublicationResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetRepositoryVersion() string {
	if o == nil || o.RepositoryVersion == nil {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || o.RepositoryVersion == nil {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion != nil {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *DebAptPublicationResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebAptPublicationResponse) SetRepository(v string) {
	o.Repository = &v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetSimple() bool {
	if o == nil || o.Simple == nil {
		var ret bool
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetSimpleOk() (*bool, bool) {
	if o == nil || o.Simple == nil {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasSimple() bool {
	if o != nil && o.Simple != nil {
		return true
	}

	return false
}

// SetSimple gets a reference to the given bool and assigns it to the Simple field.
func (o *DebAptPublicationResponse) SetSimple(v bool) {
	o.Simple = &v
}

// GetStructured returns the Structured field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetStructured() bool {
	if o == nil || o.Structured == nil {
		var ret bool
		return ret
	}
	return *o.Structured
}

// GetStructuredOk returns a tuple with the Structured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetStructuredOk() (*bool, bool) {
	if o == nil || o.Structured == nil {
		return nil, false
	}
	return o.Structured, true
}

// HasStructured returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasStructured() bool {
	if o != nil && o.Structured != nil {
		return true
	}

	return false
}

// SetStructured gets a reference to the given bool and assigns it to the Structured field.
func (o *DebAptPublicationResponse) SetStructured(v bool) {
	o.Structured = &v
}

// GetSigningService returns the SigningService field value if set, zero value otherwise.
func (o *DebAptPublicationResponse) GetSigningService() string {
	if o == nil || o.SigningService == nil {
		var ret string
		return ret
	}
	return *o.SigningService
}

// GetSigningServiceOk returns a tuple with the SigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptPublicationResponse) GetSigningServiceOk() (*string, bool) {
	if o == nil || o.SigningService == nil {
		return nil, false
	}
	return o.SigningService, true
}

// HasSigningService returns a boolean if a field has been set.
func (o *DebAptPublicationResponse) HasSigningService() bool {
	if o != nil && o.SigningService != nil {
		return true
	}

	return false
}

// SetSigningService gets a reference to the given string and assigns it to the SigningService field.
func (o *DebAptPublicationResponse) SetSigningService(v string) {
	o.SigningService = &v
}

func (o DebAptPublicationResponse) MarshalJSON() ([]byte, error) {
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
	if o.Simple != nil {
		toSerialize["simple"] = o.Simple
	}
	if o.Structured != nil {
		toSerialize["structured"] = o.Structured
	}
	if o.SigningService != nil {
		toSerialize["signing_service"] = o.SigningService
	}
	return json.Marshal(toSerialize)
}

type NullableDebAptPublicationResponse struct {
	value *DebAptPublicationResponse
	isSet bool
}

func (v NullableDebAptPublicationResponse) Get() *DebAptPublicationResponse {
	return v.value
}

func (v *NullableDebAptPublicationResponse) Set(val *DebAptPublicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebAptPublicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebAptPublicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebAptPublicationResponse(val *DebAptPublicationResponse) *NullableDebAptPublicationResponse {
	return &NullableDebAptPublicationResponse{value: val, isSet: true}
}

func (v NullableDebAptPublicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebAptPublicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


