# ToolExecutionDoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool.execution.done"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**Name** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewToolExecutionDoneEvent

`func NewToolExecutionDoneEvent(id string, name BuiltInConnectors, ) *ToolExecutionDoneEvent`

NewToolExecutionDoneEvent instantiates a new ToolExecutionDoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolExecutionDoneEventWithDefaults

`func NewToolExecutionDoneEventWithDefaults() *ToolExecutionDoneEvent`

NewToolExecutionDoneEventWithDefaults instantiates a new ToolExecutionDoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolExecutionDoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolExecutionDoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolExecutionDoneEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolExecutionDoneEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolExecutionDoneEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolExecutionDoneEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolExecutionDoneEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolExecutionDoneEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *ToolExecutionDoneEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ToolExecutionDoneEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ToolExecutionDoneEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *ToolExecutionDoneEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *ToolExecutionDoneEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolExecutionDoneEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolExecutionDoneEvent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ToolExecutionDoneEvent) GetName() BuiltInConnectors`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolExecutionDoneEvent) GetNameOk() (*BuiltInConnectors, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolExecutionDoneEvent) SetName(v BuiltInConnectors)`

SetName sets Name field to given value.


### GetInfo

`func (o *ToolExecutionDoneEvent) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ToolExecutionDoneEvent) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ToolExecutionDoneEvent) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ToolExecutionDoneEvent) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


