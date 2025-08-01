/*
Mistral AI API

Our Chat Completion and Embeddings APIs specification. Create your account on [La Plateforme](https://console.mistral.ai) to get access and read the [docs](https://docs.mistral.ai) to learn how to use it.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistral

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the LibraryOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryOut{}

// LibraryOut struct for LibraryOut
type LibraryOut struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OwnerId string `json:"owner_id"`
	OwnerType string `json:"owner_type"`
	TotalSize int32 `json:"total_size"`
	NbDocuments int32 `json:"nb_documents"`
	ChunkSize NullableInt32 `json:"chunk_size"`
	Emoji NullableString `json:"emoji,omitempty"`
	Description NullableString `json:"description,omitempty"`
	GeneratedName NullableString `json:"generated_name,omitempty"`
	GeneratedDescription NullableString `json:"generated_description,omitempty"`
	ExplicitUserMembersCount NullableInt32 `json:"explicit_user_members_count,omitempty"`
	ExplicitWorkspaceMembersCount NullableInt32 `json:"explicit_workspace_members_count,omitempty"`
	OrgSharingRole NullableString `json:"org_sharing_role,omitempty"`
}

type _LibraryOut LibraryOut

// NewLibraryOut instantiates a new LibraryOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryOut(id string, name string, createdAt time.Time, updatedAt time.Time, ownerId string, ownerType string, totalSize int32, nbDocuments int32, chunkSize NullableInt32) *LibraryOut {
	this := LibraryOut{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.OwnerId = ownerId
	this.OwnerType = ownerType
	this.TotalSize = totalSize
	this.NbDocuments = nbDocuments
	this.ChunkSize = chunkSize
	return &this
}

// NewLibraryOutWithDefaults instantiates a new LibraryOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryOutWithDefaults() *LibraryOut {
	this := LibraryOut{}
	return &this
}

// GetId returns the Id field value
func (o *LibraryOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LibraryOut) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LibraryOut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LibraryOut) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LibraryOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LibraryOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LibraryOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LibraryOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetOwnerId returns the OwnerId field value
func (o *LibraryOut) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *LibraryOut) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerType returns the OwnerType field value
func (o *LibraryOut) GetOwnerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetOwnerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value
func (o *LibraryOut) SetOwnerType(v string) {
	o.OwnerType = v
}

// GetTotalSize returns the TotalSize field value
func (o *LibraryOut) GetTotalSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetTotalSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSize, true
}

// SetTotalSize sets field value
func (o *LibraryOut) SetTotalSize(v int32) {
	o.TotalSize = v
}

// GetNbDocuments returns the NbDocuments field value
func (o *LibraryOut) GetNbDocuments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbDocuments
}

// GetNbDocumentsOk returns a tuple with the NbDocuments field value
// and a boolean to check if the value has been set.
func (o *LibraryOut) GetNbDocumentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbDocuments, true
}

// SetNbDocuments sets field value
func (o *LibraryOut) SetNbDocuments(v int32) {
	o.NbDocuments = v
}

// GetChunkSize returns the ChunkSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *LibraryOut) GetChunkSize() int32 {
	if o == nil || o.ChunkSize.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// SetChunkSize sets field value
func (o *LibraryOut) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetEmoji() string {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret string
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *LibraryOut) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableString and assigns it to the Emoji field.
func (o *LibraryOut) SetEmoji(v string) {
	o.Emoji.Set(&v)
}
// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *LibraryOut) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *LibraryOut) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *LibraryOut) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *LibraryOut) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *LibraryOut) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *LibraryOut) UnsetDescription() {
	o.Description.Unset()
}

// GetGeneratedName returns the GeneratedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetGeneratedName() string {
	if o == nil || IsNil(o.GeneratedName.Get()) {
		var ret string
		return ret
	}
	return *o.GeneratedName.Get()
}

// GetGeneratedNameOk returns a tuple with the GeneratedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetGeneratedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeneratedName.Get(), o.GeneratedName.IsSet()
}

// HasGeneratedName returns a boolean if a field has been set.
func (o *LibraryOut) HasGeneratedName() bool {
	if o != nil && o.GeneratedName.IsSet() {
		return true
	}

	return false
}

// SetGeneratedName gets a reference to the given NullableString and assigns it to the GeneratedName field.
func (o *LibraryOut) SetGeneratedName(v string) {
	o.GeneratedName.Set(&v)
}
// SetGeneratedNameNil sets the value for GeneratedName to be an explicit nil
func (o *LibraryOut) SetGeneratedNameNil() {
	o.GeneratedName.Set(nil)
}

// UnsetGeneratedName ensures that no value is present for GeneratedName, not even an explicit nil
func (o *LibraryOut) UnsetGeneratedName() {
	o.GeneratedName.Unset()
}

// GetGeneratedDescription returns the GeneratedDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetGeneratedDescription() string {
	if o == nil || IsNil(o.GeneratedDescription.Get()) {
		var ret string
		return ret
	}
	return *o.GeneratedDescription.Get()
}

// GetGeneratedDescriptionOk returns a tuple with the GeneratedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetGeneratedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeneratedDescription.Get(), o.GeneratedDescription.IsSet()
}

// HasGeneratedDescription returns a boolean if a field has been set.
func (o *LibraryOut) HasGeneratedDescription() bool {
	if o != nil && o.GeneratedDescription.IsSet() {
		return true
	}

	return false
}

// SetGeneratedDescription gets a reference to the given NullableString and assigns it to the GeneratedDescription field.
func (o *LibraryOut) SetGeneratedDescription(v string) {
	o.GeneratedDescription.Set(&v)
}
// SetGeneratedDescriptionNil sets the value for GeneratedDescription to be an explicit nil
func (o *LibraryOut) SetGeneratedDescriptionNil() {
	o.GeneratedDescription.Set(nil)
}

// UnsetGeneratedDescription ensures that no value is present for GeneratedDescription, not even an explicit nil
func (o *LibraryOut) UnsetGeneratedDescription() {
	o.GeneratedDescription.Unset()
}

// GetExplicitUserMembersCount returns the ExplicitUserMembersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetExplicitUserMembersCount() int32 {
	if o == nil || IsNil(o.ExplicitUserMembersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ExplicitUserMembersCount.Get()
}

// GetExplicitUserMembersCountOk returns a tuple with the ExplicitUserMembersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetExplicitUserMembersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplicitUserMembersCount.Get(), o.ExplicitUserMembersCount.IsSet()
}

// HasExplicitUserMembersCount returns a boolean if a field has been set.
func (o *LibraryOut) HasExplicitUserMembersCount() bool {
	if o != nil && o.ExplicitUserMembersCount.IsSet() {
		return true
	}

	return false
}

// SetExplicitUserMembersCount gets a reference to the given NullableInt32 and assigns it to the ExplicitUserMembersCount field.
func (o *LibraryOut) SetExplicitUserMembersCount(v int32) {
	o.ExplicitUserMembersCount.Set(&v)
}
// SetExplicitUserMembersCountNil sets the value for ExplicitUserMembersCount to be an explicit nil
func (o *LibraryOut) SetExplicitUserMembersCountNil() {
	o.ExplicitUserMembersCount.Set(nil)
}

// UnsetExplicitUserMembersCount ensures that no value is present for ExplicitUserMembersCount, not even an explicit nil
func (o *LibraryOut) UnsetExplicitUserMembersCount() {
	o.ExplicitUserMembersCount.Unset()
}

// GetExplicitWorkspaceMembersCount returns the ExplicitWorkspaceMembersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetExplicitWorkspaceMembersCount() int32 {
	if o == nil || IsNil(o.ExplicitWorkspaceMembersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ExplicitWorkspaceMembersCount.Get()
}

// GetExplicitWorkspaceMembersCountOk returns a tuple with the ExplicitWorkspaceMembersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetExplicitWorkspaceMembersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplicitWorkspaceMembersCount.Get(), o.ExplicitWorkspaceMembersCount.IsSet()
}

// HasExplicitWorkspaceMembersCount returns a boolean if a field has been set.
func (o *LibraryOut) HasExplicitWorkspaceMembersCount() bool {
	if o != nil && o.ExplicitWorkspaceMembersCount.IsSet() {
		return true
	}

	return false
}

// SetExplicitWorkspaceMembersCount gets a reference to the given NullableInt32 and assigns it to the ExplicitWorkspaceMembersCount field.
func (o *LibraryOut) SetExplicitWorkspaceMembersCount(v int32) {
	o.ExplicitWorkspaceMembersCount.Set(&v)
}
// SetExplicitWorkspaceMembersCountNil sets the value for ExplicitWorkspaceMembersCount to be an explicit nil
func (o *LibraryOut) SetExplicitWorkspaceMembersCountNil() {
	o.ExplicitWorkspaceMembersCount.Set(nil)
}

// UnsetExplicitWorkspaceMembersCount ensures that no value is present for ExplicitWorkspaceMembersCount, not even an explicit nil
func (o *LibraryOut) UnsetExplicitWorkspaceMembersCount() {
	o.ExplicitWorkspaceMembersCount.Unset()
}

// GetOrgSharingRole returns the OrgSharingRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOut) GetOrgSharingRole() string {
	if o == nil || IsNil(o.OrgSharingRole.Get()) {
		var ret string
		return ret
	}
	return *o.OrgSharingRole.Get()
}

// GetOrgSharingRoleOk returns a tuple with the OrgSharingRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOut) GetOrgSharingRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgSharingRole.Get(), o.OrgSharingRole.IsSet()
}

// HasOrgSharingRole returns a boolean if a field has been set.
func (o *LibraryOut) HasOrgSharingRole() bool {
	if o != nil && o.OrgSharingRole.IsSet() {
		return true
	}

	return false
}

// SetOrgSharingRole gets a reference to the given NullableString and assigns it to the OrgSharingRole field.
func (o *LibraryOut) SetOrgSharingRole(v string) {
	o.OrgSharingRole.Set(&v)
}
// SetOrgSharingRoleNil sets the value for OrgSharingRole to be an explicit nil
func (o *LibraryOut) SetOrgSharingRoleNil() {
	o.OrgSharingRole.Set(nil)
}

// UnsetOrgSharingRole ensures that no value is present for OrgSharingRole, not even an explicit nil
func (o *LibraryOut) UnsetOrgSharingRole() {
	o.OrgSharingRole.Unset()
}

func (o LibraryOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_type"] = o.OwnerType
	toSerialize["total_size"] = o.TotalSize
	toSerialize["nb_documents"] = o.NbDocuments
	toSerialize["chunk_size"] = o.ChunkSize.Get()
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.GeneratedName.IsSet() {
		toSerialize["generated_name"] = o.GeneratedName.Get()
	}
	if o.GeneratedDescription.IsSet() {
		toSerialize["generated_description"] = o.GeneratedDescription.Get()
	}
	if o.ExplicitUserMembersCount.IsSet() {
		toSerialize["explicit_user_members_count"] = o.ExplicitUserMembersCount.Get()
	}
	if o.ExplicitWorkspaceMembersCount.IsSet() {
		toSerialize["explicit_workspace_members_count"] = o.ExplicitWorkspaceMembersCount.Get()
	}
	if o.OrgSharingRole.IsSet() {
		toSerialize["org_sharing_role"] = o.OrgSharingRole.Get()
	}
	return toSerialize, nil
}

func (o *LibraryOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"owner_id",
		"owner_type",
		"total_size",
		"nb_documents",
		"chunk_size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLibraryOut := _LibraryOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLibraryOut)

	if err != nil {
		return err
	}

	*o = LibraryOut(varLibraryOut)

	return err
}

type NullableLibraryOut struct {
	value *LibraryOut
	isSet bool
}

func (v NullableLibraryOut) Get() *LibraryOut {
	return v.value
}

func (v *NullableLibraryOut) Set(val *LibraryOut) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryOut) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryOut(val *LibraryOut) *NullableLibraryOut {
	return &NullableLibraryOut{value: val, isSet: true}
}

func (v NullableLibraryOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


