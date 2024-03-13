# GroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Group ID. | [optional] 
**Name** | Pointer to **string** | Name of the group. | [optional] 
**Members** | Pointer to **int32** | Number of members in the group. | [optional] 

## Methods

### NewGroupInfo

`func NewGroupInfo() *GroupInfo`

NewGroupInfo instantiates a new GroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInfoWithDefaults

`func NewGroupInfoWithDefaults() *GroupInfo`

NewGroupInfoWithDefaults instantiates a new GroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GroupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *GroupInfo) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupInfo) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupInfo) SetMembers(v int32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupInfo) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


