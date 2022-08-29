# PulpExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to **NullableString** | A URI of the task that ran the Export. | [optional] 
**Full** | Pointer to **bool** | Do a Full (true) or Incremental (false) export. | [optional] [default to true]
**DryRun** | Pointer to **bool** | Generate report on what would be exported and disk-space required. | [optional] [default to false]
**Versions** | Pointer to **[]string** | List of explicit repo-version hrefs to export (replaces current_version). | [optional] 
**ChunkSize** | Pointer to **string** | Chunk export-tarfile into pieces of chunk_size bytes. Recognizes units of B/KB/MB/GB/TB. A chunk has a maximum size of 1TB. | [optional] 
**StartVersions** | Pointer to **[]string** | List of explicit last-exported-repo-version hrefs (replaces last_export). | [optional] 

## Methods

### NewPulpExport

`func NewPulpExport() *PulpExport`

NewPulpExport instantiates a new PulpExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpExportWithDefaults

`func NewPulpExportWithDefaults() *PulpExport`

NewPulpExportWithDefaults instantiates a new PulpExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *PulpExport) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PulpExport) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PulpExport) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *PulpExport) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *PulpExport) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *PulpExport) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetFull

`func (o *PulpExport) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *PulpExport) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *PulpExport) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *PulpExport) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetDryRun

`func (o *PulpExport) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PulpExport) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PulpExport) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PulpExport) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVersions

`func (o *PulpExport) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PulpExport) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PulpExport) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PulpExport) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetChunkSize

`func (o *PulpExport) GetChunkSize() string`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *PulpExport) GetChunkSizeOk() (*string, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *PulpExport) SetChunkSize(v string)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *PulpExport) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### GetStartVersions

`func (o *PulpExport) GetStartVersions() []string`

GetStartVersions returns the StartVersions field if non-nil, zero value otherwise.

### GetStartVersionsOk

`func (o *PulpExport) GetStartVersionsOk() (*[]string, bool)`

GetStartVersionsOk returns a tuple with the StartVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVersions

`func (o *PulpExport) SetStartVersions(v []string)`

SetStartVersions sets StartVersions field to given value.

### HasStartVersions

`func (o *PulpExport) HasStartVersions() bool`

HasStartVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


