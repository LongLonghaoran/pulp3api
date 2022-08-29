# PaginateddebAptPublicationResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]DebAptPublicationResponse**](DebAptPublicationResponse.md) |  | [optional] 

## Methods

### NewPaginateddebAptPublicationResponseList

`func NewPaginateddebAptPublicationResponseList() *PaginateddebAptPublicationResponseList`

NewPaginateddebAptPublicationResponseList instantiates a new PaginateddebAptPublicationResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginateddebAptPublicationResponseListWithDefaults

`func NewPaginateddebAptPublicationResponseListWithDefaults() *PaginateddebAptPublicationResponseList`

NewPaginateddebAptPublicationResponseListWithDefaults instantiates a new PaginateddebAptPublicationResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginateddebAptPublicationResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginateddebAptPublicationResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginateddebAptPublicationResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginateddebAptPublicationResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginateddebAptPublicationResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginateddebAptPublicationResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginateddebAptPublicationResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginateddebAptPublicationResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginateddebAptPublicationResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginateddebAptPublicationResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginateddebAptPublicationResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginateddebAptPublicationResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginateddebAptPublicationResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginateddebAptPublicationResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginateddebAptPublicationResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginateddebAptPublicationResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginateddebAptPublicationResponseList) GetResults() []DebAptPublicationResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginateddebAptPublicationResponseList) GetResultsOk() (*[]DebAptPublicationResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginateddebAptPublicationResponseList) SetResults(v []DebAptPublicationResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginateddebAptPublicationResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


