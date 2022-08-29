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

// TaskSchedulesApiService TaskSchedulesApi service
type TaskSchedulesApiService service

type ApiTaskSchedulesAddRoleRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	pulpId string
	nestedRole *NestedRole
}

func (r ApiTaskSchedulesAddRoleRequest) NestedRole(nestedRole NestedRole) ApiTaskSchedulesAddRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r ApiTaskSchedulesAddRoleRequest) Execute() (NestedRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesAddRoleExecute(r)
}

/*
 * TaskSchedulesAddRole Method for TaskSchedulesAddRole
 * Add a role for this object to users/groups.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this task schedule.
 * @return ApiTaskSchedulesAddRoleRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesAddRole(ctx _context.Context, pulpId string) ApiTaskSchedulesAddRoleRequest {
	return ApiTaskSchedulesAddRoleRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return NestedRoleResponse
 */
func (a *TaskSchedulesApiService) TaskSchedulesAddRoleExecute(r ApiTaskSchedulesAddRoleRequest) (NestedRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesAddRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/{pulp_id}/add_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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

type ApiTaskSchedulesListRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	limit *int32
	name *string
	nameContains *string
	offset *int32
	ordering *[]string
	taskName *string
	taskNameContains *string
	fields *string
	excludeFields *string
}

func (r ApiTaskSchedulesListRequest) Limit(limit int32) ApiTaskSchedulesListRequest {
	r.limit = &limit
	return r
}
func (r ApiTaskSchedulesListRequest) Name(name string) ApiTaskSchedulesListRequest {
	r.name = &name
	return r
}
func (r ApiTaskSchedulesListRequest) NameContains(nameContains string) ApiTaskSchedulesListRequest {
	r.nameContains = &nameContains
	return r
}
func (r ApiTaskSchedulesListRequest) Offset(offset int32) ApiTaskSchedulesListRequest {
	r.offset = &offset
	return r
}
func (r ApiTaskSchedulesListRequest) Ordering(ordering []string) ApiTaskSchedulesListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiTaskSchedulesListRequest) TaskName(taskName string) ApiTaskSchedulesListRequest {
	r.taskName = &taskName
	return r
}
func (r ApiTaskSchedulesListRequest) TaskNameContains(taskNameContains string) ApiTaskSchedulesListRequest {
	r.taskNameContains = &taskNameContains
	return r
}
func (r ApiTaskSchedulesListRequest) Fields(fields string) ApiTaskSchedulesListRequest {
	r.fields = &fields
	return r
}
func (r ApiTaskSchedulesListRequest) ExcludeFields(excludeFields string) ApiTaskSchedulesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiTaskSchedulesListRequest) Execute() (PaginatedTaskScheduleResponseList, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesListExecute(r)
}

/*
 * TaskSchedulesList List task schedules
 * ViewSet to monitor task schedules.

NOTE: This feature is in tech-preview and may change in backwards incompatible ways.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTaskSchedulesListRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesList(ctx _context.Context) ApiTaskSchedulesListRequest {
	return ApiTaskSchedulesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedTaskScheduleResponseList
 */
func (a *TaskSchedulesApiService) TaskSchedulesListExecute(r ApiTaskSchedulesListRequest) (PaginatedTaskScheduleResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedTaskScheduleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/"

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
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
	}
	if r.taskName != nil {
		localVarQueryParams.Add("task_name", parameterToString(*r.taskName, ""))
	}
	if r.taskNameContains != nil {
		localVarQueryParams.Add("task_name__contains", parameterToString(*r.taskNameContains, ""))
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

type ApiTaskSchedulesListRolesRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiTaskSchedulesListRolesRequest) Fields(fields string) ApiTaskSchedulesListRolesRequest {
	r.fields = &fields
	return r
}
func (r ApiTaskSchedulesListRolesRequest) ExcludeFields(excludeFields string) ApiTaskSchedulesListRolesRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiTaskSchedulesListRolesRequest) Execute() (ObjectRolesResponse, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesListRolesExecute(r)
}

