# SharingOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**OrgId** | **string** |  | 
**Role** | **string** |  | 
**ShareWithType** | **string** |  | 
**ShareWithUuid** | **string** |  | 

## Methods

### NewSharingOut

`func NewSharingOut(libraryId string, orgId string, role string, shareWithType string, shareWithUuid string, ) *SharingOut`

NewSharingOut instantiates a new SharingOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingOutWithDefaults

`func NewSharingOutWithDefaults() *SharingOut`

NewSharingOutWithDefaults instantiates a new SharingOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryId

`func (o *SharingOut) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *SharingOut) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *SharingOut) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.


### GetUserId

`func (o *SharingOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SharingOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SharingOut) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SharingOut) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SharingOut) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SharingOut) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetOrgId

`func (o *SharingOut) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SharingOut) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SharingOut) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *SharingOut) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SharingOut) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SharingOut) SetRole(v string)`

SetRole sets Role field to given value.


### GetShareWithType

`func (o *SharingOut) GetShareWithType() string`

GetShareWithType returns the ShareWithType field if non-nil, zero value otherwise.

### GetShareWithTypeOk

`func (o *SharingOut) GetShareWithTypeOk() (*string, bool)`

GetShareWithTypeOk returns a tuple with the ShareWithType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithType

`func (o *SharingOut) SetShareWithType(v string)`

SetShareWithType sets ShareWithType field to given value.


### GetShareWithUuid

`func (o *SharingOut) GetShareWithUuid() string`

GetShareWithUuid returns the ShareWithUuid field if non-nil, zero value otherwise.

### GetShareWithUuidOk

`func (o *SharingOut) GetShareWithUuidOk() (*string, bool)`

GetShareWithUuidOk returns a tuple with the ShareWithUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithUuid

`func (o *SharingOut) SetShareWithUuid(v string)`

SetShareWithUuid sets ShareWithUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


