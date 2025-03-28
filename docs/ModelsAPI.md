# \ModelsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteModelV1ModelsModelIdDelete**](ModelsAPI.md#DeleteModelV1ModelsModelIdDelete) | **Delete** /v1/models/{model_id} | Delete Model
[**JobsApiRoutesFineTuningArchiveFineTunedModel**](ModelsAPI.md#JobsApiRoutesFineTuningArchiveFineTunedModel) | **Post** /v1/fine_tuning/models/{model_id}/archive | Archive Fine Tuned Model
[**JobsApiRoutesFineTuningUnarchiveFineTunedModel**](ModelsAPI.md#JobsApiRoutesFineTuningUnarchiveFineTunedModel) | **Delete** /v1/fine_tuning/models/{model_id}/archive | Unarchive Fine Tuned Model
[**JobsApiRoutesFineTuningUpdateFineTunedModel**](ModelsAPI.md#JobsApiRoutesFineTuningUpdateFineTunedModel) | **Patch** /v1/fine_tuning/models/{model_id} | Update Fine Tuned Model
[**ListModelsV1ModelsGet**](ModelsAPI.md#ListModelsV1ModelsGet) | **Get** /v1/models | List Models
[**RetrieveModelV1ModelsModelIdGet**](ModelsAPI.md#RetrieveModelV1ModelsModelIdGet) | **Get** /v1/models/{model_id} | Retrieve Model



## DeleteModelV1ModelsModelIdDelete

> DeleteModelOut DeleteModelV1ModelsModelIdDelete(ctx, modelId).Execute()

Delete Model



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
	modelId := "ft:open-mistral-7b:587a6b29:20240514:7e773925" // string | The ID of the model to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.DeleteModelV1ModelsModelIdDelete(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.DeleteModelV1ModelsModelIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModelV1ModelsModelIdDelete`: DeleteModelOut
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.DeleteModelV1ModelsModelIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | The ID of the model to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModelV1ModelsModelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteModelOut**](DeleteModelOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningArchiveFineTunedModel

> ArchiveFTModelOut JobsApiRoutesFineTuningArchiveFineTunedModel(ctx, modelId).Execute()

Archive Fine Tuned Model



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
	modelId := "ft:open-mistral-7b:587a6b29:20240514:7e773925" // string | The ID of the model to archive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningArchiveFineTunedModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.JobsApiRoutesFineTuningArchiveFineTunedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningArchiveFineTunedModel`: ArchiveFTModelOut
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.JobsApiRoutesFineTuningArchiveFineTunedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | The ID of the model to archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningArchiveFineTunedModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArchiveFTModelOut**](ArchiveFTModelOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningUnarchiveFineTunedModel

> UnarchiveFTModelOut JobsApiRoutesFineTuningUnarchiveFineTunedModel(ctx, modelId).Execute()

Unarchive Fine Tuned Model



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
	modelId := "ft:open-mistral-7b:587a6b29:20240514:7e773925" // string | The ID of the model to unarchive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningUnarchiveFineTunedModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.JobsApiRoutesFineTuningUnarchiveFineTunedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningUnarchiveFineTunedModel`: UnarchiveFTModelOut
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.JobsApiRoutesFineTuningUnarchiveFineTunedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | The ID of the model to unarchive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningUnarchiveFineTunedModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnarchiveFTModelOut**](UnarchiveFTModelOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningUpdateFineTunedModel

> FTModelOut JobsApiRoutesFineTuningUpdateFineTunedModel(ctx, modelId).UpdateFTModelIn(updateFTModelIn).Execute()

Update Fine Tuned Model



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
	modelId := "ft:open-mistral-7b:587a6b29:20240514:7e773925" // string | The ID of the model to update.
	updateFTModelIn := *openapiclient.NewUpdateFTModelIn() // UpdateFTModelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningUpdateFineTunedModel(context.Background(), modelId).UpdateFTModelIn(updateFTModelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.JobsApiRoutesFineTuningUpdateFineTunedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningUpdateFineTunedModel`: FTModelOut
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.JobsApiRoutesFineTuningUpdateFineTunedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | The ID of the model to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningUpdateFineTunedModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFTModelIn** | [**UpdateFTModelIn**](UpdateFTModelIn.md) |  | 

### Return type

[**FTModelOut**](FTModelOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelsV1ModelsGet

> ModelList ListModelsV1ModelsGet(ctx).Execute()

List Models



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.ListModelsV1ModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ListModelsV1ModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelsV1ModelsGet`: ModelList
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ListModelsV1ModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListModelsV1ModelsGetRequest struct via the builder pattern


### Return type

[**ModelList**](ModelList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveModelV1ModelsModelIdGet

> ResponseRetrieveModelV1ModelsModelIdGet RetrieveModelV1ModelsModelIdGet(ctx, modelId).Execute()

Retrieve Model



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
	modelId := "ft:open-mistral-7b:587a6b29:20240514:7e773925" // string | The ID of the model to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.RetrieveModelV1ModelsModelIdGet(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.RetrieveModelV1ModelsModelIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveModelV1ModelsModelIdGet`: ResponseRetrieveModelV1ModelsModelIdGet
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.RetrieveModelV1ModelsModelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | The ID of the model to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveModelV1ModelsModelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseRetrieveModelV1ModelsModelIdGet**](ResponseRetrieveModelV1ModelsModelIdGet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

