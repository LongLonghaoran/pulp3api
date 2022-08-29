# TaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the task. The possible values include: &#39;waiting&#39;, &#39;skipped&#39;, &#39;running&#39;, &#39;completed&#39;, &#39;failed&#39;, &#39;canceled&#39; and &#39;canceling&#39;. | [optional] [readonly] 
**Name** | **string** | The name of task. | 
**LoggingCid** | **string** | The logging correlation id associated with this task | 
**StartedAt** | Pointer to **time.Time** | Timestamp of the when this task started execution. | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Timestamp of the when this task stopped execution. | [optional] [readonly] 
**Error** | Pointer to **map[string]map[string]interface{}** | A JSON Object of a fatal error encountered during the execution of this task. | [optional] [readonly] 
**Worker** | Pointer to **string** | The worker associated with this task. This field is empty if a worker is not yet assigned. | [optional] [readonly] 
**ParentTask** | Pointer to **string** | The parent task that spawned this task. | [optional] [readonly] 
**ChildTasks** | Pointer to **[]string** | Any tasks spawned by this task. | [optional] [readonly] 
**TaskGroup** | Pointer to **string** | The task group that this task is a member of. | [optional] [readonly] 
**ProgressReports** | Pointer to [**[]ProgressReportResponse**](ProgressReportResponse.md) |  | [optional] [readonly] 
**CreatedResources** | Pointer to **[]string** | Resources created by this task. | [optional] [readonly] 
**ReservedResourcesRecord** | Pointer to **[]string** | A list of resources required by that task. | [optional] [readonly] 

## Methods

### NewTaskResponse

`func NewTaskResponse(name string, loggingCid string, ) *TaskResponse`

NewTaskResponse instantiates a new TaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseWithDefaults

`func NewTaskResponseWithDefaults() *TaskResponse`

NewTaskResponseWithDefaults instantiates a new TaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *TaskResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *TaskResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *TaskResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *TaskResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *TaskResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *TaskResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *TaskResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *TaskResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetState

`func (o *TaskResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TaskResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *TaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLoggingCid

`func (o *TaskResponse) GetLoggingCid() string`

GetLoggingCid returns the LoggingCid field if non-nil, zero value otherwise.

### GetLoggingCidOk

`func (o *TaskResponse) GetLoggingCidOk() (*string, bool)`

GetLoggingCidOk returns a tuple with the LoggingCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCid

`func (o *TaskResponse) SetLoggingCid(v string)`

SetLoggingCid sets LoggingCid field to given value.


### GetStartedAt

`func (o *TaskResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TaskResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TaskResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TaskResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *TaskResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *TaskResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *TaskResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *TaskResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetError

`func (o *TaskResponse) GetError() map[string]map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskResponse) GetErrorOk() (*map[string]map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskResponse) SetError(v map[string]map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *TaskResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWorker

`func (o *TaskResponse) GetWorker() string`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *TaskResponse) GetWorkerOk() (*string, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *TaskResponse) SetWorker(v string)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *TaskResponse) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetParentTask

`func (o *TaskResponse) GetParentTask() string`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *TaskResponse) GetParentTaskOk() (*string, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *TaskResponse) SetParentTask(v string)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *TaskResponse) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### GetChildTasks

`func (o *TaskResponse) GetChildTasks() []string`

GetChildTasks returns the ChildTasks field if non-nil, zero value otherwise.

### GetChildTasksOk

`func (o *TaskResponse) GetChildTasksOk() (*[]string, bool)`

GetChildTasksOk returns a tuple with the ChildTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasks

`func (o *TaskResponse) SetChildTasks(v []string)`

SetChildTasks sets ChildTasks field to given value.

### HasChildTasks

`func (o *TaskResponse) HasChildTasks() bool`

HasChildTasks returns a boolean if a field has been set.

### GetTaskGroup

`func (o *TaskResponse) GetTaskGroup() string`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *TaskResponse) GetTaskGroupOk() (*string, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *TaskResponse) SetTaskGroup(v string)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *TaskResponse) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetProgressReports

`func (o *TaskResponse) GetProgressReports() []ProgressReportResponse`

GetProgressReports returns the ProgressReports field if non-nil, zero value otherwise.

### GetProgressReportsOk

`func (o *TaskResponse) GetProgressReportsOk() (*[]ProgressReportResponse, bool)`

GetProgressReportsOk returns a tuple with the ProgressReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressReports

`func (o *TaskResponse) SetProgressReports(v []ProgressReportResponse)`

SetProgressReports sets ProgressReports field to given value.

### HasProgressReports

`func (o *TaskResponse) HasProgressReports() bool`

HasProgressReports returns a boolean if a field has been set.

### GetCreatedResources

`func (o *TaskResponse) GetCreatedResources() []string`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *TaskResponse) GetCreatedResourcesOk() (*[]string, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *TaskResponse) SetCreatedResources(v []string)`

SetCreatedResources sets CreatedResources field to given value.

### HasCreatedResources

`func (o *TaskResponse) HasCreatedResources() bool`

HasCreatedResources returns a boolean if a field has been set.

### GetReservedResourcesRecord

`func (o *TaskResponse) GetReservedResourcesRecord() []string`

GetReservedResourcesRecord returns the ReservedResourcesRecord field if non-nil, zero value otherwise.

### GetReservedResourcesRecordOk

`func (o *TaskResponse) GetReservedResourcesRecordOk() (*[]string, bool)`

GetReservedResourcesRecordOk returns a tuple with the ReservedResourcesRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedResourcesRecord

`func (o *TaskResponse) SetReservedResourcesRecord(v []string)`

SetReservedResourcesRecord sets ReservedResourcesRecord field to given value.

### HasReservedResourcesRecord

`func (o *TaskResponse) HasReservedResourcesRecord() bool`

HasReservedResourcesRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


