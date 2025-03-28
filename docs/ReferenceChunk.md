# ReferenceChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceIds** | **[]int32** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "reference"]

## Methods

### NewReferenceChunk

`func NewReferenceChunk(referenceIds []int32, ) *ReferenceChunk`

NewReferenceChunk instantiates a new ReferenceChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceChunkWithDefaults

`func NewReferenceChunkWithDefaults() *ReferenceChunk`

NewReferenceChunkWithDefaults instantiates a new ReferenceChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceIds

`func (o *ReferenceChunk) GetReferenceIds() []int32`

GetReferenceIds returns the ReferenceIds field if non-nil, zero value otherwise.

### GetReferenceIdsOk

`func (o *ReferenceChunk) GetReferenceIdsOk() (*[]int32, bool)`

GetReferenceIdsOk returns a tuple with the ReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIds

`func (o *ReferenceChunk) SetReferenceIds(v []int32)`

SetReferenceIds sets ReferenceIds field to given value.


### GetType

`func (o *ReferenceChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReferenceChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReferenceChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReferenceChunk) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


