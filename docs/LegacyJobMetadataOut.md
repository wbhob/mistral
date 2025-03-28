# LegacyJobMetadataOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Cost** | Pointer to **NullableFloat32** |  | [optional] 
**CostCurrency** | Pointer to **NullableString** |  | [optional] 
**TrainTokensPerStep** | Pointer to **NullableInt32** |  | [optional] 
**TrainTokens** | Pointer to **NullableInt32** |  | [optional] 
**DataTokens** | Pointer to **NullableInt32** |  | [optional] 
**EstimatedStartTime** | Pointer to **NullableInt32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] [default to true]
**Details** | **string** |  | 
**Epochs** | Pointer to **NullableFloat32** |  | [optional] 
**TrainingSteps** | Pointer to **NullableInt32** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "job.metadata"]

## Methods

### NewLegacyJobMetadataOut

`func NewLegacyJobMetadataOut(details string, ) *LegacyJobMetadataOut`

NewLegacyJobMetadataOut instantiates a new LegacyJobMetadataOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyJobMetadataOutWithDefaults

`func NewLegacyJobMetadataOutWithDefaults() *LegacyJobMetadataOut`

NewLegacyJobMetadataOutWithDefaults instantiates a new LegacyJobMetadataOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDurationSeconds

`func (o *LegacyJobMetadataOut) GetExpectedDurationSeconds() int32`

GetExpectedDurationSeconds returns the ExpectedDurationSeconds field if non-nil, zero value otherwise.

### GetExpectedDurationSecondsOk

`func (o *LegacyJobMetadataOut) GetExpectedDurationSecondsOk() (*int32, bool)`

GetExpectedDurationSecondsOk returns a tuple with the ExpectedDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDurationSeconds

`func (o *LegacyJobMetadataOut) SetExpectedDurationSeconds(v int32)`

SetExpectedDurationSeconds sets ExpectedDurationSeconds field to given value.

### HasExpectedDurationSeconds

`func (o *LegacyJobMetadataOut) HasExpectedDurationSeconds() bool`

HasExpectedDurationSeconds returns a boolean if a field has been set.

### SetExpectedDurationSecondsNil

`func (o *LegacyJobMetadataOut) SetExpectedDurationSecondsNil(b bool)`

 SetExpectedDurationSecondsNil sets the value for ExpectedDurationSeconds to be an explicit nil

### UnsetExpectedDurationSeconds
`func (o *LegacyJobMetadataOut) UnsetExpectedDurationSeconds()`

UnsetExpectedDurationSeconds ensures that no value is present for ExpectedDurationSeconds, not even an explicit nil
### GetCost

