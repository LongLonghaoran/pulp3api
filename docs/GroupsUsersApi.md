# \GroupsUsersApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsUsersCreate**](GroupsUsersApi.md#GroupsUsersCreate) | **Post** /pulp/api/v3/groups/{group_pk}/users/ | Create an user
[**GroupsUsersDelete**](GroupsUsersApi.md#GroupsUsersDelete) | **Delete** /pulp/api/v3/groups/{group_pk}/users/{id}/ | Delete an user
[**GroupsUsersList**](GroupsUsersApi.md#GroupsUsersList) | **Get** /pulp/api/v3/groups/{group_pk}/users/ | List users



## GroupsUsersCreate

> GroupUserResponse GroupsUsersCreate(ctx, groupPk).GroupUser(groupUser).Execute()

Create an user



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
    groupUser := *openapiclient.NewGroupUser("Username_example") // GroupUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsUsersApi.GroupsUsersCreate(context.Background(), groupPk).GroupUser(groupUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersApi.GroupsUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUsersCreate`: GroupUserResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsUsersApi.GroupsUsersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUser** | [**GroupUser**](GroupUser.md) |  | 

### Return type

[**GroupUserResponse**](GroupUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUsersDelete

> GroupsUsersDelete(ctx, groupPk, id).Execute()

Delete an user



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
    id := int32(56) // int32 | A unique integer value identifying this user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsUsersApi.GroupsUsersDelete(context.Background(), groupPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersApi.GroupsUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersDeleteRequest struct via the builder pattern


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


## GroupsUsersList

> PaginatedGroupUserResponseList GroupsUsersList(ctx, groupPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List users



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsUsersApi.GroupsUsersList(context.Background(), groupPk).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsUsersApi.GroupsUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUsersList`: PaginatedGroupUserResponseList
    fmt.Fprintf(os.Stdout, "Response from `GroupsUsersApi.GroupsUsersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGroupUserResponseList**](PaginatedGroupUserResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

