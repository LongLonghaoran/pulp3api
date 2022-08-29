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

// PublicationsApiService PublicationsApi service
type PublicationsApiService service

type ApiPublicationsDeleteRequest struct {
	ctx _context.Context
	ApiService *PublicationsApiService
	pulpId string
}


func (r ApiPublicationsDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PublicationsDeleteExecute(r)
}

/*
 * PublicationsDelete Delete a publication
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
 * @param pulpId A UUID string identifying this publication.
 * @return ApiPublicationsDeleteRequest
 */
func (a *PublicationsApiService) PublicationsDelete(ctx _context.Context, pulpId string) ApiPublicationsDeleteRequest {
	return ApiPublicationsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 */
func (a *PublicationsApiService) PublicationsDeleteExecute(r ApiPublicationsDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsApiService.PublicationsDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/{pulp_id}/"
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

type ApiPublicationsListRequest struct {
	ctx _context.Context
	ApiService *PublicationsApiService
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	repository *string
	repositoryVersion *string
	fields *string
	excludeFields *string
}

func (r ApiPublicationsListRequest) Content(content string) ApiPublicationsListRequest {
	r.content = &content
	return r
}
func (r ApiPublicationsListRequest) ContentIn(contentIn string) ApiPublicationsListRequest {
	r.contentIn = &contentIn
	return r
}
func (r ApiPublicationsListRequest) Limit(limit int32) ApiPublicationsListRequest {
	r.limit = &limit
	return r
}
func (r ApiPublicationsListRequest) Offset(offset int32) ApiPublicationsListRequest {
	r.offset = &offset
	return r
}
func (r ApiPublicationsListRequest) Ordering(ordering []string) ApiPublicationsListRequest {
	r.ordering = &ordering
	return r
}
func (r ApiPublicationsListRequest) PulpCreated(pulpCreated time.Time) ApiPublicationsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}
func (r ApiPublicationsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) ApiPublicationsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}
func (r ApiPublicationsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) ApiPublicationsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}
func (r ApiPublicationsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) ApiPublicationsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}
func (r ApiPublicationsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) ApiPublicationsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}
func (r ApiPublicationsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) ApiPublicationsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}
func (r ApiPublicationsListRequest) Repository(repository string) ApiPublicationsListRequest {
	r.repository = &repository
	return r
}
func (r ApiPublicationsListRequest) RepositoryVersion(repositoryVersion string) ApiPublicationsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}
func (r ApiPublicationsListRequest) Fields(fields string) ApiPublicationsListRequest {
	r.fields = &fields
	return r
}
func (r ApiPublicationsListRequest) ExcludeFields(excludeFields string) ApiPublicationsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiPublicationsListRequest) Execute() (PaginatedPublicationResponseList, *_nethttp.Response, error) {
	return r.ApiService.PublicationsListExecute(r)
}

/*
 * PublicationsList List publications
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
 * @return ApiPublicationsListRequest
 */
func (a *PublicationsApiService) PublicationsList(ctx _context.Context) ApiPublicationsListRequest {
	return ApiPublicationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedPublicationResponseList
 */
func (a *PublicationsApiService) PublicationsListExecute(r ApiPublicationsListRequest) (PaginatedPublicationResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsApiService.PublicationsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.content != nil {
		localVarQueryParams.Add("content", parameterToString(*r.content, ""))
	}
	if r.contentIn != nil {
		localVarQueryParams.Add("content__in", parameterToString(*r.contentIn, ""))
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
	if r.repository != nil {
		localVarQueryParams.Add("repository", parameterToString(*r.repository, ""))
	}
	if r.repositoryVersion != nil {
		localVarQueryParams.Add("repository_version", parameterToString(*r.repositoryVersion, ""))
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

type ApiPublicationsReadRequest struct {
	ctx _context.Context
	ApiService *PublicationsApiService
	pulpId string
	fields *string
	excludeFields *string
}

func (r ApiPublicationsReadRequest) Fields(fields string) ApiPublicationsReadRequest {
	r.fields = &fields
	return r
}
func (r ApiPublicationsReadRequest) ExcludeFields(excludeFields string) ApiPublicationsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiPublicationsReadRequest) Execute() (PublicationResponse, *_nethttp.Response, error) {
	return r.ApiService.PublicationsReadExecute(r)
}

/*
 * PublicationsRead Inspect a publication
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
 * @param pulpId A UUID string identifying this publication.
 * @return ApiPublicationsReadRequest
 */
func (a *PublicationsApiService) PublicationsRead(ctx _context.Context, pulpId string) ApiPublicationsReadRequest {
	return ApiPublicationsReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

/*
 * Execute executes the request
 * @return PublicationResponse
 */
func (a *PublicationsApiService) PublicationsReadExecute(r ApiPublicationsReadRequest) (PublicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PublicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsApiService.PublicationsRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/{pulp_id}/"
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
