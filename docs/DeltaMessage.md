# DeltaMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to [**NullableDeltaMessageContent**](DeltaMessageContent.md) |  | [optional] 
**ToolCalls** | Pointer to [**[]ToolCall**](ToolCall.md) |  | [optional] 

## Methods

### NewDeltaMessage

`func NewDeltaMessage() *DeltaMessage`

NewDeltaMessage instantiates a new DeltaMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeltaMessageWithDefaults

`func NewDeltaMessageWithDefaults() *DeltaMessage`

NewDeltaMessageWithDefaults instantiates a new DeltaMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *DeltaMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DeltaMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DeltaMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DeltaMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *DeltaMessage) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *DeltaMessage) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetContent

`func (o *DeltaMessage) GetContent() DeltaMessageContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DeltaMessage) GetContentOk() (*DeltaMessageContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DeltaMessage) SetContent(v DeltaMessageContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *DeltaMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DeltaMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DeltaMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCalls

`func (o *DeltaMessage) GetToolCalls() []ToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *DeltaMessage) GetToolCallsOk() (*[]ToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *DeltaMessage) SetToolCalls(v []ToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *DeltaMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### SetToolCallsNil

`func (o *DeltaMessage) SetToolCallsNil(b bool)`

 SetToolCallsNil sets the value for ToolCalls to be an explicit nil

### UnsetToolCalls
`func (o *DeltaMessage) UnsetToolCalls()`

UnsetToolCalls ensures that no value is present for ToolCalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


