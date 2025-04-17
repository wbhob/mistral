# CompletionTrainingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainingSteps** | Pointer to **NullableInt32** |  | [optional] 
**LearningRate** | Pointer to **float32** |  | [optional] [default to 1.0E-4]
**WeightDecay** | Pointer to **NullableFloat32** |  | [optional] 
**WarmupFraction** | Pointer to **NullableFloat32** |  | [optional] 
**Epochs** | Pointer to **NullableFloat32** |  | [optional] 
**SeqLen** | Pointer to **NullableInt32** |  | [optional] 
**FimRatio** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewCompletionTrainingParameters

`func NewCompletionTrainingParameters() *CompletionTrainingParameters`

NewCompletionTrainingParameters instantiates a new CompletionTrainingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionTrainingParametersWithDefaults

`func NewCompletionTrainingParametersWithDefaults() *CompletionTrainingParameters`

NewCompletionTrainingParametersWithDefaults instantiates a new CompletionTrainingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainingSteps

`func (o *CompletionTrainingParameters) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *CompletionTrainingParameters) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *CompletionTrainingParameters) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *CompletionTrainingParameters) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.

### SetTrainingStepsNil

`func (o *CompletionTrainingParameters) SetTrainingStepsNil(b bool)`

 SetTrainingStepsNil sets the value for TrainingSteps to be an explicit nil

### UnsetTrainingSteps
`func (o *CompletionTrainingParameters) UnsetTrainingSteps()`

UnsetTrainingSteps ensures that no value is present for TrainingSteps, not even an explicit nil
### GetLearningRate

`func (o *CompletionTrainingParameters) GetLearningRate() float32`

GetLearningRate returns the LearningRate field if non-nil, zero value otherwise.

### GetLearningRateOk

`func (o *CompletionTrainingParameters) GetLearningRateOk() (*float32, bool)`

GetLearningRateOk returns a tuple with the LearningRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearningRate

`func (o *CompletionTrainingParameters) SetLearningRate(v float32)`

SetLearningRate sets LearningRate field to given value.

### HasLearningRate

`func (o *CompletionTrainingParameters) HasLearningRate() bool`

HasLearningRate returns a boolean if a field has been set.

### GetWeightDecay

`func (o *CompletionTrainingParameters) GetWeightDecay() float32`

GetWeightDecay returns the WeightDecay field if non-nil, zero value otherwise.

### GetWeightDecayOk

`func (o *CompletionTrainingParameters) GetWeightDecayOk() (*float32, bool)`

GetWeightDecayOk returns a tuple with the WeightDecay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightDecay

`func (o *CompletionTrainingParameters) SetWeightDecay(v float32)`

SetWeightDecay sets WeightDecay field to given value.

### HasWeightDecay

`func (o *CompletionTrainingParameters) HasWeightDecay() bool`

HasWeightDecay returns a boolean if a field has been set.

### SetWeightDecayNil

`func (o *CompletionTrainingParameters) SetWeightDecayNil(b bool)`

 SetWeightDecayNil sets the value for WeightDecay to be an explicit nil

### UnsetWeightDecay
`func (o *CompletionTrainingParameters) UnsetWeightDecay()`

UnsetWeightDecay ensures that no value is present for WeightDecay, not even an explicit nil
### GetWarmupFraction

`func (o *CompletionTrainingParameters) GetWarmupFraction() float32`

GetWarmupFraction returns the WarmupFraction field if non-nil, zero value otherwise.

### GetWarmupFractionOk

`func (o *CompletionTrainingParameters) GetWarmupFractionOk() (*float32, bool)`

GetWarmupFractionOk returns a tuple with the WarmupFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupFraction

`func (o *CompletionTrainingParameters) SetWarmupFraction(v float32)`

SetWarmupFraction sets WarmupFraction field to given value.

### HasWarmupFraction

`func (o *CompletionTrainingParameters) HasWarmupFraction() bool`

HasWarmupFraction returns a boolean if a field has been set.

### SetWarmupFractionNil

`func (o *CompletionTrainingParameters) SetWarmupFractionNil(b bool)`

 SetWarmupFractionNil sets the value for WarmupFraction to be an explicit nil

### UnsetWarmupFraction
`func (o *CompletionTrainingParameters) UnsetWarmupFraction()`

UnsetWarmupFraction ensures that no value is present for WarmupFraction, not even an explicit nil
### GetEpochs

`func (o *CompletionTrainingParameters) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *CompletionTrainingParameters) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *CompletionTrainingParameters) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *CompletionTrainingParameters) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### SetEpochsNil

`func (o *CompletionTrainingParameters) SetEpochsNil(b bool)`

 SetEpochsNil sets the value for Epochs to be an explicit nil

### UnsetEpochs
`func (o *CompletionTrainingParameters) UnsetEpochs()`

UnsetEpochs ensures that no value is present for Epochs, not even an explicit nil
### GetSeqLen

`func (o *CompletionTrainingParameters) GetSeqLen() int32`

GetSeqLen returns the SeqLen field if non-nil, zero value otherwise.

### GetSeqLenOk

`func (o *CompletionTrainingParameters) GetSeqLenOk() (*int32, bool)`

GetSeqLenOk returns a tuple with the SeqLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqLen

`func (o *CompletionTrainingParameters) SetSeqLen(v int32)`

SetSeqLen sets SeqLen field to given value.

### HasSeqLen

`func (o *CompletionTrainingParameters) HasSeqLen() bool`

HasSeqLen returns a boolean if a field has been set.

### SetSeqLenNil

`func (o *CompletionTrainingParameters) SetSeqLenNil(b bool)`

 SetSeqLenNil sets the value for SeqLen to be an explicit nil

### UnsetSeqLen
`func (o *CompletionTrainingParameters) UnsetSeqLen()`

UnsetSeqLen ensures that no value is present for SeqLen, not even an explicit nil
### GetFimRatio

`func (o *CompletionTrainingParameters) GetFimRatio() float32`

GetFimRatio returns the FimRatio field if non-nil, zero value otherwise.

### GetFimRatioOk

`func (o *CompletionTrainingParameters) GetFimRatioOk() (*float32, bool)`

GetFimRatioOk returns a tuple with the FimRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFimRatio

`func (o *CompletionTrainingParameters) SetFimRatio(v float32)`

SetFimRatio sets FimRatio field to given value.

### HasFimRatio

`func (o *CompletionTrainingParameters) HasFimRatio() bool`

HasFimRatio returns a boolean if a field has been set.

### SetFimRatioNil

`func (o *CompletionTrainingParameters) SetFimRatioNil(b bool)`

 SetFimRatioNil sets the value for FimRatio to be an explicit nil

### UnsetFimRatio
`func (o *CompletionTrainingParameters) UnsetFimRatio()`

UnsetFimRatio ensures that no value is present for FimRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


