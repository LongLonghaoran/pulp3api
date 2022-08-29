# Copy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** | A JSON document describing sources, destinations, and content to be copied | 
**Structured** | Pointer to **bool** | Also copy any distributions, components, and releases as needed for any packages being copied. This will allow for structured publications of the target repository.Default is set to True | [optional] [default to true]
**DependencySolving** | Pointer to **bool** | Also copy dependencies of any packages being copied. NOT YETIMPLEMENTED! You must keep this at \&quot;False\&quot;! | [optional] [default to false]

## Methods

### NewCopy

`func NewCopy(config map[string]interface{}, ) *Copy`

NewCopy instantiates a new Copy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyWithDefaults

`func NewCopyWithDefaults() *Copy`

NewCopyWithDefaults instantiates a new Copy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Copy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Copy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Copy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetStructured

`func (o *Copy) GetStructured() bool`

GetStructured returns the Structured field if non-nil, zero value otherwise.

### GetStructuredOk

`func (o *Copy) GetStructuredOk() (*bool, bool)`

GetStructuredOk returns a tuple with the Structured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructured

`func (o *Copy) SetStructured(v bool)`

SetStructured sets Structured field to given value.

### HasStructured

`func (o *Copy) HasStructured() bool`

HasStructured returns a boolean if a field has been set.

### GetDependencySolving

`func (o *Copy) GetDependencySolving() bool`

GetDependencySolving returns the DependencySolving field if non-nil, zero value otherwise.

### GetDependencySolvingOk

`func (o *Copy) GetDependencySolvingOk() (*bool, bool)`

GetDependencySolvingOk returns a tuple with the DependencySolving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencySolving

`func (o *Copy) SetDependencySolving(v bool)`

SetDependencySolving sets DependencySolving field to given value.

### HasDependencySolving

`func (o *Copy) HasDependencySolving() bool`

HasDependencySolving returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


