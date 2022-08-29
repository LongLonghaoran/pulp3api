# \WorkersApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkersList**](WorkersApi.md#WorkersList) | **Get** /pulp/api/v3/workers/ | List workers
[**WorkersRead**](WorkersApi.md#WorkersRead) | **Get** /pulp/api/v3/workers/{pulp_id}/ | Inspect a worker



## WorkersList

> PaginatedWorkerResponseList WorkersList(ctx).LastHeartbeat(lastHeartbeat).LastHeartbeatGt(lastHeartbeatGt).LastHeartbeatGte(lastHeartbeatGte).LastHeartbeatLt(lastHeartbeatLt).LastHeartbeatLte(lastHeartbeatLte).LastHeartbeatRange(lastHeartbeatRange).Limit(limit).Missing(missing).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Online(online).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List workers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    lastHeartbeat := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    lastHeartbeatGt := time.Now() // time.Time | Filter results where last_heartbeat is greater than value (optional)
    lastHeartbeatGte := time.Now() // time.Time | Filter results where last_heartbeat is greater than or equal to value (optional)
    lastHeartbeatLt := time.Now() // time.Time | Filter results where last_heartbeat is less than value (optional)
    lastHeartbeatLte := time.Now() // time.Time | Filter results where last_heartbeat is less than or equal to value (optional)
    lastHeartbeatRange := []time.Time{time.Now()} // []time.Time | Filter results where last_heartbeat is between two comma separated values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    missing := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    online := true // bool |  (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.WorkersList(context.Background()).LastHeartbeat(lastHeartbeat).LastHeartbeatGt(lastHeartbeatGt).LastHeartbeatGte(lastHeartbeatGte).LastHeartbeatLt(lastHeartbeatLt).LastHeartbeatLte(lastHeartbeatLte).LastHeartbeatRange(lastHeartbeatRange).Limit(limit).Missing(missing).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Online(online).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.WorkersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersList`: PaginatedWorkerResponseList
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.WorkersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastHeartbeat** | **time.Time** | ISO 8601 formatted dates are supported | 
 **lastHeartbeatGt** | **time.Time** | Filter results where last_heartbeat is greater than value | 
 **lastHeartbeatGte** | **time.Time** | Filter results where last_heartbeat is greater than or equal to value | 
 **lastHeartbeatLt** | **time.Time** | Filter results where last_heartbeat is less than value | 
 **lastHeartbeatLte** | **time.Time** | Filter results where last_heartbeat is less than or equal to value | 
 **lastHeartbeatRange** | [**[]time.Time**](time.Time.md) | Filter results where last_heartbeat is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **missing** | **bool** |  | 
 **name** | **string** |  | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **online** | **bool** |  | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedWorkerResponseList**](PaginatedWorkerResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersRead

> WorkerResponse WorkersRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a worker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pulpId := TODO // string | A UUID string identifying this worker.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.WorkersRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.WorkersRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersRead`: WorkerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.WorkersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this worker. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**WorkerResponse**](WorkerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

