// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/openmev-go/internal/apijson"
	"github.com/stainless-sdks/openmev-go/internal/apiquery"
	"github.com/stainless-sdks/openmev-go/internal/requestconfig"
	"github.com/stainless-sdks/openmev-go/option"
	"github.com/stainless-sdks/openmev-go/packages/param"
	"github.com/stainless-sdks/openmev-go/packages/respjson"
)

// APIV1BidService contains methods and other services that help with interacting
// with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1BidService] method instead.
type APIV1BidService struct {
	Options []option.RequestOption
}

// NewAPIV1BidService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1BidService(opts ...option.RequestOption) (r APIV1BidService) {
	r = APIV1BidService{}
	r.Options = opts
	return
}

// Retrieve the current status of a bid using its tracking ID. Includes processing
// information and settlement details. Requires authentication with read
// permissions.
func (r *APIV1BidService) GetStatus(ctx context.Context, trackingID string, query APIV1BidGetStatusParams, opts ...option.RequestOption) (res *Apiv1BidGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/bids/%s/status", trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1BidGetStatusResponse struct {
	Bid Apiv1BidGetStatusResponseBid `json:"bid,required"`
	// Any of "pending", "queued", "processing", "accepted", "rejected", "partial".
	Status Apiv1BidGetStatusResponseStatus `json:"status,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// Tracking identifier
	TrackingID string                              `json:"trackingId,required"`
	Processing Apiv1BidGetStatusResponseProcessing `json:"processing"`
	Settlement Apiv1BidGetStatusResponseSettlement `json:"settlement"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bid         respjson.Field
		Status      respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		TrackingID  respjson.Field
		Processing  respjson.Field
		Settlement  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1BidGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1BidGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1BidGetStatusResponseBid struct {
	// Bid amount in wei
	Amount string `json:"amount,required"`
	// Associated auction ID
	AuctionID string `json:"auctionId,required"`
	Quantity  int64  `json:"quantity,required"`
	// Bid submission timestamp
	Timestamp float64 `json:"timestamp,required"`
	// Bid identifier
	ID string `json:"id"`
	// Allocated quantity
	AllocatedQuantity int64 `json:"allocatedQuantity"`
	// Final bid amount in wei
	FinalAmount string `json:"finalAmount"`
	// Processing completion timestamp
	ProcessedAt float64 `json:"processedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		AuctionID         respjson.Field
		Quantity          respjson.Field
		Timestamp         respjson.Field
		ID                respjson.Field
		AllocatedQuantity respjson.Field
		FinalAmount       respjson.Field
		ProcessedAt       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1BidGetStatusResponseBid) RawJSON() string { return r.JSON.raw }
func (r *Apiv1BidGetStatusResponseBid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1BidGetStatusResponseStatus string

const (
	Apiv1BidGetStatusResponseStatusPending    Apiv1BidGetStatusResponseStatus = "pending"
	Apiv1BidGetStatusResponseStatusQueued     Apiv1BidGetStatusResponseStatus = "queued"
	Apiv1BidGetStatusResponseStatusProcessing Apiv1BidGetStatusResponseStatus = "processing"
	Apiv1BidGetStatusResponseStatusAccepted   Apiv1BidGetStatusResponseStatus = "accepted"
	Apiv1BidGetStatusResponseStatusRejected   Apiv1BidGetStatusResponseStatus = "rejected"
	Apiv1BidGetStatusResponseStatusPartial    Apiv1BidGetStatusResponseStatus = "partial"
)

type Apiv1BidGetStatusResponseProcessing struct {
	// Number of processing attempts
	Attempts int64 `json:"attempts,required"`
	// Estimated processing time in milliseconds
	EstimatedTime float64 `json:"estimatedTime,required"`
	// Current queue position
	QueuePosition int64 `json:"queuePosition,required"`
	// Last error message if any
	LastError string `json:"lastError"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempts      respjson.Field
		EstimatedTime respjson.Field
		QueuePosition respjson.Field
		LastError     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1BidGetStatusResponseProcessing) RawJSON() string { return r.JSON.raw }
func (r *Apiv1BidGetStatusResponseProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1BidGetStatusResponseSettlement struct {
	// Whether bid has been settled
	IsSettled bool `json:"isSettled,required"`
	// Clearing price in wei
	ClearingPrice string `json:"clearingPrice"`
	// Refund amount in wei
	RefundAmount string `json:"refundAmount"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSettled     respjson.Field
		ClearingPrice respjson.Field
		RefundAmount  respjson.Field
		SettledAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1BidGetStatusResponseSettlement) RawJSON() string { return r.JSON.raw }
func (r *Apiv1BidGetStatusResponseSettlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1BidGetStatusParams struct {
	// Include detailed processing information
	IncludeDetails param.Opt[bool] `query:"include_details,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1BidGetStatusParams]'s query parameters as
// `url.Values`.
func (r APIV1BidGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
