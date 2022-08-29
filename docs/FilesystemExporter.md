# FilesystemExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the file system exporter. | 
**Path** | **string** | File system location to export to. | 
**Method** | Pointer to [**MethodEnum**](MethodEnum.md) | Method of exporting | [optional] 

## Methods

### NewFilesystemExporter

`func NewFilesystemExporter(name string, path string, ) *FilesystemExporter`

NewFilesystemExporter instantiates a new FilesystemExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemExporterWithDefaults

`func NewFilesystemExporterWithDefaults() *FilesystemExporter`

NewFilesystemExporterWithDefaults instantiates a new FilesystemExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FilesystemExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesystemExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesystemExporter) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *FilesystemExporter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemExporter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemExporter) SetPath(v string)`

SetPath sets Path field to given value.


### GetMethod

`func (o *FilesystemExporter) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FilesystemExporter) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FilesystemExporter) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FilesystemExporter) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


