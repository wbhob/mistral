# CompletionArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stop** | Pointer to [**NullableCompletionArgsStop**](CompletionArgsStop.md) |  | [optional] 
**PresencePenalty** | Pointer to **NullableFloat32** |  | [optional] 
**FrequencyPenalty** | Pointer to **NullableFloat32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] [default to 0.3]
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** |  | [optional] 
**Prediction** | Pointer to [**NullablePrediction**](Prediction.md) |  | [optional] 
**ResponseFormat** | Pointer to [**NullableResponseFormat**](ResponseFormat.md) |  | [optional] 
**ToolChoice** | Pointer to [**ToolChoiceEnum**](ToolChoiceEnum.md) |  | [optional] 

## Methods

### NewCompletionArgs

`func NewCompletionArgs() *CompletionArgs`

NewCompletionArgs instantiates a new CompletionArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionArgsWithDefaults

`func NewCompletionArgsWithDefaults() *CompletionArgs`

NewCompletionArgsWithDefaults instantiates a new CompletionArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStop

`func (o *CompletionArgs) GetStop() CompletionArgsStop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *CompletionArgs) GetStopOk() (*CompletionArgsStop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *CompletionArgs) SetStop(v CompletionArgsStop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *CompletionArgs) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *CompletionArgs) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *CompletionArgs) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetPresencePenalty

`func (o *CompletionArgs) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *CompletionArgs) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *CompletionArgs) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *CompletionArgs) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *CompletionArgs) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *CompletionArgs) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetFrequencyPenalty

`func (o *CompletionArgs) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *CompletionArgs) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *CompletionArgs) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *CompletionArgs) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *CompletionArgs) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *CompletionArgs) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetTemperature

`func (o *CompletionArgs) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CompletionArgs) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CompletionArgs) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CompletionArgs) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *CompletionArgs) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *CompletionArgs) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *CompletionArgs) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *CompletionArgs) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *CompletionArgs) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *CompletionArgs) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetMaxTokens

`func (o *CompletionArgs) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *CompletionArgs) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *CompletionArgs) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *CompletionArgs) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *CompletionArgs) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *CompletionArgs) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetRandomSeed

`func (o *CompletionArgs) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *CompletionArgs) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *CompletionArgs) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *CompletionArgs) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *CompletionArgs) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *CompletionArgs) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetPrediction

`func (o *CompletionArgs) GetPrediction() Prediction`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *CompletionArgs) GetPredictionOk() (*Prediction, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *CompletionArgs) SetPrediction(v Prediction)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *CompletionArgs) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### SetPredictionNil

`func (o *CompletionArgs) SetPredictionNil(b bool)`

 SetPredictionNil sets the value for Prediction to be an explicit nil

### UnsetPrediction
`func (o *CompletionArgs) UnsetPrediction()`

UnsetPrediction ensures that no value is present for Prediction, not even an explicit nil
### GetResponseFormat

`func (o *CompletionArgs) GetResponseFormat() ResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *CompletionArgs) GetResponseFormatOk() (*ResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *CompletionArgs) SetResponseFormat(v ResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *CompletionArgs) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### SetResponseFormatNil

`func (o *CompletionArgs) SetResponseFormatNil(b bool)`

 SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil

### UnsetResponseFormat
`func (o *CompletionArgs) UnsetResponseFormat()`

UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
### GetToolChoice

`func (o *CompletionArgs) GetToolChoice() ToolChoiceEnum`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *CompletionArgs) GetToolChoiceOk() (*ToolChoiceEnum, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *CompletionArgs) SetToolChoice(v ToolChoiceEnum)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *CompletionArgs) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


