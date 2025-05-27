# \BetaAgentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsApiV1AgentsCreate**](BetaAgentsAPI.md#AgentsApiV1AgentsCreate) | **Post** /v1/agents | Create a agent that can be used within a conversation.
[**AgentsApiV1AgentsGet**](BetaAgentsAPI.md#AgentsApiV1AgentsGet) | **Get** /v1/agents/{agent_id} | Retrieve an agent entity.
[**AgentsApiV1AgentsList**](BetaAgentsAPI.md#AgentsApiV1AgentsList) | **Get** /v1/agents | List agent entities.
[**AgentsApiV1AgentsUpdate**](BetaAgentsAPI.md#AgentsApiV1AgentsUpdate) | **Patch** /v1/agents/{agent_id} | Update an agent entity.
[**AgentsApiV1AgentsUpdateVersion**](BetaAgentsAPI.md#AgentsApiV1AgentsUpdateVersion) | **Patch** /v1/agents/{agent_id}/version | Update an agent version.



## AgentsApiV1AgentsCreate

> Agent AgentsApiV1AgentsCreate(ctx).AgentCreationRequest(agentCreationRequest).Execute()

Create a agent that can be used within a conversation.



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
	agentCreationRequest := *openapiclient.NewAgentCreationRequest("Model_example", "Name_example") // AgentCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsCreate(context.Background()).AgentCreationRequest(agentCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsCreate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentCreationRequest** | [**AgentCreationRequest**](AgentCreationRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsGet

> Agent AgentsApiV1AgentsGet(ctx, agentId).Execute()

Retrieve an agent entity.



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
	agentId := "agentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsGet(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsGet`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsList

> []Agent AgentsApiV1AgentsList(ctx).Page(page).PageSize(pageSize).Execute()

List agent entities.



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
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsList(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsList`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**[]Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsUpdate

> Agent AgentsApiV1AgentsUpdate(ctx, agentId).AgentUpdateRequest(agentUpdateRequest).Execute()

Update an agent entity.



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
	agentId := "agentId_example" // string | 
	agentUpdateRequest := *openapiclient.NewAgentUpdateRequest() // AgentUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsUpdate(context.Background(), agentId).AgentUpdateRequest(agentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsUpdate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentUpdateRequest** | [**AgentUpdateRequest**](AgentUpdateRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsUpdateVersion

> Agent AgentsApiV1AgentsUpdateVersion(ctx, agentId).Version(version).Execute()

Update an agent version.



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
	agentId := "agentId_example" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion(context.Background(), agentId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsUpdateVersion`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

