# \ContentGenericContentsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebGenericContentsCreate**](ContentGenericContentsApi.md#ContentDebGenericContentsCreate) | **Post** /pulp/api/v3/content/deb/generic_contents/ | Create a generic content
[**ContentDebGenericContentsList**](ContentGenericContentsApi.md#ContentDebGenericContentsList) | **Get** /pulp/api/v3/content/deb/generic_contents/ | List generic contents
[**ContentDebGenericContentsRead**](ContentGenericContentsApi.md#ContentDebGenericContentsRead) | **Get** /pulp/api/v3/content/deb/generic_contents/{pulp_id}/ | Inspect a generic content



## ContentDebGenericContentsCreate

> AsyncOperationResponse ContentDebGenericContentsCreate(ctx).RelativePath(relativePath).Artifact(artifact).File(file).Repository(repository).Upload(upload).Execute()

Create a generic content



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
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsApi.ContentDebGenericContentsCreate(context.Background()).RelativePath(relativePath).Artifact(artifact).File(file).Repository(repository).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsApi.ContentDebGenericContentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsApi.ContentDebGenericContentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebGenericContentsList

> PaginateddebGenericContentResponseList ContentDebGenericContentsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List generic contents



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
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsApi.ContentDebGenericContentsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsApi.ContentDebGenericContentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsList`: PaginateddebGenericContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsApi.ContentDebGenericContentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebGenericContentResponseList**](PaginateddebGenericContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebGenericContentsRead

> DebGenericContentResponse ContentDebGenericContentsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a generic content



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this generic content.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsApi.ContentDebGenericContentsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsApi.ContentDebGenericContentsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsRead`: DebGenericContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsApi.ContentDebGenericContentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this generic content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebGenericContentResponse**](DebGenericContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

