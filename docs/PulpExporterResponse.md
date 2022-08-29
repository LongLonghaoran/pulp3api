# PulpExporterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Unique name of the file system exporter. | 
**Path** | **string** | File system directory to store exported tar.gzs. | 
**Repositories** | **[]string** |  | 
**LastExport** | Pointer to **NullableString** | Last attempted export for this PulpExporter | [optional] 

## Methods

### NewPulpExporterResponse

`func NewPulpExporterResponse(name string, path string, repositories []string, ) *PulpExporterResponse`

NewPulpExporterResponse instantiates a new PulpExporterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpExporterResponseWithDefaults

`func NewPulpExporterResponseWithDefaults() *PulpExporterResponse`

NewPulpExporterResponseWithDefaults instantiates a new PulpExporterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PulpExporterResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PulpExporterResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PulpExporterResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PulpExporterResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PulpExporterResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PulpExporterResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PulpExporterResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PulpExporterResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *PulpExporterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PulpExporterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PulpExporterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *PulpExporterResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PulpExporterResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PulpExporterResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetRepositories

`func (o *PulpExporterResponse) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PulpExporterResponse) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PulpExporterResponse) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetLastExport

`func (o *PulpExporterResponse) GetLastExport() string`

GetLastExport returns the LastExport field if non-nil, zero value otherwise.

### GetLastExportOk

`func (o *PulpExporterResponse) GetLastExportOk() (*string, bool)`

GetLastExportOk returns a tuple with the LastExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExport

`func (o *PulpExporterResponse) SetLastExport(v string)`

SetLastExport sets LastExport field to given value.

### HasLastExport

`func (o *PulpExporterResponse) HasLastExport() bool`

HasLastExport returns a boolean if a field has been set.

### SetLastExportNil

`func (o *PulpExporterResponse) SetLastExportNil(b bool)`

 SetLastExportNil sets the value for LastExport to be an explicit nil

### UnsetLastExport
`func (o *PulpExporterResponse) UnsetLastExport()`

UnsetLastExport ensures that no value is present for LastExport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


