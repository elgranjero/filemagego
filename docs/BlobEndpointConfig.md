# BlobEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | Pointer to **string** | Azure storage account secret key. | [optional] 
**AccountName** | Pointer to **string** | Azure storage account name. | [optional] 
**ContainerName** | Pointer to **string** | Azure storage Blob container. | [optional] 
**IsHNS** | Pointer to **bool** | Set to true if the Blob container has Hierarchical namespaces enabled. | [optional] 
**UseMSI** | Pointer to **bool** | Enabled to use the managed service identify associated with the instance to authenticate. | [optional] 
**UserAssignedId** | Pointer to **string** | Client ID of a user assigned managed system identity to use for authentication. | [optional] 
**Type** | Pointer to **string** | Storage endpoint type. Use \&quot;azure\&quot; when connecting to Azure Blob Storage. | [optional] 

## Methods

### NewBlobEndpointConfig

`func NewBlobEndpointConfig() *BlobEndpointConfig`

NewBlobEndpointConfig instantiates a new BlobEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobEndpointConfigWithDefaults

`func NewBlobEndpointConfigWithDefaults() *BlobEndpointConfig`

NewBlobEndpointConfigWithDefaults instantiates a new BlobEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *BlobEndpointConfig) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *BlobEndpointConfig) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *BlobEndpointConfig) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *BlobEndpointConfig) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetAccountName

`func (o *BlobEndpointConfig) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *BlobEndpointConfig) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *BlobEndpointConfig) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *BlobEndpointConfig) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetContainerName

`func (o *BlobEndpointConfig) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *BlobEndpointConfig) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *BlobEndpointConfig) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *BlobEndpointConfig) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetIsHNS

`func (o *BlobEndpointConfig) GetIsHNS() bool`

GetIsHNS returns the IsHNS field if non-nil, zero value otherwise.

### GetIsHNSOk

`func (o *BlobEndpointConfig) GetIsHNSOk() (*bool, bool)`

GetIsHNSOk returns a tuple with the IsHNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHNS

`func (o *BlobEndpointConfig) SetIsHNS(v bool)`

SetIsHNS sets IsHNS field to given value.

### HasIsHNS

`func (o *BlobEndpointConfig) HasIsHNS() bool`

HasIsHNS returns a boolean if a field has been set.

### GetUseMSI

`func (o *BlobEndpointConfig) GetUseMSI() bool`

GetUseMSI returns the UseMSI field if non-nil, zero value otherwise.

### GetUseMSIOk

`func (o *BlobEndpointConfig) GetUseMSIOk() (*bool, bool)`

GetUseMSIOk returns a tuple with the UseMSI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMSI

`func (o *BlobEndpointConfig) SetUseMSI(v bool)`

SetUseMSI sets UseMSI field to given value.

### HasUseMSI

`func (o *BlobEndpointConfig) HasUseMSI() bool`

HasUseMSI returns a boolean if a field has been set.

### GetUserAssignedId

`func (o *BlobEndpointConfig) GetUserAssignedId() string`

GetUserAssignedId returns the UserAssignedId field if non-nil, zero value otherwise.

### GetUserAssignedIdOk

`func (o *BlobEndpointConfig) GetUserAssignedIdOk() (*string, bool)`

GetUserAssignedIdOk returns a tuple with the UserAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignedId

`func (o *BlobEndpointConfig) SetUserAssignedId(v string)`

SetUserAssignedId sets UserAssignedId field to given value.

### HasUserAssignedId

`func (o *BlobEndpointConfig) HasUserAssignedId() bool`

HasUserAssignedId returns a boolean if a field has been set.

### GetType

`func (o *BlobEndpointConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobEndpointConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobEndpointConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlobEndpointConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


