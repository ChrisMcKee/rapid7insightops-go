/*
InsightOps REST API

Testing QueryLogDataAPIService

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

func Test_insightops_QueryLogDataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QueryLogDataAPIService GetContextEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var idSequenceNumber int32

		resp, httpRes, err := apiClient.QueryLogDataAPI.GetContextEvents(context.Background(), idSequenceNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService GetQueryEndpoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QueryLogDataAPI.GetQueryEndpoints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService GetQueryLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logKeys string

		resp, httpRes, err := apiClient.QueryLogDataAPI.GetQueryLogs(context.Background(), logKeys).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService GetSearchStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QueryLogDataAPI.GetSearchStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService PollQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.QueryLogDataAPI.PollQuery(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService PostQueryLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QueryLogDataAPI.PostQueryLogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService QueryLogsetsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logsetId string

		resp, httpRes, err := apiClient.QueryLogDataAPI.QueryLogsetsById(context.Background(), logsetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryLogDataAPIService QueryLogsetsByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QueryLogDataAPI.QueryLogsetsByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
