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

// APIV1OrganizationService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1OrganizationService] method instead.
type APIV1OrganizationService struct {
	Options []option.RequestOption
	Members APIV1OrganizationMemberService
}

// NewAPIV1OrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV1OrganizationService(opts ...option.RequestOption) (r APIV1OrganizationService) {
	r = APIV1OrganizationService{}
	r.Options = opts
	r.Members = NewAPIV1OrganizationMemberService(opts...)
	return
}

// Get organization details by public ID (external access)
func (r *APIV1OrganizationService) Get(ctx context.Context, publicID string, opts ...option.RequestOption) (res *Apiv1OrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if publicID == "" {
		err = errors.New("missing required publicId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/organizations/public/%s", publicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update organization details
func (r *APIV1OrganizationService) Update(ctx context.Context, id string, body APIV1OrganizationUpdateParams, opts ...option.RequestOption) (res *Apiv1OrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/organizations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get organizations the current user belongs to
func (r *APIV1OrganizationService) List(ctx context.Context, query APIV1OrganizationListParams, opts ...option.RequestOption) (res *Apiv1OrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Creates an invite code for a user to join the organization
func (r *APIV1OrganizationService) Invites(ctx context.Context, id string, body APIV1OrganizationInvitesParams, opts ...option.RequestOption) (res *Apiv1OrganizationInvitesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/organizations/%s/invites", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Apiv1OrganizationGetResponse struct {
	Result  Apiv1OrganizationGetResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationGetResponseResult struct {
	ID        string  `json:"id,required"`
	CreatedAt float64 `json:"createdAt,required"`
	Name      string  `json:"name,required"`
	OwnerID   string  `json:"ownerId,required"`
	PublicID  string  `json:"publicId,required"`
	Settings  string  `json:"settings,required"`
	UpdatedAt float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		OwnerID     respjson.Field
		PublicID    respjson.Field
		Settings    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationUpdateResponse struct {
	Result  Apiv1OrganizationUpdateResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationUpdateResponseResult struct {
	ID        string  `json:"id,required"`
	CreatedAt float64 `json:"createdAt,required"`
	Name      string  `json:"name,required"`
	OwnerID   string  `json:"ownerId,required"`
	PublicID  string  `json:"publicId,required"`
	Settings  string  `json:"settings,required"`
	UpdatedAt float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		OwnerID     respjson.Field
		PublicID    respjson.Field
		Settings    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationUpdateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationUpdateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationListResponse struct {
	Result  []Apiv1OrganizationListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationListResponseResult struct {
	ID        string  `json:"id,required"`
	CreatedAt float64 `json:"createdAt,required"`
	Name      string  `json:"name,required"`
	OwnerID   string  `json:"ownerId,required"`
	PublicID  string  `json:"publicId,required"`
	Settings  string  `json:"settings,required"`
	UpdatedAt float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		OwnerID     respjson.Field
		PublicID    respjson.Field
		Settings    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationInvitesResponse struct {
	Properties Apiv1OrganizationInvitesResponseProperties `json:"properties,required"`
	Type       string                                     `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationInvitesResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationInvitesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationInvitesResponseProperties struct {
	ExpiresAt  Apiv1OrganizationInvitesResponsePropertiesExpiresAt  `json:"expiresAt,required"`
	InviteCode Apiv1OrganizationInvitesResponsePropertiesInviteCode `json:"inviteCode,required"`
	Message    Apiv1OrganizationInvitesResponsePropertiesMessage    `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		InviteCode  respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationInvitesResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationInvitesResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationInvitesResponsePropertiesExpiresAt struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationInvitesResponsePropertiesExpiresAt) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationInvitesResponsePropertiesExpiresAt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationInvitesResponsePropertiesInviteCode struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationInvitesResponsePropertiesInviteCode) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationInvitesResponsePropertiesInviteCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationInvitesResponsePropertiesMessage struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationInvitesResponsePropertiesMessage) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationInvitesResponsePropertiesMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1OrganizationUpdateParams struct {
	Name     param.Opt[string]                     `json:"name,omitzero"`
	Settings APIV1OrganizationUpdateParamsSettings `json:"settings,omitzero"`
	paramObj
}

func (r APIV1OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1OrganizationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1OrganizationUpdateParamsSettings struct {
	APIKeyQuota    param.Opt[int64] `json:"apiKeyQuota,omitzero"`
	BiddingEnabled param.Opt[bool]  `json:"biddingEnabled,omitzero"`
	// Any of "basic", "standard", "premium", "enterprise".
	RateLimitTier string `json:"rateLimitTier,omitzero"`
	paramObj
}

func (r APIV1OrganizationUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	type shadow APIV1OrganizationUpdateParamsSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1OrganizationUpdateParamsSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[APIV1OrganizationUpdateParamsSettings](
		"rateLimitTier", "basic", "standard", "premium", "enterprise",
	)
}

type APIV1OrganizationListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1OrganizationListParams]'s query parameters as
// `url.Values`.
func (r APIV1OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1OrganizationInvitesParams struct {
	Email string `json:"email,required" format:"email"`
	// Any of "admin", "member".
	Role      APIV1OrganizationInvitesParamsRole `json:"role,omitzero,required"`
	ExpiresIn param.Opt[int64]                   `json:"expiresIn,omitzero"`
	paramObj
}

func (r APIV1OrganizationInvitesParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1OrganizationInvitesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1OrganizationInvitesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1OrganizationInvitesParamsRole string

const (
	APIV1OrganizationInvitesParamsRoleAdmin  APIV1OrganizationInvitesParamsRole = "admin"
	APIV1OrganizationInvitesParamsRoleMember APIV1OrganizationInvitesParamsRole = "member"
)
