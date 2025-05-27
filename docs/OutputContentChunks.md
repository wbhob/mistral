# OutputContentChunks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**ImageUrl** | [**ImageUrl**](ImageUrl.md) |  | 
**Tool** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**FileId** | **string** |  | 
**FileName** | Pointer to **string** |  | [optional] 
**FileType** | Pointer to **string** |  | [optional] 
**DocumentUrl** | **string** |  | 
**DocumentName** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewOutputContentChunks

`func NewOutputContentChunks(text string, imageUrl ImageUrl, tool BuiltInConnectors, fileId string, documentUrl string, title string, ) *OutputContentChunks`

NewOutputContentChunks instantiates a new OutputContentChunks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputContentChunksWithDefaults

`func NewOutputContentChunksWithDefaults() *OutputContentChunks`

NewOutputContentChunksWithDefaults instantiates a new OutputContentChunks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *OutputContentChunks) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OutputContentChunks) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OutputContentChunks) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *OutputContentChunks) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputContentChunks) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputContentChunks) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OutputContentChunks) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageUrl

`func (o *OutputContentChunks) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *OutputContentChunks) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *OutputContentChunks) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.


### GetTool

`func (o *OutputContentChunks) GetTool() BuiltInConnectors`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *OutputContentChunks) GetToolOk() (*BuiltInConnectors, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *OutputContentChunks) SetTool(v BuiltInConnectors)`

SetTool sets Tool field to given value.


### GetFileId

`func (o *OutputContentChunks) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *OutputContentChunks) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *OutputContentChunks) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileName

`func (o *OutputContentChunks) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *OutputContentChunks) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *OutputContentChunks) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *OutputContentChunks) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *OutputContentChunks) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *OutputContentChunks) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *OutputContentChunks) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *OutputContentChunks) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *OutputContentChunks) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *OutputContentChunks) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *OutputContentChunks) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *OutputContentChunks) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *OutputContentChunks) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *OutputContentChunks) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *OutputContentChunks) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetTitle

`func (o *OutputContentChunks) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutputContentChunks) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutputContentChunks) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *OutputContentChunks) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputContentChunks) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputContentChunks) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OutputContentChunks) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSource

`func (o *OutputContentChunks) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OutputContentChunks) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OutputContentChunks) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OutputContentChunks) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


