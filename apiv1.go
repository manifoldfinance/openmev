// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"github.com/manifoldfinance/openmev/option"
)

// APIV1Service contains methods and other services that help with interacting with
// the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV1Service] method instead.
type APIV1Service struct {
	Options       []option.RequestOption
	Public        APIV1PublicService
	Auctions      APIV1AuctionService
	Bids          APIV1BidService
	Auth          APIV1AuthService
	Scheduler     APIV1SchedulerService
	Users         APIV1UserService
	Organizations APIV1OrganizationService
	Invites       APIV1InviteService
	Analytics     APIV1AnalyticsService
}

// NewAPIV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIV1Service(opts ...option.RequestOption) (r APIV1Service) {
	r = APIV1Service{}
	r.Options = opts
	r.Public = NewAPIV1PublicService(opts...)
	r.Auctions = NewAPIV1AuctionService(opts...)
	r.Bids = NewAPIV1BidService(opts...)
	r.Auth = NewAPIV1AuthService(opts...)
	r.Scheduler = NewAPIV1SchedulerService(opts...)
	r.Users = NewAPIV1UserService(opts...)
	r.Organizations = NewAPIV1OrganizationService(opts...)
	r.Invites = NewAPIV1InviteService(opts...)
	r.Analytics = NewAPIV1AnalyticsService(opts...)
	return
}
