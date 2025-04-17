# Response1

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
**FineTunedModel** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**[]ClassifierJobOutIntegrationsInner**](ClassifierJobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**JobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] [default to "classifier"]
**Hyperparameters** | [**ClassifierTrainingParameters**](ClassifierTrainingParameters.md) |  | 
**Repositories** | Pointer to [**[]CompletionJobOutRepositoriesInner**](CompletionJobOutRepositoriesInner.md) |  | [optional] [default to []]
**Events** | Pointer to [**[]EventOut**](EventOut.md) | Event items are created every time the status of a fine-tuning job changes. The timestamped list of all events is accessible here. | [optional] [default to []]
**Checkpoints** | Pointer to [**[]CheckpointOut**](CheckpointOut.md) |  | [optional] [default to []]
**ClassifierTargets** | [**[]ClassifierTargetOut**](ClassifierTargetOut.md) |  | 

## Methods

### NewResponse1

`func NewResponse1(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters ClassifierTrainingParameters, classifierTargets []ClassifierTargetOut, ) *Response1`

NewResponse1 instantiates a new Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse1WithDefaults

`func NewResponse1WithDefaults() *Response1`

NewResponse1WithDefaults instantiates a new Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Response1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Response1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Response1) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *Response1) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *Response1) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *Response1) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetModel

`func (o *Response1) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Response1) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Response1) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *Response1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Response1) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Response1) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Response1) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Response1) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Response1) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Response1) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *Response1) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *Response1) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *Response1) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *Response1) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *Response1) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *Response1) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *Response1) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### GetObject

`func (o *Response1) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Response1) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Response1) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Response1) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *Response1) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *Response1) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *Response1) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *Response1) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### GetSuffix

`func (o *Response1) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *Response1) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *Response1) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *Response1) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetIntegrations

`func (o *Response1) GetIntegrations() []ClassifierJobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *Response1) GetIntegrationsOk() (*[]ClassifierJobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *Response1) SetIntegrations(v []ClassifierJobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *Response1) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetTrainedTokens

`func (o *Response1) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *Response1) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *Response1) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *Response1) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### GetMetadata

`func (o *Response1) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Response1) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Response1) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Response1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetJobType

`func (o *Response1) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *Response1) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *Response1) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *Response1) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperparameters

`func (o *Response1) GetHyperparameters() ClassifierTrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *Response1) GetHyperparametersOk() (*ClassifierTrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *Response1) SetHyperparameters(v ClassifierTrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetRepositories

`func (o *Response1) GetRepositories() []CompletionJobOutRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *Response1) GetRepositoriesOk() (*[]CompletionJobOutRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *Response1) SetRepositories(v []CompletionJobOutRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *Response1) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetEvents

`func (o *Response1) GetEvents() []EventOut`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Response1) GetEventsOk() (*[]EventOut, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Response1) SetEvents(v []EventOut)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Response1) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCheckpoints

`func (o *Response1) GetCheckpoints() []CheckpointOut`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *Response1) GetCheckpointsOk() (*[]CheckpointOut, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *Response1) SetCheckpoints(v []CheckpointOut)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *Response1) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.

### GetClassifierTargets

`func (o *Response1) GetClassifierTargets() []ClassifierTargetOut`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *Response1) GetClassifierTargetsOk() (*[]ClassifierTargetOut, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *Response1) SetClassifierTargets(v []ClassifierTargetOut)`

SetClassifierTargets sets ClassifierTargets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


