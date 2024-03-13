# CreateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of returned endpoint. | [optional] 
**Name** | Pointer to **string** | Name of returned endpoint. | [optional] 
**Config** | Pointer to [**GCSEndpointConfig**](GCSEndpointConfig.md) |  | [optional] 

## Methods

### NewCreateEndpointRequest

`func NewCreateEndpointRequest() *CreateEndpointRequest`

NewCreateEndpointRequest instantiates a new CreateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointRequestWithDefaults

`func NewCreateEndpointRequestWithDefaults() *CreateEndpointRequest`

NewCreateEndpointRequestWithDefaults instantiates a new CreateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEndpointRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEndpointRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEndpointRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEndpointRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateEndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEndpointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEndpointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *CreateEndpointRequest) GetConfig() GCSEndpointConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateEndpointRequest) GetConfigOk() (*GCSEndpointConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateEndpointRequest) SetConfig(v GCSEndpointConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateEndpointRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


