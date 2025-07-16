// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"net/http"
	"net/url"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/apiquery"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/param"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// APIV1PublicService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1PublicService] method instead.
type APIV1PublicService struct {
	Options  []option.RequestOption
	Auctions APIV1PublicAuctionService
}

// NewAPIV1PublicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1PublicService(opts ...option.RequestOption) (r APIV1PublicService) {
	r = APIV1PublicService{}
	r.Options = opts
	r.Auctions = NewAPIV1PublicAuctionService(opts...)
	return
}

// Retrieve a comprehensive overview of the auction system including active
// auctions, upcoming auctions, recent results, and market statistics.
func (r *APIV1PublicService) GetOverview(ctx context.Context, query APIV1PublicGetOverviewParams, opts ...option.RequestOption) (res *Apiv1PublicGetOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/public/overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1PublicGetOverviewResponse struct {
	Metadata Apiv1PublicGetOverviewResponseMetadata `json:"metadata,required"`
	Overview Apiv1PublicGetOverviewResponseOverview `json:"overview,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Overview    respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseMetadata struct {
	CacheInfo     Apiv1PublicGetOverviewResponseMetadataCacheInfo     `json:"cacheInfo,required"`
	DataFreshness Apiv1PublicGetOverviewResponseMetadataDataFreshness `json:"dataFreshness,required"`
	QueryParams   map[string]any                                      `json:"queryParams,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheInfo     respjson.Field
		DataFreshness respjson.Field
		QueryParams   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseMetadataCacheInfo struct {
	Cached bool    `json:"cached,required"`
	MaxAge float64 `json:"maxAge,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cached      respjson.Field
		MaxAge      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseMetadataCacheInfo) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseMetadataCacheInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseMetadataDataFreshness struct {
	ActiveAuctions int64 `json:"activeAuctions,required"`
	// Last data update timestamp
	LastUpdate        float64 `json:"lastUpdate,required"`
	RecentlyCompleted int64   `json:"recentlyCompleted,required"`
	UpcomingAuctions  int64   `json:"upcomingAuctions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveAuctions    respjson.Field
		LastUpdate        respjson.Field
		RecentlyCompleted respjson.Field
		UpcomingAuctions  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseMetadataDataFreshness) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseMetadataDataFreshness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverview struct {
	CurrentActiveAuctions   []Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuction   `json:"currentActiveAuctions,required"`
	NextScheduledAuctions   []Apiv1PublicGetOverviewResponseOverviewNextScheduledAuction   `json:"nextScheduledAuctions,required"`
	RecentCompletedAuctions []Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuction `json:"recentCompletedAuctions,required"`
	SystemHealth            Apiv1PublicGetOverviewResponseOverviewSystemHealth             `json:"systemHealth,required"`
	MarketStatistics        Apiv1PublicGetOverviewResponseOverviewMarketStatistics         `json:"marketStatistics"`
	SlotTiming              Apiv1PublicGetOverviewResponseOverviewSlotTiming               `json:"slotTiming"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentActiveAuctions   respjson.Field
		NextScheduledAuctions   respjson.Field
		RecentCompletedAuctions respjson.Field
		SystemHealth            respjson.Field
		MarketStatistics        respjson.Field
		SlotTiming              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverview) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseOverview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                                                   `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuctionElasticSupply `json:"elasticSupply,required"`
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
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CurrentSupply    respjson.Field
		ElasticSupply    respjson.Field
		EndTime          respjson.Field
		MinBid           respjson.Field
		ScheduledStart   respjson.Field
		Slot             respjson.Field
		State            respjson.Field
		TotalBids        respjson.Field
		ActualStart      respjson.Field
		ClearingPrice    respjson.Field
		HighestBidAmount respjson.Field
		SettledAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuction) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuctionElasticSupply struct {
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
func (r Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuctionElasticSupply) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewCurrentActiveAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewNextScheduledAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                                                   `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicGetOverviewResponseOverviewNextScheduledAuctionElasticSupply `json:"elasticSupply,required"`
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
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CurrentSupply    respjson.Field
		ElasticSupply    respjson.Field
		EndTime          respjson.Field
		MinBid           respjson.Field
		ScheduledStart   respjson.Field
		Slot             respjson.Field
		State            respjson.Field
		TotalBids        respjson.Field
		ActualStart      respjson.Field
		ClearingPrice    respjson.Field
		HighestBidAmount respjson.Field
		SettledAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewNextScheduledAuction) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewNextScheduledAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewNextScheduledAuctionElasticSupply struct {
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
func (r Apiv1PublicGetOverviewResponseOverviewNextScheduledAuctionElasticSupply) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewNextScheduledAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                                                     `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuctionElasticSupply `json:"elasticSupply,required"`
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
	// Highest bid amount in wei (no bidder info)
	HighestBidAmount string `json:"highestBidAmount"`
	// Settlement timestamp
	SettledAt float64 `json:"settledAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CurrentSupply    respjson.Field
		ElasticSupply    respjson.Field
		EndTime          respjson.Field
		MinBid           respjson.Field
		ScheduledStart   respjson.Field
		Slot             respjson.Field
		State            respjson.Field
		TotalBids        respjson.Field
		ActualStart      respjson.Field
		ClearingPrice    respjson.Field
		HighestBidAmount respjson.Field
		SettledAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuction) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuctionElasticSupply struct {
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
func (r Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuctionElasticSupply) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicGetOverviewResponseOverviewRecentCompletedAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewSystemHealth struct {
	// Any of "healthy", "degraded", "unhealthy".
	AuctionSystemStatus string `json:"auctionSystemStatus,required"`
	// Any of "healthy", "degraded", "unhealthy".
	BiddingSystemStatus string `json:"biddingSystemStatus,required"`
	// Last health check timestamp
	LastHealthCheck float64 `json:"lastHealthCheck,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuctionSystemStatus respjson.Field
		BiddingSystemStatus respjson.Field
		LastHealthCheck     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewSystemHealth) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseOverviewSystemHealth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewMarketStatistics struct {
	AverageBidsPerAuction float64 `json:"averageBidsPerAuction,required"`
	// Average clearing price in wei
	AverageClearingPrice string `json:"averageClearingPrice,required"`
	TotalAuctions        int64  `json:"totalAuctions,required"`
	TotalBids            int64  `json:"totalBids,required"`
	// Total revenue in wei
	TotalRevenue string `json:"totalRevenue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageBidsPerAuction respjson.Field
		AverageClearingPrice  respjson.Field
		TotalAuctions         respjson.Field
		TotalBids             respjson.Field
		TotalRevenue          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewMarketStatistics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseOverviewMarketStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicGetOverviewResponseOverviewSlotTiming struct {
	// Average slot duration in milliseconds
	AverageSlotTime float64 `json:"averageSlotTime,required"`
	CurrentSlot     int64   `json:"currentSlot,required"`
	// Time until next slot in milliseconds
	NextSlotIn float64 `json:"nextSlotIn,required"`
	// Slot efficiency ratio
	SlotEfficiency float64 `json:"slotEfficiency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageSlotTime respjson.Field
		CurrentSlot     respjson.Field
		NextSlotIn      respjson.Field
		SlotEfficiency  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicGetOverviewResponseOverviewSlotTiming) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicGetOverviewResponseOverviewSlotTiming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1PublicGetOverviewParams struct {
	// Include market statistics
	IncludeMarketStats param.Opt[bool] `query:"include_market_stats,omitzero" json:"-"`
	// Include slot timing information
	IncludeSlotTiming param.Opt[bool] `query:"include_slot_timing,omitzero" json:"-"`
	// Number of completed auctions to include
	CompletedAuctions param.Opt[int64] `query:"completed_auctions,omitzero" json:"-"`
	// Number of next auctions to include
	NextAuctions param.Opt[int64] `query:"next_auctions,omitzero" json:"-"`
	// Time range for statistics in hours
	TimeRangeHours param.Opt[int64] `query:"time_range_hours,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1PublicGetOverviewParams]'s query parameters as
// `url.Values`.
func (r APIV1PublicGetOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
