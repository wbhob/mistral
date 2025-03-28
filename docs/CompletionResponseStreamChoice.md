# CompletionResponseStreamChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Delta** | [**DeltaMessage**](DeltaMessage.md) |  | 
**FinishReason** | **NullableString** |  | 

## Methods

### NewCompletionResponseStreamChoice

`func NewCompletionResponseStreamChoice(index int32, delta DeltaMessage, finishReason NullableString, ) *CompletionResponseStreamChoice`

NewCompletionResponseStreamChoice instantiates a new CompletionResponseStreamChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionResponseStreamChoiceWithDefaults

`func NewCompletionResponseStreamChoiceWithDefaults() *CompletionResponseStreamChoice`

NewCompletionResponseStreamChoiceWithDefaults instantiates a new CompletionResponseStreamChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *CompletionResponseStreamChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CompletionResponseStreamChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CompletionResponseStreamChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetDelta

`func (o *CompletionResponseStreamChoice) GetDelta() DeltaMessage`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *CompletionResponseStreamChoice) GetDeltaOk() (*DeltaMessage, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *CompletionResponseStreamChoice) SetDelta(v DeltaMessage)`

SetDelta sets Delta field to given value.


### GetFinishReason

`func (o *CompletionResponseStreamChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *CompletionResponseStreamChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *CompletionResponseStreamChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.


### SetFinishReasonNil

`func (o *CompletionResponseStreamChoice) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *CompletionResponseStreamChoice) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


