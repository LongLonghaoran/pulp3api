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

// DebReleaseResponse A Serializer for Release.
type DebReleaseResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Codename string `json:"codename"`
	Suite string `json:"suite"`
	Distribution string `json:"distribution"`
}

// NewDebReleaseResponse instantiates a new DebReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebReleaseResponse(codename string, suite string, distribution string, ) *DebReleaseResponse {
	this := DebReleaseResponse{}
	this.Codename = codename
	this.Suite = suite
	this.Distribution = distribution
	return &this
}

// NewDebReleaseResponseWithDefaults instantiates a new DebReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseResponseWithDefaults() *DebReleaseResponse {
	this := DebReleaseResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebReleaseResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebReleaseResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebReleaseResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebReleaseResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebReleaseResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebReleaseResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetCodename returns the Codename field value
func (o *DebReleaseResponse) GetCodename() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *DebReleaseResponse) GetCodenameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *DebReleaseResponse) SetCodename(v string) {
	o.Codename = v
}

// GetSuite returns the Suite field value
func (o *DebReleaseResponse) GetSuite() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value
// and a boolean to check if the value has been set.
func (o *DebReleaseResponse) GetSuiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Suite, true
}

// SetSuite sets field value
func (o *DebReleaseResponse) SetSuite(v string) {
	o.Suite = v
}

// GetDistribution returns the Distribution field value
func (o *DebReleaseResponse) GetDistribution() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *DebReleaseResponse) GetDistributionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *DebReleaseResponse) SetDistribution(v string) {
	o.Distribution = v
}

func (o DebReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if true {
		toSerialize["codename"] = o.Codename
	}
	if true {
		toSerialize["suite"] = o.Suite
	}
	if true {
		toSerialize["distribution"] = o.Distribution
	}
	return json.Marshal(toSerialize)
}

type NullableDebReleaseResponse struct {
	value *DebReleaseResponse
	isSet bool
}

func (v NullableDebReleaseResponse) Get() *DebReleaseResponse {
	return v.value
}

func (v *NullableDebReleaseResponse) Set(val *DebReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebReleaseResponse(val *DebReleaseResponse) *NullableDebReleaseResponse {
	return &NullableDebReleaseResponse{value: val, isSet: true}
}

func (v NullableDebReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


