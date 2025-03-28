/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GithubRepositoryIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubRepositoryIn{}

// GithubRepositoryIn struct for GithubRepositoryIn
type GithubRepositoryIn struct {
	Type *string `json:"type,omitempty"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	Ref NullableString `json:"ref,omitempty"`
	Weight *float32 `json:"weight,omitempty"`
	Token string `json:"token"`
}

type _GithubRepositoryIn GithubRepositoryIn

// NewGithubRepositoryIn instantiates a new GithubRepositoryIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubRepositoryIn(name string, owner string, token string) *GithubRepositoryIn {
	this := GithubRepositoryIn{}
	var type_ string = "github"
	this.Type = &type_
	this.Name = name
	this.Owner = owner
	var weight float32 = 1.0
	this.Weight = &weight
	this.Token = token
	return &this
}

// NewGithubRepositoryInWithDefaults instantiates a new GithubRepositoryIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubRepositoryInWithDefaults() *GithubRepositoryIn {
	this := GithubRepositoryIn{}
	var type_ string = "github"
	this.Type = &type_
	var weight float32 = 1.0
	this.Weight = &weight
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GithubRepositoryIn) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubRepositoryIn) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GithubRepositoryIn) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GithubRepositoryIn) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *GithubRepositoryIn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GithubRepositoryIn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GithubRepositoryIn) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *GithubRepositoryIn) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *GithubRepositoryIn) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *GithubRepositoryIn) SetOwner(v string) {
	o.Owner = v
}

// GetRef returns the Ref field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubRepositoryIn) GetRef() string {
	if o == nil || IsNil(o.Ref.Get()) {
		var ret string
		return ret
	}
	return *o.Ref.Get()
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubRepositoryIn) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ref.Get(), o.Ref.IsSet()
}

// HasRef returns a boolean if a field has been set.
func (o *GithubRepositoryIn) HasRef() bool {
	if o != nil && o.Ref.IsSet() {
		return true
	}

	return false
}

// SetRef gets a reference to the given NullableString and assigns it to the Ref field.
func (o *GithubRepositoryIn) SetRef(v string) {
	o.Ref.Set(&v)
}
// SetRefNil sets the value for Ref to be an explicit nil
func (o *GithubRepositoryIn) SetRefNil() {
	o.Ref.Set(nil)
}

// UnsetRef ensures that no value is present for Ref, not even an explicit nil
func (o *GithubRepositoryIn) UnsetRef() {
	o.Ref.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GithubRepositoryIn) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubRepositoryIn) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GithubRepositoryIn) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *GithubRepositoryIn) SetWeight(v float32) {
	o.Weight = &v
}

// GetToken returns the Token field value
func (o *GithubRepositoryIn) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GithubRepositoryIn) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GithubRepositoryIn) SetToken(v string) {
	o.Token = v
}

func (o GithubRepositoryIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubRepositoryIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	toSerialize["owner"] = o.Owner
	if o.Ref.IsSet() {
		toSerialize["ref"] = o.Ref.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *GithubRepositoryIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
		"token",
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

	varGithubRepositoryIn := _GithubRepositoryIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubRepositoryIn)

	if err != nil {
		return err
	}

	*o = GithubRepositoryIn(varGithubRepositoryIn)

	return err
}

type NullableGithubRepositoryIn struct {
	value *GithubRepositoryIn
	isSet bool
}

func (v NullableGithubRepositoryIn) Get() *GithubRepositoryIn {
	return v.value
}

func (v *NullableGithubRepositoryIn) Set(val *GithubRepositoryIn) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubRepositoryIn) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubRepositoryIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubRepositoryIn(val *GithubRepositoryIn) *NullableGithubRepositoryIn {
	return &NullableGithubRepositoryIn{value: val, isSet: true}
}

func (v NullableGithubRepositoryIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubRepositoryIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


