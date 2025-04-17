/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"fmt"
)

// FineTuneableModel The name of the model to fine-tune.
type FineTuneableModel string

// List of FineTuneableModel
const (
	FINETUNEABLEMODEL_OPEN_MISTRAL_7B FineTuneableModel = "open-mistral-7b"
	FINETUNEABLEMODEL_MISTRAL_SMALL_LATEST FineTuneableModel = "mistral-small-latest"
	FINETUNEABLEMODEL_CODESTRAL_LATEST FineTuneableModel = "codestral-latest"
	FINETUNEABLEMODEL_MISTRAL_LARGE_LATEST FineTuneableModel = "mistral-large-latest"
	FINETUNEABLEMODEL_OPEN_MISTRAL_NEMO FineTuneableModel = "open-mistral-nemo"
	FINETUNEABLEMODEL_MINISTRAL_3B_LATEST FineTuneableModel = "ministral-3b-latest"
	FINETUNEABLEMODEL_MINISTRAL_8B_LATEST FineTuneableModel = "ministral-8b-latest"
)

// All allowed values of FineTuneableModel enum
var AllowedFineTuneableModelEnumValues = []FineTuneableModel{
	"open-mistral-7b",
	"mistral-small-latest",
	"codestral-latest",
	"mistral-large-latest",
	"open-mistral-nemo",
	"ministral-3b-latest",
	"ministral-8b-latest",
}

func (v *FineTuneableModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FineTuneableModel(value)
	for _, existing := range AllowedFineTuneableModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FineTuneableModel", value)
}

// NewFineTuneableModelFromValue returns a pointer to a valid FineTuneableModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFineTuneableModelFromValue(v string) (*FineTuneableModel, error) {
	ev := FineTuneableModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FineTuneableModel: valid values are %v", v, AllowedFineTuneableModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FineTuneableModel) IsValid() bool {
	for _, existing := range AllowedFineTuneableModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FineTuneableModel value
func (v FineTuneableModel) Ptr() *FineTuneableModel {
	return &v
}

type NullableFineTuneableModel struct {
	value *FineTuneableModel
	isSet bool
}

func (v NullableFineTuneableModel) Get() *FineTuneableModel {
	return v.value
}

func (v *NullableFineTuneableModel) Set(val *FineTuneableModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFineTuneableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFineTuneableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFineTuneableModel(val *FineTuneableModel) *NullableFineTuneableModel {
	return &NullableFineTuneableModel{value: val, isSet: true}
}

func (v NullableFineTuneableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFineTuneableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

