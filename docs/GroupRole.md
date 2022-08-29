# GroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**ContentObject** | **NullableString** | pulp_href of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on the model-level. | 

## Methods

### NewGroupRole

`func NewGroupRole(role string, contentObject NullableString, ) *GroupRole`

NewGroupRole instantiates a new GroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleWithDefaults

`func NewGroupRoleWithDefaults() *GroupRole`

NewGroupRoleWithDefaults instantiates a new GroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *GroupRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GroupRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GroupRole) SetRole(v string)`

SetRole sets Role field to given value.


### GetContentObject

`func (o *GroupRole) GetContentObject() string`

GetContentObject returns the ContentObject field if non-nil, zero value otherwise.

### GetContentObjectOk

`func (o *GroupRole) GetContentObjectOk() (*string, bool)`

GetContentObjectOk returns a tuple with the ContentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObject

`func (o *GroupRole) SetContentObject(v string)`

SetContentObject sets ContentObject field to given value.


### SetContentObjectNil

`func (o *GroupRole) SetContentObjectNil(b bool)`

 SetContentObjectNil sets the value for ContentObject to be an explicit nil

### UnsetContentObject
`func (o *GroupRole) UnsetContentObject()`

UnsetContentObject ensures that no value is present for ContentObject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


