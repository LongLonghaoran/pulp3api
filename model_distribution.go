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
)

// Distribution The Serializer for the Distribution model.  The serializer deliberately omits the `publication` and `repository_version` field due to plugins typically requiring one or the other but not both.  To include the ``publication`` field, it is recommended plugins define the field::    publication = DetailRelatedField(       required=False,       help_text=_(\"Publication to be served\"),       view_name_pattern=r\"publications(-.*_/.*)?-detail\",       queryset=models.Publication.objects.exclude(complete=False),       allow_null=True,   )  To include the ``repository_version`` field, it is recommended plugins define the field::    repository_version = RepositoryVersionRelatedField(       required=False, help_text=_(\"RepositoryVersion to be served\"), allow_null=True   )  Additionally, the serializer omits the ``remote`` field, which is used for pull-through caching feature and only by plugins which use publications. Plugins implementing a pull-through caching should define the field in their derived serializer class like this::    remote = DetailRelatedField(       required=False,       help_text=_('Remote that can be used to fetch content when using pull-through caching.'),       queryset=models.Remote.objects.all(),       allow_null=True   )
type Distribution struct {
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	PulpLabels map[string]interface{} `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
}

// NewDistribution instantiates a new Distribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistribution(basePath string, name string) *Distribution {
	this := Distribution{}
	this.BasePath = basePath
	this.Name = name
	return &this
}

// NewDistributionWithDefaults instantiates a new Distribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionWithDefaults() *Distribution {
	this := Distribution{}
	return &this
}

// GetBasePath returns the BasePath field value
func (o *Distribution) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *Distribution) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *Distribution) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Distribution) GetContentGuard() string {
	if o == nil || o.ContentGuard.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Distribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *Distribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *Distribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *Distribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *Distribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *Distribution) GetPulpLabels() map[string]interface{} {
	if o == nil || o.PulpLabels == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetPulpLabelsOk() (map[string]interface{}, bool) {
	if o == nil || o.PulpLabels == nil {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *Distribution) HasPulpLabels() bool {
	if o != nil && o.PulpLabels != nil {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]interface{} and assigns it to the PulpLabels field.
func (o *Distribution) SetPulpLabels(v map[string]interface{}) {
	o.PulpLabels = v
}

// GetName returns the Name field value
func (o *Distribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Distribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Distribution) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Distribution) GetRepository() string {
	if o == nil || o.Repository.Get() == nil {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Distribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *Distribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *Distribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *Distribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *Distribution) UnsetRepository() {
	o.Repository.Unset()
}

func (o Distribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["base_path"] = o.BasePath
	}
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if o.PulpLabels != nil {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDistribution struct {
	value *Distribution
	isSet bool
}

func (v NullableDistribution) Get() *Distribution {
	return v.value
}

func (v *NullableDistribution) Set(val *Distribution) {
	v.value = val
	v.isSet = true
}

func (v NullableDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistribution(val *Distribution) *NullableDistribution {
	return &NullableDistribution{value: val, isSet: true}
}

func (v NullableDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


