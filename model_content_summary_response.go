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

// ContentSummaryResponse Serializer for the RepositoryVersion content summary
type ContentSummaryResponse struct {
	Added map[string]map[string]interface{} `json:"added"`
	Removed map[string]map[string]interface{} `json:"removed"`
	Present map[string]map[string]interface{} `json:"present"`
}

// NewContentSummaryResponse instantiates a new ContentSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSummaryResponse(added map[string]map[string]interface{}, removed map[string]map[string]interface{}, present map[string]map[string]interface{}, ) *ContentSummaryResponse {
	this := ContentSummaryResponse{}
	this.Added = added
	this.Removed = removed
	this.Present = present
	return &this
}

// NewContentSummaryResponseWithDefaults instantiates a new ContentSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSummaryResponseWithDefaults() *ContentSummaryResponse {
	this := ContentSummaryResponse{}
	return &this
}

// GetAdded returns the Added field value
func (o *ContentSummaryResponse) GetAdded() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *ContentSummaryResponse) GetAddedOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *ContentSummaryResponse) SetAdded(v map[string]map[string]interface{}) {
	o.Added = v
}

// GetRemoved returns the Removed field value
func (o *ContentSummaryResponse) GetRemoved() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *ContentSummaryResponse) GetRemovedOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *ContentSummaryResponse) SetRemoved(v map[string]map[string]interface{}) {
	o.Removed = v
}

// GetPresent returns the Present field value
func (o *ContentSummaryResponse) GetPresent() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Present
}

// GetPresentOk returns a tuple with the Present field value
// and a boolean to check if the value has been set.
func (o *ContentSummaryResponse) GetPresentOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Present, true
}

// SetPresent sets field value
func (o *ContentSummaryResponse) SetPresent(v map[string]map[string]interface{}) {
	o.Present = v
}

func (o ContentSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["added"] = o.Added
	}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableContentSummaryResponse struct {
	value *ContentSummaryResponse
	isSet bool
}

func (v NullableContentSummaryResponse) Get() *ContentSummaryResponse {
	return v.value
}

func (v *NullableContentSummaryResponse) Set(val *ContentSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSummaryResponse(val *ContentSummaryResponse) *NullableContentSummaryResponse {
	return &NullableContentSummaryResponse{value: val, isSet: true}
}

func (v NullableContentSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


