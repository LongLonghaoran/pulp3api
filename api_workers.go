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

// WorkersApiService WorkersApi service
type WorkersApiService service

type ApiWorkersListRequest struct {
	ctx _context.Context
	ApiService *WorkersApiService
	lastHeartbeat *time.Time
	lastHeartbeatGt *time.Time
	lastHeartbeatGte *time.Time
	lastHeartbeatLt *time.Time
	lastHeartbeatLte *time.Time
	lastHeartbeatRange *[]time.Time
	limit *int32
	missing *bool
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	online *bool
	ordering *[]string
	fields *string
	excludeFields *string
}

func (r ApiWorkersListRequest) LastHeartbeat(lastHeartbeat time.Time) ApiWorkersListRequest {
	r.lastHeartbeat = &lastHeartbeat
	return r
}
func (r ApiWorkersListRequest) LastHeartbeatGt(lastHeartbeatGt time.Time) ApiWorkersListRequest {
	r.lastHeartbeatGt = &lastHeartbeatGt
	return r
}
func (r ApiWorkersListRequest) LastHeartbeatGte(lastHeartbeatGte time.Time) ApiWorkersListRequest {
	r.lastHeartbeatGte = &lastHeartbeatGte
	return r
}
func (r ApiWorkersListRequest) LastHeartbeatLt(lastHeartbeatLt time.Time) ApiWorkersListRequest {
	r.lastHeartbeatLt = &lastHeartbeatLt
	return r
}
func (r ApiWorkersListRequest) LastHeartbeatLte(lastHeartbeatLte time.Time) ApiWorkersListRequest {
	r.lastHeartbeatLte = &lastHeartbeatLte
	return r
}
func (r ApiWorkersListRequest) LastHeartbeatRange(lastHeartbeatRange []time.Time) ApiWorkersListRequest {
	r.lastHeartbeatRange = &lastHeartbeatRange
	return r
}
func (r ApiWorkersListRequest) Limit(limit int32) ApiWorkersListRequest {
	r.limit = &limit
	return r
}
func (r ApiWorkersListRequest) Missing(missing bool) ApiWorkersListRequest {
	r.missing = &missing
	return r
}
func (r ApiWorkersListRequest) Name(name string) ApiWorkersListRequest {
	r.name = &name
	return r
}
func (r ApiWorkersListRequest) NameContains(nameContains string) ApiWorkersListRequest {
	r.nameContains = &nameContains
	return r
}
func (r ApiWorkersListRequest) NameIcontains(nameIcontains string) ApiWorkersListRequest {
	r.nameIcontains = &nameIcontains
	return r
}
func (r ApiWorkersListRequest) NameIn(nameIn []string) ApiWorkersListRequest {
	r.nameIn = &nameIn
	return r
}
func (r ApiWorkersListRequest) NameStartswith(nameStartswith string) ApiWorkersListRequest {
	r.nameStartswith = &nameStartswith
	return r
}
func (r ApiWorkersListRequest) Offset(offset int32) ApiWorkersListRequest {
	r.offset = &offset
	return r
}
func (r ApiWorkersListRequest) Online(online bool) ApiWorkersListRequest {
	r.online = &online
	return r
}
func (r ApiWorkersListRequest) Ordering(ordering []string) ApiWorkersListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiWorkersListRequest) Fields(fields string) ApiWorkersListRequest {
	r.fields = &fields
	return r
}
func (r ApiWorkersListRequest) ExcludeFields(excludeFields string) ApiWorkersListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiWorkersListRequest) Execute() (PaginatedWorkerResponseList, *_nethttp.Response, error) {
	return r.ApiService.WorkersListExecute(r)
}

/*
 * WorkersList List workers
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
 * @return ApiWorkersListRequest
 */
func (a *WorkersApiService) WorkersList(ctx _context.Context) ApiWorkersListRequest {
	return ApiWorkersListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedWorkerResponseList
 */
func (a *WorkersApiService) WorkersListExecute(r ApiWorkersListRequest) (PaginatedWorkerResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedWorkerResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkersApiService.WorkersList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/workers/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.lastHeartbeat != nil {
		localVarQueryParams.Add("last_heartbeat", parameterToString(*r.lastHeartbeat, ""))
	}
	if r.lastHeartbeatGt != nil {
		localVarQueryParams.Add("last_heartbeat__gt", parameterToString(*r.lastHeartbeatGt, ""))
	}
	if r.lastHeartbeatGte != nil {
		localVarQueryParams.Add("last_heartbeat__gte", parameterToString(*r.lastHeartbeatGte, ""))
	}
	if r.lastHeartbeatLt != nil {
		localVarQueryParams.Add("last_heartbeat__lt", parameterToString(*r.lastHeartbeatLt, ""))
	}
	if r.lastHeartbeatLte != nil {
		localVarQueryParams.Add("last_heartbeat__lte", parameterToString(*r.lastHeartbeatLte, ""))
	}
	if r.lastHeartbeatRange != nil {
		localVarQueryParams.Add("last_heartbeat__range", parameterToString(*r.lastHeartbeatRange, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.missing != nil {
		localVarQueryParams.Add("missing", parameterToString(*r.missing, ""))
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
	if r.online != nil {
		localVarQueryParams.Add("online", parameterToString(*r.online, ""))
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

type ApiWorkersReadRequest struct {
	ctx _context.Context
	ApiService *WorkersApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiWorkersReadRequest) Fields(fields string) ApiWorkersReadRequest {
	r.fields = &fields
	return r
}
func (r ApiWorkersReadRequest) ExcludeFields(excludeFields string) ApiWorkersReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiWorkersReadRequest) Execute() (WorkerResponse, *_nethttp.Response, error) {
	return r.ApiService.WorkersReadExecute(r)
}

/*
 * WorkersRead Inspect a worker
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
 * @param pulpId A UUID string identifying this worker.
 * @return ApiWorkersReadRequest
 */
func (a *WorkersApiService) WorkersRead(ctx _context.Context, pulpId string) ApiWorkersReadRequest {
	return ApiWorkersReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return WorkerResponse
 */
func (a *WorkersApiService) WorkersReadExecute(r ApiWorkersReadRequest) (WorkerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkersApiService.WorkersRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/workers/{pulp_id}/"
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
