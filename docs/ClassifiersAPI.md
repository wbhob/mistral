# \ClassifiersAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatClassificationsV1ChatClassificationsPost**](ClassifiersAPI.md#ChatClassificationsV1ChatClassificationsPost) | **Post** /v1/chat/classifications | Chat Classifications
[**ChatModerationsV1ChatModerationsPost**](ClassifiersAPI.md#ChatModerationsV1ChatModerationsPost) | **Post** /v1/chat/moderations | Chat Moderations
[**ClassificationsV1ClassificationsPost**](ClassifiersAPI.md#ClassificationsV1ClassificationsPost) | **Post** /v1/classifications | Classifications
[**ModerationsV1ModerationsPost**](ClassifiersAPI.md#ModerationsV1ModerationsPost) | **Post** /v1/moderations | Moderations



## ChatClassificationsV1ChatClassificationsPost

> ClassificationResponse ChatClassificationsV1ChatClassificationsPost(ctx).ChatClassificationRequest(chatClassificationRequest).Execute()

Chat Classifications

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
	chatClassificationRequest := *openapiclient.NewChatClassificationRequest("Model_example", *openapiclient.NewChatClassificationRequestInputs([]openapiclient.MessagesInner{openapiclient.Messages_inner{AssistantMessage: openapiclient.NewAssistantMessage()}})) // ChatClassificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassifiersAPI.ChatClassificationsV1ChatClassificationsPost(context.Background()).ChatClassificationRequest(chatClassificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassifiersAPI.ChatClassificationsV1ChatClassificationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatClassificationsV1ChatClassificationsPost`: ClassificationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClassifiersAPI.ChatClassificationsV1ChatClassificationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatClassificationsV1ChatClassificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatClassificationRequest** | [**ChatClassificationRequest**](ChatClassificationRequest.md) |  | 

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


## ChatModerationsV1ChatModerationsPost

> ModerationResponse ChatModerationsV1ChatModerationsPost(ctx).ChatModerationRequest(chatModerationRequest).Execute()

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
	chatModerationRequest := *openapiclient.NewChatModerationRequest(*openapiclient.NewInput(), "Model_example") // ChatModerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassifiersAPI.ChatModerationsV1ChatModerationsPost(context.Background()).ChatModerationRequest(chatModerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassifiersAPI.ChatModerationsV1ChatModerationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatModerationsV1ChatModerationsPost`: ModerationResponse
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

[**ModerationResponse**](ModerationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClassificationsV1ClassificationsPost

> ClassificationResponse ClassificationsV1ClassificationsPost(ctx).ClassificationRequest(classificationRequest).Execute()

Classifications

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
	resp, r, err := apiClient.ClassifiersAPI.ClassificationsV1ClassificationsPost(context.Background()).ClassificationRequest(classificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassifiersAPI.ClassificationsV1ClassificationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClassificationsV1ClassificationsPost`: ClassificationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClassifiersAPI.ClassificationsV1ClassificationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClassificationsV1ClassificationsPostRequest struct via the builder pattern


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


## ModerationsV1ModerationsPost

> ModerationResponse ModerationsV1ModerationsPost(ctx).ClassificationRequest(classificationRequest).Execute()

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
	// response from `ModerationsV1ModerationsPost`: ModerationResponse
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

[**ModerationResponse**](ModerationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

