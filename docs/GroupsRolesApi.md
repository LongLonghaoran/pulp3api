# \GroupsRolesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsRolesCreate**](GroupsRolesApi.md#GroupsRolesCreate) | **Post** /pulp/api/v3/groups/{group_pk}/roles/ | Create a group role
[**GroupsRolesDelete**](GroupsRolesApi.md#GroupsRolesDelete) | **Delete** /pulp/api/v3/groups/{group_pk}/roles/{pulp_id}/ | Delete a group role
[**GroupsRolesList**](GroupsRolesApi.md#GroupsRolesList) | **Get** /pulp/api/v3/groups/{group_pk}/roles/ | List group roles
[**GroupsRolesRead**](GroupsRolesApi.md#GroupsRolesRead) | **Get** /pulp/api/v3/groups/{group_pk}/roles/{pulp_id}/ | Inspect a group role



## GroupsRolesCreate

> GroupRoleResponse GroupsRolesCreate(ctx, groupPk).GroupRole(groupRole).Execute()

Create a group role



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
    groupPk := "groupPk_example" // string | 
    groupRole := *openapiclient.NewGroupRole("Role_example", "ContentObject_example") // GroupRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsRolesApi.GroupsRolesCreate(context.Background(), groupPk).GroupRole(groupRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesCreate`: GroupRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRole** | [**GroupRole**](GroupRole.md) |  | 

### Return type

[**GroupRoleResponse**](GroupRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsRolesDelete

> GroupsRolesDelete(ctx, groupPk, pulpId).Execute()

Delete a group role



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
    groupPk := "groupPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this group role.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsRolesApi.GroupsRolesDelete(context.Background(), groupPk, pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this group role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesDeleteRequest struct via the builder pattern


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


## GroupsRolesList

> PaginatedGroupRoleResponseList GroupsRolesList(ctx, groupPk).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List group roles



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
    groupPk := "groupPk_example" // string | 
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
    resp, r, err := api_client.GroupsRolesApi.GroupsRolesList(context.Background(), groupPk).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesList`: PaginatedGroupRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesListRequest struct via the builder pattern


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

[**PaginatedGroupRoleResponseList**](PaginatedGroupRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsRolesRead

> GroupRoleResponse GroupsRolesRead(ctx, groupPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a group role



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
    groupPk := "groupPk_example" // string | 
    pulpId := TODO // string | A UUID string identifying this group role.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsRolesApi.GroupsRolesRead(context.Background(), groupPk, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesRead`: GroupRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 
**pulpId** | [**string**](.md) | A UUID string identifying this group role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**GroupRoleResponse**](GroupRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

