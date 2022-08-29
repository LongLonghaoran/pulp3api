# \RepositoriesAptApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesDebAptCreate**](RepositoriesAptApi.md#RepositoriesDebAptCreate) | **Post** /pulp/api/v3/repositories/deb/apt/ | Create an apt repository
[**RepositoriesDebAptDelete**](RepositoriesAptApi.md#RepositoriesDebAptDelete) | **Delete** /pulp/api/v3/repositories/deb/apt/{pulp_id}/ | Delete an apt repository
[**RepositoriesDebAptList**](RepositoriesAptApi.md#RepositoriesDebAptList) | **Get** /pulp/api/v3/repositories/deb/apt/ | List apt repositorys
[**RepositoriesDebAptModify**](RepositoriesAptApi.md#RepositoriesDebAptModify) | **Post** /pulp/api/v3/repositories/deb/apt/{pulp_id}/modify/ | Modify Repository Content
[**RepositoriesDebAptPartialUpdate**](RepositoriesAptApi.md#RepositoriesDebAptPartialUpdate) | **Patch** /pulp/api/v3/repositories/deb/apt/{pulp_id}/ | Update an apt repository
[**RepositoriesDebAptRead**](RepositoriesAptApi.md#RepositoriesDebAptRead) | **Get** /pulp/api/v3/repositories/deb/apt/{pulp_id}/ | Inspect an apt repository
[**RepositoriesDebAptSync**](RepositoriesAptApi.md#RepositoriesDebAptSync) | **Post** /pulp/api/v3/repositories/deb/apt/{pulp_id}/sync/ | Sync from remote
[**RepositoriesDebAptUpdate**](RepositoriesAptApi.md#RepositoriesDebAptUpdate) | **Put** /pulp/api/v3/repositories/deb/apt/{pulp_id}/ | Update an apt repository



## RepositoriesDebAptCreate

> DebAptRepositoryResponse RepositoriesDebAptCreate(ctx).DebAptRepository(debAptRepository).Execute()

Create an apt repository



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
    debAptRepository := *openapiclient.NewDebAptRepository("Name_example") // DebAptRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptCreate(context.Background()).DebAptRepository(debAptRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptCreate`: DebAptRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debAptRepository** | [**DebAptRepository**](DebAptRepository.md) |  | 

### Return type

[**DebAptRepositoryResponse**](DebAptRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptDelete

> AsyncOperationResponse RepositoriesDebAptDelete(ctx, pulpId).Execute()

Delete an apt repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptList

> PaginateddebAptRepositoryResponseList RepositoriesDebAptList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()

List apt repositorys



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
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptList`: PaginateddebAptRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptListRequest struct via the builder pattern


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
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebAptRepositoryResponseList**](PaginateddebAptRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptModify

> AsyncOperationResponse RepositoriesDebAptModify(ctx, pulpId).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify Repository Content



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.
    repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptModify(context.Background(), pulpId).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptModify`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptPartialUpdate

> AsyncOperationResponse RepositoriesDebAptPartialUpdate(ctx, pulpId).PatcheddebAptRepository(patcheddebAptRepository).Execute()

Update an apt repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.
    patcheddebAptRepository := *openapiclient.NewPatcheddebAptRepository() // PatcheddebAptRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptPartialUpdate(context.Background(), pulpId).PatcheddebAptRepository(patcheddebAptRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patcheddebAptRepository** | [**PatcheddebAptRepository**](PatcheddebAptRepository.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptRead

> DebAptRepositoryResponse RepositoriesDebAptRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an apt repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptRead`: DebAptRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebAptRepositoryResponse**](DebAptRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptSync

> AsyncOperationResponse RepositoriesDebAptSync(ctx, pulpId).RepositorySyncURL(repositorySyncURL).Execute()

Sync from remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.
    repositorySyncURL := *openapiclient.NewRepositorySyncURL() // RepositorySyncURL | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptSync(context.Background(), pulpId).RepositorySyncURL(repositorySyncURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptSync`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySyncURL** | [**RepositorySyncURL**](RepositorySyncURL.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptUpdate

> AsyncOperationResponse RepositoriesDebAptUpdate(ctx, pulpId).DebAptRepository(debAptRepository).Execute()

Update an apt repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this apt repository.
    debAptRepository := *openapiclient.NewDebAptRepository("Name_example") // DebAptRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptApi.RepositoriesDebAptUpdate(context.Background(), pulpId).DebAptRepository(debAptRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptApi.RepositoriesDebAptUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptApi.RepositoriesDebAptUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this apt repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **debAptRepository** | [**DebAptRepository**](DebAptRepository.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

