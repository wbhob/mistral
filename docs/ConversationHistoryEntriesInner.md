# ConversationHistoryEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "message.input"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | [default to "assistant"]
**Content** | [**Content1**](Content1.md) |  | 
**AgentId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ToolCallId** | **string** |  | 
**Result** | **string** |  | 
**Name** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Arguments** | [**FunctionCallEntryArguments**](FunctionCallEntryArguments.md) |  | 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewConversationHistoryEntriesInner

`func NewConversationHistoryEntriesInner(role string, content Content1, toolCallId string, result string, name BuiltInConnectors, arguments FunctionCallEntryArguments, previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string, ) *ConversationHistoryEntriesInner`

NewConversationHistoryEntriesInner instantiates a new ConversationHistoryEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationHistoryEntriesInnerWithDefaults

`func NewConversationHistoryEntriesInnerWithDefaults() *ConversationHistoryEntriesInner`

NewConversationHistoryEntriesInnerWithDefaults instantiates a new ConversationHistoryEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConversationHistoryEntriesInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConversationHistoryEntriesInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConversationHistoryEntriesInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ConversationHistoryEntriesInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *ConversationHistoryEntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConversationHistoryEntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConversationHistoryEntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConversationHistoryEntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConversationHistoryEntriesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationHistoryEntriesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationHistoryEntriesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConversationHistoryEntriesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ConversationHistoryEntriesInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ConversationHistoryEntriesInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ConversationHistoryEntriesInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ConversationHistoryEntriesInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetId

`func (o *ConversationHistoryEntriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationHistoryEntriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationHistoryEntriesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConversationHistoryEntriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *ConversationHistoryEntriesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConversationHistoryEntriesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConversationHistoryEntriesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *ConversationHistoryEntriesInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConversationHistoryEntriesInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConversationHistoryEntriesInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetAgentId

`func (o *ConversationHistoryEntriesInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConversationHistoryEntriesInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConversationHistoryEntriesInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ConversationHistoryEntriesInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetModel

`func (o *ConversationHistoryEntriesInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConversationHistoryEntriesInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConversationHistoryEntriesInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConversationHistoryEntriesInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetToolCallId

`func (o *ConversationHistoryEntriesInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ConversationHistoryEntriesInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ConversationHistoryEntriesInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetResult

`func (o *ConversationHistoryEntriesInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConversationHistoryEntriesInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConversationHistoryEntriesInner) SetResult(v string)`

SetResult sets Result field to given value.


### GetName

`func (o *ConversationHistoryEntriesInner) GetName() BuiltInConnectors`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversationHistoryEntriesInner) GetNameOk() (*BuiltInConnectors, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversationHistoryEntriesInner) SetName(v BuiltInConnectors)`

SetName sets Name field to given value.


### GetArguments

`func (o *ConversationHistoryEntriesInner) GetArguments() FunctionCallEntryArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ConversationHistoryEntriesInner) GetArgumentsOk() (*FunctionCallEntryArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ConversationHistoryEntriesInner) SetArguments(v FunctionCallEntryArguments)`

SetArguments sets Arguments field to given value.


### GetInfo

`func (o *ConversationHistoryEntriesInner) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConversationHistoryEntriesInner) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConversationHistoryEntriesInner) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConversationHistoryEntriesInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetPreviousAgentId

`func (o *ConversationHistoryEntriesInner) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *ConversationHistoryEntriesInner) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *ConversationHistoryEntriesInner) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *ConversationHistoryEntriesInner) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *ConversationHistoryEntriesInner) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *ConversationHistoryEntriesInner) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.


### GetNextAgentId

`func (o *ConversationHistoryEntriesInner) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *ConversationHistoryEntriesInner) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *ConversationHistoryEntriesInner) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *ConversationHistoryEntriesInner) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *ConversationHistoryEntriesInner) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *ConversationHistoryEntriesInner) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


