# \ExportersPulpExportsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCorePulpExportsCreate**](ExportersPulpExportsApi.md#ExportersCorePulpExportsCreate) | **Post** /pulp/api/v3/exporters/core/pulp/{exporter_pk}/exports/ | Create a pulp export
[**ExportersCorePulpExportsDelete**](ExportersPulpExportsApi.md#ExportersCorePulpExportsDelete) | **Delete** /pulp/api/v3/exporters/core/pulp/{exporter_pk}/exports/{pulp_id}/ | Delete a pulp export
[**ExportersCorePulpExportsList**](ExportersPulpExportsApi.md#ExportersCorePulpExportsList) | **Get** /pulp/api/v3/exporters/core/pulp/{exporter_pk}/exports/ | List pulp exports
[**ExportersCorePulpExportsRead**](ExportersPulpExportsApi.md#ExportersCorePulpExportsRead) | **Get** /pulp/api/v3/exporters/core/pulp/{exporter_pk}/exports/{pulp_id}/ | Inspect a pulp export



## ExportersCorePulpExportsCreate

> AsyncOperationResponse ExportersCorePulpExportsCreate(ctx, exporterPk).PulpExport(pulpExport).Execute()

Create a pulp export



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
    pulpExport := *openapiclient.NewPulpExport() // PulpExport | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersPulpExportsApi.ExportersCorePulpExportsCreate(context.Background(), exporterPk).PulpExport(pulpExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsApi.ExportersCorePulpExportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsApi.ExportersCorePulpExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpExport** | [**PulpExport**](PulpExport.md) |  | 

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


## ExportersCorePulpExportsDelete

> ExportersCorePulpExportsDelete(ctx, exporterPk, pulpId).Execute()

Delete a pulp export



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
    pulpId := TODO // string | A UUID string identifying this pulp export.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersPulpExportsApi.ExportersCorePulpExportsDelete(context.Background(), exporterPk, pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsApi.ExportersCorePulpExportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this pulp export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsDeleteRequest struct via the builder pattern


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


## ExportersCorePulpExportsList

> PaginatedPulpExportResponseList ExportersCorePulpExportsList(ctx, exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp exports



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
    resp, r, err := api_client.ExportersPulpExportsApi.ExportersCorePulpExportsList(context.Background(), exporterPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsApi.ExportersCorePulpExportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsList`: PaginatedPulpExportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsApi.ExportersCorePulpExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedPulpExportResponseList**](PaginatedPulpExportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpExportsRead

> PulpExportResponse ExportersCorePulpExportsRead(ctx, exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp export



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
    pulpId := TODO // string | A UUID string identifying this pulp export.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersPulpExportsApi.ExportersCorePulpExportsRead(context.Background(), exporterPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpExportsApi.ExportersCorePulpExportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpExportsRead`: PulpExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpExportsApi.ExportersCorePulpExportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exporterPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this pulp export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpExportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PulpExportResponse**](PulpExportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

