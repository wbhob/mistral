# BatchJobsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BatchJobOut**](BatchJobOut.md) |  | [optional] [default to []]
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Total** | **int32** |  | 

## Methods

### NewBatchJobsOut

`func NewBatchJobsOut(total int32, ) *BatchJobsOut`

NewBatchJobsOut instantiates a new BatchJobsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobsOutWithDefaults

`func NewBatchJobsOutWithDefaults() *BatchJobsOut`

NewBatchJobsOutWithDefaults instantiates a new BatchJobsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BatchJobsOut) GetData() []BatchJobOut`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BatchJobsOut) GetDataOk() (*[]BatchJobOut, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BatchJobsOut) SetData(v []BatchJobOut)`

SetData sets Data field to given value.

### HasData

`func (o *BatchJobsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetObject

`func (o *BatchJobsOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BatchJobsOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BatchJobsOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BatchJobsOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTotal

`func (o *BatchJobsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BatchJobsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BatchJobsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


