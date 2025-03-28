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
	"gopkg.in/validator.v2"
)

// JobInRepositoriesInner - struct for JobInRepositoriesInner
type JobInRepositoriesInner struct {
	GithubRepositoryIn *GithubRepositoryIn
}

// GithubRepositoryInAsJobInRepositoriesInner is a convenience function that returns GithubRepositoryIn wrapped in JobInRepositoriesInner
func GithubRepositoryInAsJobInRepositoriesInner(v *GithubRepositoryIn) JobInRepositoriesInner {
	return JobInRepositoriesInner{
		GithubRepositoryIn: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *JobInRepositoriesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GithubRepositoryIn
	err = newStrictDecoder(data).Decode(&dst.GithubRepositoryIn)
	if err == nil {
		jsonGithubRepositoryIn, _ := json.Marshal(dst.GithubRepositoryIn)
		if string(jsonGithubRepositoryIn) == "{}" { // empty struct
			dst.GithubRepositoryIn = nil
		} else {
			if err = validator.Validate(dst.GithubRepositoryIn); err != nil {
				dst.GithubRepositoryIn = nil
			} else {
				match++
			}
		}
	} else {
		dst.GithubRepositoryIn = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GithubRepositoryIn = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JobInRepositoriesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JobInRepositoriesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JobInRepositoriesInner) MarshalJSON() ([]byte, error) {
	if src.GithubRepositoryIn != nil {
		return json.Marshal(&src.GithubRepositoryIn)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JobInRepositoriesInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GithubRepositoryIn != nil {
		return obj.GithubRepositoryIn
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj JobInRepositoriesInner) GetActualInstanceValue() (interface{}) {
	if obj.GithubRepositoryIn != nil {
		return *obj.GithubRepositoryIn
	}

	// all schemas are nil
	return nil
}

type NullableJobInRepositoriesInner struct {
	value *JobInRepositoriesInner
	isSet bool
}

func (v NullableJobInRepositoriesInner) Get() *JobInRepositoriesInner {
	return v.value
}

func (v *NullableJobInRepositoriesInner) Set(val *JobInRepositoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInRepositoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInRepositoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInRepositoriesInner(val *JobInRepositoriesInner) *NullableJobInRepositoriesInner {
	return &NullableJobInRepositoriesInner{value: val, isSet: true}
}

func (v NullableJobInRepositoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInRepositoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


