# EvaluationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Parameter value being evaluated. | 
**IsValid** | **bool** | True if evaluation passed, false otherwise. | 
**Messages** | **[]string** | Messages describing results of all evaluations done. May be an empty list. | 

## Methods

### NewEvaluationResponse

`func NewEvaluationResponse(context string, isValid bool, messages []string, ) *EvaluationResponse`

NewEvaluationResponse instantiates a new EvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationResponseWithDefaults

`func NewEvaluationResponseWithDefaults() *EvaluationResponse`

NewEvaluationResponseWithDefaults instantiates a new EvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *EvaluationResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EvaluationResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EvaluationResponse) SetContext(v string)`

SetContext sets Context field to given value.


### GetIsValid

`func (o *EvaluationResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *EvaluationResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *EvaluationResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetMessages

`func (o *EvaluationResponse) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *EvaluationResponse) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *EvaluationResponse) SetMessages(v []string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


