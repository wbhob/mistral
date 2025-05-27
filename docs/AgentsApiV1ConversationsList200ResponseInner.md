# AgentsApiV1ConversationsList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]AgentToolsInner**](AgentToolsInner.md) | List of tools which are available to the model during the conversation. | [optional] 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "conversation"]
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Model** | **string** |  | 
**AgentId** | **string** |  | 

## Methods

### NewAgentsApiV1ConversationsList200ResponseInner

`func NewAgentsApiV1ConversationsList200ResponseInner(id string, createdAt time.Time, updatedAt time.Time, model string, agentId string, ) *AgentsApiV1ConversationsList200ResponseInner`

NewAgentsApiV1ConversationsList200ResponseInner instantiates a new AgentsApiV1ConversationsList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsApiV1ConversationsList200ResponseInnerWithDefaults

`func NewAgentsApiV1ConversationsList200ResponseInnerWithDefaults() *AgentsApiV1ConversationsList200ResponseInner`

NewAgentsApiV1ConversationsList200ResponseInnerWithDefaults instantiates a new AgentsApiV1ConversationsList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetTools

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetTools() []AgentToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetToolsOk() (*[]AgentToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetTools(v []AgentToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetName

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObject

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AgentsApiV1ConversationsList200ResponseInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetId

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetModel

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.


### GetAgentId

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentsApiV1ConversationsList200ResponseInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentsApiV1ConversationsList200ResponseInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


