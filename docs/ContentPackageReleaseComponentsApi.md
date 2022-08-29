# \ContentPackageReleaseComponentsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebPackageReleaseComponentsCreate**](ContentPackageReleaseComponentsApi.md#ContentDebPackageReleaseComponentsCreate) | **Post** /pulp/api/v3/content/deb/package_release_components/ | Create a package release component
[**ContentDebPackageReleaseComponentsList**](ContentPackageReleaseComponentsApi.md#ContentDebPackageReleaseComponentsList) | **Get** /pulp/api/v3/content/deb/package_release_components/ | List package release components
[**ContentDebPackageReleaseComponentsRead**](ContentPackageReleaseComponentsApi.md#ContentDebPackageReleaseComponentsRead) | **Get** /pulp/api/v3/content/deb/package_release_components/{pulp_id}/ | Inspect a package release component



## ContentDebPackageReleaseComponentsCreate

> DebPackageReleaseComponentResponse ContentDebPackageReleaseComponentsCreate(ctx).DebPackageReleaseComponent(debPackageReleaseComponent).Execute()

Create a package release component



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
    debPackageReleaseComponent := *openapiclient.NewDebPackageReleaseComponent("Package_example", "ReleaseComponent_example") // DebPackageReleaseComponent | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsCreate(context.Background()).DebPackageReleaseComponent(debPackageReleaseComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsCreate`: DebPackageReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debPackageReleaseComponent** | [**DebPackageReleaseComponent**](DebPackageReleaseComponent.md) |  | 

### Return type

[**DebPackageReleaseComponentResponse**](deb.PackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageReleaseComponentsList

> PaginateddebPackageReleaseComponentResponseList ContentDebPackageReleaseComponentsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Package_(package_).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List package release components



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
    package_ := TODO // string | Filter results where package matches value (optional)
    releaseComponent := TODO // string | Filter results where release_component matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Package_(package_).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsList`: PaginateddebPackageReleaseComponentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **package_** | [**string**](string.md) | Filter results where package matches value | 
 **releaseComponent** | [**string**](string.md) | Filter results where release_component matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebPackageReleaseComponentResponseList**](Paginateddeb.PackageReleaseComponentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackageReleaseComponentsRead

> DebPackageReleaseComponentResponse ContentDebPackageReleaseComponentsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package release component



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
    pulpId := TODO // string | A UUID string identifying this package release component.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackageReleaseComponentsRead`: DebPackageReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this package release component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackageReleaseComponentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebPackageReleaseComponentResponse**](deb.PackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

