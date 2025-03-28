# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the job. | 
**AutoStart** | **bool** |  | 
**Hyperparameters** | [**TrainingParameters**](TrainingParameters.md) |  | 
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**Status** | **string** | The current status of the fine-tuning job. | 
**JobType** | **string** | The type of job (&#x60;FT&#x60; for fine-tuning). | 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) for when the fine-tuning job was created. | 
**ModifiedAt** | **int32** | The UNIX timestamp (in seconds) for when the fine-tuning job was last modified. | 
**TrainingFiles** | **[]string** | A list containing the IDs of uploaded files that contain training data. | 
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "job.metadata"]
**FineTunedModel** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**[]JobOutIntegrationsInner**](JobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **int32** |  | [optional] 
**Repositories** | Pointer to [**[]JobOutRepositoriesInner**](JobOutRepositoriesInner.md) |  | [optional] [default to []]
**Metadata** | Pointer to [**JobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**ExpectedDurationSeconds** | Pointer to **int32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**CostCurrency** | Pointer to **string** |  | [optional] 
**TrainTokensPerStep** | Pointer to **int32** |  | [optional] 
**TrainTokens** | Pointer to **int32** |  | [optional] 
**DataTokens** | Pointer to **int32** |  | [optional] 
**EstimatedStartTime** | Pointer to **int32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] [default to true]
**Details** | **string** |  | 
**Epochs** | Pointer to **float32** |  | [optional] 
**TrainingSteps** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponse

`func NewResponse(id string, autoStart bool, hyperparameters TrainingParameters, model FineTuneableModel, status string, jobType string, createdAt int32, modifiedAt int32, trainingFiles []string, details string, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Response) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *Response) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *Response) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *Response) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetHyperparameters

`func (o *Response) GetHyperparameters() TrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *Response) GetHyperparametersOk() (*TrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *Response) SetHyperparameters(v TrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetModel

`func (o *Response) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Response) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Response) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJobType

`func (o *Response) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *Response) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *Response) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetCreatedAt

`func (o *Response) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Response) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Response) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Response) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Response) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Response) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *Response) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *Response) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *Response) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *Response) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *Response) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *Response) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *Response) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### GetObject

`func (o *Response) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Response) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Response) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *Response) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *Response) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *Response) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *Response) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### GetSuffix

`func (o *Response) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *Response) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *Response) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *Response) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetIntegrations

`func (o *Response) GetIntegrations() []JobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *Response) GetIntegrationsOk() (*[]JobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *Response) SetIntegrations(v []JobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *Response) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetTrainedTokens

`func (o *Response) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *Response) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *Response) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *Response) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### GetRepositories

`func (o *Response) GetRepositories() []JobOutRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *Response) GetRepositoriesOk() (*[]JobOutRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *Response) SetRepositories(v []JobOutRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *Response) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetMetadata

`func (o *Response) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Response) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Response) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpectedDurationSeconds

`func (o *Response) GetExpectedDurationSeconds() int32`

GetExpectedDurationSeconds returns the ExpectedDurationSeconds field if non-nil, zero value otherwise.

### GetExpectedDurationSecondsOk

`func (o *Response) GetExpectedDurationSecondsOk() (*int32, bool)`

GetExpectedDurationSecondsOk returns a tuple with the ExpectedDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDurationSeconds

`func (o *Response) SetExpectedDurationSeconds(v int32)`

SetExpectedDurationSeconds sets ExpectedDurationSeconds field to given value.

### HasExpectedDurationSeconds

`func (o *Response) HasExpectedDurationSeconds() bool`

HasExpectedDurationSeconds returns a boolean if a field has been set.

### GetCost

`func (o *Response) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Response) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Response) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Response) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostCurrency

`func (o *Response) GetCostCurrency() string`

GetCostCurrency returns the CostCurrency field if non-nil, zero value otherwise.

### GetCostCurrencyOk

`func (o *Response) GetCostCurrencyOk() (*string, bool)`

GetCostCurrencyOk returns a tuple with the CostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCurrency

`func (o *Response) SetCostCurrency(v string)`

SetCostCurrency sets CostCurrency field to given value.

### HasCostCurrency

`func (o *Response) HasCostCurrency() bool`

HasCostCurrency returns a boolean if a field has been set.

### GetTrainTokensPerStep

`func (o *Response) GetTrainTokensPerStep() int32`

GetTrainTokensPerStep returns the TrainTokensPerStep field if non-nil, zero value otherwise.

### GetTrainTokensPerStepOk

`func (o *Response) GetTrainTokensPerStepOk() (*int32, bool)`

GetTrainTokensPerStepOk returns a tuple with the TrainTokensPerStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokensPerStep

`func (o *Response) SetTrainTokensPerStep(v int32)`

SetTrainTokensPerStep sets TrainTokensPerStep field to given value.

### HasTrainTokensPerStep

`func (o *Response) HasTrainTokensPerStep() bool`

HasTrainTokensPerStep returns a boolean if a field has been set.

### GetTrainTokens

`func (o *Response) GetTrainTokens() int32`

GetTrainTokens returns the TrainTokens field if non-nil, zero value otherwise.

### GetTrainTokensOk

`func (o *Response) GetTrainTokensOk() (*int32, bool)`

GetTrainTokensOk returns a tuple with the TrainTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokens

`func (o *Response) SetTrainTokens(v int32)`

SetTrainTokens sets TrainTokens field to given value.

### HasTrainTokens

`func (o *Response) HasTrainTokens() bool`

HasTrainTokens returns a boolean if a field has been set.

### GetDataTokens

`func (o *Response) GetDataTokens() int32`

GetDataTokens returns the DataTokens field if non-nil, zero value otherwise.

### GetDataTokensOk

`func (o *Response) GetDataTokensOk() (*int32, bool)`

GetDataTokensOk returns a tuple with the DataTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTokens

`func (o *Response) SetDataTokens(v int32)`

SetDataTokens sets DataTokens field to given value.

### HasDataTokens

`func (o *Response) HasDataTokens() bool`

HasDataTokens returns a boolean if a field has been set.

### GetEstimatedStartTime

`func (o *Response) GetEstimatedStartTime() int32`

GetEstimatedStartTime returns the EstimatedStartTime field if non-nil, zero value otherwise.

### GetEstimatedStartTimeOk

`func (o *Response) GetEstimatedStartTimeOk() (*int32, bool)`

GetEstimatedStartTimeOk returns a tuple with the EstimatedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartTime

`func (o *Response) SetEstimatedStartTime(v int32)`

SetEstimatedStartTime sets EstimatedStartTime field to given value.

### HasEstimatedStartTime

`func (o *Response) HasEstimatedStartTime() bool`

HasEstimatedStartTime returns a boolean if a field has been set.

### GetDeprecated

`func (o *Response) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Response) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Response) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Response) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDetails

`func (o *Response) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Response) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Response) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetEpochs

`func (o *Response) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *Response) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *Response) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *Response) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### GetTrainingSteps

`func (o *Response) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *Response) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *Response) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *Response) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


