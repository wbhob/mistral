# ResponseStartedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "conversation.response.started"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ConversationId** | **string** |  | 

## Methods

### NewResponseStartedEvent

`func NewResponseStartedEvent(conversationId string, ) *ResponseStartedEvent`

NewResponseStartedEvent instantiates a new ResponseStartedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStartedEventWithDefaults

`func NewResponseStartedEventWithDefaults() *ResponseStartedEvent`

NewResponseStartedEventWithDefaults instantiates a new ResponseStartedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseStartedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseStartedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseStartedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseStartedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseStartedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseStartedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseStartedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponseStartedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetConversationId

`func (o *ResponseStartedEvent) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ResponseStartedEvent) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ResponseStartedEvent) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


