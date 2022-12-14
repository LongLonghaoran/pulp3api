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

// MinimalTaskResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type MinimalTaskResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The name of task.
	Name string `json:"name"`
	// The current state of the task. The possible values include: 'waiting', 'skipped', 'running', 'completed', 'failed', 'canceled' and 'canceling'.
	State *string `json:"state,omitempty"`
	// Timestamp of the when this task started execution.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// Timestamp of the when this task stopped execution.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// The worker associated with this task. This field is empty if a worker is not yet assigned.
	Worker *string `json:"worker,omitempty"`
}

// NewMinimalTaskResponse instantiates a new MinimalTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalTaskResponse(name string) *MinimalTaskResponse {
	this := MinimalTaskResponse{}
	this.Name = name
	return &this
}

// NewMinimalTaskResponseWithDefaults instantiates a new MinimalTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalTaskResponseWithDefaults() *MinimalTaskResponse {
	this := MinimalTaskResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetPulpHref() string {
	if o == nil || o.PulpHref == nil {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || o.PulpHref == nil {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasPulpHref() bool {
	if o != nil && o.PulpHref != nil {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *MinimalTaskResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetPulpCreated() time.Time {
	if o == nil || o.PulpCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || o.PulpCreated == nil {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasPulpCreated() bool {
	if o != nil && o.PulpCreated != nil {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *MinimalTaskResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *MinimalTaskResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MinimalTaskResponse) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MinimalTaskResponse) SetState(v string) {
	o.State = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *MinimalTaskResponse) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *MinimalTaskResponse) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetWorker returns the Worker field value if set, zero value otherwise.
func (o *MinimalTaskResponse) GetWorker() string {
	if o == nil || o.Worker == nil {
		var ret string
		return ret
	}
	return *o.Worker
}

// GetWorkerOk returns a tuple with the Worker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTaskResponse) GetWorkerOk() (*string, bool) {
	if o == nil || o.Worker == nil {
		return nil, false
	}
	return o.Worker, true
}

// HasWorker returns a boolean if a field has been set.
func (o *MinimalTaskResponse) HasWorker() bool {
	if o != nil && o.Worker != nil {
		return true
	}

	return false
}

// SetWorker gets a reference to the given string and assigns it to the Worker field.
func (o *MinimalTaskResponse) SetWorker(v string) {
	o.Worker = &v
}

func (o MinimalTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PulpHref != nil {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if o.PulpCreated != nil {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Worker != nil {
		toSerialize["worker"] = o.Worker
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalTaskResponse struct {
	value *MinimalTaskResponse
	isSet bool
}

func (v NullableMinimalTaskResponse) Get() *MinimalTaskResponse {
	return v.value
}

func (v *NullableMinimalTaskResponse) Set(val *MinimalTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalTaskResponse(val *MinimalTaskResponse) *NullableMinimalTaskResponse {
	return &NullableMinimalTaskResponse{value: val, isSet: true}
}

func (v NullableMinimalTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


