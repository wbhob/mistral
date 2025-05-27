# OCRRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **NullableString** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Document** | [**Document**](Document.md) |  | 
**Pages** | Pointer to **[]int32** |  | [optional] 
**IncludeImageBase64** | Pointer to **NullableBool** |  | [optional] 
**ImageLimit** | Pointer to **NullableInt32** |  | [optional] 
**ImageMinSize** | Pointer to **NullableInt32** |  | [optional] 
**BboxAnnotationFormat** | Pointer to [**NullableResponseFormat**](ResponseFormat.md) |  | [optional] 
**DocumentAnnotationFormat** | Pointer to [**NullableResponseFormat**](ResponseFormat.md) |  | [optional] 

## Methods

### NewOCRRequest

`func NewOCRRequest(model NullableString, document Document, ) *OCRRequest`

NewOCRRequest instantiates a new OCRRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRRequestWithDefaults

`func NewOCRRequestWithDefaults() *OCRRequest`

NewOCRRequestWithDefaults instantiates a new OCRRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *OCRRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OCRRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OCRRequest) SetModel(v string)`

SetModel sets Model field to given value.


### SetModelNil

`func (o *OCRRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *OCRRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetId

`func (o *OCRRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OCRRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OCRRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OCRRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocument

`func (o *OCRRequest) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *OCRRequest) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *OCRRequest) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetPages

`func (o *OCRRequest) GetPages() []int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OCRRequest) GetPagesOk() (*[]int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OCRRequest) SetPages(v []int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OCRRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPagesNil

`func (o *OCRRequest) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *OCRRequest) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil
### GetIncludeImageBase64

`func (o *OCRRequest) GetIncludeImageBase64() bool`

GetIncludeImageBase64 returns the IncludeImageBase64 field if non-nil, zero value otherwise.

### GetIncludeImageBase64Ok

`func (o *OCRRequest) GetIncludeImageBase64Ok() (*bool, bool)`

GetIncludeImageBase64Ok returns a tuple with the IncludeImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeImageBase64

`func (o *OCRRequest) SetIncludeImageBase64(v bool)`

SetIncludeImageBase64 sets IncludeImageBase64 field to given value.

### HasIncludeImageBase64

`func (o *OCRRequest) HasIncludeImageBase64() bool`

HasIncludeImageBase64 returns a boolean if a field has been set.

### SetIncludeImageBase64Nil

`func (o *OCRRequest) SetIncludeImageBase64Nil(b bool)`

 SetIncludeImageBase64Nil sets the value for IncludeImageBase64 to be an explicit nil

### UnsetIncludeImageBase64
`func (o *OCRRequest) UnsetIncludeImageBase64()`

UnsetIncludeImageBase64 ensures that no value is present for IncludeImageBase64, not even an explicit nil
### GetImageLimit

`func (o *OCRRequest) GetImageLimit() int32`

GetImageLimit returns the ImageLimit field if non-nil, zero value otherwise.

### GetImageLimitOk

`func (o *OCRRequest) GetImageLimitOk() (*int32, bool)`

GetImageLimitOk returns a tuple with the ImageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLimit

`func (o *OCRRequest) SetImageLimit(v int32)`

SetImageLimit sets ImageLimit field to given value.

### HasImageLimit

`func (o *OCRRequest) HasImageLimit() bool`

HasImageLimit returns a boolean if a field has been set.

### SetImageLimitNil

`func (o *OCRRequest) SetImageLimitNil(b bool)`

 SetImageLimitNil sets the value for ImageLimit to be an explicit nil

### UnsetImageLimit
`func (o *OCRRequest) UnsetImageLimit()`

UnsetImageLimit ensures that no value is present for ImageLimit, not even an explicit nil
### GetImageMinSize

`func (o *OCRRequest) GetImageMinSize() int32`

GetImageMinSize returns the ImageMinSize field if non-nil, zero value otherwise.

### GetImageMinSizeOk

`func (o *OCRRequest) GetImageMinSizeOk() (*int32, bool)`

GetImageMinSizeOk returns a tuple with the ImageMinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMinSize

`func (o *OCRRequest) SetImageMinSize(v int32)`

SetImageMinSize sets ImageMinSize field to given value.

### HasImageMinSize

`func (o *OCRRequest) HasImageMinSize() bool`

HasImageMinSize returns a boolean if a field has been set.

### SetImageMinSizeNil

`func (o *OCRRequest) SetImageMinSizeNil(b bool)`

 SetImageMinSizeNil sets the value for ImageMinSize to be an explicit nil

### UnsetImageMinSize
`func (o *OCRRequest) UnsetImageMinSize()`

UnsetImageMinSize ensures that no value is present for ImageMinSize, not even an explicit nil
### GetBboxAnnotationFormat

`func (o *OCRRequest) GetBboxAnnotationFormat() ResponseFormat`

GetBboxAnnotationFormat returns the BboxAnnotationFormat field if non-nil, zero value otherwise.

### GetBboxAnnotationFormatOk

`func (o *OCRRequest) GetBboxAnnotationFormatOk() (*ResponseFormat, bool)`

GetBboxAnnotationFormatOk returns a tuple with the BboxAnnotationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBboxAnnotationFormat

`func (o *OCRRequest) SetBboxAnnotationFormat(v ResponseFormat)`

SetBboxAnnotationFormat sets BboxAnnotationFormat field to given value.

### HasBboxAnnotationFormat

`func (o *OCRRequest) HasBboxAnnotationFormat() bool`

HasBboxAnnotationFormat returns a boolean if a field has been set.

### SetBboxAnnotationFormatNil

`func (o *OCRRequest) SetBboxAnnotationFormatNil(b bool)`

 SetBboxAnnotationFormatNil sets the value for BboxAnnotationFormat to be an explicit nil

### UnsetBboxAnnotationFormat
`func (o *OCRRequest) UnsetBboxAnnotationFormat()`

UnsetBboxAnnotationFormat ensures that no value is present for BboxAnnotationFormat, not even an explicit nil
### GetDocumentAnnotationFormat

`func (o *OCRRequest) GetDocumentAnnotationFormat() ResponseFormat`

GetDocumentAnnotationFormat returns the DocumentAnnotationFormat field if non-nil, zero value otherwise.

### GetDocumentAnnotationFormatOk

`func (o *OCRRequest) GetDocumentAnnotationFormatOk() (*ResponseFormat, bool)`

GetDocumentAnnotationFormatOk returns a tuple with the DocumentAnnotationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentAnnotationFormat

`func (o *OCRRequest) SetDocumentAnnotationFormat(v ResponseFormat)`

SetDocumentAnnotationFormat sets DocumentAnnotationFormat field to given value.

### HasDocumentAnnotationFormat

`func (o *OCRRequest) HasDocumentAnnotationFormat() bool`

HasDocumentAnnotationFormat returns a boolean if a field has been set.

### SetDocumentAnnotationFormatNil

`func (o *OCRRequest) SetDocumentAnnotationFormatNil(b bool)`

 SetDocumentAnnotationFormatNil sets the value for DocumentAnnotationFormat to be an explicit nil

### UnsetDocumentAnnotationFormat
`func (o *OCRRequest) UnsetDocumentAnnotationFormat()`

UnsetDocumentAnnotationFormat ensures that no value is present for DocumentAnnotationFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


