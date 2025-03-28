# DetailedJobOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AutoStart** | **bool** |  | 
**Hyperparameters** | [**TrainingParameters**](TrainingParameters.md) |  | 
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**Status** | **string** |  | 
**JobType** | **string** |  | 
**CreatedAt** | **int32** |  | 
**ModifiedAt** | **int32** |  | 
**TrainingFiles** | **[]string** |  | 
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "job"]
**FineTunedModel** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Integrations** | Pointer to [**[]JobOutIntegrationsInner**](JobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **NullableInt32** |  | [optional] 
**Repositories** | Pointer to [**[]JobOutRepositoriesInner**](JobOutRepositoriesInner.md) |  | [optional] [default to []]
**Metadata** | Pointer to [**NullableJobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**Events** | Pointer to [**[]EventOut**](EventOut.md) | Event items are created every time the status of a fine-tuning job changes. The timestamped list of all events is accessible here. | [optional] [default to []]
**Checkpoints** | Pointer to [**[]CheckpointOut**](CheckpointOut.md) |  | [optional] [default to []]

## Methods

### NewDetailedJobOut

`func NewDetailedJobOut(id string, autoStart bool, hyperparameters TrainingParameters, model FineTuneableModel, status string, jobType string, createdAt int32, modifiedAt int32, trainingFiles []string, ) *DetailedJobOut`

NewDetailedJobOut instantiates a new DetailedJobOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedJobOutWithDefaults

`func NewDetailedJobOutWithDefaults() *DetailedJobOut`

NewDetailedJobOutWithDefaults instantiates a new DetailedJobOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedJobOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedJobOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedJobOut) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *DetailedJobOut) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *DetailedJobOut) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *DetailedJobOut) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetHyperparameters

`func (o *DetailedJobOut) GetHyperparameters() TrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *DetailedJobOut) GetHyperparametersOk() (*TrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *DetailedJobOut) SetHyperparameters(v TrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetModel

`func (o *DetailedJobOut) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DetailedJobOut) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DetailedJobOut) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *DetailedJobOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedJobOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedJobOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJobType

`func (o *DetailedJobOut) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *DetailedJobOut) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *DetailedJobOut) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetCreatedAt

`func (o *DetailedJobOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailedJobOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailedJobOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *DetailedJobOut) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DetailedJobOut) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DetailedJobOut) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *DetailedJobOut) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *DetailedJobOut) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *DetailedJobOut) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *DetailedJobOut) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *DetailedJobOut) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *DetailedJobOut) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *DetailedJobOut) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### SetValidationFilesNil

`func (o *DetailedJobOut) SetValidationFilesNil(b bool)`

 SetValidationFilesNil sets the value for ValidationFiles to be an explicit nil

### UnsetValidationFiles
`func (o *DetailedJobOut) UnsetValidationFiles()`

UnsetValidationFiles ensures that no value is present for ValidationFiles, not even an explicit nil
### GetObject

`func (o *DetailedJobOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DetailedJobOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DetailedJobOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DetailedJobOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *DetailedJobOut) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *DetailedJobOut) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *DetailedJobOut) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *DetailedJobOut) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### SetFineTunedModelNil

`func (o *DetailedJobOut) SetFineTunedModelNil(b bool)`

 SetFineTunedModelNil sets the value for FineTunedModel to be an explicit nil

### UnsetFineTunedModel
`func (o *DetailedJobOut) UnsetFineTunedModel()`

UnsetFineTunedModel ensures that no value is present for FineTunedModel, not even an explicit nil
### GetSuffix

`func (o *DetailedJobOut) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *DetailedJobOut) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *DetailedJobOut) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *DetailedJobOut) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *DetailedJobOut) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *DetailedJobOut) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIntegrations

`func (o *DetailedJobOut) GetIntegrations() []JobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *DetailedJobOut) GetIntegrationsOk() (*[]JobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *DetailedJobOut) SetIntegrations(v []JobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *DetailedJobOut) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *DetailedJobOut) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *DetailedJobOut) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetTrainedTokens

`func (o *DetailedJobOut) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *DetailedJobOut) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *DetailedJobOut) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *DetailedJobOut) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### SetTrainedTokensNil

`func (o *DetailedJobOut) SetTrainedTokensNil(b bool)`

 SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil

### UnsetTrainedTokens
`func (o *DetailedJobOut) UnsetTrainedTokens()`

UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
### GetRepositories

`func (o *DetailedJobOut) GetRepositories() []JobOutRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *DetailedJobOut) GetRepositoriesOk() (*[]JobOutRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *DetailedJobOut) SetRepositories(v []JobOutRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *DetailedJobOut) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetMetadata

`func (o *DetailedJobOut) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DetailedJobOut) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DetailedJobOut) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DetailedJobOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DetailedJobOut) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DetailedJobOut) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetEvents

`func (o *DetailedJobOut) GetEvents() []EventOut`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DetailedJobOut) GetEventsOk() (*[]EventOut, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DetailedJobOut) SetEvents(v []EventOut)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *DetailedJobOut) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCheckpoints

`func (o *DetailedJobOut) GetCheckpoints() []CheckpointOut`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *DetailedJobOut) GetCheckpointsOk() (*[]CheckpointOut, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *DetailedJobOut) SetCheckpoints(v []CheckpointOut)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *DetailedJobOut) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


