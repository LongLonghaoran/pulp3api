# ContentAppStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the worker. | [optional] [readonly] 
**LastHeartbeat** | Pointer to **time.Time** | Timestamp of the last time the worker talked to the service. | [optional] [readonly] 

## Methods

### NewContentAppStatusResponse

`func NewContentAppStatusResponse() *ContentAppStatusResponse`

NewContentAppStatusResponse instantiates a new ContentAppStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentAppStatusResponseWithDefaults

`func NewContentAppStatusResponseWithDefaults() *ContentAppStatusResponse`

NewContentAppStatusResponseWithDefaults instantiates a new ContentAppStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentAppStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentAppStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentAppStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentAppStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *ContentAppStatusResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ContentAppStatusResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ContentAppStatusResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ContentAppStatusResponse) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


