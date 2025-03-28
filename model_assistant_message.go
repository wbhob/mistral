/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
)

// checks if the AssistantMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantMessage{}

// AssistantMessage struct for AssistantMessage
type AssistantMessage struct {
	Content NullableContent `json:"content,omitempty"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// Set this to `true` when adding an assistant message as prefix to condition the model response. The role of the prefix message is to force the model to start its answer by the content of the message.
	Prefix *bool `json:"prefix,omitempty"`
	Role *string `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssistantMessage AssistantMessage

// NewAssistantMessage instantiates a new AssistantMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantMessage() *AssistantMessage {
	this := AssistantMessage{}
	var prefix bool = false
	this.Prefix = &prefix
	var role string = "assistant"
	this.Role = &role
	return &this
}

// NewAssistantMessageWithDefaults instantiates a new AssistantMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantMessageWithDefaults() *AssistantMessage {
	this := AssistantMessage{}
	var prefix bool = false
	this.Prefix = &prefix
	var role string = "assistant"
	this.Role = &role
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssistantMessage) GetContent() Content {
	if o == nil || IsNil(o.Content.Get()) {
		var ret Content
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssistantMessage) GetContentOk() (*Content, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *AssistantMessage) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableContent and assigns it to the Content field.
func (o *AssistantMessage) SetContent(v Content) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *AssistantMessage) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *AssistantMessage) UnsetContent() {
	o.Content.Unset()
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssistantMessage) GetToolCalls() []ToolCall {
	if o == nil {
		var ret []ToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssistantMessage) GetToolCallsOk() ([]ToolCall, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *AssistantMessage) HasToolCalls() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []ToolCall and assigns it to the ToolCalls field.
func (o *AssistantMessage) SetToolCalls(v []ToolCall) {
	o.ToolCalls = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *AssistantMessage) GetPrefix() bool {
	if o == nil || IsNil(o.Prefix) {
		var ret bool
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantMessage) GetPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *AssistantMessage) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given bool and assigns it to the Prefix field.
func (o *AssistantMessage) SetPrefix(v bool) {
	o.Prefix = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AssistantMessage) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantMessage) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AssistantMessage) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AssistantMessage) SetRole(v string) {
	o.Role = &v
}

func (o AssistantMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.ToolCalls != nil {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssistantMessage) UnmarshalJSON(data []byte) (err error) {
	varAssistantMessage := _AssistantMessage{}

	err = json.Unmarshal(data, &varAssistantMessage)

	if err != nil {
		return err
	}

	*o = AssistantMessage(varAssistantMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "tool_calls")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssistantMessage struct {
	value *AssistantMessage
	isSet bool
}

func (v NullableAssistantMessage) Get() *AssistantMessage {
	return v.value
}

func (v *NullableAssistantMessage) Set(val *AssistantMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantMessage(val *AssistantMessage) *NullableAssistantMessage {
	return &NullableAssistantMessage{value: val, isSet: true}
}

func (v NullableAssistantMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


