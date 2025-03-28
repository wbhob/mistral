# JobInRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "github"]
**Name** | **string** |  | 
**Owner** | **string** |  | 
**Ref** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]
**Token** | **string** |  | 

## Methods

### NewJobInRepositoriesInner

`func NewJobInRepositoriesInner(name string, owner string, token string, ) *JobInRepositoriesInner`

NewJobInRepositoriesInner instantiates a new JobInRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInRepositoriesInnerWithDefaults

`func NewJobInRepositoriesInnerWithDefaults() *JobInRepositoriesInner`

NewJobInRepositoriesInnerWithDefaults instantiates a new JobInRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JobInRepositoriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobInRepositoriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobInRepositoriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobInRepositoriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *JobInRepositoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobInRepositoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobInRepositoriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *JobInRepositoriesInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *JobInRepositoriesInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *JobInRepositoriesInner) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRef

`func (o *JobInRepositoriesInner) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *JobInRepositoriesInner) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *JobInRepositoriesInner) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *JobInRepositoriesInner) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetWeight

`func (o *JobInRepositoriesInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *JobInRepositoriesInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *JobInRepositoriesInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *JobInRepositoriesInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetToken

`func (o *JobInRepositoriesInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *JobInRepositoriesInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *JobInRepositoriesInner) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


