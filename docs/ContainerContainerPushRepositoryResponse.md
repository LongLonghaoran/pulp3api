# ContainerContainerPushRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**ManifestSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt32** | Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future. | [optional] 
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewContainerContainerPushRepositoryResponse

`func NewContainerContainerPushRepositoryResponse(name string, ) *ContainerContainerPushRepositoryResponse`

NewContainerContainerPushRepositoryResponse instantiates a new ContainerContainerPushRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerPushRepositoryResponseWithDefaults

`func NewContainerContainerPushRepositoryResponseWithDefaults() *ContainerContainerPushRepositoryResponse`

NewContainerContainerPushRepositoryResponseWithDefaults instantiates a new ContainerContainerPushRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerContainerPushRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerPushRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerPushRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContainerContainerPushRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerContainerPushRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerContainerPushRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerContainerPushRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerContainerPushRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerContainerPushRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetManifestSigningService

`func (o *ContainerContainerPushRepositoryResponse) GetManifestSigningService() string`

GetManifestSigningService returns the ManifestSigningService field if non-nil, zero value otherwise.

### GetManifestSigningServiceOk

`func (o *ContainerContainerPushRepositoryResponse) GetManifestSigningServiceOk() (*string, bool)`

GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSigningService

`func (o *ContainerContainerPushRepositoryResponse) SetManifestSigningService(v string)`

SetManifestSigningService sets ManifestSigningService field to given value.

### HasManifestSigningService

`func (o *ContainerContainerPushRepositoryResponse) HasManifestSigningService() bool`

HasManifestSigningService returns a boolean if a field has been set.

### SetManifestSigningServiceNil

`func (o *ContainerContainerPushRepositoryResponse) SetManifestSigningServiceNil(b bool)`

 SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil

### UnsetManifestSigningService
`func (o *ContainerContainerPushRepositoryResponse) UnsetManifestSigningService()`

UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
### GetRetainRepoVersions

`func (o *ContainerContainerPushRepositoryResponse) GetRetainRepoVersions() int32`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *ContainerContainerPushRepositoryResponse) GetRetainRepoVersionsOk() (*int32, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *ContainerContainerPushRepositoryResponse) SetRetainRepoVersions(v int32)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *ContainerContainerPushRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *ContainerContainerPushRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *ContainerContainerPushRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetPulpHref

`func (o *ContainerContainerPushRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerContainerPushRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerContainerPushRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerContainerPushRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *ContainerContainerPushRepositoryResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerPushRepositoryResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerPushRepositoryResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerPushRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *ContainerContainerPushRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *ContainerContainerPushRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *ContainerContainerPushRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *ContainerContainerPushRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerContainerPushRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerContainerPushRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerContainerPushRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerContainerPushRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *ContainerContainerPushRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *ContainerContainerPushRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *ContainerContainerPushRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *ContainerContainerPushRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


