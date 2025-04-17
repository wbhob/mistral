# Go API client for mistral

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.2
- Package version: 0.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import mistral "github.com/wbhob/mistral"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `mistral.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), mistral.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `mistral.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), mistral.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `mistral.ContextOperationServerIndices` and `mistral.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), mistral.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), mistral.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.mistral.ai*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentsAPI* | [**AgentsCompletionV1AgentsCompletionsPost**](docs/AgentsAPI.md#agentscompletionv1agentscompletionspost) | **Post** /v1/agents/completions | Agents Completion
*BatchAPI* | [**JobsApiRoutesBatchCancelBatchJob**](docs/BatchAPI.md#jobsapiroutesbatchcancelbatchjob) | **Post** /v1/batch/jobs/{job_id}/cancel | Cancel Batch Job
*BatchAPI* | [**JobsApiRoutesBatchCreateBatchJob**](docs/BatchAPI.md#jobsapiroutesbatchcreatebatchjob) | **Post** /v1/batch/jobs | Create Batch Job
*BatchAPI* | [**JobsApiRoutesBatchGetBatchJob**](docs/BatchAPI.md#jobsapiroutesbatchgetbatchjob) | **Get** /v1/batch/jobs/{job_id} | Get Batch Job
*BatchAPI* | [**JobsApiRoutesBatchGetBatchJobs**](docs/BatchAPI.md#jobsapiroutesbatchgetbatchjobs) | **Get** /v1/batch/jobs | Get Batch Jobs
*ChatAPI* | [**ChatCompletionV1ChatCompletionsPost**](docs/ChatAPI.md#chatcompletionv1chatcompletionspost) | **Post** /v1/chat/completions | Chat Completion
*ClassifiersAPI* | [**ChatClassificationsV1ChatClassificationsPost**](docs/ClassifiersAPI.md#chatclassificationsv1chatclassificationspost) | **Post** /v1/chat/classifications | Chat Classifications
*ClassifiersAPI* | [**ChatModerationsV1ChatModerationsPost**](docs/ClassifiersAPI.md#chatmoderationsv1chatmoderationspost) | **Post** /v1/chat/moderations | Chat Moderations
*ClassifiersAPI* | [**ClassificationsV1ClassificationsPost**](docs/ClassifiersAPI.md#classificationsv1classificationspost) | **Post** /v1/classifications | Classifications
*ClassifiersAPI* | [**ModerationsV1ModerationsPost**](docs/ClassifiersAPI.md#moderationsv1moderationspost) | **Post** /v1/moderations | Moderations
*EmbeddingsAPI* | [**EmbeddingsV1EmbeddingsPost**](docs/EmbeddingsAPI.md#embeddingsv1embeddingspost) | **Post** /v1/embeddings | Embeddings
*FilesAPI* | [**FilesApiRoutesDeleteFile**](docs/FilesAPI.md#filesapiroutesdeletefile) | **Delete** /v1/files/{file_id} | Delete File
*FilesAPI* | [**FilesApiRoutesDownloadFile**](docs/FilesAPI.md#filesapiroutesdownloadfile) | **Get** /v1/files/{file_id}/content | Download File
*FilesAPI* | [**FilesApiRoutesGetSignedUrl**](docs/FilesAPI.md#filesapiroutesgetsignedurl) | **Get** /v1/files/{file_id}/url | Get Signed Url
*FilesAPI* | [**FilesApiRoutesListFiles**](docs/FilesAPI.md#filesapirouteslistfiles) | **Get** /v1/files | List Files
*FilesAPI* | [**FilesApiRoutesRetrieveFile**](docs/FilesAPI.md#filesapiroutesretrievefile) | **Get** /v1/files/{file_id} | Retrieve File
*FilesAPI* | [**FilesApiRoutesUploadFile**](docs/FilesAPI.md#filesapiroutesuploadfile) | **Post** /v1/files | Upload File
*FimAPI* | [**FimCompletionV1FimCompletionsPost**](docs/FimAPI.md#fimcompletionv1fimcompletionspost) | **Post** /v1/fim/completions | Fim Completion
*FineTuningAPI* | [**JobsApiRoutesFineTuningCancelFineTuningJob**](docs/FineTuningAPI.md#jobsapiroutesfinetuningcancelfinetuningjob) | **Post** /v1/fine_tuning/jobs/{job_id}/cancel | Cancel Fine Tuning Job
*FineTuningAPI* | [**JobsApiRoutesFineTuningCreateFineTuningJob**](docs/FineTuningAPI.md#jobsapiroutesfinetuningcreatefinetuningjob) | **Post** /v1/fine_tuning/jobs | Create Fine Tuning Job
*FineTuningAPI* | [**JobsApiRoutesFineTuningGetFineTuningJob**](docs/FineTuningAPI.md#jobsapiroutesfinetuninggetfinetuningjob) | **Get** /v1/fine_tuning/jobs/{job_id} | Get Fine Tuning Job
*FineTuningAPI* | [**JobsApiRoutesFineTuningGetFineTuningJobs**](docs/FineTuningAPI.md#jobsapiroutesfinetuninggetfinetuningjobs) | **Get** /v1/fine_tuning/jobs | Get Fine Tuning Jobs
*FineTuningAPI* | [**JobsApiRoutesFineTuningStartFineTuningJob**](docs/FineTuningAPI.md#jobsapiroutesfinetuningstartfinetuningjob) | **Post** /v1/fine_tuning/jobs/{job_id}/start | Start Fine Tuning Job
*ModelsAPI* | [**DeleteModelV1ModelsModelIdDelete**](docs/ModelsAPI.md#deletemodelv1modelsmodeliddelete) | **Delete** /v1/models/{model_id} | Delete Model
*ModelsAPI* | [**JobsApiRoutesFineTuningArchiveFineTunedModel**](docs/ModelsAPI.md#jobsapiroutesfinetuningarchivefinetunedmodel) | **Post** /v1/fine_tuning/models/{model_id}/archive | Archive Fine Tuned Model
*ModelsAPI* | [**JobsApiRoutesFineTuningUnarchiveFineTunedModel**](docs/ModelsAPI.md#jobsapiroutesfinetuningunarchivefinetunedmodel) | **Delete** /v1/fine_tuning/models/{model_id}/archive | Unarchive Fine Tuned Model
*ModelsAPI* | [**JobsApiRoutesFineTuningUpdateFineTunedModel**](docs/ModelsAPI.md#jobsapiroutesfinetuningupdatefinetunedmodel) | **Patch** /v1/fine_tuning/models/{model_id} | Update Fine Tuned Model
*ModelsAPI* | [**ListModelsV1ModelsGet**](docs/ModelsAPI.md#listmodelsv1modelsget) | **Get** /v1/models | List Models
*ModelsAPI* | [**RetrieveModelV1ModelsModelIdGet**](docs/ModelsAPI.md#retrievemodelv1modelsmodelidget) | **Get** /v1/models/{model_id} | Retrieve Model
*OcrAPI* | [**OcrV1OcrPost**](docs/OcrAPI.md#ocrv1ocrpost) | **Post** /v1/ocr | OCR


## Documentation For Models

 - [AgentsCompletionRequest](docs/AgentsCompletionRequest.md)
 - [ApiEndpoint](docs/ApiEndpoint.md)
 - [ArchiveFTModelOut](docs/ArchiveFTModelOut.md)
 - [Arguments](docs/Arguments.md)
 - [AssistantMessage](docs/AssistantMessage.md)
 - [BaseModelCard](docs/BaseModelCard.md)
 - [BatchError](docs/BatchError.md)
 - [BatchJobIn](docs/BatchJobIn.md)
 - [BatchJobOut](docs/BatchJobOut.md)
 - [BatchJobStatus](docs/BatchJobStatus.md)
 - [BatchJobsOut](docs/BatchJobsOut.md)
 - [ChatClassificationRequest](docs/ChatClassificationRequest.md)
 - [ChatClassificationRequestInputs](docs/ChatClassificationRequestInputs.md)
 - [ChatCompletionChoice](docs/ChatCompletionChoice.md)
 - [ChatCompletionRequest](docs/ChatCompletionRequest.md)
 - [ChatCompletionResponse](docs/ChatCompletionResponse.md)
 - [ChatCompletionResponseBase](docs/ChatCompletionResponseBase.md)
 - [ChatModerationRequest](docs/ChatModerationRequest.md)
 - [CheckpointOut](docs/CheckpointOut.md)
 - [ClassificationRequest](docs/ClassificationRequest.md)
 - [ClassificationResponse](docs/ClassificationResponse.md)
 - [ClassificationTargetResult](docs/ClassificationTargetResult.md)
 - [ClassifierDetailedJobOut](docs/ClassifierDetailedJobOut.md)
 - [ClassifierFTModelOut](docs/ClassifierFTModelOut.md)
 - [ClassifierJobOut](docs/ClassifierJobOut.md)
 - [ClassifierJobOutIntegrationsInner](docs/ClassifierJobOutIntegrationsInner.md)
 - [ClassifierTargetIn](docs/ClassifierTargetIn.md)
 - [ClassifierTargetOut](docs/ClassifierTargetOut.md)
 - [ClassifierTrainingParameters](docs/ClassifierTrainingParameters.md)
 - [ClassifierTrainingParametersIn](docs/ClassifierTrainingParametersIn.md)
 - [CompletionChunk](docs/CompletionChunk.md)
 - [CompletionDetailedJobOut](docs/CompletionDetailedJobOut.md)
 - [CompletionEvent](docs/CompletionEvent.md)
 - [CompletionFTModelOut](docs/CompletionFTModelOut.md)
 - [CompletionJobOut](docs/CompletionJobOut.md)
 - [CompletionJobOutRepositoriesInner](docs/CompletionJobOutRepositoriesInner.md)
 - [CompletionResponseStreamChoice](docs/CompletionResponseStreamChoice.md)
 - [CompletionTrainingParameters](docs/CompletionTrainingParameters.md)
 - [CompletionTrainingParametersIn](docs/CompletionTrainingParametersIn.md)
 - [Content](docs/Content.md)
 - [Content1](docs/Content1.md)
 - [ContentChunk](docs/ContentChunk.md)
 - [DeleteFileOut](docs/DeleteFileOut.md)
 - [DeleteModelOut](docs/DeleteModelOut.md)
 - [DeltaMessage](docs/DeltaMessage.md)
 - [DeltaMessageContent](docs/DeltaMessageContent.md)
 - [Description](docs/Description.md)
 - [Document](docs/Document.md)
 - [DocumentURLChunk](docs/DocumentURLChunk.md)
 - [EmbeddingRequest](docs/EmbeddingRequest.md)
 - [EmbeddingResponse](docs/EmbeddingResponse.md)
 - [EmbeddingResponseData](docs/EmbeddingResponseData.md)
 - [EventOut](docs/EventOut.md)
 - [FIMCompletionRequest](docs/FIMCompletionRequest.md)
 - [FIMCompletionResponse](docs/FIMCompletionResponse.md)
 - [FTClassifierLossFunction](docs/FTClassifierLossFunction.md)
 - [FTModelCapabilitiesOut](docs/FTModelCapabilitiesOut.md)
 - [FTModelCard](docs/FTModelCard.md)
 - [FilePurpose](docs/FilePurpose.md)
 - [FileSchema](docs/FileSchema.md)
 - [FileSignedURL](docs/FileSignedURL.md)
 - [FineTuneableModel](docs/FineTuneableModel.md)
 - [FineTuneableModelType](docs/FineTuneableModelType.md)
 - [Function](docs/Function.md)
 - [FunctionCall](docs/FunctionCall.md)
 - [FunctionName](docs/FunctionName.md)
 - [GithubRepositoryIn](docs/GithubRepositoryIn.md)
 - [GithubRepositoryOut](docs/GithubRepositoryOut.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [Hyperparameters](docs/Hyperparameters.md)
 - [ImageURL](docs/ImageURL.md)
 - [ImageURLChunk](docs/ImageURLChunk.md)
 - [ImageUrl](docs/ImageUrl.md)
 - [Input](docs/Input.md)
 - [Input1](docs/Input1.md)
 - [Input2](docs/Input2.md)
 - [InstructRequest](docs/InstructRequest.md)
 - [JobIn](docs/JobIn.md)
 - [JobInIntegrationsInner](docs/JobInIntegrationsInner.md)
 - [JobInRepositoriesInner](docs/JobInRepositoriesInner.md)
 - [JobMetadataOut](docs/JobMetadataOut.md)
 - [JobsOut](docs/JobsOut.md)
 - [JsonSchema](docs/JsonSchema.md)
 - [LegacyJobMetadataOut](docs/LegacyJobMetadataOut.md)
 - [ListFilesOut](docs/ListFilesOut.md)
 - [MaxTokens](docs/MaxTokens.md)
 - [MessagesInner](docs/MessagesInner.md)
 - [MetricOut](docs/MetricOut.md)
 - [ModelCapabilities](docs/ModelCapabilities.md)
 - [ModelList](docs/ModelList.md)
 - [ModelListDataInner](docs/ModelListDataInner.md)
 - [ModerationObject](docs/ModerationObject.md)
 - [ModerationResponse](docs/ModerationResponse.md)
 - [N](docs/N.md)
 - [OCRImageObject](docs/OCRImageObject.md)
 - [OCRPageDimensions](docs/OCRPageDimensions.md)
 - [OCRPageObject](docs/OCRPageObject.md)
 - [OCRRequest](docs/OCRRequest.md)
 - [OCRResponse](docs/OCRResponse.md)
 - [OCRUsageInfo](docs/OCRUsageInfo.md)
 - [Prediction](docs/Prediction.md)
 - [RandomSeed](docs/RandomSeed.md)
 - [ReferenceChunk](docs/ReferenceChunk.md)
 - [Response](docs/Response.md)
 - [Response1](docs/Response1.md)
 - [Response2](docs/Response2.md)
 - [ResponseAnyOf](docs/ResponseAnyOf.md)
 - [ResponseBase](docs/ResponseBase.md)
 - [ResponseFormat](docs/ResponseFormat.md)
 - [ResponseFormats](docs/ResponseFormats.md)
 - [ResponseRetrieveModelV1ModelsModelIdGet](docs/ResponseRetrieveModelV1ModelsModelIdGet.md)
 - [RetrieveFileOut](docs/RetrieveFileOut.md)
 - [SampleType](docs/SampleType.md)
 - [Source](docs/Source.md)
 - [Stop](docs/Stop.md)
 - [Stop1](docs/Stop1.md)
 - [SystemMessage](docs/SystemMessage.md)
 - [Temperature](docs/Temperature.md)
 - [TextChunk](docs/TextChunk.md)
 - [Tool](docs/Tool.md)
 - [ToolCall](docs/ToolCall.md)
 - [ToolChoice](docs/ToolChoice.md)
 - [ToolChoiceEnum](docs/ToolChoiceEnum.md)
 - [ToolMessage](docs/ToolMessage.md)
 - [ToolTypes](docs/ToolTypes.md)
 - [Tools](docs/Tools.md)
 - [TrainingFile](docs/TrainingFile.md)
 - [UnarchiveFTModelOut](docs/UnarchiveFTModelOut.md)
 - [UpdateFTModelIn](docs/UpdateFTModelIn.md)
 - [UploadFileOut](docs/UploadFileOut.md)
 - [UsageInfo](docs/UsageInfo.md)
 - [UserMessage](docs/UserMessage.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)
 - [WandbIntegration](docs/WandbIntegration.md)
 - [WandbIntegrationOut](docs/WandbIntegrationOut.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKey

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), mistral.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



