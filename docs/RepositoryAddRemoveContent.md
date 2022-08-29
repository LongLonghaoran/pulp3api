# RepositoryAddRemoveContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddContentUnits** | Pointer to **[]interface{}** | A list of content units to add to a new repository version. This content is added after remove_content_units are removed. | [optional] 
**RemoveContentUnits** | Pointer to **[]interface{}** | A list of content units to remove from the latest repository version. You may also specify &#39;*&#39; as an entry to remove all content. This content is removed before add_content_units are added. | [optional] 
**BaseVersion** | Pointer to **string** | A repository version whose content will be used as the initial set of content for the new repository version | [optional] 

## Methods

### NewRepositoryAddRemoveContent

`func NewRepositoryAddRemoveContent() *RepositoryAddRemoveContent`

NewRepositoryAddRemoveContent instantiates a new RepositoryAddRemoveContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryAddRemoveContentWithDefaults

`func NewRepositoryAddRemoveContentWithDefaults() *RepositoryAddRemoveContent`

NewRepositoryAddRemoveContentWithDefaults instantiates a new RepositoryAddRemoveContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddContentUnits

`func (o *RepositoryAddRemoveContent) GetAddContentUnits() []interface{}`

GetAddContentUnits returns the AddContentUnits field if non-nil, zero value otherwise.

### GetAddContentUnitsOk

`func (o *RepositoryAddRemoveContent) GetAddContentUnitsOk() (*[]interface{}, bool)`

GetAddContentUnitsOk returns a tuple with the AddContentUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddContentUnits

`func (o *RepositoryAddRemoveContent) SetAddContentUnits(v []interface{})`

SetAddContentUnits sets AddContentUnits field to given value.

### HasAddContentUnits

`func (o *RepositoryAddRemoveContent) HasAddContentUnits() bool`

HasAddContentUnits returns a boolean if a field has been set.

### GetRemoveContentUnits

`func (o *RepositoryAddRemoveContent) GetRemoveContentUnits() []interface{}`

GetRemoveContentUnits returns the RemoveContentUnits field if non-nil, zero value otherwise.

### GetRemoveContentUnitsOk

`func (o *RepositoryAddRemoveContent) GetRemoveContentUnitsOk() (*[]interface{}, bool)`

GetRemoveContentUnitsOk returns a tuple with the RemoveContentUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveContentUnits

`func (o *RepositoryAddRemoveContent) SetRemoveContentUnits(v []interface{})`

SetRemoveContentUnits sets RemoveContentUnits field to given value.

### HasRemoveContentUnits

`func (o *RepositoryAddRemoveContent) HasRemoveContentUnits() bool`

HasRemoveContentUnits returns a boolean if a field has been set.

### GetBaseVersion

`func (o *RepositoryAddRemoveContent) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *RepositoryAddRemoveContent) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *RepositoryAddRemoveContent) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.

### HasBaseVersion

`func (o *RepositoryAddRemoveContent) HasBaseVersion() bool`

HasBaseVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


