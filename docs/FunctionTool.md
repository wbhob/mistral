# FunctionTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "function"]
**Function** | [**Function**](Function.md) |  | 

## Methods

### NewFunctionTool

`func NewFunctionTool(function Function, ) *FunctionTool`

NewFunctionTool instantiates a new FunctionTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionToolWithDefaults

`func NewFunctionToolWithDefaults() *FunctionTool`

NewFunctionToolWithDefaults instantiates a new FunctionTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionTool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *FunctionTool) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *FunctionTool) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *FunctionTool) SetFunction(v Function)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


