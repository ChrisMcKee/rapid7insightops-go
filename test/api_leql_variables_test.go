/*
InsightOps REST API

Testing LEQLVariablesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package insightops

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chrismckee/rapid7insightops-go"
)

func Test_insightops_LEQLVariablesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LEQLVariablesAPIService CreateVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LEQLVariablesAPI.CreateVariable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LEQLVariablesAPIService DeleteVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LEQLVariablesAPI.DeleteVariable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LEQLVariablesAPIService ListVariables", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LEQLVariablesAPI.ListVariables(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LEQLVariablesAPIService RetrieveVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LEQLVariablesAPI.RetrieveVariable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LEQLVariablesAPIService UpdateVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LEQLVariablesAPI.UpdateVariable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}