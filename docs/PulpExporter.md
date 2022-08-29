# PulpExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the file system exporter. | 
**Path** | **string** | File system directory to store exported tar.gzs. | 
**Repositories** | **[]string** |  | 
**LastExport** | Pointer to **NullableString** | Last attempted export for this PulpExporter | [optional] 

## Methods

### NewPulpExporter

`func NewPulpExporter(name string, path string, repositories []string, ) *PulpExporter`

NewPulpExporter instantiates a new PulpExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpExporterWithDefaults

`func NewPulpExporterWithDefaults() *PulpExporter`

NewPulpExporterWithDefaults instantiates a new PulpExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PulpExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PulpExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PulpExporter) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *PulpExporter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PulpExporter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PulpExporter) SetPath(v string)`

SetPath sets Path field to given value.


### GetRepositories

`func (o *PulpExporter) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PulpExporter) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PulpExporter) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetLastExport

`func (o *PulpExporter) GetLastExport() string`

GetLastExport returns the LastExport field if non-nil, zero value otherwise.

### GetLastExportOk

`func (o *PulpExporter) GetLastExportOk() (*string, bool)`

GetLastExportOk returns a tuple with the LastExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExport

`func (o *PulpExporter) SetLastExport(v string)`

SetLastExport sets LastExport field to given value.

### HasLastExport

`func (o *PulpExporter) HasLastExport() bool`

HasLastExport returns a boolean if a field has been set.

### SetLastExportNil

`func (o *PulpExporter) SetLastExportNil(b bool)`

 SetLastExportNil sets the value for LastExport to be an explicit nil

### UnsetLastExport
`func (o *PulpExporter) UnsetLastExport()`

UnsetLastExport ensures that no value is present for LastExport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


