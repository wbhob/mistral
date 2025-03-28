# Tool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ToolTypes**](ToolTypes.md) |  | [optional] 
**Function** | [**Function**](Function.md) |  | 

## Methods

### NewTool

`func NewTool(function Function, ) *Tool`

NewTool instantiates a new Tool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolWithDefaults

`func NewToolWithDefaults() *Tool`

NewToolWithDefaults instantiates a new Tool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Tool) GetType() ToolTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tool) GetTypeOk() (*ToolTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tool) SetType(v ToolTypes)`

SetType sets Type field to given value.

### HasType

`func (o *Tool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *Tool) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Tool) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Tool) SetFunction(v Function)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


