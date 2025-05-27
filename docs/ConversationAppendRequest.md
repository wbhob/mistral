# ConversationAppendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**ConversationInputs**](ConversationInputs.md) |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**Store** | Pointer to **bool** | Whether to store the results into our servers or not. | [optional] [default to true]
**HandoffExecution** | Pointer to **string** |  | [optional] [default to "server"]
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 

## Methods

### NewConversationAppendRequest

`func NewConversationAppendRequest(inputs ConversationInputs, ) *ConversationAppendRequest`

NewConversationAppendRequest instantiates a new ConversationAppendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationAppendRequestWithDefaults

`func NewConversationAppendRequestWithDefaults() *ConversationAppendRequest`

NewConversationAppendRequestWithDefaults instantiates a new ConversationAppendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationAppendRequest) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationAppendRequest) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationAppendRequest) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.


### GetStream

`func (o *ConversationAppendRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationAppendRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationAppendRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationAppendRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *ConversationAppendRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationAppendRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationAppendRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationAppendRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *ConversationAppendRequest) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationAppendRequest) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationAppendRequest) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationAppendRequest) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *ConversationAppendRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationAppendRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationAppendRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationAppendRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


