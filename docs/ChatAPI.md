# \ChatAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatCompletionV1ChatCompletionsPost**](ChatAPI.md#ChatCompletionV1ChatCompletionsPost) | **Post** /v1/chat/completions | Chat Completion



## ChatCompletionV1ChatCompletionsPost

> ChatCompletionResponse ChatCompletionV1ChatCompletionsPost(ctx).ChatCompletionRequest(chatCompletionRequest).Execute()

Chat Completion

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
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest("Model_example", []openapiclient.MessagesInner{openapiclient.Messages_inner{AssistantMessage: openapiclient.NewAssistantMessage()}}) // ChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).ChatCompletionRequest(chatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.ChatCompletionV1ChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatCompletionV1ChatCompletionsPost`: ChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.ChatCompletionV1ChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCompletionV1ChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 

### Return type

[**ChatCompletionResponse**](ChatCompletionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

