# TaskScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | The name of the task schedule. | 
**TaskName** | **string** | The name of the task to be scheduled. | 
**DispatchInterval** | **string** | Periodicity of the schedule. | 
**NextDispatch** | Pointer to **time.Time** | Timestamp of the next time the task will be dispatched. | [optional] [readonly] 
**LastTask** | Pointer to **string** | The last task dispatched by this schedule. | [optional] [readonly] 

## Methods

### NewTaskScheduleResponse

`func NewTaskScheduleResponse(name string, taskName string, dispatchInterval string, ) *TaskScheduleResponse`

NewTaskScheduleResponse instantiates a new TaskScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskScheduleResponseWithDefaults

`func NewTaskScheduleResponseWithDefaults() *TaskScheduleResponse`

NewTaskScheduleResponseWithDefaults instantiates a new TaskScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *TaskScheduleResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *TaskScheduleResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *TaskScheduleResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *TaskScheduleResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *TaskScheduleResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *TaskScheduleResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *TaskScheduleResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *TaskScheduleResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *TaskScheduleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskScheduleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskScheduleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTaskName

`func (o *TaskScheduleResponse) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskScheduleResponse) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskScheduleResponse) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetDispatchInterval

`func (o *TaskScheduleResponse) GetDispatchInterval() string`

GetDispatchInterval returns the DispatchInterval field if non-nil, zero value otherwise.

### GetDispatchIntervalOk

`func (o *TaskScheduleResponse) GetDispatchIntervalOk() (*string, bool)`

GetDispatchIntervalOk returns a tuple with the DispatchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchInterval

`func (o *TaskScheduleResponse) SetDispatchInterval(v string)`

SetDispatchInterval sets DispatchInterval field to given value.


### GetNextDispatch

`func (o *TaskScheduleResponse) GetNextDispatch() time.Time`

GetNextDispatch returns the NextDispatch field if non-nil, zero value otherwise.

### GetNextDispatchOk

`func (o *TaskScheduleResponse) GetNextDispatchOk() (*time.Time, bool)`

GetNextDispatchOk returns a tuple with the NextDispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDispatch

`func (o *TaskScheduleResponse) SetNextDispatch(v time.Time)`

SetNextDispatch sets NextDispatch field to given value.

### HasNextDispatch

`func (o *TaskScheduleResponse) HasNextDispatch() bool`

HasNextDispatch returns a boolean if a field has been set.

### GetLastTask

`func (o *TaskScheduleResponse) GetLastTask() string`

GetLastTask returns the LastTask field if non-nil, zero value otherwise.

### GetLastTaskOk

`func (o *TaskScheduleResponse) GetLastTaskOk() (*string, bool)`

GetLastTaskOk returns a tuple with the LastTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTask

`func (o *TaskScheduleResponse) SetLastTask(v string)`

SetLastTask sets LastTask field to given value.

### HasLastTask

`func (o *TaskScheduleResponse) HasLastTask() bool`

HasLastTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


