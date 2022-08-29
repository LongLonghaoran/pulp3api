# \RepositoriesAptVersionsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesDebAptVersionsDelete**](RepositoriesAptVersionsApi.md#RepositoriesDebAptVersionsDelete) | **Delete** /pulp/api/v3/repositories/deb/apt/{repository_pk}/versions/{number}/ | Delete a repository version
[**RepositoriesDebAptVersionsList**](RepositoriesAptVersionsApi.md#RepositoriesDebAptVersionsList) | **Get** /pulp/api/v3/repositories/deb/apt/{repository_pk}/versions/ | List repository versions
[**RepositoriesDebAptVersionsRead**](RepositoriesAptVersionsApi.md#RepositoriesDebAptVersionsRead) | **Get** /pulp/api/v3/repositories/deb/apt/{repository_pk}/versions/{number}/ | Inspect a repository version
[**RepositoriesDebAptVersionsRepair**](RepositoriesAptVersionsApi.md#RepositoriesDebAptVersionsRepair) | **Post** /pulp/api/v3/repositories/deb/apt/{repository_pk}/versions/{number}/repair/ | 



## RepositoriesDebAptVersionsDelete

> AsyncOperationResponse RepositoriesDebAptVersionsDelete(ctx, number, repositoryPk).Execute()

Delete a repository version



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
    number := int32(56) // int32 | 
    repositoryPk := "repositoryPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsDelete(context.Background(), number, repositoryPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **int32** |  | 
**repositoryPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptVersionsDeleteRequest struct via the builder pattern


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


## RepositoriesDebAptVersionsList

> PaginatedRepositoryVersionResponseList RepositoriesDebAptVersionsList(ctx, repositoryPk).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List repository versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    repositoryPk := "repositoryPk_example" // string | 
    content := "content_example" // string | Content Unit referenced by HREF (optional)
    contentIn := "contentIn_example" // string | Content Unit referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    number := int32(56) // int32 |  (optional)
    numberGt := int32(56) // int32 | Filter results where number is greater than value (optional)
    numberGte := int32(56) // int32 | Filter results where number is greater than or equal to value (optional)
    numberLt := int32(56) // int32 | Filter results where number is less than value (optional)
    numberLte := int32(56) // int32 | Filter results where number is less than or equal to value (optional)
    numberRange := []int32{int32(123)} // []int32 | Filter results where number is between two comma separated values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpCreated := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
    pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
    pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
    pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
    pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsList(context.Background(), repositoryPk).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptVersionsList`: PaginatedRepositoryVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Content Unit referenced by HREF | 
 **contentIn** | **string** | Content Unit referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **number** | **int32** |  | 
 **numberGt** | **int32** | Filter results where number is greater than value | 
 **numberGte** | **int32** | Filter results where number is greater than or equal to value | 
 **numberLt** | **int32** | Filter results where number is less than value | 
 **numberLte** | **int32** | Filter results where number is less than or equal to value | 
 **numberRange** | **[]int32** | Filter results where number is between two comma separated values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpCreated** | **time.Time** | ISO 8601 formatted dates are supported | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRepositoryVersionResponseList**](PaginatedRepositoryVersionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptVersionsRead

> RepositoryVersionResponse RepositoriesDebAptVersionsRead(ctx, number, repositoryPk).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a repository version



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
    number := int32(56) // int32 | 
    repositoryPk := "repositoryPk_example" // string | 
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRead(context.Background(), number, repositoryPk).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptVersionsRead`: RepositoryVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **int32** |  | 
**repositoryPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**RepositoryVersionResponse**](RepositoryVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesDebAptVersionsRepair

> AsyncOperationResponse RepositoriesDebAptVersionsRepair(ctx, number, repositoryPk).Repair(repair).Execute()





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
    number := int32(56) // int32 | 
    repositoryPk := "repositoryPk_example" // string | 
    repair := *openapiclient.NewRepair() // Repair | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRepair(context.Background(), number, repositoryPk).Repair(repair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRepair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesDebAptVersionsRepair`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRepair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **int32** |  | 
**repositoryPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesDebAptVersionsRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repair** | [**Repair**](Repair.md) |  | 

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

