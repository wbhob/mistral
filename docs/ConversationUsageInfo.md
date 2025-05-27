# ConversationUsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CompletionTokens** | Pointer to **int32** |  | [optional] [default to 0]
**TotalTokens** | Pointer to **int32** |  | [optional] [default to 0]
**ConnectorTokens** | Pointer to [**ConnectorTokens**](ConnectorTokens.md) |  | [optional] 
**Connectors** | Pointer to [**Connectors**](Connectors.md) |  | [optional] 

## Methods

### NewConversationUsageInfo

`func NewConversationUsageInfo() *ConversationUsageInfo`

NewConversationUsageInfo instantiates a new ConversationUsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationUsageInfoWithDefaults

`func NewConversationUsageInfoWithDefaults() *ConversationUsageInfo`

NewConversationUsageInfoWithDefaults instantiates a new ConversationUsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptTokens

`func (o *ConversationUsageInfo) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ConversationUsageInfo) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ConversationUsageInfo) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *ConversationUsageInfo) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *ConversationUsageInfo) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ConversationUsageInfo) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ConversationUsageInfo) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *ConversationUsageInfo) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ConversationUsageInfo) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ConversationUsageInfo) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ConversationUsageInfo) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ConversationUsageInfo) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetConnectorTokens

`func (o *ConversationUsageInfo) GetConnectorTokens() ConnectorTokens`

GetConnectorTokens returns the ConnectorTokens field if non-nil, zero value otherwise.

### GetConnectorTokensOk

`func (o *ConversationUsageInfo) GetConnectorTokensOk() (*ConnectorTokens, bool)`

GetConnectorTokensOk returns a tuple with the ConnectorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTokens

`func (o *ConversationUsageInfo) SetConnectorTokens(v ConnectorTokens)`

SetConnectorTokens sets ConnectorTokens field to given value.

### HasConnectorTokens

`func (o *ConversationUsageInfo) HasConnectorTokens() bool`

HasConnectorTokens returns a boolean if a field has been set.

### GetConnectors

`func (o *ConversationUsageInfo) GetConnectors() Connectors`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ConversationUsageInfo) GetConnectorsOk() (*Connectors, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ConversationUsageInfo) SetConnectors(v Connectors)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ConversationUsageInfo) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


