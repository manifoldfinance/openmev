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

// APIV1InviteService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1InviteService] method instead.
type APIV1InviteService struct {
	Options []option.RequestOption
}

// NewAPIV1InviteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1InviteService(opts ...option.RequestOption) (r APIV1InviteService) {
	r = APIV1InviteService{}
	r.Options = opts
	return
}

// Get invitation details by invite code
func (r *APIV1InviteService) Get(ctx context.Context, code string, opts ...option.RequestOption) (res *Apiv1InviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("api/v1/invites/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get invitations for the current organization
func (r *APIV1InviteService) List(ctx context.Context, query APIV1InviteListParams, opts ...option.RequestOption) (res *Apiv1InviteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an invitation by invite code
func (r *APIV1InviteService) Delete(ctx context.Context, code string, opts ...option.RequestOption) (res *Apiv1InviteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return
	}
	path := fmt.Sprintf("api/v1/invites/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Apiv1InviteGetResponse struct {
	Result  Apiv1InviteGetResponseResult `json:"result,required"`
	Success bool                         `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1InviteGetResponseResult struct {
	Code      string  `json:"code,required"`
	CreatedBy string  `json:"createdBy,required"`
	Data      string  `json:"data,required"`
	ExpiresAt float64 `json:"expiresAt,required"`
	OrgID     string  `json:"orgId,required"`
	// Any of "organization", "member".
	Type string `json:"type,required"`
	Used bool   `json:"used,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		CreatedBy   respjson.Field
		Data        respjson.Field
		ExpiresAt   respjson.Field
		OrgID       respjson.Field
		Type        respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1InviteListResponse struct {
	Result  []Apiv1InviteListResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1InviteListResponseResult struct {
	Code      string  `json:"code,required"`
	CreatedBy string  `json:"createdBy,required"`
	Data      string  `json:"data,required"`
	ExpiresAt float64 `json:"expiresAt,required"`
	OrgID     string  `json:"orgId,required"`
	// Any of "organization", "member".
	Type string `json:"type,required"`
	Used bool   `json:"used,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		CreatedBy   respjson.Field
		Data        respjson.Field
		ExpiresAt   respjson.Field
		OrgID       respjson.Field
		Type        respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1InviteDeleteResponse struct {
	Result  Apiv1InviteDeleteResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1InviteDeleteResponseResult struct {
	Code      string  `json:"code,required"`
	CreatedBy string  `json:"createdBy,required"`
	Data      string  `json:"data,required"`
	ExpiresAt float64 `json:"expiresAt,required"`
	OrgID     string  `json:"orgId,required"`
	// Any of "organization", "member".
	Type string `json:"type,required"`
	Used bool   `json:"used,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		CreatedBy   respjson.Field
		Data        respjson.Field
		ExpiresAt   respjson.Field
		OrgID       respjson.Field
		Type        respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1InviteDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1InviteDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1InviteListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1InviteListParams]'s query parameters as `url.Values`.
func (r APIV1InviteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
