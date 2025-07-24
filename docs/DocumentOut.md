# DocumentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LibraryId** | **string** |  | 
**Hash** | **string** |  | 
**MimeType** | **string** |  | 
**Extension** | **string** |  | 
**Size** | **int32** |  | 
**Name** | **string** |  | 
**Summary** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**LastProcessedAt** | Pointer to **NullableTime** |  | [optional] 
**NumberOfPages** | Pointer to **NullableInt32** |  | [optional] 
**ProcessingStatus** | **string** |  | 
**UploadedById** | **string** |  | 
**UploadedByType** | **string** |  | 
**TokensProcessingMainContent** | Pointer to **NullableInt32** |  | [optional] 
**TokensProcessingSummary** | Pointer to **NullableInt32** |  | [optional] 
**TokensProcessingTotal** | **int32** |  | [readonly] 

## Methods

### NewDocumentOut

`func NewDocumentOut(id string, libraryId string, hash string, mimeType string, extension string, size int32, name string, createdAt time.Time, processingStatus string, uploadedById string, uploadedByType string, tokensProcessingTotal int32, ) *DocumentOut`

NewDocumentOut instantiates a new DocumentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentOutWithDefaults

`func NewDocumentOutWithDefaults() *DocumentOut`

NewDocumentOutWithDefaults instantiates a new DocumentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentOut) SetId(v string)`

SetId sets Id field to given value.


### GetLibraryId

`func (o *DocumentOut) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *DocumentOut) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *DocumentOut) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.


### GetHash

`func (o *DocumentOut) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DocumentOut) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DocumentOut) SetHash(v string)`

SetHash sets Hash field to given value.


### GetMimeType

`func (o *DocumentOut) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DocumentOut) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DocumentOut) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetExtension

`func (o *DocumentOut) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *DocumentOut) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *DocumentOut) SetExtension(v string)`

SetExtension sets Extension field to given value.


### GetSize

`func (o *DocumentOut) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DocumentOut) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DocumentOut) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *DocumentOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentOut) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *DocumentOut) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DocumentOut) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DocumentOut) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DocumentOut) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *DocumentOut) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *DocumentOut) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCreatedAt

`func (o *DocumentOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastProcessedAt

`func (o *DocumentOut) GetLastProcessedAt() time.Time`

GetLastProcessedAt returns the LastProcessedAt field if non-nil, zero value otherwise.

### GetLastProcessedAtOk

`func (o *DocumentOut) GetLastProcessedAtOk() (*time.Time, bool)`

GetLastProcessedAtOk returns a tuple with the LastProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedAt

`func (o *DocumentOut) SetLastProcessedAt(v time.Time)`

SetLastProcessedAt sets LastProcessedAt field to given value.

### HasLastProcessedAt

`func (o *DocumentOut) HasLastProcessedAt() bool`

HasLastProcessedAt returns a boolean if a field has been set.

### SetLastProcessedAtNil

`func (o *DocumentOut) SetLastProcessedAtNil(b bool)`

 SetLastProcessedAtNil sets the value for LastProcessedAt to be an explicit nil

### UnsetLastProcessedAt
`func (o *DocumentOut) UnsetLastProcessedAt()`

UnsetLastProcessedAt ensures that no value is present for LastProcessedAt, not even an explicit nil
### GetNumberOfPages

`func (o *DocumentOut) GetNumberOfPages() int32`

GetNumberOfPages returns the NumberOfPages field if non-nil, zero value otherwise.

### GetNumberOfPagesOk

`func (o *DocumentOut) GetNumberOfPagesOk() (*int32, bool)`

GetNumberOfPagesOk returns a tuple with the NumberOfPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPages

`func (o *DocumentOut) SetNumberOfPages(v int32)`

SetNumberOfPages sets NumberOfPages field to given value.

### HasNumberOfPages

`func (o *DocumentOut) HasNumberOfPages() bool`

HasNumberOfPages returns a boolean if a field has been set.

### SetNumberOfPagesNil

`func (o *DocumentOut) SetNumberOfPagesNil(b bool)`

 SetNumberOfPagesNil sets the value for NumberOfPages to be an explicit nil

### UnsetNumberOfPages
`func (o *DocumentOut) UnsetNumberOfPages()`

UnsetNumberOfPages ensures that no value is present for NumberOfPages, not even an explicit nil
### GetProcessingStatus

`func (o *DocumentOut) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *DocumentOut) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *DocumentOut) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.


