# ResponseV1ConversationsGet

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

### NewResponseV1ConversationsGet

`func NewResponseV1ConversationsGet(id string, createdAt time.Time, updatedAt time.Time, model string, agentId string, ) *ResponseV1ConversationsGet`

NewResponseV1ConversationsGet instantiates a new ResponseV1ConversationsGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseV1ConversationsGetWithDefaults

`func NewResponseV1ConversationsGetWithDefaults() *ResponseV1ConversationsGet`

NewResponseV1ConversationsGetWithDefaults instantiates a new ResponseV1ConversationsGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *ResponseV1ConversationsGet) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ResponseV1ConversationsGet) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ResponseV1ConversationsGet) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ResponseV1ConversationsGet) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetTools

`func (o *ResponseV1ConversationsGet) GetTools() []AgentToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ResponseV1ConversationsGet) GetToolsOk() (*[]AgentToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ResponseV1ConversationsGet) SetTools(v []AgentToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ResponseV1ConversationsGet) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *ResponseV1ConversationsGet) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ResponseV1ConversationsGet) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ResponseV1ConversationsGet) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ResponseV1ConversationsGet) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetName

`func (o *ResponseV1ConversationsGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseV1ConversationsGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseV1ConversationsGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseV1ConversationsGet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseV1ConversationsGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseV1ConversationsGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseV1ConversationsGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseV1ConversationsGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObject

`func (o *ResponseV1ConversationsGet) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseV1ConversationsGet) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseV1ConversationsGet) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseV1ConversationsGet) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetId

`func (o *ResponseV1ConversationsGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseV1ConversationsGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseV1ConversationsGet) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ResponseV1ConversationsGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseV1ConversationsGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseV1ConversationsGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResponseV1ConversationsGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponseV1ConversationsGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponseV1ConversationsGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetModel

`func (o *ResponseV1ConversationsGet) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponseV1ConversationsGet) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponseV1ConversationsGet) SetModel(v string)`

SetModel sets Model field to given value.


### GetAgentId

`func (o *ResponseV1ConversationsGet) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ResponseV1ConversationsGet) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ResponseV1ConversationsGet) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


