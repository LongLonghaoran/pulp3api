# DebReleaseFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 
**Codename** | Pointer to **string** | Codename of the release, i.e. \&quot;buster\&quot;. | [optional] 
**Suite** | Pointer to **string** | Suite of the release, i.e. \&quot;stable\&quot;. | [optional] 
**Distribution** | **string** | Distribution of the release, i.e. \&quot;stable/updates\&quot;. | 
**RelativePath** | Pointer to **string** | Path of file relative to url. | [optional] 

## Methods

### NewDebReleaseFileResponse

`func NewDebReleaseFileResponse(artifacts map[string]interface{}, distribution string, ) *DebReleaseFileResponse`

NewDebReleaseFileResponse instantiates a new DebReleaseFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseFileResponseWithDefaults

`func NewDebReleaseFileResponseWithDefaults() *DebReleaseFileResponse`

NewDebReleaseFileResponseWithDefaults instantiates a new DebReleaseFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebReleaseFileResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebReleaseFileResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebReleaseFileResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebReleaseFileResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebReleaseFileResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebReleaseFileResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebReleaseFileResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebReleaseFileResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifacts

`func (o *DebReleaseFileResponse) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebReleaseFileResponse) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebReleaseFileResponse) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetCodename

`func (o *DebReleaseFileResponse) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseFileResponse) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseFileResponse) SetCodename(v string)`

SetCodename sets Codename field to given value.

### HasCodename

`func (o *DebReleaseFileResponse) HasCodename() bool`

HasCodename returns a boolean if a field has been set.

### GetSuite

`func (o *DebReleaseFileResponse) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseFileResponse) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseFileResponse) SetSuite(v string)`

SetSuite sets Suite field to given value.

### HasSuite

`func (o *DebReleaseFileResponse) HasSuite() bool`

HasSuite returns a boolean if a field has been set.

### GetDistribution

`func (o *DebReleaseFileResponse) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseFileResponse) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseFileResponse) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetRelativePath

`func (o *DebReleaseFileResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebReleaseFileResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebReleaseFileResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebReleaseFileResponse) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


