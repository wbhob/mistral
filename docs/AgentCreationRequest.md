# AgentCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Tools** | Pointer to [**[]AgentToolsInner**](AgentToolsInner.md) | List of tools which are available to the model during the conversation. | [optional] 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Model** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Handoffs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAgentCreationRequest

`func NewAgentCreationRequest(model string, name string, ) *AgentCreationRequest`

NewAgentCreationRequest instantiates a new AgentCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreationRequestWithDefaults

`func NewAgentCreationRequestWithDefaults() *AgentCreationRequest`

NewAgentCreationRequestWithDefaults instantiates a new AgentCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *AgentCreationRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentCreationRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentCreationRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentCreationRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *AgentCreationRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *AgentCreationRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *AgentCreationRequest) GetTools() []AgentToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentCreationRequest) GetToolsOk() (*[]AgentToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentCreationRequest) SetTools(v []AgentToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentCreationRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *AgentCreationRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *AgentCreationRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *AgentCreationRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *AgentCreationRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetModel

`func (o *AgentCreationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentCreationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentCreationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *AgentCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentCreationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AgentCreationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AgentCreationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *AgentCreationRequest) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *AgentCreationRequest) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *AgentCreationRequest) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *AgentCreationRequest) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *AgentCreationRequest) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *AgentCreationRequest) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


