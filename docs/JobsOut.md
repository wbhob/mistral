# JobsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseAnyOf**](ResponseAnyOf.md) |  | [optional] [default to []]
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Total** | **int32** |  | 

## Methods

### NewJobsOut

`func NewJobsOut(total int32, ) *JobsOut`

NewJobsOut instantiates a new JobsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsOutWithDefaults

`func NewJobsOutWithDefaults() *JobsOut`

NewJobsOutWithDefaults instantiates a new JobsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JobsOut) GetData() []ResponseAnyOf`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobsOut) GetDataOk() (*[]ResponseAnyOf, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobsOut) SetData(v []ResponseAnyOf)`

SetData sets Data field to given value.

### HasData

`func (o *JobsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetObject

`func (o *JobsOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *JobsOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *JobsOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *JobsOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTotal

`func (o *JobsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *JobsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *JobsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


