# PatchedPulpExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name of the file system exporter. | [optional] 
**Path** | Pointer to **string** | File system directory to store exported tar.gzs. | [optional] 
**Repositories** | Pointer to **[]string** |  | [optional] 
**LastExport** | Pointer to **NullableString** | Last attempted export for this PulpExporter | [optional] 

## Methods

### NewPatchedPulpExporter

`func NewPatchedPulpExporter() *PatchedPulpExporter`

NewPatchedPulpExporter instantiates a new PatchedPulpExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPulpExporterWithDefaults

`func NewPatchedPulpExporterWithDefaults() *PatchedPulpExporter`

NewPatchedPulpExporterWithDefaults instantiates a new PatchedPulpExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPulpExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPulpExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPulpExporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPulpExporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *PatchedPulpExporter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchedPulpExporter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchedPulpExporter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchedPulpExporter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepositories

`func (o *PatchedPulpExporter) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PatchedPulpExporter) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PatchedPulpExporter) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PatchedPulpExporter) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetLastExport

`func (o *PatchedPulpExporter) GetLastExport() string`

GetLastExport returns the LastExport field if non-nil, zero value otherwise.

### GetLastExportOk

`func (o *PatchedPulpExporter) GetLastExportOk() (*string, bool)`

GetLastExportOk returns a tuple with the LastExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExport

`func (o *PatchedPulpExporter) SetLastExport(v string)`

SetLastExport sets LastExport field to given value.

### HasLastExport

`func (o *PatchedPulpExporter) HasLastExport() bool`

HasLastExport returns a boolean if a field has been set.

### SetLastExportNil

`func (o *PatchedPulpExporter) SetLastExportNil(b bool)`

 SetLastExportNil sets the value for LastExport to be an explicit nil

### UnsetLastExport
`func (o *PatchedPulpExporter) UnsetLastExport()`

UnsetLastExport ensures that no value is present for LastExport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


