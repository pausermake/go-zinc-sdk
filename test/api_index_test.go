/*
Zinc Search engine API

Testing IndexAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_IndexAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IndexAPIService AddOrRemoveESAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.AddOrRemoveESAlias(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService Analyze", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.Analyze(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService AnalyzeIndex", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.AnalyzeIndex(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService CreateIndex", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.CreateIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService CreateTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.CreateTemplate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService DeleteIndex", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.DeleteIndex(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService DeleteTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.IndexAPI.DeleteTemplate(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService ESCreateIndex", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.ESCreateIndex(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService ESGetMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.ESGetMapping(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService EsExists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.EsExists(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService Exists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.Exists(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService GetESAliases", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var target string
		var targetAlias string

		resp, httpRes, err := apiClient.IndexAPI.GetESAliases(context.Background(), target, targetAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService GetIndex", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.GetIndex(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService GetMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.GetMapping(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService GetSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.GetSettings(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService GetTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.IndexAPI.GetTemplate(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService IndexNameList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.IndexNameList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService ListIndexes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.ListIndexes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService ListTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IndexAPI.ListTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService Refresh", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.Refresh(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService SetMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.SetMapping(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService SetSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var index string

		resp, httpRes, err := apiClient.IndexAPI.SetSettings(context.Background(), index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IndexAPIService UpdateTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.IndexAPI.UpdateTemplate(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
