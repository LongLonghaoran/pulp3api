# \ContentReleaseComponentsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleaseComponentsCreate**](ContentReleaseComponentsApi.md#ContentDebReleaseComponentsCreate) | **Post** /pulp/api/v3/content/deb/release_components/ | Create a release component
[**ContentDebReleaseComponentsList**](ContentReleaseComponentsApi.md#ContentDebReleaseComponentsList) | **Get** /pulp/api/v3/content/deb/release_components/ | List release components
[**ContentDebReleaseComponentsRead**](ContentReleaseComponentsApi.md#ContentDebReleaseComponentsRead) | **Get** /pulp/api/v3/content/deb/release_components/{pulp_id}/ | Inspect a release component



## ContentDebReleaseComponentsCreate

> DebReleaseComponentResponse ContentDebReleaseComponentsCreate(ctx).DebReleaseComponent(debReleaseComponent).Execute()

Create a release component



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
    debReleaseComponent := *openapiclient.NewDebReleaseComponent("Component_example", "Release_example") // DebReleaseComponent | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentReleaseComponentsApi.ContentDebReleaseComponentsCreate(context.Background()).DebReleaseComponent(debReleaseComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseComponentsApi.ContentDebReleaseComponentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseComponentsCreate`: DebReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseComponentsApi.ContentDebReleaseComponentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseComponentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debReleaseComponent** | [**DebReleaseComponent**](DebReleaseComponent.md) |  | 

### Return type

[**DebReleaseComponentResponse**](deb.ReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseComponentsList

> PaginateddebReleaseComponentResponseList ContentDebReleaseComponentsList(ctx).Component(component).Limit(limit).Offset(offset).Ordering(ordering).Release(release).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List release components



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
    component := "component_example" // string | Filter results where component matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    release := TODO // string | Filter results where release matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentReleaseComponentsApi.ContentDebReleaseComponentsList(context.Background()).Component(component).Limit(limit).Offset(offset).Ordering(ordering).Release(release).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseComponentsApi.ContentDebReleaseComponentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseComponentsList`: PaginateddebReleaseComponentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseComponentsApi.ContentDebReleaseComponentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseComponentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **component** | **string** | Filter results where component matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **release** | [**string**](string.md) | Filter results where release matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseComponentResponseList**](Paginateddeb.ReleaseComponentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseComponentsRead

> DebReleaseComponentResponse ContentDebReleaseComponentsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release component



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
    pulpId := TODO // string | A UUID string identifying this release component.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentReleaseComponentsApi.ContentDebReleaseComponentsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseComponentsApi.ContentDebReleaseComponentsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseComponentsRead`: DebReleaseComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseComponentsApi.ContentDebReleaseComponentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this release component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseComponentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseComponentResponse**](deb.ReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

