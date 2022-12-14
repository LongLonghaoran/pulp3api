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
	"time"
)


// RepositoryVersionsApiService RepositoryVersionsApi service
type RepositoryVersionsApiService service

type ApiRepositoryVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoryVersionsApiService
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
func (r ApiRepositoryVersionsListRequest) Content(content string) ApiRepositoryVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r ApiRepositoryVersionsListRequest) ContentIn(contentIn string) ApiRepositoryVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r ApiRepositoryVersionsListRequest) Limit(limit int32) ApiRepositoryVersionsListRequest {
	r.limit = &limit
	return r
}

func (r ApiRepositoryVersionsListRequest) Number(number int32) ApiRepositoryVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r ApiRepositoryVersionsListRequest) NumberGt(numberGt int32) ApiRepositoryVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r ApiRepositoryVersionsListRequest) NumberGte(numberGte int32) ApiRepositoryVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r ApiRepositoryVersionsListRequest) NumberLt(numberLt int32) ApiRepositoryVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r ApiRepositoryVersionsListRequest) NumberLte(numberLte int32) ApiRepositoryVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r ApiRepositoryVersionsListRequest) NumberRange(numberRange []int32) ApiRepositoryVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r ApiRepositoryVersionsListRequest) Offset(offset int32) ApiRepositoryVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ApiRepositoryVersionsListRequest) Ordering(ordering []string) ApiRepositoryVersionsListRequest {
	r.ordering = &ordering
	return r
}

// ISO 8601 formatted dates are supported
func (r ApiRepositoryVersionsListRequest) PulpCreated(pulpCreated time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r ApiRepositoryVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r ApiRepositoryVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r ApiRepositoryVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r ApiRepositoryVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r ApiRepositoryVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) ApiRepositoryVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r ApiRepositoryVersionsListRequest) Fields(fields string) ApiRepositoryVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiRepositoryVersionsListRequest) ExcludeFields(excludeFields string) ApiRepositoryVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiRepositoryVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoryVersionsListExecute(r)
}

/*
RepositoryVersionsList List repository versions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRepositoryVersionsListRequest
*/
func (a *RepositoryVersionsApiService) RepositoryVersionsList(ctx context.Context) ApiRepositoryVersionsListRequest {
	return ApiRepositoryVersionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoryVersionsApiService) RepositoryVersionsListExecute(r ApiRepositoryVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoryVersionsApiService.RepositoryVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repository_versions/"

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
