# ClassifierTargetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Labels** | **[]string** |  | 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]
**LossFunction** | Pointer to [**NullableFTClassifierLossFunction**](FTClassifierLossFunction.md) |  | [optional] 

## Methods

### NewClassifierTargetIn

`func NewClassifierTargetIn(name string, labels []string, ) *ClassifierTargetIn`

NewClassifierTargetIn instantiates a new ClassifierTargetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierTargetInWithDefaults

`func NewClassifierTargetInWithDefaults() *ClassifierTargetIn`

NewClassifierTargetInWithDefaults instantiates a new ClassifierTargetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClassifierTargetIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassifierTargetIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassifierTargetIn) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *ClassifierTargetIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClassifierTargetIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClassifierTargetIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetWeight

`func (o *ClassifierTargetIn) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ClassifierTargetIn) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ClassifierTargetIn) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ClassifierTargetIn) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetLossFunction

`func (o *ClassifierTargetIn) GetLossFunction() FTClassifierLossFunction`

GetLossFunction returns the LossFunction field if non-nil, zero value otherwise.

### GetLossFunctionOk

`func (o *ClassifierTargetIn) GetLossFunctionOk() (*FTClassifierLossFunction, bool)`

GetLossFunctionOk returns a tuple with the LossFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossFunction

`func (o *ClassifierTargetIn) SetLossFunction(v FTClassifierLossFunction)`

SetLossFunction sets LossFunction field to given value.

### HasLossFunction

`func (o *ClassifierTargetIn) HasLossFunction() bool`

HasLossFunction returns a boolean if a field has been set.

### SetLossFunctionNil

`func (o *ClassifierTargetIn) SetLossFunctionNil(b bool)`

 SetLossFunctionNil sets the value for LossFunction to be an explicit nil

### UnsetLossFunction
`func (o *ClassifierTargetIn) UnsetLossFunction()`

UnsetLossFunction ensures that no value is present for LossFunction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


