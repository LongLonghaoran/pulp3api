# RepositorySyncURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** | A remote to sync from. This will override a remote set on repository. | [optional] 
**Mirror** | Pointer to **bool** | If &#x60;&#x60;True&#x60;&#x60;, synchronization will remove all content that is not present in the remote repository. If &#x60;&#x60;False&#x60;&#x60;, sync will be additive only. | [optional] [default to false]

## Methods

### NewRepositorySyncURL

`func NewRepositorySyncURL() *RepositorySyncURL`

NewRepositorySyncURL instantiates a new RepositorySyncURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySyncURLWithDefaults

`func NewRepositorySyncURLWithDefaults() *RepositorySyncURL`

NewRepositorySyncURLWithDefaults instantiates a new RepositorySyncURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *RepositorySyncURL) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RepositorySyncURL) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RepositorySyncURL) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *RepositorySyncURL) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetMirror

`func (o *RepositorySyncURL) GetMirror() bool`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *RepositorySyncURL) GetMirrorOk() (*bool, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *RepositorySyncURL) SetMirror(v bool)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *RepositorySyncURL) HasMirror() bool`

HasMirror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


