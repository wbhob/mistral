# Response2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | **int32** |  | 
**OwnedBy** | **string** |  | 
**Root** | **string** |  | 
**Archived** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Capabilities** | [**FTModelCapabilitiesOut**](FTModelCapabilitiesOut.md) |  | 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**Job** | **string** |  | 
**ModelType** | Pointer to **string** |  | [optional] [default to "classifier"]
**ClassifierTargets** | [**[]ClassifierTargetOut**](ClassifierTargetOut.md) |  | 

## Methods

### NewResponse2

`func NewResponse2(id string, created int32, ownedBy string, root string, archived bool, capabilities FTModelCapabilitiesOut, job string, classifierTargets []ClassifierTargetOut, ) *Response2`

NewResponse2 instantiates a new Response2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse2WithDefaults

`func NewResponse2WithDefaults() *Response2`

NewResponse2WithDefaults instantiates a new Response2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Response2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Response2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Response2) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Response2) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Response2) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Response2) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Response2) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *Response2) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Response2) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Response2) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *Response2) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Response2) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Response2) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetRoot

`func (o *Response2) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Response2) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Response2) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *Response2) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Response2) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Response2) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *Response2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Response2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Response2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Response2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Response2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Response2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Response2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Response2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCapabilities

`func (o *Response2) GetCapabilities() FTModelCapabilitiesOut`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Response2) GetCapabilitiesOk() (*FTModelCapabilitiesOut, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Response2) SetCapabilities(v FTModelCapabilitiesOut)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *Response2) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *Response2) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *Response2) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *Response2) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *Response2) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Response2) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Response2) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Response2) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *Response2) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Response2) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Response2) SetJob(v string)`

SetJob sets Job field to given value.


### GetModelType

`func (o *Response2) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *Response2) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *Response2) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *Response2) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetClassifierTargets

`func (o *Response2) GetClassifierTargets() []ClassifierTargetOut`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *Response2) GetClassifierTargetsOk() (*[]ClassifierTargetOut, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *Response2) SetClassifierTargets(v []ClassifierTargetOut)`

SetClassifierTargets sets ClassifierTargets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


