# CheckpointOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**MetricOut**](MetricOut.md) |  | 
**StepNumber** | **int32** | The step number that the checkpoint was created at. | 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) for when the checkpoint was created. | 

## Methods

### NewCheckpointOut

`func NewCheckpointOut(metrics MetricOut, stepNumber int32, createdAt int32, ) *CheckpointOut`

NewCheckpointOut instantiates a new CheckpointOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointOutWithDefaults

`func NewCheckpointOutWithDefaults() *CheckpointOut`

NewCheckpointOutWithDefaults instantiates a new CheckpointOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *CheckpointOut) GetMetrics() MetricOut`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CheckpointOut) GetMetricsOk() (*MetricOut, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CheckpointOut) SetMetrics(v MetricOut)`

SetMetrics sets Metrics field to given value.


### GetStepNumber

`func (o *CheckpointOut) GetStepNumber() int32`

GetStepNumber returns the StepNumber field if non-nil, zero value otherwise.

### GetStepNumberOk

`func (o *CheckpointOut) GetStepNumberOk() (*int32, bool)`

GetStepNumberOk returns a tuple with the StepNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepNumber

`func (o *CheckpointOut) SetStepNumber(v int32)`

SetStepNumber sets StepNumber field to given value.


### GetCreatedAt

`func (o *CheckpointOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckpointOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckpointOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


