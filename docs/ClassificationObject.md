# ClassificationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string]bool** | Classifier result thresholded | [optional] 
**CategoryScores** | Pointer to **map[string]float32** | Classifier result | [optional] 

## Methods

### NewClassificationObject

`func NewClassificationObject() *ClassificationObject`

NewClassificationObject instantiates a new ClassificationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationObjectWithDefaults

`func NewClassificationObjectWithDefaults() *ClassificationObject`

NewClassificationObjectWithDefaults instantiates a new ClassificationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ClassificationObject) GetCategories() map[string]bool`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ClassificationObject) GetCategoriesOk() (*map[string]bool, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ClassificationObject) SetCategories(v map[string]bool)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ClassificationObject) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoryScores

`func (o *ClassificationObject) GetCategoryScores() map[string]float32`

GetCategoryScores returns the CategoryScores field if non-nil, zero value otherwise.

### GetCategoryScoresOk

`func (o *ClassificationObject) GetCategoryScoresOk() (*map[string]float32, bool)`

GetCategoryScoresOk returns a tuple with the CategoryScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryScores

`func (o *ClassificationObject) SetCategoryScores(v map[string]float32)`

SetCategoryScores sets CategoryScores field to given value.

### HasCategoryScores

`func (o *ClassificationObject) HasCategoryScores() bool`

HasCategoryScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


