# JobIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**FineTuneableModel**](FineTuneableModel.md) |  | 
**TrainingFiles** | Pointer to [**[]TrainingFile**](TrainingFile.md) |  | [optional] [default to []]
**ValidationFiles** | Pointer to **[]string** |  | [optional] 
**Hyperparameters** | [**TrainingParametersIn**](TrainingParametersIn.md) |  | 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Integrations** | Pointer to [**[]JobInIntegrationsInner**](JobInIntegrationsInner.md) |  | [optional] 
**Repositories** | Pointer to [**[]JobInRepositoriesInner**](JobInRepositoriesInner.md) |  | [optional] [default to []]
**AutoStart** | Pointer to **bool** | This field will be required in a future release. | [optional] 

## Methods

### NewJobIn

`func NewJobIn(model FineTuneableModel, hyperparameters TrainingParametersIn, ) *JobIn`

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
### GetHyperparameters

`func (o *JobIn) GetHyperparameters() TrainingParametersIn`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *JobIn) GetHyperparametersOk() (*TrainingParametersIn, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *JobIn) SetHyperparameters(v TrainingParametersIn)`

SetHyperparameters sets Hyperparameters field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


