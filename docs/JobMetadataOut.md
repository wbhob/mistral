# JobMetadataOut

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

## Methods

### NewJobMetadataOut

`func NewJobMetadataOut() *JobMetadataOut`

NewJobMetadataOut instantiates a new JobMetadataOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobMetadataOutWithDefaults

`func NewJobMetadataOutWithDefaults() *JobMetadataOut`

NewJobMetadataOutWithDefaults instantiates a new JobMetadataOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDurationSeconds

`func (o *JobMetadataOut) GetExpectedDurationSeconds() int32`

GetExpectedDurationSeconds returns the ExpectedDurationSeconds field if non-nil, zero value otherwise.

### GetExpectedDurationSecondsOk

`func (o *JobMetadataOut) GetExpectedDurationSecondsOk() (*int32, bool)`

GetExpectedDurationSecondsOk returns a tuple with the ExpectedDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDurationSeconds

`func (o *JobMetadataOut) SetExpectedDurationSeconds(v int32)`

SetExpectedDurationSeconds sets ExpectedDurationSeconds field to given value.

### HasExpectedDurationSeconds

`func (o *JobMetadataOut) HasExpectedDurationSeconds() bool`

HasExpectedDurationSeconds returns a boolean if a field has been set.

### SetExpectedDurationSecondsNil

`func (o *JobMetadataOut) SetExpectedDurationSecondsNil(b bool)`

 SetExpectedDurationSecondsNil sets the value for ExpectedDurationSeconds to be an explicit nil

### UnsetExpectedDurationSeconds
`func (o *JobMetadataOut) UnsetExpectedDurationSeconds()`

UnsetExpectedDurationSeconds ensures that no value is present for ExpectedDurationSeconds, not even an explicit nil
### GetCost

`func (o *JobMetadataOut) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *JobMetadataOut) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *JobMetadataOut) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *JobMetadataOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *JobMetadataOut) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *JobMetadataOut) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostCurrency

`func (o *JobMetadataOut) GetCostCurrency() string`

GetCostCurrency returns the CostCurrency field if non-nil, zero value otherwise.

### GetCostCurrencyOk

`func (o *JobMetadataOut) GetCostCurrencyOk() (*string, bool)`

GetCostCurrencyOk returns a tuple with the CostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCurrency

`func (o *JobMetadataOut) SetCostCurrency(v string)`

SetCostCurrency sets CostCurrency field to given value.

### HasCostCurrency

`func (o *JobMetadataOut) HasCostCurrency() bool`

HasCostCurrency returns a boolean if a field has been set.

### SetCostCurrencyNil

`func (o *JobMetadataOut) SetCostCurrencyNil(b bool)`

 SetCostCurrencyNil sets the value for CostCurrency to be an explicit nil

### UnsetCostCurrency
`func (o *JobMetadataOut) UnsetCostCurrency()`

UnsetCostCurrency ensures that no value is present for CostCurrency, not even an explicit nil
### GetTrainTokensPerStep

`func (o *JobMetadataOut) GetTrainTokensPerStep() int32`

GetTrainTokensPerStep returns the TrainTokensPerStep field if non-nil, zero value otherwise.

### GetTrainTokensPerStepOk

`func (o *JobMetadataOut) GetTrainTokensPerStepOk() (*int32, bool)`

GetTrainTokensPerStepOk returns a tuple with the TrainTokensPerStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokensPerStep

`func (o *JobMetadataOut) SetTrainTokensPerStep(v int32)`

SetTrainTokensPerStep sets TrainTokensPerStep field to given value.

### HasTrainTokensPerStep

`func (o *JobMetadataOut) HasTrainTokensPerStep() bool`

HasTrainTokensPerStep returns a boolean if a field has been set.

### SetTrainTokensPerStepNil

`func (o *JobMetadataOut) SetTrainTokensPerStepNil(b bool)`

 SetTrainTokensPerStepNil sets the value for TrainTokensPerStep to be an explicit nil

### UnsetTrainTokensPerStep
`func (o *JobMetadataOut) UnsetTrainTokensPerStep()`

UnsetTrainTokensPerStep ensures that no value is present for TrainTokensPerStep, not even an explicit nil
### GetTrainTokens

`func (o *JobMetadataOut) GetTrainTokens() int32`

GetTrainTokens returns the TrainTokens field if non-nil, zero value otherwise.

### GetTrainTokensOk

`func (o *JobMetadataOut) GetTrainTokensOk() (*int32, bool)`

GetTrainTokensOk returns a tuple with the TrainTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainTokens

`func (o *JobMetadataOut) SetTrainTokens(v int32)`

SetTrainTokens sets TrainTokens field to given value.

### HasTrainTokens

`func (o *JobMetadataOut) HasTrainTokens() bool`

HasTrainTokens returns a boolean if a field has been set.

### SetTrainTokensNil

`func (o *JobMetadataOut) SetTrainTokensNil(b bool)`

 SetTrainTokensNil sets the value for TrainTokens to be an explicit nil

### UnsetTrainTokens
`func (o *JobMetadataOut) UnsetTrainTokens()`

UnsetTrainTokens ensures that no value is present for TrainTokens, not even an explicit nil
### GetDataTokens

`func (o *JobMetadataOut) GetDataTokens() int32`

GetDataTokens returns the DataTokens field if non-nil, zero value otherwise.

### GetDataTokensOk

`func (o *JobMetadataOut) GetDataTokensOk() (*int32, bool)`

GetDataTokensOk returns a tuple with the DataTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTokens

`func (o *JobMetadataOut) SetDataTokens(v int32)`

SetDataTokens sets DataTokens field to given value.

### HasDataTokens

`func (o *JobMetadataOut) HasDataTokens() bool`

HasDataTokens returns a boolean if a field has been set.

### SetDataTokensNil

`func (o *JobMetadataOut) SetDataTokensNil(b bool)`

 SetDataTokensNil sets the value for DataTokens to be an explicit nil

### UnsetDataTokens
`func (o *JobMetadataOut) UnsetDataTokens()`

UnsetDataTokens ensures that no value is present for DataTokens, not even an explicit nil
### GetEstimatedStartTime

`func (o *JobMetadataOut) GetEstimatedStartTime() int32`

GetEstimatedStartTime returns the EstimatedStartTime field if non-nil, zero value otherwise.

### GetEstimatedStartTimeOk

`func (o *JobMetadataOut) GetEstimatedStartTimeOk() (*int32, bool)`

GetEstimatedStartTimeOk returns a tuple with the EstimatedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartTime

`func (o *JobMetadataOut) SetEstimatedStartTime(v int32)`

SetEstimatedStartTime sets EstimatedStartTime field to given value.

### HasEstimatedStartTime

`func (o *JobMetadataOut) HasEstimatedStartTime() bool`

HasEstimatedStartTime returns a boolean if a field has been set.

### SetEstimatedStartTimeNil

`func (o *JobMetadataOut) SetEstimatedStartTimeNil(b bool)`

 SetEstimatedStartTimeNil sets the value for EstimatedStartTime to be an explicit nil

### UnsetEstimatedStartTime
`func (o *JobMetadataOut) UnsetEstimatedStartTime()`

UnsetEstimatedStartTime ensures that no value is present for EstimatedStartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


