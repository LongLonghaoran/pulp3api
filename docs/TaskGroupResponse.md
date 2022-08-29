# TaskGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | A description of the task group. | 
**AllTasksDispatched** | **bool** | Whether all tasks have been spawned for this task group. | 
**Waiting** | Pointer to **int32** | Number of tasks in the &#39;waiting&#39; state | [optional] [readonly] 
**Skipped** | Pointer to **int32** | Number of tasks in the &#39;skipped&#39; state | [optional] [readonly] 
**Running** | Pointer to **int32** | Number of tasks in the &#39;running&#39; state | [optional] [readonly] 
**Completed** | Pointer to **int32** | Number of tasks in the &#39;completed&#39; state | [optional] [readonly] 
**Canceled** | Pointer to **int32** | Number of tasks in the &#39;canceled&#39; state | [optional] [readonly] 
**Failed** | Pointer to **int32** | Number of tasks in the &#39;failed&#39; state | [optional] [readonly] 
**Canceling** | Pointer to **int32** | Number of tasks in the &#39;canceling&#39; state | [optional] [readonly] 
**GroupProgressReports** | Pointer to [**[]GroupProgressReportResponse**](GroupProgressReportResponse.md) |  | [optional] [readonly] 
**Tasks** | Pointer to [**[]MinimalTaskResponse**](MinimalTaskResponse.md) |  | [optional] [readonly] 

## Methods

### NewTaskGroupResponse

`func NewTaskGroupResponse(description string, allTasksDispatched bool, ) *TaskGroupResponse`

NewTaskGroupResponse instantiates a new TaskGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGroupResponseWithDefaults

`func NewTaskGroupResponseWithDefaults() *TaskGroupResponse`

NewTaskGroupResponseWithDefaults instantiates a new TaskGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *TaskGroupResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *TaskGroupResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *TaskGroupResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *TaskGroupResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetDescription

`func (o *TaskGroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskGroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskGroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAllTasksDispatched

`func (o *TaskGroupResponse) GetAllTasksDispatched() bool`

GetAllTasksDispatched returns the AllTasksDispatched field if non-nil, zero value otherwise.

### GetAllTasksDispatchedOk

`func (o *TaskGroupResponse) GetAllTasksDispatchedOk() (*bool, bool)`

GetAllTasksDispatchedOk returns a tuple with the AllTasksDispatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasksDispatched

`func (o *TaskGroupResponse) SetAllTasksDispatched(v bool)`

SetAllTasksDispatched sets AllTasksDispatched field to given value.


### GetWaiting

`func (o *TaskGroupResponse) GetWaiting() int32`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *TaskGroupResponse) GetWaitingOk() (*int32, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *TaskGroupResponse) SetWaiting(v int32)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *TaskGroupResponse) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.

### GetSkipped

`func (o *TaskGroupResponse) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *TaskGroupResponse) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *TaskGroupResponse) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *TaskGroupResponse) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetRunning

`func (o *TaskGroupResponse) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *TaskGroupResponse) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *TaskGroupResponse) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *TaskGroupResponse) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetCompleted

`func (o *TaskGroupResponse) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TaskGroupResponse) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TaskGroupResponse) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *TaskGroupResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCanceled

`func (o *TaskGroupResponse) GetCanceled() int32`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *TaskGroupResponse) GetCanceledOk() (*int32, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *TaskGroupResponse) SetCanceled(v int32)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *TaskGroupResponse) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetFailed

`func (o *TaskGroupResponse) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TaskGroupResponse) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TaskGroupResponse) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *TaskGroupResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetCanceling

`func (o *TaskGroupResponse) GetCanceling() int32`

GetCanceling returns the Canceling field if non-nil, zero value otherwise.

### GetCancelingOk

`func (o *TaskGroupResponse) GetCancelingOk() (*int32, bool)`

GetCancelingOk returns a tuple with the Canceling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceling

`func (o *TaskGroupResponse) SetCanceling(v int32)`

SetCanceling sets Canceling field to given value.

### HasCanceling

`func (o *TaskGroupResponse) HasCanceling() bool`

HasCanceling returns a boolean if a field has been set.

### GetGroupProgressReports

`func (o *TaskGroupResponse) GetGroupProgressReports() []GroupProgressReportResponse`

GetGroupProgressReports returns the GroupProgressReports field if non-nil, zero value otherwise.

### GetGroupProgressReportsOk

`func (o *TaskGroupResponse) GetGroupProgressReportsOk() (*[]GroupProgressReportResponse, bool)`

GetGroupProgressReportsOk returns a tuple with the GroupProgressReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProgressReports

`func (o *TaskGroupResponse) SetGroupProgressReports(v []GroupProgressReportResponse)`

SetGroupProgressReports sets GroupProgressReports field to given value.

### HasGroupProgressReports

`func (o *TaskGroupResponse) HasGroupProgressReports() bool`

HasGroupProgressReports returns a boolean if a field has been set.

### GetTasks

`func (o *TaskGroupResponse) GetTasks() []MinimalTaskResponse`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskGroupResponse) GetTasksOk() (*[]MinimalTaskResponse, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskGroupResponse) SetTasks(v []MinimalTaskResponse)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TaskGroupResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


