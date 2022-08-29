# RemoveImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** | sha256 of the Manifest file | 

## Methods

### NewRemoveImage

`func NewRemoveImage(digest string, ) *RemoveImage`

NewRemoveImage instantiates a new RemoveImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveImageWithDefaults

`func NewRemoveImageWithDefaults() *RemoveImage`

NewRemoveImageWithDefaults instantiates a new RemoveImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *RemoveImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RemoveImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RemoveImage) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


