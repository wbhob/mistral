# \FineTuningAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobsApiRoutesFineTuningCancelFineTuningJob**](FineTuningAPI.md#JobsApiRoutesFineTuningCancelFineTuningJob) | **Post** /v1/fine_tuning/jobs/{job_id}/cancel | Cancel Fine Tuning Job
[**JobsApiRoutesFineTuningCreateFineTuningJob**](FineTuningAPI.md#JobsApiRoutesFineTuningCreateFineTuningJob) | **Post** /v1/fine_tuning/jobs | Create Fine Tuning Job
[**JobsApiRoutesFineTuningGetFineTuningJob**](FineTuningAPI.md#JobsApiRoutesFineTuningGetFineTuningJob) | **Get** /v1/fine_tuning/jobs/{job_id} | Get Fine Tuning Job
[**JobsApiRoutesFineTuningGetFineTuningJobs**](FineTuningAPI.md#JobsApiRoutesFineTuningGetFineTuningJobs) | **Get** /v1/fine_tuning/jobs | Get Fine Tuning Jobs
[**JobsApiRoutesFineTuningStartFineTuningJob**](FineTuningAPI.md#JobsApiRoutesFineTuningStartFineTuningJob) | **Post** /v1/fine_tuning/jobs/{job_id}/start | Start Fine Tuning Job



## JobsApiRoutesFineTuningCancelFineTuningJob

> Response1 JobsApiRoutesFineTuningCancelFineTuningJob(ctx, jobId).Execute()

Cancel Fine Tuning Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the job to cancel.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.JobsApiRoutesFineTuningCancelFineTuningJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.JobsApiRoutesFineTuningCancelFineTuningJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningCancelFineTuningJob`: Response1
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.JobsApiRoutesFineTuningCancelFineTuningJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningCancelFineTuningJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response1**](Response1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningCreateFineTuningJob

> Response JobsApiRoutesFineTuningCreateFineTuningJob(ctx).JobIn(jobIn).DryRun(dryRun).Execute()

Create Fine Tuning Job



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
	jobIn := *openapiclient.NewJobIn(openapiclient.FineTuneableModel("open-mistral-7b"), *openapiclient.NewHyperparameters()) // JobIn | 
	dryRun := true // bool | * If `true` the job is not spawned, instead the query returns a handful of useful metadata   for the user to perform sanity checks (see `LegacyJobMetadataOut` response). * Otherwise, the job is started and the query returns the job ID along with some of the   input parameters (see `JobOut` response).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.JobsApiRoutesFineTuningCreateFineTuningJob(context.Background()).JobIn(jobIn).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.JobsApiRoutesFineTuningCreateFineTuningJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningCreateFineTuningJob`: Response
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.JobsApiRoutesFineTuningCreateFineTuningJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningCreateFineTuningJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIn** | [**JobIn**](JobIn.md) |  | 
 **dryRun** | **bool** | * If &#x60;true&#x60; the job is not spawned, instead the query returns a handful of useful metadata   for the user to perform sanity checks (see &#x60;LegacyJobMetadataOut&#x60; response). * Otherwise, the job is started and the query returns the job ID along with some of the   input parameters (see &#x60;JobOut&#x60; response).  | 

### Return type

[**Response**](Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningGetFineTuningJob

> Response1 JobsApiRoutesFineTuningGetFineTuningJob(ctx, jobId).Execute()

Get Fine Tuning Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the job to analyse.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningGetFineTuningJob`: Response1
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job to analyse. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningGetFineTuningJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response1**](Response1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningGetFineTuningJobs

> JobsOut JobsApiRoutesFineTuningGetFineTuningJobs(ctx).Page(page).PageSize(pageSize).Model(model).CreatedAfter(createdAfter).CreatedBefore(createdBefore).CreatedByMe(createdByMe).Status(status).WandbProject(wandbProject).WandbName(wandbName).Suffix(suffix).Execute()

Get Fine Tuning Jobs



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
	page := int32(56) // int32 | The page number of the results to be returned. (optional) (default to 0)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)
	model := "model_example" // string | The model name used for fine-tuning to filter on. When set, the other results are not displayed. (optional)
	createdAfter := time.Now() // time.Time | The date/time to filter on. When set, the results for previous creation times are not displayed. (optional)
	createdBefore := time.Now() // time.Time |  (optional)
	createdByMe := true // bool | When set, only return results for jobs created by the API caller. Other results are not displayed. (optional) (default to false)
	status := "status_example" // string | The current job state to filter on. When set, the other results are not displayed. (optional)
	wandbProject := "wandbProject_example" // string | The Weights and Biases project to filter on. When set, the other results are not displayed. (optional)
	wandbName := "wandbName_example" // string | The Weight and Biases run name to filter on. When set, the other results are not displayed. (optional)
	suffix := "suffix_example" // string | The model suffix to filter on. When set, the other results are not displayed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJobs(context.Background()).Page(page).PageSize(pageSize).Model(model).CreatedAfter(createdAfter).CreatedBefore(createdBefore).CreatedByMe(createdByMe).Status(status).WandbProject(wandbProject).WandbName(wandbName).Suffix(suffix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningGetFineTuningJobs`: JobsOut
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.JobsApiRoutesFineTuningGetFineTuningJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningGetFineTuningJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number of the results to be returned. | [default to 0]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]
 **model** | **string** | The model name used for fine-tuning to filter on. When set, the other results are not displayed. | 
 **createdAfter** | **time.Time** | The date/time to filter on. When set, the results for previous creation times are not displayed. | 
 **createdBefore** | **time.Time** |  | 
 **createdByMe** | **bool** | When set, only return results for jobs created by the API caller. Other results are not displayed. | [default to false]
 **status** | **string** | The current job state to filter on. When set, the other results are not displayed. | 
 **wandbProject** | **string** | The Weights and Biases project to filter on. When set, the other results are not displayed. | 
 **wandbName** | **string** | The Weight and Biases run name to filter on. When set, the other results are not displayed. | 
 **suffix** | **string** | The model suffix to filter on. When set, the other results are not displayed. | 

### Return type

[**JobsOut**](JobsOut.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsApiRoutesFineTuningStartFineTuningJob

> Response1 JobsApiRoutesFineTuningStartFineTuningJob(ctx, jobId).Execute()

Start Fine Tuning Job



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
	resp, r, err := apiClient.FineTuningAPI.JobsApiRoutesFineTuningStartFineTuningJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.JobsApiRoutesFineTuningStartFineTuningJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobsApiRoutesFineTuningStartFineTuningJob`: Response1
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.JobsApiRoutesFineTuningStartFineTuningJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsApiRoutesFineTuningStartFineTuningJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response1**](Response1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

