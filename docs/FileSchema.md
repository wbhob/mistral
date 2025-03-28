# FileSchema

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

## Methods

### NewFileSchema

`func NewFileSchema(id string, object string, bytes int32, createdAt int32, filename string, purpose FilePurpose, sampleType SampleType, source Source, ) *FileSchema`

NewFileSchema instantiates a new FileSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSchemaWithDefaults

`func NewFileSchemaWithDefaults() *FileSchema`

NewFileSchemaWithDefaults instantiates a new FileSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSchema) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *FileSchema) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FileSchema) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FileSchema) SetObject(v string)`

SetObject sets Object field to given value.


### GetBytes

`func (o *FileSchema) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *FileSchema) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *FileSchema) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *FileSchema) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileSchema) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileSchema) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetFilename

`func (o *FileSchema) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileSchema) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileSchema) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPurpose

`func (o *FileSchema) GetPurpose() FilePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *FileSchema) GetPurposeOk() (*FilePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *FileSchema) SetPurpose(v FilePurpose)`

SetPurpose sets Purpose field to given value.


### GetSampleType

`func (o *FileSchema) GetSampleType() SampleType`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *FileSchema) GetSampleTypeOk() (*SampleType, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *FileSchema) SetSampleType(v SampleType)`

SetSampleType sets SampleType field to given value.


### GetNumLines

`func (o *FileSchema) GetNumLines() int32`

GetNumLines returns the NumLines field if non-nil, zero value otherwise.

### GetNumLinesOk

`func (o *FileSchema) GetNumLinesOk() (*int32, bool)`

GetNumLinesOk returns a tuple with the NumLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLines

`func (o *FileSchema) SetNumLines(v int32)`

SetNumLines sets NumLines field to given value.

### HasNumLines

`func (o *FileSchema) HasNumLines() bool`

HasNumLines returns a boolean if a field has been set.

### SetNumLinesNil

`func (o *FileSchema) SetNumLinesNil(b bool)`

 SetNumLinesNil sets the value for NumLines to be an explicit nil

### UnsetNumLines
`func (o *FileSchema) UnsetNumLines()`

UnsetNumLines ensures that no value is present for NumLines, not even an explicit nil
### GetSource

`func (o *FileSchema) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FileSchema) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FileSchema) SetSource(v Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


