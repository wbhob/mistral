# AgentHandoffEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "agent.handoff"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewAgentHandoffEntry

`func NewAgentHandoffEntry(previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string, ) *AgentHandoffEntry`

NewAgentHandoffEntry instantiates a new AgentHandoffEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentHandoffEntryWithDefaults

`func NewAgentHandoffEntryWithDefaults() *AgentHandoffEntry`

NewAgentHandoffEntryWithDefaults instantiates a new AgentHandoffEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AgentHandoffEntry) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AgentHandoffEntry) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AgentHandoffEntry) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AgentHandoffEntry) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *AgentHandoffEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentHandoffEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentHandoffEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentHandoffEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentHandoffEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentHandoffEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentHandoffEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentHandoffEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *AgentHandoffEntry) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *AgentHandoffEntry) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *AgentHandoffEntry) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *AgentHandoffEntry) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *AgentHandoffEntry) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *AgentHandoffEntry) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetId

`func (o *AgentHandoffEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentHandoffEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentHandoffEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentHandoffEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPreviousAgentId

`func (o *AgentHandoffEntry) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *AgentHandoffEntry) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *AgentHandoffEntry) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *AgentHandoffEntry) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *AgentHandoffEntry) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *AgentHandoffEntry) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.


### GetNextAgentId

`func (o *AgentHandoffEntry) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *AgentHandoffEntry) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *AgentHandoffEntry) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *AgentHandoffEntry) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *AgentHandoffEntry) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *AgentHandoffEntry) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


