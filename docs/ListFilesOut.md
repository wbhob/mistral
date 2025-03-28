# ListFilesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FileSchema**](FileSchema.md) |  | 
**Object** | **string** |  | 
**Total** | **int32** |  | 

## Methods

### NewListFilesOut

`func NewListFilesOut(data []FileSchema, object string, total int32, ) *ListFilesOut`

NewListFilesOut instantiates a new ListFilesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFilesOutWithDefaults

`func NewListFilesOutWithDefaults() *ListFilesOut`

NewListFilesOutWithDefaults instantiates a new ListFilesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFilesOut) GetData() []FileSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFilesOut) GetDataOk() (*[]FileSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFilesOut) SetData(v []FileSchema)`

SetData sets Data field to given value.


### GetObject

`func (o *ListFilesOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListFilesOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListFilesOut) SetObject(v string)`

SetObject sets Object field to given value.


### GetTotal

`func (o *ListFilesOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListFilesOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListFilesOut) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


