# \ContentBlobsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerBlobsList**](ContentBlobsApi.md#ContentContainerBlobsList) | **Get** /pulp/api/v3/content/container/blobs/ | List blobs
[**ContentContainerBlobsRead**](ContentBlobsApi.md#ContentContainerBlobsRead) | **Get** /pulp/api/v3/content/container/blobs/{pulp_id}/ | Inspect a blob



## ContentContainerBlobsList

> PaginatedcontainerBlobResponseList ContentContainerBlobsList(ctx).Digest(digest).DigestIn(digestIn).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List blobs



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
    digest := "digest_example" // string | Filter results where digest matches value (optional)
    digestIn := []string{"Inner_example"} // []string | Filter results where digest is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentBlobsApi.ContentContainerBlobsList(context.Background()).Digest(digest).DigestIn(digestIn).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentBlobsApi.ContentContainerBlobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerBlobsList`: PaginatedcontainerBlobResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentBlobsApi.ContentContainerBlobsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerBlobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digest** | **string** | Filter results where digest matches value | 
 **digestIn** | **[]string** | Filter results where digest is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerBlobResponseList**](PaginatedcontainerBlobResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerBlobsRead

> ContainerBlobResponse ContentContainerBlobsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a blob



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this blob.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentBlobsApi.ContentContainerBlobsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentBlobsApi.ContentContainerBlobsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerBlobsRead`: ContainerBlobResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentBlobsApi.ContentContainerBlobsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this blob. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerBlobsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerBlobResponse**](ContainerBlobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

