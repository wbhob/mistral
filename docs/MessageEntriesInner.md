# MessageEntriesInner

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
**Prefix** | Pointer to **bool** |  | [optional] [default to false]
**AgentId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageEntriesInner

`func NewMessageEntriesInner(role string, content Content1, ) *MessageEntriesInner`

NewMessageEntriesInner instantiates a new MessageEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageEntriesInnerWithDefaults

`func NewMessageEntriesInnerWithDefaults() *MessageEntriesInner`

NewMessageEntriesInnerWithDefaults instantiates a new MessageEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MessageEntriesInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MessageEntriesInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MessageEntriesInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *MessageEntriesInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *MessageEntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageEntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageEntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageEntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageEntriesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageEntriesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageEntriesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageEntriesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MessageEntriesInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MessageEntriesInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MessageEntriesInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MessageEntriesInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetId

`func (o *MessageEntriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageEntriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageEntriesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageEntriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *MessageEntriesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageEntriesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageEntriesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *MessageEntriesInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageEntriesInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageEntriesInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetPrefix

`func (o *MessageEntriesInner) GetPrefix() bool`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MessageEntriesInner) GetPrefixOk() (*bool, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MessageEntriesInner) SetPrefix(v bool)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *MessageEntriesInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAgentId

`func (o *MessageEntriesInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *MessageEntriesInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *MessageEntriesInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *MessageEntriesInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetModel

`func (o *MessageEntriesInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MessageEntriesInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MessageEntriesInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MessageEntriesInner) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


