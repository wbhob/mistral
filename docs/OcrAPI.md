# \OcrAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OcrV1OcrPost**](OcrAPI.md#OcrV1OcrPost) | **Post** /v1/ocr | OCR



## OcrV1OcrPost

> OCRResponse OcrV1OcrPost(ctx).OCRRequest(oCRRequest).Execute()

OCR

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
	oCRRequest := *openapiclient.NewOCRRequest("Model_example", *openapiclient.NewDocument("FileId_example", "DocumentUrl_example", *openapiclient.NewImageUrl("Url_example"))) // OCRRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OcrAPI.OcrV1OcrPost(context.Background()).OCRRequest(oCRRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OcrAPI.OcrV1OcrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OcrV1OcrPost`: OCRResponse
	fmt.Fprintf(os.Stdout, "Response from `OcrAPI.OcrV1OcrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcrV1OcrPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRequest** | [**OCRRequest**](OCRRequest.md) |  | 

### Return type

[**OCRResponse**](OCRResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

