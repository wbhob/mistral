/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UsageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageInfo{}

// UsageInfo struct for UsageInfo
type UsageInfo struct {
	PromptTokens int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens int32 `json:"total_tokens"`
}

type _UsageInfo UsageInfo

// NewUsageInfo instantiates a new UsageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageInfo(promptTokens int32, completionTokens int32, totalTokens int32) *UsageInfo {
	this := UsageInfo{}
	this.PromptTokens = promptTokens
	this.CompletionTokens = completionTokens
	this.TotalTokens = totalTokens
	return &this
}

// NewUsageInfoWithDefaults instantiates a new UsageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageInfoWithDefaults() *UsageInfo {
	this := UsageInfo{}
	return &this
}

// GetPromptTokens returns the PromptTokens field value
func (o *UsageInfo) GetPromptTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptTokens
}

// GetPromptTokensOk returns a tuple with the PromptTokens field value
// and a boolean to check if the value has been set.
func (o *UsageInfo) GetPromptTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptTokens, true
}

// SetPromptTokens sets field value
func (o *UsageInfo) SetPromptTokens(v int32) {
	o.PromptTokens = v
}

// GetCompletionTokens returns the CompletionTokens field value
func (o *UsageInfo) GetCompletionTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompletionTokens
}

// GetCompletionTokensOk returns a tuple with the CompletionTokens field value
// and a boolean to check if the value has been set.
func (o *UsageInfo) GetCompletionTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionTokens, true
}

// SetCompletionTokens sets field value
func (o *UsageInfo) SetCompletionTokens(v int32) {
	o.CompletionTokens = v
}

// GetTotalTokens returns the TotalTokens field value
func (o *UsageInfo) GetTotalTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *UsageInfo) GetTotalTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value
func (o *UsageInfo) SetTotalTokens(v int32) {
	o.TotalTokens = v
}

func (o UsageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt_tokens"] = o.PromptTokens
	toSerialize["completion_tokens"] = o.CompletionTokens
	toSerialize["total_tokens"] = o.TotalTokens
	return toSerialize, nil
}

func (o *UsageInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt_tokens",
		"completion_tokens",
		"total_tokens",
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

	varUsageInfo := _UsageInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageInfo)

	if err != nil {
		return err
	}

	*o = UsageInfo(varUsageInfo)

	return err
}

type NullableUsageInfo struct {
	value *UsageInfo
	isSet bool
}

func (v NullableUsageInfo) Get() *UsageInfo {
	return v.value
}

func (v *NullableUsageInfo) Set(val *UsageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageInfo(val *UsageInfo) *NullableUsageInfo {
	return &NullableUsageInfo{value: val, isSet: true}
}

func (v NullableUsageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


