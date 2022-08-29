# \RepositoriesContainerPushApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesContainerContainerPushAddRole**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushAddRole) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/add_role/ | 
[**RepositoriesContainerContainerPushList**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushList) | **Get** /pulp/api/v3/repositories/container/container-push/ | List container push repositorys
[**RepositoriesContainerContainerPushListRoles**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushListRoles) | **Get** /pulp/api/v3/repositories/container/container-push/{pulp_id}/list_roles/ | 
[**RepositoriesContainerContainerPushMyPermissions**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushMyPermissions) | **Get** /pulp/api/v3/repositories/container/container-push/{pulp_id}/my_permissions/ | 
[**RepositoriesContainerContainerPushPartialUpdate**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushPartialUpdate) | **Patch** /pulp/api/v3/repositories/container/container-push/{pulp_id}/ | Update a container push repository
[**RepositoriesContainerContainerPushRead**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushRead) | **Get** /pulp/api/v3/repositories/container/container-push/{pulp_id}/ | Inspect a container push repository
[**RepositoriesContainerContainerPushRemoveImage**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushRemoveImage) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/remove_image/ | Delete an image from a repository
[**RepositoriesContainerContainerPushRemoveRole**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushRemoveRole) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/remove_role/ | 
[**RepositoriesContainerContainerPushRemoveSignatures**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushRemoveSignatures) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/remove_signatures/ | 
[**RepositoriesContainerContainerPushSign**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushSign) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/sign/ | Sign images in the repo
[**RepositoriesContainerContainerPushTag**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushTag) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/tag/ | Create a Tag
[**RepositoriesContainerContainerPushUntag**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushUntag) | **Post** /pulp/api/v3/repositories/container/container-push/{pulp_id}/untag/ | Delete a tag
[**RepositoriesContainerContainerPushUpdate**](RepositoriesContainerPushApi.md#RepositoriesContainerContainerPushUpdate) | **Put** /pulp/api/v3/repositories/container/container-push/{pulp_id}/ | Update a container push repository



## RepositoriesContainerContainerPushAddRole

> NestedRoleResponse RepositoriesContainerContainerPushAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushAddRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushList

> PaginatedcontainerContainerPushRepositoryResponseList RepositoriesContainerContainerPushList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()

List container push repositorys



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
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushList`: PaginatedcontainerContainerPushRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushListRequest struct via the builder pattern


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
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerPushRepositoryResponseList**](PaginatedcontainerContainerPushRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushListRoles

> ObjectRolesResponse RepositoriesContainerContainerPushListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushListRolesRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushMyPermissions

> MyPermissionsResponse RepositoriesContainerContainerPushMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushMyPermissionsRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushPartialUpdate

> AsyncOperationResponse RepositoriesContainerContainerPushPartialUpdate(ctx, pulpId).PatchedcontainerContainerPushRepository(patchedcontainerContainerPushRepository).Execute()

Update a container push repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    patchedcontainerContainerPushRepository := *openapiclient.NewPatchedcontainerContainerPushRepository() // PatchedcontainerContainerPushRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushPartialUpdate(context.Background(), pulpId).PatchedcontainerContainerPushRepository(patchedcontainerContainerPushRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerPushRepository** | [**PatchedcontainerContainerPushRepository**](PatchedcontainerContainerPushRepository.md) |  | 

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


## RepositoriesContainerContainerPushRead

> ContainerContainerPushRepositoryResponse RepositoriesContainerContainerPushRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container push repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushRead`: ContainerContainerPushRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerPushRepositoryResponse**](ContainerContainerPushRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushRemoveImage

> AsyncOperationResponse RepositoriesContainerContainerPushRemoveImage(ctx, pulpId).RemoveImage(removeImage).Execute()

Delete an image from a repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    removeImage := *openapiclient.NewRemoveImage("Digest_example") // RemoveImage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveImage(context.Background(), pulpId).RemoveImage(removeImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushRemoveImage`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeImage** | [**RemoveImage**](RemoveImage.md) |  | 

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


## RepositoriesContainerContainerPushRemoveRole

> NestedRoleResponse RepositoriesContainerContainerPushRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushRemoveSignatures

> RemoveSignaturesResponse RepositoriesContainerContainerPushRemoveSignatures(ctx, pulpId).RemoveSignatures(removeSignatures).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    removeSignatures := *openapiclient.NewRemoveSignatures("SignedWithKeyId_example") // RemoveSignatures | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveSignatures(context.Background(), pulpId).RemoveSignatures(removeSignatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveSignatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushRemoveSignatures`: RemoveSignaturesResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveSignatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveSignaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeSignatures** | [**RemoveSignatures**](RemoveSignatures.md) |  | 

### Return type

[**RemoveSignaturesResponse**](RemoveSignaturesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushSign

> AsyncOperationResponse RepositoriesContainerContainerPushSign(ctx, pulpId).RepositorySign(repositorySign).Execute()

Sign images in the repo



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    repositorySign := *openapiclient.NewRepositorySign() // RepositorySign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushSign(context.Background(), pulpId).RepositorySign(repositorySign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushSign`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySign** | [**RepositorySign**](RepositorySign.md) |  | 

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


## RepositoriesContainerContainerPushTag

> AsyncOperationResponse RepositoriesContainerContainerPushTag(ctx, pulpId).TagImage(tagImage).Execute()

Create a Tag



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    tagImage := *openapiclient.NewTagImage("Tag_example", "Digest_example") // TagImage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushTag(context.Background(), pulpId).TagImage(tagImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushTag`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagImage** | [**TagImage**](TagImage.md) |  | 

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


## RepositoriesContainerContainerPushUntag

> AsyncOperationResponse RepositoriesContainerContainerPushUntag(ctx, pulpId).UnTagImage(unTagImage).Execute()

Delete a tag



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    unTagImage := *openapiclient.NewUnTagImage("Tag_example") // UnTagImage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushUntag(context.Background(), pulpId).UnTagImage(unTagImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushUntag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushUntag`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushUntag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushUntagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unTagImage** | [**UnTagImage**](UnTagImage.md) |  | 

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


## RepositoriesContainerContainerPushUpdate

> AsyncOperationResponse RepositoriesContainerContainerPushUpdate(ctx, pulpId).ContainerContainerPushRepository(containerContainerPushRepository).Execute()

Update a container push repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container push repository.
    containerContainerPushRepository := *openapiclient.NewContainerContainerPushRepository("Name_example") // ContainerContainerPushRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushUpdate(context.Background(), pulpId).ContainerContainerPushRepository(containerContainerPushRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushApi.RepositoriesContainerContainerPushUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushApi.RepositoriesContainerContainerPushUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container push repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerPushRepository** | [**ContainerContainerPushRepository**](ContainerContainerPushRepository.md) |  | 

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

