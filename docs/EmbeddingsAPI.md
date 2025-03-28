# \EmbeddingsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmbeddingsV1EmbeddingsPost**](EmbeddingsAPI.md#EmbeddingsV1EmbeddingsPost) | **Post** /v1/embeddings | Embeddings



## EmbeddingsV1EmbeddingsPost

> EmbeddingResponse EmbeddingsV1EmbeddingsPost(ctx).EmbeddingRequest(embeddingRequest).Execute()

Embeddings



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
	embeddingRequest := *openapiclient.NewEmbeddingRequest("mistral-embed", *openapiclient.NewInput2()) // EmbeddingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.EmbeddingsV1EmbeddingsPost(context.Background()).EmbeddingRequest(embeddingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.EmbeddingsV1EmbeddingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmbeddingsV1EmbeddingsPost`: EmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.EmbeddingsV1EmbeddingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddingsV1EmbeddingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingRequest** | [**EmbeddingRequest**](EmbeddingRequest.md) |  | 

### Return type

[**EmbeddingResponse**](EmbeddingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

