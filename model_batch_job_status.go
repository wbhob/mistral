/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"fmt"
)

// BatchJobStatus the model 'BatchJobStatus'
type BatchJobStatus string

// List of BatchJobStatus
const (
	BATCHJOBSTATUS_QUEUED BatchJobStatus = "QUEUED"
	BATCHJOBSTATUS_RUNNING BatchJobStatus = "RUNNING"
	BATCHJOBSTATUS_SUCCESS BatchJobStatus = "SUCCESS"
	BATCHJOBSTATUS_FAILED BatchJobStatus = "FAILED"
	BATCHJOBSTATUS_TIMEOUT_EXCEEDED BatchJobStatus = "TIMEOUT_EXCEEDED"
	BATCHJOBSTATUS_CANCELLATION_REQUESTED BatchJobStatus = "CANCELLATION_REQUESTED"
	BATCHJOBSTATUS_CANCELLED BatchJobStatus = "CANCELLED"
)

// All allowed values of BatchJobStatus enum
var AllowedBatchJobStatusEnumValues = []BatchJobStatus{
	"QUEUED",
	"RUNNING",
	"SUCCESS",
	"FAILED",
	"TIMEOUT_EXCEEDED",
	"CANCELLATION_REQUESTED",
	"CANCELLED",
}

func (v *BatchJobStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BatchJobStatus(value)
	for _, existing := range AllowedBatchJobStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BatchJobStatus", value)
}

// NewBatchJobStatusFromValue returns a pointer to a valid BatchJobStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBatchJobStatusFromValue(v string) (*BatchJobStatus, error) {
	ev := BatchJobStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BatchJobStatus: valid values are %v", v, AllowedBatchJobStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BatchJobStatus) IsValid() bool {
	for _, existing := range AllowedBatchJobStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BatchJobStatus value
func (v BatchJobStatus) Ptr() *BatchJobStatus {
	return &v
}

type NullableBatchJobStatus struct {
	value *BatchJobStatus
	isSet bool
}

func (v NullableBatchJobStatus) Get() *BatchJobStatus {
	return v.value
}

func (v *NullableBatchJobStatus) Set(val *BatchJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchJobStatus(val *BatchJobStatus) *NullableBatchJobStatus {
	return &NullableBatchJobStatus{value: val, isSet: true}
}

func (v NullableBatchJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