`func (o *LegacyJobMetadataOut) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *LegacyJobMetadataOut) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *LegacyJobMetadataOut) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *LegacyJobMetadataOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *LegacyJobMetadataOut) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *LegacyJobMetadataOut) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostCurrency

`func (o *LegacyJobMetadataOut) GetCostCurrency() string`

GetCostCurrency returns the CostCurrency field if non-nil, zero value otherwise.

### GetCostCurrencyOk

`func (o *LegacyJobMetadataOut) GetCostCurrencyOk() (*string, bool)`

GetCostCurrencyOk returns a tuple with the CostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCurrency

`func (o *LegacyJobMetadataOut) SetCostCurrency(v string)`

SetCostCurrency sets CostCurrency field to given value.

### HasCostCurrency

`func (o *LegacyJobMetadataOut) HasCostCurrency() bool`

HasCostCurrency returns a boolean if a field has been set.

### SetCostCurrencyNil

`func (o *LegacyJobMetadataOut) SetCostCurrencyNil(b bool)`

 SetCostCurrencyNil sets the value for CostCurrency to be an explicit nil

### UnsetCostCurrency
`func (o *LegacyJobMetadataOut) UnsetCostCurrency()`

UnsetCostCurrency ensures that no value is present for CostCurrency, not even an explicit nil
### GetTrainTokensPerStep

`func (o *LegacyJobMetadataOut) GetTrainTokensPerStep() int32`

GetTrainTokensPerStep returns the TrainTokensPerStep field if non-nil, zero value otherwise.

### GetTrainTokensPerStepOk

`func (o *LegacyJobMetadataOut) GetTrainTokensPerStepOk() (*int32, bool)`

GetTrainTokensPerStepOk returns a tuple with the TrainTokensPerStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokensPerStep

`func (o *LegacyJobMetadataOut) SetTrainTokensPerStep(v int32)`

SetTrainTokensPerStep sets TrainTokensPerStep field to given value.

### HasTrainTokensPerStep

`func (o *LegacyJobMetadataOut) HasTrainTokensPerStep() bool`

HasTrainTokensPerStep returns a boolean if a field has been set.

### SetTrainTokensPerStepNil

`func (o *LegacyJobMetadataOut) SetTrainTokensPerStepNil(b bool)`

 SetTrainTokensPerStepNil sets the value for TrainTokensPerStep to be an explicit nil

### UnsetTrainTokensPerStep
`func (o *LegacyJobMetadataOut) UnsetTrainTokensPerStep()`

UnsetTrainTokensPerStep ensures that no value is present for TrainTokensPerStep, not even an explicit nil
### GetTrainTokens

`func (o *LegacyJobMetadataOut) GetTrainTokens() int32`

GetTrainTokens returns the TrainTokens field if non-nil, zero value otherwise.

### GetTrainTokensOk

`func (o *LegacyJobMetadataOut) GetTrainTokensOk() (*int32, bool)`

GetTrainTokensOk returns a tuple with the TrainTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokens

`func (o *LegacyJobMetadataOut) SetTrainTokens(v int32)`

SetTrainTokens sets TrainTokens field to given value.

### HasTrainTokens

`func (o *LegacyJobMetadataOut) HasTrainTokens() bool`

HasTrainTokens returns a boolean if a field has been set.

### SetTrainTokensNil

`func (o *LegacyJobMetadataOut) SetTrainTokensNil(b bool)`

 SetTrainTokensNil sets the value for TrainTokens to be an explicit nil

### UnsetTrainTokens
`func (o *LegacyJobMetadataOut) UnsetTrainTokens()`

UnsetTrainTokens ensures that no value is present for TrainTokens, not even an explicit nil
### GetDataTokens

`func (o *LegacyJobMetadataOut) GetDataTokens() int32`

GetDataTokens returns the DataTokens field if non-nil, zero value otherwise.

### GetDataTokensOk

`func (o *LegacyJobMetadataOut) GetDataTokensOk() (*int32, bool)`

GetDataTokensOk returns a tuple with the DataTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTokens

`func (o *LegacyJobMetadataOut) SetDataTokens(v int32)`

SetDataTokens sets DataTokens field to given value.

### HasDataTokens

`func (o *LegacyJobMetadataOut) HasDataTokens() bool`

HasDataTokens returns a boolean if a field has been set.

### SetDataTokensNil

`func (o *LegacyJobMetadataOut) SetDataTokensNil(b bool)`

 SetDataTokensNil sets the value for DataTokens to be an explicit nil

### UnsetDataTokens
`func (o *LegacyJobMetadataOut) UnsetDataTokens()`

UnsetDataTokens ensures that no value is present for DataTokens, not even an explicit nil
### GetEstimatedStartTime

`func (o *LegacyJobMetadataOut) GetEstimatedStartTime() int32`

GetEstimatedStartTime returns the EstimatedStartTime field if non-nil, zero value otherwise.

### GetEstimatedStartTimeOk

`func (o *LegacyJobMetadataOut) GetEstimatedStartTimeOk() (*int32, bool)`

GetEstimatedStartTimeOk returns a tuple with the EstimatedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartTime

`func (o *LegacyJobMetadataOut) SetEstimatedStartTime(v int32)`

SetEstimatedStartTime sets EstimatedStartTime field to given value.

### HasEstimatedStartTime

`func (o *LegacyJobMetadataOut) HasEstimatedStartTime() bool`

HasEstimatedStartTime returns a boolean if a field has been set.

### SetEstimatedStartTimeNil

`func (o *LegacyJobMetadataOut) SetEstimatedStartTimeNil(b bool)`

 SetEstimatedStartTimeNil sets the value for EstimatedStartTime to be an explicit nil

### UnsetEstimatedStartTime
`func (o *LegacyJobMetadataOut) UnsetEstimatedStartTime()`

UnsetEstimatedStartTime ensures that no value is present for EstimatedStartTime, not even an explicit nil
### GetDeprecated

`func (o *LegacyJobMetadataOut) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *LegacyJobMetadataOut) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *LegacyJobMetadataOut) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *LegacyJobMetadataOut) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDetails

`func (o *LegacyJobMetadataOut) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LegacyJobMetadataOut) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LegacyJobMetadataOut) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetEpochs

`func (o *LegacyJobMetadataOut) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *LegacyJobMetadataOut) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *LegacyJobMetadataOut) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *LegacyJobMetadataOut) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### SetEpochsNil

`func (o *LegacyJobMetadataOut) SetEpochsNil(b bool)`

 SetEpochsNil sets the value for Epochs to be an explicit nil

### UnsetEpochs
`func (o *LegacyJobMetadataOut) UnsetEpochs()`

UnsetEpochs ensures that no value is present for Epochs, not even an explicit nil
### GetTrainingSteps

`func (o *LegacyJobMetadataOut) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *LegacyJobMetadataOut) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *LegacyJobMetadataOut) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *LegacyJobMetadataOut) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.

### SetTrainingStepsNil

`func (o *LegacyJobMetadataOut) SetTrainingStepsNil(b bool)`

 SetTrainingStepsNil sets the value for TrainingSteps to be an explicit nil

### UnsetTrainingSteps
`func (o *LegacyJobMetadataOut) UnsetTrainingSteps()`

UnsetTrainingSteps ensures that no value is present for TrainingSteps, not even an explicit nil
### GetObject

`func (o *LegacyJobMetadataOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *LegacyJobMetadataOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *LegacyJobMetadataOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *LegacyJobMetadataOut) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


