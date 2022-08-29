# \UsersRolesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersRolesCreate**](UsersRolesApi.md#UsersRolesCreate) | **Post** /pulp/api/v3/users/{user_pk}/roles/ | Create an user role
[**UsersRolesDelete**](UsersRolesApi.md#UsersRolesDelete) | **Delete** /pulp/api/v3/users/{user_pk}/roles/{pulp_id}/ | Delete an user role
[**UsersRolesList**](UsersRolesApi.md#UsersRolesList) | **Get** /pulp/api/v3/users/{user_pk}/roles/ | List user roles
[**UsersRolesRead**](UsersRolesApi.md#UsersRolesRead) | **Get** /pulp/api/v3/users/{user_pk}/roles/{pulp_id}/ | Inspect an user role



## UsersRolesCreate

> UserRoleResponse UsersRolesCreate(ctx, userPk).UserRole(userRole).Execute()

Create an user role



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
    userPk := "userPk_example" // string | 
    userRole := *openapiclient.NewUserRole("Role_example", "ContentObject_example") // UserRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersRolesApi.UsersRolesCreate(context.Background(), userPk).UserRole(userRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesApi.UsersRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesCreate`: UserRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesApi.UsersRolesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRole** | [**UserRole**](UserRole.md) |  | 

### Return type

[**UserRoleResponse**](UserRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRolesDelete

> UsersRolesDelete(ctx, pulpId, userPk).Execute()

Delete an user role



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
    pulpId := TODO // string | A UUID string identifying this user role.
    userPk := "userPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersRolesApi.UsersRolesDelete(context.Background(), pulpId, userPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesApi.UsersRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this user role. | 
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesDeleteRequest struct via the builder pattern


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


## UsersRolesList

> PaginatedUserRoleResponseList UsersRolesList(ctx, userPk).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List user roles



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
    userPk := "userPk_example" // string | 
    contentObject := "contentObject_example" // string | content_object (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    role := "role_example" // string |  (optional)
    roleContains := "roleContains_example" // string |  (optional)
    roleIcontains := "roleIcontains_example" // string |  (optional)
    roleIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    roleStartswith := "roleStartswith_example" // string |  (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersRolesApi.UsersRolesList(context.Background(), userPk).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesApi.UsersRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesList`: PaginatedUserRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesApi.UsersRolesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentObject** | **string** | content_object | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **role** | **string** |  | 
 **roleContains** | **string** |  | 
 **roleIcontains** | **string** |  | 
 **roleIn** | **[]string** | Multiple values may be separated by commas. | 
 **roleStartswith** | **string** |  | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUserRoleResponseList**](PaginatedUserRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRolesRead

> UserRoleResponse UsersRolesRead(ctx, pulpId, userPk).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an user role



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
    pulpId := TODO // string | A UUID string identifying this user role.
    userPk := "userPk_example" // string | 
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersRolesApi.UsersRolesRead(context.Background(), pulpId, userPk).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersRolesApi.UsersRolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRolesRead`: UserRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersRolesApi.UsersRolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this user role. | 
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**UserRoleResponse**](UserRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

