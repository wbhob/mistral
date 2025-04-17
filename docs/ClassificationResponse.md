# ClassificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Model** | **string** |  | 
**Results** | [**[]map[string]ClassificationTargetResult**](map[string]ClassificationTargetResult.md) |  | 

## Methods

### NewClassificationResponse

`func NewClassificationResponse(id string, model string, results []map[string]ClassificationTargetResult, ) *ClassificationResponse`

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


### GetResults

`func (o *ClassificationResponse) GetResults() []map[string]ClassificationTargetResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ClassificationResponse) GetResultsOk() (*[]map[string]ClassificationTargetResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ClassificationResponse) SetResults(v []map[string]ClassificationTargetResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


