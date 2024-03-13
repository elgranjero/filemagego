# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path of the folder the permission is applied to. | [optional] 
**Sub** | Pointer to **bool** | Enable to recursively apply permissions to all sub-folders. | [optional] 
**Grants** | Pointer to **string** | Specifies the permission level granted on the path. | [optional] 
**TargetPath** | Pointer to **string** | When a target path is specified, any operations on the path of the permission folder are aliased to the target path, which can be on a different storage endpoint. | [optional] 
**EndpointId** | Pointer to **int32** | Required when target path is specified. The endpoint ID of storage endpoint to which the target path is aliased. | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Permission) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Permission) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Permission) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Permission) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSub

`func (o *Permission) GetSub() bool`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *Permission) GetSubOk() (*bool, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *Permission) SetSub(v bool)`

SetSub sets Sub field to given value.

### HasSub

`func (o *Permission) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetGrants

`func (o *Permission) GetGrants() string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *Permission) GetGrantsOk() (*string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *Permission) SetGrants(v string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *Permission) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetTargetPath

`func (o *Permission) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *Permission) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *Permission) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *Permission) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetEndpointId

`func (o *Permission) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *Permission) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *Permission) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *Permission) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


