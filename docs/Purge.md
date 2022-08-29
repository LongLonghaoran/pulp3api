# Purge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishedBefore** | Pointer to **time.Time** | Purge tasks completed earlier than this timestamp. Format &#39;%Y-%m-%d[T%H:%M:%S]&#39; | [optional] 
**States** | Pointer to [**[]StatesEnum**](StatesEnum.md) | List of task-states to be purged. Only &#39;final&#39; states are allowed. | [optional] [default to ["completed"]]

## Methods

### NewPurge

`func NewPurge() *Purge`

NewPurge instantiates a new Purge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeWithDefaults

`func NewPurgeWithDefaults() *Purge`

NewPurgeWithDefaults instantiates a new Purge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishedBefore

`func (o *Purge) GetFinishedBefore() time.Time`

GetFinishedBefore returns the FinishedBefore field if non-nil, zero value otherwise.

### GetFinishedBeforeOk

`func (o *Purge) GetFinishedBeforeOk() (*time.Time, bool)`

GetFinishedBeforeOk returns a tuple with the FinishedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedBefore

`func (o *Purge) SetFinishedBefore(v time.Time)`

SetFinishedBefore sets FinishedBefore field to given value.

### HasFinishedBefore

`func (o *Purge) HasFinishedBefore() bool`

HasFinishedBefore returns a boolean if a field has been set.

### GetStates

`func (o *Purge) GetStates() []StatesEnum`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *Purge) GetStatesOk() (*[]StatesEnum, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *Purge) SetStates(v []StatesEnum)`

SetStates sets States field to given value.

### HasStates

`func (o *Purge) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


