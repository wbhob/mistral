# AssistantMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**NullableContent**](Content.md) |  | [optional] 
**ToolCalls** | Pointer to [**[]ToolCall**](ToolCall.md) |  | [optional] 
**Prefix** | Pointer to **bool** | Set this to &#x60;true&#x60; when adding an assistant message as prefix to condition the model response. The role of the prefix message is to force the model to start its answer by the content of the message. | [optional] [default to false]
**Role** | Pointer to **string** |  | [optional] [default to "assistant"]

## Methods

### NewAssistantMessage

`func NewAssistantMessage() *AssistantMessage`

NewAssistantMessage instantiates a new AssistantMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantMessageWithDefaults

`func NewAssistantMessageWithDefaults() *AssistantMessage`

NewAssistantMessageWithDefaults instantiates a new AssistantMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AssistantMessage) GetContent() Content`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AssistantMessage) GetContentOk() (*Content, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AssistantMessage) SetContent(v Content)`

SetContent sets Content field to given value.

### HasContent

`func (o *AssistantMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *AssistantMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *AssistantMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCalls

`func (o *AssistantMessage) GetToolCalls() []ToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AssistantMessage) GetToolCallsOk() (*[]ToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AssistantMessage) SetToolCalls(v []ToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AssistantMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### SetToolCallsNil

`func (o *AssistantMessage) SetToolCallsNil(b bool)`

 SetToolCallsNil sets the value for ToolCalls to be an explicit nil

### UnsetToolCalls
`func (o *AssistantMessage) UnsetToolCalls()`

UnsetToolCalls ensures that no value is present for ToolCalls, not even an explicit nil
### GetPrefix

`func (o *AssistantMessage) GetPrefix() bool`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AssistantMessage) GetPrefixOk() (*bool, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AssistantMessage) SetPrefix(v bool)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AssistantMessage) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRole

`func (o *AssistantMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssistantMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssistantMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssistantMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


