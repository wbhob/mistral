# ImageURLChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrl** | [**ImageUrl**](ImageUrl.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "image_url"]

## Methods

### NewImageURLChunk

`func NewImageURLChunk(imageUrl ImageUrl, ) *ImageURLChunk`

NewImageURLChunk instantiates a new ImageURLChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageURLChunkWithDefaults

`func NewImageURLChunkWithDefaults() *ImageURLChunk`

NewImageURLChunkWithDefaults instantiates a new ImageURLChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrl

`func (o *ImageURLChunk) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ImageURLChunk) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ImageURLChunk) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.


### GetType

`func (o *ImageURLChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageURLChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageURLChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageURLChunk) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


