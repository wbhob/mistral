# WandbIntegrationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "wandb"]
**Project** | **string** | The name of the project that the new run will be created under. | 
**Name** | Pointer to **NullableString** |  | [optional] 
**RunName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWandbIntegrationOut

`func NewWandbIntegrationOut(project string, ) *WandbIntegrationOut`

NewWandbIntegrationOut instantiates a new WandbIntegrationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWandbIntegrationOutWithDefaults

`func NewWandbIntegrationOutWithDefaults() *WandbIntegrationOut`

NewWandbIntegrationOutWithDefaults instantiates a new WandbIntegrationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WandbIntegrationOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WandbIntegrationOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WandbIntegrationOut) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WandbIntegrationOut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *WandbIntegrationOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WandbIntegrationOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WandbIntegrationOut) SetProject(v string)`

SetProject sets Project field to given value.


### GetName

`func (o *WandbIntegrationOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WandbIntegrationOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WandbIntegrationOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WandbIntegrationOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WandbIntegrationOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WandbIntegrationOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRunName

`func (o *WandbIntegrationOut) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *WandbIntegrationOut) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *WandbIntegrationOut) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *WandbIntegrationOut) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### SetRunNameNil

`func (o *WandbIntegrationOut) SetRunNameNil(b bool)`

 SetRunNameNil sets the value for RunName to be an explicit nil

### UnsetRunName
`func (o *WandbIntegrationOut) UnsetRunName()`

UnsetRunName ensures that no value is present for RunName, not even an explicit nil
### GetUrl

`func (o *WandbIntegrationOut) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WandbIntegrationOut) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WandbIntegrationOut) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WandbIntegrationOut) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WandbIntegrationOut) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WandbIntegrationOut) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


