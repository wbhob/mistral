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

// checks if the BatchJobIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchJobIn{}

// BatchJobIn struct for BatchJobIn
type BatchJobIn struct {
	InputFiles []string `json:"input_files"`
	Endpoint ApiEndpoint `json:"endpoint"`
	Model string `json:"model"`
	Metadata map[string]string `json:"metadata,omitempty"`
	TimeoutHours *int32 `json:"timeout_hours,omitempty"`
}

type _BatchJobIn BatchJobIn

// NewBatchJobIn instantiates a new BatchJobIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchJobIn(inputFiles []string, endpoint ApiEndpoint, model string) *BatchJobIn {
	this := BatchJobIn{}
	this.InputFiles = inputFiles
	this.Endpoint = endpoint
	this.Model = model
	var timeoutHours int32 = 24
	this.TimeoutHours = &timeoutHours
	return &this
}

// NewBatchJobInWithDefaults instantiates a new BatchJobIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchJobInWithDefaults() *BatchJobIn {
	this := BatchJobIn{}
	var timeoutHours int32 = 24
	this.TimeoutHours = &timeoutHours
	return &this
}

// GetInputFiles returns the InputFiles field value
func (o *BatchJobIn) GetInputFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InputFiles
}

// GetInputFilesOk returns a tuple with the InputFiles field value
// and a boolean to check if the value has been set.
func (o *BatchJobIn) GetInputFilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputFiles, true
}

// SetInputFiles sets field value
func (o *BatchJobIn) SetInputFiles(v []string) {
	o.InputFiles = v
}

// GetEndpoint returns the Endpoint field value
func (o *BatchJobIn) GetEndpoint() ApiEndpoint {
	if o == nil {
		var ret ApiEndpoint
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *BatchJobIn) GetEndpointOk() (*ApiEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *BatchJobIn) SetEndpoint(v ApiEndpoint) {
	o.Endpoint = v
}

// GetModel returns the Model field value
func (o *BatchJobIn) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BatchJobIn) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BatchJobIn) SetModel(v string) {
	o.Model = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchJobIn) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchJobIn) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BatchJobIn) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BatchJobIn) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetTimeoutHours returns the TimeoutHours field value if set, zero value otherwise.
func (o *BatchJobIn) GetTimeoutHours() int32 {
	if o == nil || IsNil(o.TimeoutHours) {
		var ret int32
		return ret
	}
	return *o.TimeoutHours
}

// GetTimeoutHoursOk returns a tuple with the TimeoutHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchJobIn) GetTimeoutHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutHours) {
		return nil, false
	}
	return o.TimeoutHours, true
}

// HasTimeoutHours returns a boolean if a field has been set.
func (o *BatchJobIn) HasTimeoutHours() bool {
	if o != nil && !IsNil(o.TimeoutHours) {
		return true
	}

	return false
}

// SetTimeoutHours gets a reference to the given int32 and assigns it to the TimeoutHours field.
func (o *BatchJobIn) SetTimeoutHours(v int32) {
	o.TimeoutHours = &v
}

func (o BatchJobIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchJobIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input_files"] = o.InputFiles
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["model"] = o.Model
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.TimeoutHours) {
		toSerialize["timeout_hours"] = o.TimeoutHours
	}
	return toSerialize, nil
}

func (o *BatchJobIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input_files",
		"endpoint",
		"model",
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

	varBatchJobIn := _BatchJobIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchJobIn)

	if err != nil {
		return err
	}

	*o = BatchJobIn(varBatchJobIn)

	return err
}

type NullableBatchJobIn struct {
	value *BatchJobIn
	isSet bool
}

func (v NullableBatchJobIn) Get() *BatchJobIn {
	return v.value
}

func (v *NullableBatchJobIn) Set(val *BatchJobIn) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchJobIn) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchJobIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchJobIn(val *BatchJobIn) *NullableBatchJobIn {
	return &NullableBatchJobIn{value: val, isSet: true}
}

func (v NullableBatchJobIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchJobIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


