# S3Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of returned endpoint. | [optional] 
**Name** | Pointer to **string** | Name of returned endpoint. | [optional] 
**Config** | Pointer to [**S3EndpointConfig**](S3EndpointConfig.md) |  | [optional] 

## Methods

### NewS3Endpoint

`func NewS3Endpoint() *S3Endpoint`

NewS3Endpoint instantiates a new S3Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EndpointWithDefaults

`func NewS3EndpointWithDefaults() *S3Endpoint`

NewS3EndpointWithDefaults instantiates a new S3Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Endpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Endpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Endpoint) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *S3Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *S3Endpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3Endpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3Endpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3Endpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *S3Endpoint) GetConfig() S3EndpointConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *S3Endpoint) GetConfigOk() (*S3EndpointConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *S3Endpoint) SetConfig(v S3EndpointConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *S3Endpoint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


