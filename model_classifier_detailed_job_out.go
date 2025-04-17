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

// checks if the ClassifierDetailedJobOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassifierDetailedJobOut{}

// ClassifierDetailedJobOut struct for ClassifierDetailedJobOut
type ClassifierDetailedJobOut struct {
	Id string `json:"id"`
	AutoStart bool `json:"auto_start"`
	Model FineTuneableModel `json:"model"`
	Status string `json:"status"`
	CreatedAt int32 `json:"created_at"`
	ModifiedAt int32 `json:"modified_at"`
	TrainingFiles []string `json:"training_files"`
	ValidationFiles []string `json:"validation_files,omitempty"`
	Object *string `json:"object,omitempty"`
	FineTunedModel NullableString `json:"fine_tuned_model,omitempty"`
	Suffix NullableString `json:"suffix,omitempty"`
	Integrations []ClassifierJobOutIntegrationsInner `json:"integrations,omitempty"`
	TrainedTokens NullableInt32 `json:"trained_tokens,omitempty"`
	Metadata NullableJobMetadataOut `json:"metadata,omitempty"`
	JobType *string `json:"job_type,omitempty"`
	Hyperparameters ClassifierTrainingParameters `json:"hyperparameters"`
	// Event items are created every time the status of a fine-tuning job changes. The timestamped list of all events is accessible here.
	Events []EventOut `json:"events,omitempty"`
	Checkpoints []CheckpointOut `json:"checkpoints,omitempty"`
	ClassifierTargets []ClassifierTargetOut `json:"classifier_targets"`
}

type _ClassifierDetailedJobOut ClassifierDetailedJobOut

// NewClassifierDetailedJobOut instantiates a new ClassifierDetailedJobOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifierDetailedJobOut(id string, autoStart bool, model FineTuneableModel, status string, createdAt int32, modifiedAt int32, trainingFiles []string, hyperparameters ClassifierTrainingParameters, classifierTargets []ClassifierTargetOut) *ClassifierDetailedJobOut {
	this := ClassifierDetailedJobOut{}
	this.Id = id
	this.AutoStart = autoStart
	this.Model = model
	this.Status = status
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.TrainingFiles = trainingFiles
	var object string = "job"
	this.Object = &object
	var jobType string = "classifier"
	this.JobType = &jobType
	this.Hyperparameters = hyperparameters
	this.ClassifierTargets = classifierTargets
	return &this
}

// NewClassifierDetailedJobOutWithDefaults instantiates a new ClassifierDetailedJobOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifierDetailedJobOutWithDefaults() *ClassifierDetailedJobOut {
	this := ClassifierDetailedJobOut{}
	var object string = "job"
	this.Object = &object
	var jobType string = "classifier"
	this.JobType = &jobType
	return &this
}

// GetId returns the Id field value
func (o *ClassifierDetailedJobOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClassifierDetailedJobOut) SetId(v string) {
	o.Id = v
}

// GetAutoStart returns the AutoStart field value
func (o *ClassifierDetailedJobOut) GetAutoStart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoStart
}

// GetAutoStartOk returns a tuple with the AutoStart field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetAutoStartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoStart, true
}

// SetAutoStart sets field value
func (o *ClassifierDetailedJobOut) SetAutoStart(v bool) {
	o.AutoStart = v
}

// GetModel returns the Model field value
func (o *ClassifierDetailedJobOut) GetModel() FineTuneableModel {
	if o == nil {
		var ret FineTuneableModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetModelOk() (*FineTuneableModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ClassifierDetailedJobOut) SetModel(v FineTuneableModel) {
	o.Model = v
}

// GetStatus returns the Status field value
func (o *ClassifierDetailedJobOut) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ClassifierDetailedJobOut) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ClassifierDetailedJobOut) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ClassifierDetailedJobOut) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ClassifierDetailedJobOut) GetModifiedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetModifiedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ClassifierDetailedJobOut) SetModifiedAt(v int32) {
	o.ModifiedAt = v
}

// GetTrainingFiles returns the TrainingFiles field value
func (o *ClassifierDetailedJobOut) GetTrainingFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrainingFiles
}

// GetTrainingFilesOk returns a tuple with the TrainingFiles field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetTrainingFilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainingFiles, true
}

// SetTrainingFiles sets field value
func (o *ClassifierDetailedJobOut) SetTrainingFiles(v []string) {
	o.TrainingFiles = v
}

// GetValidationFiles returns the ValidationFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetValidationFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ValidationFiles
}

