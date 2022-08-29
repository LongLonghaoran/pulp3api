# ArtifactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**File** | **string** | The stored file. | 
**Size** | Pointer to **int32** | The size of the file in bytes. | [optional] 
**Md5** | Pointer to **NullableString** | The MD5 checksum of the file if available. | [optional] 
**Sha1** | Pointer to **NullableString** | The SHA-1 checksum of the file if available. | [optional] 
**Sha224** | Pointer to **NullableString** | The SHA-224 checksum of the file if available. | [optional] 
**Sha256** | Pointer to **NullableString** | The SHA-256 checksum of the file if available. | [optional] 
**Sha384** | Pointer to **NullableString** | The SHA-384 checksum of the file if available. | [optional] 
**Sha512** | Pointer to **NullableString** | The SHA-512 checksum of the file if available. | [optional] 

## Methods

### NewArtifactResponse

`func NewArtifactResponse(file string, ) *ArtifactResponse`

NewArtifactResponse instantiates a new ArtifactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactResponseWithDefaults

`func NewArtifactResponseWithDefaults() *ArtifactResponse`

NewArtifactResponseWithDefaults instantiates a new ArtifactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ArtifactResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ArtifactResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ArtifactResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ArtifactResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ArtifactResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ArtifactResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ArtifactResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ArtifactResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetFile

`func (o *ArtifactResponse) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ArtifactResponse) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ArtifactResponse) SetFile(v string)`

SetFile sets File field to given value.


### GetSize

`func (o *ArtifactResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *ArtifactResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *ArtifactResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *ArtifactResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *ArtifactResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *ArtifactResponse) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *ArtifactResponse) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSha1

`func (o *ArtifactResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ArtifactResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ArtifactResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ArtifactResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### SetSha1Nil

`func (o *ArtifactResponse) SetSha1Nil(b bool)`

 SetSha1Nil sets the value for Sha1 to be an explicit nil

### UnsetSha1
`func (o *ArtifactResponse) UnsetSha1()`

UnsetSha1 ensures that no value is present for Sha1, not even an explicit nil
### GetSha224

`func (o *ArtifactResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *ArtifactResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *ArtifactResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *ArtifactResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### SetSha224Nil

`func (o *ArtifactResponse) SetSha224Nil(b bool)`

 SetSha224Nil sets the value for Sha224 to be an explicit nil

### UnsetSha224
`func (o *ArtifactResponse) UnsetSha224()`

UnsetSha224 ensures that no value is present for Sha224, not even an explicit nil
### GetSha256

`func (o *ArtifactResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ArtifactResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ArtifactResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ArtifactResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *ArtifactResponse) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *ArtifactResponse) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetSha384

`func (o *ArtifactResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *ArtifactResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *ArtifactResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *ArtifactResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### SetSha384Nil

`func (o *ArtifactResponse) SetSha384Nil(b bool)`

 SetSha384Nil sets the value for Sha384 to be an explicit nil

### UnsetSha384
`func (o *ArtifactResponse) UnsetSha384()`

UnsetSha384 ensures that no value is present for Sha384, not even an explicit nil
### GetSha512

`func (o *ArtifactResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *ArtifactResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *ArtifactResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *ArtifactResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### SetSha512Nil

`func (o *ArtifactResponse) SetSha512Nil(b bool)`

 SetSha512Nil sets the value for Sha512 to be an explicit nil

### UnsetSha512
`func (o *ArtifactResponse) UnsetSha512()`

UnsetSha512 ensures that no value is present for Sha512, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


