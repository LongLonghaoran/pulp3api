# \ContentTagsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerTagsList**](ContentTagsApi.md#ContentContainerTagsList) | **Get** /pulp/api/v3/content/container/tags/ | List tags
[**ContentContainerTagsRead**](ContentTagsApi.md#ContentContainerTagsRead) | **Get** /pulp/api/v3/content/container/tags/{pulp_id}/ | Inspect a tag



## ContentContainerTagsList

> PaginatedcontainerTagResponseList ContentContainerTagsList(ctx).Digest(digest).Limit(limit).MediaType(mediaType).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List tags



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
    digest := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    mediaType := []string{"MediaType_example"} // []string |  (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentTagsApi.ContentContainerTagsList(context.Background()).Digest(digest).Limit(limit).MediaType(mediaType).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentTagsApi.ContentContainerTagsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerTagsList`: PaginatedcontainerTagResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentTagsApi.ContentContainerTagsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerTagsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digest** | **[]string** | Multiple values may be separated by commas. | 
 **limit** | **int32** | Number of results to return per page. | 
 **mediaType** | **[]string** |  | 
 **name** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerTagResponseList**](PaginatedcontainerTagResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerTagsRead

> ContainerTagResponse ContentContainerTagsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a tag



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tag.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentTagsApi.ContentContainerTagsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentTagsApi.ContentContainerTagsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentContainerTagsRead`: ContainerTagResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentTagsApi.ContentContainerTagsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerTagsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerTagResponse**](ContainerTagResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

