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

// DebGenericContentResponse A serializer for GenericContent.
type DebGenericContentResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	// The MD5 checksum if available.
	Md5 *string `json:"md5,omitempty"`
	// The SHA-1 checksum if available.
	Sha1 *string `json:"sha1,omitempty"`
	// The SHA-224 checksum if available.
	Sha224 *string `json:"sha224,omitempty"`
	// The SHA-256 checksum if available.
	Sha256 *string `json:"sha256,omitempty"`
	// The SHA-384 checksum if available.
	Sha384 *string `json:"sha384,omitempty"`
	// The SHA-512 checksum if available.
	Sha512 *string `json:"sha512,omitempty"`
}

// NewDebGenericContentResponse instantiates a new DebGenericContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebGenericContentResponse(relativePath string) *DebGenericContentResponse {
	this := DebGenericContentResponse{}
	this.RelativePath = relativePath
	return &this
}

// NewDebGenericContentResponseWithDefaults instantiates a new DebGenericContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebGenericContentResponseWithDefaults() *DebGenericContentResponse {
	this := DebGenericContentResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebGenericContentResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebGenericContentResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetArtifact() string {
	if o == nil || o.Artifact == nil {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetArtifactOk() (*string, bool) {
	if o == nil || o.Artifact == nil {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasArtifact() bool {
	if o != nil && o.Artifact != nil {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *DebGenericContentResponse) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value
func (o *DebGenericContentResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *DebGenericContentResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *DebGenericContentResponse) SetMd5(v string) {
	o.Md5 = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetSha1() string {
	if o == nil || o.Sha1 == nil {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetSha1Ok() (*string, bool) {
	if o == nil || o.Sha1 == nil {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasSha1() bool {
	if o != nil && o.Sha1 != nil {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *DebGenericContentResponse) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSha224 returns the Sha224 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetSha224() string {
	if o == nil || o.Sha224 == nil {
		var ret string
		return ret
	}
	return *o.Sha224
}

// GetSha224Ok returns a tuple with the Sha224 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetSha224Ok() (*string, bool) {
	if o == nil || o.Sha224 == nil {
		return nil, false
	}
	return o.Sha224, true
}

// HasSha224 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasSha224() bool {
	if o != nil && o.Sha224 != nil {
		return true
	}

	return false
}

// SetSha224 gets a reference to the given string and assigns it to the Sha224 field.
func (o *DebGenericContentResponse) SetSha224(v string) {
	o.Sha224 = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetSha256() string {
	if o == nil || o.Sha256 == nil {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetSha256Ok() (*string, bool) {
	if o == nil || o.Sha256 == nil {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasSha256() bool {
	if o != nil && o.Sha256 != nil {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *DebGenericContentResponse) SetSha256(v string) {
	o.Sha256 = &v
}

// GetSha384 returns the Sha384 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetSha384() string {
	if o == nil || o.Sha384 == nil {
		var ret string
		return ret
	}
	return *o.Sha384
}

// GetSha384Ok returns a tuple with the Sha384 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetSha384Ok() (*string, bool) {
	if o == nil || o.Sha384 == nil {
		return nil, false
	}
	return o.Sha384, true
}

// HasSha384 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasSha384() bool {
	if o != nil && o.Sha384 != nil {
		return true
	}

	return false
}

// SetSha384 gets a reference to the given string and assigns it to the Sha384 field.
func (o *DebGenericContentResponse) SetSha384(v string) {
	o.Sha384 = &v
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise.
func (o *DebGenericContentResponse) GetSha512() string {
	if o == nil || o.Sha512 == nil {
		var ret string
		return ret
	}
	return *o.Sha512
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebGenericContentResponse) GetSha512Ok() (*string, bool) {
	if o == nil || o.Sha512 == nil {
		return nil, false
	}
	return o.Sha512, true
}

// HasSha512 returns a boolean if a field has been set.
func (o *DebGenericContentResponse) HasSha512() bool {
	if o != nil && o.Sha512 != nil {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given string and assigns it to the Sha512 field.
func (o *DebGenericContentResponse) SetSha512(v string) {
	o.Sha512 = &v
}

func (o DebGenericContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.Artifact != nil {
		toSerialize["artifact"] = o.Artifact
	}
	if true {
		toSerialize["relative_path"] = o.RelativePath
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.Sha1 != nil {
		toSerialize["sha1"] = o.Sha1
	}
	if o.Sha224 != nil {
		toSerialize["sha224"] = o.Sha224
	}
	if o.Sha256 != nil {
		toSerialize["sha256"] = o.Sha256
	}
	if o.Sha384 != nil {
		toSerialize["sha384"] = o.Sha384
	}
	if o.Sha512 != nil {
		toSerialize["sha512"] = o.Sha512
	}
	return json.Marshal(toSerialize)
}

type NullableDebGenericContentResponse struct {
	value *DebGenericContentResponse
	isSet bool
}

func (v NullableDebGenericContentResponse) Get() *DebGenericContentResponse {
	return v.value
}

func (v *NullableDebGenericContentResponse) Set(val *DebGenericContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebGenericContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebGenericContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebGenericContentResponse(val *DebGenericContentResponse) *NullableDebGenericContentResponse {
	return &NullableDebGenericContentResponse{value: val, isSet: true}
}

func (v NullableDebGenericContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebGenericContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


