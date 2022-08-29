# \ExportersFilesystemExportsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCoreFilesystemExportsCreate**](ExportersFilesystemExportsApi.md#ExportersCoreFilesystemExportsCreate) | **Post** /pulp/api/v3/exporters/core/filesystem/{exporter_pk}/exports/ | Create a filesystem export
[**ExportersCoreFilesystemExportsDelete**](ExportersFilesystemExportsApi.md#ExportersCoreFilesystemExportsDelete) | **Delete** /pulp/api/v3/exporters/core/filesystem/{exporter_pk}/exports/{pulp_id}/ | Delete a filesystem export
[**ExportersCoreFilesystemExportsList**](ExportersFilesystemExportsApi.md#ExportersCoreFilesystemExportsList) | **Get** /pulp/api/v3/exporters/core/filesystem/{exporter_pk}/exports/ | List filesystem exports
[**ExportersCoreFilesystemExportsRead**](ExportersFilesystemExportsApi.md#ExportersCoreFilesystemExportsRead) | **Get** /pulp/api/v3/exporters/core/filesystem/{exporter_pk}/exports/{pulp_id}/ | Inspect a filesystem export



## ExportersCoreFilesystemExportsCreate

> AsyncOperationResponse ExportersCoreFilesystemExportsCreate(ctx, exporterPk).FilesystemExport(filesystemExport).Execute()

Create a filesystem export



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
    filesystemExport := *openapiclient.NewFilesystemExport() // FilesystemExport | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsCreate(context.Background(), exporterPk).FilesystemExport(filesystemExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filesystemExport** | [**FilesystemExport**](FilesystemExport.md) |  | 

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


## ExportersCoreFilesystemExportsDelete

> ExportersCoreFilesystemExportsDelete(ctx, exporterPk, pulpId).Execute()

Delete a filesystem export



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
    pulpId := TODO // string | A UUID string identifying this filesystem export.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsDelete(context.Background(), exporterPk, pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsDeleteRequest struct via the builder pattern


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


## ExportersCoreFilesystemExportsList

> PaginatedFilesystemExportResponseList ExportersCoreFilesystemExportsList(ctx, exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List filesystem exports



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
    resp, r, err := api_client.ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsList(context.Background(), exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsList`: PaginatedFilesystemExportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedFilesystemExportResponseList**](PaginatedFilesystemExportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemExportsRead

> FilesystemExportResponse ExportersCoreFilesystemExportsRead(ctx, exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a filesystem export



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
    pulpId := TODO // string | A UUID string identifying this filesystem export.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsRead(context.Background(), exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemExportsRead`: FilesystemExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemExportsApi.ExportersCoreFilesystemExportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemExportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**FilesystemExportResponse**](FilesystemExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

