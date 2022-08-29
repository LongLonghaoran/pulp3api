# ContainerTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Tag name | 
**TaggedManifest** | **string** | Manifest that is tagged | 

## Methods

### NewContainerTagResponse

`func NewContainerTagResponse(name string, taggedManifest string, ) *ContainerTagResponse`

NewContainerTagResponse instantiates a new ContainerTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTagResponseWithDefaults

`func NewContainerTagResponseWithDefaults() *ContainerTagResponse`

NewContainerTagResponseWithDefaults instantiates a new ContainerTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerTagResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerTagResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerTagResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerTagResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerTagResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerTagResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerTagResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerTagResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *ContainerTagResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTagResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTagResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTaggedManifest

`func (o *ContainerTagResponse) GetTaggedManifest() string`

GetTaggedManifest returns the TaggedManifest field if non-nil, zero value otherwise.

### GetTaggedManifestOk

`func (o *ContainerTagResponse) GetTaggedManifestOk() (*string, bool)`

GetTaggedManifestOk returns a tuple with the TaggedManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedManifest

`func (o *ContainerTagResponse) SetTaggedManifest(v string)`

SetTaggedManifest sets TaggedManifest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


