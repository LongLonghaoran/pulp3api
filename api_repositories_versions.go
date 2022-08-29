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


// RepositoriesVersionsApiService RepositoriesVersionsApi service
type RepositoriesVersionsApiService service

type ApiRepositoriesVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesVersionsApiService
	number int32
	repositoryPk string
}

func (r ApiRepositoriesVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesVersionsDeleteExecute(r)
}

/*
RepositoriesVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesVersionsDeleteRequest
*/
func (a *RepositoriesVersionsApiService) RepositoriesVersionsDelete(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesVersionsDeleteRequest {
	return ApiRepositoriesVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesVersionsApiService) RepositoriesVersionsDeleteExecute(r ApiRepositoriesVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesVersionsApiService.RepositoriesVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/{repository_pk}/versions/{number}/"
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

type ApiRepositoriesVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesVersionsApiService
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
func (r ApiRepositoriesVersionsListRequest) Content(content string) ApiRepositoriesVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r ApiRepositoriesVersionsListRequest) ContentIn(contentIn string) ApiRepositoriesVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r ApiRepositoriesVersionsListRequest) Limit(limit int32) ApiRepositoriesVersionsListRequest {
	r.limit = &limit
	return r
}

func (r ApiRepositoriesVersionsListRequest) Number(number int32) ApiRepositoriesVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r ApiRepositoriesVersionsListRequest) NumberGt(numberGt int32) ApiRepositoriesVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r ApiRepositoriesVersionsListRequest) NumberGte(numberGte int32) ApiRepositoriesVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r ApiRepositoriesVersionsListRequest) NumberLt(numberLt int32) ApiRepositoriesVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r ApiRepositoriesVersionsListRequest) NumberLte(numberLte int32) ApiRepositoriesVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r ApiRepositoriesVersionsListRequest) NumberRange(numberRange []int32) ApiRepositoriesVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r ApiRepositoriesVersionsListRequest) Offset(offset int32) ApiRepositoriesVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiRepositoriesVersionsListRequest) Ordering(ordering []string) ApiRepositoriesVersionsListRequest {
	r.ordering = &ordering
	return r
}

// ISO 8601 formatted dates are supported
func (r ApiRepositoriesVersionsListRequest) PulpCreated(pulpCreated time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r ApiRepositoriesVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r ApiRepositoriesVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r ApiRepositoriesVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r ApiRepositoriesVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r ApiRepositoriesVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) ApiRepositoriesVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r ApiRepositoriesVersionsListRequest) Fields(fields string) ApiRepositoriesVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiRepositoriesVersionsListRequest) ExcludeFields(excludeFields string) ApiRepositoriesVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRepositoriesVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesVersionsListExecute(r)
}

/*
RepositoriesVersionsList List repository versions

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param repositoryPk
 @return ApiRepositoriesVersionsListRequest
*/
func (a *RepositoriesVersionsApiService) RepositoriesVersionsList(ctx context.Context, repositoryPk string) ApiRepositoriesVersionsListRequest {
	return ApiRepositoriesVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesVersionsApiService) RepositoriesVersionsListExecute(r ApiRepositoriesVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesVersionsApiService.RepositoriesVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/{repository_pk}/versions/"
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

type ApiRepositoriesVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesVersionsApiService
	number int32
	repositoryPk string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiRepositoriesVersionsReadRequest) Fields(fields string) ApiRepositoriesVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiRepositoriesVersionsReadRequest) ExcludeFields(excludeFields string) ApiRepositoriesVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRepositoriesVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesVersionsReadExecute(r)
}

/*
RepositoriesVersionsRead Inspect a repository version

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesVersionsReadRequest
*/
func (a *RepositoriesVersionsApiService) RepositoriesVersionsRead(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesVersionsReadRequest {
	return ApiRepositoriesVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesVersionsApiService) RepositoriesVersionsReadExecute(r ApiRepositoriesVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesVersionsApiService.RepositoriesVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/{repository_pk}/versions/{number}/"
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

type ApiRepositoriesVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesVersionsApiService
	number int32
	repositoryPk string
	repair *Repair
}

func (r ApiRepositoriesVersionsRepairRequest) Repair(repair Repair) ApiRepositoriesVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r ApiRepositoriesVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesVersionsRepairExecute(r)
}

/*
RepositoriesVersionsRepair Method for RepositoriesVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param number
 @param repositoryPk
 @return ApiRepositoriesVersionsRepairRequest
*/
func (a *RepositoriesVersionsApiService) RepositoriesVersionsRepair(ctx context.Context, number int32, repositoryPk string) ApiRepositoriesVersionsRepairRequest {
	return ApiRepositoriesVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		number: number,
		repositoryPk: repositoryPk,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesVersionsApiService) RepositoriesVersionsRepairExecute(r ApiRepositoriesVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesVersionsApiService.RepositoriesVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/{repository_pk}/versions/{number}/repair/"
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
