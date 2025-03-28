# \FimAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FimCompletionV1FimCompletionsPost**](FimAPI.md#FimCompletionV1FimCompletionsPost) | **Post** /v1/fim/completions | Fim Completion



## FimCompletionV1FimCompletionsPost

> FIMCompletionResponse FimCompletionV1FimCompletionsPost(ctx).FIMCompletionRequest(fIMCompletionRequest).Execute()

Fim Completion



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
	fIMCompletionRequest := *openapiclient.NewFIMCompletionRequest("Model_example", "Prompt_example") // FIMCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FimAPI.FimCompletionV1FimCompletionsPost(context.Background()).FIMCompletionRequest(fIMCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FimAPI.FimCompletionV1FimCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FimCompletionV1FimCompletionsPost`: FIMCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `FimAPI.FimCompletionV1FimCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFimCompletionV1FimCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fIMCompletionRequest** | [**FIMCompletionRequest**](FIMCompletionRequest.md) |  | 

### Return type

[**FIMCompletionResponse**](FIMCompletionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

