# StatusResponseRedisConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** | Info about whether the app can connect to Redis | 

## Methods

### NewStatusResponseRedisConnection

`func NewStatusResponseRedisConnection(connected bool, ) *StatusResponseRedisConnection`

NewStatusResponseRedisConnection instantiates a new StatusResponseRedisConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseRedisConnectionWithDefaults

`func NewStatusResponseRedisConnectionWithDefaults() *StatusResponseRedisConnection`

NewStatusResponseRedisConnectionWithDefaults instantiates a new StatusResponseRedisConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *StatusResponseRedisConnection) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *StatusResponseRedisConnection) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *StatusResponseRedisConnection) SetConnected(v bool)`

SetConnected sets Connected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


