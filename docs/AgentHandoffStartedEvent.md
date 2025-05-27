# AgentHandoffStartedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "agent.handoff.started"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 

## Methods

### NewAgentHandoffStartedEvent

`func NewAgentHandoffStartedEvent(id string, previousAgentId string, previousAgentName string, ) *AgentHandoffStartedEvent`

NewAgentHandoffStartedEvent instantiates a new AgentHandoffStartedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentHandoffStartedEventWithDefaults

`func NewAgentHandoffStartedEventWithDefaults() *AgentHandoffStartedEvent`

NewAgentHandoffStartedEventWithDefaults instantiates a new AgentHandoffStartedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgentHandoffStartedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentHandoffStartedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentHandoffStartedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentHandoffStartedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentHandoffStartedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentHandoffStartedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentHandoffStartedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentHandoffStartedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *AgentHandoffStartedEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *AgentHandoffStartedEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *AgentHandoffStartedEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *AgentHandoffStartedEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *AgentHandoffStartedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentHandoffStartedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentHandoffStartedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetPreviousAgentId

`func (o *AgentHandoffStartedEvent) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *AgentHandoffStartedEvent) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *AgentHandoffStartedEvent) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *AgentHandoffStartedEvent) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *AgentHandoffStartedEvent) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *AgentHandoffStartedEvent) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


