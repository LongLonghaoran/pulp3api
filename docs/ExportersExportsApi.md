# \ExportersExportsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersExportsCreate**](ExportersExportsApi.md#ExportersExportsCreate) | **Post** /pulp/api/v3/exporters/{exporter_pk}/exports/ | Create an export
[**ExportersExportsDelete**](ExportersExportsApi.md#ExportersExportsDelete) | **Delete** /pulp/api/v3/exporters/{exporter_pk}/exports/{pulp_id}/ | Delete an export
[**ExportersExportsList**](ExportersExportsApi.md#ExportersExportsList) | **Get** /pulp/api/v3/exporters/{exporter_pk}/exports/ | List exports
[**ExportersExportsRead**](ExportersExportsApi.md#ExportersExportsRead) | **Get** /pulp/api/v3/exporters/{exporter_pk}/exports/{pulp_id}/ | Inspect an export



## ExportersExportsCreate

> ExportResponse ExportersExportsCreate(ctx, exporterPk).Export(export).Execute()

Create an export



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
    exporterPk := "exporterPk_example" // string | 
    export := *openapiclient.NewExport() // Export | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersExportsApi.ExportersExportsCreate(context.Background(), exporterPk).Export(export).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersExportsApi.ExportersExportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersExportsCreate`: ExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersExportsApi.ExportersExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **export** | [**Export**](Export.md) |  | 

### Return type

[**ExportResponse**](ExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersExportsDelete

> ExportersExportsDelete(ctx, exporterPk, pulpId).Execute()

Delete an export



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
    exporterPk := "exporterPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this export.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersExportsApi.ExportersExportsDelete(context.Background(), exporterPk, pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersExportsApi.ExportersExportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersExportsDeleteRequest struct via the builder pattern


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


## ExportersExportsList

> PaginatedExportResponseList ExportersExportsList(ctx, exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List exports



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
    exporterPk := "exporterPk_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersExportsApi.ExportersExportsList(context.Background(), exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersExportsApi.ExportersExportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersExportsList`: PaginatedExportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersExportsApi.ExportersExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedExportResponseList**](PaginatedExportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersExportsRead

> ExportResponse ExportersExportsRead(ctx, exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an export



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
    exporterPk := "exporterPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this export.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersExportsApi.ExportersExportsRead(context.Background(), exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersExportsApi.ExportersExportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersExportsRead`: ExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersExportsApi.ExportersExportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersExportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ExportResponse**](ExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

