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

// checks if the Agent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agent{}

// Agent struct for Agent
type Agent struct {
	Instructions NullableString `json:"instructions,omitempty"`
	// List of tools which are available to the model during the conversation.
	Tools []AgentToolsInner `json:"tools,omitempty"`
	// Completion arguments that will be used to generate assistant responses. Can be overridden at each message request.
	CompletionArgs *CompletionArgs `json:"completion_args,omitempty"`
	Model string `json:"model"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Handoffs []string `json:"handoffs,omitempty"`
	Object *string `json:"object,omitempty"`
	Id string `json:"id"`
	Version int32 `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Agent Agent

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent(model string, name string, id string, version int32, createdAt time.Time, updatedAt time.Time) *Agent {
	this := Agent{}
	this.Model = model
	this.Name = name
	var object string = "agent"
	this.Object = &object
	this.Id = id
	this.Version = version
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	var object string = "agent"
	this.Object = &object
	return &this
}

// GetInstructions returns the Instructions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetInstructions() string {
	if o == nil || IsNil(o.Instructions.Get()) {
		var ret string
		return ret
	}
	return *o.Instructions.Get()
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instructions.Get(), o.Instructions.IsSet()
}

// HasInstructions returns a boolean if a field has been set.
func (o *Agent) HasInstructions() bool {
	if o != nil && o.Instructions.IsSet() {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given NullableString and assigns it to the Instructions field.
func (o *Agent) SetInstructions(v string) {
	o.Instructions.Set(&v)
}
// SetInstructionsNil sets the value for Instructions to be an explicit nil
func (o *Agent) SetInstructionsNil() {
	o.Instructions.Set(nil)
}

// UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
func (o *Agent) UnsetInstructions() {
	o.Instructions.Unset()
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *Agent) GetTools() []AgentToolsInner {
	if o == nil || IsNil(o.Tools) {
		var ret []AgentToolsInner
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetToolsOk() ([]AgentToolsInner, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *Agent) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []AgentToolsInner and assigns it to the Tools field.
func (o *Agent) SetTools(v []AgentToolsInner) {
	o.Tools = v
}

// GetCompletionArgs returns the CompletionArgs field value if set, zero value otherwise.
func (o *Agent) GetCompletionArgs() CompletionArgs {
	if o == nil || IsNil(o.CompletionArgs) {
		var ret CompletionArgs
		return ret
	}
	return *o.CompletionArgs
}

// GetCompletionArgsOk returns a tuple with the CompletionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetCompletionArgsOk() (*CompletionArgs, bool) {
	if o == nil || IsNil(o.CompletionArgs) {
		return nil, false
	}
	return o.CompletionArgs, true
}

// HasCompletionArgs returns a boolean if a field has been set.
func (o *Agent) HasCompletionArgs() bool {
	if o != nil && !IsNil(o.CompletionArgs) {
		return true
	}

	return false
}

// SetCompletionArgs gets a reference to the given CompletionArgs and assigns it to the CompletionArgs field.
func (o *Agent) SetCompletionArgs(v CompletionArgs) {
	o.CompletionArgs = &v
}

// GetModel returns the Model field value
func (o *Agent) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Agent) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Agent) SetModel(v string) {
	o.Model = v
}

// GetName returns the Name field value
func (o *Agent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Agent) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Agent) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Agent) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Agent) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Agent) UnsetDescription() {
	o.Description.Unset()
}

// GetHandoffs returns the Handoffs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetHandoffs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Handoffs
}

// GetHandoffsOk returns a tuple with the Handoffs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetHandoffsOk() ([]string, bool) {
	if o == nil || IsNil(o.Handoffs) {
		return nil, false
	}
	return o.Handoffs, true
}

// HasHandoffs returns a boolean if a field has been set.
func (o *Agent) HasHandoffs() bool {
	if o != nil && !IsNil(o.Handoffs) {
		return true
	}

	return false
}

// SetHandoffs gets a reference to the given []string and assigns it to the Handoffs field.
func (o *Agent) SetHandoffs(v []string) {
	o.Handoffs = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Agent) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Agent) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Agent) SetObject(v string) {
	o.Object = &v
}

// GetId returns the Id field value
func (o *Agent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Agent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Agent) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *Agent) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Agent) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Agent) SetVersion(v int32) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Agent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Agent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Agent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Agent) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Agent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Agent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Instructions.IsSet() {
		toSerialize["instructions"] = o.Instructions.Get()
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.CompletionArgs) {
		toSerialize["completion_args"] = o.CompletionArgs
	}
	toSerialize["model"] = o.Model
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Handoffs != nil {
		toSerialize["handoffs"] = o.Handoffs
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["id"] = o.Id
	toSerialize["version"] = o.Version
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Agent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"name",
		"id",
		"version",
		"created_at",
		"updated_at",
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

	varAgent := _Agent{}

	err = json.Unmarshal(data, &varAgent)

	if err != nil {
		return err
	}

	*o = Agent(varAgent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instructions")
		delete(additionalProperties, "tools")
		delete(additionalProperties, "completion_args")
		delete(additionalProperties, "model")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "handoffs")
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


