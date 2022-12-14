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


// ExportersPulpApiService ExportersPulpApi service
type ExportersPulpApiService service

type ApiExportersCorePulpCreateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
	pulpExporter *PulpExporter
}

func (r ApiExportersCorePulpCreateRequest) PulpExporter(pulpExporter PulpExporter) ApiExportersCorePulpCreateRequest {
	r.pulpExporter = &pulpExporter
	return r
}

func (r ApiExportersCorePulpCreateRequest) Execute() (*PulpExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpCreateExecute(r)
}

/*
ExportersCorePulpCreate Create a pulp exporter

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportersCorePulpCreateRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpCreate(ctx context.Context) ApiExportersCorePulpCreateRequest {
	return ApiExportersCorePulpCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PulpExporterResponse
func (a *ExportersPulpApiService) ExportersCorePulpCreateExecute(r ApiExportersCorePulpCreateRequest) (*PulpExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpExporter == nil {
		return localVarReturnValue, nil, reportError("pulpExporter is required and must be specified")
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
	localVarPostBody = r.pulpExporter
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

type ApiExportersCorePulpDeleteRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
	pulpId string
}

func (r ApiExportersCorePulpDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpDeleteExecute(r)
}

/*
ExportersCorePulpDelete Delete a pulp exporter

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this pulp exporter.
 @return ApiExportersCorePulpDeleteRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpDelete(ctx context.Context, pulpId string) ApiExportersCorePulpDeleteRequest {
	return ApiExportersCorePulpDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpApiService) ExportersCorePulpDeleteExecute(r ApiExportersCorePulpDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiExportersCorePulpListRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
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

// Number of results to return per page.
func (r ApiExportersCorePulpListRequest) Limit(limit int32) ApiExportersCorePulpListRequest {
	r.limit = &limit
	return r
}

func (r ApiExportersCorePulpListRequest) Name(name string) ApiExportersCorePulpListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ApiExportersCorePulpListRequest) NameContains(nameContains string) ApiExportersCorePulpListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ApiExportersCorePulpListRequest) NameIcontains(nameIcontains string) ApiExportersCorePulpListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ApiExportersCorePulpListRequest) NameIn(nameIn []string) ApiExportersCorePulpListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ApiExportersCorePulpListRequest) NameStartswith(nameStartswith string) ApiExportersCorePulpListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ApiExportersCorePulpListRequest) Offset(offset int32) ApiExportersCorePulpListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiExportersCorePulpListRequest) Ordering(ordering []string) ApiExportersCorePulpListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r ApiExportersCorePulpListRequest) Fields(fields string) ApiExportersCorePulpListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiExportersCorePulpListRequest) ExcludeFields(excludeFields string) ApiExportersCorePulpListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiExportersCorePulpListRequest) Execute() (*PaginatedPulpExporterResponseList, *http.Response, error) {
	return r.ApiService.ExportersCorePulpListExecute(r)
}

/*
ExportersCorePulpList List pulp exporters

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportersCorePulpListRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpList(ctx context.Context) ApiExportersCorePulpListRequest {
	return ApiExportersCorePulpListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedPulpExporterResponseList
func (a *ExportersPulpApiService) ExportersCorePulpListExecute(r ApiExportersCorePulpListRequest) (*PaginatedPulpExporterResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPulpExporterResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiExportersCorePulpPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
	pulpId string
	patchedPulpExporter *PatchedPulpExporter
}

func (r ApiExportersCorePulpPartialUpdateRequest) PatchedPulpExporter(patchedPulpExporter PatchedPulpExporter) ApiExportersCorePulpPartialUpdateRequest {
	r.patchedPulpExporter = &patchedPulpExporter
	return r
}

func (r ApiExportersCorePulpPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpPartialUpdateExecute(r)
}

/*
ExportersCorePulpPartialUpdate Update a pulp exporter

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this pulp exporter.
 @return ApiExportersCorePulpPartialUpdateRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpPartialUpdate(ctx context.Context, pulpId string) ApiExportersCorePulpPartialUpdateRequest {
	return ApiExportersCorePulpPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpApiService) ExportersCorePulpPartialUpdateExecute(r ApiExportersCorePulpPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedPulpExporter == nil {
		return localVarReturnValue, nil, reportError("patchedPulpExporter is required and must be specified")
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
	localVarPostBody = r.patchedPulpExporter
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

type ApiExportersCorePulpReadRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
	pulpId string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiExportersCorePulpReadRequest) Fields(fields string) ApiExportersCorePulpReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiExportersCorePulpReadRequest) ExcludeFields(excludeFields string) ApiExportersCorePulpReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiExportersCorePulpReadRequest) Execute() (*PulpExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpReadExecute(r)
}

/*
ExportersCorePulpRead Inspect a pulp exporter

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this pulp exporter.
 @return ApiExportersCorePulpReadRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpRead(ctx context.Context, pulpId string) ApiExportersCorePulpReadRequest {
	return ApiExportersCorePulpReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return PulpExporterResponse
func (a *ExportersPulpApiService) ExportersCorePulpReadExecute(r ApiExportersCorePulpReadRequest) (*PulpExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/{pulp_id}/"
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

type ApiExportersCorePulpUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpApiService
	pulpId string
	pulpExporter *PulpExporter
}

func (r ApiExportersCorePulpUpdateRequest) PulpExporter(pulpExporter PulpExporter) ApiExportersCorePulpUpdateRequest {
	r.pulpExporter = &pulpExporter
	return r
}

func (r ApiExportersCorePulpUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpUpdateExecute(r)
}

/*
ExportersCorePulpUpdate Update a pulp exporter

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this pulp exporter.
 @return ApiExportersCorePulpUpdateRequest
*/
func (a *ExportersPulpApiService) ExportersCorePulpUpdate(ctx context.Context, pulpId string) ApiExportersCorePulpUpdateRequest {
	return ApiExportersCorePulpUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpApiService) ExportersCorePulpUpdateExecute(r ApiExportersCorePulpUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpApiService.ExportersCorePulpUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/pulp/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpExporter == nil {
		return localVarReturnValue, nil, reportError("pulpExporter is required and must be specified")
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
	localVarPostBody = r.pulpExporter
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
