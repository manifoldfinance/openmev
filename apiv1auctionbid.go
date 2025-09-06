// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/param"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// APIV1AuctionBidService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AuctionBidService] method instead.
type APIV1AuctionBidService struct {
	Options []option.RequestOption
}

// NewAPIV1AuctionBidService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1AuctionBidService(opts ...option.RequestOption) (r APIV1AuctionBidService) {
	r = APIV1AuctionBidService{}
	r.Options = opts
	return
}

// Submit a bid for an auction with immediate processing. Requires authentication
// with bid permissions.
func (r *APIV1AuctionBidService) New(ctx context.Context, id string, body APIV1AuctionBidNewParams, opts ...option.RequestOption) (res *Apiv1AuctionBidNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/auctions/%s/bids", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submit a bid for an auction with queue-based processing. Returns a tracking ID
// for status monitoring. Requires authentication with bid permissions.
func (r *APIV1AuctionBidService) Async(ctx context.Context, id string, body APIV1AuctionBidAsyncParams, opts ...option.RequestOption) (res *Apiv1AuctionBidAsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/auctions/%s/bids/async", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Apiv1AuctionBidNewResponse struct {
	Bid Apiv1AuctionBidNewResponseBid `json:"bid,required"`
	// Success message
	Message string `json:"message,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bid         respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionBidNewResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionBidNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionBidNewResponseBid struct {
	// Unique bid identifier
	ID string `json:"id,required"`
	// Bid amount in wei
	Amount string `json:"amount,required"`
	// Associated auction ID
	AuctionID string `json:"auctionId,required"`
	// Estimated gas cost in wei
	EstimatedGasCost string `json:"estimatedGasCost,required"`
	Quantity         int64  `json:"quantity,required"`
	// Position in processing queue
	QueuePosition int64 `json:"queuePosition,required"`
	// Any of "pending", "queued", "processing", "accepted", "rejected", "partial".
	Status string `json:"status,required"`
	// Bid submission timestamp
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		AuctionID        respjson.Field
		EstimatedGasCost respjson.Field
		Quantity         respjson.Field
		QueuePosition    respjson.Field
		Status           respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionBidNewResponseBid) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionBidNewResponseBid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionBidAsyncResponse struct {
	// Estimated processing time in milliseconds
	EstimatedProcessingTime float64 `json:"estimatedProcessingTime,required"`
	// Status message
	Message string `json:"message,required"`
	// Current position in queue
	QueuePosition int64 `json:"queuePosition,required"`
	// Current processing status
	//
	// Any of "queued", "processing".
	Status Apiv1AuctionBidAsyncResponseStatus `json:"status,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// Unique tracking identifier for the bid
	TrackingID string `json:"trackingId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EstimatedProcessingTime respjson.Field
		Message                 respjson.Field
		QueuePosition           respjson.Field
		Status                  respjson.Field
		Success                 respjson.Field
		Timestamp               respjson.Field
		TrackingID              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionBidAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionBidAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current processing status
type Apiv1AuctionBidAsyncResponseStatus string

const (
	Apiv1AuctionBidAsyncResponseStatusQueued     Apiv1AuctionBidAsyncResponseStatus = "queued"
	Apiv1AuctionBidAsyncResponseStatusProcessing Apiv1AuctionBidAsyncResponseStatus = "processing"
)

type APIV1AuctionBidNewParams struct {
	// Bid amount in wei
	Amount string `json:"amount,required"`
	// Requested quantity
	Quantity int64 `json:"quantity,required"`
	// Maximum gas price in wei
	MaxGasPrice param.Opt[string] `json:"maxGasPrice,omitzero"`
	// Additional bid metadata
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r APIV1AuctionBidNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuctionBidNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuctionBidNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuctionBidAsyncParams struct {
	// Bid amount in wei
	Amount string `json:"amount,required"`
	// Requested quantity
	Quantity int64 `json:"quantity,required"`
	// Callback URL for status updates
	Callback param.Opt[string] `json:"callback,omitzero" format:"uri"`
	// Maximum gas price in wei
	MaxGasPrice param.Opt[string] `json:"maxGasPrice,omitzero"`
	// Additional bid metadata
	Metadata map[string]any `json:"metadata,omitzero"`
	// Processing priority
	//
	// Any of "low", "normal", "high".
	Priority APIV1AuctionBidAsyncParamsPriority `json:"priority,omitzero"`
	paramObj
}

func (r APIV1AuctionBidAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuctionBidAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuctionBidAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Processing priority
type APIV1AuctionBidAsyncParamsPriority string

const (
	APIV1AuctionBidAsyncParamsPriorityLow    APIV1AuctionBidAsyncParamsPriority = "low"
	APIV1AuctionBidAsyncParamsPriorityNormal APIV1AuctionBidAsyncParamsPriority = "normal"
	APIV1AuctionBidAsyncParamsPriorityHigh   APIV1AuctionBidAsyncParamsPriority = "high"
)
