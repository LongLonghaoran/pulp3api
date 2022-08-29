# \AccessPoliciesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessPoliciesList**](AccessPoliciesApi.md#AccessPoliciesList) | **Get** /pulp/api/v3/access_policies/ | List access policys
[**AccessPoliciesPartialUpdate**](AccessPoliciesApi.md#AccessPoliciesPartialUpdate) | **Patch** /pulp/api/v3/access_policies/{pulp_id}/ | Update an access policy
[**AccessPoliciesRead**](AccessPoliciesApi.md#AccessPoliciesRead) | **Get** /pulp/api/v3/access_policies/{pulp_id}/ | Inspect an access policy
[**AccessPoliciesReset**](AccessPoliciesApi.md#AccessPoliciesReset) | **Post** /pulp/api/v3/access_policies/{pulp_id}/reset/ | 
[**AccessPoliciesUpdate**](AccessPoliciesApi.md#AccessPoliciesUpdate) | **Put** /pulp/api/v3/access_policies/{pulp_id}/ | Update an access policy



## AccessPoliciesList

> PaginatedAccessPolicyResponseList AccessPoliciesList(ctx).Customized(customized).Limit(limit).Offset(offset).Ordering(ordering).ViewsetName(viewsetName).ViewsetNameContains(viewsetNameContains).ViewsetNameIcontains(viewsetNameIcontains).ViewsetNameIn(viewsetNameIn).ViewsetNameStartswith(viewsetNameStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List access policys



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
    customized := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    viewsetName := "viewsetName_example" // string | Filter results where viewset_name matches value (optional)
    viewsetNameContains := "viewsetNameContains_example" // string | Filter results where viewset_name contains value (optional)
    viewsetNameIcontains := "viewsetNameIcontains_example" // string | Filter results where viewset_name contains value (optional)
    viewsetNameIn := []string{"Inner_example"} // []string | Filter results where viewset_name is in a comma-separated list of values (optional)
    viewsetNameStartswith := "viewsetNameStartswith_example" // string | Filter results where viewset_name starts with value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.AccessPoliciesList(context.Background()).Customized(customized).Limit(limit).Offset(offset).Ordering(ordering).ViewsetName(viewsetName).ViewsetNameContains(viewsetNameContains).ViewsetNameIcontains(viewsetNameIcontains).ViewsetNameIn(viewsetNameIn).ViewsetNameStartswith(viewsetNameStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.AccessPoliciesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesList`: PaginatedAccessPolicyResponseList
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.AccessPoliciesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customized** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **viewsetName** | **string** | Filter results where viewset_name matches value | 
 **viewsetNameContains** | **string** | Filter results where viewset_name contains value | 
 **viewsetNameIcontains** | **string** | Filter results where viewset_name contains value | 
 **viewsetNameIn** | **[]string** | Filter results where viewset_name is in a comma-separated list of values | 
 **viewsetNameStartswith** | **string** | Filter results where viewset_name starts with value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedAccessPolicyResponseList**](PaginatedAccessPolicyResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesPartialUpdate

> AccessPolicyResponse AccessPoliciesPartialUpdate(ctx, pulpId).PatchedAccessPolicy(patchedAccessPolicy).Execute()

Update an access policy



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this access policy.
    patchedAccessPolicy := *openapiclient.NewPatchedAccessPolicy() // PatchedAccessPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.AccessPoliciesPartialUpdate(context.Background(), pulpId).PatchedAccessPolicy(patchedAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.AccessPoliciesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesPartialUpdate`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.AccessPoliciesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAccessPolicy** | [**PatchedAccessPolicy**](PatchedAccessPolicy.md) |  | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesRead

> AccessPolicyResponse AccessPoliciesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an access policy



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this access policy.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.AccessPoliciesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.AccessPoliciesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesRead`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.AccessPoliciesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesReset

> AccessPolicyResponse AccessPoliciesReset(ctx, pulpId).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this access policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.AccessPoliciesReset(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.AccessPoliciesReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesReset`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.AccessPoliciesReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessPoliciesUpdate

> AccessPolicyResponse AccessPoliciesUpdate(ctx, pulpId).AccessPolicy(accessPolicy).Execute()

Update an access policy



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this access policy.
    accessPolicy := *openapiclient.NewAccessPolicy([]map[string]interface{}{map[string]interface{}(123)}) // AccessPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.AccessPoliciesUpdate(context.Background(), pulpId).AccessPolicy(accessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.AccessPoliciesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessPoliciesUpdate`: AccessPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.AccessPoliciesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessPoliciesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessPolicy** | [**AccessPolicy**](AccessPolicy.md) |  | 

### Return type

[**AccessPolicyResponse**](AccessPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

