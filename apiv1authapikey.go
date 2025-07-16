// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/openmev-go/internal/apijson"
	"github.com/stainless-sdks/openmev-go/internal/apiquery"
	"github.com/stainless-sdks/openmev-go/internal/requestconfig"
	"github.com/stainless-sdks/openmev-go/option"
	"github.com/stainless-sdks/openmev-go/packages/param"
	"github.com/stainless-sdks/openmev-go/packages/respjson"
)

// APIV1AuthAPIKeyService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AuthAPIKeyService] method instead.
type APIV1AuthAPIKeyService struct {
	Options []option.RequestOption
}

// NewAPIV1AuthAPIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1AuthAPIKeyService(opts ...option.RequestOption) (r APIV1AuthAPIKeyService) {
	r = APIV1AuthAPIKeyService{}
	r.Options = opts
	return
}

// Create a new API key with specified permissions and optional expiration.
// Requires authentication with admin permissions.
func (r *APIV1AuthAPIKeyService) APIKeys(ctx context.Context, body APIV1AuthAPIKeyAPIKeysParams, opts ...option.RequestOption) (res *Apiv1AuthAPIKeyAPIKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of API keys for the current organization with optional
// filtering. Requires authentication with admin permissions.
func (r *APIV1AuthAPIKeyService) GetAPIKeys(ctx context.Context, query APIV1AuthAPIKeyGetAPIKeysParams, opts ...option.RequestOption) (res *Apiv1AuthAPIKeyGetAPIKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1AuthAPIKeyAPIKeysResponse struct {
	APIKey Apiv1AuthAPIKeyAPIKeysResponseAPIKey `json:"apiKey,required"`
	// Success message
	Message string `json:"message,required"`
	// Security reminder message
	SecurityReminder string `json:"securityReminder,required"`
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey           respjson.Field
		Message          respjson.Field
		SecurityReminder respjson.Field
		Success          respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthAPIKeyAPIKeysResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyAPIKeysResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthAPIKeyAPIKeysResponseAPIKey struct {
	// API key identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Key status
	IsActive bool `json:"isActive,required"`
	// The actual API key value (only shown once)
	Key string `json:"key,required"`
	// Preview of the key (last 8 characters)
	KeyPreview string `json:"keyPreview,required"`
	// API key name
	Name string `json:"name,required"`
	// API key permissions
	//
	// Any of "read", "bid", "auction_manage", "admin".
	Permissions []string `json:"permissions,required"`
	// Expiration timestamp
	ExpiresAt float64 `json:"expiresAt"`
	// Last usage timestamp
	LastUsed float64 `json:"lastUsed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		Key         respjson.Field
		KeyPreview  respjson.Field
		Name        respjson.Field
		Permissions respjson.Field
		ExpiresAt   respjson.Field
		LastUsed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthAPIKeyAPIKeysResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyAPIKeysResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthAPIKeyGetAPIKeysResponse struct {
	Count    float64                                   `json:"count,required"`
	Data     []Apiv1AuthAPIKeyGetAPIKeysResponseData   `json:"data,required"`
	HasMore  bool                                      `json:"hasMore,required"`
	Metadata Apiv1AuthAPIKeyGetAPIKeysResponseMetadata `json:"metadata,required"`
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
func (r Apiv1AuthAPIKeyGetAPIKeysResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyGetAPIKeysResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthAPIKeyGetAPIKeysResponseData struct {
	// API key identifier
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Key status
	IsActive bool `json:"isActive,required"`
	// Preview of the key (last 8 characters)
	KeyPreview string `json:"keyPreview,required"`
	// API key name
	Name string `json:"name,required"`
	// API key permissions
	//
	// Any of "read", "bid", "auction_manage", "admin".
	Permissions []string `json:"permissions,required"`
	// Expiration timestamp
	ExpiresAt float64 `json:"expiresAt"`
	// Last usage timestamp
	LastUsed float64                                    `json:"lastUsed"`
	Usage    Apiv1AuthAPIKeyGetAPIKeysResponseDataUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		KeyPreview  respjson.Field
		Name        respjson.Field
		Permissions respjson.Field
		ExpiresAt   respjson.Field
		LastUsed    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthAPIKeyGetAPIKeysResponseData) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyGetAPIKeysResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthAPIKeyGetAPIKeysResponseDataUsage struct {
	DailyRequests   int64  `json:"dailyRequests,required"`
	HourlyRequests  int64  `json:"hourlyRequests,required"`
	TotalRequests   int64  `json:"totalRequests,required"`
	LastRequestPath string `json:"lastRequestPath"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyRequests   respjson.Field
		HourlyRequests  respjson.Field
		TotalRequests   respjson.Field
		LastRequestPath respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthAPIKeyGetAPIKeysResponseDataUsage) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyGetAPIKeysResponseDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthAPIKeyGetAPIKeysResponseMetadata struct {
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
func (r Apiv1AuthAPIKeyGetAPIKeysResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthAPIKeyGetAPIKeysResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuthAPIKeyAPIKeysParams struct {
	// API key name
	Name string `json:"name,required"`
	// API key permissions
	//
	// Any of "read", "bid", "auction_manage", "admin".
	Permissions []string `json:"permissions,omitzero,required"`
	// API key description
	Description param.Opt[string] `json:"description,omitzero"`
	// Expiration timestamp
	ExpiresAt param.Opt[float64] `json:"expiresAt,omitzero"`
	// Additional metadata
	Metadata map[string]any `json:"metadata,omitzero"`
	// Custom rate limits for this key
	RateLimitOverride APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride `json:"rateLimitOverride,omitzero"`
	paramObj
}

func (r APIV1AuthAPIKeyAPIKeysParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthAPIKeyAPIKeysParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthAPIKeyAPIKeysParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom rate limits for this key
type APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride struct {
	RequestsPerDay    param.Opt[int64] `json:"requestsPerDay,omitzero"`
	RequestsPerHour   param.Opt[int64] `json:"requestsPerHour,omitzero"`
	RequestsPerMinute param.Opt[int64] `json:"requestsPerMinute,omitzero"`
	paramObj
}

func (r APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthAPIKeyAPIKeysParamsRateLimitOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuthAPIKeyGetAPIKeysParams struct {
	// Include inactive keys
	IncludeInactive param.Opt[bool] `query:"include_inactive,omitzero" json:"-"`
	// Filter by creator user ID
	CreatedBy param.Opt[string] `query:"created_by,omitzero" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by permission
	//
	// Any of "read", "bid", "auction_manage", "admin".
	Permission APIV1AuthAPIKeyGetAPIKeysParamsPermission `query:"permission,omitzero" json:"-"`
	// Sort field
	//
	// Any of "name", "createdAt", "lastUsed", "expiresAt".
	SortBy APIV1AuthAPIKeyGetAPIKeysParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	SortOrder APIV1AuthAPIKeyGetAPIKeysParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AuthAPIKeyGetAPIKeysParams]'s query parameters as
// `url.Values`.
func (r APIV1AuthAPIKeyGetAPIKeysParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by permission
type APIV1AuthAPIKeyGetAPIKeysParamsPermission string

const (
	APIV1AuthAPIKeyGetAPIKeysParamsPermissionRead          APIV1AuthAPIKeyGetAPIKeysParamsPermission = "read"
	APIV1AuthAPIKeyGetAPIKeysParamsPermissionBid           APIV1AuthAPIKeyGetAPIKeysParamsPermission = "bid"
	APIV1AuthAPIKeyGetAPIKeysParamsPermissionAuctionManage APIV1AuthAPIKeyGetAPIKeysParamsPermission = "auction_manage"
	APIV1AuthAPIKeyGetAPIKeysParamsPermissionAdmin         APIV1AuthAPIKeyGetAPIKeysParamsPermission = "admin"
)

// Sort field
type APIV1AuthAPIKeyGetAPIKeysParamsSortBy string

const (
	APIV1AuthAPIKeyGetAPIKeysParamsSortByName      APIV1AuthAPIKeyGetAPIKeysParamsSortBy = "name"
	APIV1AuthAPIKeyGetAPIKeysParamsSortByCreatedAt APIV1AuthAPIKeyGetAPIKeysParamsSortBy = "createdAt"
	APIV1AuthAPIKeyGetAPIKeysParamsSortByLastUsed  APIV1AuthAPIKeyGetAPIKeysParamsSortBy = "lastUsed"
	APIV1AuthAPIKeyGetAPIKeysParamsSortByExpiresAt APIV1AuthAPIKeyGetAPIKeysParamsSortBy = "expiresAt"
)

// Sort order
type APIV1AuthAPIKeyGetAPIKeysParamsSortOrder string

const (
	APIV1AuthAPIKeyGetAPIKeysParamsSortOrderAsc  APIV1AuthAPIKeyGetAPIKeysParamsSortOrder = "asc"
	APIV1AuthAPIKeyGetAPIKeysParamsSortOrderDesc APIV1AuthAPIKeyGetAPIKeysParamsSortOrder = "desc"
)
