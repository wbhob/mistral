# CompletionJobOutRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "github"]
**Name** | **string** |  | 
**Owner** | **string** |  | 
**Ref** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]
**CommitId** | **string** |  | 

## Methods

### NewCompletionJobOutRepositoriesInner

`func NewCompletionJobOutRepositoriesInner(name string, owner string, commitId string, ) *CompletionJobOutRepositoriesInner`

NewCompletionJobOutRepositoriesInner instantiates a new CompletionJobOutRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionJobOutRepositoriesInnerWithDefaults

`func NewCompletionJobOutRepositoriesInnerWithDefaults() *CompletionJobOutRepositoriesInner`

NewCompletionJobOutRepositoriesInnerWithDefaults instantiates a new CompletionJobOutRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompletionJobOutRepositoriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompletionJobOutRepositoriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompletionJobOutRepositoriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompletionJobOutRepositoriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CompletionJobOutRepositoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompletionJobOutRepositoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompletionJobOutRepositoriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CompletionJobOutRepositoriesInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CompletionJobOutRepositoriesInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CompletionJobOutRepositoriesInner) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRef

`func (o *CompletionJobOutRepositoriesInner) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CompletionJobOutRepositoriesInner) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CompletionJobOutRepositoriesInner) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CompletionJobOutRepositoriesInner) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetWeight

`func (o *CompletionJobOutRepositoriesInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CompletionJobOutRepositoriesInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CompletionJobOutRepositoriesInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CompletionJobOutRepositoriesInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCommitId

`func (o *CompletionJobOutRepositoriesInner) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *CompletionJobOutRepositoriesInner) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *CompletionJobOutRepositoriesInner) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


