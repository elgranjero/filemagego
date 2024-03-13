# EndpointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the endpoint. | [optional] 
**Name** | Pointer to **string** | Name of the endpoint. | [optional] 

## Methods

### NewEndpointInfo

`func NewEndpointInfo() *EndpointInfo`

NewEndpointInfo instantiates a new EndpointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInfoWithDefaults

`func NewEndpointInfoWithDefaults() *EndpointInfo`

NewEndpointInfoWithDefaults instantiates a new EndpointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EndpointInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


