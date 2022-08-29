# DebInstallerFileIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 
**Release** | **string** | Release this index file belongs to. | 
**Component** | **string** | Component of the component - architecture combination. | 
**Architecture** | **string** | Architecture of the component - architecture combination. | 
**RelativePath** | Pointer to **string** | Path of directory containing MD5SUMS and SHA256SUMS relative to url. | [optional] 

## Methods

### NewDebInstallerFileIndexResponse

`func NewDebInstallerFileIndexResponse(artifacts map[string]interface{}, release string, component string, architecture string, ) *DebInstallerFileIndexResponse`

NewDebInstallerFileIndexResponse instantiates a new DebInstallerFileIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebInstallerFileIndexResponseWithDefaults

`func NewDebInstallerFileIndexResponseWithDefaults() *DebInstallerFileIndexResponse`

NewDebInstallerFileIndexResponseWithDefaults instantiates a new DebInstallerFileIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebInstallerFileIndexResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebInstallerFileIndexResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebInstallerFileIndexResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebInstallerFileIndexResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebInstallerFileIndexResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebInstallerFileIndexResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebInstallerFileIndexResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebInstallerFileIndexResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArtifacts

`func (o *DebInstallerFileIndexResponse) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebInstallerFileIndexResponse) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebInstallerFileIndexResponse) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetRelease

`func (o *DebInstallerFileIndexResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *DebInstallerFileIndexResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *DebInstallerFileIndexResponse) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetComponent

`func (o *DebInstallerFileIndexResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebInstallerFileIndexResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebInstallerFileIndexResponse) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetArchitecture

`func (o *DebInstallerFileIndexResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebInstallerFileIndexResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebInstallerFileIndexResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetRelativePath

`func (o *DebInstallerFileIndexResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebInstallerFileIndexResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebInstallerFileIndexResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebInstallerFileIndexResponse) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


