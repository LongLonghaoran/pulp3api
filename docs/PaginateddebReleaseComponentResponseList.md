# PaginateddebReleaseComponentResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]DebReleaseComponentResponse**](DebReleaseComponentResponse.md) |  | [optional] 

## Methods

### NewPaginateddebReleaseComponentResponseList

`func NewPaginateddebReleaseComponentResponseList() *PaginateddebReleaseComponentResponseList`

NewPaginateddebReleaseComponentResponseList instantiates a new PaginateddebReleaseComponentResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginateddebReleaseComponentResponseListWithDefaults

`func NewPaginateddebReleaseComponentResponseListWithDefaults() *PaginateddebReleaseComponentResponseList`

NewPaginateddebReleaseComponentResponseListWithDefaults instantiates a new PaginateddebReleaseComponentResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginateddebReleaseComponentResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginateddebReleaseComponentResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginateddebReleaseComponentResponseList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginateddebReleaseComponentResponseList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginateddebReleaseComponentResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginateddebReleaseComponentResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginateddebReleaseComponentResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginateddebReleaseComponentResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginateddebReleaseComponentResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginateddebReleaseComponentResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginateddebReleaseComponentResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginateddebReleaseComponentResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginateddebReleaseComponentResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginateddebReleaseComponentResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginateddebReleaseComponentResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginateddebReleaseComponentResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginateddebReleaseComponentResponseList) GetResults() []DebReleaseComponentResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginateddebReleaseComponentResponseList) GetResultsOk() (*[]DebReleaseComponentResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginateddebReleaseComponentResponseList) SetResults(v []DebReleaseComponentResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginateddebReleaseComponentResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


