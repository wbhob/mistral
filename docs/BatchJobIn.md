# BatchJobIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFiles** | **[]string** |  | 
**Endpoint** | [**ApiEndpoint**](ApiEndpoint.md) |  | 
**Model** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**TimeoutHours** | Pointer to **int32** |  | [optional] [default to 24]

## Methods

### NewBatchJobIn

`func NewBatchJobIn(inputFiles []string, endpoint ApiEndpoint, model string, ) *BatchJobIn`

NewBatchJobIn instantiates a new BatchJobIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobInWithDefaults

`func NewBatchJobInWithDefaults() *BatchJobIn`

NewBatchJobInWithDefaults instantiates a new BatchJobIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFiles

`func (o *BatchJobIn) GetInputFiles() []string`

GetInputFiles returns the InputFiles field if non-nil, zero value otherwise.

### GetInputFilesOk

`func (o *BatchJobIn) GetInputFilesOk() (*[]string, bool)`

GetInputFilesOk returns a tuple with the InputFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFiles

`func (o *BatchJobIn) SetInputFiles(v []string)`

SetInputFiles sets InputFiles field to given value.


### GetEndpoint

`func (o *BatchJobIn) GetEndpoint() ApiEndpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BatchJobIn) GetEndpointOk() (*ApiEndpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BatchJobIn) SetEndpoint(v ApiEndpoint)`

SetEndpoint sets Endpoint field to given value.


### GetModel

`func (o *BatchJobIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BatchJobIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BatchJobIn) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetadata

`func (o *BatchJobIn) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BatchJobIn) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BatchJobIn) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BatchJobIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BatchJobIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BatchJobIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTimeoutHours

`func (o *BatchJobIn) GetTimeoutHours() int32`

GetTimeoutHours returns the TimeoutHours field if non-nil, zero value otherwise.

### GetTimeoutHoursOk

`func (o *BatchJobIn) GetTimeoutHoursOk() (*int32, bool)`

GetTimeoutHoursOk returns a tuple with the TimeoutHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutHours

`func (o *BatchJobIn) SetTimeoutHours(v int32)`

SetTimeoutHours sets TimeoutHours field to given value.

### HasTimeoutHours

`func (o *BatchJobIn) HasTimeoutHours() bool`

HasTimeoutHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


