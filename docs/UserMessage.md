# UserMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**NullableContent3**](Content3.md) |  | 
**Role** | Pointer to **string** |  | [optional] [default to "user"]

## Methods

### NewUserMessage

`func NewUserMessage(content NullableContent3, ) *UserMessage`

NewUserMessage instantiates a new UserMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMessageWithDefaults

`func NewUserMessageWithDefaults() *UserMessage`

NewUserMessageWithDefaults instantiates a new UserMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UserMessage) GetContent() Content3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UserMessage) GetContentOk() (*Content3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UserMessage) SetContent(v Content3)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *UserMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *UserMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRole

`func (o *UserMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


