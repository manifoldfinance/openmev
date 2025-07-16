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

func TestAPIV1AuthRegisterMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Auth.Register.Member(context.TODO(), openmev.APIV1AuthRegisterMemberParams{
		InviteCode: "xxxxxxxxxxxxxxxx",
		User: openmev.APIV1AuthRegisterMemberParamsUser{
			Email:      "dev@stainless.com",
			EthAddress: openmev.String("0x2c02efDd09B3BA1AEaDd3dCAa7aC7A37C1CBDA8A"),
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

func TestAPIV1AuthRegisterOrganizationWithOptionalParams(t *testing.T) {
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
	_, err := client.API.V1.Auth.Register.Organization(context.TODO(), openmev.APIV1AuthRegisterOrganizationParams{
		InviteCode: "xxxxxxxxxxxxxxxx",
		Organization: openmev.APIV1AuthRegisterOrganizationParamsOrganization{
			Name: "xxx",
			Settings: openmev.APIV1AuthRegisterOrganizationParamsOrganizationSettings{
				APIKeyQuota:    openmev.Int(1),
				BiddingEnabled: openmev.Bool(true),
				RateLimitTier:  "basic",
			},
		},
		User: openmev.APIV1AuthRegisterOrganizationParamsUser{
			Email:      "dev@stainless.com",
			EthAddress: openmev.String("0x2c02efDd09B3BA1AEaDd3dCAa7aC7A37C1CBDA8A"),
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
