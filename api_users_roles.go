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

// UsersRolesApiService UsersRolesApi service
type UsersRolesApiService service

type ApiUsersRolesCreateRequest struct {
	ctx _context.Context
	ApiService *UsersRolesApiService
	userPk string
	userRole *UserRole
}

func (r ApiUsersRolesCreateRequest) UserRole(userRole UserRole) ApiUsersRolesCreateRequest {
	r.userRole = &userRole
	return r
}

func (r ApiUsersRolesCreateRequest) Execute() (UserRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.UsersRolesCreateExecute(r)
}

/*
 * UsersRolesCreate Create an user role
 * ViewSet for UserRole.

NOTE: This API endpoint is in "tech preview" and subject to change
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userPk
 * @return ApiUsersRolesCreateRequest
 */
func (a *UsersRolesApiService) UsersRolesCreate(ctx _context.Context, userPk string) ApiUsersRolesCreateRequest {
	return ApiUsersRolesCreateRequest{
		ApiService: a,
		ctx: ctx,
		userPk: userPk,
	}
}

/*
 * Execute executes the request
 * @return UserRoleResponse
 */
func (a *UsersRolesApiService) UsersRolesCreateExecute(r ApiUsersRolesCreateRequest) (UserRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersRolesApiService.UsersRolesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/{user_pk}/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"user_pk"+"}", _neturl.PathEscape(parameterToString(r.userPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.userRole == nil {
		return localVarReturnValue, nil, reportError("userRole is required and must be specified")
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
	localVarPostBody = r.userRole
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

type ApiUsersRolesDeleteRequest struct {
	ctx _context.Context
	ApiService *UsersRolesApiService
	pulpId string
	userPk string
}


func (r ApiUsersRolesDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersRolesDeleteExecute(r)
}

/*
 * UsersRolesDelete Delete an user role
 * ViewSet for UserRole.

NOTE: This API endpoint is in "tech preview" and subject to change
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this user role.
 * @param userPk
 * @return ApiUsersRolesDeleteRequest
 */
func (a *UsersRolesApiService) UsersRolesDelete(ctx _context.Context, pulpId string, userPk string) ApiUsersRolesDeleteRequest {
	return ApiUsersRolesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
		userPk: userPk,
	}
}

/*
 * Execute executes the request
 */
func (a *UsersRolesApiService) UsersRolesDeleteExecute(r ApiUsersRolesDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersRolesApiService.UsersRolesDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/{user_pk}/roles/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_pk"+"}", _neturl.PathEscape(parameterToString(r.userPk, "")), -1)

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

type ApiUsersRolesListRequest struct {
	ctx _context.Context
	ApiService *UsersRolesApiService
	userPk string
	contentObject *string
	limit *int32
	offset *int32
	ordering *[]string
	role *string
	roleContains *string
	roleIcontains *string
	roleIn *[]string
	roleStartswith *string
	fields *string
	excludeFields *string
}

func (r ApiUsersRolesListRequest) ContentObject(contentObject string) ApiUsersRolesListRequest {
	r.contentObject = &contentObject
	return r
}
func (r ApiUsersRolesListRequest) Limit(limit int32) ApiUsersRolesListRequest {
	r.limit = &limit
	return r
}
func (r ApiUsersRolesListRequest) Offset(offset int32) ApiUsersRolesListRequest {
	r.offset = &offset
	return r
}
func (r ApiUsersRolesListRequest) Ordering(ordering []string) ApiUsersRolesListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiUsersRolesListRequest) Role(role string) ApiUsersRolesListRequest {
	r.role = &role
	return r
}
func (r ApiUsersRolesListRequest) RoleContains(roleContains string) ApiUsersRolesListRequest {
	r.roleContains = &roleContains
	return r
}
func (r ApiUsersRolesListRequest) RoleIcontains(roleIcontains string) ApiUsersRolesListRequest {
	r.roleIcontains = &roleIcontains
	return r
}
func (r ApiUsersRolesListRequest) RoleIn(roleIn []string) ApiUsersRolesListRequest {
	r.roleIn = &roleIn
	return r
}
func (r ApiUsersRolesListRequest) RoleStartswith(roleStartswith string) ApiUsersRolesListRequest {
	r.roleStartswith = &roleStartswith
	return r
}
func (r ApiUsersRolesListRequest) Fields(fields string) ApiUsersRolesListRequest {
	r.fields = &fields
	return r
}
func (r ApiUsersRolesListRequest) ExcludeFields(excludeFields string) ApiUsersRolesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiUsersRolesListRequest) Execute() (PaginatedUserRoleResponseList, *_nethttp.Response, error) {
	return r.ApiService.UsersRolesListExecute(r)
}

/*
 * UsersRolesList List user roles
 * ViewSet for UserRole.

NOTE: This API endpoint is in "tech preview" and subject to change
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userPk
 * @return ApiUsersRolesListRequest
 */
func (a *UsersRolesApiService) UsersRolesList(ctx _context.Context, userPk string) ApiUsersRolesListRequest {
	return ApiUsersRolesListRequest{
		ApiService: a,
		ctx: ctx,
		userPk: userPk,
	}
}

/*
 * Execute executes the request
 * @return PaginatedUserRoleResponseList
 */
func (a *UsersRolesApiService) UsersRolesListExecute(r ApiUsersRolesListRequest) (PaginatedUserRoleResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedUserRoleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersRolesApiService.UsersRolesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/{user_pk}/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"user_pk"+"}", _neturl.PathEscape(parameterToString(r.userPk, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.contentObject != nil {
		localVarQueryParams.Add("content_object", parameterToString(*r.contentObject, ""))
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
	if r.role != nil {
		localVarQueryParams.Add("role", parameterToString(*r.role, ""))
	}
	if r.roleContains != nil {
		localVarQueryParams.Add("role__contains", parameterToString(*r.roleContains, ""))
	}
	if r.roleIcontains != nil {
		localVarQueryParams.Add("role__icontains", parameterToString(*r.roleIcontains, ""))
	}
	if r.roleIn != nil {
		localVarQueryParams.Add("role__in", parameterToString(*r.roleIn, "csv"))
	}
	if r.roleStartswith != nil {
		localVarQueryParams.Add("role__startswith", parameterToString(*r.roleStartswith, ""))
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

type ApiUsersRolesReadRequest struct {
	ctx _context.Context
	ApiService *UsersRolesApiService
	pulpId string
	userPk string
	fields *string
	excludeFields *string
}

func (r ApiUsersRolesReadRequest) Fields(fields string) ApiUsersRolesReadRequest {
	r.fields = &fields
	return r
}
func (r ApiUsersRolesReadRequest) ExcludeFields(excludeFields string) ApiUsersRolesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiUsersRolesReadRequest) Execute() (UserRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.UsersRolesReadExecute(r)
}

/*
 * UsersRolesRead Inspect an user role
 * ViewSet for UserRole.

NOTE: This API endpoint is in "tech preview" and subject to change
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulpId A UUID string identifying this user role.
 * @param userPk
 * @return ApiUsersRolesReadRequest
 */
func (a *UsersRolesApiService) UsersRolesRead(ctx _context.Context, pulpId string, userPk string) ApiUsersRolesReadRequest {
	return ApiUsersRolesReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
		userPk: userPk,
	}
}

/*
 * Execute executes the request
 * @return UserRoleResponse
 */
func (a *UsersRolesApiService) UsersRolesReadExecute(r ApiUsersRolesReadRequest) (UserRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersRolesApiService.UsersRolesRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/{user_pk}/roles/{pulp_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_id"+"}", _neturl.PathEscape(parameterToString(r.pulpId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_pk"+"}", _neturl.PathEscape(parameterToString(r.userPk, "")), -1)

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
