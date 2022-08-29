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
)

// Linger please
var (
	_ _context.Context
)

// RepositoriesReclaimSpaceApiService RepositoriesReclaimSpaceApi service
type RepositoriesReclaimSpaceApiService service

type ApiRepositoriesReclaimSpaceReclaimRequest struct {
	ctx _context.Context
	ApiService *RepositoriesReclaimSpaceApiService
	reclaimSpace *ReclaimSpace
}

func (r ApiRepositoriesReclaimSpaceReclaimRequest) ReclaimSpace(reclaimSpace ReclaimSpace) ApiRepositoriesReclaimSpaceReclaimRequest {
	r.reclaimSpace = &reclaimSpace
	return r
}

func (r ApiRepositoriesReclaimSpaceReclaimRequest) Execute() (AsyncOperationResponse, *_nethttp.Response, error) {
	return r.ApiService.RepositoriesReclaimSpaceReclaimExecute(r)
}

/*
 * RepositoriesReclaimSpaceReclaim Method for RepositoriesReclaimSpaceReclaim
 * Trigger an asynchronous space reclaim operation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRepositoriesReclaimSpaceReclaimRequest
 */
func (a *RepositoriesReclaimSpaceApiService) RepositoriesReclaimSpaceReclaim(ctx _context.Context) ApiRepositoriesReclaimSpaceReclaimRequest {
	return ApiRepositoriesReclaimSpaceReclaimRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AsyncOperationResponse
 */
func (a *RepositoriesReclaimSpaceApiService) RepositoriesReclaimSpaceReclaimExecute(r ApiRepositoriesReclaimSpaceReclaimRequest) (AsyncOperationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesReclaimSpaceApiService.RepositoriesReclaimSpaceReclaim")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/reclaim_space/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.reclaimSpace == nil {
		return localVarReturnValue, nil, reportError("reclaimSpace is required and must be specified")
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
	localVarPostBody = r.reclaimSpace
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
