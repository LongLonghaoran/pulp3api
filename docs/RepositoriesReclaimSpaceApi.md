# \RepositoriesReclaimSpaceApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesReclaimSpaceReclaim**](RepositoriesReclaimSpaceApi.md#RepositoriesReclaimSpaceReclaim) | **Post** /pulp/api/v3/repositories/reclaim_space/ | 



## RepositoriesReclaimSpaceReclaim

> AsyncOperationResponse RepositoriesReclaimSpaceReclaim(ctx).ReclaimSpace(reclaimSpace).Execute()





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
    reclaimSpace := *openapiclient.NewReclaimSpace([]interface{}{nil}) // ReclaimSpace | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesReclaimSpaceApi.RepositoriesReclaimSpaceReclaim(context.Background()).ReclaimSpace(reclaimSpace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesReclaimSpaceApi.RepositoriesReclaimSpaceReclaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesReclaimSpaceReclaim`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesReclaimSpaceApi.RepositoriesReclaimSpaceReclaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesReclaimSpaceReclaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reclaimSpace** | [**ReclaimSpace**](ReclaimSpace.md) |  | 

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

