# ContainerContainerRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt32** | Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**ManifestSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 

## Methods

### NewContainerContainerRepositoryResponse

`func NewContainerContainerRepositoryResponse(name string, ) *ContainerContainerRepositoryResponse`

NewContainerContainerRepositoryResponse instantiates a new ContainerContainerRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerRepositoryResponseWithDefaults

`func NewContainerContainerRepositoryResponseWithDefaults() *ContainerContainerRepositoryResponse`

NewContainerContainerRepositoryResponseWithDefaults instantiates a new ContainerContainerRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerContainerRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerContainerRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerContainerRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerContainerRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerContainerRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerContainerRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerContainerRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerContainerRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *ContainerContainerRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *ContainerContainerRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *ContainerContainerRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *ContainerContainerRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *ContainerContainerRepositoryResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerRepositoryResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerRepositoryResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *ContainerContainerRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *ContainerContainerRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *ContainerContainerRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *ContainerContainerRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *ContainerContainerRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContainerContainerRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerContainerRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerContainerRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerContainerRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerContainerRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerContainerRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersions() int32`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersionsOk() (*int32, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersions(v int32)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *ContainerContainerRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *ContainerContainerRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *ContainerContainerRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ContainerContainerRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ContainerContainerRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ContainerContainerRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *ContainerContainerRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *ContainerContainerRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetManifestSigningService

`func (o *ContainerContainerRepositoryResponse) GetManifestSigningService() string`

GetManifestSigningService returns the ManifestSigningService field if non-nil, zero value otherwise.

### GetManifestSigningServiceOk

`func (o *ContainerContainerRepositoryResponse) GetManifestSigningServiceOk() (*string, bool)`

GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSigningService

`func (o *ContainerContainerRepositoryResponse) SetManifestSigningService(v string)`

SetManifestSigningService sets ManifestSigningService field to given value.

### HasManifestSigningService

`func (o *ContainerContainerRepositoryResponse) HasManifestSigningService() bool`

HasManifestSigningService returns a boolean if a field has been set.

### SetManifestSigningServiceNil

`func (o *ContainerContainerRepositoryResponse) SetManifestSigningServiceNil(b bool)`

 SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil

### UnsetManifestSigningService
`func (o *ContainerContainerRepositoryResponse) UnsetManifestSigningService()`

UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


