# NestedRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to **[]string** |  | [optional] [default to []]
**Groups** | Pointer to **[]string** |  | [optional] [default to []]
**Role** | **string** |  | 

## Methods

### NewNestedRole

`func NewNestedRole(role string, ) *NestedRole`

NewNestedRole instantiates a new NestedRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRoleWithDefaults

`func NewNestedRoleWithDefaults() *NestedRole`

NewNestedRoleWithDefaults instantiates a new NestedRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *NestedRole) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NestedRole) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NestedRole) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NestedRole) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *NestedRole) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *NestedRole) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *NestedRole) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *NestedRole) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetRole

`func (o *NestedRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NestedRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NestedRole) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


