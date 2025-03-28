# OCRUsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagesProcessed** | **int32** | Number of pages processed | 
**DocSizeBytes** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewOCRUsageInfo

`func NewOCRUsageInfo(pagesProcessed int32, ) *OCRUsageInfo`

NewOCRUsageInfo instantiates a new OCRUsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRUsageInfoWithDefaults

`func NewOCRUsageInfoWithDefaults() *OCRUsageInfo`

NewOCRUsageInfoWithDefaults instantiates a new OCRUsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagesProcessed

`func (o *OCRUsageInfo) GetPagesProcessed() int32`

GetPagesProcessed returns the PagesProcessed field if non-nil, zero value otherwise.

### GetPagesProcessedOk

`func (o *OCRUsageInfo) GetPagesProcessedOk() (*int32, bool)`

GetPagesProcessedOk returns a tuple with the PagesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesProcessed

`func (o *OCRUsageInfo) SetPagesProcessed(v int32)`

SetPagesProcessed sets PagesProcessed field to given value.


### GetDocSizeBytes

`func (o *OCRUsageInfo) GetDocSizeBytes() int32`

GetDocSizeBytes returns the DocSizeBytes field if non-nil, zero value otherwise.

### GetDocSizeBytesOk

`func (o *OCRUsageInfo) GetDocSizeBytesOk() (*int32, bool)`

GetDocSizeBytesOk returns a tuple with the DocSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocSizeBytes

`func (o *OCRUsageInfo) SetDocSizeBytes(v int32)`

SetDocSizeBytes sets DocSizeBytes field to given value.

### HasDocSizeBytes

`func (o *OCRUsageInfo) HasDocSizeBytes() bool`

HasDocSizeBytes returns a boolean if a field has been set.

### SetDocSizeBytesNil

`func (o *OCRUsageInfo) SetDocSizeBytesNil(b bool)`

 SetDocSizeBytesNil sets the value for DocSizeBytes to be an explicit nil

### UnsetDocSizeBytes
`func (o *OCRUsageInfo) UnsetDocSizeBytes()`

UnsetDocSizeBytes ensures that no value is present for DocSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


