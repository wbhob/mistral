# ToolExecutionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "tool.execution"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Arguments** | **string** |  | 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewToolExecutionEntry

`func NewToolExecutionEntry(name BuiltInConnectors, arguments string, ) *ToolExecutionEntry`

NewToolExecutionEntry instantiates a new ToolExecutionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolExecutionEntryWithDefaults

`func NewToolExecutionEntryWithDefaults() *ToolExecutionEntry`

NewToolExecutionEntryWithDefaults instantiates a new ToolExecutionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ToolExecutionEntry) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ToolExecutionEntry) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ToolExecutionEntry) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ToolExecutionEntry) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *ToolExecutionEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolExecutionEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolExecutionEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolExecutionEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolExecutionEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolExecutionEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolExecutionEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolExecutionEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ToolExecutionEntry) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ToolExecutionEntry) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ToolExecutionEntry) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ToolExecutionEntry) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ToolExecutionEntry) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ToolExecutionEntry) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetId

`func (o *ToolExecutionEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolExecutionEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolExecutionEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ToolExecutionEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ToolExecutionEntry) GetName() BuiltInConnectors`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolExecutionEntry) GetNameOk() (*BuiltInConnectors, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolExecutionEntry) SetName(v BuiltInConnectors)`

SetName sets Name field to given value.


### GetArguments

`func (o *ToolExecutionEntry) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ToolExecutionEntry) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ToolExecutionEntry) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetInfo

`func (o *ToolExecutionEntry) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ToolExecutionEntry) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ToolExecutionEntry) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ToolExecutionEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


