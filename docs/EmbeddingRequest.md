# EmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. | 
**Input** | [**Input2**](Input2.md) |  | 
**OutputDimension** | Pointer to **NullableInt32** |  | [optional] 
**OutputDtype** | Pointer to [**EmbeddingDtype**](EmbeddingDtype.md) | The data type of the output embeddings. | [optional] 

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


### GetOutputDimension

`func (o *EmbeddingRequest) GetOutputDimension() int32`

GetOutputDimension returns the OutputDimension field if non-nil, zero value otherwise.

### GetOutputDimensionOk

`func (o *EmbeddingRequest) GetOutputDimensionOk() (*int32, bool)`

GetOutputDimensionOk returns a tuple with the OutputDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDimension

`func (o *EmbeddingRequest) SetOutputDimension(v int32)`

SetOutputDimension sets OutputDimension field to given value.

### HasOutputDimension

`func (o *EmbeddingRequest) HasOutputDimension() bool`

HasOutputDimension returns a boolean if a field has been set.

### SetOutputDimensionNil

`func (o *EmbeddingRequest) SetOutputDimensionNil(b bool)`

 SetOutputDimensionNil sets the value for OutputDimension to be an explicit nil

### UnsetOutputDimension
`func (o *EmbeddingRequest) UnsetOutputDimension()`

UnsetOutputDimension ensures that no value is present for OutputDimension, not even an explicit nil
### GetOutputDtype

`func (o *EmbeddingRequest) GetOutputDtype() EmbeddingDtype`

GetOutputDtype returns the OutputDtype field if non-nil, zero value otherwise.

### GetOutputDtypeOk

`func (o *EmbeddingRequest) GetOutputDtypeOk() (*EmbeddingDtype, bool)`

GetOutputDtypeOk returns a tuple with the OutputDtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDtype

`func (o *EmbeddingRequest) SetOutputDtype(v EmbeddingDtype)`

SetOutputDtype sets OutputDtype field to given value.

### HasOutputDtype

`func (o *EmbeddingRequest) HasOutputDtype() bool`

HasOutputDtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


