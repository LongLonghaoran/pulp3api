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
	"reflect"
)


// ContentTagsApiService ContentTagsApi service
type ContentTagsApiService service

type ApiContentContainerTagsListRequest struct {
	ctx context.Context
	ApiService *ContentTagsApiService
	digest *[]string
	limit *int32
	mediaType *[]string
	name *string
	nameIn *[]string
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *string
	excludeFields *string
}

// Multiple values may be separated by commas.
func (r ApiContentContainerTagsListRequest) Digest(digest []string) ApiContentContainerTagsListRequest {
	r.digest = &digest
	return r
}

// Number of results to return per page.
func (r ApiContentContainerTagsListRequest) Limit(limit int32) ApiContentContainerTagsListRequest {
	r.limit = &limit
	return r
}

func (r ApiContentContainerTagsListRequest) MediaType(mediaType []string) ApiContentContainerTagsListRequest {
	r.mediaType = &mediaType
	return r
}

// Filter results where name matches value
func (r ApiContentContainerTagsListRequest) Name(name string) ApiContentContainerTagsListRequest {
	r.name = &name
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ApiContentContainerTagsListRequest) NameIn(nameIn []string) ApiContentContainerTagsListRequest {
	r.nameIn = &nameIn
	return r
}

// The initial index from which to return the results.
func (r ApiContentContainerTagsListRequest) Offset(offset int32) ApiContentContainerTagsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiContentContainerTagsListRequest) Ordering(ordering []string) ApiContentContainerTagsListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ApiContentContainerTagsListRequest) RepositoryVersion(repositoryVersion string) ApiContentContainerTagsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ApiContentContainerTagsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ApiContentContainerTagsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ApiContentContainerTagsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ApiContentContainerTagsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ApiContentContainerTagsListRequest) Fields(fields string) ApiContentContainerTagsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiContentContainerTagsListRequest) ExcludeFields(excludeFields string) ApiContentContainerTagsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentContainerTagsListRequest) Execute() (*PaginatedcontainerTagResponseList, *http.Response, error) {
	return r.ApiService.ContentContainerTagsListExecute(r)
}

/*
ContentContainerTagsList List tags

ViewSet for Tag.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContentContainerTagsListRequest
*/
func (a *ContentTagsApiService) ContentContainerTagsList(ctx context.Context) ApiContentContainerTagsListRequest {
	return ApiContentContainerTagsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedcontainerTagResponseList
func (a *ContentTagsApiService) ContentContainerTagsListExecute(r ApiContentContainerTagsListRequest) (*PaginatedcontainerTagResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcontainerTagResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentTagsApiService.ContentContainerTagsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/container/tags/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.digest != nil {
		localVarQueryParams.Add("digest", parameterToString(*r.digest, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.mediaType != nil {
		t := *r.mediaType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("media_type", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("media_type", parameterToString(t, "multi"))
		}
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.nameIn != nil {
		localVarQueryParams.Add("name__in", parameterToString(*r.nameIn, "csv"))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, "csv"))
	}
	if r.repositoryVersion != nil {
		localVarQueryParams.Add("repository_version", parameterToString(*r.repositoryVersion, ""))
	}
	if r.repositoryVersionAdded != nil {
		localVarQueryParams.Add("repository_version_added", parameterToString(*r.repositoryVersionAdded, ""))
	}
	if r.repositoryVersionRemoved != nil {
		localVarQueryParams.Add("repository_version_removed", parameterToString(*r.repositoryVersionRemoved, ""))
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

type ApiContentContainerTagsReadRequest struct {
	ctx context.Context
	ApiService *ContentTagsApiService
	pulpId string
	fields *string
	excludeFields *string
}

// A list of fields to include in the response.
func (r ApiContentContainerTagsReadRequest) Fields(fields string) ApiContentContainerTagsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiContentContainerTagsReadRequest) ExcludeFields(excludeFields string) ApiContentContainerTagsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiContentContainerTagsReadRequest) Execute() (*ContainerTagResponse, *http.Response, error) {
	return r.ApiService.ContentContainerTagsReadExecute(r)
}

/*
ContentContainerTagsRead Inspect a tag

ViewSet for Tag.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpId A UUID string identifying this tag.
 @return ApiContentContainerTagsReadRequest
*/
func (a *ContentTagsApiService) ContentContainerTagsRead(ctx context.Context, pulpId string) ApiContentContainerTagsReadRequest {
	return ApiContentContainerTagsReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpId: pulpId,
	}
}

// Execute executes the request
//  @return ContainerTagResponse
func (a *ContentTagsApiService) ContentContainerTagsReadExecute(r ApiContentContainerTagsReadRequest) (*ContainerTagResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContainerTagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentTagsApiService.ContentContainerTagsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/container/tags/{pulp_id}/"
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
