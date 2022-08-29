/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ContentReleasesApiService ContentReleasesApi service
type ContentReleasesApiService service

type ApiContentDebReleasesCreateRequest struct {
	ctx context.Context
	ApiService *ContentReleasesApiService
	debRelease *DebRelease
}

func (r ApiContentDebReleasesCreateRequest) DebRelease(debRelease DebRelease) ApiContentDebReleasesCreateRequest {
	r.debRelease = &debRelease
	return r
}

func (r ApiContentDebReleasesCreateRequest) Execute() (*DebReleaseResponse, *http.Response, error) {
	return r.ApiService.ContentDebReleasesCreateExecute(r)
}

/*
ContentDebReleasesCreate Create a release

A Release represents a single APT release/distribution.

Associated artifacts: None; contains only metadata.

Note that in the context of the "Release content", the terms "distribution" and "release"
are synonyms. An "APT repository release/distribution" is associated with a single 'Release'
file below the 'dists/' folder. The "distribution" refers to the path between 'dists/' and the
'Release' file. The "distribution" could be considered the name of the "release". It is often
(but not always) equal to the "codename" or "suite".

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContentDebReleasesCreateRequest
*/
func (a *ContentReleasesApiService) ContentDebReleasesCreate(ctx context.Context) ApiContentDebReleasesCreateRequest {
	return ApiContentDebReleasesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebReleaseResponse
func (a *ContentReleasesApiService) ContentDebReleasesCreateExecute(r ApiContentDebReleasesCreateRequest) (*DebReleaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebReleaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleasesApiService.ContentDebReleasesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/releases/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debRelease == nil {
		return localVarReturnValue, nil, reportError("debRelease is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.debRelease
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContentDebReleasesListRequest struct {
	ctx context.Context
	ApiService *ContentReleasesApiService
	codename *string
	distribution *string
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	suite *string
	fields *string
	excludeFields *string
}

// Filter results where codename matches value
func (r ApiContentDebReleasesListRequest) Codename(codename string) ApiContentDebReleasesListRequest {
	r.codename = &codename
	return r
}

// Filter results where distribution matches value
func (r ApiContentDebReleasesListRequest) Distribution(distribution string) ApiContentDebReleasesListRequest {
	r.distribution = &distribution
	return r
}

// Number of results to return per page.
func (r ApiContentDebReleasesListRequest) Limit(limit int32) ApiContentDebReleasesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiContentDebReleasesListRequest) Offset(offset int32) ApiContentDebReleasesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiContentDebReleasesListRequest) Ordering(ordering []string) ApiContentDebReleasesListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ApiContentDebReleasesListRequest) RepositoryVersion(repositoryVersion string) ApiContentDebReleasesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ApiContentDebReleasesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ApiContentDebReleasesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ApiContentDebReleasesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ApiContentDebReleasesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where suite matches value
func (r ApiContentDebReleasesListRequest) Suite(suite string) ApiContentDebReleasesListRequest {
	r.suite = &suite
	return r
}

// A list of fields to include in the response.
func (r ApiContentDebReleasesListRequest) Fields(fields string) ApiContentDebReleasesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiContentDebReleasesListRequest) ExcludeFields(excludeFields string) ApiContentDebReleasesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleasesListRequest) Execute() (*PaginateddebReleaseResponseList, *http.Response, error) {
	return r.ApiService.ContentDebReleasesListExecute(r)
}

/*
ContentDebReleasesList List releases

A Release represents a single APT release/distribution.

Associated artifacts: None; contains only metadata.

Note that in the context of the "Release content", the terms "distribution" and "release"
are synonyms. An "APT repository release/distribution" is associated with a single 'Release'
file below the 'dists/' folder. The "distribution" refers to the path between 'dists/' and the
'Release' file. The "distribution" could be considered the name of the "release". It is often
(but not always) equal to the "codename" or "suite".

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContentDebReleasesListRequest
*/
func (a *ContentReleasesApiService) ContentDebReleasesList(ctx context.Context) ApiContentDebReleasesListRequest {
	return ApiContentDebReleasesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebReleaseResponseList
func (a *ContentReleasesApiService) ContentDebReleasesListExecute(r ApiContentDebReleasesListRequest) (*PaginateddebReleaseResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebReleaseResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleasesApiService.ContentDebReleasesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/releases/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.codename != nil {
		localVarQueryParams.Add("codename", parameterToString(*r.codename, ""))
	}
	if r.distribution != nil {
		localVarQueryParams.Add("distribution", parameterToString(*r.distribution, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
	}
	if r.repositoryVersion != nil {
		localVarQueryParams.Add("repository_version", parameterToString(*r.repositoryVersion, ""))
	}
	if r.repositoryVersionAdded != nil {
		localVarQueryParams.Add("repository_version_added", parameterToString(*r.repositoryVersionAdded, ""))
	}
	if r.repositoryVersionRemoved != nil {
		localVarQueryParams.Add("repository_version_removed", parameterToString(*r.repositoryVersionRemoved, ""))
	}
	if r.suite != nil {
		localVarQueryParams.Add("suite", parameterToString(*r.suite, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.excludeFields != nil {
		localVarQueryParams.Add("exclude_fields", parameterToString(*r.excludeFields, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContentDebReleasesReadRequest struct {
	ctx context.Context
	ApiService *ContentReleasesApiService
	pulpId string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiContentDebReleasesReadRequest) Fields(fields string) ApiContentDebReleasesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiContentDebReleasesReadRequest) ExcludeFields(excludeFields string) ApiContentDebReleasesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleasesReadRequest) Execute() (*DebReleaseResponse, *http.Response, error) {
	return r.ApiService.ContentDebReleasesReadExecute(r)
}

/*
ContentDebReleasesRead Inspect a release

A Release represents a single APT release/distribution.

Associated artifacts: None; contains only metadata.

Note that in the context of the "Release content", the terms "distribution" and "release"
are synonyms. An "APT repository release/distribution" is associated with a single 'Release'
file below the 'dists/' folder. The "distribution" refers to the path between 'dists/' and the
'Release' file. The "distribution" could be considered the name of the "release". It is often
(but not always) equal to the "codename" or "suite".

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this release.
 @return ApiContentDebReleasesReadRequest
*/
func (a *ContentReleasesApiService) ContentDebReleasesRead(ctx context.Context, pulpId string) ApiContentDebReleasesReadRequest {
	return ApiContentDebReleasesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return DebReleaseResponse
func (a *ContentReleasesApiService) ContentDebReleasesReadExecute(r ApiContentDebReleasesReadRequest) (*DebReleaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebReleaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleasesApiService.ContentDebReleasesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/releases/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.excludeFields != nil {
		localVarQueryParams.Add("exclude_fields", parameterToString(*r.excludeFields, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
