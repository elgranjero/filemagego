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

// checks if the KeyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyInfo{}

// KeyInfo struct for KeyInfo
type KeyInfo struct {
	// Key ID.
	Id *int32 `json:"id,omitempty"`
	// MD5 key fingerprint.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// The public key associated with the key pair.
	KeyData *string `json:"keyData,omitempty"`
	// A title to display for the key.
	Title *string `json:"title,omitempty"`
	// A UTC timestamp of when the key was created.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewKeyInfo instantiates a new KeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyInfo() *KeyInfo {
	this := KeyInfo{}
	return &this
}

// NewKeyInfoWithDefaults instantiates a new KeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyInfoWithDefaults() *KeyInfo {
	this := KeyInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyInfo) SetId(v int32) {
	o.Id = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *KeyInfo) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *KeyInfo) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *KeyInfo) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *KeyInfo) GetKeyData() string {
	if o == nil || IsNil(o.KeyData) {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetKeyDataOk() (*string, bool) {
	if o == nil || IsNil(o.KeyData) {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *KeyInfo) HasKeyData() bool {
	if o != nil && !IsNil(o.KeyData) {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *KeyInfo) SetKeyData(v string) {
	o.KeyData = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *KeyInfo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *KeyInfo) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *KeyInfo) SetTitle(v string) {
	o.Title = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KeyInfo) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KeyInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *KeyInfo) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o KeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.KeyData) {
		toSerialize["keyData"] = o.KeyData
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableKeyInfo struct {
	value *KeyInfo
	isSet bool
}

func (v NullableKeyInfo) Get() *KeyInfo {
	return v.value
}

func (v *NullableKeyInfo) Set(val *KeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyInfo(val *KeyInfo) *NullableKeyInfo {
	return &NullableKeyInfo{value: val, isSet: true}
}

func (v NullableKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

