# ContentChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**ImageUrl** | [**ImageUrl**](ImageUrl.md) |  | 
**DocumentUrl** | **string** |  | 
**DocumentName** | Pointer to **string** |  | [optional] 
**ReferenceIds** | **[]int32** |  | 

## Methods

### NewContentChunk

`func NewContentChunk(text string, imageUrl ImageUrl, documentUrl string, referenceIds []int32, ) *ContentChunk`

NewContentChunk instantiates a new ContentChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentChunkWithDefaults

`func NewContentChunkWithDefaults() *ContentChunk`

NewContentChunkWithDefaults instantiates a new ContentChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ContentChunk) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ContentChunk) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ContentChunk) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ContentChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageUrl

`func (o *ContentChunk) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ContentChunk) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ContentChunk) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.


### GetDocumentUrl

`func (o *ContentChunk) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *ContentChunk) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *ContentChunk) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *ContentChunk) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *ContentChunk) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *ContentChunk) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *ContentChunk) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetReferenceIds

`func (o *ContentChunk) GetReferenceIds() []int32`

GetReferenceIds returns the ReferenceIds field if non-nil, zero value otherwise.

### GetReferenceIdsOk

`func (o *ContentChunk) GetReferenceIdsOk() (*[]int32, bool)`

GetReferenceIdsOk returns a tuple with the ReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIds

`func (o *ContentChunk) SetReferenceIds(v []int32)`

SetReferenceIds sets ReferenceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


