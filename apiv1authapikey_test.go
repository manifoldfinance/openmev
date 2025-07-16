// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/openmev-go"
	"github.com/stainless-sdks/openmev-go/internal/testutil"
	"github.com/stainless-sdks/openmev-go/option"
)

func TestAPIV1AuthAPIKeyAPIKeysWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.V1.Auth.APIKeys.APIKeys(context.TODO(), openmev.APIV1AuthAPIKeyAPIKeysParams{
		Name:        "x",
		Permissions: []string{"read"},
		Description: openmev.String("description"),
		ExpiresAt:   openmev.Float(0),
		Metadata: map[string]any{
			"foo": "bar",
		},
		RateLimitOverride: openmev.APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride{
			RequestsPerDay:    openmev.Int(1),
			RequestsPerHour:   openmev.Int(1),
			RequestsPerMinute: openmev.Int(1),
		},
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV1AuthAPIKeyGetAPIKeysWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.V1.Auth.APIKeys.GetAPIKeys(context.TODO(), openmev.APIV1AuthAPIKeyGetAPIKeysParams{
		CreatedBy:       openmev.String("created_by"),
		Cursor:          openmev.String("cursor"),
		IncludeInactive: openmev.Bool(true),
		Limit:           openmev.Int(1),
		Permission:      openmev.APIV1AuthAPIKeyGetAPIKeysParamsPermissionRead,
		SortBy:          openmev.APIV1AuthAPIKeyGetAPIKeysParamsSortByName,
		SortOrder:       openmev.APIV1AuthAPIKeyGetAPIKeysParamsSortOrderAsc,
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
