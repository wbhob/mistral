# ToolReferenceChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool_reference"]
**Tool** | [**BuiltInConnectors**](BuiltInConnectors.md) |  | 
**Title** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Favicon** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

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
### GetFavicon

`func (o *ToolReferenceChunk) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *ToolReferenceChunk) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *ToolReferenceChunk) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *ToolReferenceChunk) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *ToolReferenceChunk) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *ToolReferenceChunk) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetDescription

`func (o *ToolReferenceChunk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolReferenceChunk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolReferenceChunk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolReferenceChunk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolReferenceChunk) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolReferenceChunk) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


