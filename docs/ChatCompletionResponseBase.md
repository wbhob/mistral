# ChatCompletionResponseBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**UsageInfo**](UsageInfo.md) |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 

## Methods

### NewChatCompletionResponseBase

`func NewChatCompletionResponseBase() *ChatCompletionResponseBase`

NewChatCompletionResponseBase instantiates a new ChatCompletionResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseBaseWithDefaults

`func NewChatCompletionResponseBaseWithDefaults() *ChatCompletionResponseBase`

NewChatCompletionResponseBaseWithDefaults instantiates a new ChatCompletionResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatCompletionResponseBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatCompletionResponseBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatCompletionResponseBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatCompletionResponseBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ChatCompletionResponseBase) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatCompletionResponseBase) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatCompletionResponseBase) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChatCompletionResponseBase) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *ChatCompletionResponseBase) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionResponseBase) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionResponseBase) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatCompletionResponseBase) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *ChatCompletionResponseBase) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatCompletionResponseBase) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatCompletionResponseBase) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatCompletionResponseBase) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreated

`func (o *ChatCompletionResponseBase) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatCompletionResponseBase) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatCompletionResponseBase) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChatCompletionResponseBase) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


