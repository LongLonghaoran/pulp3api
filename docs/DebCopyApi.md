# \DebCopyApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyContent**](DebCopyApi.md#CopyContent) | **Post** /pulp/api/v3/deb/copy/ | Copy content



## CopyContent

> AsyncOperationResponse CopyContent(ctx).Copy(copy).Execute()

Copy content



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
    copy := *openapiclient.NewCopy(map[string]interface{}(123)) // Copy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DebCopyApi.CopyContent(context.Background()).Copy(copy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebCopyApi.CopyContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyContent`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DebCopyApi.CopyContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopyContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **copy** | [**Copy**](Copy.md) |  | 

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