/*
 * TaskSchedulesListRoles Method for TaskSchedulesListRoles
 * List roles assigned to this object.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this task schedule.
 * @return ApiTaskSchedulesListRolesRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesListRoles(ctx _context.Context, pulpId string) ApiTaskSchedulesListRolesRequest {
	return ApiTaskSchedulesListRolesRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return ObjectRolesResponse
 */
func (a *TaskSchedulesApiService) TaskSchedulesListRolesExecute(r ApiTaskSchedulesListRolesRequest) (ObjectRolesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesListRoles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/{pulp_id}/list_roles/"
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

type ApiTaskSchedulesMyPermissionsRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiTaskSchedulesMyPermissionsRequest) Fields(fields string) ApiTaskSchedulesMyPermissionsRequest {
	r.fields = &fields
	return r
}
func (r ApiTaskSchedulesMyPermissionsRequest) ExcludeFields(excludeFields string) ApiTaskSchedulesMyPermissionsRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiTaskSchedulesMyPermissionsRequest) Execute() (MyPermissionsResponse, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesMyPermissionsExecute(r)
}

/*
 * TaskSchedulesMyPermissions Method for TaskSchedulesMyPermissions
 * List permissions available to the current user on this object.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this task schedule.
 * @return ApiTaskSchedulesMyPermissionsRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesMyPermissions(ctx _context.Context, pulpId string) ApiTaskSchedulesMyPermissionsRequest {
	return ApiTaskSchedulesMyPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return MyPermissionsResponse
 */
func (a *TaskSchedulesApiService) TaskSchedulesMyPermissionsExecute(r ApiTaskSchedulesMyPermissionsRequest) (MyPermissionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MyPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesMyPermissions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/{pulp_id}/my_permissions/"
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

type ApiTaskSchedulesReadRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiTaskSchedulesReadRequest) Fields(fields string) ApiTaskSchedulesReadRequest {
	r.fields = &fields
	return r
}
func (r ApiTaskSchedulesReadRequest) ExcludeFields(excludeFields string) ApiTaskSchedulesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiTaskSchedulesReadRequest) Execute() (TaskScheduleResponse, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesReadExecute(r)
}

/*
 * TaskSchedulesRead Inspect a task schedule
 * ViewSet to monitor task schedules.

NOTE: This feature is in tech-preview and may change in backwards incompatible ways.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this task schedule.
 * @return ApiTaskSchedulesReadRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesRead(ctx _context.Context, pulpId string) ApiTaskSchedulesReadRequest {
	return ApiTaskSchedulesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return TaskScheduleResponse
 */
func (a *TaskSchedulesApiService) TaskSchedulesReadExecute(r ApiTaskSchedulesReadRequest) (TaskScheduleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TaskScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/{pulp_id}/"
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

type ApiTaskSchedulesRemoveRoleRequest struct {
	ctx _context.Context
	ApiService *TaskSchedulesApiService
	pulpId string
	nestedRole *NestedRole
}

func (r ApiTaskSchedulesRemoveRoleRequest) NestedRole(nestedRole NestedRole) ApiTaskSchedulesRemoveRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r ApiTaskSchedulesRemoveRoleRequest) Execute() (NestedRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.TaskSchedulesRemoveRoleExecute(r)
}

/*
 * TaskSchedulesRemoveRole Method for TaskSchedulesRemoveRole
 * Remove a role for this object from users/groups.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this task schedule.
 * @return ApiTaskSchedulesRemoveRoleRequest
 */
func (a *TaskSchedulesApiService) TaskSchedulesRemoveRole(ctx _context.Context, pulpId string) ApiTaskSchedulesRemoveRoleRequest {
	return ApiTaskSchedulesRemoveRoleRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return NestedRoleResponse
 */
func (a *TaskSchedulesApiService) TaskSchedulesRemoveRoleExecute(r ApiTaskSchedulesRemoveRoleRequest) (NestedRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesRemoveRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/{pulp_id}/remove_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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