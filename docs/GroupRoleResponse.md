# GroupRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Role** | **string** |  | 
**ContentObject** | **NullableString** | pulp_href of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on the model-level. | 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Permissions** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewGroupRoleResponse

`func NewGroupRoleResponse(role string, contentObject NullableString, ) *GroupRoleResponse`

NewGroupRoleResponse instantiates a new GroupRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleResponseWithDefaults

`func NewGroupRoleResponseWithDefaults() *GroupRoleResponse`

NewGroupRoleResponseWithDefaults instantiates a new GroupRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *GroupRoleResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *GroupRoleResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *GroupRoleResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *GroupRoleResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *GroupRoleResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *GroupRoleResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *GroupRoleResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *GroupRoleResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetRole

`func (o *GroupRoleResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GroupRoleResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GroupRoleResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetContentObject

`func (o *GroupRoleResponse) GetContentObject() string`

GetContentObject returns the ContentObject field if non-nil, zero value otherwise.

### GetContentObjectOk

`func (o *GroupRoleResponse) GetContentObjectOk() (*string, bool)`

GetContentObjectOk returns a tuple with the ContentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObject

`func (o *GroupRoleResponse) SetContentObject(v string)`

SetContentObject sets ContentObject field to given value.


### SetContentObjectNil

`func (o *GroupRoleResponse) SetContentObjectNil(b bool)`

 SetContentObjectNil sets the value for ContentObject to be an explicit nil

### UnsetContentObject
`func (o *GroupRoleResponse) UnsetContentObject()`

UnsetContentObject ensures that no value is present for ContentObject, not even an explicit nil
### GetDescription

`func (o *GroupRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *GroupRoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupRoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupRoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupRoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


