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

// checks if the ChatClassificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatClassificationRequest{}

// ChatClassificationRequest struct for ChatClassificationRequest
type ChatClassificationRequest struct {
	Model string `json:"model"`
	Input ChatClassificationRequestInputs `json:"input"`
	AdditionalProperties map[string]interface{}
}

type _ChatClassificationRequest ChatClassificationRequest

// NewChatClassificationRequest instantiates a new ChatClassificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatClassificationRequest(model string, input ChatClassificationRequestInputs) *ChatClassificationRequest {
	this := ChatClassificationRequest{}
	this.Model = model
	this.Input = input
	return &this
}

// NewChatClassificationRequestWithDefaults instantiates a new ChatClassificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatClassificationRequestWithDefaults() *ChatClassificationRequest {
	this := ChatClassificationRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *ChatClassificationRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ChatClassificationRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ChatClassificationRequest) SetModel(v string) {
	o.Model = v
}

// GetInput returns the Input field value
func (o *ChatClassificationRequest) GetInput() ChatClassificationRequestInputs {
	if o == nil {
		var ret ChatClassificationRequestInputs
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *ChatClassificationRequest) GetInputOk() (*ChatClassificationRequestInputs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *ChatClassificationRequest) SetInput(v ChatClassificationRequestInputs) {
	o.Input = v
}

func (o ChatClassificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatClassificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["input"] = o.Input

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatClassificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"input",
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

	varChatClassificationRequest := _ChatClassificationRequest{}

	err = json.Unmarshal(data, &varChatClassificationRequest)

	if err != nil {
		return err
	}

	*o = ChatClassificationRequest(varChatClassificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "model")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatClassificationRequest struct {
	value *ChatClassificationRequest
	isSet bool
}

func (v NullableChatClassificationRequest) Get() *ChatClassificationRequest {
	return v.value
}

func (v *NullableChatClassificationRequest) Set(val *ChatClassificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatClassificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatClassificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatClassificationRequest(val *ChatClassificationRequest) *NullableChatClassificationRequest {
	return &NullableChatClassificationRequest{value: val, isSet: true}
}

func (v NullableChatClassificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatClassificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


