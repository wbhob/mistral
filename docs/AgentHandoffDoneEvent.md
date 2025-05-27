# AgentHandoffDoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "agent.handoff.done"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewAgentHandoffDoneEvent

`func NewAgentHandoffDoneEvent(id string, nextAgentId string, nextAgentName string, ) *AgentHandoffDoneEvent`

NewAgentHandoffDoneEvent instantiates a new AgentHandoffDoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentHandoffDoneEventWithDefaults

`func NewAgentHandoffDoneEventWithDefaults() *AgentHandoffDoneEvent`

NewAgentHandoffDoneEventWithDefaults instantiates a new AgentHandoffDoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgentHandoffDoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentHandoffDoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentHandoffDoneEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentHandoffDoneEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentHandoffDoneEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentHandoffDoneEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentHandoffDoneEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentHandoffDoneEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *AgentHandoffDoneEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *AgentHandoffDoneEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *AgentHandoffDoneEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *AgentHandoffDoneEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *AgentHandoffDoneEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentHandoffDoneEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentHandoffDoneEvent) SetId(v string)`

SetId sets Id field to given value.


### GetNextAgentId

`func (o *AgentHandoffDoneEvent) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *AgentHandoffDoneEvent) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *AgentHandoffDoneEvent) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *AgentHandoffDoneEvent) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *AgentHandoffDoneEvent) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *AgentHandoffDoneEvent) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


