# PulpImportCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Toc** | Pointer to [**EvaluationResponse**](EvaluationResponse.md) | Evaluation of proposed &#39;toc&#39; file for PulpImport | [optional] 
**Path** | Pointer to [**EvaluationResponse**](EvaluationResponse.md) | Evaluation of proposed &#39;path&#39; file for PulpImport | [optional] 
**RepoMapping** | Pointer to [**EvaluationResponse**](EvaluationResponse.md) | Evaluation of proposed &#39;repo_mapping&#39; file for PulpImport | [optional] 

## Methods

### NewPulpImportCheckResponse

`func NewPulpImportCheckResponse() *PulpImportCheckResponse`

NewPulpImportCheckResponse instantiates a new PulpImportCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImportCheckResponseWithDefaults

`func NewPulpImportCheckResponseWithDefaults() *PulpImportCheckResponse`

NewPulpImportCheckResponseWithDefaults instantiates a new PulpImportCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToc

`func (o *PulpImportCheckResponse) GetToc() EvaluationResponse`

GetToc returns the Toc field if non-nil, zero value otherwise.

### GetTocOk

`func (o *PulpImportCheckResponse) GetTocOk() (*EvaluationResponse, bool)`

GetTocOk returns a tuple with the Toc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToc

`func (o *PulpImportCheckResponse) SetToc(v EvaluationResponse)`

SetToc sets Toc field to given value.

### HasToc

`func (o *PulpImportCheckResponse) HasToc() bool`

HasToc returns a boolean if a field has been set.

### GetPath

`func (o *PulpImportCheckResponse) GetPath() EvaluationResponse`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PulpImportCheckResponse) GetPathOk() (*EvaluationResponse, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PulpImportCheckResponse) SetPath(v EvaluationResponse)`

SetPath sets Path field to given value.

### HasPath

`func (o *PulpImportCheckResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepoMapping

`func (o *PulpImportCheckResponse) GetRepoMapping() EvaluationResponse`

GetRepoMapping returns the RepoMapping field if non-nil, zero value otherwise.

### GetRepoMappingOk

`func (o *PulpImportCheckResponse) GetRepoMappingOk() (*EvaluationResponse, bool)`

GetRepoMappingOk returns a tuple with the RepoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoMapping

`func (o *PulpImportCheckResponse) SetRepoMapping(v EvaluationResponse)`

SetRepoMapping sets RepoMapping field to given value.

### HasRepoMapping

`func (o *PulpImportCheckResponse) HasRepoMapping() bool`

HasRepoMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


