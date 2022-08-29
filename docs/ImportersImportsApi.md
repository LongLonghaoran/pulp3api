# \ImportersImportsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportersImportsCreate**](ImportersImportsApi.md#ImportersImportsCreate) | **Post** /pulp/api/v3/importers/{importer_pk}/imports/ | Create an import
[**ImportersImportsDelete**](ImportersImportsApi.md#ImportersImportsDelete) | **Delete** /pulp/api/v3/importers/{importer_pk}/imports/{pulp_id}/ | Delete an import
[**ImportersImportsList**](ImportersImportsApi.md#ImportersImportsList) | **Get** /pulp/api/v3/importers/{importer_pk}/imports/ | List imports
[**ImportersImportsRead**](ImportersImportsApi.md#ImportersImportsRead) | **Get** /pulp/api/v3/importers/{importer_pk}/imports/{pulp_id}/ | Inspect an import



## ImportersImportsCreate

> ImportResponse ImportersImportsCreate(ctx, importerPk).Import_(import_).Execute()

Create an import



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
    importerPk := "importerPk_example" // string | 
    import_ := *openapiclient.NewImport("Task_example", map[string]interface{}(123)) // Import | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportersImportsApi.ImportersImportsCreate(context.Background(), importerPk).Import_(import_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersImportsApi.ImportersImportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersImportsCreate`: ImportResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersImportsApi.ImportersImportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importerPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersImportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **import_** | [**Import**](Import.md) |  | 

### Return type

[**ImportResponse**](ImportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersImportsDelete

> ImportersImportsDelete(ctx, importerPk, pulpId).Execute()

Delete an import



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
    importerPk := "importerPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this import.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportersImportsApi.ImportersImportsDelete(context.Background(), importerPk, pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersImportsApi.ImportersImportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importerPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersImportsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersImportsList

> PaginatedImportResponseList ImportersImportsList(ctx, importerPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List imports



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
    importerPk := "importerPk_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportersImportsApi.ImportersImportsList(context.Background(), importerPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersImportsApi.ImportersImportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersImportsList`: PaginatedImportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ImportersImportsApi.ImportersImportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importerPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersImportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedImportResponseList**](PaginatedImportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersImportsRead

> ImportResponse ImportersImportsRead(ctx, importerPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an import



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
    importerPk := "importerPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this import.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportersImportsApi.ImportersImportsRead(context.Background(), importerPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersImportsApi.ImportersImportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersImportsRead`: ImportResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersImportsApi.ImportersImportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importerPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersImportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ImportResponse**](ImportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

