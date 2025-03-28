# DeleteFileOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the deleted file. | 
**Object** | **string** | The object type that was deleted | 
**Deleted** | **bool** | The deletion status. | 

## Methods

### NewDeleteFileOut

`func NewDeleteFileOut(id string, object string, deleted bool, ) *DeleteFileOut`

NewDeleteFileOut instantiates a new DeleteFileOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFileOutWithDefaults

`func NewDeleteFileOutWithDefaults() *DeleteFileOut`

NewDeleteFileOutWithDefaults instantiates a new DeleteFileOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteFileOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteFileOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteFileOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *DeleteFileOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeleteFileOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeleteFileOut) SetObject(v string)`

SetObject sets Object field to given value.


### GetDeleted

`func (o *DeleteFileOut) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteFileOut) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteFileOut) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


