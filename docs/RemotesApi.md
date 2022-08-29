# \RemotesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesCreate**](RemotesApi.md#RemotesCreate) | **Post** /pulp/api/v3/remotes/ | Create a remote
[**RemotesDelete**](RemotesApi.md#RemotesDelete) | **Delete** /pulp/api/v3/remotes/{pulp_id}/ | Delete a remote
[**RemotesList**](RemotesApi.md#RemotesList) | **Get** /pulp/api/v3/remotes/ | List remotes
[**RemotesPartialUpdate**](RemotesApi.md#RemotesPartialUpdate) | **Patch** /pulp/api/v3/remotes/{pulp_id}/ | Update a remote
[**RemotesRead**](RemotesApi.md#RemotesRead) | **Get** /pulp/api/v3/remotes/{pulp_id}/ | Inspect a remote
[**RemotesUpdate**](RemotesApi.md#RemotesUpdate) | **Put** /pulp/api/v3/remotes/{pulp_id}/ | Update a remote



## RemotesCreate

> RemoteResponse RemotesCreate(ctx).Remote(remote).Execute()

Create a remote



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
    remote := *openapiclient.NewRemote("Name_example", "Url_example") // Remote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesCreate(context.Background()).Remote(remote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesCreate`: RemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remote** | [**Remote**](Remote.md) |  | 

### Return type

[**RemoteResponse**](RemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesDelete

> AsyncOperationResponse RemotesDelete(ctx, pulpId).Execute()

Delete a remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this remote.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesList

> PaginatedRemoteResponseList RemotesList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List remotes



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    pulpLastUpdated := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
    pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
    pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
    pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
    pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesList`: PaginatedRemoteResponseList
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | ISO 8601 formatted dates are supported | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRemoteResponseList**](PaginatedRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesPartialUpdate

> AsyncOperationResponse RemotesPartialUpdate(ctx, pulpId).PatchedRemote(patchedRemote).Execute()

Update a remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this remote.
    patchedRemote := *openapiclient.NewPatchedRemote() // PatchedRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesPartialUpdate(context.Background(), pulpId).PatchedRemote(patchedRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRemote** | [**PatchedRemote**](PatchedRemote.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesRead

> RemoteResponse RemotesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this remote.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesRead`: RemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**RemoteResponse**](RemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesUpdate

> AsyncOperationResponse RemotesUpdate(ctx, pulpId).Remote(remote).Execute()

Update a remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this remote.
    remote := *openapiclient.NewRemote("Name_example", "Url_example") // Remote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesApi.RemotesUpdate(context.Background(), pulpId).Remote(remote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesApi.RemotesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesApi.RemotesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remote** | [**Remote**](Remote.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

