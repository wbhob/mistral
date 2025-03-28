# WandbIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "wandb"]
**Project** | **string** | The name of the project that the new run will be created under. | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | **string** | The WandB API key to use for authentication. | 
**RunName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWandbIntegration

`func NewWandbIntegration(project string, apiKey string, ) *WandbIntegration`

NewWandbIntegration instantiates a new WandbIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWandbIntegrationWithDefaults

`func NewWandbIntegrationWithDefaults() *WandbIntegration`

NewWandbIntegrationWithDefaults instantiates a new WandbIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WandbIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WandbIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WandbIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WandbIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *WandbIntegration) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WandbIntegration) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WandbIntegration) SetProject(v string)`

SetProject sets Project field to given value.


### GetName

`func (o *WandbIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WandbIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WandbIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WandbIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WandbIntegration) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WandbIntegration) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetApiKey

`func (o *WandbIntegration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *WandbIntegration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *WandbIntegration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRunName

`func (o *WandbIntegration) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *WandbIntegration) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *WandbIntegration) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *WandbIntegration) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### SetRunNameNil

`func (o *WandbIntegration) SetRunNameNil(b bool)`

 SetRunNameNil sets the value for RunName to be an explicit nil

### UnsetRunName
`func (o *WandbIntegration) UnsetRunName()`

UnsetRunName ensures that no value is present for RunName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


