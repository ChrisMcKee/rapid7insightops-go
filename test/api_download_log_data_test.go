/*
InsightOps REST API

Testing DownloadLogDataAPIService

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

func Test_insightops_DownloadLogDataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DownloadLogDataAPIService GetDownloadLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ids string

		httpRes, err := apiClient.DownloadLogDataAPI.GetDownloadLogs(context.Background(), ids).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
