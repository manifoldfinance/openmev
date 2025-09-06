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

func TestAPIV1AuctionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Auctions.New(context.TODO(), openmev.APIV1AuctionNewParams{
		Duration: 1000,
		ElasticSupply: openmev.APIV1AuctionNewParamsElasticSupply{
			BasePrice:   "basePrice",
			Elasticity:  0,
			MaxCapacity: 0,
			MinCapacity: 0,
		},
		MinBid:         "minBid",
		ScheduledStart: 0,
		Slot:           0,
		Metadata: map[string]any{
			"foo": "bar",
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

func TestAPIV1AuctionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Auctions.Get(
		context.TODO(),
		"x",
		openmev.APIV1AuctionGetParams{
			IncludeAnalytics: openmev.Bool(true),
			IncludeBids:      openmev.Bool(true),
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

func TestAPIV1AuctionListWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Auctions.List(context.TODO(), openmev.APIV1AuctionListParams{
		CreatedBy:       openmev.String("created_by"),
		Cursor:          openmev.String("cursor"),
		FromDate:        openmev.Float(0),
		IncludeMetadata: openmev.Bool(true),
		IncludeRevenue:  openmev.Bool(true),
		Limit:           openmev.Int(1),
		SlotFrom:        openmev.Int(0),
		SlotTo:          openmev.Int(0),
		SortBy:          openmev.APIV1AuctionListParamsSortByCreatedAt,
		SortOrder:       openmev.APIV1AuctionListParamsSortOrderAsc,
		State:           openmev.APIV1AuctionListParamsStateScheduled,
		ToDate:          openmev.Float(0),
	})
	if err != nil {
		var apierr *openmev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
