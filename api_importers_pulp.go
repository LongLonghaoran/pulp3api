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

// ImportersPulpApiService ImportersPulpApi service
type ImportersPulpApiService service

type ApiImportersCorePulpCreateRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	pulpImporter *PulpImporter
}

func (r ApiImportersCorePulpCreateRequest) PulpImporter(pulpImporter PulpImporter) ApiImportersCorePulpCreateRequest {
	r.pulpImporter = &pulpImporter
	return r
}

func (r ApiImportersCorePulpCreateRequest) Execute() (PulpImporterResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpCreateExecute(r)
}

/*
 * ImportersCorePulpCreate Create a pulp importer
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiImportersCorePulpCreateRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpCreate(ctx _context.Context) ApiImportersCorePulpCreateRequest {
	return ApiImportersCorePulpCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PulpImporterResponse
 */
func (a *ImportersPulpApiService) ImportersCorePulpCreateExecute(r ApiImportersCorePulpCreateRequest) (PulpImporterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pulpImporter == nil {
		return localVarReturnValue, nil, reportError("pulpImporter is required and must be specified")
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
	localVarPostBody = r.pulpImporter
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

type ApiImportersCorePulpDeleteRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	pulpId string
}


func (r ApiImportersCorePulpDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpDeleteExecute(r)
}

/*
 * ImportersCorePulpDelete Delete a pulp importer
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this pulp importer.
 * @return ApiImportersCorePulpDeleteRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpDelete(ctx _context.Context, pulpId string) ApiImportersCorePulpDeleteRequest {
	return ApiImportersCorePulpDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 */
func (a *ImportersPulpApiService) ImportersCorePulpDeleteExecute(r ApiImportersCorePulpDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiImportersCorePulpListRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	fields *string
	excludeFields *string
}

func (r ApiImportersCorePulpListRequest) Limit(limit int32) ApiImportersCorePulpListRequest {
	r.limit = &limit
	return r
}
func (r ApiImportersCorePulpListRequest) Name(name string) ApiImportersCorePulpListRequest {
	r.name = &name
	return r
}
func (r ApiImportersCorePulpListRequest) NameContains(nameContains string) ApiImportersCorePulpListRequest {
	r.nameContains = &nameContains
	return r
}
func (r ApiImportersCorePulpListRequest) NameIcontains(nameIcontains string) ApiImportersCorePulpListRequest {
	r.nameIcontains = &nameIcontains
	return r
}
func (r ApiImportersCorePulpListRequest) NameIn(nameIn []string) ApiImportersCorePulpListRequest {
	r.nameIn = &nameIn
	return r
}
func (r ApiImportersCorePulpListRequest) NameStartswith(nameStartswith string) ApiImportersCorePulpListRequest {
	r.nameStartswith = &nameStartswith
	return r
}
func (r ApiImportersCorePulpListRequest) Offset(offset int32) ApiImportersCorePulpListRequest {
	r.offset = &offset
	return r
}
func (r ApiImportersCorePulpListRequest) Ordering(ordering []string) ApiImportersCorePulpListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiImportersCorePulpListRequest) Fields(fields string) ApiImportersCorePulpListRequest {
	r.fields = &fields
	return r
}
func (r ApiImportersCorePulpListRequest) ExcludeFields(excludeFields string) ApiImportersCorePulpListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiImportersCorePulpListRequest) Execute() (PaginatedPulpImporterResponseList, *_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpListExecute(r)
}

/*
 * ImportersCorePulpList List pulp importers
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiImportersCorePulpListRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpList(ctx _context.Context) ApiImportersCorePulpListRequest {
	return ApiImportersCorePulpListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedPulpImporterResponseList
 */
