# \TaskSchedulesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskSchedulesAddRole**](TaskSchedulesApi.md#TaskSchedulesAddRole) | **Post** /pulp/api/v3/task-schedules/{pulp_id}/add_role/ | 
[**TaskSchedulesList**](TaskSchedulesApi.md#TaskSchedulesList) | **Get** /pulp/api/v3/task-schedules/ | List task schedules
[**TaskSchedulesListRoles**](TaskSchedulesApi.md#TaskSchedulesListRoles) | **Get** /pulp/api/v3/task-schedules/{pulp_id}/list_roles/ | 
[**TaskSchedulesMyPermissions**](TaskSchedulesApi.md#TaskSchedulesMyPermissions) | **Get** /pulp/api/v3/task-schedules/{pulp_id}/my_permissions/ | 
[**TaskSchedulesRead**](TaskSchedulesApi.md#TaskSchedulesRead) | **Get** /pulp/api/v3/task-schedules/{pulp_id}/ | Inspect a task schedule
[**TaskSchedulesRemoveRole**](TaskSchedulesApi.md#TaskSchedulesRemoveRole) | **Post** /pulp/api/v3/task-schedules/{pulp_id}/remove_role/ | 



## TaskSchedulesAddRole

> NestedRoleResponse TaskSchedulesAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task schedule.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this task schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesAddRoleRequest struct via the builder pattern


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


## TaskSchedulesList

> PaginatedTaskScheduleResponseList TaskSchedulesList(ctx).Limit(limit).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).TaskName(taskName).TaskNameContains(taskNameContains).Fields(fields).ExcludeFields(excludeFields).Execute()

List task schedules



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    taskName := "taskName_example" // string |  (optional)
    taskNameContains := "taskNameContains_example" // string | Filter results where task_name contains value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).TaskName(taskName).TaskNameContains(taskNameContains).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesList`: PaginatedTaskScheduleResponseList
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **nameContains** | **string** | Filter results where name contains value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **taskName** | **string** |  | 
 **taskNameContains** | **string** | Filter results where task_name contains value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskScheduleResponseList**](PaginatedTaskScheduleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskSchedulesListRoles

> ObjectRolesResponse TaskSchedulesListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task schedule.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this task schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesListRolesRequest struct via the builder pattern


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


## TaskSchedulesMyPermissions

> MyPermissionsResponse TaskSchedulesMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task schedule.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this task schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesMyPermissionsRequest struct via the builder pattern


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


## TaskSchedulesRead

> TaskScheduleResponse TaskSchedulesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task schedule



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task schedule.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesRead`: TaskScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this task schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**TaskScheduleResponse**](TaskScheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskSchedulesRemoveRole

> NestedRoleResponse TaskSchedulesRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task schedule.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSchedulesApi.TaskSchedulesRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulesApi.TaskSchedulesRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSchedulesRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSchedulesApi.TaskSchedulesRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this task schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskSchedulesRemoveRoleRequest struct via the builder pattern


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

