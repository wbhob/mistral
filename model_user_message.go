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

// checks if the UserMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMessage{}

// UserMessage struct for UserMessage
type UserMessage struct {
	Content NullableContent3 `json:"content"`
	Role *string `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserMessage UserMessage

// NewUserMessage instantiates a new UserMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMessage(content NullableContent3) *UserMessage {
	this := UserMessage{}
	this.Content = content
	var role string = "user"
	this.Role = &role
	return &this
}

// NewUserMessageWithDefaults instantiates a new UserMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMessageWithDefaults() *UserMessage {
	this := UserMessage{}
	var role string = "user"
	this.Role = &role
	return &this
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for Content3 will be returned
func (o *UserMessage) GetContent() Content3 {
	if o == nil || o.Content.Get() == nil {
		var ret Content3
		return ret
	}

	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserMessage) GetContentOk() (*Content3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// SetContent sets field value
func (o *UserMessage) SetContent(v Content3) {
	o.Content.Set(&v)
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserMessage) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMessage) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserMessage) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserMessage) SetRole(v string) {
	o.Role = &v
}

func (o UserMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content.Get()
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
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

	varUserMessage := _UserMessage{}

	err = json.Unmarshal(data, &varUserMessage)

	if err != nil {
		return err
	}

	*o = UserMessage(varUserMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserMessage struct {
	value *UserMessage
	isSet bool
}

func (v NullableUserMessage) Get() *UserMessage {
	return v.value
}

func (v *NullableUserMessage) Set(val *UserMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMessage(val *UserMessage) *NullableUserMessage {
	return &NullableUserMessage{value: val, isSet: true}
}

func (v NullableUserMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


