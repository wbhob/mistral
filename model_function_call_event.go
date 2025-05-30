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

// checks if the FunctionCallEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCallEvent{}

// FunctionCallEvent struct for FunctionCallEvent
type FunctionCallEvent struct {
	Type *string `json:"type,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	OutputIndex *int32 `json:"output_index,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	ToolCallId string `json:"tool_call_id"`
	Arguments string `json:"arguments"`
	AdditionalProperties map[string]interface{}
}

type _FunctionCallEvent FunctionCallEvent

// NewFunctionCallEvent instantiates a new FunctionCallEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCallEvent(id string, name string, toolCallId string, arguments string) *FunctionCallEvent {
	this := FunctionCallEvent{}
	var type_ string = "function.call.delta"
	this.Type = &type_
	var outputIndex int32 = 0
	this.OutputIndex = &outputIndex
	this.Id = id
	this.Name = name
	this.ToolCallId = toolCallId
	this.Arguments = arguments
	return &this
}

// NewFunctionCallEventWithDefaults instantiates a new FunctionCallEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCallEventWithDefaults() *FunctionCallEvent {
	this := FunctionCallEvent{}
	var type_ string = "function.call.delta"
	this.Type = &type_
	var outputIndex int32 = 0
	this.OutputIndex = &outputIndex
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FunctionCallEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FunctionCallEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FunctionCallEvent) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FunctionCallEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FunctionCallEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FunctionCallEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetOutputIndex returns the OutputIndex field value if set, zero value otherwise.
func (o *FunctionCallEvent) GetOutputIndex() int32 {
	if o == nil || IsNil(o.OutputIndex) {
		var ret int32
		return ret
	}
	return *o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetOutputIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputIndex) {
		return nil, false
	}
	return o.OutputIndex, true
}

// HasOutputIndex returns a boolean if a field has been set.
func (o *FunctionCallEvent) HasOutputIndex() bool {
	if o != nil && !IsNil(o.OutputIndex) {
		return true
	}

	return false
}

// SetOutputIndex gets a reference to the given int32 and assigns it to the OutputIndex field.
func (o *FunctionCallEvent) SetOutputIndex(v int32) {
	o.OutputIndex = &v
}

// GetId returns the Id field value
func (o *FunctionCallEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FunctionCallEvent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FunctionCallEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionCallEvent) SetName(v string) {
	o.Name = v
}

// GetToolCallId returns the ToolCallId field value
func (o *FunctionCallEvent) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value
func (o *FunctionCallEvent) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetArguments returns the Arguments field value
func (o *FunctionCallEvent) GetArguments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *FunctionCallEvent) GetArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value
func (o *FunctionCallEvent) SetArguments(v string) {
	o.Arguments = v
}

func (o FunctionCallEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCallEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.OutputIndex) {
		toSerialize["output_index"] = o.OutputIndex
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["tool_call_id"] = o.ToolCallId
	toSerialize["arguments"] = o.Arguments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FunctionCallEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"tool_call_id",
		"arguments",
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

	varFunctionCallEvent := _FunctionCallEvent{}

	err = json.Unmarshal(data, &varFunctionCallEvent)

	if err != nil {
		return err
	}

	*o = FunctionCallEvent(varFunctionCallEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "output_index")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tool_call_id")
		delete(additionalProperties, "arguments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFunctionCallEvent struct {
	value *FunctionCallEvent
	isSet bool
}

func (v NullableFunctionCallEvent) Get() *FunctionCallEvent {
	return v.value
}

func (v *NullableFunctionCallEvent) Set(val *FunctionCallEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCallEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCallEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCallEvent(val *FunctionCallEvent) *NullableFunctionCallEvent {
	return &NullableFunctionCallEvent{value: val, isSet: true}
}

func (v NullableFunctionCallEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCallEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


