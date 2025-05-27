# ResponseDoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "conversation.response.done"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Usage** | [**ConversationUsageInfo**](ConversationUsageInfo.md) |  | 

## Methods

### NewResponseDoneEvent

`func NewResponseDoneEvent(usage ConversationUsageInfo, ) *ResponseDoneEvent`

NewResponseDoneEvent instantiates a new ResponseDoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDoneEventWithDefaults

`func NewResponseDoneEventWithDefaults() *ResponseDoneEvent`

NewResponseDoneEventWithDefaults instantiates a new ResponseDoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseDoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDoneEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDoneEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseDoneEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseDoneEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseDoneEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponseDoneEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *ResponseDoneEvent) GetUsage() ConversationUsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ResponseDoneEvent) GetUsageOk() (*ConversationUsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ResponseDoneEvent) SetUsage(v ConversationUsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


