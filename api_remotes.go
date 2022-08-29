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
	"time"
)

// Linger please
var (
	_ _context.Context
)

// RemotesApiService RemotesApi service
type RemotesApiService service

type ApiRemotesCreateRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	remote *Remote
}

func (r ApiRemotesCreateRequest) Remote(remote Remote) ApiRemotesCreateRequest {
	r.remote = &remote
	return r
}

func (r ApiRemotesCreateRequest) Execute() (RemoteResponse, *_nethttp.Response, error) {
	return r.ApiService.RemotesCreateExecute(r)
}

/*
 * RemotesCreate Create a remote
 * A customized named ModelViewSet that knows how to register itself with the Pulp API router.

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
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRemotesCreateRequest
 */
func (a *RemotesApiService) RemotesCreate(ctx _context.Context) ApiRemotesCreateRequest {
	return ApiRemotesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RemoteResponse
 */
func (a *RemotesApiService) RemotesCreateExecute(r ApiRemotesCreateRequest) (RemoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.remote == nil {
		return localVarReturnValue, nil, reportError("remote is required and must be specified")
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
	localVarPostBody = r.remote
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

type ApiRemotesDeleteRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	pulpId string
}


func (r ApiRemotesDeleteRequest) Execute() (AsyncOperationResponse, *_nethttp.Response, error) {
	return r.ApiService.RemotesDeleteExecute(r)
}

/*
 * RemotesDelete Delete a remote
 * Trigger an asynchronous delete task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this remote.
 * @return ApiRemotesDeleteRequest
 */
func (a *RemotesApiService) RemotesDelete(ctx _context.Context, pulpId string) ApiRemotesDeleteRequest {
	return ApiRemotesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return AsyncOperationResponse
 */
func (a *RemotesApiService) RemotesDeleteExecute(r ApiRemotesDeleteRequest) (AsyncOperationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesDelete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/{pulp_id}/"
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

type ApiRemotesListRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	pulpLastUpdated *time.Time
	pulpLastUpdatedGt *time.Time
	pulpLastUpdatedGte *time.Time
	pulpLastUpdatedLt *time.Time
	pulpLastUpdatedLte *time.Time
	pulpLastUpdatedRange *[]time.Time
	fields *string
	excludeFields *string
}

func (r ApiRemotesListRequest) Limit(limit int32) ApiRemotesListRequest {
	r.limit = &limit
	return r
}
func (r ApiRemotesListRequest) Name(name string) ApiRemotesListRequest {
	r.name = &name
	return r
}
func (r ApiRemotesListRequest) NameContains(nameContains string) ApiRemotesListRequest {
	r.nameContains = &nameContains
	return r
}
func (r ApiRemotesListRequest) NameIcontains(nameIcontains string) ApiRemotesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}
func (r ApiRemotesListRequest) NameIn(nameIn []string) ApiRemotesListRequest {
	r.nameIn = &nameIn
	return r
}
func (r ApiRemotesListRequest) NameStartswith(nameStartswith string) ApiRemotesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}
func (r ApiRemotesListRequest) Offset(offset int32) ApiRemotesListRequest {
	r.offset = &offset
	return r
}
func (r ApiRemotesListRequest) Ordering(ordering []string) ApiRemotesListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiRemotesListRequest) PulpLabelSelect(pulpLabelSelect string) ApiRemotesListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdated(pulpLastUpdated time.Time) ApiRemotesListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) ApiRemotesListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) ApiRemotesListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) ApiRemotesListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) ApiRemotesListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}
func (r ApiRemotesListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) ApiRemotesListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}
func (r ApiRemotesListRequest) Fields(fields string) ApiRemotesListRequest {
	r.fields = &fields
	return r
}
func (r ApiRemotesListRequest) ExcludeFields(excludeFields string) ApiRemotesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRemotesListRequest) Execute() (PaginatedRemoteResponseList, *_nethttp.Response, error) {
	return r.ApiService.RemotesListExecute(r)
}

/*
 * RemotesList List remotes
 * A customized named ModelViewSet that knows how to register itself with the Pulp API router.

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
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRemotesListRequest
 */
