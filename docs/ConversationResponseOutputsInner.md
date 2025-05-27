# ConversationResponseOutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "message.output"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "assistant"]
**Content** | [**Content1**](Content1.md) |  | 
**Name** | **string** |  | 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolCallId** | **string** |  | 
**Arguments** | [**FunctionCallEntryArguments**](FunctionCallEntryArguments.md) |  | 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewConversationResponseOutputsInner

`func NewConversationResponseOutputsInner(content Content1, name string, toolCallId string, arguments FunctionCallEntryArguments, previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string, ) *ConversationResponseOutputsInner`

NewConversationResponseOutputsInner instantiates a new ConversationResponseOutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationResponseOutputsInnerWithDefaults

`func NewConversationResponseOutputsInnerWithDefaults() *ConversationResponseOutputsInner`

NewConversationResponseOutputsInnerWithDefaults instantiates a new ConversationResponseOutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConversationResponseOutputsInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConversationResponseOutputsInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConversationResponseOutputsInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ConversationResponseOutputsInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *ConversationResponseOutputsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConversationResponseOutputsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConversationResponseOutputsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConversationResponseOutputsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConversationResponseOutputsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationResponseOutputsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationResponseOutputsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConversationResponseOutputsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ConversationResponseOutputsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ConversationResponseOutputsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ConversationResponseOutputsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ConversationResponseOutputsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetId

`func (o *ConversationResponseOutputsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationResponseOutputsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationResponseOutputsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConversationResponseOutputsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgentId

`func (o *ConversationResponseOutputsInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConversationResponseOutputsInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConversationResponseOutputsInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ConversationResponseOutputsInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetModel

`func (o *ConversationResponseOutputsInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConversationResponseOutputsInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConversationResponseOutputsInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConversationResponseOutputsInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRole

`func (o *ConversationResponseOutputsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConversationResponseOutputsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConversationResponseOutputsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConversationResponseOutputsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *ConversationResponseOutputsInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConversationResponseOutputsInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConversationResponseOutputsInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetName

`func (o *ConversationResponseOutputsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversationResponseOutputsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversationResponseOutputsInner) SetName(v string)`

SetName sets Name field to given value.


### GetInfo

`func (o *ConversationResponseOutputsInner) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConversationResponseOutputsInner) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConversationResponseOutputsInner) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConversationResponseOutputsInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetToolCallId

`func (o *ConversationResponseOutputsInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ConversationResponseOutputsInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ConversationResponseOutputsInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetArguments

`func (o *ConversationResponseOutputsInner) GetArguments() FunctionCallEntryArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ConversationResponseOutputsInner) GetArgumentsOk() (*FunctionCallEntryArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ConversationResponseOutputsInner) SetArguments(v FunctionCallEntryArguments)`

SetArguments sets Arguments field to given value.


### GetPreviousAgentId

`func (o *ConversationResponseOutputsInner) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *ConversationResponseOutputsInner) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *ConversationResponseOutputsInner) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *ConversationResponseOutputsInner) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *ConversationResponseOutputsInner) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *ConversationResponseOutputsInner) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.


### GetNextAgentId

`func (o *ConversationResponseOutputsInner) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *ConversationResponseOutputsInner) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *ConversationResponseOutputsInner) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *ConversationResponseOutputsInner) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *ConversationResponseOutputsInner) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *ConversationResponseOutputsInner) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


