# ClassifierTrainingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainingSteps** | Pointer to **NullableInt32** |  | [optional] 
**LearningRate** | Pointer to **float32** |  | [optional] [default to 1.0E-4]
**WeightDecay** | Pointer to **NullableFloat32** |  | [optional] 
**WarmupFraction** | Pointer to **NullableFloat32** |  | [optional] 
**Epochs** | Pointer to **NullableFloat32** |  | [optional] 
**SeqLen** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewClassifierTrainingParameters

`func NewClassifierTrainingParameters() *ClassifierTrainingParameters`

NewClassifierTrainingParameters instantiates a new ClassifierTrainingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierTrainingParametersWithDefaults

`func NewClassifierTrainingParametersWithDefaults() *ClassifierTrainingParameters`

NewClassifierTrainingParametersWithDefaults instantiates a new ClassifierTrainingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainingSteps

`func (o *ClassifierTrainingParameters) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *ClassifierTrainingParameters) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *ClassifierTrainingParameters) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *ClassifierTrainingParameters) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.

### SetTrainingStepsNil

`func (o *ClassifierTrainingParameters) SetTrainingStepsNil(b bool)`

 SetTrainingStepsNil sets the value for TrainingSteps to be an explicit nil

### UnsetTrainingSteps
`func (o *ClassifierTrainingParameters) UnsetTrainingSteps()`

UnsetTrainingSteps ensures that no value is present for TrainingSteps, not even an explicit nil
### GetLearningRate

`func (o *ClassifierTrainingParameters) GetLearningRate() float32`

GetLearningRate returns the LearningRate field if non-nil, zero value otherwise.

### GetLearningRateOk

`func (o *ClassifierTrainingParameters) GetLearningRateOk() (*float32, bool)`

GetLearningRateOk returns a tuple with the LearningRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearningRate

`func (o *ClassifierTrainingParameters) SetLearningRate(v float32)`

SetLearningRate sets LearningRate field to given value.

### HasLearningRate

`func (o *ClassifierTrainingParameters) HasLearningRate() bool`

HasLearningRate returns a boolean if a field has been set.

### GetWeightDecay

`func (o *ClassifierTrainingParameters) GetWeightDecay() float32`

GetWeightDecay returns the WeightDecay field if non-nil, zero value otherwise.

### GetWeightDecayOk

`func (o *ClassifierTrainingParameters) GetWeightDecayOk() (*float32, bool)`

GetWeightDecayOk returns a tuple with the WeightDecay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightDecay

`func (o *ClassifierTrainingParameters) SetWeightDecay(v float32)`

SetWeightDecay sets WeightDecay field to given value.

### HasWeightDecay

`func (o *ClassifierTrainingParameters) HasWeightDecay() bool`

HasWeightDecay returns a boolean if a field has been set.

### SetWeightDecayNil

`func (o *ClassifierTrainingParameters) SetWeightDecayNil(b bool)`

 SetWeightDecayNil sets the value for WeightDecay to be an explicit nil

### UnsetWeightDecay
`func (o *ClassifierTrainingParameters) UnsetWeightDecay()`

UnsetWeightDecay ensures that no value is present for WeightDecay, not even an explicit nil
### GetWarmupFraction

`func (o *ClassifierTrainingParameters) GetWarmupFraction() float32`

GetWarmupFraction returns the WarmupFraction field if non-nil, zero value otherwise.

### GetWarmupFractionOk

`func (o *ClassifierTrainingParameters) GetWarmupFractionOk() (*float32, bool)`

GetWarmupFractionOk returns a tuple with the WarmupFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupFraction

`func (o *ClassifierTrainingParameters) SetWarmupFraction(v float32)`

SetWarmupFraction sets WarmupFraction field to given value.

### HasWarmupFraction

`func (o *ClassifierTrainingParameters) HasWarmupFraction() bool`

HasWarmupFraction returns a boolean if a field has been set.

### SetWarmupFractionNil

`func (o *ClassifierTrainingParameters) SetWarmupFractionNil(b bool)`

 SetWarmupFractionNil sets the value for WarmupFraction to be an explicit nil

### UnsetWarmupFraction
`func (o *ClassifierTrainingParameters) UnsetWarmupFraction()`

UnsetWarmupFraction ensures that no value is present for WarmupFraction, not even an explicit nil
### GetEpochs

`func (o *ClassifierTrainingParameters) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *ClassifierTrainingParameters) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *ClassifierTrainingParameters) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *ClassifierTrainingParameters) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### SetEpochsNil

`func (o *ClassifierTrainingParameters) SetEpochsNil(b bool)`

 SetEpochsNil sets the value for Epochs to be an explicit nil

### UnsetEpochs
`func (o *ClassifierTrainingParameters) UnsetEpochs()`

UnsetEpochs ensures that no value is present for Epochs, not even an explicit nil
### GetSeqLen

`func (o *ClassifierTrainingParameters) GetSeqLen() int32`

GetSeqLen returns the SeqLen field if non-nil, zero value otherwise.

### GetSeqLenOk

`func (o *ClassifierTrainingParameters) GetSeqLenOk() (*int32, bool)`

GetSeqLenOk returns a tuple with the SeqLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqLen

`func (o *ClassifierTrainingParameters) SetSeqLen(v int32)`

SetSeqLen sets SeqLen field to given value.

### HasSeqLen

`func (o *ClassifierTrainingParameters) HasSeqLen() bool`

HasSeqLen returns a boolean if a field has been set.

### SetSeqLenNil

`func (o *ClassifierTrainingParameters) SetSeqLenNil(b bool)`

 SetSeqLenNil sets the value for SeqLen to be an explicit nil

### UnsetSeqLen
`func (o *ClassifierTrainingParameters) UnsetSeqLen()`

UnsetSeqLen ensures that no value is present for SeqLen, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


