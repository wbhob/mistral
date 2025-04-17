# CompletionFTModelOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | **int32** |  | 
**OwnedBy** | **string** |  | 
**Root** | **string** |  | 
**Archived** | **bool** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Capabilities** | [**FTModelCapabilitiesOut**](FTModelCapabilitiesOut.md) |  | 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**Job** | **string** |  | 
**ModelType** | Pointer to **string** |  | [optional] [default to "completion"]

## Methods

### NewCompletionFTModelOut

`func NewCompletionFTModelOut(id string, created int32, ownedBy string, root string, archived bool, capabilities FTModelCapabilitiesOut, job string, ) *CompletionFTModelOut`

NewCompletionFTModelOut instantiates a new CompletionFTModelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionFTModelOutWithDefaults

`func NewCompletionFTModelOutWithDefaults() *CompletionFTModelOut`

NewCompletionFTModelOutWithDefaults instantiates a new CompletionFTModelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompletionFTModelOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompletionFTModelOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompletionFTModelOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CompletionFTModelOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompletionFTModelOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompletionFTModelOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompletionFTModelOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *CompletionFTModelOut) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CompletionFTModelOut) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CompletionFTModelOut) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *CompletionFTModelOut) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *CompletionFTModelOut) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *CompletionFTModelOut) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetRoot

`func (o *CompletionFTModelOut) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CompletionFTModelOut) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CompletionFTModelOut) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *CompletionFTModelOut) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CompletionFTModelOut) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CompletionFTModelOut) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *CompletionFTModelOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompletionFTModelOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompletionFTModelOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompletionFTModelOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CompletionFTModelOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompletionFTModelOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CompletionFTModelOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompletionFTModelOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompletionFTModelOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompletionFTModelOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CompletionFTModelOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CompletionFTModelOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilities

`func (o *CompletionFTModelOut) GetCapabilities() FTModelCapabilitiesOut`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CompletionFTModelOut) GetCapabilitiesOk() (*FTModelCapabilitiesOut, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CompletionFTModelOut) SetCapabilities(v FTModelCapabilitiesOut)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *CompletionFTModelOut) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *CompletionFTModelOut) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *CompletionFTModelOut) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *CompletionFTModelOut) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *CompletionFTModelOut) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CompletionFTModelOut) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CompletionFTModelOut) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CompletionFTModelOut) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *CompletionFTModelOut) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CompletionFTModelOut) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CompletionFTModelOut) SetJob(v string)`

SetJob sets Job field to given value.


### GetModelType

`func (o *CompletionFTModelOut) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *CompletionFTModelOut) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *CompletionFTModelOut) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *CompletionFTModelOut) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


