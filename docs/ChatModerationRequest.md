# ChatModerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | [**Input**](Input.md) |  | 
**TruncateForContextLength** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewChatModerationRequest

`func NewChatModerationRequest(model string, input Input, ) *ChatModerationRequest`

NewChatModerationRequest instantiates a new ChatModerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatModerationRequestWithDefaults

`func NewChatModerationRequestWithDefaults() *ChatModerationRequest`

NewChatModerationRequestWithDefaults instantiates a new ChatModerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatModerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatModerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatModerationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *ChatModerationRequest) GetInput() Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ChatModerationRequest) GetInputOk() (*Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ChatModerationRequest) SetInput(v Input)`

SetInput sets Input field to given value.


### GetTruncateForContextLength

`func (o *ChatModerationRequest) GetTruncateForContextLength() bool`

GetTruncateForContextLength returns the TruncateForContextLength field if non-nil, zero value otherwise.

### GetTruncateForContextLengthOk

`func (o *ChatModerationRequest) GetTruncateForContextLengthOk() (*bool, bool)`

GetTruncateForContextLengthOk returns a tuple with the TruncateForContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateForContextLength

`func (o *ChatModerationRequest) SetTruncateForContextLength(v bool)`

SetTruncateForContextLength sets TruncateForContextLength field to given value.

### HasTruncateForContextLength

`func (o *ChatModerationRequest) HasTruncateForContextLength() bool`

HasTruncateForContextLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


