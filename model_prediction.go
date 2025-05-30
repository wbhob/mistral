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

// checks if the Prediction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prediction{}

// Prediction struct for Prediction
type Prediction struct {
	Type *string `json:"type,omitempty"`
	Content *string `json:"content,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Prediction Prediction

// NewPrediction instantiates a new Prediction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrediction() *Prediction {
	this := Prediction{}
	var type_ string = "content"
	this.Type = &type_
	var content string = ""
	this.Content = &content
	return &this
}

// NewPredictionWithDefaults instantiates a new Prediction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionWithDefaults() *Prediction {
	this := Prediction{}
	var type_ string = "content"
	this.Type = &type_
	var content string = ""
	this.Content = &content
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Prediction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prediction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Prediction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Prediction) SetType(v string) {
	o.Type = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Prediction) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prediction) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Prediction) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Prediction) SetContent(v string) {
	o.Content = &v
}

func (o Prediction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prediction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Prediction) UnmarshalJSON(data []byte) (err error) {
	varPrediction := _Prediction{}

	err = json.Unmarshal(data, &varPrediction)

	if err != nil {
		return err
	}

	*o = Prediction(varPrediction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrediction struct {
	value *Prediction
	isSet bool
}

func (v NullablePrediction) Get() *Prediction {
	return v.value
}

func (v *NullablePrediction) Set(val *Prediction) {
	v.value = val
	v.isSet = true
}

func (v NullablePrediction) IsSet() bool {
	return v.isSet
}

func (v *NullablePrediction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrediction(val *Prediction) *NullablePrediction {
	return &NullablePrediction{value: val, isSet: true}
}

func (v NullablePrediction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrediction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


