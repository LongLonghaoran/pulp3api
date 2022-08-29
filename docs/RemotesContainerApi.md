# \RemotesContainerApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesContainerContainerAddRole**](RemotesContainerApi.md#RemotesContainerContainerAddRole) | **Post** /pulp/api/v3/remotes/container/container/{pulp_id}/add_role/ | 
[**RemotesContainerContainerCreate**](RemotesContainerApi.md#RemotesContainerContainerCreate) | **Post** /pulp/api/v3/remotes/container/container/ | Create a container remote
[**RemotesContainerContainerDelete**](RemotesContainerApi.md#RemotesContainerContainerDelete) | **Delete** /pulp/api/v3/remotes/container/container/{pulp_id}/ | Delete a container remote
[**RemotesContainerContainerList**](RemotesContainerApi.md#RemotesContainerContainerList) | **Get** /pulp/api/v3/remotes/container/container/ | List container remotes
[**RemotesContainerContainerListRoles**](RemotesContainerApi.md#RemotesContainerContainerListRoles) | **Get** /pulp/api/v3/remotes/container/container/{pulp_id}/list_roles/ | 
[**RemotesContainerContainerMyPermissions**](RemotesContainerApi.md#RemotesContainerContainerMyPermissions) | **Get** /pulp/api/v3/remotes/container/container/{pulp_id}/my_permissions/ | 
[**RemotesContainerContainerPartialUpdate**](RemotesContainerApi.md#RemotesContainerContainerPartialUpdate) | **Patch** /pulp/api/v3/remotes/container/container/{pulp_id}/ | Update a container remote
[**RemotesContainerContainerRead**](RemotesContainerApi.md#RemotesContainerContainerRead) | **Get** /pulp/api/v3/remotes/container/container/{pulp_id}/ | Inspect a container remote
[**RemotesContainerContainerRemoveRole**](RemotesContainerApi.md#RemotesContainerContainerRemoveRole) | **Post** /pulp/api/v3/remotes/container/container/{pulp_id}/remove_role/ | 
[**RemotesContainerContainerUpdate**](RemotesContainerApi.md#RemotesContainerContainerUpdate) | **Put** /pulp/api/v3/remotes/container/container/{pulp_id}/ | Update a container remote



## RemotesContainerContainerAddRole

> NestedRoleResponse RemotesContainerContainerAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerAddRoleRequest struct via the builder pattern


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


## RemotesContainerContainerCreate

> ContainerContainerRemoteResponse RemotesContainerContainerCreate(ctx).ContainerContainerRemote(containerContainerRemote).Execute()

Create a container remote



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
    containerContainerRemote := *openapiclient.NewContainerContainerRemote("Name_example", "Url_example", "UpstreamName_example") // ContainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerCreate(context.Background()).ContainerContainerRemote(containerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerCreate`: ContainerContainerRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerContainerRemote** | [**ContainerContainerRemote**](ContainerContainerRemote.md) |  | 

### Return type

[**ContainerContainerRemoteResponse**](ContainerContainerRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerDelete

> AsyncOperationResponse RemotesContainerContainerDelete(ctx, pulpId).Execute()

Delete a container remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerDeleteRequest struct via the builder pattern


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


## RemotesContainerContainerList

> PaginatedcontainerContainerRemoteResponseList RemotesContainerContainerList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List container remotes



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    pulpLastUpdated := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
    pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
    pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
    pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
    pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerList`: PaginatedcontainerContainerRemoteResponseList
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerListRequest struct via the builder pattern


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
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | ISO 8601 formatted dates are supported | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerRemoteResponseList**](PaginatedcontainerContainerRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerListRoles

> ObjectRolesResponse RemotesContainerContainerListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerListRolesRequest struct via the builder pattern


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


## RemotesContainerContainerMyPermissions

> MyPermissionsResponse RemotesContainerContainerMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerMyPermissionsRequest struct via the builder pattern


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


## RemotesContainerContainerPartialUpdate

> AsyncOperationResponse RemotesContainerContainerPartialUpdate(ctx, pulpId).PatchedcontainerContainerRemote(patchedcontainerContainerRemote).Execute()

Update a container remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    patchedcontainerContainerRemote := *openapiclient.NewPatchedcontainerContainerRemote() // PatchedcontainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerPartialUpdate(context.Background(), pulpId).PatchedcontainerContainerRemote(patchedcontainerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerRemote** | [**PatchedcontainerContainerRemote**](PatchedcontainerContainerRemote.md) |  | 

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


## RemotesContainerContainerRead

> ContainerContainerRemoteResponse RemotesContainerContainerRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerRead`: ContainerContainerRemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerRemoteResponse**](ContainerContainerRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerContainerRemoveRole

> NestedRoleResponse RemotesContainerContainerRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerRemoveRoleRequest struct via the builder pattern


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


## RemotesContainerContainerUpdate

> AsyncOperationResponse RemotesContainerContainerUpdate(ctx, pulpId).ContainerContainerRemote(containerContainerRemote).Execute()

Update a container remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container remote.
    containerContainerRemote := *openapiclient.NewContainerContainerRemote("Name_example", "Url_example", "UpstreamName_example") // ContainerContainerRemote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotesContainerApi.RemotesContainerContainerUpdate(context.Background(), pulpId).ContainerContainerRemote(containerContainerRemote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotesContainerApi.RemotesContainerContainerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemotesContainerContainerUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RemotesContainerApi.RemotesContainerContainerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container remote. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerContainerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerRemote** | [**ContainerContainerRemote**](ContainerContainerRemote.md) |  | 

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

