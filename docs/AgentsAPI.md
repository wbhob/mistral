# \AgentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsCompletionV1AgentsCompletionsPost**](AgentsAPI.md#AgentsCompletionV1AgentsCompletionsPost) | **Post** /v1/agents/completions | Agents Completion



## AgentsCompletionV1AgentsCompletionsPost

> ChatCompletionResponse AgentsCompletionV1AgentsCompletionsPost(ctx).AgentsCompletionRequest(agentsCompletionRequest).Execute()

Agents Completion

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
	agentsCompletionRequest := *openapiclient.NewAgentsCompletionRequest([]openapiclient.MessagesInner{openapiclient.Messages_inner{AssistantMessage: openapiclient.NewAssistantMessage()}}, "AgentId_example") // AgentsCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsCompletionV1AgentsCompletionsPost(context.Background()).AgentsCompletionRequest(agentsCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsCompletionV1AgentsCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsCompletionV1AgentsCompletionsPost`: ChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsCompletionV1AgentsCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsCompletionV1AgentsCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsCompletionRequest** | [**AgentsCompletionRequest**](AgentsCompletionRequest.md) |  | 

### Return type

[**ChatCompletionResponse**](ChatCompletionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

