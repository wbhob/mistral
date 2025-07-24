# FTModelCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | Pointer to **int32** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] [default to "mistralai"]
**Capabilities** | [**ModelCapabilities**](ModelCapabilities.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**Deprecation** | Pointer to **NullableTime** |  | [optional] 
**DeprecationReplacementModel** | Pointer to **NullableString** |  | [optional] 
**DefaultModelTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "fine-tuned"]
**Job** | **string** |  | 
**Root** | **string** |  | 
**Archived** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewFTModelCard

`func NewFTModelCard(id string, capabilities ModelCapabilities, job string, root string, ) *FTModelCard`

NewFTModelCard instantiates a new FTModelCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTModelCardWithDefaults

`func NewFTModelCardWithDefaults() *FTModelCard`

NewFTModelCardWithDefaults instantiates a new FTModelCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FTModelCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FTModelCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FTModelCard) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *FTModelCard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FTModelCard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FTModelCard) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FTModelCard) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *FTModelCard) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FTModelCard) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FTModelCard) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FTModelCard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FTModelCard) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FTModelCard) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FTModelCard) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FTModelCard) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetCapabilities

`func (o *FTModelCard) GetCapabilities() ModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *FTModelCard) GetCapabilitiesOk() (*ModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *FTModelCard) SetCapabilities(v ModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetName

`func (o *FTModelCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FTModelCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FTModelCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FTModelCard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FTModelCard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FTModelCard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FTModelCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FTModelCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FTModelCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FTModelCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FTModelCard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FTModelCard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxContextLength

`func (o *FTModelCard) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *FTModelCard) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *FTModelCard) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *FTModelCard) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *FTModelCard) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *FTModelCard) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *FTModelCard) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *FTModelCard) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDeprecation

`func (o *FTModelCard) GetDeprecation() time.Time`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *FTModelCard) GetDeprecationOk() (*time.Time, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *FTModelCard) SetDeprecation(v time.Time)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *FTModelCard) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### SetDeprecationNil

`func (o *FTModelCard) SetDeprecationNil(b bool)`

 SetDeprecationNil sets the value for Deprecation to be an explicit nil

### UnsetDeprecation
`func (o *FTModelCard) UnsetDeprecation()`

UnsetDeprecation ensures that no value is present for Deprecation, not even an explicit nil
### GetDeprecationReplacementModel

`func (o *FTModelCard) GetDeprecationReplacementModel() string`

GetDeprecationReplacementModel returns the DeprecationReplacementModel field if non-nil, zero value otherwise.

### GetDeprecationReplacementModelOk

`func (o *FTModelCard) GetDeprecationReplacementModelOk() (*string, bool)`

GetDeprecationReplacementModelOk returns a tuple with the DeprecationReplacementModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationReplacementModel

`func (o *FTModelCard) SetDeprecationReplacementModel(v string)`

SetDeprecationReplacementModel sets DeprecationReplacementModel field to given value.

### HasDeprecationReplacementModel

`func (o *FTModelCard) HasDeprecationReplacementModel() bool`

HasDeprecationReplacementModel returns a boolean if a field has been set.

### SetDeprecationReplacementModelNil

`func (o *FTModelCard) SetDeprecationReplacementModelNil(b bool)`

 SetDeprecationReplacementModelNil sets the value for DeprecationReplacementModel to be an explicit nil

### UnsetDeprecationReplacementModel
`func (o *FTModelCard) UnsetDeprecationReplacementModel()`

UnsetDeprecationReplacementModel ensures that no value is present for DeprecationReplacementModel, not even an explicit nil
### GetDefaultModelTemperature

`func (o *FTModelCard) GetDefaultModelTemperature() float32`

GetDefaultModelTemperature returns the DefaultModelTemperature field if non-nil, zero value otherwise.

### GetDefaultModelTemperatureOk

`func (o *FTModelCard) GetDefaultModelTemperatureOk() (*float32, bool)`

GetDefaultModelTemperatureOk returns a tuple with the DefaultModelTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelTemperature

`func (o *FTModelCard) SetDefaultModelTemperature(v float32)`

SetDefaultModelTemperature sets DefaultModelTemperature field to given value.

### HasDefaultModelTemperature

`func (o *FTModelCard) HasDefaultModelTemperature() bool`

HasDefaultModelTemperature returns a boolean if a field has been set.

### SetDefaultModelTemperatureNil

`func (o *FTModelCard) SetDefaultModelTemperatureNil(b bool)`

 SetDefaultModelTemperatureNil sets the value for DefaultModelTemperature to be an explicit nil

### UnsetDefaultModelTemperature
`func (o *FTModelCard) UnsetDefaultModelTemperature()`

UnsetDefaultModelTemperature ensures that no value is present for DefaultModelTemperature, not even an explicit nil
### GetType

`func (o *FTModelCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FTModelCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FTModelCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FTModelCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJob

`func (o *FTModelCard) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *FTModelCard) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *FTModelCard) SetJob(v string)`

SetJob sets Job field to given value.


### GetRoot

`func (o *FTModelCard) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *FTModelCard) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *FTModelCard) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *FTModelCard) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FTModelCard) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FTModelCard) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *FTModelCard) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


