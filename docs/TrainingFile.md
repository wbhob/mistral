# TrainingFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** |  | 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]

## Methods

### NewTrainingFile

`func NewTrainingFile(fileId string, ) *TrainingFile`

NewTrainingFile instantiates a new TrainingFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingFileWithDefaults

`func NewTrainingFileWithDefaults() *TrainingFile`

NewTrainingFileWithDefaults instantiates a new TrainingFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *TrainingFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *TrainingFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *TrainingFile) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetWeight

`func (o *TrainingFile) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TrainingFile) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TrainingFile) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *TrainingFile) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


