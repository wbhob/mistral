# ResponseFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ResponseFormats**](ResponseFormats.md) |  | [optional] 
**JsonSchema** | Pointer to [**NullableJsonSchema**](JsonSchema.md) |  | [optional] 

## Methods

### NewResponseFormat

`func NewResponseFormat() *ResponseFormat`

NewResponseFormat instantiates a new ResponseFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFormatWithDefaults

`func NewResponseFormatWithDefaults() *ResponseFormat`

NewResponseFormatWithDefaults instantiates a new ResponseFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseFormat) GetType() ResponseFormats`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseFormat) GetTypeOk() (*ResponseFormats, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseFormat) SetType(v ResponseFormats)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseFormat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJsonSchema

`func (o *ResponseFormat) GetJsonSchema() JsonSchema`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *ResponseFormat) GetJsonSchemaOk() (*JsonSchema, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *ResponseFormat) SetJsonSchema(v JsonSchema)`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *ResponseFormat) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.

### SetJsonSchemaNil

`func (o *ResponseFormat) SetJsonSchemaNil(b bool)`

 SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil

### UnsetJsonSchema
`func (o *ResponseFormat) UnsetJsonSchema()`

UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


