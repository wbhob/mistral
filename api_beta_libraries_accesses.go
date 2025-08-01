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


type BetaLibrariesAccessesAPI interface {

	/*
	LibrariesShareCreateV1 Create or update an access level.

	Given a library id, you can create or update the access level of an entity. You have to be owner of the library to share a library. An owner cannot change their own role. A library cannot be shared outside of the organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesShareCreateV1Request
	*/
	LibrariesShareCreateV1(ctx context.Context, libraryId string) ApiLibrariesShareCreateV1Request

	// LibrariesShareCreateV1Execute executes the request
	//  @return SharingOut
	LibrariesShareCreateV1Execute(r ApiLibrariesShareCreateV1Request) (*SharingOut, *http.Response, error)

	/*
	LibrariesShareDeleteV1 Delete an access level.

	Given a library id, you can delete the access level of an entity. An owner cannot delete it's own access. You have to be the owner of the library to delete an acces other than yours.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesShareDeleteV1Request
	*/
	LibrariesShareDeleteV1(ctx context.Context, libraryId string) ApiLibrariesShareDeleteV1Request

	// LibrariesShareDeleteV1Execute executes the request
	//  @return SharingOut
	LibrariesShareDeleteV1Execute(r ApiLibrariesShareDeleteV1Request) (*SharingOut, *http.Response, error)

	/*
	LibrariesShareListV1 List all of the access to this library.

	Given a library, list all of the Entity that have access and to what level.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param libraryId
	@return ApiLibrariesShareListV1Request
	*/
	LibrariesShareListV1(ctx context.Context, libraryId string) ApiLibrariesShareListV1Request

	// LibrariesShareListV1Execute executes the request
	//  @return ListSharingOut
	LibrariesShareListV1Execute(r ApiLibrariesShareListV1Request) (*ListSharingOut, *http.Response, error)
}

// BetaLibrariesAccessesAPIService BetaLibrariesAccessesAPI service
type BetaLibrariesAccessesAPIService service

type ApiLibrariesShareCreateV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAccessesAPI
	libraryId string
	sharingIn *SharingIn
}

func (r ApiLibrariesShareCreateV1Request) SharingIn(sharingIn SharingIn) ApiLibrariesShareCreateV1Request {
	r.sharingIn = &sharingIn
	return r
}

func (r ApiLibrariesShareCreateV1Request) Execute() (*SharingOut, *http.Response, error) {
	return r.ApiService.LibrariesShareCreateV1Execute(r)
}

/*
LibrariesShareCreateV1 Create or update an access level.

Given a library id, you can create or update the access level of an entity. You have to be owner of the library to share a library. An owner cannot change their own role. A library cannot be shared outside of the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesShareCreateV1Request
*/
func (a *BetaLibrariesAccessesAPIService) LibrariesShareCreateV1(ctx context.Context, libraryId string) ApiLibrariesShareCreateV1Request {
	return ApiLibrariesShareCreateV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return SharingOut
func (a *BetaLibrariesAccessesAPIService) LibrariesShareCreateV1Execute(r ApiLibrariesShareCreateV1Request) (*SharingOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SharingOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAccessesAPIService.LibrariesShareCreateV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"library_id"+"}", url.PathEscape(parameterValueToString(r.libraryId, "libraryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sharingIn == nil {
		return localVarReturnValue, nil, reportError("sharingIn is required and must be specified")
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
	localVarPostBody = r.sharingIn
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

type ApiLibrariesShareDeleteV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAccessesAPI
	libraryId string
	sharingDelete *SharingDelete
}

func (r ApiLibrariesShareDeleteV1Request) SharingDelete(sharingDelete SharingDelete) ApiLibrariesShareDeleteV1Request {
	r.sharingDelete = &sharingDelete
	return r
}

func (r ApiLibrariesShareDeleteV1Request) Execute() (*SharingOut, *http.Response, error) {
	return r.ApiService.LibrariesShareDeleteV1Execute(r)
}

/*
LibrariesShareDeleteV1 Delete an access level.

Given a library id, you can delete the access level of an entity. An owner cannot delete it's own access. You have to be the owner of the library to delete an acces other than yours.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesShareDeleteV1Request
*/
func (a *BetaLibrariesAccessesAPIService) LibrariesShareDeleteV1(ctx context.Context, libraryId string) ApiLibrariesShareDeleteV1Request {
	return ApiLibrariesShareDeleteV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return SharingOut
func (a *BetaLibrariesAccessesAPIService) LibrariesShareDeleteV1Execute(r ApiLibrariesShareDeleteV1Request) (*SharingOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SharingOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAccessesAPIService.LibrariesShareDeleteV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"library_id"+"}", url.PathEscape(parameterValueToString(r.libraryId, "libraryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sharingDelete == nil {
		return localVarReturnValue, nil, reportError("sharingDelete is required and must be specified")
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
	localVarPostBody = r.sharingDelete
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

type ApiLibrariesShareListV1Request struct {
	ctx context.Context
	ApiService BetaLibrariesAccessesAPI
	libraryId string
}

func (r ApiLibrariesShareListV1Request) Execute() (*ListSharingOut, *http.Response, error) {
	return r.ApiService.LibrariesShareListV1Execute(r)
}

/*
LibrariesShareListV1 List all of the access to this library.

Given a library, list all of the Entity that have access and to what level.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param libraryId
 @return ApiLibrariesShareListV1Request
*/
func (a *BetaLibrariesAccessesAPIService) LibrariesShareListV1(ctx context.Context, libraryId string) ApiLibrariesShareListV1Request {
	return ApiLibrariesShareListV1Request{
		ApiService: a,
		ctx: ctx,
		libraryId: libraryId,
	}
}

// Execute executes the request
//  @return ListSharingOut
func (a *BetaLibrariesAccessesAPIService) LibrariesShareListV1Execute(r ApiLibrariesShareListV1Request) (*ListSharingOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListSharingOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLibrariesAccessesAPIService.LibrariesShareListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/libraries/{library_id}/share"
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
