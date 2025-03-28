# CompletionChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Model** | **string** |  | 
**Usage** | Pointer to [**UsageInfo**](UsageInfo.md) |  | [optional] 
**Choices** | [**[]CompletionResponseStreamChoice**](CompletionResponseStreamChoice.md) |  | 

## Methods

### NewCompletionChunk

`func NewCompletionChunk(id string, model string, choices []CompletionResponseStreamChoice, ) *CompletionChunk`

NewCompletionChunk instantiates a new CompletionChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionChunkWithDefaults

`func NewCompletionChunkWithDefaults() *CompletionChunk`

NewCompletionChunkWithDefaults instantiates a new CompletionChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompletionChunk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompletionChunk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompletionChunk) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CompletionChunk) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompletionChunk) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompletionChunk) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompletionChunk) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *CompletionChunk) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CompletionChunk) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CompletionChunk) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CompletionChunk) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModel

`func (o *CompletionChunk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CompletionChunk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CompletionChunk) SetModel(v string)`

SetModel sets Model field to given value.


### GetUsage

`func (o *CompletionChunk) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CompletionChunk) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CompletionChunk) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CompletionChunk) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetChoices

`func (o *CompletionChunk) GetChoices() []CompletionResponseStreamChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *CompletionChunk) GetChoicesOk() (*[]CompletionResponseStreamChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *CompletionChunk) SetChoices(v []CompletionResponseStreamChoice)`

SetChoices sets Choices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


