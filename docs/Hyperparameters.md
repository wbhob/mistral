# Hyperparameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainingSteps** | Pointer to **int32** |  | [optional] 
**LearningRate** | Pointer to **float32** | A parameter describing how much to adjust the pre-trained model&#39;s weights in response to the estimated error each time the weights are updated during the fine-tuning process. | [optional] [default to 1.0E-4]
**WeightDecay** | Pointer to **float32** |  | [optional] 
**WarmupFraction** | Pointer to **float32** |  | [optional] 
**Epochs** | Pointer to **float32** |  | [optional] 
**SeqLen** | Pointer to **int32** |  | [optional] 
**FimRatio** | Pointer to **float32** |  | [optional] 

## Methods

### NewHyperparameters

`func NewHyperparameters() *Hyperparameters`

NewHyperparameters instantiates a new Hyperparameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperparametersWithDefaults

`func NewHyperparametersWithDefaults() *Hyperparameters`

NewHyperparametersWithDefaults instantiates a new Hyperparameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainingSteps

`func (o *Hyperparameters) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *Hyperparameters) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *Hyperparameters) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *Hyperparameters) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.

### GetLearningRate

`func (o *Hyperparameters) GetLearningRate() float32`

GetLearningRate returns the LearningRate field if non-nil, zero value otherwise.

### GetLearningRateOk

`func (o *Hyperparameters) GetLearningRateOk() (*float32, bool)`

GetLearningRateOk returns a tuple with the LearningRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearningRate

`func (o *Hyperparameters) SetLearningRate(v float32)`

SetLearningRate sets LearningRate field to given value.

### HasLearningRate

`func (o *Hyperparameters) HasLearningRate() bool`

HasLearningRate returns a boolean if a field has been set.

### GetWeightDecay

`func (o *Hyperparameters) GetWeightDecay() float32`

GetWeightDecay returns the WeightDecay field if non-nil, zero value otherwise.

### GetWeightDecayOk

`func (o *Hyperparameters) GetWeightDecayOk() (*float32, bool)`

GetWeightDecayOk returns a tuple with the WeightDecay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightDecay

`func (o *Hyperparameters) SetWeightDecay(v float32)`

SetWeightDecay sets WeightDecay field to given value.

### HasWeightDecay

`func (o *Hyperparameters) HasWeightDecay() bool`

HasWeightDecay returns a boolean if a field has been set.

### GetWarmupFraction

`func (o *Hyperparameters) GetWarmupFraction() float32`

GetWarmupFraction returns the WarmupFraction field if non-nil, zero value otherwise.

### GetWarmupFractionOk

`func (o *Hyperparameters) GetWarmupFractionOk() (*float32, bool)`

GetWarmupFractionOk returns a tuple with the WarmupFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupFraction

`func (o *Hyperparameters) SetWarmupFraction(v float32)`

SetWarmupFraction sets WarmupFraction field to given value.

### HasWarmupFraction

`func (o *Hyperparameters) HasWarmupFraction() bool`

HasWarmupFraction returns a boolean if a field has been set.

### GetEpochs

`func (o *Hyperparameters) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *Hyperparameters) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *Hyperparameters) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *Hyperparameters) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### GetSeqLen

`func (o *Hyperparameters) GetSeqLen() int32`

GetSeqLen returns the SeqLen field if non-nil, zero value otherwise.

### GetSeqLenOk

`func (o *Hyperparameters) GetSeqLenOk() (*int32, bool)`

GetSeqLenOk returns a tuple with the SeqLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqLen

`func (o *Hyperparameters) SetSeqLen(v int32)`

SetSeqLen sets SeqLen field to given value.

### HasSeqLen

`func (o *Hyperparameters) HasSeqLen() bool`

HasSeqLen returns a boolean if a field has been set.

### GetFimRatio

`func (o *Hyperparameters) GetFimRatio() float32`

GetFimRatio returns the FimRatio field if non-nil, zero value otherwise.

### GetFimRatioOk

`func (o *Hyperparameters) GetFimRatioOk() (*float32, bool)`

GetFimRatioOk returns a tuple with the FimRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFimRatio

`func (o *Hyperparameters) SetFimRatio(v float32)`

SetFimRatio sets FimRatio field to given value.

### HasFimRatio

`func (o *Hyperparameters) HasFimRatio() bool`

HasFimRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


