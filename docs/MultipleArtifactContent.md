# MultipleArtifactContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | 

## Methods

### NewMultipleArtifactContent

`func NewMultipleArtifactContent(artifacts map[string]interface{}, ) *MultipleArtifactContent`

NewMultipleArtifactContent instantiates a new MultipleArtifactContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleArtifactContentWithDefaults

`func NewMultipleArtifactContentWithDefaults() *MultipleArtifactContent`

NewMultipleArtifactContentWithDefaults instantiates a new MultipleArtifactContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *MultipleArtifactContent) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *MultipleArtifactContent) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *MultipleArtifactContent) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


