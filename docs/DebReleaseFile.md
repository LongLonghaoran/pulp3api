# DebReleaseFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 
**Codename** | Pointer to **string** | Codename of the release, i.e. \&quot;buster\&quot;. | [optional] 
**Suite** | Pointer to **string** | Suite of the release, i.e. \&quot;stable\&quot;. | [optional] 
**Distribution** | **string** | Distribution of the release, i.e. \&quot;stable/updates\&quot;. | 
**RelativePath** | Pointer to **string** | Path of file relative to url. | [optional] 

## Methods

### NewDebReleaseFile

`func NewDebReleaseFile(artifacts map[string]interface{}, distribution string, ) *DebReleaseFile`

NewDebReleaseFile instantiates a new DebReleaseFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseFileWithDefaults

`func NewDebReleaseFileWithDefaults() *DebReleaseFile`

NewDebReleaseFileWithDefaults instantiates a new DebReleaseFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *DebReleaseFile) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebReleaseFile) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebReleaseFile) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetCodename

`func (o *DebReleaseFile) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseFile) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseFile) SetCodename(v string)`

SetCodename sets Codename field to given value.

### HasCodename

`func (o *DebReleaseFile) HasCodename() bool`

HasCodename returns a boolean if a field has been set.

### GetSuite

`func (o *DebReleaseFile) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseFile) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseFile) SetSuite(v string)`

SetSuite sets Suite field to given value.

### HasSuite

`func (o *DebReleaseFile) HasSuite() bool`

HasSuite returns a boolean if a field has been set.

### GetDistribution

`func (o *DebReleaseFile) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseFile) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseFile) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetRelativePath

`func (o *DebReleaseFile) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebReleaseFile) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebReleaseFile) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebReleaseFile) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


