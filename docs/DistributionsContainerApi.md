# \DistributionsContainerApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsContainerContainerAddRole**](DistributionsContainerApi.md#DistributionsContainerContainerAddRole) | **Post** /pulp/api/v3/distributions/container/container/{pulp_id}/add_role/ | 
[**DistributionsContainerContainerCreate**](DistributionsContainerApi.md#DistributionsContainerContainerCreate) | **Post** /pulp/api/v3/distributions/container/container/ | Create a container distribution
[**DistributionsContainerContainerDelete**](DistributionsContainerApi.md#DistributionsContainerContainerDelete) | **Delete** /pulp/api/v3/distributions/container/container/{pulp_id}/ | Delete a container distribution
[**DistributionsContainerContainerList**](DistributionsContainerApi.md#DistributionsContainerContainerList) | **Get** /pulp/api/v3/distributions/container/container/ | List container distributions
[**DistributionsContainerContainerListRoles**](DistributionsContainerApi.md#DistributionsContainerContainerListRoles) | **Get** /pulp/api/v3/distributions/container/container/{pulp_id}/list_roles/ | 
[**DistributionsContainerContainerMyPermissions**](DistributionsContainerApi.md#DistributionsContainerContainerMyPermissions) | **Get** /pulp/api/v3/distributions/container/container/{pulp_id}/my_permissions/ | 
[**DistributionsContainerContainerPartialUpdate**](DistributionsContainerApi.md#DistributionsContainerContainerPartialUpdate) | **Patch** /pulp/api/v3/distributions/container/container/{pulp_id}/ | Update a container distribution
[**DistributionsContainerContainerRead**](DistributionsContainerApi.md#DistributionsContainerContainerRead) | **Get** /pulp/api/v3/distributions/container/container/{pulp_id}/ | Inspect a container distribution
[**DistributionsContainerContainerRemoveRole**](DistributionsContainerApi.md#DistributionsContainerContainerRemoveRole) | **Post** /pulp/api/v3/distributions/container/container/{pulp_id}/remove_role/ | 
[**DistributionsContainerContainerUpdate**](DistributionsContainerApi.md#DistributionsContainerContainerUpdate) | **Put** /pulp/api/v3/distributions/container/container/{pulp_id}/ | Update a container distribution



## DistributionsContainerContainerAddRole

> NestedRoleResponse DistributionsContainerContainerAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerCreate

> AsyncOperationResponse DistributionsContainerContainerCreate(ctx).ContainerContainerDistribution(containerContainerDistribution).Execute()

Create a container distribution



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
    containerContainerDistribution := *openapiclient.NewContainerContainerDistribution("Name_example", "BasePath_example") // ContainerContainerDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerCreate(context.Background()).ContainerContainerDistribution(containerContainerDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerContainerDistribution** | [**ContainerContainerDistribution**](ContainerContainerDistribution.md) |  | 

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


## DistributionsContainerContainerDelete

> AsyncOperationResponse DistributionsContainerContainerDelete(ctx, pulpId).Execute()

Delete a container distribution



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerDeleteRequest struct via the builder pattern


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


## DistributionsContainerContainerList

> PaginatedcontainerContainerDistributionResponseList DistributionsContainerContainerList(ctx).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).NamespaceName(namespaceName).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()

List container distributions



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
    basePath := "basePath_example" // string |  (optional)
    basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
    basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
    basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    namespaceName := "namespaceName_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerList(context.Background()).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).NamespaceName(namespaceName).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerList`: PaginatedcontainerContainerDistributionResponseList
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basePath** | **string** |  | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **namespaceName** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerDistributionResponseList**](PaginatedcontainerContainerDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerListRoles

> ObjectRolesResponse DistributionsContainerContainerListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ObjectRolesResponse**](ObjectRolesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerMyPermissions

> MyPermissionsResponse DistributionsContainerContainerMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**MyPermissionsResponse**](MyPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerPartialUpdate

> AsyncOperationResponse DistributionsContainerContainerPartialUpdate(ctx, pulpId).PatchedcontainerContainerDistribution(patchedcontainerContainerDistribution).Execute()

Update a container distribution



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    patchedcontainerContainerDistribution := *openapiclient.NewPatchedcontainerContainerDistribution() // PatchedcontainerContainerDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerPartialUpdate(context.Background(), pulpId).PatchedcontainerContainerDistribution(patchedcontainerContainerDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerDistribution** | [**PatchedcontainerContainerDistribution**](PatchedcontainerContainerDistribution.md) |  | 

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


## DistributionsContainerContainerRead

> ContainerContainerDistributionResponse DistributionsContainerContainerRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container distribution



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerRead`: ContainerContainerDistributionResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerDistributionResponse**](ContainerContainerDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerRemoveRole

> NestedRoleResponse DistributionsContainerContainerRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerUpdate

> AsyncOperationResponse DistributionsContainerContainerUpdate(ctx, pulpId).ContainerContainerDistribution(containerContainerDistribution).Execute()

Update a container distribution



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container distribution.
    containerContainerDistribution := *openapiclient.NewContainerContainerDistribution("Name_example", "BasePath_example") // ContainerContainerDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerUpdate(context.Background(), pulpId).ContainerContainerDistribution(containerContainerDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerApi.DistributionsContainerContainerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsContainerContainerUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerApi.DistributionsContainerContainerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container distribution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerDistribution** | [**ContainerContainerDistribution**](ContainerContainerDistribution.md) |  | 

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

