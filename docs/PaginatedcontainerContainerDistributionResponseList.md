# PaginatedcontainerContainerDistributionResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ContainerContainerDistributionResponse**](ContainerContainerDistributionResponse.md) |  | [optional] 

## Methods

### NewPaginatedcontainerContainerDistributionResponseList

`func NewPaginatedcontainerContainerDistributionResponseList() *PaginatedcontainerContainerDistributionResponseList`

NewPaginatedcontainerContainerDistributionResponseList instantiates a new PaginatedcontainerContainerDistributionResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedcontainerContainerDistributionResponseListWithDefaults

`func NewPaginatedcontainerContainerDistributionResponseListWithDefaults() *PaginatedcontainerContainerDistributionResponseList`

NewPaginatedcontainerContainerDistributionResponseListWithDefaults instantiates a new PaginatedcontainerContainerDistributionResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedcontainerContainerDistributionResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedcontainerContainerDistributionResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedcontainerContainerDistributionResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedcontainerContainerDistributionResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedcontainerContainerDistributionResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedcontainerContainerDistributionResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedcontainerContainerDistributionResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedcontainerContainerDistributionResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedcontainerContainerDistributionResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedcontainerContainerDistributionResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedcontainerContainerDistributionResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedcontainerContainerDistributionResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedcontainerContainerDistributionResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedcontainerContainerDistributionResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedcontainerContainerDistributionResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedcontainerContainerDistributionResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedcontainerContainerDistributionResponseList) GetResults() []ContainerContainerDistributionResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedcontainerContainerDistributionResponseList) GetResultsOk() (*[]ContainerContainerDistributionResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedcontainerContainerDistributionResponseList) SetResults(v []ContainerContainerDistributionResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedcontainerContainerDistributionResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


