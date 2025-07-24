# ToolExecutionDeltaEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool.execution.delta"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**Name** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Arguments** | **string** |  | 

## Methods

### NewToolExecutionDeltaEvent

`func NewToolExecutionDeltaEvent(id string, name BuiltInConnectors, arguments string, ) *ToolExecutionDeltaEvent`

NewToolExecutionDeltaEvent instantiates a new ToolExecutionDeltaEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolExecutionDeltaEventWithDefaults

`func NewToolExecutionDeltaEventWithDefaults() *ToolExecutionDeltaEvent`

NewToolExecutionDeltaEventWithDefaults instantiates a new ToolExecutionDeltaEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolExecutionDeltaEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolExecutionDeltaEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolExecutionDeltaEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolExecutionDeltaEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolExecutionDeltaEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolExecutionDeltaEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolExecutionDeltaEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolExecutionDeltaEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *ToolExecutionDeltaEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ToolExecutionDeltaEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ToolExecutionDeltaEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *ToolExecutionDeltaEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *ToolExecutionDeltaEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolExecutionDeltaEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolExecutionDeltaEvent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ToolExecutionDeltaEvent) GetName() BuiltInConnectors`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolExecutionDeltaEvent) GetNameOk() (*BuiltInConnectors, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolExecutionDeltaEvent) SetName(v BuiltInConnectors)`

SetName sets Name field to given value.


### GetArguments

`func (o *ToolExecutionDeltaEvent) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ToolExecutionDeltaEvent) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ToolExecutionDeltaEvent) SetArguments(v string)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


