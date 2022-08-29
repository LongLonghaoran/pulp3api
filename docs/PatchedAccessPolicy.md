# PatchedAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionsAssignment** | Pointer to **[]map[string]interface{}** | List of callables that define the new permissions to be created for new objects.This is deprecated. Use &#x60;creation_hooks&#x60; instead. | [optional] 
**CreationHooks** | Pointer to **[]map[string]interface{}** | List of callables that may associate user roles for new objects. | [optional] 
**Statements** | Pointer to **[]map[string]interface{}** | List of policy statements defining the policy. | [optional] 
**QuerysetScoping** | Pointer to **map[string]interface{}** | A callable for performing queryset scoping. See plugin documentation for valid callables. Set to blank to turn off queryset scoping. | [optional] 

## Methods

### NewPatchedAccessPolicy

`func NewPatchedAccessPolicy() *PatchedAccessPolicy`

NewPatchedAccessPolicy instantiates a new PatchedAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAccessPolicyWithDefaults

`func NewPatchedAccessPolicyWithDefaults() *PatchedAccessPolicy`

NewPatchedAccessPolicyWithDefaults instantiates a new PatchedAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionsAssignment

`func (o *PatchedAccessPolicy) GetPermissionsAssignment() []map[string]interface{}`

GetPermissionsAssignment returns the PermissionsAssignment field if non-nil, zero value otherwise.

### GetPermissionsAssignmentOk

`func (o *PatchedAccessPolicy) GetPermissionsAssignmentOk() (*[]map[string]interface{}, bool)`

GetPermissionsAssignmentOk returns a tuple with the PermissionsAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsAssignment

`func (o *PatchedAccessPolicy) SetPermissionsAssignment(v []map[string]interface{})`

SetPermissionsAssignment sets PermissionsAssignment field to given value.

### HasPermissionsAssignment

`func (o *PatchedAccessPolicy) HasPermissionsAssignment() bool`

HasPermissionsAssignment returns a boolean if a field has been set.

### GetCreationHooks

`func (o *PatchedAccessPolicy) GetCreationHooks() []map[string]interface{}`

GetCreationHooks returns the CreationHooks field if non-nil, zero value otherwise.

### GetCreationHooksOk

`func (o *PatchedAccessPolicy) GetCreationHooksOk() (*[]map[string]interface{}, bool)`

GetCreationHooksOk returns a tuple with the CreationHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationHooks

`func (o *PatchedAccessPolicy) SetCreationHooks(v []map[string]interface{})`

SetCreationHooks sets CreationHooks field to given value.

### HasCreationHooks

`func (o *PatchedAccessPolicy) HasCreationHooks() bool`

HasCreationHooks returns a boolean if a field has been set.

### GetStatements

`func (o *PatchedAccessPolicy) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *PatchedAccessPolicy) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *PatchedAccessPolicy) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *PatchedAccessPolicy) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetQuerysetScoping

`func (o *PatchedAccessPolicy) GetQuerysetScoping() map[string]interface{}`

GetQuerysetScoping returns the QuerysetScoping field if non-nil, zero value otherwise.

### GetQuerysetScopingOk

`func (o *PatchedAccessPolicy) GetQuerysetScopingOk() (*map[string]interface{}, bool)`

GetQuerysetScopingOk returns a tuple with the QuerysetScoping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerysetScoping

`func (o *PatchedAccessPolicy) SetQuerysetScoping(v map[string]interface{})`

SetQuerysetScoping sets QuerysetScoping field to given value.

### HasQuerysetScoping

`func (o *PatchedAccessPolicy) HasQuerysetScoping() bool`

HasQuerysetScoping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


