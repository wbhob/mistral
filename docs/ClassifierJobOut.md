# ClassifierJobOut

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
**FineTunedModel** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Integrations** | Pointer to [**[]ClassifierJobOutIntegrationsInner**](ClassifierJobOutIntegrationsInner.md) |  | [optional] 
**TrainedTokens** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to [**NullableJobMetadataOut**](JobMetadataOut.md) |  | [optional] 
**JobType** | Pointer to **string** | The type of job (&#x60;FT&#x60; for fine-tuning). | [optional] [default to "classifier"]
**Hyperparameters** | [**ClassifierTrainingParameters**](ClassifierTrainingParameters.md) |  | 

## Methods

### NewClassifierJobOut

`func NewClassifierJobOut(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters ClassifierTrainingParameters, ) *ClassifierJobOut`

NewClassifierJobOut instantiates a new ClassifierJobOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierJobOutWithDefaults

`func NewClassifierJobOutWithDefaults() *ClassifierJobOut`

NewClassifierJobOutWithDefaults instantiates a new ClassifierJobOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassifierJobOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassifierJobOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassifierJobOut) SetId(v string)`

SetId sets Id field to given value.


### GetAutoStart

`func (o *ClassifierJobOut) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *ClassifierJobOut) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *ClassifierJobOut) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.


### GetModel

`func (o *ClassifierJobOut) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClassifierJobOut) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClassifierJobOut) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetStatus

`func (o *ClassifierJobOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClassifierJobOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClassifierJobOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ClassifierJobOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClassifierJobOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClassifierJobOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ClassifierJobOut) GetModifiedAt() int32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClassifierJobOut) GetModifiedAtOk() (*int32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClassifierJobOut) SetModifiedAt(v int32)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTrainingFiles

`func (o *ClassifierJobOut) GetTrainingFiles() []string`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *ClassifierJobOut) GetTrainingFilesOk() (*[]string, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *ClassifierJobOut) SetTrainingFiles(v []string)`

SetTrainingFiles sets TrainingFiles field to given value.


### GetValidationFiles

`func (o *ClassifierJobOut) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *ClassifierJobOut) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *ClassifierJobOut) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *ClassifierJobOut) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### SetValidationFilesNil

`func (o *ClassifierJobOut) SetValidationFilesNil(b bool)`

 SetValidationFilesNil sets the value for ValidationFiles to be an explicit nil

### UnsetValidationFiles
`func (o *ClassifierJobOut) UnsetValidationFiles()`

UnsetValidationFiles ensures that no value is present for ValidationFiles, not even an explicit nil
### GetObject

`func (o *ClassifierJobOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ClassifierJobOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ClassifierJobOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ClassifierJobOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFineTunedModel

`func (o *ClassifierJobOut) GetFineTunedModel() string`

GetFineTunedModel returns the FineTunedModel field if non-nil, zero value otherwise.

### GetFineTunedModelOk

`func (o *ClassifierJobOut) GetFineTunedModelOk() (*string, bool)`

GetFineTunedModelOk returns a tuple with the FineTunedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTunedModel

`func (o *ClassifierJobOut) SetFineTunedModel(v string)`

SetFineTunedModel sets FineTunedModel field to given value.

### HasFineTunedModel

`func (o *ClassifierJobOut) HasFineTunedModel() bool`

HasFineTunedModel returns a boolean if a field has been set.

### SetFineTunedModelNil

`func (o *ClassifierJobOut) SetFineTunedModelNil(b bool)`

 SetFineTunedModelNil sets the value for FineTunedModel to be an explicit nil

### UnsetFineTunedModel
`func (o *ClassifierJobOut) UnsetFineTunedModel()`

UnsetFineTunedModel ensures that no value is present for FineTunedModel, not even an explicit nil
### GetSuffix

`func (o *ClassifierJobOut) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *ClassifierJobOut) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *ClassifierJobOut) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *ClassifierJobOut) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *ClassifierJobOut) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *ClassifierJobOut) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIntegrations

`func (o *ClassifierJobOut) GetIntegrations() []ClassifierJobOutIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ClassifierJobOut) GetIntegrationsOk() (*[]ClassifierJobOutIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ClassifierJobOut) SetIntegrations(v []ClassifierJobOutIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ClassifierJobOut) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *ClassifierJobOut) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *ClassifierJobOut) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetTrainedTokens

`func (o *ClassifierJobOut) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *ClassifierJobOut) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *ClassifierJobOut) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.

### HasTrainedTokens

`func (o *ClassifierJobOut) HasTrainedTokens() bool`

HasTrainedTokens returns a boolean if a field has been set.

### SetTrainedTokensNil

`func (o *ClassifierJobOut) SetTrainedTokensNil(b bool)`

 SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil

### UnsetTrainedTokens
`func (o *ClassifierJobOut) UnsetTrainedTokens()`

UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
### GetMetadata

`func (o *ClassifierJobOut) GetMetadata() JobMetadataOut`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClassifierJobOut) GetMetadataOk() (*JobMetadataOut, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClassifierJobOut) SetMetadata(v JobMetadataOut)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClassifierJobOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ClassifierJobOut) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ClassifierJobOut) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetJobType

`func (o *ClassifierJobOut) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ClassifierJobOut) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ClassifierJobOut) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ClassifierJobOut) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperparameters

`func (o *ClassifierJobOut) GetHyperparameters() ClassifierTrainingParameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *ClassifierJobOut) GetHyperparametersOk() (*ClassifierTrainingParameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *ClassifierJobOut) SetHyperparameters(v ClassifierTrainingParameters)`

SetHyperparameters sets Hyperparameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


