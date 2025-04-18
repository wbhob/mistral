/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
)

// checks if the ClassificationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassificationObject{}

// ClassificationObject struct for ClassificationObject
type ClassificationObject struct {
	// Classifier result thresholded
	Categories map[string]bool `json:"categories,omitempty"`
	// Classifier result
	CategoryScores map[string]float32 `json:"category_scores,omitempty"`
}

// NewClassificationObject instantiates a new ClassificationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassificationObject() *ClassificationObject {
	this := ClassificationObject{}
	return &this
}

// NewClassificationObjectWithDefaults instantiates a new ClassificationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassificationObjectWithDefaults() *ClassificationObject {
	this := ClassificationObject{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ClassificationObject) GetCategories() map[string]bool {
	if o == nil || IsNil(o.Categories) {
		var ret map[string]bool
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationObject) GetCategoriesOk() (map[string]bool, bool) {
	if o == nil || IsNil(o.Categories) {
		return map[string]bool{}, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ClassificationObject) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given map[string]bool and assigns it to the Categories field.
func (o *ClassificationObject) SetCategories(v map[string]bool) {
	o.Categories = v
}

// GetCategoryScores returns the CategoryScores field value if set, zero value otherwise.
func (o *ClassificationObject) GetCategoryScores() map[string]float32 {
	if o == nil || IsNil(o.CategoryScores) {
		var ret map[string]float32
		return ret
	}
	return o.CategoryScores
}

// GetCategoryScoresOk returns a tuple with the CategoryScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassificationObject) GetCategoryScoresOk() (map[string]float32, bool) {
	if o == nil || IsNil(o.CategoryScores) {
		return map[string]float32{}, false
	}
	return o.CategoryScores, true
}

// HasCategoryScores returns a boolean if a field has been set.
func (o *ClassificationObject) HasCategoryScores() bool {
	if o != nil && !IsNil(o.CategoryScores) {
		return true
	}

	return false
}

// SetCategoryScores gets a reference to the given map[string]float32 and assigns it to the CategoryScores field.
func (o *ClassificationObject) SetCategoryScores(v map[string]float32) {
	o.CategoryScores = v
}

func (o ClassificationObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassificationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.CategoryScores) {
		toSerialize["category_scores"] = o.CategoryScores
	}
	return toSerialize, nil
}

type NullableClassificationObject struct {
	value *ClassificationObject
	isSet bool
}

func (v NullableClassificationObject) Get() *ClassificationObject {
	return v.value
}

func (v *NullableClassificationObject) Set(val *ClassificationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableClassificationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableClassificationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassificationObject(val *ClassificationObject) *NullableClassificationObject {
	return &NullableClassificationObject{value: val, isSet: true}
}

func (v NullableClassificationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassificationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


