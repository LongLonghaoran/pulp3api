# AccessPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PermissionsAssignment** | Pointer to **[]map[string]interface{}** | List of callables that define the new permissions to be created for new objects.This is deprecated. Use &#x60;creation_hooks&#x60; instead. | [optional] 
**CreationHooks** | Pointer to **[]map[string]interface{}** | List of callables that may associate user roles for new objects. | [optional] 
**Statements** | **[]map[string]interface{}** | List of policy statements defining the policy. | 
**ViewsetName** | Pointer to **string** | The name of ViewSet this AccessPolicy authorizes. | [optional] [readonly] 
**Customized** | Pointer to **bool** | True if the AccessPolicy has been user-modified. False otherwise. | [optional] [readonly] 
**QuerysetScoping** | Pointer to **map[string]interface{}** | A callable for performing queryset scoping. See plugin documentation for valid callables. Set to blank to turn off queryset scoping. | [optional] 

## Methods

### NewAccessPolicyResponse

`func NewAccessPolicyResponse(statements []map[string]interface{}, ) *AccessPolicyResponse`

NewAccessPolicyResponse instantiates a new AccessPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyResponseWithDefaults

`func NewAccessPolicyResponseWithDefaults() *AccessPolicyResponse`

NewAccessPolicyResponseWithDefaults instantiates a new AccessPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AccessPolicyResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AccessPolicyResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AccessPolicyResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AccessPolicyResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *AccessPolicyResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *AccessPolicyResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *AccessPolicyResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *AccessPolicyResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPermissionsAssignment

`func (o *AccessPolicyResponse) GetPermissionsAssignment() []map[string]interface{}`

GetPermissionsAssignment returns the PermissionsAssignment field if non-nil, zero value otherwise.

### GetPermissionsAssignmentOk

`func (o *AccessPolicyResponse) GetPermissionsAssignmentOk() (*[]map[string]interface{}, bool)`

GetPermissionsAssignmentOk returns a tuple with the PermissionsAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsAssignment

`func (o *AccessPolicyResponse) SetPermissionsAssignment(v []map[string]interface{})`

SetPermissionsAssignment sets PermissionsAssignment field to given value.

### HasPermissionsAssignment

`func (o *AccessPolicyResponse) HasPermissionsAssignment() bool`

HasPermissionsAssignment returns a boolean if a field has been set.

### GetCreationHooks

`func (o *AccessPolicyResponse) GetCreationHooks() []map[string]interface{}`

GetCreationHooks returns the CreationHooks field if non-nil, zero value otherwise.

### GetCreationHooksOk

`func (o *AccessPolicyResponse) GetCreationHooksOk() (*[]map[string]interface{}, bool)`

GetCreationHooksOk returns a tuple with the CreationHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationHooks

`func (o *AccessPolicyResponse) SetCreationHooks(v []map[string]interface{})`

SetCreationHooks sets CreationHooks field to given value.

### HasCreationHooks

`func (o *AccessPolicyResponse) HasCreationHooks() bool`

HasCreationHooks returns a boolean if a field has been set.

### GetStatements

`func (o *AccessPolicyResponse) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccessPolicyResponse) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccessPolicyResponse) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.


### GetViewsetName

`func (o *AccessPolicyResponse) GetViewsetName() string`

GetViewsetName returns the ViewsetName field if non-nil, zero value otherwise.

### GetViewsetNameOk

`func (o *AccessPolicyResponse) GetViewsetNameOk() (*string, bool)`

GetViewsetNameOk returns a tuple with the ViewsetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsetName

`func (o *AccessPolicyResponse) SetViewsetName(v string)`

SetViewsetName sets ViewsetName field to given value.

### HasViewsetName

`func (o *AccessPolicyResponse) HasViewsetName() bool`

HasViewsetName returns a boolean if a field has been set.

### GetCustomized

`func (o *AccessPolicyResponse) GetCustomized() bool`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *AccessPolicyResponse) GetCustomizedOk() (*bool, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *AccessPolicyResponse) SetCustomized(v bool)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *AccessPolicyResponse) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.

### GetQuerysetScoping

`func (o *AccessPolicyResponse) GetQuerysetScoping() map[string]interface{}`

GetQuerysetScoping returns the QuerysetScoping field if non-nil, zero value otherwise.

### GetQuerysetScopingOk

`func (o *AccessPolicyResponse) GetQuerysetScopingOk() (*map[string]interface{}, bool)`

GetQuerysetScopingOk returns a tuple with the QuerysetScoping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerysetScoping

`func (o *AccessPolicyResponse) SetQuerysetScoping(v map[string]interface{})`

SetQuerysetScoping sets QuerysetScoping field to given value.

### HasQuerysetScoping

`func (o *AccessPolicyResponse) HasQuerysetScoping() bool`

HasQuerysetScoping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


