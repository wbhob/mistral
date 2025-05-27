# ToolFileChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool_file"]
**Tool** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**FileId** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 
**FileType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewToolFileChunk

`func NewToolFileChunk(tool BuiltInConnectors, fileId string, ) *ToolFileChunk`

NewToolFileChunk instantiates a new ToolFileChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolFileChunkWithDefaults

`func NewToolFileChunkWithDefaults() *ToolFileChunk`

NewToolFileChunkWithDefaults instantiates a new ToolFileChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolFileChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolFileChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolFileChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolFileChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTool

`func (o *ToolFileChunk) GetTool() BuiltInConnectors`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *ToolFileChunk) GetToolOk() (*BuiltInConnectors, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *ToolFileChunk) SetTool(v BuiltInConnectors)`

SetTool sets Tool field to given value.


### GetFileId

`func (o *ToolFileChunk) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ToolFileChunk) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ToolFileChunk) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileName

`func (o *ToolFileChunk) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ToolFileChunk) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ToolFileChunk) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ToolFileChunk) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ToolFileChunk) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ToolFileChunk) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileType

`func (o *ToolFileChunk) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ToolFileChunk) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ToolFileChunk) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ToolFileChunk) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *ToolFileChunk) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *ToolFileChunk) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


