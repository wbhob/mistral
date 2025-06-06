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

// checks if the DeleteModelOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteModelOut{}

// DeleteModelOut struct for DeleteModelOut
type DeleteModelOut struct {
	// The ID of the deleted model.
	Id string `json:"id"`
	// The object type that was deleted
	Object *string `json:"object,omitempty"`
	// The deletion status
	Deleted *bool `json:"deleted,omitempty"`
}

type _DeleteModelOut DeleteModelOut

// NewDeleteModelOut instantiates a new DeleteModelOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteModelOut(id string) *DeleteModelOut {
	this := DeleteModelOut{}
	this.Id = id
	var object string = "model"
	this.Object = &object
	var deleted bool = true
	this.Deleted = &deleted
	return &this
}

// NewDeleteModelOutWithDefaults instantiates a new DeleteModelOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteModelOutWithDefaults() *DeleteModelOut {
	this := DeleteModelOut{}
	var object string = "model"
	this.Object = &object
	var deleted bool = true
	this.Deleted = &deleted
	return &this
}

// GetId returns the Id field value
func (o *DeleteModelOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteModelOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteModelOut) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DeleteModelOut) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteModelOut) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DeleteModelOut) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *DeleteModelOut) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DeleteModelOut) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteModelOut) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DeleteModelOut) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *DeleteModelOut) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o DeleteModelOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteModelOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

func (o *DeleteModelOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varDeleteModelOut := _DeleteModelOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteModelOut)

	if err != nil {
		return err
	}

	*o = DeleteModelOut(varDeleteModelOut)

	return err
}

type NullableDeleteModelOut struct {
	value *DeleteModelOut
	isSet bool
}

func (v NullableDeleteModelOut) Get() *DeleteModelOut {
	return v.value
}

func (v *NullableDeleteModelOut) Set(val *DeleteModelOut) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteModelOut) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteModelOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteModelOut(val *DeleteModelOut) *NullableDeleteModelOut {
	return &NullableDeleteModelOut{value: val, isSet: true}
}

func (v NullableDeleteModelOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteModelOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


