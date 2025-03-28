# EventOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the event. | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) of the event. | 

## Methods

### NewEventOut

`func NewEventOut(name string, createdAt int32, ) *EventOut`

NewEventOut instantiates a new EventOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventOutWithDefaults

`func NewEventOutWithDefaults() *EventOut`

NewEventOutWithDefaults instantiates a new EventOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventOut) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *EventOut) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventOut) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventOut) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventOut) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EventOut) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EventOut) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCreatedAt

`func (o *EventOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


