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

// checks if the S3Endpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3Endpoint{}

// S3Endpoint struct for S3Endpoint
type S3Endpoint struct {
	// ID of returned endpoint.
	Id *int32 `json:"id,omitempty"`
	// Name of returned endpoint.
	Name *string `json:"name,omitempty"`
	Config *S3EndpointConfig `json:"config,omitempty"`
}

// NewS3Endpoint instantiates a new S3Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Endpoint() *S3Endpoint {
	this := S3Endpoint{}
	return &this
}

// NewS3EndpointWithDefaults instantiates a new S3Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3EndpointWithDefaults() *S3Endpoint {
	this := S3Endpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *S3Endpoint) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Endpoint) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *S3Endpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *S3Endpoint) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *S3Endpoint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Endpoint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *S3Endpoint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *S3Endpoint) SetName(v string) {
	o.Name = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *S3Endpoint) GetConfig() S3EndpointConfig {
	if o == nil || IsNil(o.Config) {
		var ret S3EndpointConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Endpoint) GetConfigOk() (*S3EndpointConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *S3Endpoint) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given S3EndpointConfig and assigns it to the Config field.
func (o *S3Endpoint) SetConfig(v S3EndpointConfig) {
	o.Config = &v
}

func (o S3Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3Endpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableS3Endpoint struct {
	value *S3Endpoint
	isSet bool
}

func (v NullableS3Endpoint) Get() *S3Endpoint {
	return v.value
}

func (v *NullableS3Endpoint) Set(val *S3Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Endpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Endpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Endpoint(val *S3Endpoint) *NullableS3Endpoint {
	return &NullableS3Endpoint{value: val, isSet: true}
}

func (v NullableS3Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Endpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

