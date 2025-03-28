/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
)

// checks if the MetricOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricOut{}

// MetricOut Metrics at the step number during the fine-tuning job. Use these metrics to assess if the training is going smoothly (loss should decrease, token accuracy should increase).
type MetricOut struct {
	TrainLoss NullableFloat32 `json:"train_loss,omitempty"`
	ValidLoss NullableFloat32 `json:"valid_loss,omitempty"`
	ValidMeanTokenAccuracy NullableFloat32 `json:"valid_mean_token_accuracy,omitempty"`
}

// NewMetricOut instantiates a new MetricOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricOut() *MetricOut {
	this := MetricOut{}
	return &this
}

// NewMetricOutWithDefaults instantiates a new MetricOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricOutWithDefaults() *MetricOut {
	this := MetricOut{}
	return &this
}

// GetTrainLoss returns the TrainLoss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricOut) GetTrainLoss() float32 {
	if o == nil || IsNil(o.TrainLoss.Get()) {
		var ret float32
		return ret
	}
	return *o.TrainLoss.Get()
}

// GetTrainLossOk returns a tuple with the TrainLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricOut) GetTrainLossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainLoss.Get(), o.TrainLoss.IsSet()
}

// HasTrainLoss returns a boolean if a field has been set.
func (o *MetricOut) HasTrainLoss() bool {
	if o != nil && o.TrainLoss.IsSet() {
		return true
	}

	return false
}

// SetTrainLoss gets a reference to the given NullableFloat32 and assigns it to the TrainLoss field.
func (o *MetricOut) SetTrainLoss(v float32) {
	o.TrainLoss.Set(&v)
}
// SetTrainLossNil sets the value for TrainLoss to be an explicit nil
func (o *MetricOut) SetTrainLossNil() {
	o.TrainLoss.Set(nil)
}

// UnsetTrainLoss ensures that no value is present for TrainLoss, not even an explicit nil
func (o *MetricOut) UnsetTrainLoss() {
	o.TrainLoss.Unset()
}

// GetValidLoss returns the ValidLoss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricOut) GetValidLoss() float32 {
	if o == nil || IsNil(o.ValidLoss.Get()) {
		var ret float32
		return ret
	}
	return *o.ValidLoss.Get()
}

// GetValidLossOk returns a tuple with the ValidLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricOut) GetValidLossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidLoss.Get(), o.ValidLoss.IsSet()
}

// HasValidLoss returns a boolean if a field has been set.
func (o *MetricOut) HasValidLoss() bool {
	if o != nil && o.ValidLoss.IsSet() {
		return true
	}

	return false
}

// SetValidLoss gets a reference to the given NullableFloat32 and assigns it to the ValidLoss field.
func (o *MetricOut) SetValidLoss(v float32) {
	o.ValidLoss.Set(&v)
}
// SetValidLossNil sets the value for ValidLoss to be an explicit nil
func (o *MetricOut) SetValidLossNil() {
	o.ValidLoss.Set(nil)
}

// UnsetValidLoss ensures that no value is present for ValidLoss, not even an explicit nil
func (o *MetricOut) UnsetValidLoss() {
	o.ValidLoss.Unset()
}

// GetValidMeanTokenAccuracy returns the ValidMeanTokenAccuracy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricOut) GetValidMeanTokenAccuracy() float32 {
	if o == nil || IsNil(o.ValidMeanTokenAccuracy.Get()) {
		var ret float32
		return ret
	}
	return *o.ValidMeanTokenAccuracy.Get()
}

// GetValidMeanTokenAccuracyOk returns a tuple with the ValidMeanTokenAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricOut) GetValidMeanTokenAccuracyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidMeanTokenAccuracy.Get(), o.ValidMeanTokenAccuracy.IsSet()
}

// HasValidMeanTokenAccuracy returns a boolean if a field has been set.
func (o *MetricOut) HasValidMeanTokenAccuracy() bool {
	if o != nil && o.ValidMeanTokenAccuracy.IsSet() {
		return true
	}

	return false
}

// SetValidMeanTokenAccuracy gets a reference to the given NullableFloat32 and assigns it to the ValidMeanTokenAccuracy field.
func (o *MetricOut) SetValidMeanTokenAccuracy(v float32) {
	o.ValidMeanTokenAccuracy.Set(&v)
}
// SetValidMeanTokenAccuracyNil sets the value for ValidMeanTokenAccuracy to be an explicit nil
func (o *MetricOut) SetValidMeanTokenAccuracyNil() {
	o.ValidMeanTokenAccuracy.Set(nil)
}

// UnsetValidMeanTokenAccuracy ensures that no value is present for ValidMeanTokenAccuracy, not even an explicit nil
func (o *MetricOut) UnsetValidMeanTokenAccuracy() {
	o.ValidMeanTokenAccuracy.Unset()
}

func (o MetricOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TrainLoss.IsSet() {
		toSerialize["train_loss"] = o.TrainLoss.Get()
	}
	if o.ValidLoss.IsSet() {
		toSerialize["valid_loss"] = o.ValidLoss.Get()
	}
	if o.ValidMeanTokenAccuracy.IsSet() {
		toSerialize["valid_mean_token_accuracy"] = o.ValidMeanTokenAccuracy.Get()
	}
	return toSerialize, nil
}

type NullableMetricOut struct {
	value *MetricOut
	isSet bool
}

func (v NullableMetricOut) Get() *MetricOut {
	return v.value
}

func (v *NullableMetricOut) Set(val *MetricOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricOut(val *MetricOut) *NullableMetricOut {
	return &NullableMetricOut{value: val, isSet: true}
}

func (v NullableMetricOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


