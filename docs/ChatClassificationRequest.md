# ChatClassificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | [**ChatClassificationRequestInputs**](ChatClassificationRequestInputs.md) |  | 

## Methods

### NewChatClassificationRequest

`func NewChatClassificationRequest(model string, input ChatClassificationRequestInputs, ) *ChatClassificationRequest`

NewChatClassificationRequest instantiates a new ChatClassificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatClassificationRequestWithDefaults

`func NewChatClassificationRequestWithDefaults() *ChatClassificationRequest`

NewChatClassificationRequestWithDefaults instantiates a new ChatClassificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatClassificationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatClassificationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatClassificationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *ChatClassificationRequest) GetInput() ChatClassificationRequestInputs`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ChatClassificationRequest) GetInputOk() (*ChatClassificationRequestInputs, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ChatClassificationRequest) SetInput(v ChatClassificationRequestInputs)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


