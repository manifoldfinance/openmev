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

// APIV1AuctionService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AuctionService] method instead.
type APIV1AuctionService struct {
	Options []option.RequestOption
	Bids    APIV1AuctionBidService
}

// NewAPIV1AuctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1AuctionService(opts ...option.RequestOption) (r APIV1AuctionService) {
	r = APIV1AuctionService{}
	r.Options = opts
	r.Bids = NewAPIV1AuctionBidService(opts...)
	return
}

// Create a new auction with the specified parameters. Requires authentication with
// auction_manage permissions.
func (r *APIV1AuctionService) New(ctx context.Context, body APIV1AuctionNewParams, opts ...option.RequestOption) (res *Apiv1AuctionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auctions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific auction including bid summary and
// analytics. Requires authentication with read permissions.
func (r *APIV1AuctionService) Get(ctx context.Context, id string, query APIV1AuctionGetParams, opts ...option.RequestOption) (res *Apiv1AuctionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/auctions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of auctions with advanced filtering options. Requires
// authentication with read permissions.
func (r *APIV1AuctionService) List(ctx context.Context, query APIV1AuctionListParams, opts ...option.RequestOption) (res *Apiv1AuctionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auctions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1AuctionNewResponse struct {
	Auction Apiv1AuctionNewResponseAuction `json:"auction,required"`
	// Success message
	Message string `json:"message,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Auction     respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionNewResponseAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Number of bidding rounds
	BiddingRounds int64 `json:"biddingRounds,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// ID of user who created the auction
	CreatedBy string `json:"createdBy,required"`
	// Current allocated supply
	CurrentSupply int64                                       `json:"currentSupply,required"`
	ElasticSupply Apiv1AuctionNewResponseAuctionElasticSupply `json:"elasticSupply,required"`
	// End timestamp
	EndTime float64 `json:"endTime,required"`
	// Minimum bid amount in wei
	MinBid string `json:"minBid,required"`
	// Scheduled start timestamp
	ScheduledStart float64 `json:"scheduledStart,required"`
	// Ethereum slot number
	Slot int64 `json:"slot,required"`
	// Any of "scheduled", "active", "settling", "settled".
	State string `json:"state,required"`
	// Total number of bids
	TotalBids int64 `json:"totalBids,required"`
	// Actual start timestamp
	ActualStart float64 `json:"actualStart"`
	// Clearing price in wei
	ClearingPrice string `json:"clearingPrice"`
	// Gas used for settlement
	GasUsed string `json:"gasUsed"`
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Timestamp of last bid
	LastBidAt float64 `json:"lastBidAt"`
	// Additional auction metadata
	Metadata            map[string]any                                    `json:"metadata"`
	RevenueDistribution Apiv1AuctionNewResponseAuctionRevenueDistribution `json:"revenueDistribution"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// ID of user who settled the auction
	SettledBy string `json:"settledBy"`
	// Reason for settlement
	SettlementReason string `json:"settlementReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BiddingRounds       respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CurrentSupply       respjson.Field
		ElasticSupply       respjson.Field
		EndTime             respjson.Field
		MinBid              respjson.Field
		ScheduledStart      respjson.Field
		Slot                respjson.Field
		State               respjson.Field
		TotalBids           respjson.Field
		ActualStart         respjson.Field
		ClearingPrice       respjson.Field
		GasUsed             respjson.Field
		HighestBidAmount    respjson.Field
		LastBidAt           respjson.Field
		Metadata            respjson.Field
		RevenueDistribution respjson.Field
		SettledAt           respjson.Field
		SettledBy           respjson.Field
		SettlementReason    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionNewResponseAuction) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionNewResponseAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionNewResponseAuctionElasticSupply struct {
	// Base price in wei
	BasePrice string `json:"basePrice,required"`
	// Elasticity factor (0-1)
	Elasticity float64 `json:"elasticity,required"`
	// Maximum capacity
	MaxCapacity int64 `json:"maxCapacity,required"`
	// Minimum capacity
	MinCapacity int64 `json:"minCapacity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BasePrice   respjson.Field
		Elasticity  respjson.Field
		MaxCapacity respjson.Field
		MinCapacity respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionNewResponseAuctionElasticSupply) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionNewResponseAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionNewResponseAuctionRevenueDistribution struct {
	// Net revenue after fees
	NetRevenue string `json:"netRevenue,required"`
	// Platform fee amount
	PlatformFee string `json:"platformFee,required"`
	// Total revenue generated
	TotalRevenue string `json:"totalRevenue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetRevenue   respjson.Field
		PlatformFee  respjson.Field
		TotalRevenue respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionNewResponseAuctionRevenueDistribution) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionNewResponseAuctionRevenueDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponse struct {
	Auction Apiv1AuctionGetResponseAuction `json:"auction,required"`
	// Any of true.
	Success    bool                              `json:"success,required"`
	Timestamp  float64                           `json:"timestamp,required"`
	Analytics  Apiv1AuctionGetResponseAnalytics  `json:"analytics"`
	BidSummary Apiv1AuctionGetResponseBidSummary `json:"bidSummary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Auction     respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Analytics   respjson.Field
		BidSummary  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Number of bidding rounds
	BiddingRounds int64 `json:"biddingRounds,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// ID of user who created the auction
	CreatedBy string `json:"createdBy,required"`
	// Current allocated supply
	CurrentSupply int64                                       `json:"currentSupply,required"`
	ElasticSupply Apiv1AuctionGetResponseAuctionElasticSupply `json:"elasticSupply,required"`
	// End timestamp
	EndTime float64 `json:"endTime,required"`
	// Minimum bid amount in wei
	MinBid string `json:"minBid,required"`
	// Scheduled start timestamp
	ScheduledStart float64 `json:"scheduledStart,required"`
	// Ethereum slot number
	Slot int64 `json:"slot,required"`
	// Any of "scheduled", "active", "settling", "settled".
	State string `json:"state,required"`
	// Total number of bids
	TotalBids int64 `json:"totalBids,required"`
	// Actual start timestamp
	ActualStart float64 `json:"actualStart"`
	// Clearing price in wei
	ClearingPrice string `json:"clearingPrice"`
	// Gas used for settlement
	GasUsed string `json:"gasUsed"`
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Timestamp of last bid
	LastBidAt float64 `json:"lastBidAt"`
	// Additional auction metadata
	Metadata            map[string]any                                    `json:"metadata"`
	RevenueDistribution Apiv1AuctionGetResponseAuctionRevenueDistribution `json:"revenueDistribution"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// ID of user who settled the auction
	SettledBy string `json:"settledBy"`
	// Reason for settlement
	SettlementReason string `json:"settlementReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BiddingRounds       respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CurrentSupply       respjson.Field
		ElasticSupply       respjson.Field
		EndTime             respjson.Field
		MinBid              respjson.Field
		ScheduledStart      respjson.Field
		Slot                respjson.Field
		State               respjson.Field
		TotalBids           respjson.Field
		ActualStart         respjson.Field
		ClearingPrice       respjson.Field
		GasUsed             respjson.Field
		HighestBidAmount    respjson.Field
		LastBidAt           respjson.Field
		Metadata            respjson.Field
		RevenueDistribution respjson.Field
		SettledAt           respjson.Field
		SettledBy           respjson.Field
		SettlementReason    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseAuction) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseAuctionElasticSupply struct {
	// Base price in wei
	BasePrice string `json:"basePrice,required"`
	// Elasticity factor (0-1)
	Elasticity float64 `json:"elasticity,required"`
	// Maximum capacity
	MaxCapacity int64 `json:"maxCapacity,required"`
	// Minimum capacity
	MinCapacity int64 `json:"minCapacity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BasePrice   respjson.Field
		Elasticity  respjson.Field
		MaxCapacity respjson.Field
		MinCapacity respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseAuctionElasticSupply) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseAuctionRevenueDistribution struct {
	// Net revenue after fees
	NetRevenue string `json:"netRevenue,required"`
	// Platform fee amount
	PlatformFee string `json:"platformFee,required"`
	// Total revenue generated
	TotalRevenue string `json:"totalRevenue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetRevenue   respjson.Field
		PlatformFee  respjson.Field
		TotalRevenue respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseAuctionRevenueDistribution) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseAuctionRevenueDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseAnalytics struct {
	// Competition index
	CompetitionIndex float64 `json:"competitionIndex,required"`
	// Participation rate
	ParticipationRate float64 `json:"participationRate,required"`
	// Price discovery efficiency
	PriceDiscoveryEfficiency float64 `json:"priceDiscoveryEfficiency,required"`
	// Time to settlement in milliseconds
	TimeToSettlement float64 `json:"timeToSettlement,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompetitionIndex         respjson.Field
		ParticipationRate        respjson.Field
		PriceDiscoveryEfficiency respjson.Field
		TimeToSettlement         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseAnalytics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseAnalytics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseBidSummary struct {
	// Average bid amount in wei
	AverageBidAmount string                                       `json:"averageBidAmount,required"`
	TopBidders       []Apiv1AuctionGetResponseBidSummaryTopBidder `json:"topBidders,required"`
	TotalBids        int64                                        `json:"totalBids,required"`
	// Total bid value in wei
	TotalValue  string `json:"totalValue,required"`
	WinningBids int64  `json:"winningBids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageBidAmount respjson.Field
		TopBidders       respjson.Field
		TotalBids        respjson.Field
		TotalValue       respjson.Field
		WinningBids      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseBidSummary) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseBidSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionGetResponseBidSummaryTopBidder struct {
	// Bid amount in wei
	Amount   string `json:"amount,required"`
	BidID    string `json:"bidId,required"`
	Quantity int64  `json:"quantity,required"`
	// Bid timestamp
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		BidID       respjson.Field
		Quantity    respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionGetResponseBidSummaryTopBidder) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionGetResponseBidSummaryTopBidder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionListResponse struct {
	Count    float64                          `json:"count,required"`
	Data     []Apiv1AuctionListResponseData   `json:"data,required"`
	HasMore  bool                             `json:"hasMore,required"`
	Metadata Apiv1AuctionListResponseMetadata `json:"metadata,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	Cursor    string  `json:"cursor"`
	Total     float64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Data        respjson.Field
		HasMore     respjson.Field
		Metadata    respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Cursor      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionListResponseData struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Number of bidding rounds
	BiddingRounds int64 `json:"biddingRounds,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// ID of user who created the auction
	CreatedBy string `json:"createdBy,required"`
	// Current allocated supply
	CurrentSupply int64                                     `json:"currentSupply,required"`
	ElasticSupply Apiv1AuctionListResponseDataElasticSupply `json:"elasticSupply,required"`
	// End timestamp
	EndTime float64 `json:"endTime,required"`
	// Minimum bid amount in wei
	MinBid string `json:"minBid,required"`
	// Scheduled start timestamp
	ScheduledStart float64 `json:"scheduledStart,required"`
	// Ethereum slot number
	Slot int64 `json:"slot,required"`
	// Any of "scheduled", "active", "settling", "settled".
	State string `json:"state,required"`
	// Total number of bids
	TotalBids int64 `json:"totalBids,required"`
	// Actual start timestamp
	ActualStart float64 `json:"actualStart"`
	// Clearing price in wei
	ClearingPrice string `json:"clearingPrice"`
	// Gas used for settlement
	GasUsed string `json:"gasUsed"`
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Timestamp of last bid
	LastBidAt float64 `json:"lastBidAt"`
	// Additional auction metadata
	Metadata            map[string]any                                  `json:"metadata"`
	RevenueDistribution Apiv1AuctionListResponseDataRevenueDistribution `json:"revenueDistribution"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// ID of user who settled the auction
	SettledBy string `json:"settledBy"`
	// Reason for settlement
	SettlementReason string `json:"settlementReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BiddingRounds       respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CurrentSupply       respjson.Field
		ElasticSupply       respjson.Field
		EndTime             respjson.Field
		MinBid              respjson.Field
		ScheduledStart      respjson.Field
		Slot                respjson.Field
		State               respjson.Field
		TotalBids           respjson.Field
		ActualStart         respjson.Field
		ClearingPrice       respjson.Field
		GasUsed             respjson.Field
		HighestBidAmount    respjson.Field
		LastBidAt           respjson.Field
		Metadata            respjson.Field
		RevenueDistribution respjson.Field
		SettledAt           respjson.Field
		SettledBy           respjson.Field
		SettlementReason    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionListResponseDataElasticSupply struct {
	// Base price in wei
	BasePrice string `json:"basePrice,required"`
	// Elasticity factor (0-1)
	Elasticity float64 `json:"elasticity,required"`
	// Maximum capacity
	MaxCapacity int64 `json:"maxCapacity,required"`
	// Minimum capacity
	MinCapacity int64 `json:"minCapacity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BasePrice   respjson.Field
		Elasticity  respjson.Field
		MaxCapacity respjson.Field
		MinCapacity respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionListResponseDataElasticSupply) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionListResponseDataElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionListResponseDataRevenueDistribution struct {
	// Net revenue after fees
	NetRevenue string `json:"netRevenue,required"`
	// Platform fee amount
	PlatformFee string `json:"platformFee,required"`
	// Total revenue generated
	TotalRevenue string `json:"totalRevenue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetRevenue   respjson.Field
		PlatformFee  respjson.Field
		TotalRevenue respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionListResponseDataRevenueDistribution) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionListResponseDataRevenueDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuctionListResponseMetadata struct {
	PageSize           float64 `json:"pageSize,required"`
	CurrentPage        float64 `json:"currentPage"`
	FirstItemTimestamp float64 `json:"firstItemTimestamp"`
	LastItemTimestamp  float64 `json:"lastItemTimestamp"`
	TotalPages         float64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageSize           respjson.Field
		CurrentPage        respjson.Field
		FirstItemTimestamp respjson.Field
		LastItemTimestamp  respjson.Field
		TotalPages         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuctionListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuctionListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuctionNewParams struct {
	// Auction duration in milliseconds
	Duration      int64                              `json:"duration,required"`
	ElasticSupply APIV1AuctionNewParamsElasticSupply `json:"elasticSupply,omitzero,required"`
	// Minimum bid amount in wei
	MinBid string `json:"minBid,required"`
	// Scheduled start timestamp
	ScheduledStart float64 `json:"scheduledStart,required"`
	// Ethereum slot number for the auction
	Slot int64 `json:"slot,required"`
	// Additional auction metadata
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r APIV1AuctionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuctionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BasePrice, Elasticity, MaxCapacity, MinCapacity are required.
type APIV1AuctionNewParamsElasticSupply struct {
	// Base price in wei
	BasePrice string `json:"basePrice,required"`
	// Elasticity factor (0-1)
	Elasticity float64 `json:"elasticity,required"`
	// Maximum capacity
	MaxCapacity int64 `json:"maxCapacity,required"`
	// Minimum capacity
	MinCapacity int64 `json:"minCapacity,required"`
	paramObj
}

func (r APIV1AuctionNewParamsElasticSupply) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuctionNewParamsElasticSupply
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuctionNewParamsElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuctionGetParams struct {
	// Include auction analytics
	IncludeAnalytics param.Opt[bool] `query:"include_analytics,omitzero" json:"-"`
	// Include bid summary
	IncludeBids param.Opt[bool] `query:"include_bids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AuctionGetParams]'s query parameters as `url.Values`.
func (r APIV1AuctionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AuctionListParams struct {
	// Start date as Unix timestamp
	FromDate param.Opt[float64] `query:"from_date,omitzero" json:"-"`
	// Include auction metadata
	IncludeMetadata param.Opt[bool] `query:"include_metadata,omitzero" json:"-"`
	// Include revenue information
	IncludeRevenue param.Opt[bool] `query:"include_revenue,omitzero" json:"-"`
	// Starting slot number
	SlotFrom param.Opt[int64] `query:"slot_from,omitzero" json:"-"`
	// Ending slot number
	SlotTo param.Opt[int64] `query:"slot_to,omitzero" json:"-"`
	// End date as Unix timestamp
	ToDate param.Opt[float64] `query:"to_date,omitzero" json:"-"`
	// Filter by creator user ID
	CreatedBy param.Opt[string] `query:"created_by,omitzero" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Any of "createdAt", "slot", "state", "endTime", "settledAt", "totalBids".
	SortBy APIV1AuctionListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	SortOrder APIV1AuctionListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Filter by auction state
	//
	// Any of "scheduled", "active", "settling", "settled".
	State APIV1AuctionListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AuctionListParams]'s query parameters as `url.Values`.
func (r APIV1AuctionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AuctionListParamsSortBy string

const (
	APIV1AuctionListParamsSortByCreatedAt APIV1AuctionListParamsSortBy = "createdAt"
	APIV1AuctionListParamsSortBySlot      APIV1AuctionListParamsSortBy = "slot"
	APIV1AuctionListParamsSortByState     APIV1AuctionListParamsSortBy = "state"
	APIV1AuctionListParamsSortByEndTime   APIV1AuctionListParamsSortBy = "endTime"
	APIV1AuctionListParamsSortBySettledAt APIV1AuctionListParamsSortBy = "settledAt"
	APIV1AuctionListParamsSortByTotalBids APIV1AuctionListParamsSortBy = "totalBids"
)

// Sort order
type APIV1AuctionListParamsSortOrder string

const (
	APIV1AuctionListParamsSortOrderAsc  APIV1AuctionListParamsSortOrder = "asc"
	APIV1AuctionListParamsSortOrderDesc APIV1AuctionListParamsSortOrder = "desc"
)

// Filter by auction state
type APIV1AuctionListParamsState string

const (
	APIV1AuctionListParamsStateScheduled APIV1AuctionListParamsState = "scheduled"
	APIV1AuctionListParamsStateActive    APIV1AuctionListParamsState = "active"
	APIV1AuctionListParamsStateSettling  APIV1AuctionListParamsState = "settling"
	APIV1AuctionListParamsStateSettled   APIV1AuctionListParamsState = "settled"
)
