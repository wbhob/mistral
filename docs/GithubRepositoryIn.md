# GithubRepositoryIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "github"]
**Name** | **string** |  | 
**Owner** | **string** |  | 
**Ref** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]
**Token** | **string** |  | 

## Methods

### NewGithubRepositoryIn

`func NewGithubRepositoryIn(name string, owner string, token string, ) *GithubRepositoryIn`

NewGithubRepositoryIn instantiates a new GithubRepositoryIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubRepositoryInWithDefaults

`func NewGithubRepositoryInWithDefaults() *GithubRepositoryIn`

NewGithubRepositoryInWithDefaults instantiates a new GithubRepositoryIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GithubRepositoryIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubRepositoryIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubRepositoryIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GithubRepositoryIn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GithubRepositoryIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubRepositoryIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubRepositoryIn) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *GithubRepositoryIn) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GithubRepositoryIn) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GithubRepositoryIn) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRef

`func (o *GithubRepositoryIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GithubRepositoryIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GithubRepositoryIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GithubRepositoryIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### SetRefNil

`func (o *GithubRepositoryIn) SetRefNil(b bool)`

 SetRefNil sets the value for Ref to be an explicit nil

### UnsetRef
`func (o *GithubRepositoryIn) UnsetRef()`

UnsetRef ensures that no value is present for Ref, not even an explicit nil
### GetWeight

`func (o *GithubRepositoryIn) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GithubRepositoryIn) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GithubRepositoryIn) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GithubRepositoryIn) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetToken

`func (o *GithubRepositoryIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GithubRepositoryIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GithubRepositoryIn) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


