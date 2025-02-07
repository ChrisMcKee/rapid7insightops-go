/*
InsightOps REST API

Testing SavedQueriesAPIService

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

func Test_insightops_SavedQueriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SavedQueriesAPIService DeleteSavedQueryId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var savedQueryId string

		httpRes, err := apiClient.SavedQueriesAPI.DeleteSavedQueryId(context.Background(), savedQueryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService GetSavedQueryId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var savedQueryId string

		resp, httpRes, err := apiClient.SavedQueriesAPI.GetSavedQueryId(context.Background(), savedQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService ListSavedQueries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SavedQueriesAPI.ListSavedQueries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService PatchSavedQueryId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var savedQueryId string

		resp, httpRes, err := apiClient.SavedQueriesAPI.PatchSavedQueryId(context.Background(), savedQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService PostSavedQueryRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SavedQueriesAPI.PostSavedQueryRoot(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService PutSavedQueryId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var savedQueryId string

		resp, httpRes, err := apiClient.SavedQueriesAPI.PutSavedQueryId(context.Background(), savedQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService UseSavedQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var savedQueryId string

		resp, httpRes, err := apiClient.SavedQueriesAPI.UseSavedQuery(context.Background(), savedQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SavedQueriesAPIService UseSavedQueryNoLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logKeys string
		var savedQueryId string

		resp, httpRes, err := apiClient.SavedQueriesAPI.UseSavedQueryNoLogs(context.Background(), logKeys, savedQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
