/*
FileMage Gateway

FileMage Gateway provides a simple API to configure storage endpoints, users, and keys. Please note that all urls must end in a trailing slash, and the Content-Type header must be set to application/json.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package filemagego

import (
	"encoding/json"
)

// checks if the BlobEndpointConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlobEndpointConfig{}

// BlobEndpointConfig Settings specific to the Azure Blob Storage endpoint types.
type BlobEndpointConfig struct {
	// Azure storage account secret key.
	AccountKey *string `json:"accountKey,omitempty"`
	// Azure storage account name.
	AccountName *string `json:"accountName,omitempty"`
	// Azure storage Blob container.
	ContainerName *string `json:"containerName,omitempty"`
	// Set to true if the Blob container has Hierarchical namespaces enabled.
	IsHNS *bool `json:"isHNS,omitempty"`
	// Enabled to use the managed service identify associated with the instance to authenticate.
	UseMSI *bool `json:"useMSI,omitempty"`
	// Client ID of a user assigned managed system identity to use for authentication.
	UserAssignedId *string `json:"userAssignedId,omitempty"`
	// Storage endpoint type. Use \"azure\" when connecting to Azure Blob Storage.
	Type *string `json:"type,omitempty"`
}

// NewBlobEndpointConfig instantiates a new BlobEndpointConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobEndpointConfig() *BlobEndpointConfig {
	this := BlobEndpointConfig{}
	return &this
}

// NewBlobEndpointConfigWithDefaults instantiates a new BlobEndpointConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobEndpointConfigWithDefaults() *BlobEndpointConfig {
	this := BlobEndpointConfig{}
	return &this
}

// GetAccountKey returns the AccountKey field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetAccountKey() string {
	if o == nil || IsNil(o.AccountKey) {
		var ret string
		return ret
	}
	return *o.AccountKey
}

// GetAccountKeyOk returns a tuple with the AccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccountKey) {
		return nil, false
	}
	return o.AccountKey, true
}

// HasAccountKey returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasAccountKey() bool {
	if o != nil && !IsNil(o.AccountKey) {
		return true
	}

	return false
}

// SetAccountKey gets a reference to the given string and assigns it to the AccountKey field.
func (o *BlobEndpointConfig) SetAccountKey(v string) {
	o.AccountKey = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *BlobEndpointConfig) SetAccountName(v string) {
	o.AccountName = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *BlobEndpointConfig) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetIsHNS returns the IsHNS field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetIsHNS() bool {
	if o == nil || IsNil(o.IsHNS) {
		var ret bool
		return ret
	}
	return *o.IsHNS
}

// GetIsHNSOk returns a tuple with the IsHNS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetIsHNSOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHNS) {
		return nil, false
	}
	return o.IsHNS, true
}

// HasIsHNS returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasIsHNS() bool {
	if o != nil && !IsNil(o.IsHNS) {
		return true
	}

	return false
}

// SetIsHNS gets a reference to the given bool and assigns it to the IsHNS field.
func (o *BlobEndpointConfig) SetIsHNS(v bool) {
	o.IsHNS = &v
}

// GetUseMSI returns the UseMSI field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetUseMSI() bool {
	if o == nil || IsNil(o.UseMSI) {
		var ret bool
		return ret
	}
	return *o.UseMSI
}

// GetUseMSIOk returns a tuple with the UseMSI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetUseMSIOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMSI) {
		return nil, false
	}
	return o.UseMSI, true
}

// HasUseMSI returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasUseMSI() bool {
	if o != nil && !IsNil(o.UseMSI) {
		return true
	}

	return false
}

// SetUseMSI gets a reference to the given bool and assigns it to the UseMSI field.
func (o *BlobEndpointConfig) SetUseMSI(v bool) {
	o.UseMSI = &v
}

// GetUserAssignedId returns the UserAssignedId field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetUserAssignedId() string {
	if o == nil || IsNil(o.UserAssignedId) {
		var ret string
		return ret
	}
	return *o.UserAssignedId
}

// GetUserAssignedIdOk returns a tuple with the UserAssignedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetUserAssignedIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserAssignedId) {
		return nil, false
	}
	return o.UserAssignedId, true
}

// HasUserAssignedId returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasUserAssignedId() bool {
	if o != nil && !IsNil(o.UserAssignedId) {
		return true
	}

	return false
}

// SetUserAssignedId gets a reference to the given string and assigns it to the UserAssignedId field.
func (o *BlobEndpointConfig) SetUserAssignedId(v string) {
	o.UserAssignedId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlobEndpointConfig) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobEndpointConfig) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlobEndpointConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlobEndpointConfig) SetType(v string) {
	o.Type = &v
}

func (o BlobEndpointConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlobEndpointConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountKey) {
		toSerialize["accountKey"] = o.AccountKey
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !IsNil(o.IsHNS) {
		toSerialize["isHNS"] = o.IsHNS
	}
	if !IsNil(o.UseMSI) {
		toSerialize["useMSI"] = o.UseMSI
	}
	if !IsNil(o.UserAssignedId) {
		toSerialize["userAssignedId"] = o.UserAssignedId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBlobEndpointConfig struct {
	value *BlobEndpointConfig
	isSet bool
}

func (v NullableBlobEndpointConfig) Get() *BlobEndpointConfig {
	return v.value
}

func (v *NullableBlobEndpointConfig) Set(val *BlobEndpointConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobEndpointConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobEndpointConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobEndpointConfig(val *BlobEndpointConfig) *NullableBlobEndpointConfig {
	return &NullableBlobEndpointConfig{value: val, isSet: true}
}

func (v NullableBlobEndpointConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobEndpointConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

