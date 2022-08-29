# \ContentApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCreate**](ContentApi.md#ContentCreate) | **Post** /pulp/api/v3/content/ | Create a content
[**ContentList**](ContentApi.md#ContentList) | **Get** /pulp/api/v3/content/ | List content
[**ContentRead**](ContentApi.md#ContentRead) | **Get** /pulp/api/v3/content/{pulp_id}/ | Inspect a content



## ContentCreate

> MultipleArtifactContentResponse ContentCreate(ctx).MultipleArtifactContent(multipleArtifactContent).Execute()

Create a content



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
    multipleArtifactContent := *openapiclient.NewMultipleArtifactContent(map[string]interface{}(123)) // MultipleArtifactContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentCreate(context.Background()).MultipleArtifactContent(multipleArtifactContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentCreate`: MultipleArtifactContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleArtifactContent** | [**MultipleArtifactContent**](MultipleArtifactContent.md) |  | 

### Return type

[**MultipleArtifactContentResponse**](MultipleArtifactContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentList

> PaginatedMultipleArtifactContentResponseList ContentList(ctx).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List content



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentList`: PaginatedMultipleArtifactContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedMultipleArtifactContentResponseList**](PaginatedMultipleArtifactContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRead

> MultipleArtifactContentResponse ContentRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a content



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this content.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRead`: MultipleArtifactContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**MultipleArtifactContentResponse**](MultipleArtifactContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

