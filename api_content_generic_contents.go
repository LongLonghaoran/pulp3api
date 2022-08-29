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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// ContentGenericContentsApiService ContentGenericContentsApi service
type ContentGenericContentsApiService service

type ApiContentDebGenericContentsCreateRequest struct {
	ctx _context.Context
	ApiService *ContentGenericContentsApiService
	relativePath *string
	artifact *string
	file **os.File
	repository *string
	upload *string
}

func (r ApiContentDebGenericContentsCreateRequest) RelativePath(relativePath string) ApiContentDebGenericContentsCreateRequest {
	r.relativePath = &relativePath
	return r
}
func (r ApiContentDebGenericContentsCreateRequest) Artifact(artifact string) ApiContentDebGenericContentsCreateRequest {
	r.artifact = &artifact
	return r
}
func (r ApiContentDebGenericContentsCreateRequest) File(file *os.File) ApiContentDebGenericContentsCreateRequest {
	r.file = &file
	return r
}
func (r ApiContentDebGenericContentsCreateRequest) Repository(repository string) ApiContentDebGenericContentsCreateRequest {
	r.repository = &repository
	return r
}
func (r ApiContentDebGenericContentsCreateRequest) Upload(upload string) ApiContentDebGenericContentsCreateRequest {
	r.upload = &upload
	return r
}

func (r ApiContentDebGenericContentsCreateRequest) Execute() (AsyncOperationResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebGenericContentsCreateExecute(r)
}

/*
 * ContentDebGenericContentsCreate Create a generic content
 * Trigger an asynchronous task to create content,optionally create new repository version.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebGenericContentsCreateRequest
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsCreate(ctx _context.Context) ApiContentDebGenericContentsCreateRequest {
	return ApiContentDebGenericContentsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AsyncOperationResponse
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsCreateExecute(r ApiContentDebGenericContentsCreateRequest) (AsyncOperationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.relativePath == nil {
		return localVarReturnValue, nil, reportError("relativePath is required and must be specified")
	}
	if strlen(*r.relativePath) < 1 {
		return localVarReturnValue, nil, reportError("relativePath must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	if r.artifact != nil {
		localVarFormParams.Add("artifact", parameterToString(*r.artifact, ""))
	}
	localVarFormParams.Add("relative_path", parameterToString(*r.relativePath, ""))
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	if r.repository != nil {
		localVarFormParams.Add("repository", parameterToString(*r.repository, ""))
	}
	if r.upload != nil {
		localVarFormParams.Add("upload", parameterToString(*r.upload, ""))
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

type ApiContentDebGenericContentsListRequest struct {
	ctx _context.Context
	ApiService *ContentGenericContentsApiService
	limit *int32
	offset *int32
	ordering *[]string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	fields *string
	excludeFields *string
}

func (r ApiContentDebGenericContentsListRequest) Limit(limit int32) ApiContentDebGenericContentsListRequest {
	r.limit = &limit
	return r
}
func (r ApiContentDebGenericContentsListRequest) Offset(offset int32) ApiContentDebGenericContentsListRequest {
	r.offset = &offset
	return r
}
func (r ApiContentDebGenericContentsListRequest) Ordering(ordering []string) ApiContentDebGenericContentsListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiContentDebGenericContentsListRequest) RelativePath(relativePath string) ApiContentDebGenericContentsListRequest {
	r.relativePath = &relativePath
	return r
}
func (r ApiContentDebGenericContentsListRequest) RepositoryVersion(repositoryVersion string) ApiContentDebGenericContentsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}
func (r ApiContentDebGenericContentsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ApiContentDebGenericContentsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}
func (r ApiContentDebGenericContentsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ApiContentDebGenericContentsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}
func (r ApiContentDebGenericContentsListRequest) Sha256(sha256 string) ApiContentDebGenericContentsListRequest {
	r.sha256 = &sha256
	return r
}
func (r ApiContentDebGenericContentsListRequest) Fields(fields string) ApiContentDebGenericContentsListRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebGenericContentsListRequest) ExcludeFields(excludeFields string) ApiContentDebGenericContentsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebGenericContentsListRequest) Execute() (PaginateddebGenericContentResponseList, *_nethttp.Response, error) {
	return r.ApiService.ContentDebGenericContentsListExecute(r)
}

/*
 * ContentDebGenericContentsList List generic contents
 * GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebGenericContentsListRequest
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsList(ctx _context.Context) ApiContentDebGenericContentsListRequest {
	return ApiContentDebGenericContentsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginateddebGenericContentResponseList
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsListExecute(r ApiContentDebGenericContentsListRequest) (PaginateddebGenericContentResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginateddebGenericContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
	}
	if r.relativePath != nil {
		localVarQueryParams.Add("relative_path", parameterToString(*r.relativePath, ""))
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
	if r.sha256 != nil {
		localVarQueryParams.Add("sha256", parameterToString(*r.sha256, ""))
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

type ApiContentDebGenericContentsReadRequest struct {
	ctx _context.Context
	ApiService *ContentGenericContentsApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiContentDebGenericContentsReadRequest) Fields(fields string) ApiContentDebGenericContentsReadRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebGenericContentsReadRequest) ExcludeFields(excludeFields string) ApiContentDebGenericContentsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebGenericContentsReadRequest) Execute() (DebGenericContentResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebGenericContentsReadExecute(r)
}

/*
 * ContentDebGenericContentsRead Inspect a generic content
 * GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this generic content.
 * @return ApiContentDebGenericContentsReadRequest
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsRead(ctx _context.Context, pulpId string) ApiContentDebGenericContentsReadRequest {
	return ApiContentDebGenericContentsReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return DebGenericContentResponse
 */
func (a *ContentGenericContentsApiService) ContentDebGenericContentsReadExecute(r ApiContentDebGenericContentsReadRequest) (DebGenericContentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DebGenericContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/{pulp_id}/"
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