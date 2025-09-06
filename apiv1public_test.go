// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/manifoldfinance/openmev"
	"github.com/manifoldfinance/openmev/internal/testutil"
	"github.com/manifoldfinance/openmev/option"
)

func TestAPIV1PublicGetOverviewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.API.V1.Public.GetOverview(context.TODO(), openmev.APIV1PublicGetOverviewParams{
		CompletedAuctions:  openmev.Int(1),
		IncludeMarketStats: openmev.Bool(true),
		IncludeSlotTiming:  openmev.Bool(true),
		NextAuctions:       openmev.Int(1),
		TimeRangeHours:     openmev.Int(1),
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
