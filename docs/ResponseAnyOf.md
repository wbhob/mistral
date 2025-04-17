# ResponseAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the job. | 
**AutoStart** | **bool** |  | 
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**Status** | **string** | The current status of the fine-tuning job. | 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) for when the fine-tuning job was created. | 
**ModifiedAt** | **int32** | The UNIX timestamp (in seconds) for when the fine-tuning job was last modified. | 
**TrainingFiles** | **[]string** | A list containing the IDs of uploaded files that contain training data. | 
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Object** | Pointer to **string** | The object type of the fine-tuning job. | [optional] [default to "job"]
**FineTunedModel** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**[]ClassifierJobOutIntegrationsInner**](ClassifierJobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**JobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**JobType** | Pointer to **string** | The type of job (&#x60;FT&#x60; for fine-tuning). | [optional] [default to "classifier"]
**Hyperparameters** | [**ClassifierTrainingParameters**](ClassifierTrainingParameters.md) |  | 
**Repositories** | Pointer to [**[]CompletionJobOutRepositoriesInner**](CompletionJobOutRepositoriesInner.md) |  | [optional] [default to []]

## Methods

### NewResponseAnyOf

`func NewResponseAnyOf(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters ClassifierTrainingParameters, ) *ResponseAnyOf`

NewResponseAnyOf instantiates a new ResponseAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAnyOfWithDefaults

`func NewResponseAnyOfWithDefaults() *ResponseAnyOf`

NewResponseAnyOfWithDefaults instantiates a new ResponseAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseAnyOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseAnyOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseAnyOf) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *ResponseAnyOf) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *ResponseAnyOf) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *ResponseAnyOf) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetModel

`func (o *ResponseAnyOf) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponseAnyOf) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponseAnyOf) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *ResponseAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ResponseAnyOf) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseAnyOf) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseAnyOf) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ResponseAnyOf) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ResponseAnyOf) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ResponseAnyOf) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *ResponseAnyOf) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *ResponseAnyOf) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *ResponseAnyOf) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *ResponseAnyOf) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *ResponseAnyOf) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *ResponseAnyOf) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *ResponseAnyOf) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### GetObject

`func (o *ResponseAnyOf) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseAnyOf) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseAnyOf) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseAnyOf) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *ResponseAnyOf) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *ResponseAnyOf) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *ResponseAnyOf) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *ResponseAnyOf) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### GetSuffix

`func (o *ResponseAnyOf) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *ResponseAnyOf) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *ResponseAnyOf) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *ResponseAnyOf) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetIntegrations

`func (o *ResponseAnyOf) GetIntegrations() []ClassifierJobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ResponseAnyOf) GetIntegrationsOk() (*[]ClassifierJobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ResponseAnyOf) SetIntegrations(v []ClassifierJobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ResponseAnyOf) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetTrainedTokens

`func (o *ResponseAnyOf) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *ResponseAnyOf) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *ResponseAnyOf) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *ResponseAnyOf) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponseAnyOf) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponseAnyOf) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponseAnyOf) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponseAnyOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetJobType

`func (o *ResponseAnyOf) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ResponseAnyOf) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ResponseAnyOf) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ResponseAnyOf) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperparameters

`func (o *ResponseAnyOf) GetHyperparameters() ClassifierTrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *ResponseAnyOf) GetHyperparametersOk() (*ClassifierTrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *ResponseAnyOf) SetHyperparameters(v ClassifierTrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetRepositories

`func (o *ResponseAnyOf) GetRepositories() []CompletionJobOutRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ResponseAnyOf) GetRepositoriesOk() (*[]CompletionJobOutRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ResponseAnyOf) SetRepositories(v []CompletionJobOutRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ResponseAnyOf) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


