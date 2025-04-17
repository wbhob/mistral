# CompletionDetailedJobOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AutoStart** | **bool** |  | 
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**Status** | **string** |  | 
**CreatedAt** | **int32** |  | 
**ModifiedAt** | **int32** |  | 
**TrainingFiles** | **[]string** |  | 
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "job"]
**FineTunedModel** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Integrations** | Pointer to [**[]ClassifierJobOutIntegrationsInner**](ClassifierJobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to [**NullableJobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] [default to "completion"]
**Hyperparameters** | [**CompletionTrainingParameters**](CompletionTrainingParameters.md) |  | 
**Repositories** | Pointer to [**[]CompletionJobOutRepositoriesInner**](CompletionJobOutRepositoriesInner.md) |  | [optional] [default to []]
**Events** | Pointer to [**[]EventOut**](EventOut.md) | Event items are created every time the status of a fine-tuning job changes. The timestamped list of all events is accessible here. | [optional] [default to []]
**Checkpoints** | Pointer to [**[]CheckpointOut**](CheckpointOut.md) |  | [optional] [default to []]

## Methods

### NewCompletionDetailedJobOut

`func NewCompletionDetailedJobOut(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters CompletionTrainingParameters, ) *CompletionDetailedJobOut`

NewCompletionDetailedJobOut instantiates a new CompletionDetailedJobOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionDetailedJobOutWithDefaults

`func NewCompletionDetailedJobOutWithDefaults() *CompletionDetailedJobOut`

NewCompletionDetailedJobOutWithDefaults instantiates a new CompletionDetailedJobOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompletionDetailedJobOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompletionDetailedJobOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompletionDetailedJobOut) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *CompletionDetailedJobOut) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *CompletionDetailedJobOut) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *CompletionDetailedJobOut) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetModel

`func (o *CompletionDetailedJobOut) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CompletionDetailedJobOut) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CompletionDetailedJobOut) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *CompletionDetailedJobOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompletionDetailedJobOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompletionDetailedJobOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CompletionDetailedJobOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompletionDetailedJobOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompletionDetailedJobOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *CompletionDetailedJobOut) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CompletionDetailedJobOut) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CompletionDetailedJobOut) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *CompletionDetailedJobOut) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *CompletionDetailedJobOut) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *CompletionDetailedJobOut) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *CompletionDetailedJobOut) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *CompletionDetailedJobOut) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *CompletionDetailedJobOut) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *CompletionDetailedJobOut) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### SetValidationFilesNil

`func (o *CompletionDetailedJobOut) SetValidationFilesNil(b bool)`

 SetValidationFilesNil sets the value for ValidationFiles to be an explicit nil

### UnsetValidationFiles
`func (o *CompletionDetailedJobOut) UnsetValidationFiles()`

UnsetValidationFiles ensures that no value is present for ValidationFiles, not even an explicit nil
### GetObject

`func (o *CompletionDetailedJobOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompletionDetailedJobOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompletionDetailedJobOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompletionDetailedJobOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *CompletionDetailedJobOut) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *CompletionDetailedJobOut) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *CompletionDetailedJobOut) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *CompletionDetailedJobOut) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### SetFineTunedModelNil

`func (o *CompletionDetailedJobOut) SetFineTunedModelNil(b bool)`

 SetFineTunedModelNil sets the value for FineTunedModel to be an explicit nil

### UnsetFineTunedModel
`func (o *CompletionDetailedJobOut) UnsetFineTunedModel()`

UnsetFineTunedModel ensures that no value is present for FineTunedModel, not even an explicit nil
### GetSuffix

`func (o *CompletionDetailedJobOut) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CompletionDetailedJobOut) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CompletionDetailedJobOut) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CompletionDetailedJobOut) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *CompletionDetailedJobOut) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *CompletionDetailedJobOut) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIntegrations

`func (o *CompletionDetailedJobOut) GetIntegrations() []ClassifierJobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *CompletionDetailedJobOut) GetIntegrationsOk() (*[]ClassifierJobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *CompletionDetailedJobOut) SetIntegrations(v []ClassifierJobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *CompletionDetailedJobOut) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *CompletionDetailedJobOut) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *CompletionDetailedJobOut) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetTrainedTokens

`func (o *CompletionDetailedJobOut) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *CompletionDetailedJobOut) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *CompletionDetailedJobOut) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *CompletionDetailedJobOut) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### SetTrainedTokensNil

`func (o *CompletionDetailedJobOut) SetTrainedTokensNil(b bool)`

 SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil

### UnsetTrainedTokens
`func (o *CompletionDetailedJobOut) UnsetTrainedTokens()`

UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
### GetMetadata

`func (o *CompletionDetailedJobOut) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CompletionDetailedJobOut) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CompletionDetailedJobOut) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CompletionDetailedJobOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CompletionDetailedJobOut) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CompletionDetailedJobOut) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetJobType

`func (o *CompletionDetailedJobOut) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CompletionDetailedJobOut) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CompletionDetailedJobOut) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CompletionDetailedJobOut) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperparameters

`func (o *CompletionDetailedJobOut) GetHyperparameters() CompletionTrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *CompletionDetailedJobOut) GetHyperparametersOk() (*CompletionTrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *CompletionDetailedJobOut) SetHyperparameters(v CompletionTrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetRepositories

`func (o *CompletionDetailedJobOut) GetRepositories() []CompletionJobOutRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *CompletionDetailedJobOut) GetRepositoriesOk() (*[]CompletionJobOutRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *CompletionDetailedJobOut) SetRepositories(v []CompletionJobOutRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *CompletionDetailedJobOut) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetEvents

`func (o *CompletionDetailedJobOut) GetEvents() []EventOut`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CompletionDetailedJobOut) GetEventsOk() (*[]EventOut, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CompletionDetailedJobOut) SetEvents(v []EventOut)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CompletionDetailedJobOut) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCheckpoints

`func (o *CompletionDetailedJobOut) GetCheckpoints() []CheckpointOut`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *CompletionDetailedJobOut) GetCheckpointsOk() (*[]CheckpointOut, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *CompletionDetailedJobOut) SetCheckpoints(v []CheckpointOut)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *CompletionDetailedJobOut) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


