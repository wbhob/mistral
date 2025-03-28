# ModelListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | Pointer to **int32** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] [default to "mistralai"]
**Capabilities** | [**ModelCapabilities**](ModelCapabilities.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**Deprecation** | Pointer to **time.Time** |  | [optional] 
**DefaultModelTemperature** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "base"]
**Job** | **string** |  | 
**Root** | **string** |  | 
**Archived** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewModelListDataInner

`func NewModelListDataInner(id string, capabilities ModelCapabilities, job string, root string, ) *ModelListDataInner`

NewModelListDataInner instantiates a new ModelListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelListDataInnerWithDefaults

`func NewModelListDataInnerWithDefaults() *ModelListDataInner`

NewModelListDataInnerWithDefaults instantiates a new ModelListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ModelListDataInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelListDataInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelListDataInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelListDataInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ModelListDataInner) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelListDataInner) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelListDataInner) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelListDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ModelListDataInner) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelListDataInner) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelListDataInner) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ModelListDataInner) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetCapabilities

`func (o *ModelListDataInner) GetCapabilities() ModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelListDataInner) GetCapabilitiesOk() (*ModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelListDataInner) SetCapabilities(v ModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetName

`func (o *ModelListDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelListDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelListDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelListDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelListDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelListDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelListDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelListDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxContextLength

`func (o *ModelListDataInner) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *ModelListDataInner) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *ModelListDataInner) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *ModelListDataInner) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *ModelListDataInner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ModelListDataInner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ModelListDataInner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ModelListDataInner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDeprecation

`func (o *ModelListDataInner) GetDeprecation() time.Time`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *ModelListDataInner) GetDeprecationOk() (*time.Time, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *ModelListDataInner) SetDeprecation(v time.Time)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *ModelListDataInner) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### GetDefaultModelTemperature

`func (o *ModelListDataInner) GetDefaultModelTemperature() float32`

GetDefaultModelTemperature returns the DefaultModelTemperature field if non-nil, zero value otherwise.

### GetDefaultModelTemperatureOk

`func (o *ModelListDataInner) GetDefaultModelTemperatureOk() (*float32, bool)`

GetDefaultModelTemperatureOk returns a tuple with the DefaultModelTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelTemperature

`func (o *ModelListDataInner) SetDefaultModelTemperature(v float32)`

SetDefaultModelTemperature sets DefaultModelTemperature field to given value.

### HasDefaultModelTemperature

`func (o *ModelListDataInner) HasDefaultModelTemperature() bool`

HasDefaultModelTemperature returns a boolean if a field has been set.

### GetType

`func (o *ModelListDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelListDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelListDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelListDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJob

`func (o *ModelListDataInner) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ModelListDataInner) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ModelListDataInner) SetJob(v string)`

SetJob sets Job field to given value.


### GetRoot

`func (o *ModelListDataInner) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ModelListDataInner) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ModelListDataInner) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *ModelListDataInner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ModelListDataInner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ModelListDataInner) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ModelListDataInner) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


