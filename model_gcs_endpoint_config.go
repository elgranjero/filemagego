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

// checks if the GCSEndpointConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCSEndpointConfig{}

// GCSEndpointConfig Settings specific to the Google Cloud Storage endpoint types.
type GCSEndpointConfig struct {
	// The Google Cloud Storage bucket name.
	BucketName *string `json:"bucketName,omitempty"`
	// A base64 encoded Google Cloud JSON credentials file used to authenticate.
	Credentials *string `json:"credentials,omitempty"`
	// Enabled to use service account associated with the instance to authenticate.
	UseVMSA *bool `json:"useVMSA,omitempty"`
	// Storage endpoint type. Use \"gcp\" when connecting to Google Cloud Storage.
	Type *string `json:"type,omitempty"`
}

// NewGCSEndpointConfig instantiates a new GCSEndpointConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCSEndpointConfig() *GCSEndpointConfig {
	this := GCSEndpointConfig{}
	return &this
}

// NewGCSEndpointConfigWithDefaults instantiates a new GCSEndpointConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCSEndpointConfigWithDefaults() *GCSEndpointConfig {
	this := GCSEndpointConfig{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *GCSEndpointConfig) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSEndpointConfig) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *GCSEndpointConfig) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *GCSEndpointConfig) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GCSEndpointConfig) GetCredentials() string {
	if o == nil || IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSEndpointConfig) GetCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GCSEndpointConfig) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *GCSEndpointConfig) SetCredentials(v string) {
	o.Credentials = &v
}

// GetUseVMSA returns the UseVMSA field value if set, zero value otherwise.
func (o *GCSEndpointConfig) GetUseVMSA() bool {
	if o == nil || IsNil(o.UseVMSA) {
		var ret bool
		return ret
	}
	return *o.UseVMSA
}

// GetUseVMSAOk returns a tuple with the UseVMSA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSEndpointConfig) GetUseVMSAOk() (*bool, bool) {
	if o == nil || IsNil(o.UseVMSA) {
		return nil, false
	}
	return o.UseVMSA, true
}

// HasUseVMSA returns a boolean if a field has been set.
func (o *GCSEndpointConfig) HasUseVMSA() bool {
	if o != nil && !IsNil(o.UseVMSA) {
		return true
	}

	return false
}

// SetUseVMSA gets a reference to the given bool and assigns it to the UseVMSA field.
func (o *GCSEndpointConfig) SetUseVMSA(v bool) {
	o.UseVMSA = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCSEndpointConfig) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSEndpointConfig) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCSEndpointConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GCSEndpointConfig) SetType(v string) {
	o.Type = &v
}

func (o GCSEndpointConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCSEndpointConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.UseVMSA) {
		toSerialize["useVMSA"] = o.UseVMSA
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGCSEndpointConfig struct {
	value *GCSEndpointConfig
	isSet bool
}

func (v NullableGCSEndpointConfig) Get() *GCSEndpointConfig {
	return v.value
}

func (v *NullableGCSEndpointConfig) Set(val *GCSEndpointConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGCSEndpointConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGCSEndpointConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCSEndpointConfig(val *GCSEndpointConfig) *NullableGCSEndpointConfig {
	return &NullableGCSEndpointConfig{value: val, isSet: true}
}

func (v NullableGCSEndpointConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCSEndpointConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

