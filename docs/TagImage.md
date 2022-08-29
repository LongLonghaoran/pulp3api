# TagImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | A tag name | 
**Digest** | **string** | sha256 of the Manifest file | 

## Methods

### NewTagImage

`func NewTagImage(tag string, digest string, ) *TagImage`

NewTagImage instantiates a new TagImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagImageWithDefaults

`func NewTagImageWithDefaults() *TagImage`

NewTagImageWithDefaults instantiates a new TagImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *TagImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagImage) SetTag(v string)`

SetTag sets Tag field to given value.


### GetDigest

`func (o *TagImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *TagImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *TagImage) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


