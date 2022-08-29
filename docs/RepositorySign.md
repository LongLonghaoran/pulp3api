# RepositorySign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManifestSigningService** | Pointer to **NullableString** | A signing service to sign with. This will override a signing service set on the repo. | [optional] 
**FutureBasePath** | Pointer to **string** | Future base path content will be distributed at for sync repos | [optional] 
**TagsList** | Pointer to **[]interface{}** | A list of tags to sign. | [optional] 

## Methods

### NewRepositorySign

`func NewRepositorySign() *RepositorySign`

NewRepositorySign instantiates a new RepositorySign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySignWithDefaults

`func NewRepositorySignWithDefaults() *RepositorySign`

NewRepositorySignWithDefaults instantiates a new RepositorySign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifestSigningService

`func (o *RepositorySign) GetManifestSigningService() string`

GetManifestSigningService returns the ManifestSigningService field if non-nil, zero value otherwise.

### GetManifestSigningServiceOk

`func (o *RepositorySign) GetManifestSigningServiceOk() (*string, bool)`

GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSigningService

`func (o *RepositorySign) SetManifestSigningService(v string)`

SetManifestSigningService sets ManifestSigningService field to given value.

### HasManifestSigningService

`func (o *RepositorySign) HasManifestSigningService() bool`

HasManifestSigningService returns a boolean if a field has been set.

### SetManifestSigningServiceNil

`func (o *RepositorySign) SetManifestSigningServiceNil(b bool)`

 SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil

### UnsetManifestSigningService
`func (o *RepositorySign) UnsetManifestSigningService()`

UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
### GetFutureBasePath

`func (o *RepositorySign) GetFutureBasePath() string`

GetFutureBasePath returns the FutureBasePath field if non-nil, zero value otherwise.

### GetFutureBasePathOk

`func (o *RepositorySign) GetFutureBasePathOk() (*string, bool)`

GetFutureBasePathOk returns a tuple with the FutureBasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureBasePath

`func (o *RepositorySign) SetFutureBasePath(v string)`

SetFutureBasePath sets FutureBasePath field to given value.

### HasFutureBasePath

`func (o *RepositorySign) HasFutureBasePath() bool`

HasFutureBasePath returns a boolean if a field has been set.

### GetTagsList

`func (o *RepositorySign) GetTagsList() []interface{}`

GetTagsList returns the TagsList field if non-nil, zero value otherwise.

### GetTagsListOk

`func (o *RepositorySign) GetTagsListOk() (*[]interface{}, bool)`

GetTagsListOk returns a tuple with the TagsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsList

`func (o *RepositorySign) SetTagsList(v []interface{})`

SetTagsList sets TagsList field to given value.

### HasTagsList

`func (o *RepositorySign) HasTagsList() bool`

HasTagsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


