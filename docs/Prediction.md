# Prediction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "content"]
**Content** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewPrediction

`func NewPrediction() *Prediction`

NewPrediction instantiates a new Prediction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionWithDefaults

`func NewPredictionWithDefaults() *Prediction`

NewPredictionWithDefaults instantiates a new Prediction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Prediction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Prediction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Prediction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Prediction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *Prediction) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Prediction) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Prediction) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Prediction) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


