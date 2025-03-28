# BaseModelCard

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
**DefaultModelTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "base"]

## Methods

### NewBaseModelCard

`func NewBaseModelCard(id string, capabilities ModelCapabilities, ) *BaseModelCard`

NewBaseModelCard instantiates a new BaseModelCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseModelCardWithDefaults

`func NewBaseModelCardWithDefaults() *BaseModelCard`

NewBaseModelCardWithDefaults instantiates a new BaseModelCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseModelCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseModelCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseModelCard) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *BaseModelCard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BaseModelCard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BaseModelCard) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BaseModelCard) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *BaseModelCard) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseModelCard) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseModelCard) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseModelCard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *BaseModelCard) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *BaseModelCard) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *BaseModelCard) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *BaseModelCard) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetCapabilities

`func (o *BaseModelCard) GetCapabilities() ModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *BaseModelCard) GetCapabilitiesOk() (*ModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *BaseModelCard) SetCapabilities(v ModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetName

`func (o *BaseModelCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseModelCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseModelCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseModelCard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BaseModelCard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BaseModelCard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BaseModelCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseModelCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseModelCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseModelCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BaseModelCard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BaseModelCard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxContextLength

`func (o *BaseModelCard) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *BaseModelCard) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *BaseModelCard) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *BaseModelCard) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *BaseModelCard) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *BaseModelCard) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *BaseModelCard) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *BaseModelCard) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDeprecation

`func (o *BaseModelCard) GetDeprecation() time.Time`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *BaseModelCard) GetDeprecationOk() (*time.Time, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *BaseModelCard) SetDeprecation(v time.Time)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *BaseModelCard) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### SetDeprecationNil

`func (o *BaseModelCard) SetDeprecationNil(b bool)`

 SetDeprecationNil sets the value for Deprecation to be an explicit nil

### UnsetDeprecation
`func (o *BaseModelCard) UnsetDeprecation()`

UnsetDeprecation ensures that no value is present for Deprecation, not even an explicit nil
### GetDefaultModelTemperature

`func (o *BaseModelCard) GetDefaultModelTemperature() float32`

GetDefaultModelTemperature returns the DefaultModelTemperature field if non-nil, zero value otherwise.

### GetDefaultModelTemperatureOk

`func (o *BaseModelCard) GetDefaultModelTemperatureOk() (*float32, bool)`

GetDefaultModelTemperatureOk returns a tuple with the DefaultModelTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelTemperature

`func (o *BaseModelCard) SetDefaultModelTemperature(v float32)`

SetDefaultModelTemperature sets DefaultModelTemperature field to given value.

### HasDefaultModelTemperature

`func (o *BaseModelCard) HasDefaultModelTemperature() bool`

HasDefaultModelTemperature returns a boolean if a field has been set.

### SetDefaultModelTemperatureNil

`func (o *BaseModelCard) SetDefaultModelTemperatureNil(b bool)`

 SetDefaultModelTemperatureNil sets the value for DefaultModelTemperature to be an explicit nil

### UnsetDefaultModelTemperature
`func (o *BaseModelCard) UnsetDefaultModelTemperature()`

UnsetDefaultModelTemperature ensures that no value is present for DefaultModelTemperature, not even an explicit nil
### GetType

`func (o *BaseModelCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseModelCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseModelCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseModelCard) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


