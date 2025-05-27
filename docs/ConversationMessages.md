# ConversationMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "conversation.messages"]
**ConversationId** | **string** |  | 
**Messages** | [**[]MessageEntriesInner**](MessageEntriesInner.md) |  | 

## Methods

### NewConversationMessages

`func NewConversationMessages(conversationId string, messages []MessageEntriesInner, ) *ConversationMessages`

NewConversationMessages instantiates a new ConversationMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMessagesWithDefaults

`func NewConversationMessagesWithDefaults() *ConversationMessages`

NewConversationMessagesWithDefaults instantiates a new ConversationMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConversationMessages) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConversationMessages) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConversationMessages) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ConversationMessages) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetConversationId

`func (o *ConversationMessages) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationMessages) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationMessages) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetMessages

`func (o *ConversationMessages) GetMessages() []MessageEntriesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ConversationMessages) GetMessagesOk() (*[]MessageEntriesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ConversationMessages) SetMessages(v []MessageEntriesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


