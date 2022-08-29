# \TasksApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksAddRole**](TasksApi.md#TasksAddRole) | **Post** /pulp/api/v3/tasks/{pulp_id}/add_role/ | 
[**TasksCancel**](TasksApi.md#TasksCancel) | **Patch** /pulp/api/v3/tasks/{pulp_id}/ | Cancel a task
[**TasksDelete**](TasksApi.md#TasksDelete) | **Delete** /pulp/api/v3/tasks/{pulp_id}/ | Delete a task
[**TasksList**](TasksApi.md#TasksList) | **Get** /pulp/api/v3/tasks/ | List tasks
[**TasksListRoles**](TasksApi.md#TasksListRoles) | **Get** /pulp/api/v3/tasks/{pulp_id}/list_roles/ | 
[**TasksMyPermissions**](TasksApi.md#TasksMyPermissions) | **Get** /pulp/api/v3/tasks/{pulp_id}/my_permissions/ | 
[**TasksPurge**](TasksApi.md#TasksPurge) | **Post** /pulp/api/v3/tasks/purge/ | Purge Completed Tasks
[**TasksRead**](TasksApi.md#TasksRead) | **Get** /pulp/api/v3/tasks/{pulp_id}/ | Inspect a task
[**TasksRemoveRole**](TasksApi.md#TasksRemoveRole) | **Post** /pulp/api/v3/tasks/{pulp_id}/remove_role/ | 



## TasksAddRole

> NestedRoleResponse TasksAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this task.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksAddRoleRequest struct via the builder pattern


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


## TasksCancel

> TaskResponse TasksCancel(ctx, pulpId).PatchedTaskCancel(patchedTaskCancel).Execute()

Cancel a task



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
    pulpId := TODO // string | A UUID string identifying this task.
    patchedTaskCancel := *openapiclient.NewPatchedTaskCancel() // PatchedTaskCancel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksCancel(context.Background(), pulpId).PatchedTaskCancel(patchedTaskCancel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksCancel`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTaskCancel** | [**PatchedTaskCancel**](PatchedTaskCancel.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksDelete

> TasksDelete(ctx, pulpId).Execute()

Delete a task



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
    pulpId := TODO // string | A UUID string identifying this task.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksDeleteRequest struct via the builder pattern


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


## TasksList

> PaginatedTaskResponseList TasksList(ctx).ChildTasks(childTasks).CreatedResources(createdResources).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).ParentTask(parentTask).ReservedResourcesRecord(reservedResourcesRecord).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).TaskGroup(taskGroup).Worker(worker).WorkerIn(workerIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List tasks



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
    childTasks := TODO // string | Foreign Key referenced by HREF (optional)
    createdResources := TODO // string |  (optional)
    finishedAt := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    finishedAtGt := time.Now() // time.Time | Filter results where finished_at is greater than value (optional)
    finishedAtGte := time.Now() // time.Time | Filter results where finished_at is greater than or equal to value (optional)
    finishedAtLt := time.Now() // time.Time | Filter results where finished_at is less than value (optional)
    finishedAtLte := time.Now() // time.Time | Filter results where finished_at is less than or equal to value (optional)
    finishedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where finished_at is between two comma separated values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    loggingCid := "loggingCid_example" // string |  (optional)
    loggingCidContains := "loggingCidContains_example" // string | Filter results where logging_cid contains value (optional)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    parentTask := TODO // string | Foreign Key referenced by HREF (optional)
    reservedResourcesRecord := []string{"Inner_example"} // []string |  (optional)
    startedAt := time.Now() // time.Time | ISO 8601 formatted dates are supported (optional)
    startedAtGt := time.Now() // time.Time | Filter results where started_at is greater than value (optional)
    startedAtGte := time.Now() // time.Time | Filter results where started_at is greater than or equal to value (optional)
    startedAtLt := time.Now() // time.Time | Filter results where started_at is less than value (optional)
    startedAtLte := time.Now() // time.Time | Filter results where started_at is less than or equal to value (optional)
    startedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where started_at is between two comma separated values (optional)
    state := "state_example" // string |  (optional)
    stateIn := []string{"Inner_example"} // []string | Filter results where state is in a comma-separated list of values (optional)
    taskGroup := TODO // string | Foreign Key referenced by HREF (optional)
    worker := TODO // string | Foreign Key referenced by HREF (optional)
    workerIn := []string{"Inner_example"} // []string | Filter results where worker is in a comma-separated list of values (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksList(context.Background()).ChildTasks(childTasks).CreatedResources(createdResources).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).Offset(offset).Ordering(ordering).ParentTask(parentTask).ReservedResourcesRecord(reservedResourcesRecord).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).TaskGroup(taskGroup).Worker(worker).WorkerIn(workerIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksList`: PaginatedTaskResponseList
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childTasks** | [**string**](string.md) | Foreign Key referenced by HREF | 
 **createdResources** | [**string**](string.md) |  | 
 **finishedAt** | **time.Time** | ISO 8601 formatted dates are supported | 
 **finishedAtGt** | **time.Time** | Filter results where finished_at is greater than value | 
 **finishedAtGte** | **time.Time** | Filter results where finished_at is greater than or equal to value | 
 **finishedAtLt** | **time.Time** | Filter results where finished_at is less than value | 
 **finishedAtLte** | **time.Time** | Filter results where finished_at is less than or equal to value | 
 **finishedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where finished_at is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **loggingCid** | **string** |  | 
 **loggingCidContains** | **string** | Filter results where logging_cid contains value | 
 **name** | **string** |  | 
 **nameContains** | **string** | Filter results where name contains value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **parentTask** | [**string**](string.md) | Foreign Key referenced by HREF | 
 **reservedResourcesRecord** | **[]string** |  | 
 **startedAt** | **time.Time** | ISO 8601 formatted dates are supported | 
 **startedAtGt** | **time.Time** | Filter results where started_at is greater than value | 
 **startedAtGte** | **time.Time** | Filter results where started_at is greater than or equal to value | 
 **startedAtLt** | **time.Time** | Filter results where started_at is less than value | 
 **startedAtLte** | **time.Time** | Filter results where started_at is less than or equal to value | 
 **startedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where started_at is between two comma separated values | 
 **state** | **string** |  | 
 **stateIn** | **[]string** | Filter results where state is in a comma-separated list of values | 
 **taskGroup** | [**string**](string.md) | Foreign Key referenced by HREF | 
 **worker** | [**string**](string.md) | Foreign Key referenced by HREF | 
 **workerIn** | **[]string** | Filter results where worker is in a comma-separated list of values | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskResponseList**](PaginatedTaskResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksListRoles

> ObjectRolesResponse TasksListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this task.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRolesRequest struct via the builder pattern


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


## TasksMyPermissions

> MyPermissionsResponse TasksMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this task.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksMyPermissionsRequest struct via the builder pattern


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


## TasksPurge

> AsyncOperationResponse TasksPurge(ctx).Purge(purge).Execute()

Purge Completed Tasks



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
    purge := *openapiclient.NewPurge() // Purge | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksPurge(context.Background()).Purge(purge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksPurge`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksPurge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purge** | [**Purge**](Purge.md) |  | 

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


## TasksRead

> TaskResponse TasksRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task



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
    pulpId := TODO // string | A UUID string identifying this task.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRead`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRemoveRole

> NestedRoleResponse TasksRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this task.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.TasksRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRemoveRoleRequest struct via the builder pattern


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

