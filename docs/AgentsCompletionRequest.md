# AgentsCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. If set, tokens will be sent as data-only server-side events as they become available, with the stream terminated by a data: [DONE] message. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Stop** | Pointer to [**Stop1**](Stop1.md) |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** |  | [optional] 
**Messages** | [**[]MessagesInner**](MessagesInner.md) | The prompt(s) to generate completions for, encoded as a list of dict with role and content. | 
**ResponseFormat** | Pointer to [**ResponseFormat**](ResponseFormat.md) |  | [optional] 
**Tools** | Pointer to [**[]Tool**](Tool.md) |  | [optional] 
**ToolChoice** | Pointer to [**ToolChoice**](ToolChoice.md) |  | [optional] [default to auto]
**PresencePenalty** | Pointer to **float32** | presence_penalty determines how much the model penalizes the repetition of words or phrases. A higher presence penalty encourages the model to use a wider variety of words and phrases, making the output more diverse and creative. | [optional] [default to 0.0]
**FrequencyPenalty** | Pointer to **float32** | frequency_penalty penalizes the repetition of words based on their frequency in the generated text. A higher frequency penalty discourages the model from repeating words that have already appeared frequently in the output, promoting diversity and reducing repetition. | [optional] [default to 0.0]
**N** | Pointer to **NullableInt32** |  | [optional] 
**Prediction** | Pointer to [**Prediction**](Prediction.md) | Enable users to specify expected results, optimizing response times by leveraging known or predictable content. This approach is especially effective for updating text documents or code files with minimal changes, reducing latency while maintaining high-quality results. | [optional] 
**ParallelToolCalls** | Pointer to **bool** |  | [optional] [default to true]
**PromptMode** | Pointer to [**NullableMistralPromptMode**](MistralPromptMode.md) |  | [optional] 
**AgentId** | **string** | The ID of the agent to use for this completion. | 

## Methods

### NewAgentsCompletionRequest

`func NewAgentsCompletionRequest(messages []MessagesInner, agentId string, ) *AgentsCompletionRequest`

NewAgentsCompletionRequest instantiates a new AgentsCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsCompletionRequestWithDefaults

`func NewAgentsCompletionRequestWithDefaults() *AgentsCompletionRequest`

NewAgentsCompletionRequestWithDefaults instantiates a new AgentsCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokens

`func (o *AgentsCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AgentsCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AgentsCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AgentsCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *AgentsCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *AgentsCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetStream

`func (o *AgentsCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AgentsCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AgentsCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AgentsCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStop

`func (o *AgentsCompletionRequest) GetStop() Stop1`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *AgentsCompletionRequest) GetStopOk() (*Stop1, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *AgentsCompletionRequest) SetStop(v Stop1)`

SetStop sets Stop field to given value.

### HasStop

`func (o *AgentsCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetRandomSeed

`func (o *AgentsCompletionRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *AgentsCompletionRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *AgentsCompletionRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *AgentsCompletionRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *AgentsCompletionRequest) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *AgentsCompletionRequest) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetMessages

`func (o *AgentsCompletionRequest) GetMessages() []MessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AgentsCompletionRequest) GetMessagesOk() (*[]MessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AgentsCompletionRequest) SetMessages(v []MessagesInner)`

SetMessages sets Messages field to given value.


### GetResponseFormat

`func (o *AgentsCompletionRequest) GetResponseFormat() ResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AgentsCompletionRequest) GetResponseFormatOk() (*ResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AgentsCompletionRequest) SetResponseFormat(v ResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AgentsCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetTools

`func (o *AgentsCompletionRequest) GetTools() []Tool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentsCompletionRequest) GetToolsOk() (*[]Tool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentsCompletionRequest) SetTools(v []Tool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentsCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *AgentsCompletionRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *AgentsCompletionRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolChoice

`func (o *AgentsCompletionRequest) GetToolChoice() ToolChoice`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AgentsCompletionRequest) GetToolChoiceOk() (*ToolChoice, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AgentsCompletionRequest) SetToolChoice(v ToolChoice)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AgentsCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *AgentsCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *AgentsCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *AgentsCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *AgentsCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *AgentsCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *AgentsCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *AgentsCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *AgentsCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetN

`func (o *AgentsCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentsCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentsCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *AgentsCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *AgentsCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *AgentsCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetPrediction

`func (o *AgentsCompletionRequest) GetPrediction() Prediction`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *AgentsCompletionRequest) GetPredictionOk() (*Prediction, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *AgentsCompletionRequest) SetPrediction(v Prediction)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *AgentsCompletionRequest) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### GetParallelToolCalls

`func (o *AgentsCompletionRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *AgentsCompletionRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *AgentsCompletionRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *AgentsCompletionRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetPromptMode

`func (o *AgentsCompletionRequest) GetPromptMode() MistralPromptMode`

GetPromptMode returns the PromptMode field if non-nil, zero value otherwise.

### GetPromptModeOk

`func (o *AgentsCompletionRequest) GetPromptModeOk() (*MistralPromptMode, bool)`

GetPromptModeOk returns a tuple with the PromptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptMode

`func (o *AgentsCompletionRequest) SetPromptMode(v MistralPromptMode)`

SetPromptMode sets PromptMode field to given value.

### HasPromptMode

`func (o *AgentsCompletionRequest) HasPromptMode() bool`

HasPromptMode returns a boolean if a field has been set.

### SetPromptModeNil

`func (o *AgentsCompletionRequest) SetPromptModeNil(b bool)`

 SetPromptModeNil sets the value for PromptMode to be an explicit nil

### UnsetPromptMode
`func (o *AgentsCompletionRequest) UnsetPromptMode()`

UnsetPromptMode ensures that no value is present for PromptMode, not even an explicit nil
### GetAgentId

`func (o *AgentsCompletionRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentsCompletionRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentsCompletionRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


