# PermissionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Grants** | Pointer to **string** | Specifies the permission level granted on the path. | [optional] 
**Sub** | Pointer to **bool** | Enable to recursively apply permissions to all sub-folders. | [optional] 
**Name** | Pointer to **string** | The name of the principal this permission is applied to. | [optional] 
**Type** | Pointer to **string** | Indicates if the principal is a user or group. | [optional] 

## Methods

### NewPermissionDetail

`func NewPermissionDetail() *PermissionDetail`

NewPermissionDetail instantiates a new PermissionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDetailWithDefaults

`func NewPermissionDetailWithDefaults() *PermissionDetail`

NewPermissionDetailWithDefaults instantiates a new PermissionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGrants

`func (o *PermissionDetail) GetGrants() string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *PermissionDetail) GetGrantsOk() (*string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *PermissionDetail) SetGrants(v string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *PermissionDetail) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetSub

`func (o *PermissionDetail) GetSub() bool`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *PermissionDetail) GetSubOk() (*bool, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *PermissionDetail) SetSub(v bool)`

SetSub sets Sub field to given value.

### HasSub

`func (o *PermissionDetail) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetName

`func (o *PermissionDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PermissionDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PermissionDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PermissionDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PermissionDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


