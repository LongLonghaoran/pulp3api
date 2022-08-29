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


// DistributionsAptApiService DistributionsAptApi service
type DistributionsAptApiService service

type ApiDistributionsDebAptCreateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistribution *DebAptDistribution
}

func (r ApiDistributionsDebAptCreateRequest) DebAptDistribution(debAptDistribution DebAptDistribution) ApiDistributionsDebAptCreateRequest {
	r.debAptDistribution = &debAptDistribution
	return r
}

func (r ApiDistributionsDebAptCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptCreateExecute(r)
}

/*
DistributionsDebAptCreate Create an apt distribution

Trigger an asynchronous create task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDistributionsDebAptCreateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptCreate(ctx context.Context) ApiDistributionsDebAptCreateRequest {
	return ApiDistributionsDebAptCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptCreateExecute(r ApiDistributionsDebAptCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptDistribution == nil {
		return localVarReturnValue, nil, reportError("debAptDistribution is required and must be specified")
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
	localVarPostBody = r.debAptDistribution
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

type ApiDistributionsDebAptDeleteRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	pulpId string
}

func (r ApiDistributionsDebAptDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptDeleteExecute(r)
}

/*
DistributionsDebAptDelete Delete an apt distribution

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this apt distribution.
 @return ApiDistributionsDebAptDeleteRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptDelete(ctx context.Context, pulpId string) ApiDistributionsDebAptDeleteRequest {
	return ApiDistributionsDebAptDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptDeleteExecute(r ApiDistributionsDebAptDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/{pulp_id}/"
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

type ApiDistributionsDebAptListRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	fields *string
	excludeFields *string
}

func (r ApiDistributionsDebAptListRequest) BasePath(basePath string) ApiDistributionsDebAptListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r ApiDistributionsDebAptListRequest) BasePathContains(basePathContains string) ApiDistributionsDebAptListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r ApiDistributionsDebAptListRequest) BasePathIcontains(basePathIcontains string) ApiDistributionsDebAptListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r ApiDistributionsDebAptListRequest) BasePathIn(basePathIn []string) ApiDistributionsDebAptListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r ApiDistributionsDebAptListRequest) Limit(limit int32) ApiDistributionsDebAptListRequest {
	r.limit = &limit
	return r
}

func (r ApiDistributionsDebAptListRequest) Name(name string) ApiDistributionsDebAptListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ApiDistributionsDebAptListRequest) NameContains(nameContains string) ApiDistributionsDebAptListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ApiDistributionsDebAptListRequest) NameIcontains(nameIcontains string) ApiDistributionsDebAptListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ApiDistributionsDebAptListRequest) NameIn(nameIn []string) ApiDistributionsDebAptListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ApiDistributionsDebAptListRequest) NameStartswith(nameStartswith string) ApiDistributionsDebAptListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ApiDistributionsDebAptListRequest) Offset(offset int32) ApiDistributionsDebAptListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiDistributionsDebAptListRequest) Ordering(ordering []string) ApiDistributionsDebAptListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r ApiDistributionsDebAptListRequest) PulpLabelSelect(pulpLabelSelect string) ApiDistributionsDebAptListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// A list of fields to include in the response.
func (r ApiDistributionsDebAptListRequest) Fields(fields string) ApiDistributionsDebAptListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiDistributionsDebAptListRequest) ExcludeFields(excludeFields string) ApiDistributionsDebAptListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiDistributionsDebAptListRequest) Execute() (*PaginateddebAptDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsDebAptListExecute(r)
}

/*
DistributionsDebAptList List apt distributions

An AptDistribution is just an AptPublication made available via the content app.

Creating an AptDistribution is a comparatively quick action. This way Pulp users may take as
much time as is needed to prepare a VerbatimPublication or AptPublication, and then control the
exact moment when that publication is made available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDistributionsDebAptListRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptList(ctx context.Context) ApiDistributionsDebAptListRequest {
	return ApiDistributionsDebAptListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebAptDistributionResponseList
func (a *DistributionsAptApiService) DistributionsDebAptListExecute(r ApiDistributionsDebAptListRequest) (*PaginateddebAptDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebAptDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		localVarQueryParams.Add("base_path", parameterToString(*r.basePath, ""))
	}
	if r.basePathContains != nil {
		localVarQueryParams.Add("base_path__contains", parameterToString(*r.basePathContains, ""))
	}
	if r.basePathIcontains != nil {
		localVarQueryParams.Add("base_path__icontains", parameterToString(*r.basePathIcontains, ""))
	}
	if r.basePathIn != nil {
		localVarQueryParams.Add("base_path__in", parameterToString(*r.basePathIn, "csv"))
	}
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
	if r.pulpLabelSelect != nil {
		localVarQueryParams.Add("pulp_label_select", parameterToString(*r.pulpLabelSelect, ""))
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

type ApiDistributionsDebAptPartialUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	pulpId string
	patcheddebAptDistribution *PatcheddebAptDistribution
}

func (r ApiDistributionsDebAptPartialUpdateRequest) PatcheddebAptDistribution(patcheddebAptDistribution PatcheddebAptDistribution) ApiDistributionsDebAptPartialUpdateRequest {
	r.patcheddebAptDistribution = &patcheddebAptDistribution
	return r
}

func (r ApiDistributionsDebAptPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptPartialUpdateExecute(r)
}

/*
DistributionsDebAptPartialUpdate Update an apt distribution

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this apt distribution.
 @return ApiDistributionsDebAptPartialUpdateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptPartialUpdate(ctx context.Context, pulpId string) ApiDistributionsDebAptPartialUpdateRequest {
	return ApiDistributionsDebAptPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptPartialUpdateExecute(r ApiDistributionsDebAptPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patcheddebAptDistribution == nil {
		return localVarReturnValue, nil, reportError("patcheddebAptDistribution is required and must be specified")
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
	localVarPostBody = r.patcheddebAptDistribution
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

type ApiDistributionsDebAptReadRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	pulpId string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiDistributionsDebAptReadRequest) Fields(fields string) ApiDistributionsDebAptReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiDistributionsDebAptReadRequest) ExcludeFields(excludeFields string) ApiDistributionsDebAptReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiDistributionsDebAptReadRequest) Execute() (*DebAptDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptReadExecute(r)
}

/*
DistributionsDebAptRead Inspect an apt distribution

An AptDistribution is just an AptPublication made available via the content app.

Creating an AptDistribution is a comparatively quick action. This way Pulp users may take as
much time as is needed to prepare a VerbatimPublication or AptPublication, and then control the
exact moment when that publication is made available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this apt distribution.
 @return ApiDistributionsDebAptReadRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptRead(ctx context.Context, pulpId string) ApiDistributionsDebAptReadRequest {
	return ApiDistributionsDebAptReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return DebAptDistributionResponse
func (a *DistributionsAptApiService) DistributionsDebAptReadExecute(r ApiDistributionsDebAptReadRequest) (*DebAptDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/{pulp_id}/"
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

type ApiDistributionsDebAptUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	pulpId string
	debAptDistribution *DebAptDistribution
}

func (r ApiDistributionsDebAptUpdateRequest) DebAptDistribution(debAptDistribution DebAptDistribution) ApiDistributionsDebAptUpdateRequest {
	r.debAptDistribution = &debAptDistribution
	return r
}

func (r ApiDistributionsDebAptUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptUpdateExecute(r)
}

/*
DistributionsDebAptUpdate Update an apt distribution

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this apt distribution.
 @return ApiDistributionsDebAptUpdateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptUpdate(ctx context.Context, pulpId string) ApiDistributionsDebAptUpdateRequest {
	return ApiDistributionsDebAptUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptUpdateExecute(r ApiDistributionsDebAptUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", url.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptDistribution == nil {
		return localVarReturnValue, nil, reportError("debAptDistribution is required and must be specified")
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
	localVarPostBody = r.debAptDistribution
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
