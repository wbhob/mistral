/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type BetaLibrariesAPI interface {

	/*
	LibrariesCreateV1 Create a new Library.

	Create a new Library, you will be marked as the owner and only you will have the possibility to share it with others. When first created this will only be accessible by you.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLibrariesCreateV1Request
	*/
	LibrariesCreateV1(ctx context.Context) ApiLibrariesCreateV1Request

	// LibrariesCreateV1Execute executes the request
	//  @return LibraryOut
	LibrariesCreateV1Execute(r ApiLibrariesCreateV1Request) (*LibraryOut, *http.Response, error)

	/*
	LibrariesDeleteV1 Delete a library and all of it's document.

	Given a library id, deletes it together with all documents that have been uploaded to that library.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesDeleteV1Request
	*/
	LibrariesDeleteV1(ctx context.Context, libraryId string) ApiLibrariesDeleteV1Request

	// LibrariesDeleteV1Execute executes the request
	//  @return LibraryOut
	LibrariesDeleteV1Execute(r ApiLibrariesDeleteV1Request) (*LibraryOut, *http.Response, error)

	/*
	LibrariesGetV1 Detailed information about a specific Library.

	Given a library id, details information about that Library.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesGetV1Request
	*/
	LibrariesGetV1(ctx context.Context, libraryId string) ApiLibrariesGetV1Request

	// LibrariesGetV1Execute executes the request
	//  @return LibraryOut
	LibrariesGetV1Execute(r ApiLibrariesGetV1Request) (*LibraryOut, *http.Response, error)

	/*
	LibrariesListV1 List all libraries you have access to.

	List all libraries that you have created or have been shared with you.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLibrariesListV1Request
	*/
	LibrariesListV1(ctx context.Context) ApiLibrariesListV1Request

	// LibrariesListV1Execute executes the request
	//  @return ListLibraryOut
	LibrariesListV1Execute(r ApiLibrariesListV1Request) (*ListLibraryOut, *http.Response, error)

	/*
	LibrariesUpdateV1 Update a library.

	Given a library id, you can update the name and description.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesUpdateV1Request
	*/
	LibrariesUpdateV1(ctx context.Context, libraryId string) ApiLibrariesUpdateV1Request

	// LibrariesUpdateV1Execute executes the request
	//  @return LibraryOut
	LibrariesUpdateV1Execute(r ApiLibrariesUpdateV1Request) (*LibraryOut, *http.Response, error)
}

// BetaLibrariesAPIService BetaLibrariesAPI service
type BetaLibrariesAPIService service

type ApiLibrariesCreateV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAPI
	libraryIn *LibraryIn
}

func (r ApiLibrariesCreateV1Request) LibraryIn(libraryIn LibraryIn) ApiLibrariesCreateV1Request {
	r.libraryIn = &libraryIn
	return r
}

func (r ApiLibrariesCreateV1Request) Execute() (*LibraryOut, *http.Response, error) {
	return r.ApiService.LibrariesCreateV1Execute(r)
}

/*
LibrariesCreateV1 Create a new Library.

Create a new Library, you will be marked as the owner and only you will have the possibility to share it with others. When first created this will only be accessible by you.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLibrariesCreateV1Request
*/
func (a *BetaLibrariesAPIService) LibrariesCreateV1(ctx context.Context) ApiLibrariesCreateV1Request {
	return ApiLibrariesCreateV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LibraryOut
func (a *BetaLibrariesAPIService) LibrariesCreateV1Execute(r ApiLibrariesCreateV1Request) (*LibraryOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LibraryOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAPIService.LibrariesCreateV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.libraryIn == nil {
		return localVarReturnValue, nil, reportError("libraryIn is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.libraryIn
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiLibrariesDeleteV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAPI
	libraryId string
}

func (r ApiLibrariesDeleteV1Request) Execute() (*LibraryOut, *http.Response, error) {
	return r.ApiService.LibrariesDeleteV1Execute(r)
}

/*
LibrariesDeleteV1 Delete a library and all of it's document.

Given a library id, deletes it together with all documents that have been uploaded to that library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesDeleteV1Request
*/
func (a *BetaLibrariesAPIService) LibrariesDeleteV1(ctx context.Context, libraryId string) ApiLibrariesDeleteV1Request {
	return ApiLibrariesDeleteV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return LibraryOut
func (a *BetaLibrariesAPIService) LibrariesDeleteV1Execute(r ApiLibrariesDeleteV1Request) (*LibraryOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LibraryOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAPIService.LibrariesDeleteV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"library_id"+"}", url.PathEscape(parameterValueToString(r.libraryId, "libraryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiLibrariesGetV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAPI
	libraryId string
}

func (r ApiLibrariesGetV1Request) Execute() (*LibraryOut, *http.Response, error) {
	return r.ApiService.LibrariesGetV1Execute(r)
}

/*
LibrariesGetV1 Detailed information about a specific Library.

Given a library id, details information about that Library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesGetV1Request
*/
func (a *BetaLibrariesAPIService) LibrariesGetV1(ctx context.Context, libraryId string) ApiLibrariesGetV1Request {
	return ApiLibrariesGetV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return LibraryOut
func (a *BetaLibrariesAPIService) LibrariesGetV1Execute(r ApiLibrariesGetV1Request) (*LibraryOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LibraryOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAPIService.LibrariesGetV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"library_id"+"}", url.PathEscape(parameterValueToString(r.libraryId, "libraryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiLibrariesListV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAPI
}

func (r ApiLibrariesListV1Request) Execute() (*ListLibraryOut, *http.Response, error) {
	return r.ApiService.LibrariesListV1Execute(r)
}

/*
LibrariesListV1 List all libraries you have access to.

List all libraries that you have created or have been shared with you.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLibrariesListV1Request
*/
func (a *BetaLibrariesAPIService) LibrariesListV1(ctx context.Context) ApiLibrariesListV1Request {
	return ApiLibrariesListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListLibraryOut
func (a *BetaLibrariesAPIService) LibrariesListV1Execute(r ApiLibrariesListV1Request) (*ListLibraryOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListLibraryOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAPIService.LibrariesListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiLibrariesUpdateV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAPI
	libraryId string
	libraryInUpdate *LibraryInUpdate
}

func (r ApiLibrariesUpdateV1Request) LibraryInUpdate(libraryInUpdate LibraryInUpdate) ApiLibrariesUpdateV1Request {
	r.libraryInUpdate = &libraryInUpdate
	return r
}

func (r ApiLibrariesUpdateV1Request) Execute() (*LibraryOut, *http.Response, error) {
	return r.ApiService.LibrariesUpdateV1Execute(r)
}

/*
LibrariesUpdateV1 Update a library.

Given a library id, you can update the name and description.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesUpdateV1Request
*/
func (a *BetaLibrariesAPIService) LibrariesUpdateV1(ctx context.Context, libraryId string) ApiLibrariesUpdateV1Request {
	return ApiLibrariesUpdateV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return LibraryOut
func (a *BetaLibrariesAPIService) LibrariesUpdateV1Execute(r ApiLibrariesUpdateV1Request) (*LibraryOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LibraryOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAPIService.LibrariesUpdateV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"library_id"+"}", url.PathEscape(parameterValueToString(r.libraryId, "libraryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.libraryInUpdate == nil {
		return localVarReturnValue, nil, reportError("libraryInUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.libraryInUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
