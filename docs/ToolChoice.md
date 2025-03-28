# ToolChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ToolTypes**](ToolTypes.md) |  | [optional] 
**Function** | [**FunctionName**](FunctionName.md) |  | 

## Methods

### NewToolChoice

`func NewToolChoice(function FunctionName, ) *ToolChoice`

NewToolChoice instantiates a new ToolChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolChoiceWithDefaults

`func NewToolChoiceWithDefaults() *ToolChoice`

NewToolChoiceWithDefaults instantiates a new ToolChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolChoice) GetType() ToolTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolChoice) GetTypeOk() (*ToolTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolChoice) SetType(v ToolTypes)`

SetType sets Type field to given value.

### HasType

`func (o *ToolChoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *ToolChoice) GetFunction() FunctionName`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ToolChoice) GetFunctionOk() (*FunctionName, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ToolChoice) SetFunction(v FunctionName)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


