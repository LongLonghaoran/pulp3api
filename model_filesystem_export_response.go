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

// FilesystemExportResponse Serializer for FilesystemExports.
type FilesystemExportResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A URI of the task that ran the Export.
	Task NullableString `json:"task,omitempty"`
	// Resources that were exported.
	ExportedResources []string `json:"exported_resources,omitempty"`
	// Any additional parameters that were used to create the export.
	Params map[string]interface{} `json:"params,omitempty"`
}

// NewFilesystemExportResponse instantiates a new FilesystemExportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemExportResponse() *FilesystemExportResponse {
	this := FilesystemExportResponse{}
	return &this
}

// NewFilesystemExportResponseWithDefaults instantiates a new FilesystemExportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemExportResponseWithDefaults() *FilesystemExportResponse {
	this := FilesystemExportResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *FilesystemExportResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExportResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *FilesystemExportResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *FilesystemExportResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *FilesystemExportResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExportResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *FilesystemExportResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *FilesystemExportResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetTask returns the Task field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesystemExportResponse) GetTask() string {
	if o == nil || o.Task.Get() == nil {
		var ret string
		return ret
	}
	return *o.Task.Get()
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesystemExportResponse) GetTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Task.Get(), o.Task.IsSet()
}

// HasTask returns a boolean if a field has been set.
func (o *FilesystemExportResponse) HasTask() bool {
	if o != nil && o.Task.IsSet() {
		return true
	}

	return false
}

// SetTask gets a reference to the given NullableString and assigns it to the Task field.
func (o *FilesystemExportResponse) SetTask(v string) {
	o.Task.Set(&v)
}
// SetTaskNil sets the value for Task to be an explicit nil
func (o *FilesystemExportResponse) SetTaskNil() {
	o.Task.Set(nil)
}

// UnsetTask ensures that no value is present for Task, not even an explicit nil
func (o *FilesystemExportResponse) UnsetTask() {
	o.Task.Unset()
}

// GetExportedResources returns the ExportedResources field value if set, zero value otherwise.
func (o *FilesystemExportResponse) GetExportedResources() []string {
	if o == nil || o.ExportedResources == nil {
		var ret []string
		return ret
	}
	return o.ExportedResources
}

// GetExportedResourcesOk returns a tuple with the ExportedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExportResponse) GetExportedResourcesOk() ([]string, bool) {
	if o == nil || o.ExportedResources == nil {
		return nil, false
	}
	return o.ExportedResources, true
}

// HasExportedResources returns a boolean if a field has been set.
func (o *FilesystemExportResponse) HasExportedResources() bool {
	if o != nil && o.ExportedResources != nil {
		return true
	}

	return false
}

// SetExportedResources gets a reference to the given []string and assigns it to the ExportedResources field.
func (o *FilesystemExportResponse) SetExportedResources(v []string) {
	o.ExportedResources = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *FilesystemExportResponse) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExportResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *FilesystemExportResponse) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *FilesystemExportResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o FilesystemExportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.Task.IsSet() {
		toSerialize["task"] = o.Task.Get()
	}
	if o.ExportedResources != nil {
		toSerialize["exported_resources"] = o.ExportedResources
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableFilesystemExportResponse struct {
	value *FilesystemExportResponse
	isSet bool
}

func (v NullableFilesystemExportResponse) Get() *FilesystemExportResponse {
	return v.value
}

func (v *NullableFilesystemExportResponse) Set(val *FilesystemExportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemExportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemExportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemExportResponse(val *FilesystemExportResponse) *NullableFilesystemExportResponse {
	return &NullableFilesystemExportResponse{value: val, isSet: true}
}

func (v NullableFilesystemExportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemExportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


