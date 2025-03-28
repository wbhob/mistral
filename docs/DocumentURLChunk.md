# DocumentURLChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentUrl** | **string** |  | 
**DocumentName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "document_url"]

## Methods

### NewDocumentURLChunk

`func NewDocumentURLChunk(documentUrl string, ) *DocumentURLChunk`

NewDocumentURLChunk instantiates a new DocumentURLChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentURLChunkWithDefaults

`func NewDocumentURLChunkWithDefaults() *DocumentURLChunk`

NewDocumentURLChunkWithDefaults instantiates a new DocumentURLChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentUrl

`func (o *DocumentURLChunk) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *DocumentURLChunk) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *DocumentURLChunk) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *DocumentURLChunk) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *DocumentURLChunk) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *DocumentURLChunk) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *DocumentURLChunk) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### SetDocumentNameNil

`func (o *DocumentURLChunk) SetDocumentNameNil(b bool)`

 SetDocumentNameNil sets the value for DocumentName to be an explicit nil

### UnsetDocumentName
`func (o *DocumentURLChunk) UnsetDocumentName()`

UnsetDocumentName ensures that no value is present for DocumentName, not even an explicit nil
### GetType

`func (o *DocumentURLChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentURLChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentURLChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentURLChunk) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


