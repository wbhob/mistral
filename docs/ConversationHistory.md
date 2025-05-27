# ConversationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "conversation.history"]
**ConversationId** | **string** |  | 
**Entries** | [**[]ConversationHistoryEntriesInner**](ConversationHistoryEntriesInner.md) |  | 

## Methods

### NewConversationHistory

`func NewConversationHistory(conversationId string, entries []ConversationHistoryEntriesInner, ) *ConversationHistory`

NewConversationHistory instantiates a new ConversationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationHistoryWithDefaults

`func NewConversationHistoryWithDefaults() *ConversationHistory`

NewConversationHistoryWithDefaults instantiates a new ConversationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConversationHistory) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConversationHistory) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConversationHistory) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ConversationHistory) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetConversationId

`func (o *ConversationHistory) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationHistory) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationHistory) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetEntries

`func (o *ConversationHistory) GetEntries() []ConversationHistoryEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ConversationHistory) GetEntriesOk() (*[]ConversationHistoryEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ConversationHistory) SetEntries(v []ConversationHistoryEntriesInner)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


