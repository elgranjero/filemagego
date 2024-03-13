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

// checks if the GroupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInfo{}

// GroupInfo struct for GroupInfo
type GroupInfo struct {
	// Group ID.
	Id *int32 `json:"id,omitempty"`
	// Name of the group.
	Name *string `json:"name,omitempty"`
	// Number of members in the group.
	Members *int32 `json:"members,omitempty"`
}

// NewGroupInfo instantiates a new GroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInfo() *GroupInfo {
	this := GroupInfo{}
	return &this
}

// NewGroupInfoWithDefaults instantiates a new GroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInfoWithDefaults() *GroupInfo {
	this := GroupInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GroupInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupInfo) SetName(v string) {
	o.Name = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupInfo) GetMembers() int32 {
	if o == nil || IsNil(o.Members) {
		var ret int32
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfo) GetMembersOk() (*int32, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupInfo) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given int32 and assigns it to the Members field.
func (o *GroupInfo) SetMembers(v int32) {
	o.Members = &v
}

func (o GroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableGroupInfo struct {
	value *GroupInfo
	isSet bool
}

func (v NullableGroupInfo) Get() *GroupInfo {
	return v.value
}

func (v *NullableGroupInfo) Set(val *GroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInfo(val *GroupInfo) *NullableGroupInfo {
	return &NullableGroupInfo{value: val, isSet: true}
}

func (v NullableGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

