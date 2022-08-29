# ContainerContainerDistributionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**ContentGuard** | Pointer to **string** | An optional content-guard. If none is specified, a default one will be used. | [optional] 
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**RegistryPath** | Pointer to **string** | The Registry hostname/name/ to use with docker pull command defined by this distribution. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Namespace this distribution belongs to. | [optional] [readonly] 
**Private** | Pointer to **bool** | Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 

## Methods

### NewContainerContainerDistributionResponse

`func NewContainerContainerDistributionResponse(name string, basePath string, ) *ContainerContainerDistributionResponse`

NewContainerContainerDistributionResponse instantiates a new ContainerContainerDistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerDistributionResponseWithDefaults

`func NewContainerContainerDistributionResponseWithDefaults() *ContainerContainerDistributionResponse`

NewContainerContainerDistributionResponseWithDefaults instantiates a new ContainerContainerDistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerContainerDistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerDistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerDistributionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPulpHref

`func (o *ContainerContainerDistributionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerContainerDistributionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerContainerDistributionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerContainerDistributionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *ContainerContainerDistributionResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerDistributionResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerDistributionResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerDistributionResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerContainerDistributionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerContainerDistributionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerContainerDistributionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerContainerDistributionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetRepository

`func (o *ContainerContainerDistributionResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ContainerContainerDistributionResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ContainerContainerDistributionResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ContainerContainerDistributionResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ContainerContainerDistributionResponse) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ContainerContainerDistributionResponse) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetContentGuard

`func (o *ContainerContainerDistributionResponse) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *ContainerContainerDistributionResponse) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *ContainerContainerDistributionResponse) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *ContainerContainerDistributionResponse) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### GetBasePath

`func (o *ContainerContainerDistributionResponse) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *ContainerContainerDistributionResponse) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *ContainerContainerDistributionResponse) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetRepositoryVersion

`func (o *ContainerContainerDistributionResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *ContainerContainerDistributionResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *ContainerContainerDistributionResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *ContainerContainerDistributionResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *ContainerContainerDistributionResponse) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *ContainerContainerDistributionResponse) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetRegistryPath

`func (o *ContainerContainerDistributionResponse) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *ContainerContainerDistributionResponse) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *ContainerContainerDistributionResponse) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.

### HasRegistryPath

`func (o *ContainerContainerDistributionResponse) HasRegistryPath() bool`

HasRegistryPath returns a boolean if a field has been set.

### GetNamespace

`func (o *ContainerContainerDistributionResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ContainerContainerDistributionResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ContainerContainerDistributionResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ContainerContainerDistributionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPrivate

`func (o *ContainerContainerDistributionResponse) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ContainerContainerDistributionResponse) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ContainerContainerDistributionResponse) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ContainerContainerDistributionResponse) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetDescription

`func (o *ContainerContainerDistributionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerContainerDistributionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerContainerDistributionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerContainerDistributionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerContainerDistributionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerContainerDistributionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


