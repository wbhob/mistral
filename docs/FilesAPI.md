# \FilesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesApiRoutesDeleteFile**](FilesAPI.md#FilesApiRoutesDeleteFile) | **Delete** /v1/files/{file_id} | Delete File
[**FilesApiRoutesDownloadFile**](FilesAPI.md#FilesApiRoutesDownloadFile) | **Get** /v1/files/{file_id}/content | Download File
[**FilesApiRoutesGetSignedUrl**](FilesAPI.md#FilesApiRoutesGetSignedUrl) | **Get** /v1/files/{file_id}/url | Get Signed Url
[**FilesApiRoutesListFiles**](FilesAPI.md#FilesApiRoutesListFiles) | **Get** /v1/files | List Files
[**FilesApiRoutesRetrieveFile**](FilesAPI.md#FilesApiRoutesRetrieveFile) | **Get** /v1/files/{file_id} | Retrieve File
[**FilesApiRoutesUploadFile**](FilesAPI.md#FilesApiRoutesUploadFile) | **Post** /v1/files | Upload File



## FilesApiRoutesDeleteFile

> DeleteFileOut FilesApiRoutesDeleteFile(ctx, fileId).Execute()

Delete File



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesDeleteFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesDeleteFile`: DeleteFileOut
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFileOut**](DeleteFileOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiRoutesDownloadFile

> *os.File FilesApiRoutesDownloadFile(ctx, fileId).Execute()

Download File



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesDownloadFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesDownloadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesDownloadFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesDownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiRoutesGetSignedUrl

> FileSignedURL FilesApiRoutesGetSignedUrl(ctx, fileId).Expiry(expiry).Execute()

Get Signed Url

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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	expiry := int32(56) // int32 | Number of hours before the url becomes invalid. Defaults to 24h (optional) (default to 24)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesGetSignedUrl(context.Background(), fileId).Expiry(expiry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesGetSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesGetSignedUrl`: FileSignedURL
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesGetSignedUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesGetSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiry** | **int32** | Number of hours before the url becomes invalid. Defaults to 24h | [default to 24]

### Return type

[**FileSignedURL**](FileSignedURL.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiRoutesListFiles

> ListFilesOut FilesApiRoutesListFiles(ctx).Page(page).PageSize(pageSize).SampleType(sampleType).Source(source).Search(search).Purpose(purpose).Execute()

List Files



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
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	sampleType := []openapiclient.SampleType{openapiclient.SampleType("pretrain")} // []SampleType |  (optional)
	source := []openapiclient.Source{openapiclient.Source("upload")} // []Source |  (optional)
	search := "search_example" // string |  (optional)
	purpose := openapiclient.FilePurpose("fine-tune") // FilePurpose |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesListFiles(context.Background()).Page(page).PageSize(pageSize).SampleType(sampleType).Source(source).Search(search).Purpose(purpose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesListFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesListFiles`: ListFilesOut
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesListFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sampleType** | [**[]SampleType**](SampleType.md) |  | 
 **source** | [**[]Source**](Source.md) |  | 
 **search** | **string** |  | 
 **purpose** | [**FilePurpose**](FilePurpose.md) |  | 

### Return type

[**ListFilesOut**](ListFilesOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiRoutesRetrieveFile

> RetrieveFileOut FilesApiRoutesRetrieveFile(ctx, fileId).Execute()

Retrieve File



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesRetrieveFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesRetrieveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesRetrieveFile`: RetrieveFileOut
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesRetrieveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesRetrieveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveFileOut**](RetrieveFileOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiRoutesUploadFile

> UploadFileOut FilesApiRoutesUploadFile(ctx).File(file).Purpose(purpose).Execute()

Upload File



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
	file := os.NewFile(1234, "some_file") // *os.File | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  ```bash  file=@path/to/your/file.jsonl;filename=custom_name.jsonl  ```  Otherwise, you can just keep the original file name:  ```bash  file=@path/to/your/file.jsonl  ```
	purpose := openapiclient.FilePurpose("fine-tune") // FilePurpose |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesApiRoutesUploadFile(context.Background()).File(file).Purpose(purpose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesApiRoutesUploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesApiRoutesUploadFile`: UploadFileOut
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesApiRoutesUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiRoutesUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl;filename&#x3D;custom_name.jsonl  &#x60;&#x60;&#x60;  Otherwise, you can just keep the original file name:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl  &#x60;&#x60;&#x60; | 
 **purpose** | [**FilePurpose**](FilePurpose.md) |  | 

### Return type

[**UploadFileOut**](UploadFileOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

