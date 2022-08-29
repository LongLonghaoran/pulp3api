# StatusResponseStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of bytes | 
**Used** | **int32** | Number of bytes in use | 
**Free** | **int32** | Number of free bytes | 

## Methods

### NewStatusResponseStorage

`func NewStatusResponseStorage(total int32, used int32, free int32, ) *StatusResponseStorage`

NewStatusResponseStorage instantiates a new StatusResponseStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseStorageWithDefaults

`func NewStatusResponseStorageWithDefaults() *StatusResponseStorage`

NewStatusResponseStorageWithDefaults instantiates a new StatusResponseStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *StatusResponseStorage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StatusResponseStorage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StatusResponseStorage) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetUsed

`func (o *StatusResponseStorage) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StatusResponseStorage) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StatusResponseStorage) SetUsed(v int32)`

SetUsed sets Used field to given value.


### GetFree

`func (o *StatusResponseStorage) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *StatusResponseStorage) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *StatusResponseStorage) SetFree(v int32)`

SetFree sets Free field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


