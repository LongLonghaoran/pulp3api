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

// ProgressReportResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type ProgressReportResponse struct {
	// The message shown to the user for the progress report.
	Message *string `json:"message,omitempty"`
	// Identifies the type of progress report'.
	Code *string `json:"code,omitempty"`
	// The current state of the progress report. The possible values are: 'waiting', 'skipped', 'running', 'completed', 'failed', 'canceled' and 'canceling'. The default is 'waiting'.
	State *string `json:"state,omitempty"`
	// The total count of items.
	Total *int32 `json:"total,omitempty"`
	// The count of items already processed. Defaults to 0.
	Done *int32 `json:"done,omitempty"`
	// The suffix to be shown with the progress report.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewProgressReportResponse instantiates a new ProgressReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressReportResponse() *ProgressReportResponse {
	this := ProgressReportResponse{}
	return &this
}

// NewProgressReportResponseWithDefaults instantiates a new ProgressReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressReportResponseWithDefaults() *ProgressReportResponse {
	this := ProgressReportResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProgressReportResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressReportResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProgressReportResponse) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProgressReportResponse) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressReportResponse) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProgressReportResponse) SetCode(v string) {
	o.Code = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProgressReportResponse) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressReportResponse) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProgressReportResponse) SetState(v string) {
	o.State = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProgressReportResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressReportResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ProgressReportResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *ProgressReportResponse) GetDone() int32 {
	if o == nil || o.Done == nil {
		var ret int32
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressReportResponse) GetDoneOk() (*int32, bool) {
	if o == nil || o.Done == nil {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasDone() bool {
	if o != nil && o.Done != nil {
		return true
	}

	return false
}

// SetDone gets a reference to the given int32 and assigns it to the Done field.
func (o *ProgressReportResponse) SetDone(v int32) {
	o.Done = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressReportResponse) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressReportResponse) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *ProgressReportResponse) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *ProgressReportResponse) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *ProgressReportResponse) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *ProgressReportResponse) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o ProgressReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Done != nil {
		toSerialize["done"] = o.Done
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProgressReportResponse struct {
	value *ProgressReportResponse
	isSet bool
}

func (v NullableProgressReportResponse) Get() *ProgressReportResponse {
	return v.value
}

func (v *NullableProgressReportResponse) Set(val *ProgressReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressReportResponse(val *ProgressReportResponse) *NullableProgressReportResponse {
	return &NullableProgressReportResponse{value: val, isSet: true}
}

func (v NullableProgressReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


