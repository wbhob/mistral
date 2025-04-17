# ClassifierFTModelOut

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
**ClassifierTargets** | [**[]ClassifierTargetOut**](ClassifierTargetOut.md) |  | 
**ModelType** | Pointer to **string** |  | [optional] [default to "classifier"]

## Methods

### NewClassifierFTModelOut

`func NewClassifierFTModelOut(id string, created int32, ownedBy string, root string, archived bool, capabilities FTModelCapabilitiesOut, job string, classifierTargets []ClassifierTargetOut, ) *ClassifierFTModelOut`

NewClassifierFTModelOut instantiates a new ClassifierFTModelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierFTModelOutWithDefaults

`func NewClassifierFTModelOutWithDefaults() *ClassifierFTModelOut`

NewClassifierFTModelOutWithDefaults instantiates a new ClassifierFTModelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassifierFTModelOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassifierFTModelOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassifierFTModelOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ClassifierFTModelOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ClassifierFTModelOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ClassifierFTModelOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ClassifierFTModelOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ClassifierFTModelOut) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClassifierFTModelOut) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClassifierFTModelOut) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *ClassifierFTModelOut) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ClassifierFTModelOut) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ClassifierFTModelOut) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetRoot

`func (o *ClassifierFTModelOut) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ClassifierFTModelOut) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ClassifierFTModelOut) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *ClassifierFTModelOut) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ClassifierFTModelOut) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ClassifierFTModelOut) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *ClassifierFTModelOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassifierFTModelOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassifierFTModelOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClassifierFTModelOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClassifierFTModelOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClassifierFTModelOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ClassifierFTModelOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClassifierFTModelOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClassifierFTModelOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClassifierFTModelOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClassifierFTModelOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClassifierFTModelOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilities

`func (o *ClassifierFTModelOut) GetCapabilities() FTModelCapabilitiesOut`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ClassifierFTModelOut) GetCapabilitiesOk() (*FTModelCapabilitiesOut, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ClassifierFTModelOut) SetCapabilities(v FTModelCapabilitiesOut)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *ClassifierFTModelOut) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *ClassifierFTModelOut) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *ClassifierFTModelOut) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *ClassifierFTModelOut) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *ClassifierFTModelOut) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ClassifierFTModelOut) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ClassifierFTModelOut) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ClassifierFTModelOut) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *ClassifierFTModelOut) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ClassifierFTModelOut) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ClassifierFTModelOut) SetJob(v string)`

SetJob sets Job field to given value.


### GetClassifierTargets

`func (o *ClassifierFTModelOut) GetClassifierTargets() []ClassifierTargetOut`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *ClassifierFTModelOut) GetClassifierTargetsOk() (*[]ClassifierTargetOut, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *ClassifierFTModelOut) SetClassifierTargets(v []ClassifierTargetOut)`

SetClassifierTargets sets ClassifierTargets field to given value.


### GetModelType

`func (o *ClassifierFTModelOut) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ClassifierFTModelOut) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ClassifierFTModelOut) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ClassifierFTModelOut) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


