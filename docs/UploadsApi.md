# \UploadsApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadsCommit**](UploadsApi.md#UploadsCommit) | **Post** /pulp/api/v3/uploads/{pulp_id}/commit/ | Finish an Upload
[**UploadsCreate**](UploadsApi.md#UploadsCreate) | **Post** /pulp/api/v3/uploads/ | Create an upload
[**UploadsDelete**](UploadsApi.md#UploadsDelete) | **Delete** /pulp/api/v3/uploads/{pulp_id}/ | Delete an upload
[**UploadsList**](UploadsApi.md#UploadsList) | **Get** /pulp/api/v3/uploads/ | List uploads
[**UploadsRead**](UploadsApi.md#UploadsRead) | **Get** /pulp/api/v3/uploads/{pulp_id}/ | Inspect an upload
[**UploadsUpdate**](UploadsApi.md#UploadsUpdate) | **Put** /pulp/api/v3/uploads/{pulp_id}/ | Upload a file chunk



## UploadsCommit

> AsyncOperationResponse UploadsCommit(ctx, pulpId).UploadCommit(uploadCommit).Execute()

Finish an Upload



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
    pulpId := TODO // string | A UUID string identifying this upload.
    uploadCommit := *openapiclient.NewUploadCommit("Sha256_example") // UploadCommit | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsCommit(context.Background(), pulpId).UploadCommit(uploadCommit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsCommit`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadCommit** | [**UploadCommit**](UploadCommit.md) |  | 

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


## UploadsCreate

> UploadResponse UploadsCreate(ctx).Upload(upload).Execute()

Create an upload



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
    upload := *openapiclient.NewUpload(int32(123)) // Upload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsCreate(context.Background()).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsCreate`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upload** | [**Upload**](Upload.md) |  | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsDelete

> UploadsDelete(ctx, pulpId).Execute()

Delete an upload



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
    pulpId := TODO // string | A UUID string identifying this upload.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsDelete(context.Background(), pulpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsDeleteRequest struct via the builder pattern


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


## UploadsList

> PaginatedUploadResponseList UploadsList(ctx).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List uploads



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsList(context.Background()).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsList`: PaginatedUploadResponseList
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUploadResponseList**](PaginatedUploadResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsRead

> UploadDetailResponse UploadsRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upload



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
    pulpId := TODO // string | A UUID string identifying this upload.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsRead`: UploadDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**UploadDetailResponse**](UploadDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsUpdate

> UploadResponse UploadsUpdate(ctx, pulpId).ContentRange(contentRange).File(file).Sha256(sha256).Execute()

Upload a file chunk



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
    contentRange := "contentRange_example" // string | The Content-Range header specifies the location of the file chunk within the file.
    pulpId := TODO // string | A UUID string identifying this upload.
    file := os.NewFile(1234, "some_file") // *os.File | A chunk of the uploaded file.
    sha256 := "sha256_example" // string | The SHA-256 checksum of the chunk if available. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadsApi.UploadsUpdate(context.Background(), pulpId).ContentRange(contentRange).File(file).Sha256(sha256).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsUpdate`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | [**string**](.md) | A UUID string identifying this upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentRange** | **string** | The Content-Range header specifies the location of the file chunk within the file. | 

 **file** | ***os.File** | A chunk of the uploaded file. | 
 **sha256** | **string** | The SHA-256 checksum of the chunk if available. | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

