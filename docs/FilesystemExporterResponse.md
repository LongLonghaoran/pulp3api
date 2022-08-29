# FilesystemExporterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Unique name of the file system exporter. | 
**Path** | **string** | File system location to export to. | 
**Method** | Pointer to [**NullableMethodEnum**](MethodEnum.md) | Method of exporting | [optional] 

## Methods

### NewFilesystemExporterResponse

`func NewFilesystemExporterResponse(name string, path string, ) *FilesystemExporterResponse`

NewFilesystemExporterResponse instantiates a new FilesystemExporterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemExporterResponseWithDefaults

`func NewFilesystemExporterResponseWithDefaults() *FilesystemExporterResponse`

NewFilesystemExporterResponseWithDefaults instantiates a new FilesystemExporterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *FilesystemExporterResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *FilesystemExporterResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *FilesystemExporterResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *FilesystemExporterResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *FilesystemExporterResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *FilesystemExporterResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *FilesystemExporterResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *FilesystemExporterResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *FilesystemExporterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesystemExporterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesystemExporterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *FilesystemExporterResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemExporterResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemExporterResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetMethod

`func (o *FilesystemExporterResponse) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FilesystemExporterResponse) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FilesystemExporterResponse) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FilesystemExporterResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *FilesystemExporterResponse) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *FilesystemExporterResponse) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


