# OCRImageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Image ID for extracted image in a page | 
**TopLeftX** | **NullableInt32** |  | 
**TopLeftY** | **NullableInt32** |  | 
**BottomRightX** | **NullableInt32** |  | 
**BottomRightY** | **NullableInt32** |  | 
**ImageBase64** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOCRImageObject

`func NewOCRImageObject(id string, topLeftX NullableInt32, topLeftY NullableInt32, bottomRightX NullableInt32, bottomRightY NullableInt32, ) *OCRImageObject`

NewOCRImageObject instantiates a new OCRImageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRImageObjectWithDefaults

`func NewOCRImageObjectWithDefaults() *OCRImageObject`

NewOCRImageObjectWithDefaults instantiates a new OCRImageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OCRImageObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OCRImageObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OCRImageObject) SetId(v string)`

SetId sets Id field to given value.


### GetTopLeftX

`func (o *OCRImageObject) GetTopLeftX() int32`

GetTopLeftX returns the TopLeftX field if non-nil, zero value otherwise.

### GetTopLeftXOk

`func (o *OCRImageObject) GetTopLeftXOk() (*int32, bool)`

GetTopLeftXOk returns a tuple with the TopLeftX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftX

`func (o *OCRImageObject) SetTopLeftX(v int32)`

SetTopLeftX sets TopLeftX field to given value.


### SetTopLeftXNil

`func (o *OCRImageObject) SetTopLeftXNil(b bool)`

 SetTopLeftXNil sets the value for TopLeftX to be an explicit nil

### UnsetTopLeftX
`func (o *OCRImageObject) UnsetTopLeftX()`

UnsetTopLeftX ensures that no value is present for TopLeftX, not even an explicit nil
### GetTopLeftY

`func (o *OCRImageObject) GetTopLeftY() int32`

GetTopLeftY returns the TopLeftY field if non-nil, zero value otherwise.

### GetTopLeftYOk

`func (o *OCRImageObject) GetTopLeftYOk() (*int32, bool)`

GetTopLeftYOk returns a tuple with the TopLeftY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftY

`func (o *OCRImageObject) SetTopLeftY(v int32)`

SetTopLeftY sets TopLeftY field to given value.


### SetTopLeftYNil

`func (o *OCRImageObject) SetTopLeftYNil(b bool)`

 SetTopLeftYNil sets the value for TopLeftY to be an explicit nil

### UnsetTopLeftY
`func (o *OCRImageObject) UnsetTopLeftY()`

UnsetTopLeftY ensures that no value is present for TopLeftY, not even an explicit nil
### GetBottomRightX

`func (o *OCRImageObject) GetBottomRightX() int32`

GetBottomRightX returns the BottomRightX field if non-nil, zero value otherwise.

### GetBottomRightXOk

`func (o *OCRImageObject) GetBottomRightXOk() (*int32, bool)`

GetBottomRightXOk returns a tuple with the BottomRightX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightX

`func (o *OCRImageObject) SetBottomRightX(v int32)`

SetBottomRightX sets BottomRightX field to given value.


### SetBottomRightXNil

`func (o *OCRImageObject) SetBottomRightXNil(b bool)`

 SetBottomRightXNil sets the value for BottomRightX to be an explicit nil

### UnsetBottomRightX
`func (o *OCRImageObject) UnsetBottomRightX()`

UnsetBottomRightX ensures that no value is present for BottomRightX, not even an explicit nil
### GetBottomRightY

`func (o *OCRImageObject) GetBottomRightY() int32`

GetBottomRightY returns the BottomRightY field if non-nil, zero value otherwise.

### GetBottomRightYOk

`func (o *OCRImageObject) GetBottomRightYOk() (*int32, bool)`

GetBottomRightYOk returns a tuple with the BottomRightY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightY

`func (o *OCRImageObject) SetBottomRightY(v int32)`

SetBottomRightY sets BottomRightY field to given value.


### SetBottomRightYNil

`func (o *OCRImageObject) SetBottomRightYNil(b bool)`

 SetBottomRightYNil sets the value for BottomRightY to be an explicit nil

### UnsetBottomRightY
`func (o *OCRImageObject) UnsetBottomRightY()`

UnsetBottomRightY ensures that no value is present for BottomRightY, not even an explicit nil
### GetImageBase64

`func (o *OCRImageObject) GetImageBase64() string`

GetImageBase64 returns the ImageBase64 field if non-nil, zero value otherwise.

### GetImageBase64Ok

`func (o *OCRImageObject) GetImageBase64Ok() (*string, bool)`

GetImageBase64Ok returns a tuple with the ImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase64

`func (o *OCRImageObject) SetImageBase64(v string)`

SetImageBase64 sets ImageBase64 field to given value.

### HasImageBase64

`func (o *OCRImageObject) HasImageBase64() bool`

HasImageBase64 returns a boolean if a field has been set.

### SetImageBase64Nil

`func (o *OCRImageObject) SetImageBase64Nil(b bool)`

 SetImageBase64Nil sets the value for ImageBase64 to be an explicit nil

### UnsetImageBase64
`func (o *OCRImageObject) UnsetImageBase64()`

UnsetImageBase64 ensures that no value is present for ImageBase64, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


