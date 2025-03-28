# BatchJobOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "batch"]
**InputFiles** | **[]string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Endpoint** | **string** |  | 
**Model** | **string** |  | 
**OutputFile** | Pointer to **NullableString** |  | [optional] 
**ErrorFile** | Pointer to **NullableString** |  | [optional] 
**Errors** | [**[]BatchError**](BatchError.md) |  | 
**Status** | [**BatchJobStatus**](BatchJobStatus.md) |  | 
**CreatedAt** | **int32** |  | 
**TotalRequests** | **int32** |  | 
**CompletedRequests** | **int32** |  | 
**SucceededRequests** | **int32** |  | 
**FailedRequests** | **int32** |  | 
**StartedAt** | Pointer to **NullableInt32** |  | [optional] 
**CompletedAt** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBatchJobOut

`func NewBatchJobOut(id string, inputFiles []string, endpoint string, model string, errors []BatchError, status BatchJobStatus, createdAt int32, totalRequests int32, completedRequests int32, succeededRequests int32, failedRequests int32, ) *BatchJobOut`

NewBatchJobOut instantiates a new BatchJobOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobOutWithDefaults

`func NewBatchJobOutWithDefaults() *BatchJobOut`

NewBatchJobOutWithDefaults instantiates a new BatchJobOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchJobOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchJobOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchJobOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *BatchJobOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BatchJobOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BatchJobOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BatchJobOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetInputFiles

`func (o *BatchJobOut) GetInputFiles() []string`

GetInputFiles returns the InputFiles field if non-nil, zero value otherwise.

### GetInputFilesOk

`func (o *BatchJobOut) GetInputFilesOk() (*[]string, bool)`

GetInputFilesOk returns a tuple with the InputFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFiles

`func (o *BatchJobOut) SetInputFiles(v []string)`

SetInputFiles sets InputFiles field to given value.


### GetMetadata

`func (o *BatchJobOut) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BatchJobOut) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BatchJobOut) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BatchJobOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BatchJobOut) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BatchJobOut) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetEndpoint

`func (o *BatchJobOut) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BatchJobOut) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BatchJobOut) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetModel

`func (o *BatchJobOut) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BatchJobOut) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BatchJobOut) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputFile

`func (o *BatchJobOut) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *BatchJobOut) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *BatchJobOut) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *BatchJobOut) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### SetOutputFileNil

`func (o *BatchJobOut) SetOutputFileNil(b bool)`

 SetOutputFileNil sets the value for OutputFile to be an explicit nil

### UnsetOutputFile
`func (o *BatchJobOut) UnsetOutputFile()`

UnsetOutputFile ensures that no value is present for OutputFile, not even an explicit nil
### GetErrorFile

`func (o *BatchJobOut) GetErrorFile() string`

GetErrorFile returns the ErrorFile field if non-nil, zero value otherwise.

### GetErrorFileOk

`func (o *BatchJobOut) GetErrorFileOk() (*string, bool)`

GetErrorFileOk returns a tuple with the ErrorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFile

`func (o *BatchJobOut) SetErrorFile(v string)`

SetErrorFile sets ErrorFile field to given value.

### HasErrorFile

`func (o *BatchJobOut) HasErrorFile() bool`

HasErrorFile returns a boolean if a field has been set.

### SetErrorFileNil

`func (o *BatchJobOut) SetErrorFileNil(b bool)`

 SetErrorFileNil sets the value for ErrorFile to be an explicit nil

### UnsetErrorFile
`func (o *BatchJobOut) UnsetErrorFile()`

UnsetErrorFile ensures that no value is present for ErrorFile, not even an explicit nil
### GetErrors

`func (o *BatchJobOut) GetErrors() []BatchError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchJobOut) GetErrorsOk() (*[]BatchError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchJobOut) SetErrors(v []BatchError)`

SetErrors sets Errors field to given value.


### GetStatus

`func (o *BatchJobOut) GetStatus() BatchJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchJobOut) GetStatusOk() (*BatchJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchJobOut) SetStatus(v BatchJobStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BatchJobOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchJobOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchJobOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetTotalRequests

`func (o *BatchJobOut) GetTotalRequests() int32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *BatchJobOut) GetTotalRequestsOk() (*int32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *BatchJobOut) SetTotalRequests(v int32)`

SetTotalRequests sets TotalRequests field to given value.


### GetCompletedRequests

`func (o *BatchJobOut) GetCompletedRequests() int32`

GetCompletedRequests returns the CompletedRequests field if non-nil, zero value otherwise.

### GetCompletedRequestsOk

`func (o *BatchJobOut) GetCompletedRequestsOk() (*int32, bool)`

GetCompletedRequestsOk returns a tuple with the CompletedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedRequests

`func (o *BatchJobOut) SetCompletedRequests(v int32)`

SetCompletedRequests sets CompletedRequests field to given value.


### GetSucceededRequests

`func (o *BatchJobOut) GetSucceededRequests() int32`

GetSucceededRequests returns the SucceededRequests field if non-nil, zero value otherwise.

### GetSucceededRequestsOk

`func (o *BatchJobOut) GetSucceededRequestsOk() (*int32, bool)`

GetSucceededRequestsOk returns a tuple with the SucceededRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededRequests

`func (o *BatchJobOut) SetSucceededRequests(v int32)`

SetSucceededRequests sets SucceededRequests field to given value.


### GetFailedRequests

`func (o *BatchJobOut) GetFailedRequests() int32`

GetFailedRequests returns the FailedRequests field if non-nil, zero value otherwise.

### GetFailedRequestsOk

`func (o *BatchJobOut) GetFailedRequestsOk() (*int32, bool)`

GetFailedRequestsOk returns a tuple with the FailedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRequests

`func (o *BatchJobOut) SetFailedRequests(v int32)`

SetFailedRequests sets FailedRequests field to given value.


### GetStartedAt

`func (o *BatchJobOut) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchJobOut) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchJobOut) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BatchJobOut) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BatchJobOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BatchJobOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *BatchJobOut) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchJobOut) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchJobOut) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BatchJobOut) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BatchJobOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BatchJobOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


