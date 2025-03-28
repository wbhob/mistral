# ResponseBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**UsageInfo**](UsageInfo.md) |  | [optional] 

## Methods

### NewResponseBase

`func NewResponseBase() *ResponseBase`

NewResponseBase instantiates a new ResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBaseWithDefaults

`func NewResponseBaseWithDefaults() *ResponseBase`

NewResponseBaseWithDefaults instantiates a new ResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ResponseBase) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseBase) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseBase) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseBase) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *ResponseBase) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponseBase) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponseBase) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ResponseBase) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *ResponseBase) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ResponseBase) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ResponseBase) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ResponseBase) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


