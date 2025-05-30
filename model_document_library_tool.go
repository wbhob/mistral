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

// checks if the DocumentLibraryTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentLibraryTool{}

// DocumentLibraryTool struct for DocumentLibraryTool
type DocumentLibraryTool struct {
	Type *string `json:"type,omitempty"`
	// Ids of the library in which to search.
	LibraryIds []string `json:"library_ids"`
	AdditionalProperties map[string]interface{}
}

type _DocumentLibraryTool DocumentLibraryTool

// NewDocumentLibraryTool instantiates a new DocumentLibraryTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentLibraryTool(libraryIds []string) *DocumentLibraryTool {
	this := DocumentLibraryTool{}
	var type_ string = "document_library"
	this.Type = &type_
	this.LibraryIds = libraryIds
	return &this
}

// NewDocumentLibraryToolWithDefaults instantiates a new DocumentLibraryTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentLibraryToolWithDefaults() *DocumentLibraryTool {
	this := DocumentLibraryTool{}
	var type_ string = "document_library"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DocumentLibraryTool) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentLibraryTool) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DocumentLibraryTool) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DocumentLibraryTool) SetType(v string) {
	o.Type = &v
}

// GetLibraryIds returns the LibraryIds field value
func (o *DocumentLibraryTool) GetLibraryIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LibraryIds
}

// GetLibraryIdsOk returns a tuple with the LibraryIds field value
// and a boolean to check if the value has been set.
func (o *DocumentLibraryTool) GetLibraryIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibraryIds, true
}

// SetLibraryIds sets field value
func (o *DocumentLibraryTool) SetLibraryIds(v []string) {
	o.LibraryIds = v
}

func (o DocumentLibraryTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentLibraryTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["library_ids"] = o.LibraryIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentLibraryTool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"library_ids",
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

	varDocumentLibraryTool := _DocumentLibraryTool{}

	err = json.Unmarshal(data, &varDocumentLibraryTool)

	if err != nil {
		return err
	}

	*o = DocumentLibraryTool(varDocumentLibraryTool)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "library_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentLibraryTool struct {
	value *DocumentLibraryTool
	isSet bool
}

func (v NullableDocumentLibraryTool) Get() *DocumentLibraryTool {
	return v.value
}

func (v *NullableDocumentLibraryTool) Set(val *DocumentLibraryTool) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentLibraryTool) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentLibraryTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentLibraryTool(val *DocumentLibraryTool) *NullableDocumentLibraryTool {
	return &NullableDocumentLibraryTool{value: val, isSet: true}
}

func (v NullableDocumentLibraryTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentLibraryTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