// GetValidationFilesOk returns a tuple with the ValidationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetValidationFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.ValidationFiles) {
		return nil, false
	}
	return o.ValidationFiles, true
}

// HasValidationFiles returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasValidationFiles() bool {
	if o != nil && !IsNil(o.ValidationFiles) {
		return true
	}

	return false
}

// SetValidationFiles gets a reference to the given []string and assigns it to the ValidationFiles field.
func (o *ClassifierDetailedJobOut) SetValidationFiles(v []string) {
	o.ValidationFiles = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ClassifierDetailedJobOut) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ClassifierDetailedJobOut) SetObject(v string) {
	o.Object = &v
}

// GetFineTunedModel returns the FineTunedModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetFineTunedModel() string {
	if o == nil || IsNil(o.FineTunedModel.Get()) {
		var ret string
		return ret
	}
	return *o.FineTunedModel.Get()
}

// GetFineTunedModelOk returns a tuple with the FineTunedModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetFineTunedModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FineTunedModel.Get(), o.FineTunedModel.IsSet()
}

// HasFineTunedModel returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasFineTunedModel() bool {
	if o != nil && o.FineTunedModel.IsSet() {
		return true
	}

	return false
}

// SetFineTunedModel gets a reference to the given NullableString and assigns it to the FineTunedModel field.
func (o *ClassifierDetailedJobOut) SetFineTunedModel(v string) {
	o.FineTunedModel.Set(&v)
}
// SetFineTunedModelNil sets the value for FineTunedModel to be an explicit nil
func (o *ClassifierDetailedJobOut) SetFineTunedModelNil() {
	o.FineTunedModel.Set(nil)
}

