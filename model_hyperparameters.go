/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"fmt"
)


// Hyperparameters struct for Hyperparameters
type Hyperparameters struct {
	ClassifierTrainingParametersIn *ClassifierTrainingParametersIn
	CompletionTrainingParametersIn *CompletionTrainingParametersIn
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Hyperparameters) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ClassifierTrainingParametersIn
	err = json.Unmarshal(data, &dst.ClassifierTrainingParametersIn);
	if err == nil {
		jsonClassifierTrainingParametersIn, _ := json.Marshal(dst.ClassifierTrainingParametersIn)
		if string(jsonClassifierTrainingParametersIn) == "{}" { // empty struct
			dst.ClassifierTrainingParametersIn = nil
		} else {
			return nil // data stored in dst.ClassifierTrainingParametersIn, return on the first match
		}
	} else {
		dst.ClassifierTrainingParametersIn = nil
	}

	// try to unmarshal JSON data into CompletionTrainingParametersIn
	err = json.Unmarshal(data, &dst.CompletionTrainingParametersIn);
	if err == nil {
		jsonCompletionTrainingParametersIn, _ := json.Marshal(dst.CompletionTrainingParametersIn)
		if string(jsonCompletionTrainingParametersIn) == "{}" { // empty struct
			dst.CompletionTrainingParametersIn = nil
		} else {
			return nil // data stored in dst.CompletionTrainingParametersIn, return on the first match
		}
	} else {
		dst.CompletionTrainingParametersIn = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Hyperparameters)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Hyperparameters) MarshalJSON() ([]byte, error) {
	if src.ClassifierTrainingParametersIn != nil {
		return json.Marshal(&src.ClassifierTrainingParametersIn)
	}

	if src.CompletionTrainingParametersIn != nil {
		return json.Marshal(&src.CompletionTrainingParametersIn)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableHyperparameters struct {
	value *Hyperparameters
	isSet bool
}

func (v NullableHyperparameters) Get() *Hyperparameters {
	return v.value
}

func (v *NullableHyperparameters) Set(val *Hyperparameters) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperparameters) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperparameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperparameters(val *Hyperparameters) *NullableHyperparameters {
	return &NullableHyperparameters{value: val, isSet: true}
}

func (v NullableHyperparameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperparameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


