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


// Input Chat to classify
type Input struct {
	ArrayOfMessagesInner *[]MessagesInner
	ArrayOfArrayOfMessagesInner *[][]MessagesInner
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Input) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfMessagesInner
	err = json.Unmarshal(data, &dst.ArrayOfMessagesInner);
	if err == nil {
		jsonArrayOfMessagesInner, _ := json.Marshal(dst.ArrayOfMessagesInner)
		if string(jsonArrayOfMessagesInner) == "{}" { // empty struct
			dst.ArrayOfMessagesInner = nil
		} else {
			return nil // data stored in dst.ArrayOfMessagesInner, return on the first match
		}
	} else {
		dst.ArrayOfMessagesInner = nil
	}

	// try to unmarshal JSON data into ArrayOfArrayOfMessagesInner
	err = json.Unmarshal(data, &dst.ArrayOfArrayOfMessagesInner);
	if err == nil {
		jsonArrayOfArrayOfMessagesInner, _ := json.Marshal(dst.ArrayOfArrayOfMessagesInner)
		if string(jsonArrayOfArrayOfMessagesInner) == "{}" { // empty struct
			dst.ArrayOfArrayOfMessagesInner = nil
		} else {
			return nil // data stored in dst.ArrayOfArrayOfMessagesInner, return on the first match
		}
	} else {
		dst.ArrayOfArrayOfMessagesInner = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Input)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Input) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMessagesInner != nil {
		return json.Marshal(&src.ArrayOfMessagesInner)
	}

	if src.ArrayOfArrayOfMessagesInner != nil {
		return json.Marshal(&src.ArrayOfArrayOfMessagesInner)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableInput struct {
	value *Input
	isSet bool
}

func (v NullableInput) Get() *Input {
	return v.value
}

func (v *NullableInput) Set(val *Input) {
	v.value = val
	v.isSet = true
}

func (v NullableInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInput(val *Input) *NullableInput {
	return &NullableInput{value: val, isSet: true}
}

func (v NullableInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


