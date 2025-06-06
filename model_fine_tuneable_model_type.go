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

// FineTuneableModelType the model 'FineTuneableModelType'
type FineTuneableModelType string

// List of FineTuneableModelType
const (
	FINETUNEABLEMODELTYPE_COMPLETION FineTuneableModelType = "completion"
	FINETUNEABLEMODELTYPE_CLASSIFIER FineTuneableModelType = "classifier"
)

// All allowed values of FineTuneableModelType enum
var AllowedFineTuneableModelTypeEnumValues = []FineTuneableModelType{
	"completion",
	"classifier",
}

func (v *FineTuneableModelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FineTuneableModelType(value)
	for _, existing := range AllowedFineTuneableModelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FineTuneableModelType", value)
}

// NewFineTuneableModelTypeFromValue returns a pointer to a valid FineTuneableModelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFineTuneableModelTypeFromValue(v string) (*FineTuneableModelType, error) {
	ev := FineTuneableModelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FineTuneableModelType: valid values are %v", v, AllowedFineTuneableModelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FineTuneableModelType) IsValid() bool {
	for _, existing := range AllowedFineTuneableModelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FineTuneableModelType value
func (v FineTuneableModelType) Ptr() *FineTuneableModelType {
	return &v
}

type NullableFineTuneableModelType struct {
	value *FineTuneableModelType
	isSet bool
}

func (v NullableFineTuneableModelType) Get() *FineTuneableModelType {
	return v.value
}

func (v *NullableFineTuneableModelType) Set(val *FineTuneableModelType) {
	v.value = val
	v.isSet = true
}

func (v NullableFineTuneableModelType) IsSet() bool {
	return v.isSet
}

func (v *NullableFineTuneableModelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFineTuneableModelType(val *FineTuneableModelType) *NullableFineTuneableModelType {
	return &NullableFineTuneableModelType{value: val, isSet: true}
}

func (v NullableFineTuneableModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFineTuneableModelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

