# ContainerContainerRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt32** | Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**ManifestSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 

## Methods

### NewContainerContainerRepository

`func NewContainerContainerRepository(name string, ) *ContainerContainerRepository`

NewContainerContainerRepository instantiates a new ContainerContainerRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerRepositoryWithDefaults

`func NewContainerContainerRepositoryWithDefaults() *ContainerContainerRepository`

NewContainerContainerRepositoryWithDefaults instantiates a new ContainerContainerRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *ContainerContainerRepository) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerRepository) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerRepository) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *ContainerContainerRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContainerContainerRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerContainerRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerContainerRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerContainerRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerContainerRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerContainerRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *ContainerContainerRepository) GetRetainRepoVersions() int32`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *ContainerContainerRepository) GetRetainRepoVersionsOk() (*int32, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *ContainerContainerRepository) SetRetainRepoVersions(v int32)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *ContainerContainerRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *ContainerContainerRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *ContainerContainerRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *ContainerContainerRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ContainerContainerRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ContainerContainerRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ContainerContainerRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *ContainerContainerRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *ContainerContainerRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetManifestSigningService

`func (o *ContainerContainerRepository) GetManifestSigningService() string`

GetManifestSigningService returns the ManifestSigningService field if non-nil, zero value otherwise.

### GetManifestSigningServiceOk

`func (o *ContainerContainerRepository) GetManifestSigningServiceOk() (*string, bool)`

GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSigningService

`func (o *ContainerContainerRepository) SetManifestSigningService(v string)`

SetManifestSigningService sets ManifestSigningService field to given value.

### HasManifestSigningService

`func (o *ContainerContainerRepository) HasManifestSigningService() bool`

HasManifestSigningService returns a boolean if a field has been set.

### SetManifestSigningServiceNil

`func (o *ContainerContainerRepository) SetManifestSigningServiceNil(b bool)`

 SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil

### UnsetManifestSigningService
`func (o *ContainerContainerRepository) UnsetManifestSigningService()`

UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


