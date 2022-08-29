# ManifestCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRepository** | Pointer to **string** | A URI of the repository to copy content from. | [optional] 
**SourceRepositoryVersion** | Pointer to **string** | A URI of the repository version to copy content from. | [optional] 
**Digests** | Pointer to **[]interface{}** | A list of manifest digests to copy. | [optional] 
**MediaTypes** | Pointer to [**[]MediaTypesEnum**](MediaTypesEnum.md) | A list of media_types to copy. | [optional] 

## Methods

### NewManifestCopy

`func NewManifestCopy() *ManifestCopy`

NewManifestCopy instantiates a new ManifestCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestCopyWithDefaults

`func NewManifestCopyWithDefaults() *ManifestCopy`

NewManifestCopyWithDefaults instantiates a new ManifestCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRepository

`func (o *ManifestCopy) GetSourceRepository() string`

GetSourceRepository returns the SourceRepository field if non-nil, zero value otherwise.

### GetSourceRepositoryOk

`func (o *ManifestCopy) GetSourceRepositoryOk() (*string, bool)`

GetSourceRepositoryOk returns a tuple with the SourceRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepository

`func (o *ManifestCopy) SetSourceRepository(v string)`

SetSourceRepository sets SourceRepository field to given value.

### HasSourceRepository

`func (o *ManifestCopy) HasSourceRepository() bool`

HasSourceRepository returns a boolean if a field has been set.

### GetSourceRepositoryVersion

`func (o *ManifestCopy) GetSourceRepositoryVersion() string`

GetSourceRepositoryVersion returns the SourceRepositoryVersion field if non-nil, zero value otherwise.

### GetSourceRepositoryVersionOk

`func (o *ManifestCopy) GetSourceRepositoryVersionOk() (*string, bool)`

GetSourceRepositoryVersionOk returns a tuple with the SourceRepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepositoryVersion

`func (o *ManifestCopy) SetSourceRepositoryVersion(v string)`

SetSourceRepositoryVersion sets SourceRepositoryVersion field to given value.

### HasSourceRepositoryVersion

`func (o *ManifestCopy) HasSourceRepositoryVersion() bool`

HasSourceRepositoryVersion returns a boolean if a field has been set.

### GetDigests

`func (o *ManifestCopy) GetDigests() []interface{}`

GetDigests returns the Digests field if non-nil, zero value otherwise.

### GetDigestsOk

`func (o *ManifestCopy) GetDigestsOk() (*[]interface{}, bool)`

GetDigestsOk returns a tuple with the Digests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigests

`func (o *ManifestCopy) SetDigests(v []interface{})`

SetDigests sets Digests field to given value.

### HasDigests

`func (o *ManifestCopy) HasDigests() bool`

HasDigests returns a boolean if a field has been set.

### GetMediaTypes

`func (o *ManifestCopy) GetMediaTypes() []MediaTypesEnum`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *ManifestCopy) GetMediaTypesOk() (*[]MediaTypesEnum, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *ManifestCopy) SetMediaTypes(v []MediaTypesEnum)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *ManifestCopy) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


