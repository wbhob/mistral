# MessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**NullableContent3**](Content3.md) |  | 
**Role** | Pointer to **string** |  | [optional] [default to "tool"]
**ToolCalls** | Pointer to [**[]ToolCall**](ToolCall.md) |  | [optional] 
**Prefix** | Pointer to **bool** | Set this to &#x60;true&#x60; when adding an assistant message as prefix to condition the model response. The role of the prefix message is to force the model to start its answer by the content of the message. | [optional] [default to false]
**ToolCallId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewMessagesInner

`func NewMessagesInner(content NullableContent3, ) *MessagesInner`

NewMessagesInner instantiates a new MessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesInnerWithDefaults

`func NewMessagesInnerWithDefaults() *MessagesInner`

NewMessagesInnerWithDefaults instantiates a new MessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MessagesInner) GetContent() Content3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessagesInner) GetContentOk() (*Content3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessagesInner) SetContent(v Content3)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *MessagesInner) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MessagesInner) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRole

`func (o *MessagesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessagesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessagesInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MessagesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToolCalls

`func (o *MessagesInner) GetToolCalls() []ToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *MessagesInner) GetToolCallsOk() (*[]ToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *MessagesInner) SetToolCalls(v []ToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *MessagesInner) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetPrefix

`func (o *MessagesInner) GetPrefix() bool`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MessagesInner) GetPrefixOk() (*bool, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MessagesInner) SetPrefix(v bool)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *MessagesInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetToolCallId

`func (o *MessagesInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *MessagesInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *MessagesInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *MessagesInner) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### GetName

`func (o *MessagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessagesInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


