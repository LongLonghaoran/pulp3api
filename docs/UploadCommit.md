# UploadCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256** | **string** | The expected sha256 checksum for the file. | 

## Methods

### NewUploadCommit

`func NewUploadCommit(sha256 string, ) *UploadCommit`

NewUploadCommit instantiates a new UploadCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadCommitWithDefaults

`func NewUploadCommitWithDefaults() *UploadCommit`

NewUploadCommitWithDefaults instantiates a new UploadCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256

`func (o *UploadCommit) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *UploadCommit) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *UploadCommit) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


