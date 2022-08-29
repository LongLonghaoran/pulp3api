# \AcsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcsCreate**](AcsApi.md#AcsCreate) | **Post** /pulp/api/v3/acs/ | Create an alternate content source
[**AcsDelete**](AcsApi.md#AcsDelete) | **Delete** /pulp/api/v3/acs/{pulp_id}/ | Delete an alternate content source
[**AcsList**](AcsApi.md#AcsList) | **Get** /pulp/api/v3/acs/ | List acs
[**AcsPartialUpdate**](AcsApi.md#AcsPartialUpdate) | **Patch** /pulp/api/v3/acs/{pulp_id}/ | Update an alternate content source
[**AcsRead**](AcsApi.md#AcsRead) | **Get** /pulp/api/v3/acs/{pulp_id}/ | Inspect an alternate content source
[**AcsRefresh**](AcsApi.md#AcsRefresh) | **Post** /pulp/api/v3/acs/{pulp_id}/refresh/ | 
[**AcsUpdate**](AcsApi.md#AcsUpdate) | **Put** /pulp/api/v3/acs/{pulp_id}/ | Update an alternate content source



## AcsCreate

> AlternateContentSourceResponse AcsCreate(ctx).AlternateContentSource(alternateContentSource).Execute()

Create an alternate content source



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
    alternateContentSource := *openapiclient.NewAlternateContentSource("Name_example", "Remote_example") // AlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsCreate(context.Background()).AlternateContentSource(alternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsCreate`: AlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alternateContentSource** | [**AlternateContentSource**](AlternateContentSource.md) |  | 

### Return type

[**AlternateContentSourceResponse**](AlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsDelete

> AsyncOperationResponse AcsDelete(ctx, pulpId).Execute()

Delete an alternate content source



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this alternate content source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this alternate content source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsDeleteRequest struct via the builder pattern


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


## AcsList

> PaginatedAlternateContentSourceResponseList AcsList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List acs



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
    resp, r, err := apiClient.AcsApi.AcsList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsList`: PaginatedAlternateContentSourceResponseList
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsListRequest struct via the builder pattern


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

[**PaginatedAlternateContentSourceResponseList**](PaginatedAlternateContentSourceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsPartialUpdate

> AsyncOperationResponse AcsPartialUpdate(ctx, pulpId).PatchedAlternateContentSource(patchedAlternateContentSource).Execute()

Update an alternate content source



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this alternate content source.
    patchedAlternateContentSource := *openapiclient.NewPatchedAlternateContentSource() // PatchedAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsPartialUpdate(context.Background(), pulpId).PatchedAlternateContentSource(patchedAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this alternate content source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAlternateContentSource** | [**PatchedAlternateContentSource**](PatchedAlternateContentSource.md) |  | 

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


## AcsRead

> AlternateContentSourceResponse AcsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an alternate content source



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this alternate content source.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRead`: AlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this alternate content source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**AlternateContentSourceResponse**](AlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRefresh

> AlternateContentSourceResponse AcsRefresh(ctx, pulpId).AlternateContentSource(alternateContentSource).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this alternate content source.
    alternateContentSource := *openapiclient.NewAlternateContentSource("Name_example", "Remote_example") // AlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsRefresh(context.Background(), pulpId).AlternateContentSource(alternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRefresh`: AlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this alternate content source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alternateContentSource** | [**AlternateContentSource**](AlternateContentSource.md) |  | 

### Return type

[**AlternateContentSourceResponse**](AlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsUpdate

> AsyncOperationResponse AcsUpdate(ctx, pulpId).AlternateContentSource(alternateContentSource).Execute()

Update an alternate content source



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this alternate content source.
    alternateContentSource := *openapiclient.NewAlternateContentSource("Name_example", "Remote_example") // AlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsApi.AcsUpdate(context.Background(), pulpId).AlternateContentSource(alternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsApi.AcsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsApi.AcsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this alternate content source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alternateContentSource** | [**AlternateContentSource**](AlternateContentSource.md) |  | 

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

