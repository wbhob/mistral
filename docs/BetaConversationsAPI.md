# \BetaConversationsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsApiV1ConversationsAppend**](BetaConversationsAPI.md#AgentsApiV1ConversationsAppend) | **Post** /v1/conversations/{conversation_id} | Append new entries to an existing conversation.
[**AgentsApiV1ConversationsAppendStream**](BetaConversationsAPI.md#AgentsApiV1ConversationsAppendStream) | **Post** /v1/conversations/{conversation_id}#stream | Append new entries to an existing conversation.
[**AgentsApiV1ConversationsGet**](BetaConversationsAPI.md#AgentsApiV1ConversationsGet) | **Get** /v1/conversations/{conversation_id} | Retrieve a conversation information.
[**AgentsApiV1ConversationsHistory**](BetaConversationsAPI.md#AgentsApiV1ConversationsHistory) | **Get** /v1/conversations/{conversation_id}/history | Retrieve all entries in a conversation.
[**AgentsApiV1ConversationsList**](BetaConversationsAPI.md#AgentsApiV1ConversationsList) | **Get** /v1/conversations | List all created conversations.
[**AgentsApiV1ConversationsMessages**](BetaConversationsAPI.md#AgentsApiV1ConversationsMessages) | **Get** /v1/conversations/{conversation_id}/messages | Retrieve all messages in a conversation.
[**AgentsApiV1ConversationsRestart**](BetaConversationsAPI.md#AgentsApiV1ConversationsRestart) | **Post** /v1/conversations/{conversation_id}/restart | Restart a conversation starting from a given entry.
[**AgentsApiV1ConversationsRestartStream**](BetaConversationsAPI.md#AgentsApiV1ConversationsRestartStream) | **Post** /v1/conversations/{conversation_id}/restart#stream | Restart a conversation starting from a given entry.
[**AgentsApiV1ConversationsStart**](BetaConversationsAPI.md#AgentsApiV1ConversationsStart) | **Post** /v1/conversations | Create a conversation and append entries to it.
[**AgentsApiV1ConversationsStartStream**](BetaConversationsAPI.md#AgentsApiV1ConversationsStartStream) | **Post** /v1/conversations#stream | Create a conversation and append entries to it.



## AgentsApiV1ConversationsAppend

> ConversationResponse AgentsApiV1ConversationsAppend(ctx, conversationId).ConversationAppendRequest(conversationAppendRequest).Execute()

Append new entries to an existing conversation.



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
	conversationId := "conversationId_example" // string | ID of the conversation to which we append entries.
	conversationAppendRequest := *openapiclient.NewConversationAppendRequest(*openapiclient.NewConversationInputs()) // ConversationAppendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsAppend(context.Background(), conversationId).ConversationAppendRequest(conversationAppendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsAppend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsAppend`: ConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsAppend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | ID of the conversation to which we append entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsAppendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conversationAppendRequest** | [**ConversationAppendRequest**](ConversationAppendRequest.md) |  | 

### Return type

[**ConversationResponse**](ConversationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsAppendStream

> ConversationEvents AgentsApiV1ConversationsAppendStream(ctx, conversationId).ConversationAppendStreamRequest(conversationAppendStreamRequest).Execute()

Append new entries to an existing conversation.



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
	conversationId := "conversationId_example" // string | ID of the conversation to which we append entries.
	conversationAppendStreamRequest := *openapiclient.NewConversationAppendStreamRequest(*openapiclient.NewConversationInputs()) // ConversationAppendStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsAppendStream(context.Background(), conversationId).ConversationAppendStreamRequest(conversationAppendStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsAppendStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsAppendStream`: ConversationEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsAppendStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | ID of the conversation to which we append entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsAppendStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conversationAppendStreamRequest** | [**ConversationAppendStreamRequest**](ConversationAppendStreamRequest.md) |  | 

### Return type

[**ConversationEvents**](ConversationEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsGet

> ResponseV1ConversationsGet AgentsApiV1ConversationsGet(ctx, conversationId).Execute()

Retrieve a conversation information.



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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsGet(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsGet`: ResponseV1ConversationsGet
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseV1ConversationsGet**](ResponseV1ConversationsGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsHistory

> ConversationHistory AgentsApiV1ConversationsHistory(ctx, conversationId).Execute()

Retrieve all entries in a conversation.



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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsHistory(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsHistory`: ConversationHistory
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConversationHistory**](ConversationHistory.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsList

> []AgentsApiV1ConversationsList200ResponseInner AgentsApiV1ConversationsList(ctx).Page(page).PageSize(pageSize).Execute()

List all created conversations.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsList(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsList`: []AgentsApiV1ConversationsList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]

### Return type

[**[]AgentsApiV1ConversationsList200ResponseInner**](AgentsApiV1ConversationsList200ResponseInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsMessages

> ConversationMessages AgentsApiV1ConversationsMessages(ctx, conversationId).Execute()

Retrieve all messages in a conversation.



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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsMessages(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsMessages`: ConversationMessages
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConversationMessages**](ConversationMessages.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsRestart

> ConversationResponse AgentsApiV1ConversationsRestart(ctx, conversationId).ConversationRestartRequest(conversationRestartRequest).Execute()

Restart a conversation starting from a given entry.



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
	conversationId := "conversationId_example" // string | 
	conversationRestartRequest := *openapiclient.NewConversationRestartRequest(*openapiclient.NewConversationInputs(), "FromEntryId_example") // ConversationRestartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsRestart(context.Background(), conversationId).ConversationRestartRequest(conversationRestartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsRestart`: ConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conversationRestartRequest** | [**ConversationRestartRequest**](ConversationRestartRequest.md) |  | 

### Return type

[**ConversationResponse**](ConversationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsRestartStream

> ConversationEvents AgentsApiV1ConversationsRestartStream(ctx, conversationId).ConversationRestartStreamRequest(conversationRestartStreamRequest).Execute()

Restart a conversation starting from a given entry.



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
	conversationId := "conversationId_example" // string | 
	conversationRestartStreamRequest := *openapiclient.NewConversationRestartStreamRequest(*openapiclient.NewConversationInputs(), "FromEntryId_example") // ConversationRestartStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsRestartStream(context.Background(), conversationId).ConversationRestartStreamRequest(conversationRestartStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsRestartStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsRestartStream`: ConversationEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsRestartStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsRestartStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conversationRestartStreamRequest** | [**ConversationRestartStreamRequest**](ConversationRestartStreamRequest.md) |  | 

### Return type

[**ConversationEvents**](ConversationEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsStart

> ConversationResponse AgentsApiV1ConversationsStart(ctx).ConversationRequest(conversationRequest).Execute()

Create a conversation and append entries to it.



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
	conversationRequest := *openapiclient.NewConversationRequest(*openapiclient.NewConversationInputs()) // ConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsStart(context.Background()).ConversationRequest(conversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsStart`: ConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversationRequest** | [**ConversationRequest**](ConversationRequest.md) |  | 

### Return type

[**ConversationResponse**](ConversationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1ConversationsStartStream

> ConversationEvents AgentsApiV1ConversationsStartStream(ctx).ConversationStreamRequest(conversationStreamRequest).Execute()

Create a conversation and append entries to it.



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
	conversationStreamRequest := *openapiclient.NewConversationStreamRequest(*openapiclient.NewConversationInputs()) // ConversationStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConversationsAPI.AgentsApiV1ConversationsStartStream(context.Background()).ConversationStreamRequest(conversationStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConversationsAPI.AgentsApiV1ConversationsStartStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1ConversationsStartStream`: ConversationEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaConversationsAPI.AgentsApiV1ConversationsStartStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1ConversationsStartStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversationStreamRequest** | [**ConversationStreamRequest**](ConversationStreamRequest.md) |  | 

### Return type

[**ConversationEvents**](ConversationEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

