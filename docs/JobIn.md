# JobIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**TrainingFiles** | Pointer to [**[]TrainingFile**](TrainingFile.md) |  | [optional] [default to []]
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Integrations** | Pointer to [**[]JobInIntegrationsInner**](JobInIntegrationsInner.md) |  | [optional] 
**AutoStart** | Pointer to **bool** | This field will be required in a future release. | [optional] 
**InvalidSampleSkipPercentage** | Pointer to **float32** |  | [optional] [default to 0]
**JobType** | Pointer to [**NullableFineTuneableModelType**](FineTuneableModelType.md) |  | [optional] 
**Hyperparameters** | [**Hyperparameters**](Hyperparameters.md) |  | 
**Repositories** | Pointer to [**[]JobInRepositoriesInner**](JobInRepositoriesInner.md) |  | [optional] 
**ClassifierTargets** | Pointer to [**[]ClassifierTargetIn**](ClassifierTargetIn.md) |  | [optional] 

## Methods

### NewJobIn

`func NewJobIn(model FineTuneableModel, hyperparameters Hyperparameters, ) *JobIn`

NewJobIn instantiates a new JobIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInWithDefaults

`func NewJobInWithDefaults() *JobIn`

NewJobInWithDefaults instantiates a new JobIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *JobIn) GetModel() FineTuneableModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JobIn) GetModelOk() (*FineTuneableModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JobIn) SetModel(v FineTuneableModel)`

SetModel sets Model field to given value.


### GetTrainingFiles

`func (o *JobIn) GetTrainingFiles() []TrainingFile`

GetTrainingFiles returns the TrainingFiles field if non-nil, zero value otherwise.

### GetTrainingFilesOk

`func (o *JobIn) GetTrainingFilesOk() (*[]TrainingFile, bool)`

GetTrainingFilesOk returns a tuple with the TrainingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFiles

`func (o *JobIn) SetTrainingFiles(v []TrainingFile)`

SetTrainingFiles sets TrainingFiles field to given value.

### HasTrainingFiles

`func (o *JobIn) HasTrainingFiles() bool`

HasTrainingFiles returns a boolean if a field has been set.

### GetValidationFiles

`func (o *JobIn) GetValidationFiles() []string`

GetValidationFiles returns the ValidationFiles field if non-nil, zero value otherwise.

### GetValidationFilesOk

`func (o *JobIn) GetValidationFilesOk() (*[]string, bool)`

GetValidationFilesOk returns a tuple with the ValidationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFiles

`func (o *JobIn) SetValidationFiles(v []string)`

SetValidationFiles sets ValidationFiles field to given value.

### HasValidationFiles

`func (o *JobIn) HasValidationFiles() bool`

HasValidationFiles returns a boolean if a field has been set.

### SetValidationFilesNil

`func (o *JobIn) SetValidationFilesNil(b bool)`

 SetValidationFilesNil sets the value for ValidationFiles to be an explicit nil

### UnsetValidationFiles
`func (o *JobIn) UnsetValidationFiles()`

UnsetValidationFiles ensures that no value is present for ValidationFiles, not even an explicit nil
### GetSuffix

`func (o *JobIn) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *JobIn) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *JobIn) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *JobIn) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *JobIn) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *JobIn) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIntegrations

`func (o *JobIn) GetIntegrations() []JobInIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *JobIn) GetIntegrationsOk() (*[]JobInIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *JobIn) SetIntegrations(v []JobInIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *JobIn) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *JobIn) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *JobIn) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetAutoStart

`func (o *JobIn) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *JobIn) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *JobIn) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *JobIn) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetInvalidSampleSkipPercentage

`func (o *JobIn) GetInvalidSampleSkipPercentage() float32`

GetInvalidSampleSkipPercentage returns the InvalidSampleSkipPercentage field if non-nil, zero value otherwise.

### GetInvalidSampleSkipPercentageOk

`func (o *JobIn) GetInvalidSampleSkipPercentageOk() (*float32, bool)`

GetInvalidSampleSkipPercentageOk returns a tuple with the InvalidSampleSkipPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSampleSkipPercentage

`func (o *JobIn) SetInvalidSampleSkipPercentage(v float32)`

SetInvalidSampleSkipPercentage sets InvalidSampleSkipPercentage field to given value.

### HasInvalidSampleSkipPercentage

`func (o *JobIn) HasInvalidSampleSkipPercentage() bool`

HasInvalidSampleSkipPercentage returns a boolean if a field has been set.

### GetJobType

`func (o *JobIn) GetJobType() FineTuneableModelType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobIn) GetJobTypeOk() (*FineTuneableModelType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobIn) SetJobType(v FineTuneableModelType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *JobIn) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *JobIn) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *JobIn) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetHyperparameters

`func (o *JobIn) GetHyperparameters() Hyperparameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *JobIn) GetHyperparametersOk() (*Hyperparameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *JobIn) SetHyperparameters(v Hyperparameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetRepositories

`func (o *JobIn) GetRepositories() []JobInRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *JobIn) GetRepositoriesOk() (*[]JobInRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *JobIn) SetRepositories(v []JobInRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *JobIn) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### SetRepositoriesNil

`func (o *JobIn) SetRepositoriesNil(b bool)`

 SetRepositoriesNil sets the value for Repositories to be an explicit nil

### UnsetRepositories
`func (o *JobIn) UnsetRepositories()`

UnsetRepositories ensures that no value is present for Repositories, not even an explicit nil
### GetClassifierTargets

`func (o *JobIn) GetClassifierTargets() []ClassifierTargetIn`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *JobIn) GetClassifierTargetsOk() (*[]ClassifierTargetIn, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *JobIn) SetClassifierTargets(v []ClassifierTargetIn)`

SetClassifierTargets sets ClassifierTargets field to given value.

### HasClassifierTargets

`func (o *JobIn) HasClassifierTargets() bool`

HasClassifierTargets returns a boolean if a field has been set.

### SetClassifierTargetsNil

`func (o *JobIn) SetClassifierTargetsNil(b bool)`

 SetClassifierTargetsNil sets the value for ClassifierTargets to be an explicit nil

### UnsetClassifierTargets
`func (o *JobIn) UnsetClassifierTargets()`

UnsetClassifierTargets ensures that no value is present for ClassifierTargets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


