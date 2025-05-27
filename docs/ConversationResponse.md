# ConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "conversation.response"]
**ConversationId** | **string** |  | 
**Outputs** | [**[]ConversationResponseOutputsInner**](ConversationResponseOutputsInner.md) |  | 
**Usage** | [**ConversationUsageInfo**](ConversationUsageInfo.md) |  | 

## Methods

### NewConversationResponse

`func NewConversationResponse(conversationId string, outputs []ConversationResponseOutputsInner, usage ConversationUsageInfo, ) *ConversationResponse`

NewConversationResponse instantiates a new ConversationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationResponseWithDefaults

`func NewConversationResponseWithDefaults() *ConversationResponse`

NewConversationResponseWithDefaults instantiates a new ConversationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConversationResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConversationResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConversationResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ConversationResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetConversationId

`func (o *ConversationResponse) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationResponse) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationResponse) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetOutputs

`func (o *ConversationResponse) GetOutputs() []ConversationResponseOutputsInner`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ConversationResponse) GetOutputsOk() (*[]ConversationResponseOutputsInner, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ConversationResponse) SetOutputs(v []ConversationResponseOutputsInner)`

SetOutputs sets Outputs field to given value.


### GetUsage

`func (o *ConversationResponse) GetUsage() ConversationUsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ConversationResponse) GetUsageOk() (*ConversationUsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ConversationResponse) SetUsage(v ConversationUsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


