# FIMCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**UsageInfo**](UsageInfo.md) |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Choices** | Pointer to [**[]ChatCompletionChoice**](ChatCompletionChoice.md) |  | [optional] 

## Methods

### NewFIMCompletionResponse

`func NewFIMCompletionResponse() *FIMCompletionResponse`

NewFIMCompletionResponse instantiates a new FIMCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIMCompletionResponseWithDefaults

`func NewFIMCompletionResponseWithDefaults() *FIMCompletionResponse`

NewFIMCompletionResponseWithDefaults instantiates a new FIMCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FIMCompletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FIMCompletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FIMCompletionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FIMCompletionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *FIMCompletionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FIMCompletionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FIMCompletionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FIMCompletionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *FIMCompletionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FIMCompletionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FIMCompletionResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FIMCompletionResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *FIMCompletionResponse) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FIMCompletionResponse) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FIMCompletionResponse) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FIMCompletionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreated

`func (o *FIMCompletionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FIMCompletionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FIMCompletionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FIMCompletionResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetChoices

`func (o *FIMCompletionResponse) GetChoices() []ChatCompletionChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *FIMCompletionResponse) GetChoicesOk() (*[]ChatCompletionChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *FIMCompletionResponse) SetChoices(v []ChatCompletionChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *FIMCompletionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


