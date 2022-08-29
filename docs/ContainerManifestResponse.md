# ContainerManifestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifact** | **string** | Artifact file representing the physical content | 
**Digest** | **string** | sha256 of the Manifest file | 
**SchemaVersion** | **int32** | Manifest schema version | 
**MediaType** | **string** | Manifest media type of the file | 
**ListedManifests** | **[]string** | Manifests that are referenced by this Manifest List | 
**ConfigBlob** | Pointer to **string** | Blob that contains configuration for this Manifest | [optional] 
**Blobs** | **[]string** | Blobs that are referenced by this Manifest | 

## Methods

### NewContainerManifestResponse

`func NewContainerManifestResponse(artifact string, digest string, schemaVersion int32, mediaType string, listedManifests []string, blobs []string, ) *ContainerManifestResponse`

NewContainerManifestResponse instantiates a new ContainerManifestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerManifestResponseWithDefaults

`func NewContainerManifestResponseWithDefaults() *ContainerManifestResponse`

NewContainerManifestResponseWithDefaults instantiates a new ContainerManifestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerManifestResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerManifestResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerManifestResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerManifestResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerManifestResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerManifestResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerManifestResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerManifestResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifact

`func (o *ContainerManifestResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *ContainerManifestResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *ContainerManifestResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetDigest

`func (o *ContainerManifestResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ContainerManifestResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ContainerManifestResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetSchemaVersion

`func (o *ContainerManifestResponse) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ContainerManifestResponse) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ContainerManifestResponse) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetMediaType

`func (o *ContainerManifestResponse) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ContainerManifestResponse) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ContainerManifestResponse) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetListedManifests

`func (o *ContainerManifestResponse) GetListedManifests() []string`

GetListedManifests returns the ListedManifests field if non-nil, zero value otherwise.

### GetListedManifestsOk

`func (o *ContainerManifestResponse) GetListedManifestsOk() (*[]string, bool)`

GetListedManifestsOk returns a tuple with the ListedManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedManifests

`func (o *ContainerManifestResponse) SetListedManifests(v []string)`

SetListedManifests sets ListedManifests field to given value.


### GetConfigBlob

`func (o *ContainerManifestResponse) GetConfigBlob() string`

GetConfigBlob returns the ConfigBlob field if non-nil, zero value otherwise.

### GetConfigBlobOk

`func (o *ContainerManifestResponse) GetConfigBlobOk() (*string, bool)`

GetConfigBlobOk returns a tuple with the ConfigBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBlob

`func (o *ContainerManifestResponse) SetConfigBlob(v string)`

SetConfigBlob sets ConfigBlob field to given value.

### HasConfigBlob

`func (o *ContainerManifestResponse) HasConfigBlob() bool`

HasConfigBlob returns a boolean if a field has been set.

### GetBlobs

`func (o *ContainerManifestResponse) GetBlobs() []string`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *ContainerManifestResponse) GetBlobsOk() (*[]string, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *ContainerManifestResponse) SetBlobs(v []string)`

SetBlobs sets Blobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


