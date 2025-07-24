# SharingIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Level** | [**ShareEnum**](ShareEnum.md) |  | 
**ShareWithUuid** | **string** | The id of the entity (user, workspace or organization) to share with | 
**ShareWithType** | [**EntityType**](EntityType.md) |  | 

## Methods

### NewSharingIn

`func NewSharingIn(orgId string, level ShareEnum, shareWithUuid string, shareWithType EntityType, ) *SharingIn`

NewSharingIn instantiates a new SharingIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingInWithDefaults

`func NewSharingInWithDefaults() *SharingIn`

NewSharingInWithDefaults instantiates a new SharingIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SharingIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SharingIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SharingIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetLevel

`func (o *SharingIn) GetLevel() ShareEnum`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SharingIn) GetLevelOk() (*ShareEnum, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SharingIn) SetLevel(v ShareEnum)`

SetLevel sets Level field to given value.


### GetShareWithUuid

`func (o *SharingIn) GetShareWithUuid() string`

GetShareWithUuid returns the ShareWithUuid field if non-nil, zero value otherwise.

### GetShareWithUuidOk

`func (o *SharingIn) GetShareWithUuidOk() (*string, bool)`

GetShareWithUuidOk returns a tuple with the ShareWithUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithUuid

`func (o *SharingIn) SetShareWithUuid(v string)`

SetShareWithUuid sets ShareWithUuid field to given value.


### GetShareWithType

`func (o *SharingIn) GetShareWithType() EntityType`

GetShareWithType returns the ShareWithType field if non-nil, zero value otherwise.

### GetShareWithTypeOk

`func (o *SharingIn) GetShareWithTypeOk() (*EntityType, bool)`

GetShareWithTypeOk returns a tuple with the ShareWithType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithType

`func (o *SharingIn) SetShareWithType(v EntityType)`

SetShareWithType sets ShareWithType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


