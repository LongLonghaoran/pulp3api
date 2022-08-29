# \ArtifactsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactsCreate**](ArtifactsApi.md#ArtifactsCreate) | **Post** /pulp/api/v3/artifacts/ | Create an artifact
[**ArtifactsDelete**](ArtifactsApi.md#ArtifactsDelete) | **Delete** /pulp/api/v3/artifacts/{pulp_id}/ | Delete an artifact
[**ArtifactsList**](ArtifactsApi.md#ArtifactsList) | **Get** /pulp/api/v3/artifacts/ | List artifacts
[**ArtifactsRead**](ArtifactsApi.md#ArtifactsRead) | **Get** /pulp/api/v3/artifacts/{pulp_id}/ | Inspect an artifact



## ArtifactsCreate

> ArtifactResponse ArtifactsCreate(ctx).File(file).Size(size).Md5(md5).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Execute()

Create an artifact



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
    file := os.NewFile(1234, "some_file") // *os.File | The stored file.
    size := int32(56) // int32 | The size of the file in bytes. (optional)
    md5 := "md5_example" // string | The MD5 checksum of the file if available. (optional)
    sha1 := "sha1_example" // string | The SHA-1 checksum of the file if available. (optional)
    sha224 := "sha224_example" // string | The SHA-224 checksum of the file if available. (optional)
    sha256 := "sha256_example" // string | The SHA-256 checksum of the file if available. (optional)
    sha384 := "sha384_example" // string | The SHA-384 checksum of the file if available. (optional)
    sha512 := "sha512_example" // string | The SHA-512 checksum of the file if available. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ArtifactsCreate(context.Background()).File(file).Size(size).Md5(md5).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ArtifactsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsCreate`: ArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The stored file. | 
 **size** | **int32** | The size of the file in bytes. | 
 **md5** | **string** | The MD5 checksum of the file if available. | 
 **sha1** | **string** | The SHA-1 checksum of the file if available. | 
 **sha224** | **string** | The SHA-224 checksum of the file if available. | 
 **sha256** | **string** | The SHA-256 checksum of the file if available. | 
 **sha384** | **string** | The SHA-384 checksum of the file if available. | 
 **sha512** | **string** | The SHA-512 checksum of the file if available. | 

### Return type

[**ArtifactResponse**](ArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactsDelete

> ArtifactsDelete(ctx, pulpId).Execute()

Delete an artifact



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
    pulpId := TODO // string | A UUID string identifying this artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ArtifactsDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ArtifactsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsDeleteRequest struct via the builder pattern


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


## ArtifactsList

> PaginatedArtifactResponseList ArtifactsList(ctx).Limit(limit).Md5(md5).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Fields(fields).ExcludeFields(excludeFields).Execute()

List artifacts



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
    md5 := "md5_example" // string | Filter results where md5 matches value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    sha1 := "sha1_example" // string | Filter results where sha1 matches value (optional)
    sha224 := "sha224_example" // string | Filter results where sha224 matches value (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    sha384 := "sha384_example" // string | Filter results where sha384 matches value (optional)
    sha512 := "sha512_example" // string | Filter results where sha512 matches value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ArtifactsList(context.Background()).Limit(limit).Md5(md5).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).Sha1(sha1).Sha224(sha224).Sha256(sha256).Sha384(sha384).Sha512(sha512).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ArtifactsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsList`: PaginatedArtifactResponseList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ArtifactsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **md5** | **string** | Filter results where md5 matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **sha1** | **string** | Filter results where sha1 matches value | 
 **sha224** | **string** | Filter results where sha224 matches value | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **sha384** | **string** | Filter results where sha384 matches value | 
 **sha512** | **string** | Filter results where sha512 matches value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedArtifactResponseList**](PaginatedArtifactResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactsRead

> ArtifactResponse ArtifactsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an artifact



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
    pulpId := TODO // string | A UUID string identifying this artifact.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ArtifactsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ArtifactsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactsRead`: ArtifactResponse
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ArtifactsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**ArtifactResponse**](ArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

