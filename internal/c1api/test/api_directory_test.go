/*
ConductorOne API

Testing DirectoryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package c1api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/conductorone/cone/internal/c1api"
)

func Test_c1api_DirectoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DirectoryAPIService C1ApiDirectoryV1DirectoryServiceCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectoryAPIService C1ApiDirectoryV1DirectoryServiceDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId string

		resp, httpRes, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceDelete(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectoryAPIService C1ApiDirectoryV1DirectoryServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId string

		resp, httpRes, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceGet(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectoryAPIService C1ApiDirectoryV1DirectoryServiceList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
