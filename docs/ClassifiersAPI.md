# \ClassifiersAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatModerationsV1ChatModerationsPost**](ClassifiersAPI.md#ChatModerationsV1ChatModerationsPost) | **Post** /v1/chat/moderations | Chat Moderations
[**ModerationsV1ModerationsPost**](ClassifiersAPI.md#ModerationsV1ModerationsPost) | **Post** /v1/moderations | Moderations



## ChatModerationsV1ChatModerationsPost

> ClassificationResponse ChatModerationsV1ChatModerationsPost(ctx).ChatModerationRequest(chatModerationRequest).Execute()

Chat Moderations

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
	chatModerationRequest := *openapiclient.NewChatModerationRequest("Model_example", *openapiclient.NewInput()) // ChatModerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassifiersAPI.ChatModerationsV1ChatModerationsPost(context.Background()).ChatModerationRequest(chatModerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassifiersAPI.ChatModerationsV1ChatModerationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatModerationsV1ChatModerationsPost`: ClassificationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClassifiersAPI.ChatModerationsV1ChatModerationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatModerationsV1ChatModerationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatModerationRequest** | [**ChatModerationRequest**](ChatModerationRequest.md) |  | 

### Return type

[**ClassificationResponse**](ClassificationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModerationsV1ModerationsPost

> ClassificationResponse ModerationsV1ModerationsPost(ctx).ClassificationRequest(classificationRequest).Execute()

Moderations

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
	classificationRequest := *openapiclient.NewClassificationRequest("Model_example", *openapiclient.NewInput1()) // ClassificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassifiersAPI.ModerationsV1ModerationsPost(context.Background()).ClassificationRequest(classificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassifiersAPI.ModerationsV1ModerationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModerationsV1ModerationsPost`: ClassificationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClassifiersAPI.ModerationsV1ModerationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModerationsV1ModerationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classificationRequest** | [**ClassificationRequest**](ClassificationRequest.md) |  | 

### Return type

[**ClassificationResponse**](ClassificationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

