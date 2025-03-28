# ModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Data** | Pointer to [**[]ModelListDataInner**](ModelListDataInner.md) |  | [optional] 

## Methods

### NewModelList

`func NewModelList() *ModelList`

NewModelList instantiates a new ModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelListWithDefaults

`func NewModelListWithDefaults() *ModelList`

NewModelListWithDefaults instantiates a new ModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ModelList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *ModelList) GetData() []ModelListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelList) GetDataOk() (*[]ModelListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelList) SetData(v []ModelListDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ModelList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


