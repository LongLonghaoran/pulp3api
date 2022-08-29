# \ContentInstallerFileIndicesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebInstallerFileIndicesCreate**](ContentInstallerFileIndicesApi.md#ContentDebInstallerFileIndicesCreate) | **Post** /pulp/api/v3/content/deb/installer_file_indices/ | Create an installer file index
[**ContentDebInstallerFileIndicesList**](ContentInstallerFileIndicesApi.md#ContentDebInstallerFileIndicesList) | **Get** /pulp/api/v3/content/deb/installer_file_indices/ | List InstallerFileIndices
[**ContentDebInstallerFileIndicesRead**](ContentInstallerFileIndicesApi.md#ContentDebInstallerFileIndicesRead) | **Get** /pulp/api/v3/content/deb/installer_file_indices/{pulp_id}/ | Inspect an installer file index



## ContentDebInstallerFileIndicesCreate

> DebInstallerFileIndexResponse ContentDebInstallerFileIndicesCreate(ctx).DebInstallerFileIndex(debInstallerFileIndex).Execute()

Create an installer file index



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
    debInstallerFileIndex := *openapiclient.NewDebInstallerFileIndex(map[string]interface{}(123), "Release_example", "Component_example", "Architecture_example") // DebInstallerFileIndex | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesCreate(context.Background()).DebInstallerFileIndex(debInstallerFileIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesCreate`: DebInstallerFileIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debInstallerFileIndex** | [**DebInstallerFileIndex**](DebInstallerFileIndex.md) |  | 

### Return type

[**DebInstallerFileIndexResponse**](deb.InstallerFileIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerFileIndicesList

> PaginateddebInstallerFileIndexResponseList ContentDebInstallerFileIndicesList(ctx).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List InstallerFileIndices



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
    resp, r, err := api_client.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesList(context.Background()).Architecture(architecture).Component(component).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesList`: PaginateddebInstallerFileIndexResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesListRequest struct via the builder pattern


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

[**PaginateddebInstallerFileIndexResponseList**](Paginateddeb.InstallerFileIndexResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerFileIndicesRead

> DebInstallerFileIndexResponse ContentDebInstallerFileIndicesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an installer file index



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
    pulpId := TODO // string | A UUID string identifying this installer file index.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerFileIndicesRead`: DebInstallerFileIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this installer file index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerFileIndicesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebInstallerFileIndexResponse**](deb.InstallerFileIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

