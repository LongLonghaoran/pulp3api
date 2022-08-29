# \ContentInstallerPackagesApi

All URIs are relative to *http://121.37.218.63:24817*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebInstallerPackagesCreate**](ContentInstallerPackagesApi.md#ContentDebInstallerPackagesCreate) | **Post** /pulp/api/v3/content/deb/installer_packages/ | Create an installer package
[**ContentDebInstallerPackagesList**](ContentInstallerPackagesApi.md#ContentDebInstallerPackagesList) | **Get** /pulp/api/v3/content/deb/installer_packages/ | List installer packages
[**ContentDebInstallerPackagesRead**](ContentInstallerPackagesApi.md#ContentDebInstallerPackagesRead) | **Get** /pulp/api/v3/content/deb/installer_packages/{pulp_id}/ | Inspect an installer package



## ContentDebInstallerPackagesCreate

> AsyncOperationResponse ContentDebInstallerPackagesCreate(ctx).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()

Create an installer package



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
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesCreate(context.Background()).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesApi.ContentDebInstallerPackagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerPackagesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesApi.ContentDebInstallerPackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifact** | **string** | Artifact file representing the physical content | 
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 

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


## ContentDebInstallerPackagesList

> PaginateddebBasePackageResponseList ContentDebInstallerPackagesList(ctx).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).Package_(package_).Priority(priority).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List installer packages



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
    architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
    autoBuiltPackage := "autoBuiltPackage_example" // string | Filter results where auto_built_package matches value (optional)
    buildEssential := true // bool | Filter results where build_essential matches value (optional)
    builtUsing := "builtUsing_example" // string | Filter results where built_using matches value (optional)
    essential := true // bool | Filter results where essential matches value (optional)
    installedSize := int32(56) // int32 | Filter results where installed_size matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    maintainer := "maintainer_example" // string | Filter results where maintainer matches value (optional)
    multiArch := "multiArch_example" // string | Filter results where multi_arch matches value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    origin := "origin_example" // string | Filter results where origin matches value (optional)
    originalMaintainer := "originalMaintainer_example" // string | Filter results where original_maintainer matches value (optional)
    package_ := "package__example" // string | Filter results where package matches value (optional)
    priority := "priority_example" // string | Filter results where priority matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    section := "section_example" // string | Filter results where section matches value (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    source := "source_example" // string | Filter results where source matches value (optional)
    tag := "tag_example" // string | Filter results where tag matches value (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesList(context.Background()).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).Package_(package_).Priority(priority).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesApi.ContentDebInstallerPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerPackagesList`: PaginateddebBasePackageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesApi.ContentDebInstallerPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **autoBuiltPackage** | **string** | Filter results where auto_built_package matches value | 
 **buildEssential** | **bool** | Filter results where build_essential matches value | 
 **builtUsing** | **string** | Filter results where built_using matches value | 
 **essential** | **bool** | Filter results where essential matches value | 
 **installedSize** | **int32** | Filter results where installed_size matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **maintainer** | **string** | Filter results where maintainer matches value | 
 **multiArch** | **string** | Filter results where multi_arch matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **origin** | **string** | Filter results where origin matches value | 
 **originalMaintainer** | **string** | Filter results where original_maintainer matches value | 
 **package_** | **string** | Filter results where package matches value | 
 **priority** | **string** | Filter results where priority matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **section** | **string** | Filter results where section matches value | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **source** | **string** | Filter results where source matches value | 
 **tag** | **string** | Filter results where tag matches value | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebBasePackageResponseList**](PaginateddebBasePackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerPackagesRead

> DebBasePackageResponse ContentDebInstallerPackagesRead(ctx, pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an installer package



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
    pulpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this installer package.
    fields := "fields_example" // string | A list of fields to include in the response. (optional)
    excludeFields := "excludeFields_example" // string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesRead(context.Background(), pulpId).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesApi.ContentDebInstallerPackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebInstallerPackagesRead`: DebBasePackageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesApi.ContentDebInstallerPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpId** | **string** | A UUID string identifying this installer package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of fields to include in the response. | 
 **excludeFields** | **string** | A list of fields to exclude from the response. | 

### Return type

[**DebBasePackageResponse**](DebBasePackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