func (a *ImportersPulpApiService) ImportersCorePulpListExecute(r ApiImportersCorePulpListRequest) (PaginatedPulpImporterResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedPulpImporterResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.nameContains != nil {
		localVarQueryParams.Add("name__contains", parameterToString(*r.nameContains, ""))
	}
	if r.nameIcontains != nil {
		localVarQueryParams.Add("name__icontains", parameterToString(*r.nameIcontains, ""))
	}
	if r.nameIn != nil {
		localVarQueryParams.Add("name__in", parameterToString(*r.nameIn, "csv"))
	}
	if r.nameStartswith != nil {
		localVarQueryParams.Add("name__startswith", parameterToString(*r.nameStartswith, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
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

type ApiImportersCorePulpPartialUpdateRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	pulpId string
	patchedPulpImporter *PatchedPulpImporter
}

func (r ApiImportersCorePulpPartialUpdateRequest) PatchedPulpImporter(patchedPulpImporter PatchedPulpImporter) ApiImportersCorePulpPartialUpdateRequest {
	r.patchedPulpImporter = &patchedPulpImporter
	return r
}

func (r ApiImportersCorePulpPartialUpdateRequest) Execute() (PulpImporterResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpPartialUpdateExecute(r)
}

/*
 * ImportersCorePulpPartialUpdate Update a pulp importer
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this pulp importer.
 * @return ApiImportersCorePulpPartialUpdateRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpPartialUpdate(ctx _context.Context, pulpId string) ApiImportersCorePulpPartialUpdateRequest {
	return ApiImportersCorePulpPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return PulpImporterResponse
 */
func (a *ImportersPulpApiService) ImportersCorePulpPartialUpdateExecute(r ApiImportersCorePulpPartialUpdateRequest) (PulpImporterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.patchedPulpImporter == nil {
		return localVarReturnValue, nil, reportError("patchedPulpImporter is required and must be specified")
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
	localVarPostBody = r.patchedPulpImporter
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

type ApiImportersCorePulpReadRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiImportersCorePulpReadRequest) Fields(fields string) ApiImportersCorePulpReadRequest {
	r.fields = &fields
	return r
}
func (r ApiImportersCorePulpReadRequest) ExcludeFields(excludeFields string) ApiImportersCorePulpReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiImportersCorePulpReadRequest) Execute() (PulpImporterResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpReadExecute(r)
}

/*
 * ImportersCorePulpRead Inspect a pulp importer
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this pulp importer.
 * @return ApiImportersCorePulpReadRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpRead(ctx _context.Context, pulpId string) ApiImportersCorePulpReadRequest {
	return ApiImportersCorePulpReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return PulpImporterResponse
 */
func (a *ImportersPulpApiService) ImportersCorePulpReadExecute(r ApiImportersCorePulpReadRequest) (PulpImporterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/{pulp_id}/"
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

type ApiImportersCorePulpUpdateRequest struct {
	ctx _context.Context
	ApiService *ImportersPulpApiService
	pulpId string
	pulpImporter *PulpImporter
}

func (r ApiImportersCorePulpUpdateRequest) PulpImporter(pulpImporter PulpImporter) ApiImportersCorePulpUpdateRequest {
	r.pulpImporter = &pulpImporter
	return r
}

func (r ApiImportersCorePulpUpdateRequest) Execute() (PulpImporterResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportersCorePulpUpdateExecute(r)
}

/*
 * ImportersCorePulpUpdate Update a pulp importer
 * ViewSet for PulpImporters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this pulp importer.
 * @return ApiImportersCorePulpUpdateRequest
 */
func (a *ImportersPulpApiService) ImportersCorePulpUpdate(ctx _context.Context, pulpId string) ApiImportersCorePulpUpdateRequest {
	return ApiImportersCorePulpUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return PulpImporterResponse
 */
func (a *ImportersPulpApiService) ImportersCorePulpUpdateExecute(r ApiImportersCorePulpUpdateRequest) (PulpImporterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpApiService.ImportersCorePulpUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pulpImporter == nil {
		return localVarReturnValue, nil, reportError("pulpImporter is required and must be specified")
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
	localVarPostBody = r.pulpImporter
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