// UnsetFineTunedModel ensures that no value is present for FineTunedModel, not even an explicit nil
func (o *ClassifierDetailedJobOut) UnsetFineTunedModel() {
	o.FineTunedModel.Unset()
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetSuffix() string {
	if o == nil || IsNil(o.Suffix.Get()) {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *ClassifierDetailedJobOut) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *ClassifierDetailedJobOut) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *ClassifierDetailedJobOut) UnsetSuffix() {
	o.Suffix.Unset()
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetIntegrations() []ClassifierJobOutIntegrationsInner {
	if o == nil {
		var ret []ClassifierJobOutIntegrationsInner
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetIntegrationsOk() ([]ClassifierJobOutIntegrationsInner, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given []ClassifierJobOutIntegrationsInner and assigns it to the Integrations field.
func (o *ClassifierDetailedJobOut) SetIntegrations(v []ClassifierJobOutIntegrationsInner) {
	o.Integrations = v
}

// GetTrainedTokens returns the TrainedTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetTrainedTokens() int32 {
	if o == nil || IsNil(o.TrainedTokens.Get()) {
		var ret int32
		return ret
	}
	return *o.TrainedTokens.Get()
}

// GetTrainedTokensOk returns a tuple with the TrainedTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetTrainedTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainedTokens.Get(), o.TrainedTokens.IsSet()
}

// HasTrainedTokens returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasTrainedTokens() bool {
	if o != nil && o.TrainedTokens.IsSet() {
		return true
	}

	return false
}

// SetTrainedTokens gets a reference to the given NullableInt32 and assigns it to the TrainedTokens field.
func (o *ClassifierDetailedJobOut) SetTrainedTokens(v int32) {
	o.TrainedTokens.Set(&v)
}
// SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil
func (o *ClassifierDetailedJobOut) SetTrainedTokensNil() {
	o.TrainedTokens.Set(nil)
}

// UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
func (o *ClassifierDetailedJobOut) UnsetTrainedTokens() {
	o.TrainedTokens.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClassifierDetailedJobOut) GetMetadata() JobMetadataOut {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret JobMetadataOut
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClassifierDetailedJobOut) GetMetadataOk() (*JobMetadataOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableJobMetadataOut and assigns it to the Metadata field.
func (o *ClassifierDetailedJobOut) SetMetadata(v JobMetadataOut) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *ClassifierDetailedJobOut) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *ClassifierDetailedJobOut) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *ClassifierDetailedJobOut) GetJobType() string {
	if o == nil || IsNil(o.JobType) {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetJobTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasJobType() bool {
	if o != nil && !IsNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *ClassifierDetailedJobOut) SetJobType(v string) {
	o.JobType = &v
}

// GetHyperparameters returns the Hyperparameters field value
func (o *ClassifierDetailedJobOut) GetHyperparameters() ClassifierTrainingParameters {
	if o == nil {
		var ret ClassifierTrainingParameters
		return ret
	}

	return o.Hyperparameters
}

// GetHyperparametersOk returns a tuple with the Hyperparameters field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetHyperparametersOk() (*ClassifierTrainingParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hyperparameters, true
}

// SetHyperparameters sets field value
func (o *ClassifierDetailedJobOut) SetHyperparameters(v ClassifierTrainingParameters) {
	o.Hyperparameters = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ClassifierDetailedJobOut) GetEvents() []EventOut {
	if o == nil || IsNil(o.Events) {
		var ret []EventOut
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetEventsOk() ([]EventOut, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []EventOut and assigns it to the Events field.
func (o *ClassifierDetailedJobOut) SetEvents(v []EventOut) {
	o.Events = v
}

// GetCheckpoints returns the Checkpoints field value if set, zero value otherwise.
func (o *ClassifierDetailedJobOut) GetCheckpoints() []CheckpointOut {
	if o == nil || IsNil(o.Checkpoints) {
		var ret []CheckpointOut
		return ret
	}
	return o.Checkpoints
}

// GetCheckpointsOk returns a tuple with the Checkpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetCheckpointsOk() ([]CheckpointOut, bool) {
	if o == nil || IsNil(o.Checkpoints) {
		return nil, false
	}
	return o.Checkpoints, true
}

// HasCheckpoints returns a boolean if a field has been set.
func (o *ClassifierDetailedJobOut) HasCheckpoints() bool {
	if o != nil && !IsNil(o.Checkpoints) {
		return true
	}

	return false
}

// SetCheckpoints gets a reference to the given []CheckpointOut and assigns it to the Checkpoints field.
func (o *ClassifierDetailedJobOut) SetCheckpoints(v []CheckpointOut) {
	o.Checkpoints = v
}

// GetClassifierTargets returns the ClassifierTargets field value
func (o *ClassifierDetailedJobOut) GetClassifierTargets() []ClassifierTargetOut {
	if o == nil {
		var ret []ClassifierTargetOut
		return ret
	}

	return o.ClassifierTargets
}

// GetClassifierTargetsOk returns a tuple with the ClassifierTargets field value
// and a boolean to check if the value has been set.
func (o *ClassifierDetailedJobOut) GetClassifierTargetsOk() ([]ClassifierTargetOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClassifierTargets, true
}

// SetClassifierTargets sets field value
func (o *ClassifierDetailedJobOut) SetClassifierTargets(v []ClassifierTargetOut) {
	o.ClassifierTargets = v
}

func (o ClassifierDetailedJobOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassifierDetailedJobOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["auto_start"] = o.AutoStart
	toSerialize["model"] = o.Model
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["training_files"] = o.TrainingFiles
	if o.ValidationFiles != nil {
		toSerialize["validation_files"] = o.ValidationFiles
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if o.FineTunedModel.IsSet() {
		toSerialize["fine_tuned_model"] = o.FineTunedModel.Get()
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.TrainedTokens.IsSet() {
		toSerialize["trained_tokens"] = o.TrainedTokens.Get()
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	if !IsNil(o.JobType) {
		toSerialize["job_type"] = o.JobType
	}
	toSerialize["hyperparameters"] = o.Hyperparameters
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Checkpoints) {
		toSerialize["checkpoints"] = o.Checkpoints
	}
	toSerialize["classifier_targets"] = o.ClassifierTargets
	return toSerialize, nil
}

func (o *ClassifierDetailedJobOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"auto_start",
		"model",
		"status",
		"created_at",
		"modified_at",
		"training_files",
		"hyperparameters",
		"classifier_targets",
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

	varClassifierDetailedJobOut := _ClassifierDetailedJobOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClassifierDetailedJobOut)

	if err != nil {
		return err
	}

	*o = ClassifierDetailedJobOut(varClassifierDetailedJobOut)

	return err
}

type NullableClassifierDetailedJobOut struct {
	value *ClassifierDetailedJobOut
	isSet bool
}

func (v NullableClassifierDetailedJobOut) Get() *ClassifierDetailedJobOut {
	return v.value
}

func (v *NullableClassifierDetailedJobOut) Set(val *ClassifierDetailedJobOut) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifierDetailedJobOut) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifierDetailedJobOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifierDetailedJobOut(val *ClassifierDetailedJobOut) *NullableClassifierDetailedJobOut {
	return &NullableClassifierDetailedJobOut{value: val, isSet: true}
}

func (v NullableClassifierDetailedJobOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifierDetailedJobOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


