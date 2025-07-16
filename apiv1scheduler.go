// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"net/http"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// APIV1SchedulerService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1SchedulerService] method instead.
type APIV1SchedulerService struct {
	Options []option.RequestOption
}

// NewAPIV1SchedulerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1SchedulerService(opts ...option.RequestOption) (r APIV1SchedulerService) {
	r = APIV1SchedulerService{}
	r.Options = opts
	return
}

// Get current status of the auction scheduler. Requires authentication.
func (r *APIV1SchedulerService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *Apiv1SchedulerGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/scheduler/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Manually trigger the auction scheduler to create new auctions. Requires auction
// manager permissions.
func (r *APIV1SchedulerService) Trigger(ctx context.Context, opts ...option.RequestOption) (res *Apiv1SchedulerTriggerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/scheduler/trigger"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Apiv1SchedulerGetStatusResponse struct {
	IsRunning            bool    `json:"isRunning,required"`
	PendingAuctions      float64 `json:"pendingAuctions,required"`
	Success              bool    `json:"success,required"`
	Timestamp            float64 `json:"timestamp,required"`
	TotalAuctionsCreated float64 `json:"totalAuctionsCreated,required"`
	// Last run timestamp
	LastRun float64 `json:"lastRun"`
	// Next scheduled run timestamp
	NextRun float64 `json:"nextRun"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRunning            respjson.Field
		PendingAuctions      respjson.Field
		Success              respjson.Field
		Timestamp            respjson.Field
		TotalAuctionsCreated respjson.Field
		LastRun              respjson.Field
		NextRun              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1SchedulerGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1SchedulerGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1SchedulerTriggerResponse struct {
	AuctionsCreated float64 `json:"auctionsCreated,required"`
	Message         string  `json:"message,required"`
	Success         bool    `json:"success,required"`
	Timestamp       float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuctionsCreated respjson.Field
		Message         respjson.Field
		Success         respjson.Field
		Timestamp       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1SchedulerTriggerResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1SchedulerTriggerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
