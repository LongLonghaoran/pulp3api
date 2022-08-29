# \ContentReleasesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleasesCreate**](ContentReleasesApi.md#ContentDebReleasesCreate) | **Post** /pulp/api/v3/content/deb/releases/ | Create a release
[**ContentDebReleasesList**](ContentReleasesApi.md#ContentDebReleasesList) | **Get** /pulp/api/v3/content/deb/releases/ | List releases
[**ContentDebReleasesRead**](ContentReleasesApi.md#ContentDebReleasesRead) | **Get** /pulp/api/v3/content/deb/releases/{pulp_id}/ | Inspect a release



## ContentDebReleasesCreate

> DebReleaseResponse ContentDebReleasesCreate(ctx).DebRelease(debRelease).Execute()

Create a release



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
    debRelease := *openapiclient.NewDebRelease("Codename_example", "Suite_example", "Distribution_example") // DebRelease | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleasesApi.ContentDebReleasesCreate(context.Background()).DebRelease(debRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleasesApi.ContentDebReleasesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleasesCreate`: DebReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleasesApi.ContentDebReleasesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleasesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debRelease** | [**DebRelease**](DebRelease.md) |  | 

### Return type

[**DebReleaseResponse**](DebReleaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleasesList

> PaginateddebReleaseResponseList ContentDebReleasesList(ctx).Codename(codename).Distribution(distribution).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()

List releases



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
    codename := "codename_example" // string | Filter results where codename matches value (optional)
    distribution := "distribution_example" // string | Filter results where distribution matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    suite := "suite_example" // string | Filter results where suite matches value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleasesApi.ContentDebReleasesList(context.Background()).Codename(codename).Distribution(distribution).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleasesApi.ContentDebReleasesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleasesList`: PaginateddebReleaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleasesApi.ContentDebReleasesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleasesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** | Filter results where codename matches value | 
 **distribution** | **string** | Filter results where distribution matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **suite** | **string** | Filter results where suite matches value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseResponseList**](PaginateddebReleaseResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleasesRead

> DebReleaseResponse ContentDebReleasesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this release.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleasesApi.ContentDebReleasesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleasesApi.ContentDebReleasesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleasesRead`: DebReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleasesApi.ContentDebReleasesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleasesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseResponse**](DebReleaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

