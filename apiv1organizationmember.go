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

// APIV1OrganizationMemberService contains methods and other services that help
// with interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1OrganizationMemberService] method instead.
type APIV1OrganizationMemberService struct {
	Options []option.RequestOption
}

// NewAPIV1OrganizationMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV1OrganizationMemberService(opts ...option.RequestOption) (r APIV1OrganizationMemberService) {
	r = APIV1OrganizationMemberService{}
	r.Options = opts
	return
}

// Get members of an organization
func (r *APIV1OrganizationMemberService) List(ctx context.Context, orgID string, query APIV1OrganizationMemberListParams, opts ...option.RequestOption) (res *Apiv1OrganizationMemberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/organizations/%s/members", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove a user from an organization
func (r *APIV1OrganizationMemberService) Delete(ctx context.Context, userID string, body APIV1OrganizationMemberDeleteParams, opts ...option.RequestOption) (res *Apiv1OrganizationMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/organizations/%s/members/%s", body.OrgID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Apiv1OrganizationMemberListResponse struct {
	Result  []Apiv1OrganizationMemberListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationMemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationMemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationMemberListResponseResult struct {
	InvitedBy string  `json:"invitedBy,required"`
	JoinedAt  float64 `json:"joinedAt,required"`
	OrgID     string  `json:"orgId,required"`
	// Any of "owner", "admin", "member".
	Role   string `json:"role,required"`
	UserID string `json:"userId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvitedBy   respjson.Field
		JoinedAt    respjson.Field
		OrgID       respjson.Field
		Role        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationMemberListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationMemberListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationMemberDeleteResponse struct {
	Result  Apiv1OrganizationMemberDeleteResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationMemberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationMemberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1OrganizationMemberDeleteResponseResult struct {
	InvitedBy string  `json:"invitedBy,required"`
	JoinedAt  float64 `json:"joinedAt,required"`
	OrgID     string  `json:"orgId,required"`
	// Any of "owner", "admin", "member".
	Role   string `json:"role,required"`
	UserID string `json:"userId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvitedBy   respjson.Field
		JoinedAt    respjson.Field
		OrgID       respjson.Field
		Role        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1OrganizationMemberDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1OrganizationMemberDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1OrganizationMemberListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1OrganizationMemberListParams]'s query parameters as
// `url.Values`.
func (r APIV1OrganizationMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1OrganizationMemberDeleteParams struct {
	OrgID string `path:"orgId,required" json:"-"`
	paramObj
}
