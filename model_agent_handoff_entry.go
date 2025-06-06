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

// checks if the AgentHandoffEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentHandoffEntry{}

// AgentHandoffEntry struct for AgentHandoffEntry
type AgentHandoffEntry struct {
	Object *string `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	Id *string `json:"id,omitempty"`
	PreviousAgentId string `json:"previous_agent_id"`
	PreviousAgentName string `json:"previous_agent_name"`
	NextAgentId string `json:"next_agent_id"`
	NextAgentName string `json:"next_agent_name"`
	AdditionalProperties map[string]interface{}
}

type _AgentHandoffEntry AgentHandoffEntry

// NewAgentHandoffEntry instantiates a new AgentHandoffEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentHandoffEntry(previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string) *AgentHandoffEntry {
	this := AgentHandoffEntry{}
	var object string = "entry"
	this.Object = &object
	var type_ string = "agent.handoff"
	this.Type = &type_
	this.PreviousAgentId = previousAgentId
	this.PreviousAgentName = previousAgentName
	this.NextAgentId = nextAgentId
	this.NextAgentName = nextAgentName
	return &this
}

// NewAgentHandoffEntryWithDefaults instantiates a new AgentHandoffEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentHandoffEntryWithDefaults() *AgentHandoffEntry {
	this := AgentHandoffEntry{}
	var object string = "entry"
	this.Object = &object
	var type_ string = "agent.handoff"
	this.Type = &type_
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AgentHandoffEntry) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AgentHandoffEntry) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *AgentHandoffEntry) SetObject(v string) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AgentHandoffEntry) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentHandoffEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentHandoffEntry) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentHandoffEntry) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentHandoffEntry) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AgentHandoffEntry) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentHandoffEntry) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentHandoffEntry) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *AgentHandoffEntry) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *AgentHandoffEntry) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *AgentHandoffEntry) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *AgentHandoffEntry) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentHandoffEntry) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentHandoffEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentHandoffEntry) SetId(v string) {
	o.Id = &v
}

// GetPreviousAgentId returns the PreviousAgentId field value
func (o *AgentHandoffEntry) GetPreviousAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousAgentId
}

// GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field value
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetPreviousAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousAgentId, true
}

// SetPreviousAgentId sets field value
func (o *AgentHandoffEntry) SetPreviousAgentId(v string) {
	o.PreviousAgentId = v
}

// GetPreviousAgentName returns the PreviousAgentName field value
func (o *AgentHandoffEntry) GetPreviousAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousAgentName
}

// GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field value
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetPreviousAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousAgentName, true
}

// SetPreviousAgentName sets field value
func (o *AgentHandoffEntry) SetPreviousAgentName(v string) {
	o.PreviousAgentName = v
}

// GetNextAgentId returns the NextAgentId field value
func (o *AgentHandoffEntry) GetNextAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextAgentId
}

// GetNextAgentIdOk returns a tuple with the NextAgentId field value
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetNextAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextAgentId, true
}

// SetNextAgentId sets field value
func (o *AgentHandoffEntry) SetNextAgentId(v string) {
	o.NextAgentId = v
}

// GetNextAgentName returns the NextAgentName field value
func (o *AgentHandoffEntry) GetNextAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextAgentName
}

// GetNextAgentNameOk returns a tuple with the NextAgentName field value
// and a boolean to check if the value has been set.
func (o *AgentHandoffEntry) GetNextAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextAgentName, true
}

// SetNextAgentName sets field value
func (o *AgentHandoffEntry) SetNextAgentName(v string) {
	o.NextAgentName = v
}

func (o AgentHandoffEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentHandoffEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["previous_agent_id"] = o.PreviousAgentId
	toSerialize["previous_agent_name"] = o.PreviousAgentName
	toSerialize["next_agent_id"] = o.NextAgentId
	toSerialize["next_agent_name"] = o.NextAgentName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentHandoffEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"previous_agent_id",
		"previous_agent_name",
		"next_agent_id",
		"next_agent_name",
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

	varAgentHandoffEntry := _AgentHandoffEntry{}

	err = json.Unmarshal(data, &varAgentHandoffEntry)

	if err != nil {
		return err
	}

	*o = AgentHandoffEntry(varAgentHandoffEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "previous_agent_id")
		delete(additionalProperties, "previous_agent_name")
		delete(additionalProperties, "next_agent_id")
		delete(additionalProperties, "next_agent_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentHandoffEntry struct {
	value *AgentHandoffEntry
	isSet bool
}

func (v NullableAgentHandoffEntry) Get() *AgentHandoffEntry {
	return v.value
}

func (v *NullableAgentHandoffEntry) Set(val *AgentHandoffEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentHandoffEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentHandoffEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentHandoffEntry(val *AgentHandoffEntry) *NullableAgentHandoffEntry {
	return &NullableAgentHandoffEntry{value: val, isSet: true}
}

func (v NullableAgentHandoffEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentHandoffEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


