# FunctionCallEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "function.call.delta"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**Name** | **string** |  | 
**ToolCallId** | **string** |  | 
**Arguments** | **string** |  | 

## Methods

### NewFunctionCallEvent

`func NewFunctionCallEvent(id string, name string, toolCallId string, arguments string, ) *FunctionCallEvent`

NewFunctionCallEvent instantiates a new FunctionCallEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCallEventWithDefaults

`func NewFunctionCallEventWithDefaults() *FunctionCallEvent`

NewFunctionCallEventWithDefaults instantiates a new FunctionCallEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionCallEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionCallEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionCallEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionCallEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionCallEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionCallEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionCallEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionCallEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *FunctionCallEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *FunctionCallEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *FunctionCallEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *FunctionCallEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *FunctionCallEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionCallEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionCallEvent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FunctionCallEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionCallEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionCallEvent) SetName(v string)`

SetName sets Name field to given value.


### GetToolCallId

`func (o *FunctionCallEvent) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *FunctionCallEvent) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *FunctionCallEvent) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetArguments

`func (o *FunctionCallEvent) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *FunctionCallEvent) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *FunctionCallEvent) SetArguments(v string)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


