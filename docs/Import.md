# Import

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | **string** | A URI of the Task that ran the Import. | 
**Params** | **map[string]interface{}** | Any parameters that were used to create the import. | 

## Methods

### NewImport

`func NewImport(task string, params map[string]interface{}, ) *Import`

NewImport instantiates a new Import object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportWithDefaults

`func NewImportWithDefaults() *Import`

NewImportWithDefaults instantiates a new Import object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *Import) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Import) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Import) SetTask(v string)`

SetTask sets Task field to given value.


### GetParams

`func (o *Import) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Import) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Import) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


