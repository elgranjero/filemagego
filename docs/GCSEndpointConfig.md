# GCSEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The Google Cloud Storage bucket name. | [optional] 
**Credentials** | Pointer to **string** | A base64 encoded Google Cloud JSON credentials file used to authenticate. | [optional] 
**UseVMSA** | Pointer to **bool** | Enabled to use service account associated with the instance to authenticate. | [optional] 
**Type** | Pointer to **string** | Storage endpoint type. Use \&quot;gcp\&quot; when connecting to Google Cloud Storage. | [optional] 

## Methods

### NewGCSEndpointConfig

`func NewGCSEndpointConfig() *GCSEndpointConfig`

NewGCSEndpointConfig instantiates a new GCSEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSEndpointConfigWithDefaults

`func NewGCSEndpointConfigWithDefaults() *GCSEndpointConfig`

NewGCSEndpointConfigWithDefaults instantiates a new GCSEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *GCSEndpointConfig) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GCSEndpointConfig) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GCSEndpointConfig) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *GCSEndpointConfig) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCredentials

`func (o *GCSEndpointConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCSEndpointConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCSEndpointConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GCSEndpointConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetUseVMSA

`func (o *GCSEndpointConfig) GetUseVMSA() bool`

GetUseVMSA returns the UseVMSA field if non-nil, zero value otherwise.

### GetUseVMSAOk

`func (o *GCSEndpointConfig) GetUseVMSAOk() (*bool, bool)`

GetUseVMSAOk returns a tuple with the UseVMSA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVMSA

`func (o *GCSEndpointConfig) SetUseVMSA(v bool)`

SetUseVMSA sets UseVMSA field to given value.

### HasUseVMSA

`func (o *GCSEndpointConfig) HasUseVMSA() bool`

HasUseVMSA returns a boolean if a field has been set.

### GetType

`func (o *GCSEndpointConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCSEndpointConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCSEndpointConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GCSEndpointConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


