# \ContentguardsRbacApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreRbacAddRole**](ContentguardsRbacApi.md#ContentguardsCoreRbacAddRole) | **Post** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/add_role/ | 
[**ContentguardsCoreRbacCreate**](ContentguardsRbacApi.md#ContentguardsCoreRbacCreate) | **Post** /pulp/api/v3/contentguards/core/rbac/ | Create a rbac content guard
[**ContentguardsCoreRbacDelete**](ContentguardsRbacApi.md#ContentguardsCoreRbacDelete) | **Delete** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/ | Delete a rbac content guard
[**ContentguardsCoreRbacList**](ContentguardsRbacApi.md#ContentguardsCoreRbacList) | **Get** /pulp/api/v3/contentguards/core/rbac/ | List rbac content guards
[**ContentguardsCoreRbacListRoles**](ContentguardsRbacApi.md#ContentguardsCoreRbacListRoles) | **Get** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/list_roles/ | 
[**ContentguardsCoreRbacMyPermissions**](ContentguardsRbacApi.md#ContentguardsCoreRbacMyPermissions) | **Get** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/my_permissions/ | 
[**ContentguardsCoreRbacPartialUpdate**](ContentguardsRbacApi.md#ContentguardsCoreRbacPartialUpdate) | **Patch** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/ | Update a rbac content guard
[**ContentguardsCoreRbacRead**](ContentguardsRbacApi.md#ContentguardsCoreRbacRead) | **Get** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/ | Inspect a rbac content guard
[**ContentguardsCoreRbacRemoveRole**](ContentguardsRbacApi.md#ContentguardsCoreRbacRemoveRole) | **Post** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/remove_role/ | 
[**ContentguardsCoreRbacUpdate**](ContentguardsRbacApi.md#ContentguardsCoreRbacUpdate) | **Put** /pulp/api/v3/contentguards/core/rbac/{pulp_id}/ | Update a rbac content guard



## ContentguardsCoreRbacAddRole

> NestedRoleResponse ContentguardsCoreRbacAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacAddRoleRequest struct via the builder pattern


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


## ContentguardsCoreRbacCreate

> RBACContentGuardResponse ContentguardsCoreRbacCreate(ctx).RBACContentGuard(rBACContentGuard).Execute()

Create a rbac content guard



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
    rBACContentGuard := *openapiclient.NewRBACContentGuard("Name_example") // RBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacCreate(context.Background()).RBACContentGuard(rBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacCreate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rBACContentGuard** | [**RBACContentGuard**](RBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacDelete

> ContentguardsCoreRbacDelete(ctx, pulpId).Execute()

Delete a rbac content guard



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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacDeleteRequest struct via the builder pattern


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


## ContentguardsCoreRbacList

> PaginatedRBACContentGuardResponseList ContentguardsCoreRbacList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List rbac content guards



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
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacList`: PaginatedRBACContentGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacListRequest struct via the builder pattern


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
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRBACContentGuardResponseList**](PaginatedRBACContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacListRoles

> ObjectRolesResponse ContentguardsCoreRbacListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacListRolesRequest struct via the builder pattern


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


## ContentguardsCoreRbacMyPermissions

> MyPermissionsResponse ContentguardsCoreRbacMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacMyPermissionsRequest struct via the builder pattern


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


## ContentguardsCoreRbacPartialUpdate

> RBACContentGuardResponse ContentguardsCoreRbacPartialUpdate(ctx, pulpId).PatchedRBACContentGuard(patchedRBACContentGuard).Execute()

Update a rbac content guard



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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    patchedRBACContentGuard := *openapiclient.NewPatchedRBACContentGuard() // PatchedRBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacPartialUpdate(context.Background(), pulpId).PatchedRBACContentGuard(patchedRBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacPartialUpdate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRBACContentGuard** | [**PatchedRBACContentGuard**](PatchedRBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacRead

> RBACContentGuardResponse ContentguardsCoreRbacRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rbac content guard



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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacRead`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreRbacRemoveRole

> NestedRoleResponse ContentguardsCoreRbacRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacRemoveRoleRequest struct via the builder pattern


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


## ContentguardsCoreRbacUpdate

> RBACContentGuardResponse ContentguardsCoreRbacUpdate(ctx, pulpId).RBACContentGuard(rBACContentGuard).Execute()

Update a rbac content guard



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
    pulpId := TODO // string | A UUID string identifying this rbac content guard.
    rBACContentGuard := *openapiclient.NewRBACContentGuard("Name_example") // RBACContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsRbacApi.ContentguardsCoreRbacUpdate(context.Background(), pulpId).RBACContentGuard(rBACContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRbacApi.ContentguardsCoreRbacUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreRbacUpdate`: RBACContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRbacApi.ContentguardsCoreRbacUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this rbac content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreRbacUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rBACContentGuard** | [**RBACContentGuard**](RBACContentGuard.md) |  | 

### Return type

[**RBACContentGuardResponse**](RBACContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

