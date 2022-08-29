# AccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionsAssignment** | Pointer to **[]map[string]interface{}** | List of callables that define the new permissions to be created for new objects.This is deprecated. Use &#x60;creation_hooks&#x60; instead. | [optional] 
**CreationHooks** | Pointer to **[]map[string]interface{}** | List of callables that may associate user roles for new objects. | [optional] 
**Statements** | **[]map[string]interface{}** | List of policy statements defining the policy. | 
**QuerysetScoping** | Pointer to **map[string]interface{}** | A callable for performing queryset scoping. See plugin documentation for valid callables. Set to blank to turn off queryset scoping. | [optional] 

## Methods

### NewAccessPolicy

`func NewAccessPolicy(statements []map[string]interface{}, ) *AccessPolicy`

NewAccessPolicy instantiates a new AccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyWithDefaults

`func NewAccessPolicyWithDefaults() *AccessPolicy`

NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionsAssignment

`func (o *AccessPolicy) GetPermissionsAssignment() []map[string]interface{}`

GetPermissionsAssignment returns the PermissionsAssignment field if non-nil, zero value otherwise.

### GetPermissionsAssignmentOk

`func (o *AccessPolicy) GetPermissionsAssignmentOk() (*[]map[string]interface{}, bool)`

GetPermissionsAssignmentOk returns a tuple with the PermissionsAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsAssignment

`func (o *AccessPolicy) SetPermissionsAssignment(v []map[string]interface{})`

SetPermissionsAssignment sets PermissionsAssignment field to given value.

### HasPermissionsAssignment

`func (o *AccessPolicy) HasPermissionsAssignment() bool`

HasPermissionsAssignment returns a boolean if a field has been set.

### GetCreationHooks

`func (o *AccessPolicy) GetCreationHooks() []map[string]interface{}`

GetCreationHooks returns the CreationHooks field if non-nil, zero value otherwise.

### GetCreationHooksOk

`func (o *AccessPolicy) GetCreationHooksOk() (*[]map[string]interface{}, bool)`

GetCreationHooksOk returns a tuple with the CreationHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationHooks

`func (o *AccessPolicy) SetCreationHooks(v []map[string]interface{})`

SetCreationHooks sets CreationHooks field to given value.

### HasCreationHooks

`func (o *AccessPolicy) HasCreationHooks() bool`

HasCreationHooks returns a boolean if a field has been set.

### GetStatements

`func (o *AccessPolicy) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccessPolicy) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccessPolicy) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.


### GetQuerysetScoping

`func (o *AccessPolicy) GetQuerysetScoping() map[string]interface{}`

GetQuerysetScoping returns the QuerysetScoping field if non-nil, zero value otherwise.

### GetQuerysetScopingOk

`func (o *AccessPolicy) GetQuerysetScopingOk() (*map[string]interface{}, bool)`

GetQuerysetScopingOk returns a tuple with the QuerysetScoping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerysetScoping

`func (o *AccessPolicy) SetQuerysetScoping(v map[string]interface{})`

SetQuerysetScoping sets QuerysetScoping field to given value.

### HasQuerysetScoping

`func (o *AccessPolicy) HasQuerysetScoping() bool`

HasQuerysetScoping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


