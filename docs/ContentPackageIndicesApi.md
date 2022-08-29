# \ContentPackageIndicesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebPackageIndicesCreate**](ContentPackageIndicesApi.md#ContentDebPackageIndicesCreate) | **Post** /pulp/api/v3/content/deb/package_indices/ | Create a package index
[**ContentDebPackageIndicesList**](ContentPackageIndicesApi.md#ContentDebPackageIndicesList) | **Get** /pulp/api/v3/content/deb/package_indices/ | List PackageIndices
[**ContentDebPackageIndicesRead**](ContentPackageIndicesApi.md#ContentDebPackageIndicesRead) | **Get** /pulp/api/v3/content/deb/package_indices/{pulp_id}/ | Inspect a package index



## ContentDebPackageIndicesCreate

> DebPackageIndexResponse ContentDebPackageIndicesCreate(ctx).DebPackageIndex(debPackageIndex).Execute()

Create a package index



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
    debPackageIndex := *openapiclient.NewDebPackageIndex(map[string]interface{}(123), "Release_example") // DebPackageIndex | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageIndicesApi.ContentDebPackageIndicesCreate(context.Background()).DebPackageIndex(debPackageIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageIndicesApi.ContentDebPackageIndicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageIndicesCreate`: DebPackageIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageIndicesApi.ContentDebPackageIndicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageIndicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debPackageIndex** | [**DebPackageIndex**](DebPackageIndex.md) |  | 

### Return type

[**DebPackageIndexResponse**](deb.PackageIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageIndicesList

> PaginateddebPackageIndexResponseList ContentDebPackageIndicesList(ctx).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List PackageIndices



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
    architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
    component := "component_example" // string | Filter results where component matches value (optional)
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageIndicesApi.ContentDebPackageIndicesList(context.Background()).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageIndicesApi.ContentDebPackageIndicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageIndicesList`: PaginateddebPackageIndexResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageIndicesApi.ContentDebPackageIndicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageIndicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **component** | **string** | Filter results where component matches value | 
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

[**PaginateddebPackageIndexResponseList**](Paginateddeb.PackageIndexResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageIndicesRead

> DebPackageIndexResponse ContentDebPackageIndicesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package index



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
    pulpId := TODO // string | A UUID string identifying this package index.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageIndicesApi.ContentDebPackageIndicesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageIndicesApi.ContentDebPackageIndicesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageIndicesRead`: DebPackageIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageIndicesApi.ContentDebPackageIndicesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this package index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageIndicesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebPackageIndexResponse**](deb.PackageIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

