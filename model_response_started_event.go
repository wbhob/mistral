/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ResponseStartedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStartedEvent{}

// ResponseStartedEvent struct for ResponseStartedEvent
type ResponseStartedEvent struct {
	Type *string `json:"type,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ConversationId string `json:"conversation_id"`
	AdditionalProperties map[string]interface{}
}

type _ResponseStartedEvent ResponseStartedEvent

// NewResponseStartedEvent instantiates a new ResponseStartedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStartedEvent(conversationId string) *ResponseStartedEvent {
	this := ResponseStartedEvent{}
	var type_ string = "conversation.response.started"
	this.Type = &type_
	this.ConversationId = conversationId
	return &this
}

// NewResponseStartedEventWithDefaults instantiates a new ResponseStartedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStartedEventWithDefaults() *ResponseStartedEvent {
	this := ResponseStartedEvent{}
	var type_ string = "conversation.response.started"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponseStartedEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStartedEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponseStartedEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResponseStartedEvent) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ResponseStartedEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStartedEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResponseStartedEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ResponseStartedEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetConversationId returns the ConversationId field value
func (o *ResponseStartedEvent) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *ResponseStartedEvent) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *ResponseStartedEvent) SetConversationId(v string) {
	o.ConversationId = v
}

func (o ResponseStartedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStartedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["conversation_id"] = o.ConversationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseStartedEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversation_id",
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

	varResponseStartedEvent := _ResponseStartedEvent{}

	err = json.Unmarshal(data, &varResponseStartedEvent)

	if err != nil {
		return err
	}

	*o = ResponseStartedEvent(varResponseStartedEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "conversation_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseStartedEvent struct {
	value *ResponseStartedEvent
	isSet bool
}

func (v NullableResponseStartedEvent) Get() *ResponseStartedEvent {
	return v.value
}

func (v *NullableResponseStartedEvent) Set(val *ResponseStartedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStartedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStartedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStartedEvent(val *ResponseStartedEvent) *NullableResponseStartedEvent {
	return &NullableResponseStartedEvent{value: val, isSet: true}
}

func (v NullableResponseStartedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStartedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


