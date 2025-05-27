# MessageInputContentChunksInner

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

## Methods

### NewMessageInputContentChunksInner

`func NewMessageInputContentChunksInner(text string, imageUrl ImageUrl, tool BuiltInConnectors, fileId string, documentUrl string, ) *MessageInputContentChunksInner`

NewMessageInputContentChunksInner instantiates a new MessageInputContentChunksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageInputContentChunksInnerWithDefaults

`func NewMessageInputContentChunksInnerWithDefaults() *MessageInputContentChunksInner`

NewMessageInputContentChunksInnerWithDefaults instantiates a new MessageInputContentChunksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MessageInputContentChunksInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageInputContentChunksInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageInputContentChunksInner) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *MessageInputContentChunksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageInputContentChunksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageInputContentChunksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageInputContentChunksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageUrl

`func (o *MessageInputContentChunksInner) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *MessageInputContentChunksInner) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *MessageInputContentChunksInner) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.


### GetTool

`func (o *MessageInputContentChunksInner) GetTool() BuiltInConnectors`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *MessageInputContentChunksInner) GetToolOk() (*BuiltInConnectors, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *MessageInputContentChunksInner) SetTool(v BuiltInConnectors)`

SetTool sets Tool field to given value.


### GetFileId

`func (o *MessageInputContentChunksInner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MessageInputContentChunksInner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MessageInputContentChunksInner) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileName

`func (o *MessageInputContentChunksInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MessageInputContentChunksInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MessageInputContentChunksInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MessageInputContentChunksInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *MessageInputContentChunksInner) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *MessageInputContentChunksInner) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *MessageInputContentChunksInner) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *MessageInputContentChunksInner) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *MessageInputContentChunksInner) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *MessageInputContentChunksInner) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *MessageInputContentChunksInner) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *MessageInputContentChunksInner) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *MessageInputContentChunksInner) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *MessageInputContentChunksInner) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *MessageInputContentChunksInner) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


