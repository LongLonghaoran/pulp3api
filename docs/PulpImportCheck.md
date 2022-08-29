# PulpImportCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path to export-tar-gz that will be imported. | [optional] 
**Toc** | Pointer to **string** | Path to a table-of-contents file describing chunks to be validated, reassembled, and imported. | [optional] 
**RepoMapping** | Pointer to **string** | Mapping of repo names in an export file to the repo names in Pulp. For example, if the export has a repo named &#39;foo&#39; and the repo to import content into was &#39;bar&#39;, the mapping would be \&quot;{&#39;foo&#39;: &#39;bar&#39;}\&quot;. | [optional] 

## Methods

### NewPulpImportCheck

`func NewPulpImportCheck() *PulpImportCheck`

NewPulpImportCheck instantiates a new PulpImportCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImportCheckWithDefaults

`func NewPulpImportCheckWithDefaults() *PulpImportCheck`

NewPulpImportCheckWithDefaults instantiates a new PulpImportCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PulpImportCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PulpImportCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PulpImportCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PulpImportCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetToc

`func (o *PulpImportCheck) GetToc() string`

GetToc returns the Toc field if non-nil, zero value otherwise.

### GetTocOk

`func (o *PulpImportCheck) GetTocOk() (*string, bool)`

GetTocOk returns a tuple with the Toc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToc

`func (o *PulpImportCheck) SetToc(v string)`

SetToc sets Toc field to given value.

### HasToc

`func (o *PulpImportCheck) HasToc() bool`

HasToc returns a boolean if a field has been set.

### GetRepoMapping

`func (o *PulpImportCheck) GetRepoMapping() string`

GetRepoMapping returns the RepoMapping field if non-nil, zero value otherwise.

### GetRepoMappingOk

`func (o *PulpImportCheck) GetRepoMappingOk() (*string, bool)`

GetRepoMappingOk returns a tuple with the RepoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoMapping

`func (o *PulpImportCheck) SetRepoMapping(v string)`

SetRepoMapping sets RepoMapping field to given value.

### HasRepoMapping

`func (o *PulpImportCheck) HasRepoMapping() bool`

HasRepoMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


