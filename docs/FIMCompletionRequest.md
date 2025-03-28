# FIMCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. Only compatible for now with:   - &#x60;codestral-2405&#x60;   - &#x60;codestral-latest&#x60; | [default to "codestral-2405"]
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**TopP** | Pointer to **float32** | Nucleus sampling, where the model considers the results of the tokens with &#x60;top_p&#x60; probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. We generally recommend altering this or &#x60;temperature&#x60; but not both. | [optional] [default to 1.0]
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. If set, tokens will be sent as data-only server-side events as they become available, with the stream terminated by a data: [DONE] message. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Stop** | Pointer to [**Stop1**](Stop1.md) |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** |  | [optional] 
**Prompt** | **string** | The text/code to complete. | 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**MinTokens** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFIMCompletionRequest

`func NewFIMCompletionRequest(model string, prompt string, ) *FIMCompletionRequest`

NewFIMCompletionRequest instantiates a new FIMCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIMCompletionRequestWithDefaults

`func NewFIMCompletionRequestWithDefaults() *FIMCompletionRequest`

NewFIMCompletionRequestWithDefaults instantiates a new FIMCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *FIMCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FIMCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FIMCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetTemperature

`func (o *FIMCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *FIMCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *FIMCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *FIMCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *FIMCompletionRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *FIMCompletionRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *FIMCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *FIMCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *FIMCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *FIMCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetMaxTokens

`func (o *FIMCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *FIMCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *FIMCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *FIMCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *FIMCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *FIMCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetStream

`func (o *FIMCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *FIMCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *FIMCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *FIMCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStop

`func (o *FIMCompletionRequest) GetStop() Stop1`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *FIMCompletionRequest) GetStopOk() (*Stop1, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *FIMCompletionRequest) SetStop(v Stop1)`

SetStop sets Stop field to given value.

### HasStop

`func (o *FIMCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetRandomSeed

`func (o *FIMCompletionRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *FIMCompletionRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *FIMCompletionRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *FIMCompletionRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *FIMCompletionRequest) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *FIMCompletionRequest) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetPrompt

`func (o *FIMCompletionRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FIMCompletionRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FIMCompletionRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSuffix

`func (o *FIMCompletionRequest) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *FIMCompletionRequest) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *FIMCompletionRequest) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *FIMCompletionRequest) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *FIMCompletionRequest) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *FIMCompletionRequest) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetMinTokens

`func (o *FIMCompletionRequest) GetMinTokens() int32`

GetMinTokens returns the MinTokens field if non-nil, zero value otherwise.

### GetMinTokensOk

`func (o *FIMCompletionRequest) GetMinTokensOk() (*int32, bool)`

GetMinTokensOk returns a tuple with the MinTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTokens

`func (o *FIMCompletionRequest) SetMinTokens(v int32)`

SetMinTokens sets MinTokens field to given value.

### HasMinTokens

`func (o *FIMCompletionRequest) HasMinTokens() bool`

HasMinTokens returns a boolean if a field has been set.

### SetMinTokensNil

`func (o *FIMCompletionRequest) SetMinTokensNil(b bool)`

 SetMinTokensNil sets the value for MinTokens to be an explicit nil

### UnsetMinTokens
`func (o *FIMCompletionRequest) UnsetMinTokens()`

UnsetMinTokens ensures that no value is present for MinTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


