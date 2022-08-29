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
	"time"
)


// RepositoriesContainerVersionsApiService RepositoriesContainerVersionsApi service
type RepositoriesContainerVersionsApiService service

type ApiRepositoriesContainerContainerVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerVersionsApiService
	number int32
	repositoryPk string
}

func (r ApiRepositoriesContainerContainerVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerVersionsDeleteExecute(r)
}

/*
RepositoriesContainerContainerVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesContainerContainerVersionsDeleteRequest
*/
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsDelete(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesContainerContainerVersionsDeleteRequest {
	return ApiRepositoriesContainerContainerVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsDeleteExecute(r ApiRepositoriesContainerContainerVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerVersionsApiService.RepositoriesContainerContainerVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/container/container/{repository_pk}/versions/{number}/"
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository_pk"+"}", url.PathEscape(parameterToString(r.repositoryPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number < 0 {
		return localVarReturnValue, nil, reportError("number must be greater than 0")
	}
	if r.number > 2147483647 {
		return localVarReturnValue, nil, reportError("number must be less than 2147483647")
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

type ApiRepositoriesContainerContainerVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerVersionsApiService
	repositoryPk string
	content *string
	contentIn *string
	limit *int32
	number *int32
	numberGt *int32
	numberGte *int32
	numberLt *int32
	numberLte *int32
	numberRange *[]int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	fields *string
	excludeFields *string
}

// Content Unit referenced by HREF
func (r ApiRepositoriesContainerContainerVersionsListRequest) Content(content string) ApiRepositoriesContainerContainerVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r ApiRepositoriesContainerContainerVersionsListRequest) ContentIn(contentIn string) ApiRepositoriesContainerContainerVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r ApiRepositoriesContainerContainerVersionsListRequest) Limit(limit int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.limit = &limit
	return r
}

func (r ApiRepositoriesContainerContainerVersionsListRequest) Number(number int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r ApiRepositoriesContainerContainerVersionsListRequest) NumberGt(numberGt int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r ApiRepositoriesContainerContainerVersionsListRequest) NumberGte(numberGte int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r ApiRepositoriesContainerContainerVersionsListRequest) NumberLt(numberLt int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r ApiRepositoriesContainerContainerVersionsListRequest) NumberLte(numberLte int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r ApiRepositoriesContainerContainerVersionsListRequest) NumberRange(numberRange []int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r ApiRepositoriesContainerContainerVersionsListRequest) Offset(offset int32) ApiRepositoriesContainerContainerVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiRepositoriesContainerContainerVersionsListRequest) Ordering(ordering []string) ApiRepositoriesContainerContainerVersionsListRequest {
	r.ordering = &ordering
	return r
}

// ISO 8601 formatted dates are supported
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreated(pulpCreated time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r ApiRepositoriesContainerContainerVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) ApiRepositoriesContainerContainerVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r ApiRepositoriesContainerContainerVersionsListRequest) Fields(fields string) ApiRepositoriesContainerContainerVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiRepositoriesContainerContainerVersionsListRequest) ExcludeFields(excludeFields string) ApiRepositoriesContainerContainerVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRepositoriesContainerContainerVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerVersionsListExecute(r)
}

