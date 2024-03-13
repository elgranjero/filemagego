# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyData** | Pointer to **string** | Contents of key file to import. Use either keyData or create. | [optional] 
**Create** | Pointer to **bool** | Create and return a new key. Use either keyData or create. | [optional] [default to false]
**Title** | Pointer to **string** | A title to display for the key. | [optional] 

## Methods

### NewKey

`func NewKey() *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyData

`func (o *Key) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *Key) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *Key) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *Key) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetCreate

`func (o *Key) GetCreate() bool`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Key) GetCreateOk() (*bool, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Key) SetCreate(v bool)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Key) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetTitle

`func (o *Key) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Key) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Key) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Key) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


