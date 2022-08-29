# \ContentguardsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCreate**](ContentguardsApi.md#ContentguardsCreate) | **Post** /pulp/api/v3/contentguards/ | Create a content guard
[**ContentguardsDelete**](ContentguardsApi.md#ContentguardsDelete) | **Delete** /pulp/api/v3/contentguards/{pulp_id}/ | Delete a content guard
[**ContentguardsList**](ContentguardsApi.md#ContentguardsList) | **Get** /pulp/api/v3/contentguards/ | List content guards
[**ContentguardsPartialUpdate**](ContentguardsApi.md#ContentguardsPartialUpdate) | **Patch** /pulp/api/v3/contentguards/{pulp_id}/ | Update a content guard
[**ContentguardsRead**](ContentguardsApi.md#ContentguardsRead) | **Get** /pulp/api/v3/contentguards/{pulp_id}/ | Inspect a content guard
[**ContentguardsUpdate**](ContentguardsApi.md#ContentguardsUpdate) | **Put** /pulp/api/v3/contentguards/{pulp_id}/ | Update a content guard



## ContentguardsCreate

> ContentGuardResponse ContentguardsCreate(ctx).ContentGuard(contentGuard).Execute()

Create a content guard



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
    contentGuard := *openapiclient.NewContentGuard("Name_example") // ContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsCreate(context.Background()).ContentGuard(contentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCreate`: ContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsApi.ContentguardsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentGuard** | [**ContentGuard**](ContentGuard.md) |  | 

### Return type

[**ContentGuardResponse**](ContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsDelete

> ContentguardsDelete(ctx, pulpId).Execute()

Delete a content guard



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this content guard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsList

> PaginatedContentGuardResponseList ContentguardsList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List content guards



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsList`: PaginatedContentGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsApi.ContentguardsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsListRequest struct via the builder pattern


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
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentGuardResponseList**](PaginatedContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsPartialUpdate

> ContentGuardResponse ContentguardsPartialUpdate(ctx, pulpId).PatchedContentGuard(patchedContentGuard).Execute()

Update a content guard



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this content guard.
    patchedContentGuard := *openapiclient.NewPatchedContentGuard() // PatchedContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsPartialUpdate(context.Background(), pulpId).PatchedContentGuard(patchedContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsPartialUpdate`: ContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsApi.ContentguardsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedContentGuard** | [**PatchedContentGuard**](PatchedContentGuard.md) |  | 

### Return type

[**ContentGuardResponse**](ContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsRead

> ContentGuardResponse ContentguardsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a content guard



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsRead`: ContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsApi.ContentguardsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContentGuardResponse**](ContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsUpdate

> ContentGuardResponse ContentguardsUpdate(ctx, pulpId).ContentGuard(contentGuard).Execute()

Update a content guard



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this content guard.
    contentGuard := *openapiclient.NewContentGuard("Name_example") // ContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsApi.ContentguardsUpdate(context.Background(), pulpId).ContentGuard(contentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsApi.ContentguardsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsUpdate`: ContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsApi.ContentguardsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentGuard** | [**ContentGuard**](ContentGuard.md) |  | 

### Return type

[**ContentGuardResponse**](ContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

