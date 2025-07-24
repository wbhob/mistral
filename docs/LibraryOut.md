# LibraryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**OwnerId** | **string** |  | 
**OwnerType** | **string** |  | 
**TotalSize** | **int32** |  | 
**NbDocuments** | **int32** |  | 
**ChunkSize** | **NullableInt32** |  | 
**Emoji** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**GeneratedName** | Pointer to **NullableString** |  | [optional] 
**GeneratedDescription** | Pointer to **NullableString** |  | [optional] 
**ExplicitUserMembersCount** | Pointer to **NullableInt32** |  | [optional] 
**ExplicitWorkspaceMembersCount** | Pointer to **NullableInt32** |  | [optional] 
**OrgSharingRole** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLibraryOut

`func NewLibraryOut(id string, name string, createdAt time.Time, updatedAt time.Time, ownerId string, ownerType string, totalSize int32, nbDocuments int32, chunkSize NullableInt32, ) *LibraryOut`

NewLibraryOut instantiates a new LibraryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOutWithDefaults

`func NewLibraryOutWithDefaults() *LibraryOut`

NewLibraryOutWithDefaults instantiates a new LibraryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LibraryOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LibraryOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryOut) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *LibraryOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LibraryOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LibraryOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LibraryOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LibraryOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LibraryOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOwnerId

`func (o *LibraryOut) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *LibraryOut) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *LibraryOut) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerType

`func (o *LibraryOut) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *LibraryOut) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *LibraryOut) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.


### GetTotalSize

`func (o *LibraryOut) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *LibraryOut) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *LibraryOut) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.


### GetNbDocuments

`func (o *LibraryOut) GetNbDocuments() int32`

GetNbDocuments returns the NbDocuments field if non-nil, zero value otherwise.

### GetNbDocumentsOk

`func (o *LibraryOut) GetNbDocumentsOk() (*int32, bool)`

GetNbDocumentsOk returns a tuple with the NbDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbDocuments

`func (o *LibraryOut) SetNbDocuments(v int32)`

SetNbDocuments sets NbDocuments field to given value.


### GetChunkSize

`func (o *LibraryOut) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *LibraryOut) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *LibraryOut) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.


### SetChunkSizeNil

`func (o *LibraryOut) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *LibraryOut) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetEmoji

`func (o *LibraryOut) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *LibraryOut) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *LibraryOut) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *LibraryOut) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *LibraryOut) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *LibraryOut) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetDescription

