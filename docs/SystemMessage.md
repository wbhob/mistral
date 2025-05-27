# SystemMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content4**](Content4.md) |  | 
**Role** | Pointer to **string** |  | [optional] [default to "system"]

## Methods

### NewSystemMessage

`func NewSystemMessage(content Content4, ) *SystemMessage`

NewSystemMessage instantiates a new SystemMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemMessageWithDefaults

`func NewSystemMessageWithDefaults() *SystemMessage`

NewSystemMessageWithDefaults instantiates a new SystemMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SystemMessage) GetContent() Content4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SystemMessage) GetContentOk() (*Content4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SystemMessage) SetContent(v Content4)`

SetContent sets Content field to given value.


### GetRole

`func (o *SystemMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SystemMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SystemMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SystemMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


