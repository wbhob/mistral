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

// checks if the RetrieveFileOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveFileOut{}

// RetrieveFileOut struct for RetrieveFileOut
type RetrieveFileOut struct {
	// The unique identifier of the file.
	Id string `json:"id"`
	// The object type, which is always \"file\".
	Object string `json:"object"`
	// The size of the file, in bytes.
	Bytes int32 `json:"bytes"`
	// The UNIX timestamp (in seconds) of the event.
	CreatedAt int32 `json:"created_at"`
	// The name of the uploaded file.
	Filename string `json:"filename"`
	// The intended purpose of the uploaded file. Only accepts fine-tuning (`fine-tune`) for now.
	Purpose FilePurpose `json:"purpose"`
	SampleType SampleType `json:"sample_type"`
	NumLines NullableInt32 `json:"num_lines,omitempty"`
	Source Source `json:"source"`
	Deleted bool `json:"deleted"`
}

type _RetrieveFileOut RetrieveFileOut

// NewRetrieveFileOut instantiates a new RetrieveFileOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveFileOut(id string, object string, bytes int32, createdAt int32, filename string, purpose FilePurpose, sampleType SampleType, source Source, deleted bool) *RetrieveFileOut {
	this := RetrieveFileOut{}
	this.Id = id
	this.Object = object
	this.Bytes = bytes
	this.CreatedAt = createdAt
	this.Filename = filename
	this.Purpose = purpose
	this.SampleType = sampleType
	this.Source = source
	this.Deleted = deleted
	return &this
}

// NewRetrieveFileOutWithDefaults instantiates a new RetrieveFileOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveFileOutWithDefaults() *RetrieveFileOut {
	this := RetrieveFileOut{}
	return &this
}

// GetId returns the Id field value
func (o *RetrieveFileOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RetrieveFileOut) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *RetrieveFileOut) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RetrieveFileOut) SetObject(v string) {
	o.Object = v
}

// GetBytes returns the Bytes field value
func (o *RetrieveFileOut) GetBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *RetrieveFileOut) SetBytes(v int32) {
	o.Bytes = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RetrieveFileOut) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RetrieveFileOut) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetFilename returns the Filename field value
func (o *RetrieveFileOut) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *RetrieveFileOut) SetFilename(v string) {
	o.Filename = v
}

// GetPurpose returns the Purpose field value
func (o *RetrieveFileOut) GetPurpose() FilePurpose {
	if o == nil {
		var ret FilePurpose
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetPurposeOk() (*FilePurpose, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *RetrieveFileOut) SetPurpose(v FilePurpose) {
	o.Purpose = v
}

// GetSampleType returns the SampleType field value
func (o *RetrieveFileOut) GetSampleType() SampleType {
	if o == nil {
		var ret SampleType
		return ret
	}

	return o.SampleType
}

// GetSampleTypeOk returns a tuple with the SampleType field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetSampleTypeOk() (*SampleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleType, true
}

// SetSampleType sets field value
func (o *RetrieveFileOut) SetSampleType(v SampleType) {
	o.SampleType = v
}

// GetNumLines returns the NumLines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveFileOut) GetNumLines() int32 {
	if o == nil || IsNil(o.NumLines.Get()) {
		var ret int32
		return ret
	}
	return *o.NumLines.Get()
}

// GetNumLinesOk returns a tuple with the NumLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveFileOut) GetNumLinesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumLines.Get(), o.NumLines.IsSet()
}

// HasNumLines returns a boolean if a field has been set.
func (o *RetrieveFileOut) HasNumLines() bool {
	if o != nil && o.NumLines.IsSet() {
		return true
	}

	return false
}

// SetNumLines gets a reference to the given NullableInt32 and assigns it to the NumLines field.
func (o *RetrieveFileOut) SetNumLines(v int32) {
	o.NumLines.Set(&v)
}
// SetNumLinesNil sets the value for NumLines to be an explicit nil
func (o *RetrieveFileOut) SetNumLinesNil() {
	o.NumLines.Set(nil)
}

// UnsetNumLines ensures that no value is present for NumLines, not even an explicit nil
func (o *RetrieveFileOut) UnsetNumLines() {
	o.NumLines.Unset()
}

// GetSource returns the Source field value
func (o *RetrieveFileOut) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *RetrieveFileOut) SetSource(v Source) {
	o.Source = v
}

// GetDeleted returns the Deleted field value
func (o *RetrieveFileOut) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *RetrieveFileOut) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *RetrieveFileOut) SetDeleted(v bool) {
	o.Deleted = v
}

func (o RetrieveFileOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetrieveFileOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["bytes"] = o.Bytes
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["filename"] = o.Filename
	toSerialize["purpose"] = o.Purpose
	toSerialize["sample_type"] = o.SampleType
	if o.NumLines.IsSet() {
		toSerialize["num_lines"] = o.NumLines.Get()
	}
	toSerialize["source"] = o.Source
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

func (o *RetrieveFileOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"bytes",
		"created_at",
		"filename",
		"purpose",
		"sample_type",
		"source",
		"deleted",
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

	varRetrieveFileOut := _RetrieveFileOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRetrieveFileOut)

	if err != nil {
		return err
	}

	*o = RetrieveFileOut(varRetrieveFileOut)

	return err
}

type NullableRetrieveFileOut struct {
	value *RetrieveFileOut
	isSet bool
}

func (v NullableRetrieveFileOut) Get() *RetrieveFileOut {
	return v.value
}

func (v *NullableRetrieveFileOut) Set(val *RetrieveFileOut) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveFileOut) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveFileOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveFileOut(val *RetrieveFileOut) *NullableRetrieveFileOut {
	return &NullableRetrieveFileOut{value: val, isSet: true}
}

func (v NullableRetrieveFileOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveFileOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


