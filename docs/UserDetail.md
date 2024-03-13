# UserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | User ID. | [optional] 
**Username** | Pointer to **string** | Username. | [optional] 
**EndpointId** | Pointer to **int32** | Endpoint ID associated with user. | [optional] 
**Disabled** | Pointer to **bool** | Password authentication disabled. | [optional] 
**Home** | Pointer to [**UserHome**](UserHome.md) |  | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | Additional folder permissions assigned to the user. | [optional] 
**Keys** | Pointer to [**[]KeyInfo**](KeyInfo.md) | Keys associated with this user. | [optional] 
**Whitelist** | Pointer to **[]string** | IP addresses which are allowed to connect as this user. | [optional] 
**MfaRequired** | Pointer to **bool** | Require multi-factor authentication to be set up before allowing user to log in. | [optional] 
**Ldap** | Pointer to **bool** | Indicates that this user should be authneticated using LDAP. | [optional] 
**Expires** | Pointer to **string** | A timestamp in ISO 8601 format indicating when the users account should no longer be accessible. | [optional] 
**OtpEnabled** | Pointer to **bool** | Indicates that the user has configured a OTP device. | [optional] 

## Methods

### NewUserDetail

`func NewUserDetail() *UserDetail`

NewUserDetail instantiates a new UserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailWithDefaults

`func NewUserDetailWithDefaults() *UserDetail`

NewUserDetailWithDefaults instantiates a new UserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDetail) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserDetail) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpointId

`func (o *UserDetail) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *UserDetail) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *UserDetail) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *UserDetail) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetDisabled

`func (o *UserDetail) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserDetail) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserDetail) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserDetail) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHome

`func (o *UserDetail) GetHome() UserHome`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *UserDetail) GetHomeOk() (*UserHome, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *UserDetail) SetHome(v UserHome)`

SetHome sets Home field to given value.

### HasHome

`func (o *UserDetail) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetPermissions

`func (o *UserDetail) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserDetail) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserDetail) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserDetail) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetKeys

`func (o *UserDetail) GetKeys() []KeyInfo`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UserDetail) GetKeysOk() (*[]KeyInfo, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UserDetail) SetKeys(v []KeyInfo)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *UserDetail) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetWhitelist

`func (o *UserDetail) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *UserDetail) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *UserDetail) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *UserDetail) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### GetMfaRequired

`func (o *UserDetail) GetMfaRequired() bool`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *UserDetail) GetMfaRequiredOk() (*bool, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *UserDetail) SetMfaRequired(v bool)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *UserDetail) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetLdap

`func (o *UserDetail) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *UserDetail) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *UserDetail) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *UserDetail) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetExpires

`func (o *UserDetail) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UserDetail) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UserDetail) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *UserDetail) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetOtpEnabled

`func (o *UserDetail) GetOtpEnabled() bool`

GetOtpEnabled returns the OtpEnabled field if non-nil, zero value otherwise.

### GetOtpEnabledOk

`func (o *UserDetail) GetOtpEnabledOk() (*bool, bool)`

GetOtpEnabledOk returns a tuple with the OtpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpEnabled

`func (o *UserDetail) SetOtpEnabled(v bool)`

SetOtpEnabled sets OtpEnabled field to given value.

### HasOtpEnabled

`func (o *UserDetail) HasOtpEnabled() bool`

HasOtpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


