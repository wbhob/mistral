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

// ResponseFormats An object specifying the format that the model must output. Setting to `{ \"type\": \"json_object\" }` enables JSON mode, which guarantees the message the model generates is in JSON. When using JSON mode you MUST also instruct the model to produce JSON yourself with a system or a user message.
type ResponseFormats string

// List of ResponseFormats
const (
	RESPONSEFORMATS_TEXT ResponseFormats = "text"
	RESPONSEFORMATS_JSON_OBJECT ResponseFormats = "json_object"
	RESPONSEFORMATS_JSON_SCHEMA ResponseFormats = "json_schema"
)

// All allowed values of ResponseFormats enum
var AllowedResponseFormatsEnumValues = []ResponseFormats{
	"text",
	"json_object",
	"json_schema",
}

func (v *ResponseFormats) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseFormats(value)
	for _, existing := range AllowedResponseFormatsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponseFormats", value)
}

// NewResponseFormatsFromValue returns a pointer to a valid ResponseFormats
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponseFormatsFromValue(v string) (*ResponseFormats, error) {
	ev := ResponseFormats(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResponseFormats: valid values are %v", v, AllowedResponseFormatsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponseFormats) IsValid() bool {
	for _, existing := range AllowedResponseFormatsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResponseFormats value
func (v ResponseFormats) Ptr() *ResponseFormats {
	return &v
}

type NullableResponseFormats struct {
	value *ResponseFormats
	isSet bool
}

func (v NullableResponseFormats) Get() *ResponseFormats {
	return v.value
}

func (v *NullableResponseFormats) Set(val *ResponseFormats) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseFormats) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseFormats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseFormats(val *ResponseFormats) *NullableResponseFormats {
	return &NullableResponseFormats{value: val, isSet: true}
}

func (v NullableResponseFormats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseFormats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

