# EmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. | 
**Input** | [**Input2**](Input2.md) |  | 

## Methods

### NewEmbeddingRequest

`func NewEmbeddingRequest(model string, input Input2, ) *EmbeddingRequest`

NewEmbeddingRequest instantiates a new EmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingRequestWithDefaults

`func NewEmbeddingRequestWithDefaults() *EmbeddingRequest`

NewEmbeddingRequestWithDefaults instantiates a new EmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *EmbeddingRequest) GetInput() Input2`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EmbeddingRequest) GetInputOk() (*Input2, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EmbeddingRequest) SetInput(v Input2)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


