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

// ContentReleaseFilesApiService ContentReleaseFilesApi service
type ContentReleaseFilesApiService service

type ApiContentDebReleaseFilesCreateRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseFilesApiService
	debReleaseFile *DebReleaseFile
}

func (r ApiContentDebReleaseFilesCreateRequest) DebReleaseFile(debReleaseFile DebReleaseFile) ApiContentDebReleaseFilesCreateRequest {
	r.debReleaseFile = &debReleaseFile
	return r
}

func (r ApiContentDebReleaseFilesCreateRequest) Execute() (DebReleaseFileResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseFilesCreateExecute(r)
}

/*
 * ContentDebReleaseFilesCreate Create a release file
 * A ReleaseFile represents the Release file(s) from a single APT distribution.

Associated artifacts: At least one of 'Release' and 'InRelease' file. If the 'Release' file is
present, then there may also be a 'Release.gpg' detached signature file for it.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Release' files it needs when creating the
publication. It does not make use of ReleaseFile content.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebReleaseFilesCreateRequest
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesCreate(ctx _context.Context) ApiContentDebReleaseFilesCreateRequest {
	return ApiContentDebReleaseFilesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return DebReleaseFileResponse
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesCreateExecute(r ApiContentDebReleaseFilesCreateRequest) (DebReleaseFileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DebReleaseFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseFilesApiService.ContentDebReleaseFilesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_files/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.debReleaseFile == nil {
		return localVarReturnValue, nil, reportError("debReleaseFile is required and must be specified")
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
	localVarPostBody = r.debReleaseFile
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

type ApiContentDebReleaseFilesListRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseFilesApiService
	codename *string
	limit *int32
	offset *int32
	ordering *[]string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	suite *string
	fields *string
	excludeFields *string
}

func (r ApiContentDebReleaseFilesListRequest) Codename(codename string) ApiContentDebReleaseFilesListRequest {
	r.codename = &codename
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Limit(limit int32) ApiContentDebReleaseFilesListRequest {
	r.limit = &limit
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Offset(offset int32) ApiContentDebReleaseFilesListRequest {
	r.offset = &offset
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Ordering(ordering []string) ApiContentDebReleaseFilesListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiContentDebReleaseFilesListRequest) RelativePath(relativePath string) ApiContentDebReleaseFilesListRequest {
	r.relativePath = &relativePath
	return r
}
func (r ApiContentDebReleaseFilesListRequest) RepositoryVersion(repositoryVersion string) ApiContentDebReleaseFilesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}
func (r ApiContentDebReleaseFilesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ApiContentDebReleaseFilesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}
func (r ApiContentDebReleaseFilesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ApiContentDebReleaseFilesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Sha256(sha256 string) ApiContentDebReleaseFilesListRequest {
	r.sha256 = &sha256
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Suite(suite string) ApiContentDebReleaseFilesListRequest {
	r.suite = &suite
	return r
}
func (r ApiContentDebReleaseFilesListRequest) Fields(fields string) ApiContentDebReleaseFilesListRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebReleaseFilesListRequest) ExcludeFields(excludeFields string) ApiContentDebReleaseFilesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleaseFilesListRequest) Execute() (PaginateddebReleaseFileResponseList, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseFilesListExecute(r)
}

/*
 * ContentDebReleaseFilesList List release files
 * A ReleaseFile represents the Release file(s) from a single APT distribution.

Associated artifacts: At least one of 'Release' and 'InRelease' file. If the 'Release' file is
present, then there may also be a 'Release.gpg' detached signature file for it.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Release' files it needs when creating the
publication. It does not make use of ReleaseFile content.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiContentDebReleaseFilesListRequest
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesList(ctx _context.Context) ApiContentDebReleaseFilesListRequest {
	return ApiContentDebReleaseFilesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginateddebReleaseFileResponseList
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesListExecute(r ApiContentDebReleaseFilesListRequest) (PaginateddebReleaseFileResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginateddebReleaseFileResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseFilesApiService.ContentDebReleaseFilesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_files/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.codename != nil {
		localVarQueryParams.Add("codename", parameterToString(*r.codename, ""))
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

type ApiContentDebReleaseFilesReadRequest struct {
	ctx _context.Context
	ApiService *ContentReleaseFilesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiContentDebReleaseFilesReadRequest) Fields(fields string) ApiContentDebReleaseFilesReadRequest {
	r.fields = &fields
	return r
}
func (r ApiContentDebReleaseFilesReadRequest) ExcludeFields(excludeFields string) ApiContentDebReleaseFilesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentDebReleaseFilesReadRequest) Execute() (DebReleaseFileResponse, *_nethttp.Response, error) {
	return r.ApiService.ContentDebReleaseFilesReadExecute(r)
}

/*
 * ContentDebReleaseFilesRead Inspect a release file
 * A ReleaseFile represents the Release file(s) from a single APT distribution.

Associated artifacts: At least one of 'Release' and 'InRelease' file. If the 'Release' file is
present, then there may also be a 'Release.gpg' detached signature file for it.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Release' files it needs when creating the
publication. It does not make use of ReleaseFile content.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this release file.
 * @return ApiContentDebReleaseFilesReadRequest
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesRead(ctx _context.Context, pulpId string) ApiContentDebReleaseFilesReadRequest {
	return ApiContentDebReleaseFilesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return DebReleaseFileResponse
 */
func (a *ContentReleaseFilesApiService) ContentDebReleaseFilesReadExecute(r ApiContentDebReleaseFilesReadRequest) (DebReleaseFileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DebReleaseFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentReleaseFilesApiService.ContentDebReleaseFilesRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/release_files/{pulp_id}/"
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
