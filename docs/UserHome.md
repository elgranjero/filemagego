# UserHome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path of user home directory. | [optional] 
**Sub** | Pointer to **bool** | Recursively apply permissions to all sub-folders. | [optional] 
**Grants** | Pointer to **string** | Specifies permission level granted on home directory. Allowed values are &#39;list&#39;, &#39;read&#39;, &#39;write&#39;, &#39;full&#39;. | [optional] 

## Methods

### NewUserHome

`func NewUserHome() *UserHome`

NewUserHome instantiates a new UserHome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserHomeWithDefaults

`func NewUserHomeWithDefaults() *UserHome`

NewUserHomeWithDefaults instantiates a new UserHome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *UserHome) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UserHome) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UserHome) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UserHome) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSub

`func (o *UserHome) GetSub() bool`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UserHome) GetSubOk() (*bool, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UserHome) SetSub(v bool)`

SetSub sets Sub field to given value.

### HasSub

`func (o *UserHome) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetGrants

`func (o *UserHome) GetGrants() string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *UserHome) GetGrantsOk() (*string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *UserHome) SetGrants(v string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *UserHome) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


