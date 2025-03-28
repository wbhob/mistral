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

// checks if the ChatCompletionChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionChoice{}

// ChatCompletionChoice struct for ChatCompletionChoice
type ChatCompletionChoice struct {
	Index int32 `json:"index"`
	Message AssistantMessage `json:"message"`
	FinishReason string `json:"finish_reason"`
}

type _ChatCompletionChoice ChatCompletionChoice

// NewChatCompletionChoice instantiates a new ChatCompletionChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionChoice(index int32, message AssistantMessage, finishReason string) *ChatCompletionChoice {
	this := ChatCompletionChoice{}
	this.Index = index
	this.Message = message
	this.FinishReason = finishReason
	return &this
}

// NewChatCompletionChoiceWithDefaults instantiates a new ChatCompletionChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionChoiceWithDefaults() *ChatCompletionChoice {
	this := ChatCompletionChoice{}
	return &this
}

// GetIndex returns the Index field value
func (o *ChatCompletionChoice) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionChoice) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ChatCompletionChoice) SetIndex(v int32) {
	o.Index = v
}

// GetMessage returns the Message field value
func (o *ChatCompletionChoice) GetMessage() AssistantMessage {
	if o == nil {
		var ret AssistantMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionChoice) GetMessageOk() (*AssistantMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ChatCompletionChoice) SetMessage(v AssistantMessage) {
	o.Message = v
}

// GetFinishReason returns the FinishReason field value
func (o *ChatCompletionChoice) GetFinishReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishReason
}

// GetFinishReasonOk returns a tuple with the FinishReason field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionChoice) GetFinishReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishReason, true
}

// SetFinishReason sets field value
func (o *ChatCompletionChoice) SetFinishReason(v string) {
	o.FinishReason = v
}

func (o ChatCompletionChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["message"] = o.Message
	toSerialize["finish_reason"] = o.FinishReason
	return toSerialize, nil
}

func (o *ChatCompletionChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"message",
		"finish_reason",
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

	varChatCompletionChoice := _ChatCompletionChoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionChoice)

	if err != nil {
		return err
	}

	*o = ChatCompletionChoice(varChatCompletionChoice)

	return err
}

type NullableChatCompletionChoice struct {
	value *ChatCompletionChoice
	isSet bool
}

func (v NullableChatCompletionChoice) Get() *ChatCompletionChoice {
	return v.value
}

func (v *NullableChatCompletionChoice) Set(val *ChatCompletionChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionChoice(val *ChatCompletionChoice) *NullableChatCompletionChoice {
	return &NullableChatCompletionChoice{value: val, isSet: true}
}

func (v NullableChatCompletionChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


