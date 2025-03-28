# ClassificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Results** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewClassificationResponse

`func NewClassificationResponse() *ClassificationResponse`

NewClassificationResponse instantiates a new ClassificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationResponseWithDefaults

`func NewClassificationResponseWithDefaults() *ClassificationResponse`

NewClassificationResponseWithDefaults instantiates a new ClassificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassificationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassificationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassificationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClassificationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *ClassificationResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClassificationResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClassificationResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClassificationResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResults

`func (o *ClassificationResponse) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ClassificationResponse) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ClassificationResponse) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ClassificationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