`func (o *LibraryOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LibraryOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LibraryOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LibraryOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LibraryOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LibraryOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGeneratedName

`func (o *LibraryOut) GetGeneratedName() string`

GetGeneratedName returns the GeneratedName field if non-nil, zero value otherwise.

### GetGeneratedNameOk

`func (o *LibraryOut) GetGeneratedNameOk() (*string, bool)`

GetGeneratedNameOk returns a tuple with the GeneratedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedName

`func (o *LibraryOut) SetGeneratedName(v string)`

SetGeneratedName sets GeneratedName field to given value.

### HasGeneratedName

`func (o *LibraryOut) HasGeneratedName() bool`

HasGeneratedName returns a boolean if a field has been set.

### SetGeneratedNameNil

`func (o *LibraryOut) SetGeneratedNameNil(b bool)`

 SetGeneratedNameNil sets the value for GeneratedName to be an explicit nil

### UnsetGeneratedName
`func (o *LibraryOut) UnsetGeneratedName()`

UnsetGeneratedName ensures that no value is present for GeneratedName, not even an explicit nil
### GetGeneratedDescription

`func (o *LibraryOut) GetGeneratedDescription() string`

GetGeneratedDescription returns the GeneratedDescription field if non-nil, zero value otherwise.

### GetGeneratedDescriptionOk

`func (o *LibraryOut) GetGeneratedDescriptionOk() (*string, bool)`

GetGeneratedDescriptionOk returns a tuple with the GeneratedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedDescription

`func (o *LibraryOut) SetGeneratedDescription(v string)`

SetGeneratedDescription sets GeneratedDescription field to given value.

### HasGeneratedDescription

`func (o *LibraryOut) HasGeneratedDescription() bool`

HasGeneratedDescription returns a boolean if a field has been set.

### SetGeneratedDescriptionNil

`func (o *LibraryOut) SetGeneratedDescriptionNil(b bool)`

 SetGeneratedDescriptionNil sets the value for GeneratedDescription to be an explicit nil

### UnsetGeneratedDescription
`func (o *LibraryOut) UnsetGeneratedDescription()`

UnsetGeneratedDescription ensures that no value is present for GeneratedDescription, not even an explicit nil
### GetExplicitUserMembersCount

`func (o *LibraryOut) GetExplicitUserMembersCount() int32`

GetExplicitUserMembersCount returns the ExplicitUserMembersCount field if non-nil, zero value otherwise.

### GetExplicitUserMembersCountOk

`func (o *LibraryOut) GetExplicitUserMembersCountOk() (*int32, bool)`

GetExplicitUserMembersCountOk returns a tuple with the ExplicitUserMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUserMembersCount

`func (o *LibraryOut) SetExplicitUserMembersCount(v int32)`

SetExplicitUserMembersCount sets ExplicitUserMembersCount field to given value.

### HasExplicitUserMembersCount

`func (o *LibraryOut) HasExplicitUserMembersCount() bool`

HasExplicitUserMembersCount returns a boolean if a field has been set.

### SetExplicitUserMembersCountNil

`func (o *LibraryOut) SetExplicitUserMembersCountNil(b bool)`

 SetExplicitUserMembersCountNil sets the value for ExplicitUserMembersCount to be an explicit nil

### UnsetExplicitUserMembersCount
`func (o *LibraryOut) UnsetExplicitUserMembersCount()`

UnsetExplicitUserMembersCount ensures that no value is present for ExplicitUserMembersCount, not even an explicit nil
### GetExplicitWorkspaceMembersCount

`func (o *LibraryOut) GetExplicitWorkspaceMembersCount() int32`

GetExplicitWorkspaceMembersCount returns the ExplicitWorkspaceMembersCount field if non-nil, zero value otherwise.

### GetExplicitWorkspaceMembersCountOk

`func (o *LibraryOut) GetExplicitWorkspaceMembersCountOk() (*int32, bool)`

GetExplicitWorkspaceMembersCountOk returns a tuple with the ExplicitWorkspaceMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitWorkspaceMembersCount

`func (o *LibraryOut) SetExplicitWorkspaceMembersCount(v int32)`

SetExplicitWorkspaceMembersCount sets ExplicitWorkspaceMembersCount field to given value.

### HasExplicitWorkspaceMembersCount

`func (o *LibraryOut) HasExplicitWorkspaceMembersCount() bool`

HasExplicitWorkspaceMembersCount returns a boolean if a field has been set.

### SetExplicitWorkspaceMembersCountNil

`func (o *LibraryOut) SetExplicitWorkspaceMembersCountNil(b bool)`

 SetExplicitWorkspaceMembersCountNil sets the value for ExplicitWorkspaceMembersCount to be an explicit nil

### UnsetExplicitWorkspaceMembersCount
`func (o *LibraryOut) UnsetExplicitWorkspaceMembersCount()`

UnsetExplicitWorkspaceMembersCount ensures that no value is present for ExplicitWorkspaceMembersCount, not even an explicit nil
### GetOrgSharingRole

`func (o *LibraryOut) GetOrgSharingRole() string`

GetOrgSharingRole returns the OrgSharingRole field if non-nil, zero value otherwise.

### GetOrgSharingRoleOk

`func (o *LibraryOut) GetOrgSharingRoleOk() (*string, bool)`

GetOrgSharingRoleOk returns a tuple with the OrgSharingRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSharingRole

`func (o *LibraryOut) SetOrgSharingRole(v string)`

SetOrgSharingRole sets OrgSharingRole field to given value.

### HasOrgSharingRole

`func (o *LibraryOut) HasOrgSharingRole() bool`

HasOrgSharingRole returns a boolean if a field has been set.

### SetOrgSharingRoleNil

`func (o *LibraryOut) SetOrgSharingRoleNil(b bool)`

 SetOrgSharingRoleNil sets the value for OrgSharingRole to be an explicit nil

### UnsetOrgSharingRole
`func (o *LibraryOut) UnsetOrgSharingRole()`

UnsetOrgSharingRole ensures that no value is present for OrgSharingRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


