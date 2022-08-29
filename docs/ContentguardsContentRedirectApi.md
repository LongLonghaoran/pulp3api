# \ContentguardsContentRedirectApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreContentRedirectAddRole**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectAddRole) | **Post** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/add_role/ | 
[**ContentguardsCoreContentRedirectCreate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectCreate) | **Post** /pulp/api/v3/contentguards/core/content_redirect/ | Create a content redirect content guard
[**ContentguardsCoreContentRedirectDelete**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectDelete) | **Delete** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/ | Delete a content redirect content guard
[**ContentguardsCoreContentRedirectList**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectList) | **Get** /pulp/api/v3/contentguards/core/content_redirect/ | List content redirect content guards
[**ContentguardsCoreContentRedirectListRoles**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectListRoles) | **Get** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/list_roles/ | 
[**ContentguardsCoreContentRedirectMyPermissions**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectMyPermissions) | **Get** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/my_permissions/ | 
[**ContentguardsCoreContentRedirectPartialUpdate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectPartialUpdate) | **Patch** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/ | Update a content redirect content guard
[**ContentguardsCoreContentRedirectRead**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectRead) | **Get** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/ | Inspect a content redirect content guard
[**ContentguardsCoreContentRedirectRemoveRole**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectRemoveRole) | **Post** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/remove_role/ | 
[**ContentguardsCoreContentRedirectUpdate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectUpdate) | **Put** /pulp/api/v3/contentguards/core/content_redirect/{pulp_id}/ | Update a content redirect content guard



## ContentguardsCoreContentRedirectAddRole

> NestedRoleResponse ContentguardsCoreContentRedirectAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectAddRoleRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectCreate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectCreate(ctx).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()

Create a content redirect content guard



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
    contentRedirectContentGuard := *openapiclient.NewContentRedirectContentGuard("Name_example") // ContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate(context.Background()).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectCreate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentRedirectContentGuard** | [**ContentRedirectContentGuard**](ContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectDelete

> ContentguardsCoreContentRedirectDelete(ctx, pulpId).Execute()

Delete a content redirect content guard



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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectDeleteRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectList

> PaginatedContentRedirectContentGuardResponseList ContentguardsCoreContentRedirectList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List content redirect content guards



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
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectList`: PaginatedContentRedirectContentGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectListRequest struct via the builder pattern


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

[**PaginatedContentRedirectContentGuardResponseList**](PaginatedContentRedirectContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectListRoles

> ObjectRolesResponse ContentguardsCoreContentRedirectListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectListRolesRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectMyPermissions

> MyPermissionsResponse ContentguardsCoreContentRedirectMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectMyPermissionsRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectPartialUpdate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectPartialUpdate(ctx, pulpId).PatchedContentRedirectContentGuard(patchedContentRedirectContentGuard).Execute()

Update a content redirect content guard



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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    patchedContentRedirectContentGuard := *openapiclient.NewPatchedContentRedirectContentGuard() // PatchedContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate(context.Background(), pulpId).PatchedContentRedirectContentGuard(patchedContentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectPartialUpdate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedContentRedirectContentGuard** | [**PatchedContentRedirectContentGuard**](PatchedContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectRead

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a content redirect content guard



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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectRead`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectRemoveRole

> NestedRoleResponse ContentguardsCoreContentRedirectRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectRemoveRoleRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectUpdate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectUpdate(ctx, pulpId).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()

Update a content redirect content guard



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
    pulpId := TODO // string | A UUID string identifying this content redirect content guard.
    contentRedirectContentGuard := *openapiclient.NewContentRedirectContentGuard("Name_example") // ContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate(context.Background(), pulpId).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectUpdate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this content redirect content guard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentRedirectContentGuard** | [**ContentRedirectContentGuard**](ContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

