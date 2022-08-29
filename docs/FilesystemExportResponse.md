# FilesystemExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Task** | Pointer to **NullableString** | A URI of the task that ran the Export. | [optional] 
**ExportedResources** | Pointer to **[]string** | Resources that were exported. | [optional] [readonly] 
**Params** | Pointer to **map[string]interface{}** | Any additional parameters that were used to create the export. | [optional] [readonly] 

## Methods

### NewFilesystemExportResponse

`func NewFilesystemExportResponse() *FilesystemExportResponse`

NewFilesystemExportResponse instantiates a new FilesystemExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemExportResponseWithDefaults

`func NewFilesystemExportResponseWithDefaults() *FilesystemExportResponse`

NewFilesystemExportResponseWithDefaults instantiates a new FilesystemExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *FilesystemExportResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *FilesystemExportResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *FilesystemExportResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *FilesystemExportResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *FilesystemExportResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *FilesystemExportResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *FilesystemExportResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *FilesystemExportResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetTask

`func (o *FilesystemExportResponse) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *FilesystemExportResponse) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *FilesystemExportResponse) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *FilesystemExportResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *FilesystemExportResponse) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *FilesystemExportResponse) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetExportedResources

`func (o *FilesystemExportResponse) GetExportedResources() []string`

GetExportedResources returns the ExportedResources field if non-nil, zero value otherwise.

### GetExportedResourcesOk

`func (o *FilesystemExportResponse) GetExportedResourcesOk() (*[]string, bool)`

GetExportedResourcesOk returns a tuple with the ExportedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedResources

`func (o *FilesystemExportResponse) SetExportedResources(v []string)`

SetExportedResources sets ExportedResources field to given value.

### HasExportedResources

`func (o *FilesystemExportResponse) HasExportedResources() bool`

HasExportedResources returns a boolean if a field has been set.

### GetParams

`func (o *FilesystemExportResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FilesystemExportResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FilesystemExportResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *FilesystemExportResponse) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


