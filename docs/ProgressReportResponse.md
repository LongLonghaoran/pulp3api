# ProgressReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The message shown to the user for the progress report. | [optional] [readonly] 
**Code** | Pointer to **string** | Identifies the type of progress report&#39;. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the progress report. The possible values are: &#39;waiting&#39;, &#39;skipped&#39;, &#39;running&#39;, &#39;completed&#39;, &#39;failed&#39;, &#39;canceled&#39; and &#39;canceling&#39;. The default is &#39;waiting&#39;. | [optional] [readonly] 
**Total** | Pointer to **int32** | The total count of items. | [optional] [readonly] 
**Done** | Pointer to **int32** | The count of items already processed. Defaults to 0. | [optional] [readonly] 
**Suffix** | Pointer to **NullableString** | The suffix to be shown with the progress report. | [optional] [readonly] 

## Methods

### NewProgressReportResponse

`func NewProgressReportResponse() *ProgressReportResponse`

NewProgressReportResponse instantiates a new ProgressReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressReportResponseWithDefaults

`func NewProgressReportResponseWithDefaults() *ProgressReportResponse`

NewProgressReportResponseWithDefaults instantiates a new ProgressReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProgressReportResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProgressReportResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProgressReportResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProgressReportResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *ProgressReportResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProgressReportResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProgressReportResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProgressReportResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetState

`func (o *ProgressReportResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProgressReportResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProgressReportResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProgressReportResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotal

`func (o *ProgressReportResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProgressReportResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProgressReportResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProgressReportResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetDone

`func (o *ProgressReportResponse) GetDone() int32`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *ProgressReportResponse) GetDoneOk() (*int32, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *ProgressReportResponse) SetDone(v int32)`

SetDone sets Done field to given value.

### HasDone

`func (o *ProgressReportResponse) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetSuffix

`func (o *ProgressReportResponse) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *ProgressReportResponse) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *ProgressReportResponse) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *ProgressReportResponse) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *ProgressReportResponse) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *ProgressReportResponse) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


