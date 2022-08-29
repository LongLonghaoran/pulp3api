# UploadDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Size** | **int32** | The size of the upload in bytes. | 
**Completed** | Pointer to **time.Time** | Timestamp when upload is committed. | [optional] [readonly] 
**Chunks** | Pointer to [**[]UploadChunkResponse**](UploadChunkResponse.md) |  | [optional] [readonly] 

## Methods

### NewUploadDetailResponse

`func NewUploadDetailResponse(size int32, ) *UploadDetailResponse`

NewUploadDetailResponse instantiates a new UploadDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDetailResponseWithDefaults

`func NewUploadDetailResponseWithDefaults() *UploadDetailResponse`

NewUploadDetailResponseWithDefaults instantiates a new UploadDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *UploadDetailResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *UploadDetailResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *UploadDetailResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *UploadDetailResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *UploadDetailResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *UploadDetailResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *UploadDetailResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *UploadDetailResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetSize

`func (o *UploadDetailResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadDetailResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadDetailResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetCompleted

`func (o *UploadDetailResponse) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *UploadDetailResponse) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *UploadDetailResponse) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *UploadDetailResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetChunks

`func (o *UploadDetailResponse) GetChunks() []UploadChunkResponse`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *UploadDetailResponse) GetChunksOk() (*[]UploadChunkResponse, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *UploadDetailResponse) SetChunks(v []UploadChunkResponse)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *UploadDetailResponse) HasChunks() bool`

HasChunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


