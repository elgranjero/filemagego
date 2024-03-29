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

// checks if the UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfo{}

// UserInfo struct for UserInfo
type UserInfo struct {
	// ID of returned user.
	Id *int32 `json:"id,omitempty"`
	// Username
	Username *string `json:"username,omitempty"`
	// Name of endpoint associated with user.
	EndpointName *string `json:"endpointName,omitempty"`
}

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo() *UserInfo {
	this := UserInfo{}
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserInfo) SetId(v int32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserInfo) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserInfo) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserInfo) SetUsername(v string) {
	o.Username = &v
}

// GetEndpointName returns the EndpointName field value if set, zero value otherwise.
func (o *UserInfo) GetEndpointName() string {
	if o == nil || IsNil(o.EndpointName) {
		var ret string
		return ret
	}
	return *o.EndpointName
}

// GetEndpointNameOk returns a tuple with the EndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetEndpointNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointName) {
		return nil, false
	}
	return o.EndpointName, true
}

// HasEndpointName returns a boolean if a field has been set.
func (o *UserInfo) HasEndpointName() bool {
	if o != nil && !IsNil(o.EndpointName) {
		return true
	}

	return false
}

// SetEndpointName gets a reference to the given string and assigns it to the EndpointName field.
func (o *UserInfo) SetEndpointName(v string) {
	o.EndpointName = &v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.EndpointName) {
		toSerialize["endpointName"] = o.EndpointName
	}
	return toSerialize, nil
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


