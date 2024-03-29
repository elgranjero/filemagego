/*
FileMage Gateway

FileMage Gateway provides a simple API to configure storage endpoints, users, and keys. Please note that all urls must end in a trailing slash, and the Content-Type header must be set to application/json.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package filemagego

import (
	"encoding/json"
	"fmt"
)

// CreateEndpointRequest - struct for CreateEndpointRequest
type CreateEndpointRequest struct {
	BlobEndpoint *BlobEndpoint
	GCSEndpoint *GCSEndpoint
	S3Endpoint *S3Endpoint
}

// BlobEndpointAsCreateEndpointRequest is a convenience function that returns BlobEndpoint wrapped in CreateEndpointRequest
func BlobEndpointAsCreateEndpointRequest(v *BlobEndpoint) CreateEndpointRequest {
	return CreateEndpointRequest{
		BlobEndpoint: v,
	}
}

// GCSEndpointAsCreateEndpointRequest is a convenience function that returns GCSEndpoint wrapped in CreateEndpointRequest
func GCSEndpointAsCreateEndpointRequest(v *GCSEndpoint) CreateEndpointRequest {
	return CreateEndpointRequest{
		GCSEndpoint: v,
	}
}

// S3EndpointAsCreateEndpointRequest is a convenience function that returns S3Endpoint wrapped in CreateEndpointRequest
func S3EndpointAsCreateEndpointRequest(v *S3Endpoint) CreateEndpointRequest {
	return CreateEndpointRequest{
		S3Endpoint: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEndpointRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlobEndpoint
	err = newStrictDecoder(data).Decode(&dst.BlobEndpoint)
	if err == nil {
		jsonBlobEndpoint, _ := json.Marshal(dst.BlobEndpoint)
		if string(jsonBlobEndpoint) == "{}" { // empty struct
			dst.BlobEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.BlobEndpoint = nil
	}

	// try to unmarshal data into GCSEndpoint
	err = newStrictDecoder(data).Decode(&dst.GCSEndpoint)
	if err == nil {
		jsonGCSEndpoint, _ := json.Marshal(dst.GCSEndpoint)
		if string(jsonGCSEndpoint) == "{}" { // empty struct
			dst.GCSEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.GCSEndpoint = nil
	}

	// try to unmarshal data into S3Endpoint
	err = newStrictDecoder(data).Decode(&dst.S3Endpoint)
	if err == nil {
		jsonS3Endpoint, _ := json.Marshal(dst.S3Endpoint)
		if string(jsonS3Endpoint) == "{}" { // empty struct
			dst.S3Endpoint = nil
		} else {
			match++
		}
	} else {
		dst.S3Endpoint = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlobEndpoint = nil
		dst.GCSEndpoint = nil
		dst.S3Endpoint = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateEndpointRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateEndpointRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	if src.BlobEndpoint != nil {
		return json.Marshal(&src.BlobEndpoint)
	}

	if src.GCSEndpoint != nil {
		return json.Marshal(&src.GCSEndpoint)
	}

	if src.S3Endpoint != nil {
		return json.Marshal(&src.S3Endpoint)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEndpointRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlobEndpoint != nil {
		return obj.BlobEndpoint
	}

	if obj.GCSEndpoint != nil {
		return obj.GCSEndpoint
	}

	if obj.S3Endpoint != nil {
		return obj.S3Endpoint
	}

	// all schemas are nil
	return nil
}

type NullableCreateEndpointRequest struct {
	value *CreateEndpointRequest
	isSet bool
}

func (v NullableCreateEndpointRequest) Get() *CreateEndpointRequest {
	return v.value
}

func (v *NullableCreateEndpointRequest) Set(val *CreateEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndpointRequest(val *CreateEndpointRequest) *NullableCreateEndpointRequest {
	return &NullableCreateEndpointRequest{value: val, isSet: true}
}

func (v NullableCreateEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


