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
	"github.com/stainless-sdks/openmev-go/packages/respjson"
)

// APIV1AnalyticsUserService contains methods and other services that help with
// interacting with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1AnalyticsUserService] method instead.
type APIV1AnalyticsUserService struct {
	Options []option.RequestOption
}

// NewAPIV1AnalyticsUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV1AnalyticsUserService(opts ...option.RequestOption) (r APIV1AnalyticsUserService) {
	r = APIV1AnalyticsUserService{}
	r.Options = opts
	return
}

// Returns activity metrics for the current user
func (r *APIV1AnalyticsUserService) GetMe(ctx context.Context, query APIV1AnalyticsUserGetMeParams, opts ...option.RequestOption) (res *Apiv1AnalyticsUserGetMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/analytics/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Apiv1AnalyticsUserGetMeResponse struct {
	Properties Apiv1AnalyticsUserGetMeResponseProperties `json:"properties,required"`
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
func (r Apiv1AnalyticsUserGetMeResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponseProperties struct {
	Activity      Apiv1AnalyticsUserGetMeResponsePropertiesActivity      `json:"activity,required"`
	Organizations Apiv1AnalyticsUserGetMeResponsePropertiesOrganizations `json:"organizations,required"`
	User          Apiv1AnalyticsUserGetMeResponsePropertiesUser          `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activity      respjson.Field
		Organizations respjson.Field
		User          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivity struct {
	Properties Apiv1AnalyticsUserGetMeResponsePropertiesActivityProperties `json:"properties,required"`
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
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivity) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityProperties struct {
	AvgResponseTime Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesAvgResponseTime `json:"avgResponseTime,required"`
	LastActive      Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesLastActive      `json:"lastActive,required"`
	Organizations   Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesOrganizations   `json:"organizations,required"`
	SuccessfulBids  Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesSuccessfulBids  `json:"successfulBids,required"`
	TotalBids       Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesTotalBids       `json:"totalBids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgResponseTime respjson.Field
		LastActive      respjson.Field
		Organizations   respjson.Field
		SuccessfulBids  respjson.Field
		TotalBids       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesAvgResponseTime struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesAvgResponseTime) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesAvgResponseTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesLastActive struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesLastActive) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesLastActive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesOrganizations struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesOrganizations) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesOrganizations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesSuccessfulBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesSuccessfulBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesSuccessfulBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesTotalBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesTotalBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesActivityPropertiesTotalBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizations struct {
	Items Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItems `json:"items,required"`
	Type  string                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizations) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItems struct {
	Properties Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsProperties `json:"properties,required"`
	Type       string                                                                `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItems) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsProperties struct {
	ID   Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesID   `json:"id,required"`
	Bids Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesBids `json:"bids,required"`
	Name Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesName `json:"name,required"`
	Role Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bids        respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesID struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesID) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesBids struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesBids) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesBids) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesName struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesName) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesRole struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesRole) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesOrganizationsItemsPropertiesRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesUser struct {
	Properties Apiv1AnalyticsUserGetMeResponsePropertiesUserProperties `json:"properties,required"`
	Type       string                                                  `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesUser) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesUserProperties struct {
	ID    Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesID    `json:"id,required"`
	Email Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesEmail `json:"email,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesUserProperties) RawJSON() string { return r.JSON.raw }
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesUserProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesID struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesID) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesEmail struct {
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesEmail) RawJSON() string {
	return r.JSON.raw
}
func (r *Apiv1AnalyticsUserGetMeResponsePropertiesUserPropertiesEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV1AnalyticsUserGetMeParams struct {
	// Any of "7d", "30d", "90d".
	Period APIV1AnalyticsUserGetMeParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV1AnalyticsUserGetMeParams]'s query parameters as
// `url.Values`.
func (r APIV1AnalyticsUserGetMeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV1AnalyticsUserGetMeParamsPeriod string

const (
	APIV1AnalyticsUserGetMeParamsPeriod7d  APIV1AnalyticsUserGetMeParamsPeriod = "7d"
	APIV1AnalyticsUserGetMeParamsPeriod30d APIV1AnalyticsUserGetMeParamsPeriod = "30d"
	APIV1AnalyticsUserGetMeParamsPeriod90d APIV1AnalyticsUserGetMeParamsPeriod = "90d"
)
