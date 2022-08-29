# GroupProgressReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The message shown to the user for the group progress report. | [optional] [readonly] 
**Code** | Pointer to **string** | Identifies the type of group progress report&#39;. | [optional] [readonly] 
**Total** | Pointer to **int32** | The total count of items. | [optional] [readonly] 
**Done** | Pointer to **int32** | The count of items already processed. Defaults to 0. | [optional] [readonly] 
**Suffix** | Pointer to **NullableString** | The suffix to be shown with the group progress report. | [optional] [readonly] 

## Methods

### NewGroupProgressReportResponse

`func NewGroupProgressReportResponse() *GroupProgressReportResponse`

NewGroupProgressReportResponse instantiates a new GroupProgressReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupProgressReportResponseWithDefaults

`func NewGroupProgressReportResponseWithDefaults() *GroupProgressReportResponse`

NewGroupProgressReportResponseWithDefaults instantiates a new GroupProgressReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GroupProgressReportResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GroupProgressReportResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GroupProgressReportResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GroupProgressReportResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *GroupProgressReportResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GroupProgressReportResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GroupProgressReportResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GroupProgressReportResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTotal

`func (o *GroupProgressReportResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GroupProgressReportResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GroupProgressReportResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GroupProgressReportResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetDone

`func (o *GroupProgressReportResponse) GetDone() int32`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *GroupProgressReportResponse) GetDoneOk() (*int32, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *GroupProgressReportResponse) SetDone(v int32)`

SetDone sets Done field to given value.

### HasDone

`func (o *GroupProgressReportResponse) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetSuffix

`func (o *GroupProgressReportResponse) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *GroupProgressReportResponse) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *GroupProgressReportResponse) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *GroupProgressReportResponse) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *GroupProgressReportResponse) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *GroupProgressReportResponse) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


