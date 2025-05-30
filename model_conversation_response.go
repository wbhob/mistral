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

// checks if the ConversationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversationResponse{}

// ConversationResponse The response after appending new entries to the conversation.
type ConversationResponse struct {
	Object *string `json:"object,omitempty"`
	ConversationId string `json:"conversation_id"`
	Outputs []ConversationResponseOutputsInner `json:"outputs"`
	Usage ConversationUsageInfo `json:"usage"`
	AdditionalProperties map[string]interface{}
}

type _ConversationResponse ConversationResponse

// NewConversationResponse instantiates a new ConversationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversationResponse(conversationId string, outputs []ConversationResponseOutputsInner, usage ConversationUsageInfo) *ConversationResponse {
	this := ConversationResponse{}
	var object string = "conversation.response"
	this.Object = &object
	this.ConversationId = conversationId
	this.Outputs = outputs
	this.Usage = usage
	return &this
}

// NewConversationResponseWithDefaults instantiates a new ConversationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversationResponseWithDefaults() *ConversationResponse {
	this := ConversationResponse{}
	var object string = "conversation.response"
	this.Object = &object
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ConversationResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ConversationResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ConversationResponse) SetObject(v string) {
	o.Object = &v
}

// GetConversationId returns the ConversationId field value
func (o *ConversationResponse) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *ConversationResponse) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *ConversationResponse) SetConversationId(v string) {
	o.ConversationId = v
}

// GetOutputs returns the Outputs field value
func (o *ConversationResponse) GetOutputs() []ConversationResponseOutputsInner {
	if o == nil {
		var ret []ConversationResponseOutputsInner
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *ConversationResponse) GetOutputsOk() ([]ConversationResponseOutputsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *ConversationResponse) SetOutputs(v []ConversationResponseOutputsInner) {
	o.Outputs = v
}

// GetUsage returns the Usage field value
func (o *ConversationResponse) GetUsage() ConversationUsageInfo {
	if o == nil {
		var ret ConversationUsageInfo
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *ConversationResponse) GetUsageOk() (*ConversationUsageInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *ConversationResponse) SetUsage(v ConversationUsageInfo) {
	o.Usage = v
}

func (o ConversationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["conversation_id"] = o.ConversationId
	toSerialize["outputs"] = o.Outputs
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConversationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversation_id",
		"outputs",
		"usage",
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

	varConversationResponse := _ConversationResponse{}

	err = json.Unmarshal(data, &varConversationResponse)

	if err != nil {
		return err
	}

	*o = ConversationResponse(varConversationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "conversation_id")
		delete(additionalProperties, "outputs")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConversationResponse struct {
	value *ConversationResponse
	isSet bool
}

func (v NullableConversationResponse) Get() *ConversationResponse {
	return v.value
}

func (v *NullableConversationResponse) Set(val *ConversationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConversationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConversationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversationResponse(val *ConversationResponse) *NullableConversationResponse {
	return &NullableConversationResponse{value: val, isSet: true}
}

func (v NullableConversationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


