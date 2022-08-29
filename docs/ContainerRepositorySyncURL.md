# ContainerRepositorySyncURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** | A remote to sync from. This will override a remote set on repository. | [optional] 
**Mirror** | Pointer to **bool** | If &#x60;&#x60;True&#x60;&#x60;, synchronization will remove all content that is not present in the remote repository. If &#x60;&#x60;False&#x60;&#x60;, sync will be additive only. | [optional] [default to false]
**SignedOnly** | Pointer to **bool** | If &#x60;&#x60;True&#x60;&#x60;, only signed content will be synced. Signatures are not verified. | [optional] [default to false]

## Methods

### NewContainerRepositorySyncURL

`func NewContainerRepositorySyncURL() *ContainerRepositorySyncURL`

NewContainerRepositorySyncURL instantiates a new ContainerRepositorySyncURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRepositorySyncURLWithDefaults

`func NewContainerRepositorySyncURLWithDefaults() *ContainerRepositorySyncURL`

NewContainerRepositorySyncURLWithDefaults instantiates a new ContainerRepositorySyncURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *ContainerRepositorySyncURL) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ContainerRepositorySyncURL) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ContainerRepositorySyncURL) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ContainerRepositorySyncURL) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetMirror

`func (o *ContainerRepositorySyncURL) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *ContainerRepositorySyncURL) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *ContainerRepositorySyncURL) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *ContainerRepositorySyncURL) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### GetSignedOnly

`func (o *ContainerRepositorySyncURL) GetSignedOnly() bool`

GetSignedOnly returns the SignedOnly field if non-nil, zero value otherwise.

### GetSignedOnlyOk

`func (o *ContainerRepositorySyncURL) GetSignedOnlyOk() (*bool, bool)`

GetSignedOnlyOk returns a tuple with the SignedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedOnly

`func (o *ContainerRepositorySyncURL) SetSignedOnly(v bool)`

SetSignedOnly sets SignedOnly field to given value.

### HasSignedOnly

`func (o *ContainerRepositorySyncURL) HasSignedOnly() bool`

HasSignedOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


