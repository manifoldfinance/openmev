// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"net/http"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/param"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// APIV1AuthRegisterService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AuthRegisterService] method instead.
type APIV1AuthRegisterService struct {
	Options []option.RequestOption
}

// NewAPIV1AuthRegisterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV1AuthRegisterService(opts ...option.RequestOption) (r APIV1AuthRegisterService) {
	r = APIV1AuthRegisterService{}
	r.Options = opts
	return
}

// Joins an existing organization as a member using an invite code
func (r *APIV1AuthRegisterService) Member(ctx context.Context, body APIV1AuthRegisterMemberParams, opts ...option.RequestOption) (res *Apiv1AuthRegisterMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auth/register/member"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new organization and assigns the registering user as owner
func (r *APIV1AuthRegisterService) Organization(ctx context.Context, body APIV1AuthRegisterOrganizationParams, opts ...option.RequestOption) (res *Apiv1AuthRegisterOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/auth/register/organization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Apiv1AuthRegisterMemberResponse struct {
	APIKey       string                                      `json:"apiKey,required"`
	Organization Apiv1AuthRegisterMemberResponseOrganization `json:"organization,required"`
	User         Apiv1AuthRegisterMemberResponseUser         `json:"user,required"`
	// Any of "owner", "admin", "member".
	Role Apiv1AuthRegisterMemberResponseRole `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey       respjson.Field
		Organization respjson.Field
		User         respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthRegisterMemberResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterMemberResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterMemberResponseOrganization struct {
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
func (r Apiv1AuthRegisterMemberResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterMemberResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterMemberResponseUser struct {
	ID    string `json:"id,required"`
	Email string `json:"email,required" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthRegisterMemberResponseUser) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterMemberResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterMemberResponseRole string

const (
	Apiv1AuthRegisterMemberResponseRoleOwner  Apiv1AuthRegisterMemberResponseRole = "owner"
	Apiv1AuthRegisterMemberResponseRoleAdmin  Apiv1AuthRegisterMemberResponseRole = "admin"
	Apiv1AuthRegisterMemberResponseRoleMember Apiv1AuthRegisterMemberResponseRole = "member"
)

type Apiv1AuthRegisterOrganizationResponse struct {
	APIKey       string                                            `json:"apiKey,required"`
	Organization Apiv1AuthRegisterOrganizationResponseOrganization `json:"organization,required"`
	User         Apiv1AuthRegisterOrganizationResponseUser         `json:"user,required"`
	// Any of "owner", "admin", "member".
	Role Apiv1AuthRegisterOrganizationResponseRole `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey       respjson.Field
		Organization respjson.Field
		User         respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthRegisterOrganizationResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterOrganizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterOrganizationResponseOrganization struct {
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
func (r Apiv1AuthRegisterOrganizationResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterOrganizationResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterOrganizationResponseUser struct {
	ID    string `json:"id,required"`
	Email string `json:"email,required" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AuthRegisterOrganizationResponseUser) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AuthRegisterOrganizationResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AuthRegisterOrganizationResponseRole string

const (
	Apiv1AuthRegisterOrganizationResponseRoleOwner  Apiv1AuthRegisterOrganizationResponseRole = "owner"
	Apiv1AuthRegisterOrganizationResponseRoleAdmin  Apiv1AuthRegisterOrganizationResponseRole = "admin"
	Apiv1AuthRegisterOrganizationResponseRoleMember Apiv1AuthRegisterOrganizationResponseRole = "member"
)

type APIV1AuthRegisterMemberParams struct {
	InviteCode string                            `json:"inviteCode,required"`
	User       APIV1AuthRegisterMemberParamsUser `json:"user,omitzero,required"`
	paramObj
}

func (r APIV1AuthRegisterMemberParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterMemberParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterMemberParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Email is required.
type APIV1AuthRegisterMemberParamsUser struct {
	Email      string            `json:"email,required" format:"email"`
	EthAddress param.Opt[string] `json:"ethAddress,omitzero"`
	paramObj
}

func (r APIV1AuthRegisterMemberParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterMemberParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterMemberParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuthRegisterOrganizationParams struct {
	InviteCode   string                                          `json:"inviteCode,required"`
	Organization APIV1AuthRegisterOrganizationParamsOrganization `json:"organization,omitzero,required"`
	User         APIV1AuthRegisterOrganizationParamsUser         `json:"user,omitzero,required"`
	paramObj
}

func (r APIV1AuthRegisterOrganizationParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterOrganizationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterOrganizationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type APIV1AuthRegisterOrganizationParamsOrganization struct {
	Name     string                                                  `json:"name,required"`
	Settings APIV1AuthRegisterOrganizationParamsOrganizationSettings `json:"settings,omitzero"`
	paramObj
}

func (r APIV1AuthRegisterOrganizationParamsOrganization) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterOrganizationParamsOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterOrganizationParamsOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AuthRegisterOrganizationParamsOrganizationSettings struct {
	APIKeyQuota    param.Opt[int64] `json:"apiKeyQuota,omitzero"`
	BiddingEnabled param.Opt[bool]  `json:"biddingEnabled,omitzero"`
	// Any of "basic", "standard", "premium", "enterprise".
	RateLimitTier string `json:"rateLimitTier,omitzero"`
	paramObj
}

func (r APIV1AuthRegisterOrganizationParamsOrganizationSettings) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterOrganizationParamsOrganizationSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterOrganizationParamsOrganizationSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[APIV1AuthRegisterOrganizationParamsOrganizationSettings](
		"rateLimitTier", "basic", "standard", "premium", "enterprise",
	)
}

// The property Email is required.
type APIV1AuthRegisterOrganizationParamsUser struct {
	Email      string            `json:"email,required" format:"email"`
	EthAddress param.Opt[string] `json:"ethAddress,omitzero"`
	paramObj
}

func (r APIV1AuthRegisterOrganizationParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow APIV1AuthRegisterOrganizationParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1AuthRegisterOrganizationParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
