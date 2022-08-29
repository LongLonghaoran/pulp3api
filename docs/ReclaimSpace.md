# ReclaimSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoHrefs** | **[]interface{}** | Will reclaim space for the specified list of repos. Use [&#39;*&#39;] to specify all repos. | 
**RepoVersionsKeeplist** | Pointer to **[]string** | Will exclude repo versions from space reclaim. | [optional] 

## Methods

### NewReclaimSpace

`func NewReclaimSpace(repoHrefs []interface{}, ) *ReclaimSpace`

NewReclaimSpace instantiates a new ReclaimSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReclaimSpaceWithDefaults

`func NewReclaimSpaceWithDefaults() *ReclaimSpace`

NewReclaimSpaceWithDefaults instantiates a new ReclaimSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoHrefs

`func (o *ReclaimSpace) GetRepoHrefs() []interface{}`

GetRepoHrefs returns the RepoHrefs field if non-nil, zero value otherwise.

### GetRepoHrefsOk

`func (o *ReclaimSpace) GetRepoHrefsOk() (*[]interface{}, bool)`

GetRepoHrefsOk returns a tuple with the RepoHrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoHrefs

`func (o *ReclaimSpace) SetRepoHrefs(v []interface{})`

SetRepoHrefs sets RepoHrefs field to given value.


### GetRepoVersionsKeeplist

`func (o *ReclaimSpace) GetRepoVersionsKeeplist() []string`

GetRepoVersionsKeeplist returns the RepoVersionsKeeplist field if non-nil, zero value otherwise.

### GetRepoVersionsKeeplistOk

`func (o *ReclaimSpace) GetRepoVersionsKeeplistOk() (*[]string, bool)`

GetRepoVersionsKeeplistOk returns a tuple with the RepoVersionsKeeplist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersionsKeeplist

`func (o *ReclaimSpace) SetRepoVersionsKeeplist(v []string)`

SetRepoVersionsKeeplist sets RepoVersionsKeeplist field to given value.

### HasRepoVersionsKeeplist

`func (o *ReclaimSpace) HasRepoVersionsKeeplist() bool`

HasRepoVersionsKeeplist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


