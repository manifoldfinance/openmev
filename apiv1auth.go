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

// APIV1AuthService contains methods and other services that help with interacting
// with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AuthService] method instead.
type APIV1AuthService struct {
	Options  []option.RequestOption
	APIKeys  APIV1AuthAPIKeyService
	Register APIV1AuthRegisterService
}

// NewAPIV1AuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1AuthService(opts ...option.RequestOption) (r APIV1AuthService) {
	r = APIV1AuthService{}
	r.Options = opts
	r.APIKeys = NewAPIV1AuthAPIKeyService(opts...)
	r.Register = NewAPIV1AuthRegisterService(opts...)
	return
}

// Checks if an invite code is valid and returns invitation details
func (r *APIV1AuthService) Get(ctx context.Context, code string, opts ...option.RequestOption) (res *Apiv1AuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("api/v1/auth/invites/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Validate the API key provided in the request and return detailed information
// about the key, user, and organization.
func (r *APIV1AuthService) GetValidate(ctx context.Context, query APIV1AuthGetValidateParams, opts ...option.RequestOption) (res *Apiv1AuthGetValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auth/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1AuthGetResponse struct {
	Organization Apiv1AuthGetResponseOrganization `json:"organization,required"`
	Valid        bool                             `json:"valid,required"`
	Metadata     map[string]any                   `json:"metadata"`
	Reason       string                           `json:"reason"`
	// Any of "organization", "member".
	Type Apiv1AuthGetResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Valid        respjson.Field
		Metadata     respjson.Field
		Reason       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetResponseOrganization struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetResponseType string

const (
	Apiv1AuthGetResponseTypeOrganization Apiv1AuthGetResponseType = "organization"
	Apiv1AuthGetResponseTypeMember       Apiv1AuthGetResponseType = "member"
)

type Apiv1AuthGetValidateResponse struct {
	// Any of true.
	Success   bool    `json:"success,required"`
	Timestamp float64 `json:"timestamp,required"`
	// Whether the API key is valid
	Valid        bool                                     `json:"valid,required"`
	APIKey       Apiv1AuthGetValidateResponseAPIKey       `json:"apiKey"`
	Organization Apiv1AuthGetValidateResponseOrganization `json:"organization"`
	User         Apiv1AuthGetValidateResponseUser         `json:"user"`
	// Security warnings
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success      respjson.Field
		Timestamp    respjson.Field
		Valid        respjson.Field
		APIKey       respjson.Field
		Organization respjson.Field
		User         respjson.Field
		Warnings     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKey struct {
	// API key ID
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// Any of "read", "bid", "auction_manage", "admin".
	Permissions []string                                `json:"permissions,required"`
	Usage       Apiv1AuthGetValidateResponseAPIKeyUsage `json:"usage,required"`
	// Expiration timestamp
	ExpiresAt float64 `json:"expiresAt"`
	// Last usage timestamp
	LastUsed float64 `json:"lastUsed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Permissions respjson.Field
		Usage       respjson.Field
		ExpiresAt   respjson.Field
		LastUsed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKeyUsage struct {
	DailyRequests   int64                                                  `json:"dailyRequests,required"`
	HourlyRequests  int64                                                  `json:"hourlyRequests,required"`
	RateLimitStatus Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatus `json:"rateLimitStatus,required"`
	TotalRequests   int64                                                  `json:"totalRequests,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyRequests   respjson.Field
		HourlyRequests  respjson.Field
		RateLimitStatus respjson.Field
		TotalRequests   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKeyUsage) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseAPIKeyUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatus struct {
	RequestsPerDay    Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerDay    `json:"requestsPerDay,required"`
	RequestsPerHour   Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerHour   `json:"requestsPerHour,required"`
	RequestsPerMinute Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerMinute `json:"requestsPerMinute,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestsPerDay    respjson.Field
		RequestsPerHour   respjson.Field
		RequestsPerMinute respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatus) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerDay struct {
	Limit     int64 `json:"limit,required"`
	Remaining int64 `json:"remaining,required"`
	// Rate limit reset timestamp
	ResetAt float64 `json:"resetAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerDay) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerDay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerHour struct {
	Limit     int64 `json:"limit,required"`
	Remaining int64 `json:"remaining,required"`
	// Rate limit reset timestamp
	ResetAt float64 `json:"resetAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerHour) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerMinute struct {
	Limit     int64 `json:"limit,required"`
	Remaining int64 `json:"remaining,required"`
	// Rate limit reset timestamp
	ResetAt float64 `json:"resetAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerMinute) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AuthGetValidateResponseAPIKeyUsageRateLimitStatusRequestsPerMinute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseOrganization struct {
	// Internal organization ID
	ID string `json:"id,required"`
	// Organization creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Organization status
	IsActive    bool  `json:"isActive,required"`
	MemberCount int64 `json:"memberCount,required"`
	// Organization name
	Name string `json:"name,required"`
	// Public organization identifier
	PublicID string                                           `json:"publicId,required"`
	Settings Apiv1AuthGetValidateResponseOrganizationSettings `json:"settings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		MemberCount respjson.Field
		Name        respjson.Field
		PublicID    respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseOrganizationSettings struct {
	APIKeyQuota    int64 `json:"apiKeyQuota,required"`
	BiddingEnabled bool  `json:"biddingEnabled,required"`
	// Any of "basic", "standard", "premium", "enterprise".
	RateLimitTier string `json:"rateLimitTier,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyQuota    respjson.Field
		BiddingEnabled respjson.Field
		RateLimitTier  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseOrganizationSettings) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseOrganizationSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthGetValidateResponseUser struct {
	// Internal user ID
	ID string `json:"id,required"`
	// Account creation timestamp
	CreatedAt float64 `json:"createdAt,required"`
	// Account status
	IsActive bool `json:"isActive,required"`
	// Organization ID
	OrganizationID string `json:"organizationId,required"`
	// Public user identifier
	PublicID string `json:"publicId,required"`
	// Any of "admin", "auction_manager", "bidder", "read_only".
	Role string `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		IsActive       respjson.Field
		OrganizationID respjson.Field
		PublicID       respjson.Field
		Role           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthGetValidateResponseUser) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthGetValidateResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuthGetValidateParams struct {
	// Include detailed usage statistics
	IncludeUsage param.Opt[bool] `query:"include_usage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AuthGetValidateParams]'s query parameters as
// `url.Values`.
func (r APIV1AuthGetValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
