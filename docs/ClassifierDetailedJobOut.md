# ClassifierDetailedJobOut

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
**JobType** | Pointer to **string** |  | [optional] [default to "classifier"]
**Hyperparameters** | [**ClassifierTrainingParameters**](ClassifierTrainingParameters.md) |  | 
**Events** | Pointer to [**[]EventOut**](EventOut.md) | Event items are created every time the status of a fine-tuning job changes. The timestamped list of all events is accessible here. | [optional] [default to []]
**Checkpoints** | Pointer to [**[]CheckpointOut**](CheckpointOut.md) |  | [optional] [default to []]
**ClassifierTargets** | [**[]ClassifierTargetOut**](ClassifierTargetOut.md) |  | 

## Methods

### NewClassifierDetailedJobOut

`func NewClassifierDetailedJobOut(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters ClassifierTrainingParameters, classifierTargets []ClassifierTargetOut, ) *ClassifierDetailedJobOut`

NewClassifierDetailedJobOut instantiates a new ClassifierDetailedJobOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierDetailedJobOutWithDefaults

`func NewClassifierDetailedJobOutWithDefaults() *ClassifierDetailedJobOut`

NewClassifierDetailedJobOutWithDefaults instantiates a new ClassifierDetailedJobOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassifierDetailedJobOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassifierDetailedJobOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassifierDetailedJobOut) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *ClassifierDetailedJobOut) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *ClassifierDetailedJobOut) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *ClassifierDetailedJobOut) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetModel

`func (o *ClassifierDetailedJobOut) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClassifierDetailedJobOut) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClassifierDetailedJobOut) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *ClassifierDetailedJobOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClassifierDetailedJobOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClassifierDetailedJobOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ClassifierDetailedJobOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClassifierDetailedJobOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClassifierDetailedJobOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ClassifierDetailedJobOut) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClassifierDetailedJobOut) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClassifierDetailedJobOut) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *ClassifierDetailedJobOut) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *ClassifierDetailedJobOut) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *ClassifierDetailedJobOut) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *ClassifierDetailedJobOut) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *ClassifierDetailedJobOut) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *ClassifierDetailedJobOut) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *ClassifierDetailedJobOut) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### SetValidationFilesNil

`func (o *ClassifierDetailedJobOut) SetValidationFilesNil(b bool)`

 SetValidationFilesNil sets the value for ValidationFiles to be an explicit nil

### UnsetValidationFiles
`func (o *ClassifierDetailedJobOut) UnsetValidationFiles()`

UnsetValidationFiles ensures that no value is present for ValidationFiles, not even an explicit nil
### GetObject

`func (o *ClassifierDetailedJobOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ClassifierDetailedJobOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ClassifierDetailedJobOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ClassifierDetailedJobOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *ClassifierDetailedJobOut) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *ClassifierDetailedJobOut) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *ClassifierDetailedJobOut) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *ClassifierDetailedJobOut) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### SetFineTunedModelNil

`func (o *ClassifierDetailedJobOut) SetFineTunedModelNil(b bool)`

 SetFineTunedModelNil sets the value for FineTunedModel to be an explicit nil

### UnsetFineTunedModel
`func (o *ClassifierDetailedJobOut) UnsetFineTunedModel()`

UnsetFineTunedModel ensures that no value is present for FineTunedModel, not even an explicit nil
### GetSuffix

`func (o *ClassifierDetailedJobOut) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *ClassifierDetailedJobOut) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *ClassifierDetailedJobOut) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *ClassifierDetailedJobOut) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *ClassifierDetailedJobOut) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *ClassifierDetailedJobOut) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIntegrations

`func (o *ClassifierDetailedJobOut) GetIntegrations() []ClassifierJobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ClassifierDetailedJobOut) GetIntegrationsOk() (*[]ClassifierJobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ClassifierDetailedJobOut) SetIntegrations(v []ClassifierJobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ClassifierDetailedJobOut) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *ClassifierDetailedJobOut) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *ClassifierDetailedJobOut) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetTrainedTokens

`func (o *ClassifierDetailedJobOut) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *ClassifierDetailedJobOut) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *ClassifierDetailedJobOut) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *ClassifierDetailedJobOut) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### SetTrainedTokensNil

`func (o *ClassifierDetailedJobOut) SetTrainedTokensNil(b bool)`

 SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil

### UnsetTrainedTokens
`func (o *ClassifierDetailedJobOut) UnsetTrainedTokens()`

UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
### GetMetadata

`func (o *ClassifierDetailedJobOut) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClassifierDetailedJobOut) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClassifierDetailedJobOut) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClassifierDetailedJobOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ClassifierDetailedJobOut) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ClassifierDetailedJobOut) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetJobType

`func (o *ClassifierDetailedJobOut) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ClassifierDetailedJobOut) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ClassifierDetailedJobOut) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ClassifierDetailedJobOut) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperparameters

`func (o *ClassifierDetailedJobOut) GetHyperparameters() ClassifierTrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *ClassifierDetailedJobOut) GetHyperparametersOk() (*ClassifierTrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *ClassifierDetailedJobOut) SetHyperparameters(v ClassifierTrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetEvents

`func (o *ClassifierDetailedJobOut) GetEvents() []EventOut`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClassifierDetailedJobOut) GetEventsOk() (*[]EventOut, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClassifierDetailedJobOut) SetEvents(v []EventOut)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClassifierDetailedJobOut) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCheckpoints

`func (o *ClassifierDetailedJobOut) GetCheckpoints() []CheckpointOut`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *ClassifierDetailedJobOut) GetCheckpointsOk() (*[]CheckpointOut, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *ClassifierDetailedJobOut) SetCheckpoints(v []CheckpointOut)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *ClassifierDetailedJobOut) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.

### GetClassifierTargets

`func (o *ClassifierDetailedJobOut) GetClassifierTargets() []ClassifierTargetOut`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *ClassifierDetailedJobOut) GetClassifierTargetsOk() (*[]ClassifierTargetOut, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *ClassifierDetailedJobOut) SetClassifierTargets(v []ClassifierTargetOut)`

SetClassifierTargets sets ClassifierTargets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


