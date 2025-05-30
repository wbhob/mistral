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

// checks if the ResponseBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBase{}

// ResponseBase struct for ResponseBase
type ResponseBase struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Model *string `json:"model,omitempty"`
	Usage *UsageInfo `json:"usage,omitempty"`
}

// NewResponseBase instantiates a new ResponseBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBase() *ResponseBase {
	this := ResponseBase{}
	return &this
}

// NewResponseBaseWithDefaults instantiates a new ResponseBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBaseWithDefaults() *ResponseBase {
	this := ResponseBase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseBase) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseBase) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResponseBase) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ResponseBase) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBase) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ResponseBase) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ResponseBase) SetObject(v string) {
	o.Object = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ResponseBase) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBase) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ResponseBase) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ResponseBase) SetModel(v string) {
	o.Model = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ResponseBase) GetUsage() UsageInfo {
	if o == nil || IsNil(o.Usage) {
		var ret UsageInfo
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBase) GetUsageOk() (*UsageInfo, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ResponseBase) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given UsageInfo and assigns it to the Usage field.
func (o *ResponseBase) SetUsage(v UsageInfo) {
	o.Usage = &v
}

func (o ResponseBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableResponseBase struct {
	value *ResponseBase
	isSet bool
}

func (v NullableResponseBase) Get() *ResponseBase {
	return v.value
}

func (v *NullableResponseBase) Set(val *ResponseBase) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBase) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBase(val *ResponseBase) *NullableResponseBase {
	return &NullableResponseBase{value: val, isSet: true}
}

func (v NullableResponseBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


