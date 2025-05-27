# JsonSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to [**NullableDescription**](Description.md) |  | [optional] 
**Schema** | **map[string]interface{}** |  | 
**Strict** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewJsonSchema

`func NewJsonSchema(name string, schema map[string]interface{}, ) *JsonSchema`

NewJsonSchema instantiates a new JsonSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSchemaWithDefaults

`func NewJsonSchemaWithDefaults() *JsonSchema`

NewJsonSchemaWithDefaults instantiates a new JsonSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JsonSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JsonSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JsonSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JsonSchema) GetDescription() Description`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonSchema) GetDescriptionOk() (*Description, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonSchema) SetDescription(v Description)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JsonSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JsonSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSchema

`func (o *JsonSchema) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JsonSchema) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JsonSchema) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetStrict

`func (o *JsonSchema) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *JsonSchema) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *JsonSchema) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *JsonSchema) HasStrict() bool`

HasStrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


