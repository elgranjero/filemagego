# S3EndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Storage endpoint type. Use \&quot;aws\&quot; when connecting to Amazon S3. | [optional] 
**BucketName** | Pointer to **string** | The Amazon S3 bucket name. | [optional] 
**AccessKey** | Pointer to **string** | The access key to use to authenticate. | [optional] 
**SecretKey** | Pointer to **string** | The secret key to use to authenticate. | [optional] 
**UseIamRole** | Pointer to **bool** | Enabled to use the IAM role associated with the instance to authenticate. | [optional] 
**SSE** | Pointer to **string** | The server side encryption type to use when creating objects. | [optional] 
**KmsKeyArn** | Pointer to **string** | The ARN of the KMS key to use when creating objects. | [optional] 

## Methods

### NewS3EndpointConfig

`func NewS3EndpointConfig() *S3EndpointConfig`

NewS3EndpointConfig instantiates a new S3EndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EndpointConfigWithDefaults

`func NewS3EndpointConfigWithDefaults() *S3EndpointConfig`

NewS3EndpointConfigWithDefaults instantiates a new S3EndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *S3EndpointConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3EndpointConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3EndpointConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *S3EndpointConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBucketName

`func (o *S3EndpointConfig) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3EndpointConfig) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3EndpointConfig) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *S3EndpointConfig) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetAccessKey

`func (o *S3EndpointConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3EndpointConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3EndpointConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *S3EndpointConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *S3EndpointConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3EndpointConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3EndpointConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *S3EndpointConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseIamRole

`func (o *S3EndpointConfig) GetUseIamRole() bool`

GetUseIamRole returns the UseIamRole field if non-nil, zero value otherwise.

### GetUseIamRoleOk

`func (o *S3EndpointConfig) GetUseIamRoleOk() (*bool, bool)`

GetUseIamRoleOk returns a tuple with the UseIamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIamRole

`func (o *S3EndpointConfig) SetUseIamRole(v bool)`

SetUseIamRole sets UseIamRole field to given value.

### HasUseIamRole

`func (o *S3EndpointConfig) HasUseIamRole() bool`

HasUseIamRole returns a boolean if a field has been set.

### GetSSE

`func (o *S3EndpointConfig) GetSSE() string`

GetSSE returns the SSE field if non-nil, zero value otherwise.

### GetSSEOk

`func (o *S3EndpointConfig) GetSSEOk() (*string, bool)`

GetSSEOk returns a tuple with the SSE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSE

`func (o *S3EndpointConfig) SetSSE(v string)`

SetSSE sets SSE field to given value.

### HasSSE

`func (o *S3EndpointConfig) HasSSE() bool`

HasSSE returns a boolean if a field has been set.

### GetKmsKeyArn

`func (o *S3EndpointConfig) GetKmsKeyArn() string`

GetKmsKeyArn returns the KmsKeyArn field if non-nil, zero value otherwise.

### GetKmsKeyArnOk

`func (o *S3EndpointConfig) GetKmsKeyArnOk() (*string, bool)`

GetKmsKeyArnOk returns a tuple with the KmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyArn

`func (o *S3EndpointConfig) SetKmsKeyArn(v string)`

SetKmsKeyArn sets KmsKeyArn field to given value.

### HasKmsKeyArn

`func (o *S3EndpointConfig) HasKmsKeyArn() bool`

HasKmsKeyArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


