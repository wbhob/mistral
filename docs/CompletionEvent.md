# CompletionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CompletionChunk**](CompletionChunk.md) |  | 

## Methods

### NewCompletionEvent

`func NewCompletionEvent(data CompletionChunk, ) *CompletionEvent`

NewCompletionEvent instantiates a new CompletionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionEventWithDefaults

`func NewCompletionEventWithDefaults() *CompletionEvent`

NewCompletionEventWithDefaults instantiates a new CompletionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CompletionEvent) GetData() CompletionChunk`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompletionEvent) GetDataOk() (*CompletionChunk, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompletionEvent) SetData(v CompletionChunk)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


