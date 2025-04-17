/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ClassifierTargetOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassifierTargetOut{}

// ClassifierTargetOut struct for ClassifierTargetOut
type ClassifierTargetOut struct {
	Name string `json:"name"`
	Labels []string `json:"labels"`
	Weight float32 `json:"weight"`
	LossFunction FTClassifierLossFunction `json:"loss_function"`
}

type _ClassifierTargetOut ClassifierTargetOut

// NewClassifierTargetOut instantiates a new ClassifierTargetOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifierTargetOut(name string, labels []string, weight float32, lossFunction FTClassifierLossFunction) *ClassifierTargetOut {
	this := ClassifierTargetOut{}
	this.Name = name
	this.Labels = labels
	this.Weight = weight
	this.LossFunction = lossFunction
	return &this
}

// NewClassifierTargetOutWithDefaults instantiates a new ClassifierTargetOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifierTargetOutWithDefaults() *ClassifierTargetOut {
	this := ClassifierTargetOut{}
	return &this
}

// GetName returns the Name field value
func (o *ClassifierTargetOut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClassifierTargetOut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClassifierTargetOut) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *ClassifierTargetOut) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ClassifierTargetOut) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ClassifierTargetOut) SetLabels(v []string) {
	o.Labels = v
}

// GetWeight returns the Weight field value
func (o *ClassifierTargetOut) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ClassifierTargetOut) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ClassifierTargetOut) SetWeight(v float32) {
	o.Weight = v
}

// GetLossFunction returns the LossFunction field value
func (o *ClassifierTargetOut) GetLossFunction() FTClassifierLossFunction {
	if o == nil {
		var ret FTClassifierLossFunction
		return ret
	}

	return o.LossFunction
}

// GetLossFunctionOk returns a tuple with the LossFunction field value
// and a boolean to check if the value has been set.
func (o *ClassifierTargetOut) GetLossFunctionOk() (*FTClassifierLossFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LossFunction, true
}

// SetLossFunction sets field value
func (o *ClassifierTargetOut) SetLossFunction(v FTClassifierLossFunction) {
	o.LossFunction = v
}

func (o ClassifierTargetOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassifierTargetOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["labels"] = o.Labels
	toSerialize["weight"] = o.Weight
	toSerialize["loss_function"] = o.LossFunction
	return toSerialize, nil
}

func (o *ClassifierTargetOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"labels",
		"weight",
		"loss_function",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varClassifierTargetOut := _ClassifierTargetOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClassifierTargetOut)

	if err != nil {
		return err
	}

	*o = ClassifierTargetOut(varClassifierTargetOut)

	return err
}

type NullableClassifierTargetOut struct {
	value *ClassifierTargetOut
	isSet bool
}

func (v NullableClassifierTargetOut) Get() *ClassifierTargetOut {
	return v.value
}

func (v *NullableClassifierTargetOut) Set(val *ClassifierTargetOut) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifierTargetOut) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifierTargetOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifierTargetOut(val *ClassifierTargetOut) *NullableClassifierTargetOut {
	return &NullableClassifierTargetOut{value: val, isSet: true}
}

func (v NullableClassifierTargetOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifierTargetOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


