# PatchedFilesystemExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name of the file system exporter. | [optional] 
**Path** | Pointer to **string** | File system location to export to. | [optional] 
**Method** | Pointer to [**MethodEnum**](MethodEnum.md) | Method of exporting | [optional] 

## Methods

### NewPatchedFilesystemExporter

`func NewPatchedFilesystemExporter() *PatchedFilesystemExporter`

NewPatchedFilesystemExporter instantiates a new PatchedFilesystemExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFilesystemExporterWithDefaults

`func NewPatchedFilesystemExporterWithDefaults() *PatchedFilesystemExporter`

NewPatchedFilesystemExporterWithDefaults instantiates a new PatchedFilesystemExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedFilesystemExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFilesystemExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFilesystemExporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFilesystemExporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *PatchedFilesystemExporter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchedFilesystemExporter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchedFilesystemExporter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchedFilesystemExporter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *PatchedFilesystemExporter) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PatchedFilesystemExporter) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PatchedFilesystemExporter) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PatchedFilesystemExporter) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


