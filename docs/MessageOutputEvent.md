# MessageOutputEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "message.output.delta"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**ContentIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Model** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "assistant"]
**Content** | [**Content2**](Content2.md) |  | 

## Methods

### NewMessageOutputEvent

`func NewMessageOutputEvent(id string, content Content2, ) *MessageOutputEvent`

NewMessageOutputEvent instantiates a new MessageOutputEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOutputEventWithDefaults

`func NewMessageOutputEventWithDefaults() *MessageOutputEvent`

NewMessageOutputEventWithDefaults instantiates a new MessageOutputEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOutputEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOutputEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOutputEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageOutputEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageOutputEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageOutputEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageOutputEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageOutputEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *MessageOutputEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *MessageOutputEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *MessageOutputEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *MessageOutputEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *MessageOutputEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageOutputEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageOutputEvent) SetId(v string)`

SetId sets Id field to given value.


### GetContentIndex

`func (o *MessageOutputEvent) GetContentIndex() int32`

GetContentIndex returns the ContentIndex field if non-nil, zero value otherwise.

### GetContentIndexOk

`func (o *MessageOutputEvent) GetContentIndexOk() (*int32, bool)`

GetContentIndexOk returns a tuple with the ContentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentIndex

`func (o *MessageOutputEvent) SetContentIndex(v int32)`

SetContentIndex sets ContentIndex field to given value.

### HasContentIndex

`func (o *MessageOutputEvent) HasContentIndex() bool`

HasContentIndex returns a boolean if a field has been set.

### GetModel

`func (o *MessageOutputEvent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MessageOutputEvent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MessageOutputEvent) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MessageOutputEvent) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MessageOutputEvent) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MessageOutputEvent) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetAgentId

`func (o *MessageOutputEvent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *MessageOutputEvent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *MessageOutputEvent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *MessageOutputEvent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *MessageOutputEvent) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *MessageOutputEvent) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetRole

`func (o *MessageOutputEvent) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageOutputEvent) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageOutputEvent) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MessageOutputEvent) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *MessageOutputEvent) GetContent() Content2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageOutputEvent) GetContentOk() (*Content2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageOutputEvent) SetContent(v Content2)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


