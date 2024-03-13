/*
FileMage Gateway

Testing EndpointsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package filemagego

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	fm "github.com/elgranjero/filemagego"
)

func Test_filemagego_EndpointsAPIService(t *testing.T) {

	configuration := fm.NewConfiguration()
	apiClient := fm.NewAPIClient(configuration)

	t.Run("Test EndpointsAPIService CreateEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpoint(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpoint(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpoint(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpoints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpoint(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
