/*
 * Pulp 3 API
 *
 * Fetch, Upload, Organize, and Distribute Software Packages
 *
 * API version: v3
 * Contact: pulp-list@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ContentReleaseArchitecturesApiService ContentReleaseArchitecturesApi service
type ContentReleaseArchitecturesApiService service

type ApiContentDebReleaseArchitecturesCreateRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseArchitecturesApiService
	debReleaseArchitecture *DebReleaseArchitecture
}

func (r ApiContentDebReleaseArchitecturesCreateRequest) DebReleaseArchitecture(debReleaseArchitecture DebReleaseArchitecture) ApiContentDebReleaseArchitecturesCreateRequest {
	r.debReleaseArchitecture = &debReleaseArchitecture
	return r
}

func (r ApiContentDebReleaseArchitecturesCreateRequest) Execute() (DebReleaseArchitectureResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseArchitecturesCreateExecute(r)
}

/*
 * ContentDebReleaseArchitecturesCreate Create a release architecture
 * A ReleaseArchitecture represents a single dpkg architecture string.

Associated artifacts: None; contains only metadata.

Every ReleaseArchitecture is always associated with exactly one Release. This indicates that
the release/distribution in question supports this architecture.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebReleaseArchitecturesCreateRequest
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesCreate(ctx _context.Context) ApiContentDebReleaseArchitecturesCreateRequest {
	return ApiContentDebReleaseArchitecturesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return DebReleaseArchitectureResponse
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesCreateExecute(r ApiContentDebReleaseArchitecturesCreateRequest) (DebReleaseArchitectureResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DebReleaseArchitectureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseArchitecturesApiService.ContentDebReleaseArchitecturesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_architectures/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.debReleaseArchitecture == nil {
		return localVarReturnValue, nil, reportError("debReleaseArchitecture is required and must be specified")
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
	localVarPostBody = r.debReleaseArchitecture
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContentDebReleaseArchitecturesListRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseArchitecturesApiService
	architecture *string
	limit *int32
	offset *int32
	ordering *[]string
	release *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *string
	excludeFields *string
}

func (r ApiContentDebReleaseArchitecturesListRequest) Architecture(architecture string) ApiContentDebReleaseArchitecturesListRequest {
	r.architecture = &architecture
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) Limit(limit int32) ApiContentDebReleaseArchitecturesListRequest {
	r.limit = &limit
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) Offset(offset int32) ApiContentDebReleaseArchitecturesListRequest {
	r.offset = &offset
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) Ordering(ordering []string) ApiContentDebReleaseArchitecturesListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) Release(release string) ApiContentDebReleaseArchitecturesListRequest {
	r.release = &release
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) RepositoryVersion(repositoryVersion string) ApiContentDebReleaseArchitecturesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ApiContentDebReleaseArchitecturesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ApiContentDebReleaseArchitecturesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) Fields(fields string) ApiContentDebReleaseArchitecturesListRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebReleaseArchitecturesListRequest) ExcludeFields(excludeFields string) ApiContentDebReleaseArchitecturesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleaseArchitecturesListRequest) Execute() (PaginateddebReleaseArchitectureResponseList, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseArchitecturesListExecute(r)
}

/*
 * ContentDebReleaseArchitecturesList List release architectures
 * A ReleaseArchitecture represents a single dpkg architecture string.

Associated artifacts: None; contains only metadata.

Every ReleaseArchitecture is always associated with exactly one Release. This indicates that
the release/distribution in question supports this architecture.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebReleaseArchitecturesListRequest
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesList(ctx _context.Context) ApiContentDebReleaseArchitecturesListRequest {
	return ApiContentDebReleaseArchitecturesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginateddebReleaseArchitectureResponseList
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesListExecute(r ApiContentDebReleaseArchitecturesListRequest) (PaginateddebReleaseArchitectureResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginateddebReleaseArchitectureResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseArchitecturesApiService.ContentDebReleaseArchitecturesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_architectures/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.architecture != nil {
		localVarQueryParams.Add("architecture", parameterToString(*r.architecture, ""))
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
	if r.release != nil {
		localVarQueryParams.Add("release", parameterToString(*r.release, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContentDebReleaseArchitecturesReadRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseArchitecturesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiContentDebReleaseArchitecturesReadRequest) Fields(fields string) ApiContentDebReleaseArchitecturesReadRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebReleaseArchitecturesReadRequest) ExcludeFields(excludeFields string) ApiContentDebReleaseArchitecturesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleaseArchitecturesReadRequest) Execute() (DebReleaseArchitectureResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseArchitecturesReadExecute(r)
}

/*
 * ContentDebReleaseArchitecturesRead Inspect a release architecture
 * A ReleaseArchitecture represents a single dpkg architecture string.

Associated artifacts: None; contains only metadata.

Every ReleaseArchitecture is always associated with exactly one Release. This indicates that
the release/distribution in question supports this architecture.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this release architecture.
 * @return ApiContentDebReleaseArchitecturesReadRequest
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesRead(ctx _context.Context, pulpId string) ApiContentDebReleaseArchitecturesReadRequest {
	return ApiContentDebReleaseArchitecturesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return DebReleaseArchitectureResponse
 */
func (a *ContentReleaseArchitecturesApiService) ContentDebReleaseArchitecturesReadExecute(r ApiContentDebReleaseArchitecturesReadRequest) (DebReleaseArchitectureResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DebReleaseArchitectureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseArchitecturesApiService.ContentDebReleaseArchitecturesRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_architectures/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}