func (a *RemotesApiService) RemotesList(ctx _context.Context) ApiRemotesListRequest {
	return ApiRemotesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedRemoteResponseList
 */
func (a *RemotesApiService) RemotesListExecute(r ApiRemotesListRequest) (PaginatedRemoteResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/"

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
	if r.pulpLabelSelect != nil {
		localVarQueryParams.Add("pulp_label_select", parameterToString(*r.pulpLabelSelect, ""))
	}
	if r.pulpLastUpdated != nil {
		localVarQueryParams.Add("pulp_last_updated", parameterToString(*r.pulpLastUpdated, ""))
	}
	if r.pulpLastUpdatedGt != nil {
		localVarQueryParams.Add("pulp_last_updated__gt", parameterToString(*r.pulpLastUpdatedGt, ""))
	}
	if r.pulpLastUpdatedGte != nil {
		localVarQueryParams.Add("pulp_last_updated__gte", parameterToString(*r.pulpLastUpdatedGte, ""))
	}
	if r.pulpLastUpdatedLt != nil {
		localVarQueryParams.Add("pulp_last_updated__lt", parameterToString(*r.pulpLastUpdatedLt, ""))
	}
	if r.pulpLastUpdatedLte != nil {
		localVarQueryParams.Add("pulp_last_updated__lte", parameterToString(*r.pulpLastUpdatedLte, ""))
	}
	if r.pulpLastUpdatedRange != nil {
		localVarQueryParams.Add("pulp_last_updated__range", parameterToString(*r.pulpLastUpdatedRange, "csv"))
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

type ApiRemotesPartialUpdateRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	pulpId string
	patchedRemote *PatchedRemote
}

func (r ApiRemotesPartialUpdateRequest) PatchedRemote(patchedRemote PatchedRemote) ApiRemotesPartialUpdateRequest {
	r.patchedRemote = &patchedRemote
	return r
}

func (r ApiRemotesPartialUpdateRequest) Execute() (AsyncOperationResponse, *_nethttp.Response, error) {
	return r.ApiService.RemotesPartialUpdateExecute(r)
}

/*
 * RemotesPartialUpdate Update a remote
 * Trigger an asynchronous partial update task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this remote.
 * @return ApiRemotesPartialUpdateRequest
 */
func (a *RemotesApiService) RemotesPartialUpdate(ctx _context.Context, pulpId string) ApiRemotesPartialUpdateRequest {
	return ApiRemotesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return AsyncOperationResponse
 */
func (a *RemotesApiService) RemotesPartialUpdateExecute(r ApiRemotesPartialUpdateRequest) (AsyncOperationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.patchedRemote == nil {
		return localVarReturnValue, nil, reportError("patchedRemote is required and must be specified")
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
	localVarPostBody = r.patchedRemote
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

type ApiRemotesReadRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiRemotesReadRequest) Fields(fields string) ApiRemotesReadRequest {
	r.fields = &fields
	return r
}
func (r ApiRemotesReadRequest) ExcludeFields(excludeFields string) ApiRemotesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRemotesReadRequest) Execute() (RemoteResponse, *_nethttp.Response, error) {
	return r.ApiService.RemotesReadExecute(r)
}

/*
 * RemotesRead Inspect a remote
 * A customized named ModelViewSet that knows how to register itself with the Pulp API router.

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
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this remote.
 * @return ApiRemotesReadRequest
 */
func (a *RemotesApiService) RemotesRead(ctx _context.Context, pulpId string) ApiRemotesReadRequest {
	return ApiRemotesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return RemoteResponse
 */
func (a *RemotesApiService) RemotesReadExecute(r ApiRemotesReadRequest) (RemoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/{pulp_id}/"
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

type ApiRemotesUpdateRequest struct {
	ctx _context.Context
	ApiService *RemotesApiService
	pulpId string
	remote *Remote
}

func (r ApiRemotesUpdateRequest) Remote(remote Remote) ApiRemotesUpdateRequest {
	r.remote = &remote
	return r
}

func (r ApiRemotesUpdateRequest) Execute() (AsyncOperationResponse, *_nethttp.Response, error) {
	return r.ApiService.RemotesUpdateExecute(r)
}

/*
 * RemotesUpdate Update a remote
 * Trigger an asynchronous update task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this remote.
 * @return ApiRemotesUpdateRequest
 */
func (a *RemotesApiService) RemotesUpdate(ctx _context.Context, pulpId string) ApiRemotesUpdateRequest {
	return ApiRemotesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return AsyncOperationResponse
 */
func (a *RemotesApiService) RemotesUpdateExecute(r ApiRemotesUpdateRequest) (AsyncOperationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesApiService.RemotesUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.remote == nil {
		return localVarReturnValue, nil, reportError("remote is required and must be specified")
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
	localVarPostBody = r.remote
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
