# TagCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRepository** | Pointer to **string** | A URI of the repository to copy content from. | [optional] 
**SourceRepositoryVersion** | Pointer to **string** | A URI of the repository version to copy content from. | [optional] 
**Names** | Pointer to **[]interface{}** | A list of tag names to copy. | [optional] 

## Methods

### NewTagCopy

`func NewTagCopy() *TagCopy`

NewTagCopy instantiates a new TagCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCopyWithDefaults

`func NewTagCopyWithDefaults() *TagCopy`

NewTagCopyWithDefaults instantiates a new TagCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRepository

`func (o *TagCopy) GetSourceRepository() string`

GetSourceRepository returns the SourceRepository field if non-nil, zero value otherwise.

### GetSourceRepositoryOk

`func (o *TagCopy) GetSourceRepositoryOk() (*string, bool)`

GetSourceRepositoryOk returns a tuple with the SourceRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepository

`func (o *TagCopy) SetSourceRepository(v string)`

SetSourceRepository sets SourceRepository field to given value.

### HasSourceRepository

`func (o *TagCopy) HasSourceRepository() bool`

HasSourceRepository returns a boolean if a field has been set.

### GetSourceRepositoryVersion

`func (o *TagCopy) GetSourceRepositoryVersion() string`

GetSourceRepositoryVersion returns the SourceRepositoryVersion field if non-nil, zero value otherwise.

### GetSourceRepositoryVersionOk

`func (o *TagCopy) GetSourceRepositoryVersionOk() (*string, bool)`

GetSourceRepositoryVersionOk returns a tuple with the SourceRepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepositoryVersion

`func (o *TagCopy) SetSourceRepositoryVersion(v string)`

SetSourceRepositoryVersion sets SourceRepositoryVersion field to given value.

### HasSourceRepositoryVersion

`func (o *TagCopy) HasSourceRepositoryVersion() bool`

HasSourceRepositoryVersion returns a boolean if a field has been set.

### GetNames

`func (o *TagCopy) GetNames() []interface{}`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *TagCopy) GetNamesOk() (*[]interface{}, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *TagCopy) SetNames(v []interface{})`

SetNames sets Names field to given value.

### HasNames

`func (o *TagCopy) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


