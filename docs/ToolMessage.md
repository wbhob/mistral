# ToolMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**NullableContent**](Content.md) |  | 
**ToolCallId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "tool"]

## Methods

### NewToolMessage

`func NewToolMessage(content NullableContent, ) *ToolMessage`

NewToolMessage instantiates a new ToolMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMessageWithDefaults

`func NewToolMessageWithDefaults() *ToolMessage`

NewToolMessageWithDefaults instantiates a new ToolMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ToolMessage) GetContent() Content`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ToolMessage) GetContentOk() (*Content, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ToolMessage) SetContent(v Content)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *ToolMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ToolMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCallId

`func (o *ToolMessage) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ToolMessage) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ToolMessage) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *ToolMessage) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### SetToolCallIdNil

`func (o *ToolMessage) SetToolCallIdNil(b bool)`

 SetToolCallIdNil sets the value for ToolCallId to be an explicit nil

### UnsetToolCallId
`func (o *ToolMessage) UnsetToolCallId()`

UnsetToolCallId ensures that no value is present for ToolCallId, not even an explicit nil
### GetName

`func (o *ToolMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ToolMessage) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ToolMessage) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *ToolMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ToolMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ToolMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ToolMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


