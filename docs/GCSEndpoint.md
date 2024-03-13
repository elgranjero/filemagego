# GCSEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of returned endpoint. | [optional] 
**Name** | Pointer to **string** | Name of returned endpoint. | [optional] 
**Config** | Pointer to [**GCSEndpointConfig**](GCSEndpointConfig.md) |  | [optional] 

## Methods

### NewGCSEndpoint

`func NewGCSEndpoint() *GCSEndpoint`

NewGCSEndpoint instantiates a new GCSEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSEndpointWithDefaults

`func NewGCSEndpointWithDefaults() *GCSEndpoint`

NewGCSEndpointWithDefaults instantiates a new GCSEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GCSEndpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCSEndpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCSEndpoint) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GCSEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GCSEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GCSEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GCSEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GCSEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *GCSEndpoint) GetConfig() GCSEndpointConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GCSEndpoint) GetConfigOk() (*GCSEndpointConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GCSEndpoint) SetConfig(v GCSEndpointConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GCSEndpoint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