/*
RepositoriesContainerContainerVersionsList List repository versions

ContainerRepositoryVersion represents a single container repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param repositoryPk
 @return ApiRepositoriesContainerContainerVersionsListRequest
*/
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsList(ctx context.Context, repositoryPk string) ApiRepositoriesContainerContainerVersionsListRequest {
	return ApiRepositoriesContainerContainerVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsListExecute(r ApiRepositoriesContainerContainerVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerVersionsApiService.RepositoriesContainerContainerVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/container/container/{repository_pk}/versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"repository_pk"+"}", url.PathEscape(parameterToString(r.repositoryPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		localVarQueryParams.Add("content", parameterToString(*r.content, ""))
	}
	if r.contentIn != nil {
		localVarQueryParams.Add("content__in", parameterToString(*r.contentIn, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.number != nil {
		localVarQueryParams.Add("number", parameterToString(*r.number, ""))
	}
	if r.numberGt != nil {
		localVarQueryParams.Add("number__gt", parameterToString(*r.numberGt, ""))
	}
	if r.numberGte != nil {
		localVarQueryParams.Add("number__gte", parameterToString(*r.numberGte, ""))
	}
	if r.numberLt != nil {
		localVarQueryParams.Add("number__lt", parameterToString(*r.numberLt, ""))
	}
	if r.numberLte != nil {
		localVarQueryParams.Add("number__lte", parameterToString(*r.numberLte, ""))
	}
	if r.numberRange != nil {
		localVarQueryParams.Add("number__range", parameterToString(*r.numberRange, "csv"))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
	}
	if r.pulpCreated != nil {
		localVarQueryParams.Add("pulp_created", parameterToString(*r.pulpCreated, ""))
	}
	if r.pulpCreatedGt != nil {
		localVarQueryParams.Add("pulp_created__gt", parameterToString(*r.pulpCreatedGt, ""))
	}
	if r.pulpCreatedGte != nil {
		localVarQueryParams.Add("pulp_created__gte", parameterToString(*r.pulpCreatedGte, ""))
	}
	if r.pulpCreatedLt != nil {
		localVarQueryParams.Add("pulp_created__lt", parameterToString(*r.pulpCreatedLt, ""))
	}
	if r.pulpCreatedLte != nil {
		localVarQueryParams.Add("pulp_created__lte", parameterToString(*r.pulpCreatedLte, ""))
	}
	if r.pulpCreatedRange != nil {
		localVarQueryParams.Add("pulp_created__range", parameterToString(*r.pulpCreatedRange, "csv"))
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

type ApiRepositoriesContainerContainerVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerVersionsApiService
	number int32
	repositoryPk string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiRepositoriesContainerContainerVersionsReadRequest) Fields(fields string) ApiRepositoriesContainerContainerVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiRepositoriesContainerContainerVersionsReadRequest) ExcludeFields(excludeFields string) ApiRepositoriesContainerContainerVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRepositoriesContainerContainerVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerVersionsReadExecute(r)
}

/*
RepositoriesContainerContainerVersionsRead Inspect a repository version

ContainerRepositoryVersion represents a single container repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesContainerContainerVersionsReadRequest
*/
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsRead(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesContainerContainerVersionsReadRequest {
	return ApiRepositoriesContainerContainerVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsReadExecute(r ApiRepositoriesContainerContainerVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerVersionsApiService.RepositoriesContainerContainerVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/container/container/{repository_pk}/versions/{number}/"
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository_pk"+"}", url.PathEscape(parameterToString(r.repositoryPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number < 0 {
		return localVarReturnValue, nil, reportError("number must be greater than 0")
	}
	if r.number > 2147483647 {
		return localVarReturnValue, nil, reportError("number must be less than 2147483647")
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

type ApiRepositoriesContainerContainerVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesContainerVersionsApiService
	number int32
	repositoryPk string
	repair *Repair
}

func (r ApiRepositoriesContainerContainerVersionsRepairRequest) Repair(repair Repair) ApiRepositoriesContainerContainerVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r ApiRepositoriesContainerContainerVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerVersionsRepairExecute(r)
}

/*
RepositoriesContainerContainerVersionsRepair Method for RepositoriesContainerContainerVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesContainerContainerVersionsRepairRequest
*/
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsRepair(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesContainerContainerVersionsRepairRequest {
	return ApiRepositoriesContainerContainerVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesContainerVersionsApiService) RepositoriesContainerContainerVersionsRepairExecute(r ApiRepositoriesContainerContainerVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesContainerVersionsApiService.RepositoriesContainerContainerVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/container/container/{repository_pk}/versions/{number}/repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterToString(r.number, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository_pk"+"}", url.PathEscape(parameterToString(r.repositoryPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number < 0 {
		return localVarReturnValue, nil, reportError("number must be greater than 0")
	}
	if r.number > 2147483647 {
		return localVarReturnValue, nil, reportError("number must be less than 2147483647")
	}
	if r.repair == nil {
		return localVarReturnValue, nil, reportError("repair is required and must be specified")
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
	localVarPostBody = r.repair
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
