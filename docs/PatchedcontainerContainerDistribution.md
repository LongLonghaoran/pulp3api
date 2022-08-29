# PatchedcontainerContainerDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | [optional] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**ContentGuard** | Pointer to **string** | An optional content-guard. If none is specified, a default one will be used. | [optional] 
**BasePath** | Pointer to **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**Private** | Pointer to **bool** | Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 

## Methods

### NewPatchedcontainerContainerDistribution

`func NewPatchedcontainerContainerDistribution() *PatchedcontainerContainerDistribution`

NewPatchedcontainerContainerDistribution instantiates a new PatchedcontainerContainerDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcontainerContainerDistributionWithDefaults

`func NewPatchedcontainerContainerDistributionWithDefaults() *PatchedcontainerContainerDistribution`

NewPatchedcontainerContainerDistributionWithDefaults instantiates a new PatchedcontainerContainerDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedcontainerContainerDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcontainerContainerDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcontainerContainerDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcontainerContainerDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PatchedcontainerContainerDistribution) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedcontainerContainerDistribution) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedcontainerContainerDistribution) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedcontainerContainerDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedcontainerContainerDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedcontainerContainerDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedcontainerContainerDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedcontainerContainerDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedcontainerContainerDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedcontainerContainerDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetContentGuard

`func (o *PatchedcontainerContainerDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedcontainerContainerDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedcontainerContainerDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedcontainerContainerDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### GetBasePath

`func (o *PatchedcontainerContainerDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedcontainerContainerDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedcontainerContainerDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedcontainerContainerDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *PatchedcontainerContainerDistribution) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *PatchedcontainerContainerDistribution) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *PatchedcontainerContainerDistribution) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *PatchedcontainerContainerDistribution) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *PatchedcontainerContainerDistribution) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *PatchedcontainerContainerDistribution) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPrivate

`func (o *PatchedcontainerContainerDistribution) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PatchedcontainerContainerDistribution) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PatchedcontainerContainerDistribution) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *PatchedcontainerContainerDistribution) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedcontainerContainerDistribution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedcontainerContainerDistribution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedcontainerContainerDistribution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedcontainerContainerDistribution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedcontainerContainerDistribution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedcontainerContainerDistribution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


