# FunctionCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Arguments** | [**Arguments**](Arguments.md) |  | 

## Methods

### NewFunctionCall

`func NewFunctionCall(name string, arguments Arguments, ) *FunctionCall`

NewFunctionCall instantiates a new FunctionCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCallWithDefaults

`func NewFunctionCallWithDefaults() *FunctionCall`

NewFunctionCallWithDefaults instantiates a new FunctionCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionCall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionCall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionCall) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *FunctionCall) GetArguments() Arguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *FunctionCall) GetArgumentsOk() (*Arguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *FunctionCall) SetArguments(v Arguments)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


