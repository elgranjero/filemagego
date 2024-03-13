# KeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Key ID. | [optional] 
**Fingerprint** | Pointer to **string** | MD5 key fingerprint. | [optional] 
**KeyData** | Pointer to **string** | The public key associated with the key pair. | [optional] 
**Title** | Pointer to **string** | A title to display for the key. | [optional] 
**CreatedAt** | Pointer to **string** | A UTC timestamp of when the key was created. | [optional] 

## Methods

### NewKeyInfo

`func NewKeyInfo() *KeyInfo`

NewKeyInfo instantiates a new KeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoWithDefaults

`func NewKeyInfoWithDefaults() *KeyInfo`

NewKeyInfoWithDefaults instantiates a new KeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *KeyInfo) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KeyInfo) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KeyInfo) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KeyInfo) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKeyData

`func (o *KeyInfo) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *KeyInfo) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *KeyInfo) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *KeyInfo) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetTitle

`func (o *KeyInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KeyInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KeyInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KeyInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KeyInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeyInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


