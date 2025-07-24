# \BetaLibrariesDocumentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibrariesDocumentsDeleteV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsDeleteV1) | **Delete** /v1/libraries/{library_id}/documents/{document_id} | Delete a document.
[**LibrariesDocumentsGetExtractedTextSignedUrlV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsGetExtractedTextSignedUrlV1) | **Get** /v1/libraries/{library_id}/documents/{document_id}/extracted-text-signed-url | Retrieve the signed URL of text extracted from a given document.
[**LibrariesDocumentsGetSignedUrlV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsGetSignedUrlV1) | **Get** /v1/libraries/{library_id}/documents/{document_id}/signed-url | Retrieve the signed URL of a specific document.
[**LibrariesDocumentsGetStatusV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsGetStatusV1) | **Get** /v1/libraries/{library_id}/documents/{document_id}/status | Retrieve the processing status of a specific document.
[**LibrariesDocumentsGetTextContentV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsGetTextContentV1) | **Get** /v1/libraries/{library_id}/documents/{document_id}/text_content | Retrieve the text content of a specific document.
[**LibrariesDocumentsGetV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsGetV1) | **Get** /v1/libraries/{library_id}/documents/{document_id} | Retrieve the metadata of a specific document.
[**LibrariesDocumentsListV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsListV1) | **Get** /v1/libraries/{library_id}/documents | List document in a given library.
[**LibrariesDocumentsReprocessV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsReprocessV1) | **Post** /v1/libraries/{library_id}/documents/{document_id}/reprocess | Reprocess a document.
[**LibrariesDocumentsUpdateV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsUpdateV1) | **Put** /v1/libraries/{library_id}/documents/{document_id} | Update the metadata of a specific document.
[**LibrariesDocumentsUploadV1**](BetaLibrariesDocumentsAPI.md#LibrariesDocumentsUploadV1) | **Post** /v1/libraries/{library_id}/documents | Upload a new document.



## LibrariesDocumentsDeleteV1

> LibrariesDocumentsDeleteV1(ctx, libraryId, documentId).Execute()

Delete a document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsDeleteV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsGetExtractedTextSignedUrlV1

> string LibrariesDocumentsGetExtractedTextSignedUrlV1(ctx, libraryId, documentId).Execute()

Retrieve the signed URL of text extracted from a given document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsGetExtractedTextSignedUrlV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetExtractedTextSignedUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsGetExtractedTextSignedUrlV1`: string
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetExtractedTextSignedUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsGetExtractedTextSignedUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsGetSignedUrlV1

> string LibrariesDocumentsGetSignedUrlV1(ctx, libraryId, documentId).Execute()

Retrieve the signed URL of a specific document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsGetSignedUrlV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetSignedUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsGetSignedUrlV1`: string
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetSignedUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsGetSignedUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsGetStatusV1

> ProcessingStatusOut LibrariesDocumentsGetStatusV1(ctx, libraryId, documentId).Execute()

Retrieve the processing status of a specific document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsGetStatusV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsGetStatusV1`: ProcessingStatusOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetStatusV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsGetStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProcessingStatusOut**](ProcessingStatusOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsGetTextContentV1

> DocumentTextContent LibrariesDocumentsGetTextContentV1(ctx, libraryId, documentId).Execute()

Retrieve the text content of a specific document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsGetTextContentV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetTextContentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsGetTextContentV1`: DocumentTextContent
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetTextContentV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsGetTextContentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DocumentTextContent**](DocumentTextContent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsGetV1

> DocumentOut LibrariesDocumentsGetV1(ctx, libraryId, documentId).Execute()

Retrieve the metadata of a specific document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsGetV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsGetV1`: DocumentOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsGetV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsGetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DocumentOut**](DocumentOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsListV1

> ListDocumentOut LibrariesDocumentsListV1(ctx, libraryId).Search(search).PageSize(pageSize).Page(page).SortBy(sortBy).SortOrder(sortOrder).Execute()

List document in a given library.



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
	search := "search_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 0)
	sortBy := "sortBy_example" // string |  (optional) (default to "created_at")
	sortOrder := "sortOrder_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsListV1(context.Background(), libraryId).Search(search).PageSize(pageSize).Page(page).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsListV1`: ListDocumentOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **pageSize** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 0]
 **sortBy** | **string** |  | [default to &quot;created_at&quot;]
 **sortOrder** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**ListDocumentOut**](ListDocumentOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsReprocessV1

> LibrariesDocumentsReprocessV1(ctx, libraryId, documentId).Execute()

Reprocess a document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsReprocessV1(context.Background(), libraryId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsReprocessV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsReprocessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsUpdateV1

> DocumentOut LibrariesDocumentsUpdateV1(ctx, libraryId, documentId).DocumentUpdateIn(documentUpdateIn).Execute()

Update the metadata of a specific document.



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
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	documentUpdateIn := *openapiclient.NewDocumentUpdateIn() // DocumentUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsUpdateV1(context.Background(), libraryId, documentId).DocumentUpdateIn(documentUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsUpdateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsUpdateV1`: DocumentOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsUpdateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsUpdateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentUpdateIn** | [**DocumentUpdateIn**](DocumentUpdateIn.md) |  | 

### Return type

[**DocumentOut**](DocumentOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDocumentsUploadV1

> DocumentOut LibrariesDocumentsUploadV1(ctx, libraryId).File(file).Execute()

Upload a new document.



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
	file := os.NewFile(1234, "some_file") // *os.File | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  ```bash  file=@path/to/your/file.jsonl;filename=custom_name.jsonl  ```  Otherwise, you can just keep the original file name:  ```bash  file=@path/to/your/file.jsonl  ```

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesDocumentsAPI.LibrariesDocumentsUploadV1(context.Background(), libraryId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesDocumentsAPI.LibrariesDocumentsUploadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDocumentsUploadV1`: DocumentOut
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesDocumentsAPI.LibrariesDocumentsUploadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDocumentsUploadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl;filename&#x3D;custom_name.jsonl  &#x60;&#x60;&#x60;  Otherwise, you can just keep the original file name:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl  &#x60;&#x60;&#x60; | 

### Return type

[**DocumentOut**](DocumentOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

