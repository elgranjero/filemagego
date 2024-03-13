# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username. | [optional] 
**Password** | Pointer to **string** | Password. | [optional] 
**EndpointId** | Pointer to **int32** | ID of endpoint to associate with user. | [optional] 
**Email** | Pointer to **string** | The email address used for password reset and welcome emails. | [optional] 
**Disabled** | Pointer to **bool** | Password authentication disabled. | [optional] 
**Home** | Pointer to [**UserHome**](UserHome.md) |  | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | Additional folder permissions assigned to the user. | [optional] 
**Whitelist** | Pointer to **[]string** | IP addresses which are allowed to connect as this user. | [optional] 
**MfaRequired** | Pointer to **bool** | Require multi-factor authentication to be set up before allowing user to log in. | [optional] 
**Ldap** | Pointer to **bool** | Indicates that this user should be authneticated using LDAP. | [optional] 
**Expires** | Pointer to **string** | A timestamp in ISO 8601 format indicating when the users account should no longer be accessible. | [optional] 
**FtpDisabled** | Pointer to **bool** | Disable FTP protocol for this user. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEndpointId

`func (o *User) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *User) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *User) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *User) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisabled

`func (o *User) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *User) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *User) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *User) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHome

`func (o *User) GetHome() UserHome`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *User) GetHomeOk() (*UserHome, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *User) SetHome(v UserHome)`

SetHome sets Home field to given value.

### HasHome

`func (o *User) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetPermissions

`func (o *User) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *User) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetWhitelist

`func (o *User) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *User) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *User) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *User) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### GetMfaRequired

`func (o *User) GetMfaRequired() bool`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *User) GetMfaRequiredOk() (*bool, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *User) SetMfaRequired(v bool)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *User) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetLdap

`func (o *User) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *User) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *User) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *User) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetExpires

`func (o *User) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *User) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *User) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *User) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFtpDisabled

`func (o *User) GetFtpDisabled() bool`

GetFtpDisabled returns the FtpDisabled field if non-nil, zero value otherwise.

### GetFtpDisabledOk

`func (o *User) GetFtpDisabledOk() (*bool, bool)`

GetFtpDisabledOk returns a tuple with the FtpDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpDisabled

`func (o *User) SetFtpDisabled(v bool)`

SetFtpDisabled sets FtpDisabled field to given value.

### HasFtpDisabled

`func (o *User) HasFtpDisabled() bool`

HasFtpDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


