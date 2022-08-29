# FilesystemExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to **NullableString** | A URI of the task that ran the Export. | [optional] 
**Publication** | Pointer to **string** | A URI of the publication to be exported. | [optional] 
**RepositoryVersion** | Pointer to **string** | A URI of the repository version export. | [optional] 

## Methods

### NewFilesystemExport

`func NewFilesystemExport() *FilesystemExport`

NewFilesystemExport instantiates a new FilesystemExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemExportWithDefaults

`func NewFilesystemExportWithDefaults() *FilesystemExport`

NewFilesystemExportWithDefaults instantiates a new FilesystemExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *FilesystemExport) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *FilesystemExport) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *FilesystemExport) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *FilesystemExport) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *FilesystemExport) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *FilesystemExport) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetPublication

`func (o *FilesystemExport) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *FilesystemExport) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *FilesystemExport) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *FilesystemExport) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *FilesystemExport) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *FilesystemExport) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *FilesystemExport) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *FilesystemExport) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


