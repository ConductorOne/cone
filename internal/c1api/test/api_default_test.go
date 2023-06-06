/*
ConductorOne API

Testing DefaultAPIService

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

func Test_c1api_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService C1ApiAppV1AppEntitlementUserBindingServiceGetAppEntitlementUserBindingsForIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId string
		var appEntitlementId string
		var identityUserId string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiAppV1AppEntitlementUserBindingServiceGetAppEntitlementUserBindingsForIdentity(context.Background(), appId, appEntitlementId, identityUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiAppV1AppResourceServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId string
		var appResourceTypeId string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiAppV1AppResourceServiceGet(context.Background(), appId, appResourceTypeId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiAppV1AppResourceTypeServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet(context.Background(), appId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiAppV1AppsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiAppV1AppsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiAuthV1AuthIntrospect", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiAuthV1AuthIntrospect(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiTaskV1TaskSearchServiceSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiTaskV1TaskSearchServiceSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiTaskV1TaskServiceCreateGrantTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiTaskV1TaskServiceCreateRevokeTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiTaskV1TaskServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService C1ApiUserV1UserServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DefaultAPI.C1ApiUserV1UserServiceGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
