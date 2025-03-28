# DeleteModelOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the deleted model. | 
**Object** | Pointer to **string** | The object type that was deleted | [optional] [default to "model"]
**Deleted** | Pointer to **bool** | The deletion status | [optional] [default to true]

## Methods

### NewDeleteModelOut

`func NewDeleteModelOut(id string, ) *DeleteModelOut`

NewDeleteModelOut instantiates a new DeleteModelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteModelOutWithDefaults

`func NewDeleteModelOutWithDefaults() *DeleteModelOut`

NewDeleteModelOutWithDefaults instantiates a new DeleteModelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteModelOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteModelOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteModelOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *DeleteModelOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeleteModelOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeleteModelOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DeleteModelOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDeleted

`func (o *DeleteModelOut) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteModelOut) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteModelOut) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteModelOut) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


