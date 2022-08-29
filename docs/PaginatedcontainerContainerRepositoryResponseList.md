# PaginatedcontainerContainerRepositoryResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ContainerContainerRepositoryResponse**](ContainerContainerRepositoryResponse.md) |  | [optional] 

## Methods

### NewPaginatedcontainerContainerRepositoryResponseList

`func NewPaginatedcontainerContainerRepositoryResponseList() *PaginatedcontainerContainerRepositoryResponseList`

NewPaginatedcontainerContainerRepositoryResponseList instantiates a new PaginatedcontainerContainerRepositoryResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedcontainerContainerRepositoryResponseListWithDefaults

`func NewPaginatedcontainerContainerRepositoryResponseListWithDefaults() *PaginatedcontainerContainerRepositoryResponseList`

NewPaginatedcontainerContainerRepositoryResponseListWithDefaults instantiates a new PaginatedcontainerContainerRepositoryResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedcontainerContainerRepositoryResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedcontainerContainerRepositoryResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedcontainerContainerRepositoryResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedcontainerContainerRepositoryResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedcontainerContainerRepositoryResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetResults() []ContainerContainerRepositoryResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedcontainerContainerRepositoryResponseList) GetResultsOk() (*[]ContainerContainerRepositoryResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedcontainerContainerRepositoryResponseList) SetResults(v []ContainerContainerRepositoryResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedcontainerContainerRepositoryResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


