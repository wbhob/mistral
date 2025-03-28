/*
Mistral AI API

Testing ModelsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistral

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/wbhob/mistral"
)

func Test_mistral_ModelsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ModelsAPIService DeleteModelV1ModelsModelIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelsAPI.DeleteModelV1ModelsModelIdDelete(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService JobsApiRoutesFineTuningArchiveFineTunedModel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningArchiveFineTunedModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService JobsApiRoutesFineTuningUnarchiveFineTunedModel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningUnarchiveFineTunedModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService JobsApiRoutesFineTuningUpdateFineTunedModel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelsAPI.JobsApiRoutesFineTuningUpdateFineTunedModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService ListModelsV1ModelsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModelsAPI.ListModelsV1ModelsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService RetrieveModelV1ModelsModelIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelsAPI.RetrieveModelV1ModelsModelIdGet(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
