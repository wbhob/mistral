# ConversationRestartRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**ConversationInputs**](ConversationInputs.md) |  | 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Store** | Pointer to **bool** | Whether to store the results into our servers or not. | [optional] [default to true]
**HandoffExecution** | Pointer to **string** |  | [optional] [default to "server"]
**FromEntryId** | **string** |  | 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 

## Methods

### NewConversationRestartRequestBase

`func NewConversationRestartRequestBase(inputs ConversationInputs, fromEntryId string, ) *ConversationRestartRequestBase`

NewConversationRestartRequestBase instantiates a new ConversationRestartRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationRestartRequestBaseWithDefaults

`func NewConversationRestartRequestBaseWithDefaults() *ConversationRestartRequestBase`

NewConversationRestartRequestBaseWithDefaults instantiates a new ConversationRestartRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationRestartRequestBase) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationRestartRequestBase) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationRestartRequestBase) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.


### GetStream

`func (o *ConversationRestartRequestBase) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationRestartRequestBase) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationRestartRequestBase) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationRestartRequestBase) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *ConversationRestartRequestBase) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationRestartRequestBase) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationRestartRequestBase) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationRestartRequestBase) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *ConversationRestartRequestBase) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationRestartRequestBase) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationRestartRequestBase) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationRestartRequestBase) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetFromEntryId

`func (o *ConversationRestartRequestBase) GetFromEntryId() string`

GetFromEntryId returns the FromEntryId field if non-nil, zero value otherwise.

### GetFromEntryIdOk

`func (o *ConversationRestartRequestBase) GetFromEntryIdOk() (*string, bool)`

GetFromEntryIdOk returns a tuple with the FromEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntryId

`func (o *ConversationRestartRequestBase) SetFromEntryId(v string)`

SetFromEntryId sets FromEntryId field to given value.


### GetCompletionArgs

`func (o *ConversationRestartRequestBase) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationRestartRequestBase) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationRestartRequestBase) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationRestartRequestBase) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


