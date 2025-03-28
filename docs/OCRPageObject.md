# OCRPageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The page index in a pdf document starting from 0 | 
**Markdown** | **string** | The markdown string response of the page | 
**Images** | [**[]OCRImageObject**](OCRImageObject.md) | List of all extracted images in the page | 
**Dimensions** | [**NullableOCRPageDimensions**](OCRPageDimensions.md) |  | 

## Methods

### NewOCRPageObject

`func NewOCRPageObject(index int32, markdown string, images []OCRImageObject, dimensions NullableOCRPageDimensions, ) *OCRPageObject`

NewOCRPageObject instantiates a new OCRPageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRPageObjectWithDefaults

`func NewOCRPageObjectWithDefaults() *OCRPageObject`

NewOCRPageObjectWithDefaults instantiates a new OCRPageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *OCRPageObject) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OCRPageObject) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OCRPageObject) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMarkdown

`func (o *OCRPageObject) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *OCRPageObject) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *OCRPageObject) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetImages

`func (o *OCRPageObject) GetImages() []OCRImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *OCRPageObject) GetImagesOk() (*[]OCRImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *OCRPageObject) SetImages(v []OCRImageObject)`

SetImages sets Images field to given value.


### GetDimensions

`func (o *OCRPageObject) GetDimensions() OCRPageDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OCRPageObject) GetDimensionsOk() (*OCRPageDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OCRPageObject) SetDimensions(v OCRPageDimensions)`

SetDimensions sets Dimensions field to given value.


### SetDimensionsNil

`func (o *OCRPageObject) SetDimensionsNil(b bool)`

 SetDimensionsNil sets the value for Dimensions to be an explicit nil

### UnsetDimensions
`func (o *OCRPageObject) UnsetDimensions()`

UnsetDimensions ensures that no value is present for Dimensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


