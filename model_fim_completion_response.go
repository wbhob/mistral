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

// checks if the FIMCompletionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIMCompletionResponse{}

// FIMCompletionResponse struct for FIMCompletionResponse
type FIMCompletionResponse struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Model *string `json:"model,omitempty"`
	Usage *UsageInfo `json:"usage,omitempty"`
	Created *int32 `json:"created,omitempty"`
	Choices []ChatCompletionChoice `json:"choices"`
}

type _FIMCompletionResponse FIMCompletionResponse

// NewFIMCompletionResponse instantiates a new FIMCompletionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIMCompletionResponse(choices []ChatCompletionChoice) *FIMCompletionResponse {
	this := FIMCompletionResponse{}
	this.Choices = choices
	return &this
}

// NewFIMCompletionResponseWithDefaults instantiates a new FIMCompletionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIMCompletionResponseWithDefaults() *FIMCompletionResponse {
	this := FIMCompletionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FIMCompletionResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FIMCompletionResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FIMCompletionResponse) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *FIMCompletionResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *FIMCompletionResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *FIMCompletionResponse) SetObject(v string) {
	o.Object = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FIMCompletionResponse) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FIMCompletionResponse) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FIMCompletionResponse) SetModel(v string) {
	o.Model = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *FIMCompletionResponse) GetUsage() UsageInfo {
	if o == nil || IsNil(o.Usage) {
		var ret UsageInfo
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetUsageOk() (*UsageInfo, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *FIMCompletionResponse) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given UsageInfo and assigns it to the Usage field.
func (o *FIMCompletionResponse) SetUsage(v UsageInfo) {
	o.Usage = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FIMCompletionResponse) GetCreated() int32 {
	if o == nil || IsNil(o.Created) {
		var ret int32
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FIMCompletionResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int32 and assigns it to the Created field.
func (o *FIMCompletionResponse) SetCreated(v int32) {
	o.Created = &v
}

// GetChoices returns the Choices field value
func (o *FIMCompletionResponse) GetChoices() []ChatCompletionChoice {
	if o == nil {
		var ret []ChatCompletionChoice
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *FIMCompletionResponse) GetChoicesOk() ([]ChatCompletionChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *FIMCompletionResponse) SetChoices(v []ChatCompletionChoice) {
	o.Choices = v
}

func (o FIMCompletionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIMCompletionResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	toSerialize["choices"] = o.Choices
	return toSerialize, nil
}

func (o *FIMCompletionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"choices",
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

	varFIMCompletionResponse := _FIMCompletionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFIMCompletionResponse)

	if err != nil {
		return err
	}

	*o = FIMCompletionResponse(varFIMCompletionResponse)

	return err
}

type NullableFIMCompletionResponse struct {
	value *FIMCompletionResponse
	isSet bool
}

func (v NullableFIMCompletionResponse) Get() *FIMCompletionResponse {
	return v.value
}

func (v *NullableFIMCompletionResponse) Set(val *FIMCompletionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFIMCompletionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFIMCompletionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIMCompletionResponse(val *FIMCompletionResponse) *NullableFIMCompletionResponse {
	return &NullableFIMCompletionResponse{value: val, isSet: true}
}

func (v NullableFIMCompletionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIMCompletionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


