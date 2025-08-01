/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
)

// checks if the DocumentUpdateIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentUpdateIn{}

// DocumentUpdateIn struct for DocumentUpdateIn
type DocumentUpdateIn struct {
	Name NullableString `json:"name,omitempty"`
}

// NewDocumentUpdateIn instantiates a new DocumentUpdateIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentUpdateIn() *DocumentUpdateIn {
	this := DocumentUpdateIn{}
	return &this
}

// NewDocumentUpdateInWithDefaults instantiates a new DocumentUpdateIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentUpdateInWithDefaults() *DocumentUpdateIn {
	this := DocumentUpdateIn{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentUpdateIn) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentUpdateIn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DocumentUpdateIn) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DocumentUpdateIn) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DocumentUpdateIn) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DocumentUpdateIn) UnsetName() {
	o.Name.Unset()
}

func (o DocumentUpdateIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentUpdateIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableDocumentUpdateIn struct {
	value *DocumentUpdateIn
	isSet bool
}

func (v NullableDocumentUpdateIn) Get() *DocumentUpdateIn {
	return v.value
}

func (v *NullableDocumentUpdateIn) Set(val *DocumentUpdateIn) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentUpdateIn) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentUpdateIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentUpdateIn(val *DocumentUpdateIn) *NullableDocumentUpdateIn {
	return &NullableDocumentUpdateIn{value: val, isSet: true}
}

func (v NullableDocumentUpdateIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentUpdateIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


