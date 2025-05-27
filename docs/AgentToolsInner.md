# AgentToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "function"]
**Function** | [**Function**](Function.md) |  | 
**LibraryIds** | **[]string** | Ids of the library in which to search. | 

## Methods

### NewAgentToolsInner

`func NewAgentToolsInner(function Function, libraryIds []string, ) *AgentToolsInner`

NewAgentToolsInner instantiates a new AgentToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentToolsInnerWithDefaults

`func NewAgentToolsInnerWithDefaults() *AgentToolsInner`

NewAgentToolsInnerWithDefaults instantiates a new AgentToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgentToolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentToolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentToolsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentToolsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *AgentToolsInner) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AgentToolsInner) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AgentToolsInner) SetFunction(v Function)`

SetFunction sets Function field to given value.


### GetLibraryIds

`func (o *AgentToolsInner) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *AgentToolsInner) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *AgentToolsInner) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


