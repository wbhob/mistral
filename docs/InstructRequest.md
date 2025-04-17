# InstructRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]MessagesInner**](MessagesInner.md) |  | 

## Methods

### NewInstructRequest

`func NewInstructRequest(messages []MessagesInner, ) *InstructRequest`

NewInstructRequest instantiates a new InstructRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructRequestWithDefaults

`func NewInstructRequestWithDefaults() *InstructRequest`

NewInstructRequestWithDefaults instantiates a new InstructRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *InstructRequest) GetMessages() []MessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *InstructRequest) GetMessagesOk() (*[]MessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *InstructRequest) SetMessages(v []MessagesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


