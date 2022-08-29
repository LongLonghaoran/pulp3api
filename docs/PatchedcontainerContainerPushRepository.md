# PatchedcontainerContainerPushRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for this repository. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**ManifestSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt32** | Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future. | [optional] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedcontainerContainerPushRepository

`func NewPatchedcontainerContainerPushRepository() *PatchedcontainerContainerPushRepository`

NewPatchedcontainerContainerPushRepository instantiates a new PatchedcontainerContainerPushRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcontainerContainerPushRepositoryWithDefaults

`func NewPatchedcontainerContainerPushRepositoryWithDefaults() *PatchedcontainerContainerPushRepository`

NewPatchedcontainerContainerPushRepositoryWithDefaults instantiates a new PatchedcontainerContainerPushRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedcontainerContainerPushRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcontainerContainerPushRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcontainerContainerPushRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcontainerContainerPushRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedcontainerContainerPushRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedcontainerContainerPushRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedcontainerContainerPushRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedcontainerContainerPushRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedcontainerContainerPushRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedcontainerContainerPushRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetManifestSigningService

`func (o *PatchedcontainerContainerPushRepository) GetManifestSigningService() string`

GetManifestSigningService returns the ManifestSigningService field if non-nil, zero value otherwise.

### GetManifestSigningServiceOk

`func (o *PatchedcontainerContainerPushRepository) GetManifestSigningServiceOk() (*string, bool)`

GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSigningService

`func (o *PatchedcontainerContainerPushRepository) SetManifestSigningService(v string)`

SetManifestSigningService sets ManifestSigningService field to given value.

### HasManifestSigningService

`func (o *PatchedcontainerContainerPushRepository) HasManifestSigningService() bool`

HasManifestSigningService returns a boolean if a field has been set.

### SetManifestSigningServiceNil

`func (o *PatchedcontainerContainerPushRepository) SetManifestSigningServiceNil(b bool)`

 SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil

### UnsetManifestSigningService
`func (o *PatchedcontainerContainerPushRepository) UnsetManifestSigningService()`

UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
### GetRetainRepoVersions

`func (o *PatchedcontainerContainerPushRepository) GetRetainRepoVersions() int32`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PatchedcontainerContainerPushRepository) GetRetainRepoVersionsOk() (*int32, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PatchedcontainerContainerPushRepository) SetRetainRepoVersions(v int32)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PatchedcontainerContainerPushRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PatchedcontainerContainerPushRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PatchedcontainerContainerPushRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetPulpLabels

`func (o *PatchedcontainerContainerPushRepository) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedcontainerContainerPushRepository) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedcontainerContainerPushRepository) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedcontainerContainerPushRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


