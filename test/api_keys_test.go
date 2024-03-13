/*
FileMage Gateway

Testing KeysAPIService

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

func Test_filemagego_KeysAPIService(t *testing.T) {

	configuration := fm.NewConfiguration()
	apiClient := fm.NewAPIClient(configuration)

	t.Run("Test KeysAPIService CreateKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.KeysAPI.CreateKey(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KeysAPIService DeleteKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId int32
		var keyId int32

		httpRes, err := apiClient.KeysAPI.DeleteKey(context.Background(), userId, keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}