# MetricOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainLoss** | Pointer to **NullableFloat32** |  | [optional] 
**ValidLoss** | Pointer to **NullableFloat32** |  | [optional] 
**ValidMeanTokenAccuracy** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMetricOut

`func NewMetricOut() *MetricOut`

NewMetricOut instantiates a new MetricOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricOutWithDefaults

`func NewMetricOutWithDefaults() *MetricOut`

NewMetricOutWithDefaults instantiates a new MetricOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainLoss

`func (o *MetricOut) GetTrainLoss() float32`

GetTrainLoss returns the TrainLoss field if non-nil, zero value otherwise.

### GetTrainLossOk

`func (o *MetricOut) GetTrainLossOk() (*float32, bool)`

GetTrainLossOk returns a tuple with the TrainLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainLoss

`func (o *MetricOut) SetTrainLoss(v float32)`

SetTrainLoss sets TrainLoss field to given value.

### HasTrainLoss

`func (o *MetricOut) HasTrainLoss() bool`

HasTrainLoss returns a boolean if a field has been set.

### SetTrainLossNil

`func (o *MetricOut) SetTrainLossNil(b bool)`

 SetTrainLossNil sets the value for TrainLoss to be an explicit nil

### UnsetTrainLoss
`func (o *MetricOut) UnsetTrainLoss()`

UnsetTrainLoss ensures that no value is present for TrainLoss, not even an explicit nil
### GetValidLoss

`func (o *MetricOut) GetValidLoss() float32`

GetValidLoss returns the ValidLoss field if non-nil, zero value otherwise.

### GetValidLossOk

`func (o *MetricOut) GetValidLossOk() (*float32, bool)`

GetValidLossOk returns a tuple with the ValidLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLoss

`func (o *MetricOut) SetValidLoss(v float32)`

SetValidLoss sets ValidLoss field to given value.

### HasValidLoss

`func (o *MetricOut) HasValidLoss() bool`

HasValidLoss returns a boolean if a field has been set.

### SetValidLossNil

`func (o *MetricOut) SetValidLossNil(b bool)`

 SetValidLossNil sets the value for ValidLoss to be an explicit nil

### UnsetValidLoss
`func (o *MetricOut) UnsetValidLoss()`

UnsetValidLoss ensures that no value is present for ValidLoss, not even an explicit nil
### GetValidMeanTokenAccuracy

`func (o *MetricOut) GetValidMeanTokenAccuracy() float32`

GetValidMeanTokenAccuracy returns the ValidMeanTokenAccuracy field if non-nil, zero value otherwise.

### GetValidMeanTokenAccuracyOk

`func (o *MetricOut) GetValidMeanTokenAccuracyOk() (*float32, bool)`

GetValidMeanTokenAccuracyOk returns a tuple with the ValidMeanTokenAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidMeanTokenAccuracy

`func (o *MetricOut) SetValidMeanTokenAccuracy(v float32)`

SetValidMeanTokenAccuracy sets ValidMeanTokenAccuracy field to given value.

### HasValidMeanTokenAccuracy

`func (o *MetricOut) HasValidMeanTokenAccuracy() bool`

HasValidMeanTokenAccuracy returns a boolean if a field has been set.

### SetValidMeanTokenAccuracyNil

`func (o *MetricOut) SetValidMeanTokenAccuracyNil(b bool)`

 SetValidMeanTokenAccuracyNil sets the value for ValidMeanTokenAccuracy to be an explicit nil

### UnsetValidMeanTokenAccuracy
`func (o *MetricOut) UnsetValidMeanTokenAccuracy()`

UnsetValidMeanTokenAccuracy ensures that no value is present for ValidMeanTokenAccuracy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


