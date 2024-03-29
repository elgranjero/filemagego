/*
FileMage Gateway

Testing GroupsAPIService

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

func Test_filemagego_GroupsAPIService(t *testing.T) {

	configuration := fm.NewConfiguration()
	apiClient := fm.NewAPIClient(configuration)

	t.Run("Test GroupsAPIService CreateGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GroupsAPI.CreateGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService DeleteGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GetGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.GroupsAPI.GetGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService ListGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GroupsAPI.ListGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
