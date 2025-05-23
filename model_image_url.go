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


// ImageUrl struct for ImageUrl
type ImageUrl struct {
	ImageURL *ImageURL
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ImageUrl) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ImageURL
	err = json.Unmarshal(data, &dst.ImageURL);
	if err == nil {
		jsonImageURL, _ := json.Marshal(dst.ImageURL)
		if string(jsonImageURL) == "{}" { // empty struct
			dst.ImageURL = nil
		} else {
			return nil // data stored in dst.ImageURL, return on the first match
		}
	} else {
		dst.ImageURL = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ImageUrl)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ImageUrl) MarshalJSON() ([]byte, error) {
	if src.ImageURL != nil {
		return json.Marshal(&src.ImageURL)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableImageUrl struct {
	value *ImageUrl
	isSet bool
}

func (v NullableImageUrl) Get() *ImageUrl {
	return v.value
}

func (v *NullableImageUrl) Set(val *ImageUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableImageUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableImageUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageUrl(val *ImageUrl) *NullableImageUrl {
	return &NullableImageUrl{value: val, isSet: true}
}

func (v NullableImageUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


