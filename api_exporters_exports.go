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


// ExportersExportsApiService ExportersExportsApi service
type ExportersExportsApiService service

type ApiExportersExportsCreateRequest struct {
	ctx context.Context
	ApiService *ExportersExportsApiService
	exporterPk string
	export *Export
}

func (r ApiExportersExportsCreateRequest) Export(export Export) ApiExportersExportsCreateRequest {
	r.export = &export
	return r
}

func (r ApiExportersExportsCreateRequest) Execute() (*ExportResponse, *http.Response, error) {
	return r.ApiService.ExportersExportsCreateExecute(r)
}

/*
ExportersExportsCreate Create an export

ViewSet for viewing exports from an Exporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exporterPk
 @return ApiExportersExportsCreateRequest
*/
func (a *ExportersExportsApiService) ExportersExportsCreate(ctx context.Context, exporterPk string) ApiExportersExportsCreateRequest {
	return ApiExportersExportsCreateRequest{
		ApiService: a,
		ctx: ctx,
		exporterPk: exporterPk,
	}
}

// Execute executes the request
//  @return ExportResponse
func (a *ExportersExportsApiService) ExportersExportsCreateExecute(r ApiExportersExportsCreateRequest) (*ExportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersExportsApiService.ExportersExportsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/{exporter_pk}/exports/"
	localVarPath = strings.Replace(localVarPath, "{"+"exporter_pk"+"}", url.PathEscape(parameterToString(r.exporterPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.export == nil {
		return localVarReturnValue, nil, reportError("export is required and must be specified")
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
	localVarPostBody = r.export
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

type ApiExportersExportsDeleteRequest struct {
	ctx context.Context
	ApiService *ExportersExportsApiService
	exporterPk string
	pulpId string
}

func (r ApiExportersExportsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExportersExportsDeleteExecute(r)
}

/*
ExportersExportsDelete Delete an export

ViewSet for viewing exports from an Exporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exporterPk
 @param pulpId A UUID string identifying this export.
 @return ApiExportersExportsDeleteRequest
*/
func (a *ExportersExportsApiService) ExportersExportsDelete(ctx context.Context, exporterPk string, pulpId string) ApiExportersExportsDeleteRequest {
	return ApiExportersExportsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		exporterPk: exporterPk,
		pulpId: pulpId,
	}
}

// Execute executes the request
func (a *ExportersExportsApiService) ExportersExportsDeleteExecute(r ApiExportersExportsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersExportsApiService.ExportersExportsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/{exporter_pk}/exports/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"exporter_pk"+"}", url.PathEscape(parameterToString(r.exporterPk, "")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiExportersExportsListRequest struct {
	ctx context.Context
	ApiService *ExportersExportsApiService
	exporterPk string
	limit *int32
	offset *int32
	fields *string
	excludeFields *string
}

// Number of results to return per page.
func (r ApiExportersExportsListRequest) Limit(limit int32) ApiExportersExportsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiExportersExportsListRequest) Offset(offset int32) ApiExportersExportsListRequest {
	r.offset = &offset
	return r
}

// A list of fields to include in the response.
func (r ApiExportersExportsListRequest) Fields(fields string) ApiExportersExportsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiExportersExportsListRequest) ExcludeFields(excludeFields string) ApiExportersExportsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiExportersExportsListRequest) Execute() (*PaginatedExportResponseList, *http.Response, error) {
	return r.ApiService.ExportersExportsListExecute(r)
}

/*
ExportersExportsList List exports

ViewSet for viewing exports from an Exporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exporterPk
 @return ApiExportersExportsListRequest
*/
func (a *ExportersExportsApiService) ExportersExportsList(ctx context.Context, exporterPk string) ApiExportersExportsListRequest {
	return ApiExportersExportsListRequest{
		ApiService: a,
		ctx: ctx,
		exporterPk: exporterPk,
	}
}

// Execute executes the request
//  @return PaginatedExportResponseList
func (a *ExportersExportsApiService) ExportersExportsListExecute(r ApiExportersExportsListRequest) (*PaginatedExportResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedExportResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersExportsApiService.ExportersExportsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/{exporter_pk}/exports/"
	localVarPath = strings.Replace(localVarPath, "{"+"exporter_pk"+"}", url.PathEscape(parameterToString(r.exporterPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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

type ApiExportersExportsReadRequest struct {
	ctx context.Context
	ApiService *ExportersExportsApiService
	exporterPk string
	pulpId string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiExportersExportsReadRequest) Fields(fields string) ApiExportersExportsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiExportersExportsReadRequest) ExcludeFields(excludeFields string) ApiExportersExportsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiExportersExportsReadRequest) Execute() (*ExportResponse, *http.Response, error) {
	return r.ApiService.ExportersExportsReadExecute(r)
}

/*
ExportersExportsRead Inspect an export

ViewSet for viewing exports from an Exporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exporterPk
 @param pulpId A UUID string identifying this export.
 @return ApiExportersExportsReadRequest
*/
func (a *ExportersExportsApiService) ExportersExportsRead(ctx context.Context, exporterPk string, pulpId string) ApiExportersExportsReadRequest {
	return ApiExportersExportsReadRequest{
		ApiService: a,
		ctx: ctx,
		exporterPk: exporterPk,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return ExportResponse
func (a *ExportersExportsApiService) ExportersExportsReadExecute(r ApiExportersExportsReadRequest) (*ExportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersExportsApiService.ExportersExportsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/{exporter_pk}/exports/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"exporter_pk"+"}", url.PathEscape(parameterToString(r.exporterPk, "")), -1)
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
