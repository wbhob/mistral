# EmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**UsageInfo**](UsageInfo.md) |  | [optional] 
**Data** | **[]interface{}** |  | 

## Methods

### NewEmbeddingResponse

`func NewEmbeddingResponse(data []interface{}, ) *EmbeddingResponse`

NewEmbeddingResponse instantiates a new EmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingResponseWithDefaults

`func NewEmbeddingResponseWithDefaults() *EmbeddingResponse`

NewEmbeddingResponseWithDefaults instantiates a new EmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbeddingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbeddingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbeddingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmbeddingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *EmbeddingResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EmbeddingResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EmbeddingResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *EmbeddingResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *EmbeddingResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EmbeddingResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *EmbeddingResponse) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EmbeddingResponse) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EmbeddingResponse) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EmbeddingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetData

`func (o *EmbeddingResponse) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmbeddingResponse) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmbeddingResponse) SetData(v []interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


