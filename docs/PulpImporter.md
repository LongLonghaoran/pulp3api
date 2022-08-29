# PulpImporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the Importer. | 
**RepoMapping** | Pointer to **map[string]string** | Mapping of repo names in an export file to the repo names in Pulp. For example, if the export has a repo named &#39;foo&#39; and the repo to import content into was &#39;bar&#39;, the mapping would be \&quot;{&#39;foo&#39;: &#39;bar&#39;}\&quot;. | [optional] 

## Methods

### NewPulpImporter

`func NewPulpImporter(name string, ) *PulpImporter`

NewPulpImporter instantiates a new PulpImporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImporterWithDefaults

`func NewPulpImporterWithDefaults() *PulpImporter`

NewPulpImporterWithDefaults instantiates a new PulpImporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PulpImporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PulpImporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PulpImporter) SetName(v string)`

SetName sets Name field to given value.


### GetRepoMapping

`func (o *PulpImporter) GetRepoMapping() map[string]string`

GetRepoMapping returns the RepoMapping field if non-nil, zero value otherwise.

### GetRepoMappingOk

`func (o *PulpImporter) GetRepoMappingOk() (*map[string]string, bool)`

GetRepoMappingOk returns a tuple with the RepoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoMapping

`func (o *PulpImporter) SetRepoMapping(v map[string]string)`

SetRepoMapping sets RepoMapping field to given value.

### HasRepoMapping

`func (o *PulpImporter) HasRepoMapping() bool`

HasRepoMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


