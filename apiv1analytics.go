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

// APIV1AnalyticsService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AnalyticsService] method instead.
type APIV1AnalyticsService struct {
	Options []option.RequestOption
	Users   APIV1AnalyticsUserService
}

// NewAPIV1AnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1AnalyticsService(opts ...option.RequestOption) (r APIV1AnalyticsService) {
	r = APIV1AnalyticsService{}
	r.Options = opts
	r.Users = NewAPIV1AnalyticsUserService(opts...)
	return
}

// Returns detailed metrics for a specific organization
func (r *APIV1AnalyticsService) Get(ctx context.Context, id string, query APIV1AnalyticsGetParams, opts ...option.RequestOption) (res *Apiv1AnalyticsGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/analytics/organizations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns top performing organizations with public metrics
func (r *APIV1AnalyticsService) GetLeaderboard(ctx context.Context, query APIV1AnalyticsGetLeaderboardParams, opts ...option.RequestOption) (res *Apiv1AnalyticsGetLeaderboardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/analytics/leaderboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns comprehensive system metrics (admin only)
func (r *APIV1AnalyticsService) GetSystem(ctx context.Context, query APIV1AnalyticsGetSystemParams, opts ...option.RequestOption) (res *Apiv1AnalyticsGetSystemResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/analytics/system"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1AnalyticsGetResponse struct {
	Properties Apiv1AnalyticsGetResponseProperties `json:"properties,required"`
	Type       string                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponseProperties struct {
	Metrics      Apiv1AnalyticsGetResponsePropertiesMetrics      `json:"metrics,required"`
	Organization Apiv1AnalyticsGetResponsePropertiesOrganization `json:"organization,required"`
	Trends       Apiv1AnalyticsGetResponsePropertiesTrends       `json:"trends,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics      respjson.Field
		Organization respjson.Field
		Trends       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetrics struct {
	Properties Apiv1AnalyticsGetResponsePropertiesMetricsProperties `json:"properties,required"`
	Type       string                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetrics) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponsePropertiesMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsProperties struct {
	ActiveDays      Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesActiveDays      `json:"activeDays,required"`
	APICalls        Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAPICalls        `json:"apiCalls,required"`
	AvgResponseTime Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAvgResponseTime `json:"avgResponseTime,required"`
	ErrorRate       Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesErrorRate       `json:"errorRate,required"`
	MemberCount     Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesMemberCount     `json:"memberCount,required"`
	SuccessfulBids  Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesSuccessfulBids  `json:"successfulBids,required"`
	TotalBids       Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalBids       `json:"totalBids,required"`
	TotalVolume     Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalVolume     `json:"totalVolume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveDays      respjson.Field
		APICalls        respjson.Field
		AvgResponseTime respjson.Field
		ErrorRate       respjson.Field
		MemberCount     respjson.Field
		SuccessfulBids  respjson.Field
		TotalBids       respjson.Field
		TotalVolume     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesActiveDays struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesActiveDays) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesActiveDays) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAPICalls struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAPICalls) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAPICalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAvgResponseTime struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAvgResponseTime) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesAvgResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesErrorRate struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesErrorRate) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesErrorRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesMemberCount struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesMemberCount) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesMemberCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesSuccessfulBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesSuccessfulBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesSuccessfulBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalVolume struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalVolume) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesMetricsPropertiesTotalVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesOrganization struct {
	Properties Apiv1AnalyticsGetResponsePropertiesOrganizationProperties `json:"properties,required"`
	Type       string                                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesOrganization) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponsePropertiesOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesOrganizationProperties struct {
	ID   Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesID   `json:"id,required"`
	Name Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesName `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesOrganizationProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesOrganizationProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesID struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesID) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesName struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesName) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesOrganizationPropertiesName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrends struct {
	Properties Apiv1AnalyticsGetResponsePropertiesTrendsProperties `json:"properties,required"`
	Type       string                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrends) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponsePropertiesTrends) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsProperties struct {
	DailyMetrics Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetrics `json:"dailyMetrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyMetrics respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetrics struct {
	Items Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItems `json:"items,required"`
	Type  string                                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItems struct {
	Properties Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsProperties `json:"properties,required"`
	Type       string                                                                         `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItems) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsProperties struct {
	APICalls Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesAPICalls `json:"apiCalls,required"`
	BidCount Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesBidCount `json:"bidCount,required"`
	Date     Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesDate     `json:"date,required"`
	Volume   Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesVolume   `json:"volume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APICalls    respjson.Field
		BidCount    respjson.Field
		Date        respjson.Field
		Volume      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesAPICalls struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesAPICalls) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesAPICalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesBidCount struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesBidCount) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesBidCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesDate struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesDate) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesVolume struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesVolume) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetResponsePropertiesTrendsPropertiesDailyMetricsItemsPropertiesVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponse struct {
	Properties Apiv1AnalyticsGetLeaderboardResponseProperties `json:"properties,required"`
	Type       string                                         `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetLeaderboardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponseProperties struct {
	Leaderboard  Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboard  `json:"leaderboard,required"`
	Metric       Apiv1AnalyticsGetLeaderboardResponsePropertiesMetric       `json:"metric,required"`
	Period       Apiv1AnalyticsGetLeaderboardResponsePropertiesPeriod       `json:"period,required"`
	TotalEntries Apiv1AnalyticsGetLeaderboardResponsePropertiesTotalEntries `json:"totalEntries,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Leaderboard  respjson.Field
		Metric       respjson.Field
		Period       respjson.Field
		TotalEntries respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetLeaderboardResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboard struct {
	Items Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItems `json:"items,required"`
	Type  string                                                         `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboard) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItems struct {
	Properties Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsProperties `json:"properties,required"`
	Type       string                                                                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItems) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsProperties struct {
	Metrics      Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetrics      `json:"metrics,required"`
	Organization Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganization `json:"organization,required"`
	Rank         Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesRank         `json:"rank,required"`
	Score        Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesScore        `json:"score,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics      respjson.Field
		Organization respjson.Field
		Rank         respjson.Field
		Score        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetrics struct {
	Properties Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsProperties `json:"properties,required"`
	Type       string                                                                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsProperties struct {
	AvgBidTime  Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesAvgBidTime  `json:"avgBidTime,required"`
	BidCount    Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesBidCount    `json:"bidCount,required"`
	SuccessRate Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesSuccessRate `json:"successRate,required"`
	Volume      Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesVolume      `json:"volume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgBidTime  respjson.Field
		BidCount    respjson.Field
		SuccessRate respjson.Field
		Volume      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesAvgBidTime struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesAvgBidTime) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesAvgBidTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesBidCount struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesBidCount) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesBidCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesSuccessRate struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesSuccessRate) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesSuccessRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesVolume struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesVolume) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesMetricsPropertiesVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganization struct {
	Properties Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationProperties `json:"properties,required"`
	Type       string                                                                                         `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganization) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationProperties struct {
	ID   Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesID   `json:"id,required"`
	Name Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesName `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesID struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesID) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesName struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesName) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesOrganizationPropertiesName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesRank struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesRank) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesRank) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesScore struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesScore) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesLeaderboardItemsPropertiesScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesMetric struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesMetric) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesPeriod struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesPeriod) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetLeaderboardResponsePropertiesTotalEntries struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetLeaderboardResponsePropertiesTotalEntries) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetLeaderboardResponsePropertiesTotalEntries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponse struct {
	Properties Apiv1AnalyticsGetSystemResponseProperties `json:"properties,required"`
	Type       string                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetSystemResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponseProperties struct {
	Activity Apiv1AnalyticsGetSystemResponsePropertiesActivity `json:"activity,required"`
	Growth   Apiv1AnalyticsGetSystemResponsePropertiesGrowth   `json:"growth,required"`
	Overview Apiv1AnalyticsGetSystemResponsePropertiesOverview `json:"overview,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activity    respjson.Field
		Growth      respjson.Field
		Overview    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetSystemResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesActivity struct {
	Properties Apiv1AnalyticsGetSystemResponsePropertiesActivityProperties `json:"properties,required"`
	Type       string                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesActivity) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetSystemResponsePropertiesActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesActivityProperties struct {
	ActiveOrganizations Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesActiveOrganizations `json:"activeOrganizations,required"`
	AvgBidsPerOrg       Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesAvgBidsPerOrg       `json:"avgBidsPerOrg,required"`
	DailyAPICalls       Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesDailyAPICalls       `json:"dailyApiCalls,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveOrganizations respjson.Field
		AvgBidsPerOrg       respjson.Field
		DailyAPICalls       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesActivityProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesActivityProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesActiveOrganizations struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesActiveOrganizations) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesActiveOrganizations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesAvgBidsPerOrg struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesAvgBidsPerOrg) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesAvgBidsPerOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesDailyAPICalls struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesDailyAPICalls) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesActivityPropertiesDailyAPICalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesGrowth struct {
	Properties Apiv1AnalyticsGetSystemResponsePropertiesGrowthProperties `json:"properties,required"`
	Type       string                                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesGrowth) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetSystemResponsePropertiesGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesGrowthProperties struct {
	GrowthRate       Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesGrowthRate       `json:"growthRate,required"`
	NewOrganizations Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewOrganizations `json:"newOrganizations,required"`
	NewUsers         Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewUsers         `json:"newUsers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GrowthRate       respjson.Field
		NewOrganizations respjson.Field
		NewUsers         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesGrowthProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesGrowthProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesGrowthRate struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesGrowthRate) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesGrowthRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewOrganizations struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewOrganizations) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewOrganizations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewUsers struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewUsers) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesGrowthPropertiesNewUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverview struct {
	Properties Apiv1AnalyticsGetSystemResponsePropertiesOverviewProperties `json:"properties,required"`
	Type       string                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverview) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewProperties struct {
	AvgOrganizationSize Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesAvgOrganizationSize `json:"avgOrganizationSize,required"`
	TotalBids           Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalBids           `json:"totalBids,required"`
	TotalOrganizations  Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalOrganizations  `json:"totalOrganizations,required"`
	TotalUsers          Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalUsers          `json:"totalUsers,required"`
	TotalVolume         Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalVolume         `json:"totalVolume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgOrganizationSize respjson.Field
		TotalBids           respjson.Field
		TotalOrganizations  respjson.Field
		TotalUsers          respjson.Field
		TotalVolume         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesAvgOrganizationSize struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesAvgOrganizationSize) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesAvgOrganizationSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalOrganizations struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalOrganizations) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalOrganizations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalUsers struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalUsers) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalVolume struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalVolume) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsGetSystemResponsePropertiesOverviewPropertiesTotalVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AnalyticsGetParams struct {
	// Any of "7d", "30d", "90d", "1y".
	Period APIV1AnalyticsGetParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AnalyticsGetParams]'s query parameters as
// `url.Values`.
func (r APIV1AnalyticsGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AnalyticsGetParamsPeriod string

const (
	APIV1AnalyticsGetParamsPeriod7d  APIV1AnalyticsGetParamsPeriod = "7d"
	APIV1AnalyticsGetParamsPeriod30d APIV1AnalyticsGetParamsPeriod = "30d"
	APIV1AnalyticsGetParamsPeriod90d APIV1AnalyticsGetParamsPeriod = "90d"
	APIV1AnalyticsGetParamsPeriod1y  APIV1AnalyticsGetParamsPeriod = "1y"
)

type APIV1AnalyticsGetLeaderboardParams struct {
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Any of "volume", "bids", "success-rate", "activity".
	Metric APIV1AnalyticsGetLeaderboardParamsMetric `query:"metric,omitzero" json:"-"`
	// Any of "daily", "weekly", "monthly", "all-time".
	Period APIV1AnalyticsGetLeaderboardParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AnalyticsGetLeaderboardParams]'s query parameters as
// `url.Values`.
func (r APIV1AnalyticsGetLeaderboardParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AnalyticsGetLeaderboardParamsMetric string

const (
	APIV1AnalyticsGetLeaderboardParamsMetricVolume      APIV1AnalyticsGetLeaderboardParamsMetric = "volume"
	APIV1AnalyticsGetLeaderboardParamsMetricBids        APIV1AnalyticsGetLeaderboardParamsMetric = "bids"
	APIV1AnalyticsGetLeaderboardParamsMetricSuccessRate APIV1AnalyticsGetLeaderboardParamsMetric = "success-rate"
	APIV1AnalyticsGetLeaderboardParamsMetricActivity    APIV1AnalyticsGetLeaderboardParamsMetric = "activity"
)

type APIV1AnalyticsGetLeaderboardParamsPeriod string

const (
	APIV1AnalyticsGetLeaderboardParamsPeriodDaily   APIV1AnalyticsGetLeaderboardParamsPeriod = "daily"
	APIV1AnalyticsGetLeaderboardParamsPeriodWeekly  APIV1AnalyticsGetLeaderboardParamsPeriod = "weekly"
	APIV1AnalyticsGetLeaderboardParamsPeriodMonthly APIV1AnalyticsGetLeaderboardParamsPeriod = "monthly"
	APIV1AnalyticsGetLeaderboardParamsPeriodAllTime APIV1AnalyticsGetLeaderboardParamsPeriod = "all-time"
)

type APIV1AnalyticsGetSystemParams struct {
	// Any of "24h", "7d", "30d".
	Period APIV1AnalyticsGetSystemParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AnalyticsGetSystemParams]'s query parameters as
// `url.Values`.
func (r APIV1AnalyticsGetSystemParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AnalyticsGetSystemParamsPeriod string

const (
	APIV1AnalyticsGetSystemParamsPeriod24h APIV1AnalyticsGetSystemParamsPeriod = "24h"
	APIV1AnalyticsGetSystemParamsPeriod7d  APIV1AnalyticsGetSystemParamsPeriod = "7d"
	APIV1AnalyticsGetSystemParamsPeriod30d APIV1AnalyticsGetSystemParamsPeriod = "30d"
)
