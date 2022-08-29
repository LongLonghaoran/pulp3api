# \RepositoriesContainerApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesContainerContainerAdd**](RepositoriesContainerApi.md#RepositoriesContainerContainerAdd) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/add/ | Add content
[**RepositoriesContainerContainerAddRole**](RepositoriesContainerApi.md#RepositoriesContainerContainerAddRole) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/add_role/ | 
[**RepositoriesContainerContainerBuildImage**](RepositoriesContainerApi.md#RepositoriesContainerContainerBuildImage) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/build_image/ | Build an Image
[**RepositoriesContainerContainerCopyManifests**](RepositoriesContainerApi.md#RepositoriesContainerContainerCopyManifests) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/copy_manifests/ | Copy manifests
[**RepositoriesContainerContainerCopyTags**](RepositoriesContainerApi.md#RepositoriesContainerContainerCopyTags) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/copy_tags/ | Copy tags
[**RepositoriesContainerContainerCreate**](RepositoriesContainerApi.md#RepositoriesContainerContainerCreate) | **Post** /pulp/api/v3/repositories/container/container/ | Create a container repository
[**RepositoriesContainerContainerDelete**](RepositoriesContainerApi.md#RepositoriesContainerContainerDelete) | **Delete** /pulp/api/v3/repositories/container/container/{pulp_id}/ | Delete a container repository
[**RepositoriesContainerContainerList**](RepositoriesContainerApi.md#RepositoriesContainerContainerList) | **Get** /pulp/api/v3/repositories/container/container/ | List container repositorys
[**RepositoriesContainerContainerListRoles**](RepositoriesContainerApi.md#RepositoriesContainerContainerListRoles) | **Get** /pulp/api/v3/repositories/container/container/{pulp_id}/list_roles/ | 
[**RepositoriesContainerContainerMyPermissions**](RepositoriesContainerApi.md#RepositoriesContainerContainerMyPermissions) | **Get** /pulp/api/v3/repositories/container/container/{pulp_id}/my_permissions/ | 
[**RepositoriesContainerContainerPartialUpdate**](RepositoriesContainerApi.md#RepositoriesContainerContainerPartialUpdate) | **Patch** /pulp/api/v3/repositories/container/container/{pulp_id}/ | Update a container repository
[**RepositoriesContainerContainerRead**](RepositoriesContainerApi.md#RepositoriesContainerContainerRead) | **Get** /pulp/api/v3/repositories/container/container/{pulp_id}/ | Inspect a container repository
[**RepositoriesContainerContainerRemove**](RepositoriesContainerApi.md#RepositoriesContainerContainerRemove) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/remove/ | Remove content
[**RepositoriesContainerContainerRemoveRole**](RepositoriesContainerApi.md#RepositoriesContainerContainerRemoveRole) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/remove_role/ | 
[**RepositoriesContainerContainerSign**](RepositoriesContainerApi.md#RepositoriesContainerContainerSign) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/sign/ | Sign images in the repo
[**RepositoriesContainerContainerSync**](RepositoriesContainerApi.md#RepositoriesContainerContainerSync) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/sync/ | Sync from a remote
[**RepositoriesContainerContainerTag**](RepositoriesContainerApi.md#RepositoriesContainerContainerTag) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/tag/ | Create a Tag
[**RepositoriesContainerContainerUntag**](RepositoriesContainerApi.md#RepositoriesContainerContainerUntag) | **Post** /pulp/api/v3/repositories/container/container/{pulp_id}/untag/ | Delete a tag
[**RepositoriesContainerContainerUpdate**](RepositoriesContainerApi.md#RepositoriesContainerContainerUpdate) | **Put** /pulp/api/v3/repositories/container/container/{pulp_id}/ | Update a container repository



## RepositoriesContainerContainerAdd

> AsyncOperationResponse RepositoriesContainerContainerAdd(ctx, pulpId).RecursiveManage(recursiveManage).Execute()

Add content



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    recursiveManage := *openapiclient.NewRecursiveManage() // RecursiveManage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerAdd(context.Background(), pulpId).RecursiveManage(recursiveManage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerAdd`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursiveManage** | [**RecursiveManage**](RecursiveManage.md) |  | 

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


## RepositoriesContainerContainerAddRole

> NestedRoleResponse RepositoriesContainerContainerAddRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerAddRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerAddRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerBuildImage

> AsyncOperationResponse RepositoriesContainerContainerBuildImage(ctx, pulpId).ContainerfileArtifact(containerfileArtifact).Containerfile(containerfile).Tag(tag).Artifacts(artifacts).Execute()

Build an Image



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    containerfileArtifact := "containerfileArtifact_example" // string | Artifact representing the Containerfile that should be used to run podman-build. (optional)
    containerfile := os.NewFile(1234, "some_file") // *os.File | An uploaded Containerfile that should be used to run podman-build. (optional)
    tag := "tag_example" // string | A tag name for the new image being built. (optional) (default to "latest")
    artifacts := map[string]interface{}{ ... } // map[string]interface{} | A JSON string where each key is an artifact href and the value is it's relative path (name) inside the /pulp_working_directory of the build container executing the Containerfile. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerBuildImage(context.Background(), pulpId).ContainerfileArtifact(containerfileArtifact).Containerfile(containerfile).Tag(tag).Artifacts(artifacts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerBuildImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerBuildImage`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerBuildImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerBuildImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerfileArtifact** | **string** | Artifact representing the Containerfile that should be used to run podman-build. | 
 **containerfile** | ***os.File** | An uploaded Containerfile that should be used to run podman-build. | 
 **tag** | **string** | A tag name for the new image being built. | [default to &quot;latest&quot;]
 **artifacts** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string where each key is an artifact href and the value is it&#39;s relative path (name) inside the /pulp_working_directory of the build container executing the Containerfile. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerCopyManifests

> AsyncOperationResponse RepositoriesContainerContainerCopyManifests(ctx, pulpId).ManifestCopy(manifestCopy).Execute()

Copy manifests



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    manifestCopy := *openapiclient.NewManifestCopy() // ManifestCopy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCopyManifests(context.Background(), pulpId).ManifestCopy(manifestCopy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerCopyManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerCopyManifests`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerCopyManifests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerCopyManifestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manifestCopy** | [**ManifestCopy**](ManifestCopy.md) |  | 

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


## RepositoriesContainerContainerCopyTags

> AsyncOperationResponse RepositoriesContainerContainerCopyTags(ctx, pulpId).TagCopy(tagCopy).Execute()

Copy tags



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    tagCopy := *openapiclient.NewTagCopy() // TagCopy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCopyTags(context.Background(), pulpId).TagCopy(tagCopy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerCopyTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerCopyTags`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerCopyTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerCopyTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagCopy** | [**TagCopy**](TagCopy.md) |  | 

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


## RepositoriesContainerContainerCreate

> ContainerContainerRepositoryResponse RepositoriesContainerContainerCreate(ctx).ContainerContainerRepository(containerContainerRepository).Execute()

Create a container repository



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
    containerContainerRepository := *openapiclient.NewContainerContainerRepository("Name_example") // ContainerContainerRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCreate(context.Background()).ContainerContainerRepository(containerContainerRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerCreate`: ContainerContainerRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerContainerRepository** | [**ContainerContainerRepository**](ContainerContainerRepository.md) |  | 

### Return type

[**ContainerContainerRepositoryResponse**](ContainerContainerRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerDelete

> AsyncOperationResponse RepositoriesContainerContainerDelete(ctx, pulpId).Execute()

Delete a container repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerDeleteRequest struct via the builder pattern


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


## RepositoriesContainerContainerList

> PaginatedcontainerContainerRepositoryResponseList RepositoriesContainerContainerList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()

List container repositorys



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
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerList`: PaginatedcontainerContainerRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerListRequest struct via the builder pattern


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

[**PaginatedcontainerContainerRepositoryResponseList**](PaginatedcontainerContainerRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerListRoles

> ObjectRolesResponse RepositoriesContainerContainerListRoles(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerListRoles(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerListRolesRequest struct via the builder pattern


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


## RepositoriesContainerContainerMyPermissions

> MyPermissionsResponse RepositoriesContainerContainerMyPermissions(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerMyPermissions(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerMyPermissionsRequest struct via the builder pattern


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


## RepositoriesContainerContainerPartialUpdate

> AsyncOperationResponse RepositoriesContainerContainerPartialUpdate(ctx, pulpId).PatchedcontainerContainerRepository(patchedcontainerContainerRepository).Execute()

Update a container repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    patchedcontainerContainerRepository := *openapiclient.NewPatchedcontainerContainerRepository() // PatchedcontainerContainerRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerPartialUpdate(context.Background(), pulpId).PatchedcontainerContainerRepository(patchedcontainerContainerRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerRepository** | [**PatchedcontainerContainerRepository**](PatchedcontainerContainerRepository.md) |  | 

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


## RepositoriesContainerContainerRead

> ContainerContainerRepositoryResponse RepositoriesContainerContainerRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerRead`: ContainerContainerRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerRepositoryResponse**](ContainerContainerRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerRemove

> AsyncOperationResponse RepositoriesContainerContainerRemove(ctx, pulpId).RecursiveManage(recursiveManage).Execute()

Remove content



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    recursiveManage := *openapiclient.NewRecursiveManage() // RecursiveManage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRemove(context.Background(), pulpId).RecursiveManage(recursiveManage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerRemove`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursiveManage** | [**RecursiveManage**](RecursiveManage.md) |  | 

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


## RepositoriesContainerContainerRemoveRole

> NestedRoleResponse RepositoriesContainerContainerRemoveRole(ctx, pulpId).NestedRole(nestedRole).Execute()





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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRemoveRole(context.Background(), pulpId).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerRemoveRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerSign

> AsyncOperationResponse RepositoriesContainerContainerSign(ctx, pulpId).RepositorySign(repositorySign).Execute()

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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    repositorySign := *openapiclient.NewRepositorySign() // RepositorySign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerSign(context.Background(), pulpId).RepositorySign(repositorySign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerSign`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerSignRequest struct via the builder pattern


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


## RepositoriesContainerContainerSync

> AsyncOperationResponse RepositoriesContainerContainerSync(ctx, pulpId).ContainerRepositorySyncURL(containerRepositorySyncURL).Execute()

Sync from a remote



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    containerRepositorySyncURL := *openapiclient.NewContainerRepositorySyncURL() // ContainerRepositorySyncURL | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerSync(context.Background(), pulpId).ContainerRepositorySyncURL(containerRepositorySyncURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerSync`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRepositorySyncURL** | [**ContainerRepositorySyncURL**](ContainerRepositorySyncURL.md) |  | 

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


## RepositoriesContainerContainerTag

> AsyncOperationResponse RepositoriesContainerContainerTag(ctx, pulpId).TagImage(tagImage).Execute()

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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    tagImage := *openapiclient.NewTagImage("Tag_example", "Digest_example") // TagImage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerTag(context.Background(), pulpId).TagImage(tagImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerTag`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerTagRequest struct via the builder pattern


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


## RepositoriesContainerContainerUntag

> AsyncOperationResponse RepositoriesContainerContainerUntag(ctx, pulpId).UnTagImage(unTagImage).Execute()

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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    unTagImage := *openapiclient.NewUnTagImage("Tag_example") // UnTagImage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerUntag(context.Background(), pulpId).UnTagImage(unTagImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerUntag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerUntag`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerUntag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerUntagRequest struct via the builder pattern


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


## RepositoriesContainerContainerUpdate

> AsyncOperationResponse RepositoriesContainerContainerUpdate(ctx, pulpId).ContainerContainerRepository(containerContainerRepository).Execute()

Update a container repository



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this container repository.
    containerContainerRepository := *openapiclient.NewContainerContainerRepository("Name_example") // ContainerContainerRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerUpdate(context.Background(), pulpId).ContainerContainerRepository(containerContainerRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerApi.RepositoriesContainerContainerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerApi.RepositoriesContainerContainerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this container repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerRepository** | [**ContainerContainerRepository**](ContainerContainerRepository.md) |  | 

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

