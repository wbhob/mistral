# ModerationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string]bool** | Moderation result thresholds | [optional] 
**CategoryScores** | Pointer to **map[string]float32** | Moderation result | [optional] 

## Methods

### NewModerationObject

`func NewModerationObject() *ModerationObject`

NewModerationObject instantiates a new ModerationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationObjectWithDefaults

`func NewModerationObjectWithDefaults() *ModerationObject`

NewModerationObjectWithDefaults instantiates a new ModerationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ModerationObject) GetCategories() map[string]bool`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ModerationObject) GetCategoriesOk() (*map[string]bool, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ModerationObject) SetCategories(v map[string]bool)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ModerationObject) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoryScores

`func (o *ModerationObject) GetCategoryScores() map[string]float32`

GetCategoryScores returns the CategoryScores field if non-nil, zero value otherwise.

### GetCategoryScoresOk

`func (o *ModerationObject) GetCategoryScoresOk() (*map[string]float32, bool)`

GetCategoryScoresOk returns a tuple with the CategoryScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryScores

`func (o *ModerationObject) SetCategoryScores(v map[string]float32)`

SetCategoryScores sets CategoryScores field to given value.

### HasCategoryScores

`func (o *ModerationObject) HasCategoryScores() bool`

HasCategoryScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


