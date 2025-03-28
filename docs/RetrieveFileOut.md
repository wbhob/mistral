# RetrieveFileOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the file. | 
**Object** | **string** | The object type, which is always \&quot;file\&quot;. | 
**Bytes** | **int32** | The size of the file, in bytes. | 
**CreatedAt** | **int32** | The UNIX timestamp (in seconds) of the event. | 
**Filename** | **string** | The name of the uploaded file. | 
**Purpose** | [**FilePurpose**](FilePurpose.md) | The intended purpose of the uploaded file. Only accepts fine-tuning (&#x60;fine-tune&#x60;) for now. | 
**SampleType** | [**SampleType**](SampleType.md) |  | 
**NumLines** | Pointer to **NullableInt32** |  | [optional] 
**Source** | [**Source**](Source.md) |  | 
**Deleted** | **bool** |  | 

## Methods

### NewRetrieveFileOut

`func NewRetrieveFileOut(id string, object string, bytes int32, createdAt int32, filename string, purpose FilePurpose, sampleType SampleType, source Source, deleted bool, ) *RetrieveFileOut`

NewRetrieveFileOut instantiates a new RetrieveFileOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveFileOutWithDefaults

`func NewRetrieveFileOutWithDefaults() *RetrieveFileOut`

NewRetrieveFileOutWithDefaults instantiates a new RetrieveFileOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetrieveFileOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetrieveFileOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetrieveFileOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *RetrieveFileOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RetrieveFileOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RetrieveFileOut) SetObject(v string)`

SetObject sets Object field to given value.


### GetBytes

`func (o *RetrieveFileOut) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *RetrieveFileOut) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *RetrieveFileOut) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *RetrieveFileOut) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RetrieveFileOut) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RetrieveFileOut) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetFilename

`func (o *RetrieveFileOut) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RetrieveFileOut) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RetrieveFileOut) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPurpose

`func (o *RetrieveFileOut) GetPurpose() FilePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RetrieveFileOut) GetPurposeOk() (*FilePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RetrieveFileOut) SetPurpose(v FilePurpose)`

SetPurpose sets Purpose field to given value.


### GetSampleType

`func (o *RetrieveFileOut) GetSampleType() SampleType`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *RetrieveFileOut) GetSampleTypeOk() (*SampleType, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *RetrieveFileOut) SetSampleType(v SampleType)`

SetSampleType sets SampleType field to given value.


### GetNumLines

`func (o *RetrieveFileOut) GetNumLines() int32`

GetNumLines returns the NumLines field if non-nil, zero value otherwise.

### GetNumLinesOk

`func (o *RetrieveFileOut) GetNumLinesOk() (*int32, bool)`

GetNumLinesOk returns a tuple with the NumLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLines

`func (o *RetrieveFileOut) SetNumLines(v int32)`

SetNumLines sets NumLines field to given value.

### HasNumLines

`func (o *RetrieveFileOut) HasNumLines() bool`

HasNumLines returns a boolean if a field has been set.

### SetNumLinesNil

`func (o *RetrieveFileOut) SetNumLinesNil(b bool)`

 SetNumLinesNil sets the value for NumLines to be an explicit nil

### UnsetNumLines
`func (o *RetrieveFileOut) UnsetNumLines()`

UnsetNumLines ensures that no value is present for NumLines, not even an explicit nil
### GetSource

`func (o *RetrieveFileOut) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RetrieveFileOut) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RetrieveFileOut) SetSource(v Source)`

SetSource sets Source field to given value.


### GetDeleted

`func (o *RetrieveFileOut) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *RetrieveFileOut) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *RetrieveFileOut) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


