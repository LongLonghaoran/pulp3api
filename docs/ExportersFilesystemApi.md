# \ExportersFilesystemApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCoreFilesystemCreate**](ExportersFilesystemApi.md#ExportersCoreFilesystemCreate) | **Post** /pulp/api/v3/exporters/core/filesystem/ | Create a filesystem exporter
[**ExportersCoreFilesystemDelete**](ExportersFilesystemApi.md#ExportersCoreFilesystemDelete) | **Delete** /pulp/api/v3/exporters/core/filesystem/{pulp_id}/ | Delete a filesystem exporter
[**ExportersCoreFilesystemList**](ExportersFilesystemApi.md#ExportersCoreFilesystemList) | **Get** /pulp/api/v3/exporters/core/filesystem/ | List filesystem exporters
[**ExportersCoreFilesystemPartialUpdate**](ExportersFilesystemApi.md#ExportersCoreFilesystemPartialUpdate) | **Patch** /pulp/api/v3/exporters/core/filesystem/{pulp_id}/ | Update a filesystem exporter
[**ExportersCoreFilesystemRead**](ExportersFilesystemApi.md#ExportersCoreFilesystemRead) | **Get** /pulp/api/v3/exporters/core/filesystem/{pulp_id}/ | Inspect a filesystem exporter
[**ExportersCoreFilesystemUpdate**](ExportersFilesystemApi.md#ExportersCoreFilesystemUpdate) | **Put** /pulp/api/v3/exporters/core/filesystem/{pulp_id}/ | Update a filesystem exporter



## ExportersCoreFilesystemCreate

> FilesystemExporterResponse ExportersCoreFilesystemCreate(ctx).FilesystemExporter(filesystemExporter).Execute()

Create a filesystem exporter



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
    filesystemExporter := *openapiclient.NewFilesystemExporter("Name_example", "Path_example") // FilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemCreate(context.Background()).FilesystemExporter(filesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemCreate`: FilesystemExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemExporter** | [**FilesystemExporter**](FilesystemExporter.md) |  | 

### Return type

[**FilesystemExporterResponse**](FilesystemExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemDelete

> AsyncOperationResponse ExportersCoreFilesystemDelete(ctx, pulpId).Execute()

Delete a filesystem exporter



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
    pulpId := TODO // string | A UUID string identifying this filesystem exporter.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem exporter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemDeleteRequest struct via the builder pattern


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


## ExportersCoreFilesystemList

> PaginatedFilesystemExporterResponseList ExportersCoreFilesystemList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List filesystem exporters



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
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemList`: PaginatedFilesystemExporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemListRequest struct via the builder pattern


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
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedFilesystemExporterResponseList**](PaginatedFilesystemExporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemPartialUpdate

> AsyncOperationResponse ExportersCoreFilesystemPartialUpdate(ctx, pulpId).PatchedFilesystemExporter(patchedFilesystemExporter).Execute()

Update a filesystem exporter



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
    pulpId := TODO // string | A UUID string identifying this filesystem exporter.
    patchedFilesystemExporter := *openapiclient.NewPatchedFilesystemExporter() // PatchedFilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemPartialUpdate(context.Background(), pulpId).PatchedFilesystemExporter(patchedFilesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem exporter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFilesystemExporter** | [**PatchedFilesystemExporter**](PatchedFilesystemExporter.md) |  | 

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


## ExportersCoreFilesystemRead

> FilesystemExporterResponse ExportersCoreFilesystemRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a filesystem exporter



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
    pulpId := TODO // string | A UUID string identifying this filesystem exporter.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemRead`: FilesystemExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem exporter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**FilesystemExporterResponse**](FilesystemExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCoreFilesystemUpdate

> AsyncOperationResponse ExportersCoreFilesystemUpdate(ctx, pulpId).FilesystemExporter(filesystemExporter).Execute()

Update a filesystem exporter



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
    pulpId := TODO // string | A UUID string identifying this filesystem exporter.
    filesystemExporter := *openapiclient.NewFilesystemExporter("Name_example", "Path_example") // FilesystemExporter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportersFilesystemApi.ExportersCoreFilesystemUpdate(context.Background(), pulpId).FilesystemExporter(filesystemExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersFilesystemApi.ExportersCoreFilesystemUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCoreFilesystemUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersFilesystemApi.ExportersCoreFilesystemUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this filesystem exporter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCoreFilesystemUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filesystemExporter** | [**FilesystemExporter**](FilesystemExporter.md) |  | 

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

