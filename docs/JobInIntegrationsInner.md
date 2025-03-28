# JobInIntegrationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "wandb"]
**Project** | **string** | The name of the project that the new run will be created under. | 
**Name** | Pointer to **string** |  | [optional] 
**ApiKey** | **string** | The WandB API key to use for authentication. | 
**RunName** | Pointer to **string** |  | [optional] 

## Methods

### NewJobInIntegrationsInner

`func NewJobInIntegrationsInner(project string, apiKey string, ) *JobInIntegrationsInner`

NewJobInIntegrationsInner instantiates a new JobInIntegrationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInIntegrationsInnerWithDefaults

`func NewJobInIntegrationsInnerWithDefaults() *JobInIntegrationsInner`

NewJobInIntegrationsInnerWithDefaults instantiates a new JobInIntegrationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JobInIntegrationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobInIntegrationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobInIntegrationsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobInIntegrationsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProject

`func (o *JobInIntegrationsInner) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *JobInIntegrationsInner) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *JobInIntegrationsInner) SetProject(v string)`

SetProject sets Project field to given value.


### GetName

`func (o *JobInIntegrationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobInIntegrationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobInIntegrationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobInIntegrationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiKey

`func (o *JobInIntegrationsInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *JobInIntegrationsInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *JobInIntegrationsInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRunName

`func (o *JobInIntegrationsInner) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *JobInIntegrationsInner) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *JobInIntegrationsInner) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *JobInIntegrationsInner) HasRunName() bool`

HasRunName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