### GetUploadedById

`func (o *DocumentOut) GetUploadedById() string`

GetUploadedById returns the UploadedById field if non-nil, zero value otherwise.

### GetUploadedByIdOk

`func (o *DocumentOut) GetUploadedByIdOk() (*string, bool)`

GetUploadedByIdOk returns a tuple with the UploadedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedById

`func (o *DocumentOut) SetUploadedById(v string)`

SetUploadedById sets UploadedById field to given value.


### GetUploadedByType

`func (o *DocumentOut) GetUploadedByType() string`

GetUploadedByType returns the UploadedByType field if non-nil, zero value otherwise.

### GetUploadedByTypeOk

`func (o *DocumentOut) GetUploadedByTypeOk() (*string, bool)`

GetUploadedByTypeOk returns a tuple with the UploadedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedByType

`func (o *DocumentOut) SetUploadedByType(v string)`

SetUploadedByType sets UploadedByType field to given value.


### GetTokensProcessingMainContent

`func (o *DocumentOut) GetTokensProcessingMainContent() int32`

GetTokensProcessingMainContent returns the TokensProcessingMainContent field if non-nil, zero value otherwise.

### GetTokensProcessingMainContentOk

`func (o *DocumentOut) GetTokensProcessingMainContentOk() (*int32, bool)`

GetTokensProcessingMainContentOk returns a tuple with the TokensProcessingMainContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingMainContent

`func (o *DocumentOut) SetTokensProcessingMainContent(v int32)`

SetTokensProcessingMainContent sets TokensProcessingMainContent field to given value.

### HasTokensProcessingMainContent

`func (o *DocumentOut) HasTokensProcessingMainContent() bool`

HasTokensProcessingMainContent returns a boolean if a field has been set.

### SetTokensProcessingMainContentNil

`func (o *DocumentOut) SetTokensProcessingMainContentNil(b bool)`

 SetTokensProcessingMainContentNil sets the value for TokensProcessingMainContent to be an explicit nil

### UnsetTokensProcessingMainContent
`func (o *DocumentOut) UnsetTokensProcessingMainContent()`

UnsetTokensProcessingMainContent ensures that no value is present for TokensProcessingMainContent, not even an explicit nil
### GetTokensProcessingSummary

`func (o *DocumentOut) GetTokensProcessingSummary() int32`

GetTokensProcessingSummary returns the TokensProcessingSummary field if non-nil, zero value otherwise.

### GetTokensProcessingSummaryOk

`func (o *DocumentOut) GetTokensProcessingSummaryOk() (*int32, bool)`

GetTokensProcessingSummaryOk returns a tuple with the TokensProcessingSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingSummary

`func (o *DocumentOut) SetTokensProcessingSummary(v int32)`

SetTokensProcessingSummary sets TokensProcessingSummary field to given value.

### HasTokensProcessingSummary

`func (o *DocumentOut) HasTokensProcessingSummary() bool`

HasTokensProcessingSummary returns a boolean if a field has been set.

### SetTokensProcessingSummaryNil

`func (o *DocumentOut) SetTokensProcessingSummaryNil(b bool)`

 SetTokensProcessingSummaryNil sets the value for TokensProcessingSummary to be an explicit nil

### UnsetTokensProcessingSummary
`func (o *DocumentOut) UnsetTokensProcessingSummary()`

UnsetTokensProcessingSummary ensures that no value is present for TokensProcessingSummary, not even an explicit nil
### GetTokensProcessingTotal

`func (o *DocumentOut) GetTokensProcessingTotal() int32`

GetTokensProcessingTotal returns the TokensProcessingTotal field if non-nil, zero value otherwise.

### GetTokensProcessingTotalOk

`func (o *DocumentOut) GetTokensProcessingTotalOk() (*int32, bool)`

GetTokensProcessingTotalOk returns a tuple with the TokensProcessingTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingTotal

`func (o *DocumentOut) SetTokensProcessingTotal(v int32)`

SetTokensProcessingTotal sets TokensProcessingTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


