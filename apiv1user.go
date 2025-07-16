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

// APIV1UserService contains methods and other services that help with interacting
// with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1UserService] method instead.
type APIV1UserService struct {
	Options []option.RequestOption
}

// NewAPIV1UserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV1UserService(opts ...option.RequestOption) (r APIV1UserService) {
	r = APIV1UserService{}
	r.Options = opts
	return
}

// Get user details by public ID (external access)
func (r *APIV1UserService) Get(ctx context.Context, publicID string, opts ...option.RequestOption) (res *Apiv1UserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if publicID == "" {
		err = errors.New("missing required publicId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/users/public/%s", publicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update user profile information
func (r *APIV1UserService) Update(ctx context.Context, id string, body APIV1UserUpdateParams, opts ...option.RequestOption) (res *Apiv1UserUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a paginated list of users in the current organization
func (r *APIV1UserService) List(ctx context.Context, query APIV1UserListParams, opts ...option.RequestOption) (res *Apiv1UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get current authenticated user profile
func (r *APIV1UserService) GetMe(ctx context.Context, opts ...option.RequestOption) (res *Apiv1UserGetMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Apiv1UserGetResponse struct {
	Result  Apiv1UserGetResponseResult `json:"result,required"`
	Success bool                       `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserGetResponseResult struct {
	ID         string  `json:"id,required"`
	CreatedAt  float64 `json:"createdAt,required"`
	Email      string  `json:"email,required"`
	EthAddress string  `json:"ethAddress,required"`
	PublicID   string  `json:"publicId,required"`
	UpdatedAt  float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		EthAddress  respjson.Field
		PublicID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserUpdateResponse struct {
	Result  Apiv1UserUpdateResponseResult `json:"result,required"`
	Success bool                          `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserUpdateResponseResult struct {
	ID         string  `json:"id,required"`
	CreatedAt  float64 `json:"createdAt,required"`
	Email      string  `json:"email,required"`
	EthAddress string  `json:"ethAddress,required"`
	PublicID   string  `json:"publicId,required"`
	UpdatedAt  float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		EthAddress  respjson.Field
		PublicID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserUpdateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserUpdateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserListResponse struct {
	Result  []Apiv1UserListResponseResult `json:"result,required"`
	Success bool                          `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserListResponseResult struct {
	ID         string  `json:"id,required"`
	CreatedAt  float64 `json:"createdAt,required"`
	Email      string  `json:"email,required"`
	EthAddress string  `json:"ethAddress,required"`
	PublicID   string  `json:"publicId,required"`
	UpdatedAt  float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		EthAddress  respjson.Field
		PublicID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserGetMeResponse struct {
	Result  Apiv1UserGetMeResponseResult `json:"result,required"`
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
func (r Apiv1UserGetMeResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserGetMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1UserGetMeResponseResult struct {
	ID         string  `json:"id,required"`
	CreatedAt  float64 `json:"createdAt,required"`
	Email      string  `json:"email,required"`
	EthAddress string  `json:"ethAddress,required"`
	PublicID   string  `json:"publicId,required"`
	UpdatedAt  float64 `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		EthAddress  respjson.Field
		PublicID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1UserGetMeResponseResult) RawJSON() string { return r.JSON.raw }
func (r *Apiv1UserGetMeResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1UserUpdateParams struct {
	Email      param.Opt[string] `json:"email,omitzero" format:"email"`
	EthAddress param.Opt[string] `json:"ethAddress,omitzero"`
	paramObj
}

func (r APIV1UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1UserListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1UserListParams]'s query parameters as `url.Values`.
func (r APIV1UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
