/*
ConductorOne API

Testing RequestCatalogSearchAPIService

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

func Test_c1api_RequestCatalogSearchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RequestCatalogSearchAPIService C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RequestCatalogSearchAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}