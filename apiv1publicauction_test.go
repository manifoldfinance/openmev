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

func TestAPIV1PublicAuctionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Public.Auctions.Get(
		context.TODO(),
		"x",
		openmev.APIV1PublicAuctionGetParams{
			IncludeAnalytics: openmev.Bool(true),
			IncludeBidStats:  openmev.Bool(true),
		},
	)
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV1PublicAuctionGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Public.Auctions.GetHistory(context.TODO(), openmev.APIV1PublicAuctionGetHistoryParams{
		Cursor:       openmev.String("cursor"),
		FromDate:     openmev.Float(0),
		IncludeStats: openmev.Bool(true),
		Limit:        openmev.Int(1),
		SlotFrom:     openmev.Int(0),
		SlotTo:       openmev.Int(0),
		SortBy:       openmev.APIV1PublicAuctionGetHistoryParamsSortByCreatedAt,
		SortOrder:    openmev.APIV1PublicAuctionGetHistoryParamsSortOrderAsc,
		State:        openmev.APIV1PublicAuctionGetHistoryParamsStateScheduled,
		ToDate:       openmev.Float(0),
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV1PublicAuctionGetLatestResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Public.Auctions.GetLatestResults(context.TODO(), openmev.APIV1PublicAuctionGetLatestResultsParams{
		Cursor:                 openmev.String("cursor"),
		IncludeMarketMetrics:   openmev.Bool(true),
		IncludeMetrics:         openmev.Bool(true),
		IncludePerformanceData: openmev.Bool(true),
		IncludeTrendAnalysis:   openmev.Bool(true),
		Limit:                  openmev.Int(1),
		SlotFrom:               openmev.Int(0),
		SlotTo:                 openmev.Int(0),
		SortBy:                 openmev.APIV1PublicAuctionGetLatestResultsParamsSortBySettledAt,
		SortOrder:              openmev.APIV1PublicAuctionGetLatestResultsParamsSortOrderAsc,
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
