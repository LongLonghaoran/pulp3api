# ContainerContainerNamespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** |  | 

## Methods

### NewContainerContainerNamespaceResponse

`func NewContainerContainerNamespaceResponse(name string, ) *ContainerContainerNamespaceResponse`

NewContainerContainerNamespaceResponse instantiates a new ContainerContainerNamespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerNamespaceResponseWithDefaults

`func NewContainerContainerNamespaceResponseWithDefaults() *ContainerContainerNamespaceResponse`

NewContainerContainerNamespaceResponseWithDefaults instantiates a new ContainerContainerNamespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerContainerNamespaceResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerContainerNamespaceResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerContainerNamespaceResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerContainerNamespaceResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerContainerNamespaceResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerContainerNamespaceResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerContainerNamespaceResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerContainerNamespaceResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *ContainerContainerNamespaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerNamespaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerNamespaceResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


