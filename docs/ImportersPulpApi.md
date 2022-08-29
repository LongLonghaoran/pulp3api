# \ImportersPulpApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportersCorePulpCreate**](ImportersPulpApi.md#ImportersCorePulpCreate) | **Post** /pulp/api/v3/importers/core/pulp/ | Create a pulp importer
[**ImportersCorePulpDelete**](ImportersPulpApi.md#ImportersCorePulpDelete) | **Delete** /pulp/api/v3/importers/core/pulp/{pulp_id}/ | Delete a pulp importer
[**ImportersCorePulpList**](ImportersPulpApi.md#ImportersCorePulpList) | **Get** /pulp/api/v3/importers/core/pulp/ | List pulp importers
[**ImportersCorePulpPartialUpdate**](ImportersPulpApi.md#ImportersCorePulpPartialUpdate) | **Patch** /pulp/api/v3/importers/core/pulp/{pulp_id}/ | Update a pulp importer
[**ImportersCorePulpRead**](ImportersPulpApi.md#ImportersCorePulpRead) | **Get** /pulp/api/v3/importers/core/pulp/{pulp_id}/ | Inspect a pulp importer
[**ImportersCorePulpUpdate**](ImportersPulpApi.md#ImportersCorePulpUpdate) | **Put** /pulp/api/v3/importers/core/pulp/{pulp_id}/ | Update a pulp importer



## ImportersCorePulpCreate

> PulpImporterResponse ImportersCorePulpCreate(ctx).PulpImporter(pulpImporter).Execute()

Create a pulp importer



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
    pulpImporter := *openapiclient.NewPulpImporter("Name_example") // PulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpCreate(context.Background()).PulpImporter(pulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpCreate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpApi.ImportersCorePulpCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pulpImporter** | [**PulpImporter**](PulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpDelete

> ImportersCorePulpDelete(ctx, pulpId).Execute()

Delete a pulp importer



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this pulp importer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this pulp importer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpDeleteRequest struct via the builder pattern


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


## ImportersCorePulpList

> PaginatedPulpImporterResponseList ImportersCorePulpList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp importers



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpList`: PaginatedPulpImporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpApi.ImportersCorePulpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpListRequest struct via the builder pattern


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

[**PaginatedPulpImporterResponseList**](PaginatedPulpImporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpPartialUpdate

> PulpImporterResponse ImportersCorePulpPartialUpdate(ctx, pulpId).PatchedPulpImporter(patchedPulpImporter).Execute()

Update a pulp importer



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this pulp importer.
    patchedPulpImporter := *openapiclient.NewPatchedPulpImporter() // PatchedPulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpPartialUpdate(context.Background(), pulpId).PatchedPulpImporter(patchedPulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpPartialUpdate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpApi.ImportersCorePulpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this pulp importer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPulpImporter** | [**PatchedPulpImporter**](PatchedPulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpRead

> PulpImporterResponse ImportersCorePulpRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp importer



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this pulp importer.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpRead`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpApi.ImportersCorePulpRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this pulp importer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpUpdate

> PulpImporterResponse ImportersCorePulpUpdate(ctx, pulpId).PulpImporter(pulpImporter).Execute()

Update a pulp importer



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this pulp importer.
    pulpImporter := *openapiclient.NewPulpImporter("Name_example") // PulpImporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpApi.ImportersCorePulpUpdate(context.Background(), pulpId).PulpImporter(pulpImporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpApi.ImportersCorePulpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpUpdate`: PulpImporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpApi.ImportersCorePulpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this pulp importer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpImporter** | [**PulpImporter**](PulpImporter.md) |  | 

### Return type

[**PulpImporterResponse**](PulpImporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

