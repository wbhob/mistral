# \BetaLibrariesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibrariesCreateV1**](BetaLibrariesAPI.md#LibrariesCreateV1) | **Post** /v1/libraries | Create a new Library.
[**LibrariesDeleteV1**](BetaLibrariesAPI.md#LibrariesDeleteV1) | **Delete** /v1/libraries/{library_id} | Delete a library and all of it&#39;s document.
[**LibrariesGetV1**](BetaLibrariesAPI.md#LibrariesGetV1) | **Get** /v1/libraries/{library_id} | Detailed information about a specific Library.
[**LibrariesListV1**](BetaLibrariesAPI.md#LibrariesListV1) | **Get** /v1/libraries | List all libraries you have access to.
[**LibrariesUpdateV1**](BetaLibrariesAPI.md#LibrariesUpdateV1) | **Put** /v1/libraries/{library_id} | Update a library.



## LibrariesCreateV1

> LibraryOut LibrariesCreateV1(ctx).LibraryIn(libraryIn).Execute()

Create a new Library.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wbhob/mistral"
)

func main() {
	libraryIn := *openapiclient.NewLibraryIn("Name_example") // LibraryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesCreateV1(context.Background()).LibraryIn(libraryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesCreateV1`: LibraryOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryIn** | [**LibraryIn**](LibraryIn.md) |  | 

### Return type

[**LibraryOut**](LibraryOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDeleteV1

> LibraryOut LibrariesDeleteV1(ctx, libraryId).Execute()

Delete a library and all of it's document.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wbhob/mistral"
)

func main() {
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesDeleteV1(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDeleteV1`: LibraryOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesDeleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryOut**](LibraryOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesGetV1

> LibraryOut LibrariesGetV1(ctx, libraryId).Execute()

Detailed information about a specific Library.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wbhob/mistral"
)

func main() {
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesGetV1(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesGetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesGetV1`: LibraryOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesGetV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesGetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryOut**](LibraryOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesListV1

> ListLibraryOut LibrariesListV1(ctx).Execute()

List all libraries you have access to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wbhob/mistral"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesListV1`: ListLibraryOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesListV1Request struct via the builder pattern


### Return type

[**ListLibraryOut**](ListLibraryOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesUpdateV1

> LibraryOut LibrariesUpdateV1(ctx, libraryId).LibraryInUpdate(libraryInUpdate).Execute()

Update a library.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wbhob/mistral"
)

func main() {
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	libraryInUpdate := *openapiclient.NewLibraryInUpdate() // LibraryInUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesUpdateV1(context.Background(), libraryId).LibraryInUpdate(libraryInUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesUpdateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesUpdateV1`: LibraryOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesUpdateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesUpdateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryInUpdate** | [**LibraryInUpdate**](LibraryInUpdate.md) |  | 

### Return type

[**LibraryOut**](LibraryOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

