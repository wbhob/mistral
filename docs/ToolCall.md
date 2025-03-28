# ToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [default to "null"]
**Type** | Pointer to [**ToolTypes**](ToolTypes.md) |  | [optional] 
**Function** | [**FunctionCall**](FunctionCall.md) |  | 
**Index** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewToolCall

`func NewToolCall(function FunctionCall, ) *ToolCall`

NewToolCall instantiates a new ToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCallWithDefaults

`func NewToolCallWithDefaults() *ToolCall`

NewToolCallWithDefaults instantiates a new ToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ToolCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ToolCall) GetType() ToolTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolCall) GetTypeOk() (*ToolTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolCall) SetType(v ToolTypes)`

SetType sets Type field to given value.

### HasType

`func (o *ToolCall) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *ToolCall) GetFunction() FunctionCall`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ToolCall) GetFunctionOk() (*FunctionCall, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ToolCall) SetFunction(v FunctionCall)`

SetFunction sets Function field to given value.


### GetIndex

`func (o *ToolCall) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ToolCall) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ToolCall) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ToolCall) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


