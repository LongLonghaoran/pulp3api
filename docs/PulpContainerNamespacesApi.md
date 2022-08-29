# \PulpContainerNamespacesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpContainerNamespacesAddRole**](PulpContainerNamespacesApi.md#PulpContainerNamespacesAddRole) | **Post** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/add_role/ | 
[**PulpContainerNamespacesCreate**](PulpContainerNamespacesApi.md#PulpContainerNamespacesCreate) | **Post** /pulp/api/v3/pulp_container/namespaces/ | Create a container namespace
[**PulpContainerNamespacesDelete**](PulpContainerNamespacesApi.md#PulpContainerNamespacesDelete) | **Delete** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/ | Delete a container namespace
[**PulpContainerNamespacesList**](PulpContainerNamespacesApi.md#PulpContainerNamespacesList) | **Get** /pulp/api/v3/pulp_container/namespaces/ | List container namespaces
[**PulpContainerNamespacesListRoles**](PulpContainerNamespacesApi.md#PulpContainerNamespacesListRoles) | **Get** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/list_roles/ | 
[**PulpContainerNamespacesMyPermissions**](PulpContainerNamespacesApi.md#PulpContainerNamespacesMyPermissions) | **Get** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/my_permissions/ | 
[**PulpContainerNamespacesRead**](PulpContainerNamespacesApi.md#PulpContainerNamespacesRead) | **Get** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/ | Inspect a container namespace
[**PulpContainerNamespacesRemoveRole**](PulpContainerNamespacesApi.md#PulpContainerNamespacesRemoveRole) | **Post** /pulp/api/v3/pulp_container/namespaces/{pulp_id}/remove_role/ | 



## PulpContainerNamespacesAddRole

> NestedRoleResponse PulpContainerNamespacesAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesAddRoleRequest struct via the builder pattern


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


## PulpContainerNamespacesCreate

> ContainerContainerNamespaceResponse PulpContainerNamespacesCreate(ctx).ContainerContainerNamespace(containerContainerNamespace).Execute()

Create a container namespace



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
    containerContainerNamespace := *openapiclient.NewContainerContainerNamespace("Name_example") // ContainerContainerNamespace | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesCreate(context.Background()).ContainerContainerNamespace(containerContainerNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesCreate`: ContainerContainerNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerContainerNamespace** | [**ContainerContainerNamespace**](ContainerContainerNamespace.md) |  | 

### Return type

[**ContainerContainerNamespaceResponse**](ContainerContainerNamespaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpContainerNamespacesDelete

> AsyncOperationResponse PulpContainerNamespacesDelete(ctx, pulpId).Execute()

Delete a container namespace



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesDeleteRequest struct via the builder pattern


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


## PulpContainerNamespacesList

> PaginatedcontainerContainerNamespaceResponseList PulpContainerNamespacesList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List container namespaces



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
    name := "name_example" // string | Filter results where name matches value (optional)
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
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesList`: PaginatedcontainerContainerNamespaceResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerNamespaceResponseList**](PaginatedcontainerContainerNamespaceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpContainerNamespacesListRoles

> ObjectRolesResponse PulpContainerNamespacesListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesListRolesRequest struct via the builder pattern


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


## PulpContainerNamespacesMyPermissions

> MyPermissionsResponse PulpContainerNamespacesMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesMyPermissionsRequest struct via the builder pattern


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


## PulpContainerNamespacesRead

> ContainerContainerNamespaceResponse PulpContainerNamespacesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container namespace



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesRead`: ContainerContainerNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerNamespaceResponse**](ContainerContainerNamespaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpContainerNamespacesRemoveRole

> NestedRoleResponse PulpContainerNamespacesRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container namespace.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpContainerNamespacesApi.PulpContainerNamespacesRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpContainerNamespacesRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpContainerNamespacesApi.PulpContainerNamespacesRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpContainerNamespacesRemoveRoleRequest struct via the builder pattern


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

