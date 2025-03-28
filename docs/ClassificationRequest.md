# ClassificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. | 
**Input** | [**Input1**](Input1.md) |  | 

## Methods

### NewClassificationRequest

`func NewClassificationRequest(model string, input Input1, ) *ClassificationRequest`

NewClassificationRequest instantiates a new ClassificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationRequestWithDefaults

`func NewClassificationRequestWithDefaults() *ClassificationRequest`

NewClassificationRequestWithDefaults instantiates a new ClassificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ClassificationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClassificationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClassificationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *ClassificationRequest) GetInput() Input1`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ClassificationRequest) GetInputOk() (*Input1, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ClassificationRequest) SetInput(v Input1)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


