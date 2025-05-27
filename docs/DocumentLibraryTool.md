# DocumentLibraryTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "document_library"]
**LibraryIds** | **[]string** | Ids of the library in which to search. | 

## Methods

### NewDocumentLibraryTool

`func NewDocumentLibraryTool(libraryIds []string, ) *DocumentLibraryTool`

NewDocumentLibraryTool instantiates a new DocumentLibraryTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentLibraryToolWithDefaults

`func NewDocumentLibraryToolWithDefaults() *DocumentLibraryTool`

NewDocumentLibraryToolWithDefaults instantiates a new DocumentLibraryTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DocumentLibraryTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentLibraryTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentLibraryTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentLibraryTool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLibraryIds

`func (o *DocumentLibraryTool) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *DocumentLibraryTool) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *DocumentLibraryTool) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


