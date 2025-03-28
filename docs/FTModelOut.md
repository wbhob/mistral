# FTModelOut

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

## Methods

### NewFTModelOut

`func NewFTModelOut(id string, created int32, ownedBy string, root string, archived bool, capabilities FTModelCapabilitiesOut, job string, ) *FTModelOut`

NewFTModelOut instantiates a new FTModelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTModelOutWithDefaults

`func NewFTModelOutWithDefaults() *FTModelOut`

NewFTModelOutWithDefaults instantiates a new FTModelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FTModelOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FTModelOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FTModelOut) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *FTModelOut) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FTModelOut) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FTModelOut) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FTModelOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *FTModelOut) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FTModelOut) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FTModelOut) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *FTModelOut) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FTModelOut) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FTModelOut) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetRoot

`func (o *FTModelOut) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *FTModelOut) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *FTModelOut) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *FTModelOut) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FTModelOut) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FTModelOut) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *FTModelOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FTModelOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FTModelOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FTModelOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FTModelOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FTModelOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FTModelOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FTModelOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FTModelOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FTModelOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FTModelOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FTModelOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilities

`func (o *FTModelOut) GetCapabilities() FTModelCapabilitiesOut`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *FTModelOut) GetCapabilitiesOk() (*FTModelCapabilitiesOut, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *FTModelOut) SetCapabilities(v FTModelCapabilitiesOut)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *FTModelOut) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *FTModelOut) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *FTModelOut) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *FTModelOut) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *FTModelOut) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *FTModelOut) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *FTModelOut) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *FTModelOut) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *FTModelOut) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *FTModelOut) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *FTModelOut) SetJob(v string)`

SetJob sets Job field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


