# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Object** | Pointer to **string** |  | [optional] [default to "job.metadata"]

## Methods

### NewResponse

`func NewResponse(details string, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


