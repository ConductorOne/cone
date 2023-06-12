/*
ConductorOne API

Testing RolesAPIService

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

func Test_c1api_RolesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RolesAPIService C1ApiIamV1RolesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.RolesAPI.C1ApiIamV1RolesGet(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesAPIService C1ApiIamV1RolesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RolesAPI.C1ApiIamV1RolesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
