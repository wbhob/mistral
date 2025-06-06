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

// checks if the ToolFileChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolFileChunk{}

// ToolFileChunk struct for ToolFileChunk
type ToolFileChunk struct {
	Type *string `json:"type,omitempty"`
	Tool BuiltInConnectors `json:"tool"`
	FileId string `json:"file_id"`
	FileName NullableString `json:"file_name,omitempty"`
	FileType NullableString `json:"file_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ToolFileChunk ToolFileChunk

// NewToolFileChunk instantiates a new ToolFileChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolFileChunk(tool BuiltInConnectors, fileId string) *ToolFileChunk {
	this := ToolFileChunk{}
	var type_ string = "tool_file"
	this.Type = &type_
	this.Tool = tool
	this.FileId = fileId
	return &this
}

// NewToolFileChunkWithDefaults instantiates a new ToolFileChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolFileChunkWithDefaults() *ToolFileChunk {
	this := ToolFileChunk{}
	var type_ string = "tool_file"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ToolFileChunk) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolFileChunk) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ToolFileChunk) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ToolFileChunk) SetType(v string) {
	o.Type = &v
}

// GetTool returns the Tool field value
func (o *ToolFileChunk) GetTool() BuiltInConnectors {
	if o == nil {
		var ret BuiltInConnectors
		return ret
	}

	return o.Tool
}

// GetToolOk returns a tuple with the Tool field value
// and a boolean to check if the value has been set.
func (o *ToolFileChunk) GetToolOk() (*BuiltInConnectors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tool, true
}

// SetTool sets field value
func (o *ToolFileChunk) SetTool(v BuiltInConnectors) {
	o.Tool = v
}

// GetFileId returns the FileId field value
func (o *ToolFileChunk) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *ToolFileChunk) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *ToolFileChunk) SetFileId(v string) {
	o.FileId = v
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolFileChunk) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolFileChunk) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *ToolFileChunk) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *ToolFileChunk) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *ToolFileChunk) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *ToolFileChunk) UnsetFileName() {
	o.FileName.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolFileChunk) GetFileType() string {
	if o == nil || IsNil(o.FileType.Get()) {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolFileChunk) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *ToolFileChunk) HasFileType() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *ToolFileChunk) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *ToolFileChunk) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *ToolFileChunk) UnsetFileType() {
	o.FileType.Unset()
}

func (o ToolFileChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolFileChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["tool"] = o.Tool
	toSerialize["file_id"] = o.FileId
	if o.FileName.IsSet() {
		toSerialize["file_name"] = o.FileName.Get()
	}
	if o.FileType.IsSet() {
		toSerialize["file_type"] = o.FileType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ToolFileChunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tool",
		"file_id",
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

	varToolFileChunk := _ToolFileChunk{}

	err = json.Unmarshal(data, &varToolFileChunk)

	if err != nil {
		return err
	}

	*o = ToolFileChunk(varToolFileChunk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "tool")
		delete(additionalProperties, "file_id")
		delete(additionalProperties, "file_name")
		delete(additionalProperties, "file_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableToolFileChunk struct {
	value *ToolFileChunk
	isSet bool
}

func (v NullableToolFileChunk) Get() *ToolFileChunk {
	return v.value
}

func (v *NullableToolFileChunk) Set(val *ToolFileChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableToolFileChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableToolFileChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolFileChunk(val *ToolFileChunk) *NullableToolFileChunk {
	return &NullableToolFileChunk{value: val, isSet: true}
}

func (v NullableToolFileChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolFileChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


