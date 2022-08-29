# ContainerContainerDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**ContentGuard** | Pointer to **string** | An optional content-guard. If none is specified, a default one will be used. | [optional] 
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**Private** | Pointer to **bool** | Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 

## Methods

### NewContainerContainerDistribution

`func NewContainerContainerDistribution(name string, basePath string, ) *ContainerContainerDistribution`

NewContainerContainerDistribution instantiates a new ContainerContainerDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerDistributionWithDefaults

`func NewContainerContainerDistributionWithDefaults() *ContainerContainerDistribution`

NewContainerContainerDistributionWithDefaults instantiates a new ContainerContainerDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerContainerDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerDistribution) SetName(v string)`

SetName sets Name field to given value.


### GetPulpLabels

`func (o *ContainerContainerDistribution) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerDistribution) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerDistribution) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetRepository

`func (o *ContainerContainerDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ContainerContainerDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ContainerContainerDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ContainerContainerDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ContainerContainerDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ContainerContainerDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetContentGuard

`func (o *ContainerContainerDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *ContainerContainerDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *ContainerContainerDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *ContainerContainerDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### GetBasePath

`func (o *ContainerContainerDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *ContainerContainerDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *ContainerContainerDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetRepositoryVersion

`func (o *ContainerContainerDistribution) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *ContainerContainerDistribution) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *ContainerContainerDistribution) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *ContainerContainerDistribution) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *ContainerContainerDistribution) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *ContainerContainerDistribution) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPrivate

`func (o *ContainerContainerDistribution) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ContainerContainerDistribution) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ContainerContainerDistribution) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ContainerContainerDistribution) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetDescription

`func (o *ContainerContainerDistribution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerContainerDistribution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerContainerDistribution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerContainerDistribution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerContainerDistribution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerContainerDistribution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


