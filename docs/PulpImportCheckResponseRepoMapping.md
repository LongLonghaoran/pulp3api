# PulpImportCheckResponseRepoMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Parameter value being evaluated. | 
**IsValid** | **bool** | True if evaluation passed, false otherwise. | 
**Messages** | **[]string** | Messages describing results of all evaluations done. May be an empty list. | 

## Methods

### NewPulpImportCheckResponseRepoMapping

`func NewPulpImportCheckResponseRepoMapping(context string, isValid bool, messages []string, ) *PulpImportCheckResponseRepoMapping`

NewPulpImportCheckResponseRepoMapping instantiates a new PulpImportCheckResponseRepoMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImportCheckResponseRepoMappingWithDefaults

`func NewPulpImportCheckResponseRepoMappingWithDefaults() *PulpImportCheckResponseRepoMapping`

NewPulpImportCheckResponseRepoMappingWithDefaults instantiates a new PulpImportCheckResponseRepoMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PulpImportCheckResponseRepoMapping) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PulpImportCheckResponseRepoMapping) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PulpImportCheckResponseRepoMapping) SetContext(v string)`

SetContext sets Context field to given value.


### GetIsValid

`func (o *PulpImportCheckResponseRepoMapping) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *PulpImportCheckResponseRepoMapping) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *PulpImportCheckResponseRepoMapping) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetMessages

`func (o *PulpImportCheckResponseRepoMapping) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PulpImportCheckResponseRepoMapping) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PulpImportCheckResponseRepoMapping) SetMessages(v []string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


