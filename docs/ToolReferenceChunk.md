# ToolReferenceChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool_reference"]
**Tool** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Title** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewToolReferenceChunk

`func NewToolReferenceChunk(tool BuiltInConnectors, title string, ) *ToolReferenceChunk`

NewToolReferenceChunk instantiates a new ToolReferenceChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolReferenceChunkWithDefaults

`func NewToolReferenceChunkWithDefaults() *ToolReferenceChunk`

NewToolReferenceChunkWithDefaults instantiates a new ToolReferenceChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolReferenceChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolReferenceChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolReferenceChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolReferenceChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTool

`func (o *ToolReferenceChunk) GetTool() BuiltInConnectors`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *ToolReferenceChunk) GetToolOk() (*BuiltInConnectors, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *ToolReferenceChunk) SetTool(v BuiltInConnectors)`

SetTool sets Tool field to given value.


### GetTitle

`func (o *ToolReferenceChunk) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ToolReferenceChunk) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ToolReferenceChunk) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *ToolReferenceChunk) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolReferenceChunk) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolReferenceChunk) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolReferenceChunk) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ToolReferenceChunk) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ToolReferenceChunk) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSource

`func (o *ToolReferenceChunk) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolReferenceChunk) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolReferenceChunk) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolReferenceChunk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ToolReferenceChunk) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ToolReferenceChunk) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


