// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/apiquery"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/param"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// APIV1PublicAuctionService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1PublicAuctionService] method instead.
type APIV1PublicAuctionService struct {
	Options []option.RequestOption
}

// NewAPIV1PublicAuctionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV1PublicAuctionService(opts ...option.RequestOption) (r APIV1PublicAuctionService) {
	r = APIV1PublicAuctionService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific auction with optional analytics
// and bid statistics.
func (r *APIV1PublicAuctionService) Get(ctx context.Context, id string, query APIV1PublicAuctionGetParams, opts ...option.RequestOption) (res *Apiv1PublicAuctionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/public/auctions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of historical auctions with optional filtering by
// state, slot range, date range, and statistical data.
func (r *APIV1PublicAuctionService) GetHistory(ctx context.Context, query APIV1PublicAuctionGetHistoryParams, opts ...option.RequestOption) (res *Apiv1PublicAuctionGetHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/public/auctions/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the latest auction results with optional analytics, trend analysis, and
// market metrics.
func (r *APIV1PublicAuctionService) GetLatestResults(ctx context.Context, query APIV1PublicAuctionGetLatestResultsParams, opts ...option.RequestOption) (res *Apiv1PublicAuctionGetLatestResultsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/public/auctions/latest-results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1PublicAuctionGetResponse struct {
	Auction Apiv1PublicAuctionGetResponseAuction `json:"auction,required"`
	// Any of true.
	Success   bool                                   `json:"success,required"`
	Timestamp float64                                `json:"timestamp,required"`
	Analytics Apiv1PublicAuctionGetResponseAnalytics `json:"analytics"`
	BidStats  Apiv1PublicAuctionGetResponseBidStats  `json:"bidStats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Auction     respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Analytics   respjson.Field
		BidStats    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseAuction struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                             `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicAuctionGetResponseAuctionElasticSupply `json:"elasticSupply,required"`
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
func (r Apiv1PublicAuctionGetResponseAuction) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseAuction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseAuctionElasticSupply struct {
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
func (r Apiv1PublicAuctionGetResponseAuctionElasticSupply) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseAuctionElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseAnalytics struct {
	// Total auction duration (ms)
	Duration float64 `json:"duration,required"`
	// Whether auction is past end time
	IsOvertime bool                                            `json:"isOvertime,required"`
	PhaseInfo  Apiv1PublicAuctionGetResponseAnalyticsPhaseInfo `json:"phaseInfo,required"`
	// Progress percentage
	ProgressPercentage float64 `json:"progressPercentage,required"`
	// Time elapsed since start (ms)
	TimeElapsed float64 `json:"timeElapsed,required"`
	// Time remaining until end (ms)
	TimeRemaining float64 `json:"timeRemaining,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration           respjson.Field
		IsOvertime         respjson.Field
		PhaseInfo          respjson.Field
		ProgressPercentage respjson.Field
		TimeElapsed        respjson.Field
		TimeRemaining      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponseAnalytics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseAnalytics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseAnalyticsPhaseInfo struct {
	// Any of "scheduled", "active", "settling", "settled".
	Current    string `json:"current,required"`
	IsActive   bool   `json:"isActive,required"`
	IsComplete bool   `json:"isComplete,required"`
	IsPreStart bool   `json:"isPreStart,required"`
	IsSettling bool   `json:"isSettling,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current     respjson.Field
		IsActive    respjson.Field
		IsComplete  respjson.Field
		IsPreStart  respjson.Field
		IsSettling  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponseAnalyticsPhaseInfo) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseAnalyticsPhaseInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseBidStats struct {
	// Allocation ratio as percentage
	AllocationRatio float64 `json:"allocationRatio,required"`
	// Competition level (0-10)
	CompetitionLevel  float64                                             `json:"competitionLevel,required"`
	CurrentAllocation int64                                               `json:"currentAllocation,required"`
	MaxCapacity       int64                                               `json:"maxCapacity,required"`
	PriceDiscovery    Apiv1PublicAuctionGetResponseBidStatsPriceDiscovery `json:"priceDiscovery,required"`
	TotalBids         int64                                               `json:"totalBids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllocationRatio   respjson.Field
		CompetitionLevel  respjson.Field
		CurrentAllocation respjson.Field
		MaxCapacity       respjson.Field
		PriceDiscovery    respjson.Field
		TotalBids         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponseBidStats) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseBidStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseBidStatsPriceDiscovery struct {
	ClearingPrice string                                                        `json:"clearingPrice,required"`
	HasClearing   bool                                                          `json:"hasClearing,required"`
	MinBid        string                                                        `json:"minBid,required"`
	PriceRange    Apiv1PublicAuctionGetResponseBidStatsPriceDiscoveryPriceRange `json:"priceRange,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClearingPrice respjson.Field
		HasClearing   respjson.Field
		MinBid        respjson.Field
		PriceRange    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponseBidStatsPriceDiscovery) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetResponseBidStatsPriceDiscovery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetResponseBidStatsPriceDiscoveryPriceRange struct {
	Multiplier float64 `json:"multiplier,required"`
	Spread     string  `json:"spread,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Multiplier  respjson.Field
		Spread      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetResponseBidStatsPriceDiscoveryPriceRange) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicAuctionGetResponseBidStatsPriceDiscoveryPriceRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetHistoryResponse struct {
	Count    float64                                      `json:"count,required"`
	Data     []Apiv1PublicAuctionGetHistoryResponseData   `json:"data,required"`
	HasMore  bool                                         `json:"hasMore,required"`
	Metadata Apiv1PublicAuctionGetHistoryResponseMetadata `json:"metadata,required"`
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
func (r Apiv1PublicAuctionGetHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetHistoryResponseData struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                                 `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicAuctionGetHistoryResponseDataElasticSupply `json:"elasticSupply,required"`
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
func (r Apiv1PublicAuctionGetHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetHistoryResponseDataElasticSupply struct {
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
func (r Apiv1PublicAuctionGetHistoryResponseDataElasticSupply) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetHistoryResponseDataElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetHistoryResponseMetadata struct {
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
func (r Apiv1PublicAuctionGetHistoryResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetHistoryResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponse struct {
	Count    int64                                              `json:"count,required"`
	Data     []Apiv1PublicAuctionGetLatestResultsResponseData   `json:"data,required"`
	Metadata Apiv1PublicAuctionGetLatestResultsResponseMetadata `json:"metadata,required"`
	// Any of true.
	Success         bool                                                      `json:"success,required"`
	Timestamp       float64                                                   `json:"timestamp,required"`
	Total           int64                                                     `json:"total,required"`
	MarketMetrics   Apiv1PublicAuctionGetLatestResultsResponseMarketMetrics   `json:"marketMetrics"`
	PerformanceData Apiv1PublicAuctionGetLatestResultsResponsePerformanceData `json:"performanceData"`
	Statistics      Apiv1PublicAuctionGetLatestResultsResponseStatistics      `json:"statistics"`
	TrendAnalysis   Apiv1PublicAuctionGetLatestResultsResponseTrendAnalysis   `json:"trendAnalysis"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count           respjson.Field
		Data            respjson.Field
		Metadata        respjson.Field
		Success         respjson.Field
		Timestamp       respjson.Field
		Total           respjson.Field
		MarketMetrics   respjson.Field
		PerformanceData respjson.Field
		Statistics      respjson.Field
		TrendAnalysis   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseData struct {
	// Unique auction identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Current allocated supply
	CurrentSupply int64                                                       `json:"currentSupply,required"`
	ElasticSupply Apiv1PublicAuctionGetLatestResultsResponseDataElasticSupply `json:"elasticSupply,required"`
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
func (r Apiv1PublicAuctionGetLatestResultsResponseData) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseDataElasticSupply struct {
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
func (r Apiv1PublicAuctionGetLatestResultsResponseDataElasticSupply) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicAuctionGetLatestResultsResponseDataElasticSupply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseMetadata struct {
	CacheInfo      Apiv1PublicAuctionGetLatestResultsResponseMetadataCacheInfo `json:"cacheInfo,required"`
	QueryParams    map[string]any                                              `json:"queryParams,required"`
	DataSource     string                                                      `json:"dataSource"`
	ProcessingTime float64                                                     `json:"processingTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheInfo      respjson.Field
		QueryParams    respjson.Field
		DataSource     respjson.Field
		ProcessingTime respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseMetadataCacheInfo struct {
	Key      string  `json:"key,required"`
	Strategy string  `json:"strategy,required"`
	Ttl      float64 `json:"ttl,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Strategy    respjson.Field
		Ttl         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponseMetadataCacheInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicAuctionGetLatestResultsResponseMetadataCacheInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseMarketMetrics struct {
	// Average clearing price in wei
	AveragePrice string `json:"averagePrice,required"`
	// Demand intensity metric
	DemandIntensity float64 `json:"demandIntensity,required"`
	// Price volatility index
	PriceVolatility float64 `json:"priceVolatility,required"`
	// Average supply utilization
	SupplyUtilization float64 `json:"supplyUtilization,required"`
	// Total trading volume in wei
	TotalVolume string `json:"totalVolume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AveragePrice      respjson.Field
		DemandIntensity   respjson.Field
		PriceVolatility   respjson.Field
		SupplyUtilization respjson.Field
		TotalVolume       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponseMarketMetrics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponseMarketMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponsePerformanceData struct {
	// Cache hit rate
	CacheHitRate float64 `json:"cacheHitRate,required"`
	// Data source type
	//
	// Any of "live", "cached".
	DataSource string `json:"dataSource,required"`
	// Last data update timestamp
	LastUpdate float64 `json:"lastUpdate,required"`
	// Processing time in milliseconds
	ProcessingTime float64 `json:"processingTime,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHitRate   respjson.Field
		DataSource     respjson.Field
		LastUpdate     respjson.Field
		ProcessingTime respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponsePerformanceData) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1PublicAuctionGetLatestResultsResponsePerformanceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseStatistics struct {
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
func (r Apiv1PublicAuctionGetLatestResultsResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1PublicAuctionGetLatestResultsResponseTrendAnalysis struct {
	// Demand trend factor
	DemandTrend float64 `json:"demandTrend,required"`
	// Any of "up", "down", "stable".
	PriceMovement string `json:"priceMovement,required"`
	// Supply utilization rate
	SupplyUtilization float64 `json:"supplyUtilization,required"`
	// Price volatility (0-1)
	Volatility float64 `json:"volatility,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DemandTrend       respjson.Field
		PriceMovement     respjson.Field
		SupplyUtilization respjson.Field
		Volatility        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1PublicAuctionGetLatestResultsResponseTrendAnalysis) RawJSON() string { return r.JSON.raw }
func (r *Apiv1PublicAuctionGetLatestResultsResponseTrendAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1PublicAuctionGetParams struct {
	// Include analytics data
	IncludeAnalytics param.Opt[bool] `query:"include_analytics,omitzero" json:"-"`
	// Include bid statistics
	IncludeBidStats param.Opt[bool] `query:"include_bid_stats,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1PublicAuctionGetParams]'s query parameters as
// `url.Values`.
func (r APIV1PublicAuctionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1PublicAuctionGetHistoryParams struct {
	// Start date as Unix timestamp
	FromDate param.Opt[float64] `query:"from_date,omitzero" json:"-"`
	// Include statistical data
	IncludeStats param.Opt[bool] `query:"include_stats,omitzero" json:"-"`
	// Starting slot number
	SlotFrom param.Opt[int64] `query:"slot_from,omitzero" json:"-"`
	// Ending slot number
	SlotTo param.Opt[int64] `query:"slot_to,omitzero" json:"-"`
	// End date as Unix timestamp
	ToDate param.Opt[float64] `query:"to_date,omitzero" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Any of "createdAt", "slot", "state", "endTime", "settledAt".
	SortBy APIV1PublicAuctionGetHistoryParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	SortOrder APIV1PublicAuctionGetHistoryParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Filter by auction state
	//
	// Any of "scheduled", "active", "settling", "settled".
	State APIV1PublicAuctionGetHistoryParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1PublicAuctionGetHistoryParams]'s query parameters as
// `url.Values`.
func (r APIV1PublicAuctionGetHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1PublicAuctionGetHistoryParamsSortBy string

const (
	APIV1PublicAuctionGetHistoryParamsSortByCreatedAt APIV1PublicAuctionGetHistoryParamsSortBy = "createdAt"
	APIV1PublicAuctionGetHistoryParamsSortBySlot      APIV1PublicAuctionGetHistoryParamsSortBy = "slot"
	APIV1PublicAuctionGetHistoryParamsSortByState     APIV1PublicAuctionGetHistoryParamsSortBy = "state"
	APIV1PublicAuctionGetHistoryParamsSortByEndTime   APIV1PublicAuctionGetHistoryParamsSortBy = "endTime"
	APIV1PublicAuctionGetHistoryParamsSortBySettledAt APIV1PublicAuctionGetHistoryParamsSortBy = "settledAt"
)

// Sort order
type APIV1PublicAuctionGetHistoryParamsSortOrder string

const (
	APIV1PublicAuctionGetHistoryParamsSortOrderAsc  APIV1PublicAuctionGetHistoryParamsSortOrder = "asc"
	APIV1PublicAuctionGetHistoryParamsSortOrderDesc APIV1PublicAuctionGetHistoryParamsSortOrder = "desc"
)

// Filter by auction state
type APIV1PublicAuctionGetHistoryParamsState string

const (
	APIV1PublicAuctionGetHistoryParamsStateScheduled APIV1PublicAuctionGetHistoryParamsState = "scheduled"
	APIV1PublicAuctionGetHistoryParamsStateActive    APIV1PublicAuctionGetHistoryParamsState = "active"
	APIV1PublicAuctionGetHistoryParamsStateSettling  APIV1PublicAuctionGetHistoryParamsState = "settling"
	APIV1PublicAuctionGetHistoryParamsStateSettled   APIV1PublicAuctionGetHistoryParamsState = "settled"
)

type APIV1PublicAuctionGetLatestResultsParams struct {
	// Include market metrics
	IncludeMarketMetrics param.Opt[bool] `query:"include_market_metrics,omitzero" json:"-"`
	// Include advanced metrics
	IncludeMetrics param.Opt[bool] `query:"include_metrics,omitzero" json:"-"`
	// Include performance data
	IncludePerformanceData param.Opt[bool] `query:"include_performance_data,omitzero" json:"-"`
	// Include trend analysis
	IncludeTrendAnalysis param.Opt[bool] `query:"include_trend_analysis,omitzero" json:"-"`
	// Starting slot number
	SlotFrom param.Opt[int64] `query:"slot_from,omitzero" json:"-"`
	// Ending slot number
	SlotTo param.Opt[int64] `query:"slot_to,omitzero" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Any of "settledAt", "slot", "clearingPrice", "totalBids".
	SortBy APIV1PublicAuctionGetLatestResultsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	SortOrder APIV1PublicAuctionGetLatestResultsParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1PublicAuctionGetLatestResultsParams]'s query
// parameters as `url.Values`.
func (r APIV1PublicAuctionGetLatestResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1PublicAuctionGetLatestResultsParamsSortBy string

const (
	APIV1PublicAuctionGetLatestResultsParamsSortBySettledAt     APIV1PublicAuctionGetLatestResultsParamsSortBy = "settledAt"
	APIV1PublicAuctionGetLatestResultsParamsSortBySlot          APIV1PublicAuctionGetLatestResultsParamsSortBy = "slot"
	APIV1PublicAuctionGetLatestResultsParamsSortByClearingPrice APIV1PublicAuctionGetLatestResultsParamsSortBy = "clearingPrice"
	APIV1PublicAuctionGetLatestResultsParamsSortByTotalBids     APIV1PublicAuctionGetLatestResultsParamsSortBy = "totalBids"
)

// Sort order
type APIV1PublicAuctionGetLatestResultsParamsSortOrder string

const (
	APIV1PublicAuctionGetLatestResultsParamsSortOrderAsc  APIV1PublicAuctionGetLatestResultsParamsSortOrder = "asc"
	APIV1PublicAuctionGetLatestResultsParamsSortOrderDesc APIV1PublicAuctionGetLatestResultsParamsSortOrder = "desc"
)
