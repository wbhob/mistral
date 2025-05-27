# ConversationEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**SSETypes**](SSETypes.md) |  | 
**Data** | [**Data**](Data.md) |  | 

## Methods

### NewConversationEvents

`func NewConversationEvents(event SSETypes, data Data, ) *ConversationEvents`

NewConversationEvents instantiates a new ConversationEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationEventsWithDefaults

`func NewConversationEventsWithDefaults() *ConversationEvents`

NewConversationEventsWithDefaults instantiates a new ConversationEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ConversationEvents) GetEvent() SSETypes`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ConversationEvents) GetEventOk() (*SSETypes, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ConversationEvents) SetEvent(v SSETypes)`

SetEvent sets Event field to given value.


### GetData

`func (o *ConversationEvents) GetData() Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConversationEvents) GetDataOk() (*Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConversationEvents) SetData(v Data)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


