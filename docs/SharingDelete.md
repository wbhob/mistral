# SharingDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**ShareWithUuid** | **string** | The id of the entity (user, workspace or organization) to share with | 
**ShareWithType** | [**EntityType**](EntityType.md) |  | 

## Methods

### NewSharingDelete

`func NewSharingDelete(orgId string, shareWithUuid string, shareWithType EntityType, ) *SharingDelete`

NewSharingDelete instantiates a new SharingDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingDeleteWithDefaults

`func NewSharingDeleteWithDefaults() *SharingDelete`

NewSharingDeleteWithDefaults instantiates a new SharingDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SharingDelete) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SharingDelete) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SharingDelete) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetShareWithUuid

`func (o *SharingDelete) GetShareWithUuid() string`

GetShareWithUuid returns the ShareWithUuid field if non-nil, zero value otherwise.

### GetShareWithUuidOk

`func (o *SharingDelete) GetShareWithUuidOk() (*string, bool)`

GetShareWithUuidOk returns a tuple with the ShareWithUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithUuid

`func (o *SharingDelete) SetShareWithUuid(v string)`

SetShareWithUuid sets ShareWithUuid field to given value.


### GetShareWithType

`func (o *SharingDelete) GetShareWithType() EntityType`

GetShareWithType returns the ShareWithType field if non-nil, zero value otherwise.

### GetShareWithTypeOk

`func (o *SharingDelete) GetShareWithTypeOk() (*EntityType, bool)`

GetShareWithTypeOk returns a tuple with the ShareWithType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithType

`func (o *SharingDelete) SetShareWithType(v EntityType)`

SetShareWithType sets ShareWithType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


