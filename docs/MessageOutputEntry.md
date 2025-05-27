# MessageOutputEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "message.output"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "assistant"]
**Content** | [**Content1**](Content1.md) |  | 

## Methods

### NewMessageOutputEntry

`func NewMessageOutputEntry(content Content1, ) *MessageOutputEntry`

NewMessageOutputEntry instantiates a new MessageOutputEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOutputEntryWithDefaults

`func NewMessageOutputEntryWithDefaults() *MessageOutputEntry`

NewMessageOutputEntryWithDefaults instantiates a new MessageOutputEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MessageOutputEntry) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MessageOutputEntry) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MessageOutputEntry) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *MessageOutputEntry) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *MessageOutputEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOutputEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOutputEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageOutputEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageOutputEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageOutputEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageOutputEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageOutputEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MessageOutputEntry) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MessageOutputEntry) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MessageOutputEntry) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MessageOutputEntry) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *MessageOutputEntry) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *MessageOutputEntry) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetId

`func (o *MessageOutputEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageOutputEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageOutputEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageOutputEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgentId

`func (o *MessageOutputEntry) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *MessageOutputEntry) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *MessageOutputEntry) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *MessageOutputEntry) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *MessageOutputEntry) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *MessageOutputEntry) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetModel

`func (o *MessageOutputEntry) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MessageOutputEntry) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MessageOutputEntry) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MessageOutputEntry) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MessageOutputEntry) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MessageOutputEntry) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetRole

`func (o *MessageOutputEntry) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageOutputEntry) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageOutputEntry) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MessageOutputEntry) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *MessageOutputEntry) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageOutputEntry) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageOutputEntry) SetContent(v Content1)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


