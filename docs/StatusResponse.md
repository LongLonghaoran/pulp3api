# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]VersionResponse**](VersionResponse.md) | Version information of Pulp components | 
**OnlineWorkers** | [**[]WorkerResponse**](WorkerResponse.md) | List of online workers known to the application. An online worker is actively heartbeating and can respond to new work | 
**OnlineContentApps** | [**[]ContentAppStatusResponse**](ContentAppStatusResponse.md) | List of online content apps known to the application. An online content app is actively heartbeating and can serve data to clients | 
**DatabaseConnection** | [**StatusResponseDatabaseConnection**](StatusResponseDatabaseConnection.md) |  | 
**RedisConnection** | Pointer to [**StatusResponseRedisConnection**](StatusResponseRedisConnection.md) |  | [optional] 
**Storage** | Pointer to [**StatusResponseStorage**](StatusResponseStorage.md) |  | [optional] 

## Methods

### NewStatusResponse

`func NewStatusResponse(versions []VersionResponse, onlineWorkers []WorkerResponse, onlineContentApps []ContentAppStatusResponse, databaseConnection StatusResponseDatabaseConnection, ) *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *StatusResponse) GetVersions() []VersionResponse`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StatusResponse) GetVersionsOk() (*[]VersionResponse, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StatusResponse) SetVersions(v []VersionResponse)`

SetVersions sets Versions field to given value.


### GetOnlineWorkers

`func (o *StatusResponse) GetOnlineWorkers() []WorkerResponse`

GetOnlineWorkers returns the OnlineWorkers field if non-nil, zero value otherwise.

### GetOnlineWorkersOk

`func (o *StatusResponse) GetOnlineWorkersOk() (*[]WorkerResponse, bool)`

GetOnlineWorkersOk returns a tuple with the OnlineWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineWorkers

`func (o *StatusResponse) SetOnlineWorkers(v []WorkerResponse)`

SetOnlineWorkers sets OnlineWorkers field to given value.


### GetOnlineContentApps

`func (o *StatusResponse) GetOnlineContentApps() []ContentAppStatusResponse`

GetOnlineContentApps returns the OnlineContentApps field if non-nil, zero value otherwise.

### GetOnlineContentAppsOk

`func (o *StatusResponse) GetOnlineContentAppsOk() (*[]ContentAppStatusResponse, bool)`

GetOnlineContentAppsOk returns a tuple with the OnlineContentApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineContentApps

`func (o *StatusResponse) SetOnlineContentApps(v []ContentAppStatusResponse)`

SetOnlineContentApps sets OnlineContentApps field to given value.


### GetDatabaseConnection

`func (o *StatusResponse) GetDatabaseConnection() StatusResponseDatabaseConnection`

GetDatabaseConnection returns the DatabaseConnection field if non-nil, zero value otherwise.

### GetDatabaseConnectionOk

`func (o *StatusResponse) GetDatabaseConnectionOk() (*StatusResponseDatabaseConnection, bool)`

GetDatabaseConnectionOk returns a tuple with the DatabaseConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseConnection

`func (o *StatusResponse) SetDatabaseConnection(v StatusResponseDatabaseConnection)`

SetDatabaseConnection sets DatabaseConnection field to given value.


### GetRedisConnection

`func (o *StatusResponse) GetRedisConnection() StatusResponseRedisConnection`

GetRedisConnection returns the RedisConnection field if non-nil, zero value otherwise.

### GetRedisConnectionOk

`func (o *StatusResponse) GetRedisConnectionOk() (*StatusResponseRedisConnection, bool)`

GetRedisConnectionOk returns a tuple with the RedisConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisConnection

`func (o *StatusResponse) SetRedisConnection(v StatusResponseRedisConnection)`

SetRedisConnection sets RedisConnection field to given value.

### HasRedisConnection

`func (o *StatusResponse) HasRedisConnection() bool`

HasRedisConnection returns a boolean if a field has been set.

### GetStorage

`func (o *StatusResponse) GetStorage() StatusResponseStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StatusResponse) GetStorageOk() (*StatusResponseStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StatusResponse) SetStorage(v StatusResponseStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *StatusResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


