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


// Response struct for Response
type Response struct {
	LegacyJobMetadataOut *LegacyJobMetadataOut
	ResponseAnyOf *ResponseAnyOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LegacyJobMetadataOut
	err = json.Unmarshal(data, &dst.LegacyJobMetadataOut);
	if err == nil {
		jsonLegacyJobMetadataOut, _ := json.Marshal(dst.LegacyJobMetadataOut)
		if string(jsonLegacyJobMetadataOut) == "{}" { // empty struct
			dst.LegacyJobMetadataOut = nil
		} else {
			return nil // data stored in dst.LegacyJobMetadataOut, return on the first match
		}
	} else {
		dst.LegacyJobMetadataOut = nil
	}

	// try to unmarshal JSON data into ResponseAnyOf
	err = json.Unmarshal(data, &dst.ResponseAnyOf);
	if err == nil {
		jsonResponseAnyOf, _ := json.Marshal(dst.ResponseAnyOf)
		if string(jsonResponseAnyOf) == "{}" { // empty struct
			dst.ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.ResponseAnyOf, return on the first match
		}
	} else {
		dst.ResponseAnyOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Response) MarshalJSON() ([]byte, error) {
	if src.LegacyJobMetadataOut != nil {
		return json.Marshal(&src.LegacyJobMetadataOut)
	}

	if src.ResponseAnyOf != nil {
		return json.Marshal(&src.ResponseAnyOf)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


