/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"fmt"
)

// checks if the OCRImageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRImageObject{}

// OCRImageObject struct for OCRImageObject
type OCRImageObject struct {
	// Image ID for extracted image in a page
	Id string `json:"id"`
	TopLeftX NullableInt32 `json:"top_left_x"`
	TopLeftY NullableInt32 `json:"top_left_y"`
	BottomRightX NullableInt32 `json:"bottom_right_x"`
	BottomRightY NullableInt32 `json:"bottom_right_y"`
	ImageBase64 NullableString `json:"image_base64,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OCRImageObject OCRImageObject

// NewOCRImageObject instantiates a new OCRImageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRImageObject(id string, topLeftX NullableInt32, topLeftY NullableInt32, bottomRightX NullableInt32, bottomRightY NullableInt32) *OCRImageObject {
	this := OCRImageObject{}
	this.Id = id
	this.TopLeftX = topLeftX
	this.TopLeftY = topLeftY
	this.BottomRightX = bottomRightX
	this.BottomRightY = bottomRightY
	return &this
}

// NewOCRImageObjectWithDefaults instantiates a new OCRImageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRImageObjectWithDefaults() *OCRImageObject {
	this := OCRImageObject{}
	return &this
}

// GetId returns the Id field value
func (o *OCRImageObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OCRImageObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OCRImageObject) SetId(v string) {
	o.Id = v
}

// GetTopLeftX returns the TopLeftX field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OCRImageObject) GetTopLeftX() int32 {
	if o == nil || o.TopLeftX.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TopLeftX.Get()
}

// GetTopLeftXOk returns a tuple with the TopLeftX field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRImageObject) GetTopLeftXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopLeftX.Get(), o.TopLeftX.IsSet()
}

// SetTopLeftX sets field value
func (o *OCRImageObject) SetTopLeftX(v int32) {
	o.TopLeftX.Set(&v)
}

// GetTopLeftY returns the TopLeftY field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OCRImageObject) GetTopLeftY() int32 {
	if o == nil || o.TopLeftY.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TopLeftY.Get()
}

// GetTopLeftYOk returns a tuple with the TopLeftY field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRImageObject) GetTopLeftYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopLeftY.Get(), o.TopLeftY.IsSet()
}

// SetTopLeftY sets field value
func (o *OCRImageObject) SetTopLeftY(v int32) {
	o.TopLeftY.Set(&v)
}

// GetBottomRightX returns the BottomRightX field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OCRImageObject) GetBottomRightX() int32 {
	if o == nil || o.BottomRightX.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BottomRightX.Get()
}

// GetBottomRightXOk returns a tuple with the BottomRightX field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRImageObject) GetBottomRightXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BottomRightX.Get(), o.BottomRightX.IsSet()
}

// SetBottomRightX sets field value
func (o *OCRImageObject) SetBottomRightX(v int32) {
	o.BottomRightX.Set(&v)
}

// GetBottomRightY returns the BottomRightY field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OCRImageObject) GetBottomRightY() int32 {
	if o == nil || o.BottomRightY.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BottomRightY.Get()
}

// GetBottomRightYOk returns a tuple with the BottomRightY field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRImageObject) GetBottomRightYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BottomRightY.Get(), o.BottomRightY.IsSet()
}

// SetBottomRightY sets field value
func (o *OCRImageObject) SetBottomRightY(v int32) {
	o.BottomRightY.Set(&v)
}

// GetImageBase64 returns the ImageBase64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRImageObject) GetImageBase64() string {
	if o == nil || IsNil(o.ImageBase64.Get()) {
		var ret string
		return ret
	}
	return *o.ImageBase64.Get()
}

// GetImageBase64Ok returns a tuple with the ImageBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRImageObject) GetImageBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageBase64.Get(), o.ImageBase64.IsSet()
}

// HasImageBase64 returns a boolean if a field has been set.
func (o *OCRImageObject) HasImageBase64() bool {
	if o != nil && o.ImageBase64.IsSet() {
		return true
	}

	return false
}

// SetImageBase64 gets a reference to the given NullableString and assigns it to the ImageBase64 field.
func (o *OCRImageObject) SetImageBase64(v string) {
	o.ImageBase64.Set(&v)
}
// SetImageBase64Nil sets the value for ImageBase64 to be an explicit nil
func (o *OCRImageObject) SetImageBase64Nil() {
	o.ImageBase64.Set(nil)
}

// UnsetImageBase64 ensures that no value is present for ImageBase64, not even an explicit nil
func (o *OCRImageObject) UnsetImageBase64() {
	o.ImageBase64.Unset()
}

func (o OCRImageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRImageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["top_left_x"] = o.TopLeftX.Get()
	toSerialize["top_left_y"] = o.TopLeftY.Get()
	toSerialize["bottom_right_x"] = o.BottomRightX.Get()
	toSerialize["bottom_right_y"] = o.BottomRightY.Get()
	if o.ImageBase64.IsSet() {
		toSerialize["image_base64"] = o.ImageBase64.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OCRImageObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"top_left_x",
		"top_left_y",
		"bottom_right_x",
		"bottom_right_y",
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

	varOCRImageObject := _OCRImageObject{}

	err = json.Unmarshal(data, &varOCRImageObject)

	if err != nil {
		return err
	}

	*o = OCRImageObject(varOCRImageObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "top_left_x")
		delete(additionalProperties, "top_left_y")
		delete(additionalProperties, "bottom_right_x")
		delete(additionalProperties, "bottom_right_y")
		delete(additionalProperties, "image_base64")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOCRImageObject struct {
	value *OCRImageObject
	isSet bool
}

func (v NullableOCRImageObject) Get() *OCRImageObject {
	return v.value
}

func (v *NullableOCRImageObject) Set(val *OCRImageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRImageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRImageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRImageObject(val *OCRImageObject) *NullableOCRImageObject {
	return &NullableOCRImageObject{value: val, isSet: true}
}

func (v NullableOCRImageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRImageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


