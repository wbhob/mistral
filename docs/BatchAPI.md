# \BatchAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobsApiRoutesBatchCancelBatchJob**](BatchAPI.md#JobsApiRoutesBatchCancelBatchJob) | **Post** /v1/batch/jobs/{job_id}/cancel | Cancel Batch Job
[**JobsApiRoutesBatchCreateBatchJob**](BatchAPI.md#JobsApiRoutesBatchCreateBatchJob) | **Post** /v1/batch/jobs | Create Batch Job
[**JobsApiRoutesBatchGetBatchJob**](BatchAPI.md#JobsApiRoutesBatchGetBatchJob) | **Get** /v1/batch/jobs/{job_id} | Get Batch Job
[**JobsApiRoutesBatchGetBatchJobs**](BatchAPI.md#JobsApiRoutesBatchGetBatchJobs) | **Get** /v1/batch/jobs | Get Batch Jobs



## JobsApiRoutesBatchCancelBatchJob

> BatchJobOut JobsApiRoutesBatchCancelBatchJob(ctx, jobId).Execute()

Cancel Batch Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.JobsApiRoutesBatchCancelBatchJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.JobsApiRoutesBatchCancelBatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesBatchCancelBatchJob`: BatchJobOut
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.JobsApiRoutesBatchCancelBatchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesBatchCancelBatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchJobOut**](BatchJobOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesBatchCreateBatchJob

> BatchJobOut JobsApiRoutesBatchCreateBatchJob(ctx).BatchJobIn(batchJobIn).Execute()

Create Batch Job



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
	batchJobIn := *openapiclient.NewBatchJobIn([]string{"InputFiles_example"}, openapiclient.ApiEndpoint("/v1/chat/completions"), "Model_example") // BatchJobIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.JobsApiRoutesBatchCreateBatchJob(context.Background()).BatchJobIn(batchJobIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.JobsApiRoutesBatchCreateBatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesBatchCreateBatchJob`: BatchJobOut
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.JobsApiRoutesBatchCreateBatchJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesBatchCreateBatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchJobIn** | [**BatchJobIn**](BatchJobIn.md) |  | 

### Return type

[**BatchJobOut**](BatchJobOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesBatchGetBatchJob

> BatchJobOut JobsApiRoutesBatchGetBatchJob(ctx, jobId).Execute()

Get Batch Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.JobsApiRoutesBatchGetBatchJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.JobsApiRoutesBatchGetBatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesBatchGetBatchJob`: BatchJobOut
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.JobsApiRoutesBatchGetBatchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesBatchGetBatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchJobOut**](BatchJobOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesBatchGetBatchJobs

> BatchJobsOut JobsApiRoutesBatchGetBatchJobs(ctx).Page(page).PageSize(pageSize).Model(model).Metadata(metadata).CreatedAfter(createdAfter).CreatedByMe(createdByMe).Status(status).Execute()

Get Batch Jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/wbhob/mistral"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	model := "model_example" // string |  (optional)
	metadata := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)
	createdAfter := time.Now() // time.Time |  (optional)
	createdByMe := true // bool |  (optional) (default to false)
	status := []openapiclient.BatchJobStatus{openapiclient.BatchJobStatus("QUEUED")} // []BatchJobStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.JobsApiRoutesBatchGetBatchJobs(context.Background()).Page(page).PageSize(pageSize).Model(model).Metadata(metadata).CreatedAfter(createdAfter).CreatedByMe(createdByMe).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.JobsApiRoutesBatchGetBatchJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesBatchGetBatchJobs`: BatchJobsOut
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.JobsApiRoutesBatchGetBatchJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesBatchGetBatchJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **model** | **string** |  | 
 **metadata** | **map[string]interface{}** |  | 
 **createdAfter** | **time.Time** |  | 
 **createdByMe** | **bool** |  | [default to false]
 **status** | [**[]BatchJobStatus**](BatchJobStatus.md) |  | 

### Return type

[**BatchJobsOut**](BatchJobsOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

