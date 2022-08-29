# PulpExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Task** | Pointer to **NullableString** | A URI of the task that ran the Export. | [optional] 
**ExportedResources** | Pointer to **[]string** | Resources that were exported. | [optional] [readonly] 
**Params** | Pointer to **map[string]interface{}** | Any additional parameters that were used to create the export. | [optional] [readonly] 
**OutputFileInfo** | Pointer to **map[string]interface{}** | Dictionary of filename: sha256hash entries for export-output-file(s) | [optional] [readonly] 
**TocInfo** | Pointer to **map[string]interface{}** | Filename and sha256-checksum of table-of-contents for this export | [optional] [readonly] 

## Methods

### NewPulpExportResponse

`func NewPulpExportResponse() *PulpExportResponse`

NewPulpExportResponse instantiates a new PulpExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpExportResponseWithDefaults

`func NewPulpExportResponseWithDefaults() *PulpExportResponse`

NewPulpExportResponseWithDefaults instantiates a new PulpExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PulpExportResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PulpExportResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PulpExportResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PulpExportResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PulpExportResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PulpExportResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PulpExportResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PulpExportResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetTask

`func (o *PulpExportResponse) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PulpExportResponse) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PulpExportResponse) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *PulpExportResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *PulpExportResponse) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *PulpExportResponse) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetExportedResources

`func (o *PulpExportResponse) GetExportedResources() []string`

GetExportedResources returns the ExportedResources field if non-nil, zero value otherwise.

### GetExportedResourcesOk

`func (o *PulpExportResponse) GetExportedResourcesOk() (*[]string, bool)`

GetExportedResourcesOk returns a tuple with the ExportedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedResources

`func (o *PulpExportResponse) SetExportedResources(v []string)`

SetExportedResources sets ExportedResources field to given value.

### HasExportedResources

`func (o *PulpExportResponse) HasExportedResources() bool`

HasExportedResources returns a boolean if a field has been set.

### GetParams

`func (o *PulpExportResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PulpExportResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PulpExportResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *PulpExportResponse) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetOutputFileInfo

`func (o *PulpExportResponse) GetOutputFileInfo() map[string]interface{}`

GetOutputFileInfo returns the OutputFileInfo field if non-nil, zero value otherwise.

### GetOutputFileInfoOk

`func (o *PulpExportResponse) GetOutputFileInfoOk() (*map[string]interface{}, bool)`

GetOutputFileInfoOk returns a tuple with the OutputFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileInfo

`func (o *PulpExportResponse) SetOutputFileInfo(v map[string]interface{})`

SetOutputFileInfo sets OutputFileInfo field to given value.

### HasOutputFileInfo

`func (o *PulpExportResponse) HasOutputFileInfo() bool`

HasOutputFileInfo returns a boolean if a field has been set.

### GetTocInfo

`func (o *PulpExportResponse) GetTocInfo() map[string]interface{}`

GetTocInfo returns the TocInfo field if non-nil, zero value otherwise.

### GetTocInfoOk

`func (o *PulpExportResponse) GetTocInfoOk() (*map[string]interface{}, bool)`

GetTocInfoOk returns a tuple with the TocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTocInfo

`func (o *PulpExportResponse) SetTocInfo(v map[string]interface{})`

SetTocInfo sets TocInfo field to given value.

### HasTocInfo

`func (o *PulpExportResponse) HasTocInfo() bool`

HasTocInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


