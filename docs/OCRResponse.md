# OCRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | [**[]OCRPageObject**](OCRPageObject.md) | List of OCR info for pages. | 
**Model** | **string** | The model used to generate the OCR. | 
**UsageInfo** | [**OCRUsageInfo**](OCRUsageInfo.md) | Usage info for the OCR request. | 

## Methods

### NewOCRResponse

`func NewOCRResponse(pages []OCRPageObject, model string, usageInfo OCRUsageInfo, ) *OCRResponse`

NewOCRResponse instantiates a new OCRResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRResponseWithDefaults

`func NewOCRResponseWithDefaults() *OCRResponse`

NewOCRResponseWithDefaults instantiates a new OCRResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *OCRResponse) GetPages() []OCRPageObject`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OCRResponse) GetPagesOk() (*[]OCRPageObject, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OCRResponse) SetPages(v []OCRPageObject)`

SetPages sets Pages field to given value.


### GetModel

`func (o *OCRResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OCRResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OCRResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetUsageInfo

`func (o *OCRResponse) GetUsageInfo() OCRUsageInfo`

GetUsageInfo returns the UsageInfo field if non-nil, zero value otherwise.

### GetUsageInfoOk

`func (o *OCRResponse) GetUsageInfoOk() (*OCRUsageInfo, bool)`

GetUsageInfoOk returns a tuple with the UsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInfo

`func (o *OCRResponse) SetUsageInfo(v OCRUsageInfo)`

SetUsageInfo sets UsageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


