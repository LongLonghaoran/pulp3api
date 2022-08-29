# \RolesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RolesCreate**](RolesApi.md#RolesCreate) | **Post** /pulp/api/v3/roles/ | Create a role
[**RolesDelete**](RolesApi.md#RolesDelete) | **Delete** /pulp/api/v3/roles/{pulp_id}/ | Delete a role
[**RolesList**](RolesApi.md#RolesList) | **Get** /pulp/api/v3/roles/ | List roles
[**RolesPartialUpdate**](RolesApi.md#RolesPartialUpdate) | **Patch** /pulp/api/v3/roles/{pulp_id}/ | Update a role
[**RolesRead**](RolesApi.md#RolesRead) | **Get** /pulp/api/v3/roles/{pulp_id}/ | Inspect a role
[**RolesUpdate**](RolesApi.md#RolesUpdate) | **Put** /pulp/api/v3/roles/{pulp_id}/ | Update a role



## RolesCreate

> RoleResponse RolesCreate(ctx).Role(role).Execute()

Create a role



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
    role := *openapiclient.NewRole("Name_example", []string{"Permissions_example"}) // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RolesCreate(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesCreate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.RolesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesDelete

> RolesDelete(ctx, pulpId).Execute()

Delete a role



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RolesDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesDeleteRequest struct via the builder pattern


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


## RolesList

> PaginatedRoleResponseList RolesList(ctx).ContainsPermission(containsPermission).Description(description).DescriptionContains(descriptionContains).DescriptionIcontains(descriptionIcontains).DescriptionIexact(descriptionIexact).ForObjectType(forObjectType).Limit(limit).Locked(locked).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
    containsPermission := []string{"Inner_example"} // []string | Filter roles that have any of the permissions in the list. (optional)
    description := "description_example" // string | Filter results where description matches value (optional)
    descriptionContains := "descriptionContains_example" // string | Filter results where description contains value (optional)
    descriptionIcontains := "descriptionIcontains_example" // string | Filter results where description contains value (optional)
    descriptionIexact := "descriptionIexact_example" // string | Filter results where description matches value (optional)
    forObjectType := "forObjectType_example" // string | Filter roles that only have permissions for the specified object HREF. (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    locked := true // bool | Filter results where locked matches value (optional)
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
    resp, r, err := apiClient.RolesApi.RolesList(context.Background()).ContainsPermission(containsPermission).Description(description).DescriptionContains(descriptionContains).DescriptionIcontains(descriptionIcontains).DescriptionIexact(descriptionIexact).ForObjectType(forObjectType).Limit(limit).Locked(locked).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesList`: PaginatedRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.RolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containsPermission** | **[]string** | Filter roles that have any of the permissions in the list. | 
 **description** | **string** | Filter results where description matches value | 
 **descriptionContains** | **string** | Filter results where description contains value | 
 **descriptionIcontains** | **string** | Filter results where description contains value | 
 **descriptionIexact** | **string** | Filter results where description matches value | 
 **forObjectType** | **string** | Filter roles that only have permissions for the specified object HREF. | 
 **limit** | **int32** | Number of results to return per page. | 
 **locked** | **bool** | Filter results where locked matches value | 
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

[**PaginatedRoleResponseList**](PaginatedRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesPartialUpdate

> RoleResponse RolesPartialUpdate(ctx, pulpId).PatchedRole(patchedRole).Execute()

Update a role



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
    patchedRole := *openapiclient.NewPatchedRole() // PatchedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RolesPartialUpdate(context.Background(), pulpId).PatchedRole(patchedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesPartialUpdate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.RolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRole** | [**PatchedRole**](PatchedRole.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRead

> RoleResponse RolesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a role



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RolesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesRead`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.RolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesUpdate

> RoleResponse RolesUpdate(ctx, pulpId).Role(role).Execute()

Update a role



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
    role := *openapiclient.NewRole("Name_example", []string{"Permissions_example"}) // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RolesUpdate(context.Background(), pulpId).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesUpdate`: RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.RolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

