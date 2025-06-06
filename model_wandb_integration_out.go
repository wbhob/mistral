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

// checks if the WandbIntegrationOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WandbIntegrationOut{}

// WandbIntegrationOut struct for WandbIntegrationOut
type WandbIntegrationOut struct {
	Type *string `json:"type,omitempty"`
	// The name of the project that the new run will be created under.
	Project string `json:"project"`
	Name NullableString `json:"name,omitempty"`
	RunName NullableString `json:"run_name,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

type _WandbIntegrationOut WandbIntegrationOut

// NewWandbIntegrationOut instantiates a new WandbIntegrationOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWandbIntegrationOut(project string) *WandbIntegrationOut {
	this := WandbIntegrationOut{}
	var type_ string = "wandb"
	this.Type = &type_
	this.Project = project
	return &this
}

// NewWandbIntegrationOutWithDefaults instantiates a new WandbIntegrationOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWandbIntegrationOutWithDefaults() *WandbIntegrationOut {
	this := WandbIntegrationOut{}
	var type_ string = "wandb"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WandbIntegrationOut) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WandbIntegrationOut) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WandbIntegrationOut) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WandbIntegrationOut) SetType(v string) {
	o.Type = &v
}

// GetProject returns the Project field value
func (o *WandbIntegrationOut) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *WandbIntegrationOut) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *WandbIntegrationOut) SetProject(v string) {
	o.Project = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WandbIntegrationOut) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WandbIntegrationOut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WandbIntegrationOut) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WandbIntegrationOut) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WandbIntegrationOut) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WandbIntegrationOut) UnsetName() {
	o.Name.Unset()
}

// GetRunName returns the RunName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WandbIntegrationOut) GetRunName() string {
	if o == nil || IsNil(o.RunName.Get()) {
		var ret string
		return ret
	}
	return *o.RunName.Get()
}

// GetRunNameOk returns a tuple with the RunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WandbIntegrationOut) GetRunNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunName.Get(), o.RunName.IsSet()
}

// HasRunName returns a boolean if a field has been set.
func (o *WandbIntegrationOut) HasRunName() bool {
	if o != nil && o.RunName.IsSet() {
		return true
	}

	return false
}

// SetRunName gets a reference to the given NullableString and assigns it to the RunName field.
func (o *WandbIntegrationOut) SetRunName(v string) {
	o.RunName.Set(&v)
}
// SetRunNameNil sets the value for RunName to be an explicit nil
func (o *WandbIntegrationOut) SetRunNameNil() {
	o.RunName.Set(nil)
}

// UnsetRunName ensures that no value is present for RunName, not even an explicit nil
func (o *WandbIntegrationOut) UnsetRunName() {
	o.RunName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WandbIntegrationOut) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WandbIntegrationOut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *WandbIntegrationOut) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *WandbIntegrationOut) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *WandbIntegrationOut) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *WandbIntegrationOut) UnsetUrl() {
	o.Url.Unset()
}

func (o WandbIntegrationOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WandbIntegrationOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["project"] = o.Project
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.RunName.IsSet() {
		toSerialize["run_name"] = o.RunName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

func (o *WandbIntegrationOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project",
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

	varWandbIntegrationOut := _WandbIntegrationOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWandbIntegrationOut)

	if err != nil {
		return err
	}

	*o = WandbIntegrationOut(varWandbIntegrationOut)

	return err
}

type NullableWandbIntegrationOut struct {
	value *WandbIntegrationOut
	isSet bool
}

func (v NullableWandbIntegrationOut) Get() *WandbIntegrationOut {
	return v.value
}

func (v *NullableWandbIntegrationOut) Set(val *WandbIntegrationOut) {
	v.value = val
	v.isSet = true
}

func (v NullableWandbIntegrationOut) IsSet() bool {
	return v.isSet
}

func (v *NullableWandbIntegrationOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWandbIntegrationOut(val *WandbIntegrationOut) *NullableWandbIntegrationOut {
	return &NullableWandbIntegrationOut{value: val, isSet: true}
}

func (v NullableWandbIntegrationOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWandbIntegrationOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


