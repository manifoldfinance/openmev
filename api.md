# API

## V1

### Public

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicGetOverviewResponse">Apiv1PublicGetOverviewResponse</a>

Methods:

- <code title="get /api/v1/public/overview">client.API.V1.Public.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicService.GetOverview">GetOverview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicGetOverviewParams">APIV1PublicGetOverviewParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicGetOverviewResponse">Apiv1PublicGetOverviewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Auctions

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetResponse">Apiv1PublicAuctionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetHistoryResponse">Apiv1PublicAuctionGetHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetLatestResultsResponse">Apiv1PublicAuctionGetLatestResultsResponse</a>

Methods:

- <code title="get /api/v1/public/auctions/{id}">client.API.V1.Public.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionGetParams">APIV1PublicAuctionGetParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetResponse">Apiv1PublicAuctionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/public/auctions/history">client.API.V1.Public.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionGetHistoryParams">APIV1PublicAuctionGetHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetHistoryResponse">Apiv1PublicAuctionGetHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/public/auctions/latest-results">client.API.V1.Public.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionService.GetLatestResults">GetLatestResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1PublicAuctionGetLatestResultsParams">APIV1PublicAuctionGetLatestResultsParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1PublicAuctionGetLatestResultsResponse">Apiv1PublicAuctionGetLatestResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Auctions

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionNewResponse">Apiv1AuctionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionGetResponse">Apiv1AuctionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionListResponse">Apiv1AuctionListResponse</a>

Methods:

- <code title="post /api/v1/auctions">client.API.V1.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionNewParams">APIV1AuctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionNewResponse">Apiv1AuctionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/auctions/{id}">client.API.V1.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionGetParams">APIV1AuctionGetParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionGetResponse">Apiv1AuctionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/auctions">client.API.V1.Auctions.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionListParams">APIV1AuctionListParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionListResponse">Apiv1AuctionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionBidNewResponse">Apiv1AuctionBidNewResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionBidAsyncResponse">Apiv1AuctionBidAsyncResponse</a>

Methods:

- <code title="post /api/v1/auctions/{id}/bids">client.API.V1.Auctions.Bids.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionBidService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionBidNewParams">APIV1AuctionBidNewParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionBidNewResponse">Apiv1AuctionBidNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/auctions/{id}/bids/async">client.API.V1.Auctions.Bids.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionBidService.Async">Async</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuctionBidAsyncParams">APIV1AuctionBidAsyncParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuctionBidAsyncResponse">Apiv1AuctionBidAsyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1BidGetStatusResponse">Apiv1BidGetStatusResponse</a>

Methods:

- <code title="get /api/v1/bids/{trackingId}/status">client.API.V1.Bids.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1BidService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, trackingID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1BidGetStatusParams">APIV1BidGetStatusParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1BidGetStatusResponse">Apiv1BidGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthGetResponse">Apiv1AuthGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthGetValidateResponse">Apiv1AuthGetValidateResponse</a>

Methods:

- <code title="get /api/v1/auth/invites/{code}">client.API.V1.Auth.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, code <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthGetResponse">Apiv1AuthGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/auth/validate">client.API.V1.Auth.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthService.GetValidate">GetValidate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthGetValidateParams">APIV1AuthGetValidateParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthGetValidateResponse">Apiv1AuthGetValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthAPIKeyAPIKeysResponse">Apiv1AuthAPIKeyAPIKeysResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthAPIKeyGetAPIKeysResponse">Apiv1AuthAPIKeyGetAPIKeysResponse</a>

Methods:

- <code title="post /api/v1/auth/api-keys">client.API.V1.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthAPIKeyService.APIKeys">APIKeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthAPIKeyAPIKeysParams">APIV1AuthAPIKeyAPIKeysParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthAPIKeyAPIKeysResponse">Apiv1AuthAPIKeyAPIKeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/auth/api-keys">client.API.V1.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthAPIKeyService.GetAPIKeys">GetAPIKeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthAPIKeyGetAPIKeysParams">APIV1AuthAPIKeyGetAPIKeysParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthAPIKeyGetAPIKeysResponse">Apiv1AuthAPIKeyGetAPIKeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Register

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthRegisterMemberResponse">Apiv1AuthRegisterMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthRegisterOrganizationResponse">Apiv1AuthRegisterOrganizationResponse</a>

Methods:

- <code title="post /api/v1/auth/register/member">client.API.V1.Auth.Register.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthRegisterService.Member">Member</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthRegisterMemberParams">APIV1AuthRegisterMemberParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthRegisterMemberResponse">Apiv1AuthRegisterMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/auth/register/organization">client.API.V1.Auth.Register.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthRegisterService.Organization">Organization</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AuthRegisterOrganizationParams">APIV1AuthRegisterOrganizationParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AuthRegisterOrganizationResponse">Apiv1AuthRegisterOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Scheduler

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1SchedulerGetStatusResponse">Apiv1SchedulerGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1SchedulerTriggerResponse">Apiv1SchedulerTriggerResponse</a>

Methods:

- <code title="get /api/v1/scheduler/status">client.API.V1.Scheduler.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1SchedulerService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1SchedulerGetStatusResponse">Apiv1SchedulerGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/scheduler/trigger">client.API.V1.Scheduler.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1SchedulerService.Trigger">Trigger</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1SchedulerTriggerResponse">Apiv1SchedulerTriggerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserGetResponse">Apiv1UserGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserUpdateResponse">Apiv1UserUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserListResponse">Apiv1UserListResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserGetMeResponse">Apiv1UserGetMeResponse</a>

Methods:

- <code title="get /api/v1/users/public/{publicId}">client.API.V1.Users.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, publicID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserGetResponse">Apiv1UserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/users/{id}">client.API.V1.Users.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserUpdateParams">APIV1UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserUpdateResponse">Apiv1UserUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/users">client.API.V1.Users.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserListParams">APIV1UserListParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserListResponse">Apiv1UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/users/me">client.API.V1.Users.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1UserService.GetMe">GetMe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1UserGetMeResponse">Apiv1UserGetMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationGetResponse">Apiv1OrganizationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationUpdateResponse">Apiv1OrganizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationListResponse">Apiv1OrganizationListResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationInvitesResponse">Apiv1OrganizationInvitesResponse</a>

Methods:

- <code title="get /api/v1/organizations/public/{publicId}">client.API.V1.Organizations.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, publicID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationGetResponse">Apiv1OrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/v1/organizations/{id}">client.API.V1.Organizations.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationUpdateParams">APIV1OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationUpdateResponse">Apiv1OrganizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/organizations">client.API.V1.Organizations.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationListParams">APIV1OrganizationListParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationListResponse">Apiv1OrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/organizations/{id}/invites">client.API.V1.Organizations.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationService.Invites">Invites</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationInvitesParams">APIV1OrganizationInvitesParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationInvitesResponse">Apiv1OrganizationInvitesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Members

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationMemberListResponse">Apiv1OrganizationMemberListResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationMemberDeleteResponse">Apiv1OrganizationMemberDeleteResponse</a>

Methods:

- <code title="get /api/v1/organizations/{orgId}/members">client.API.V1.Organizations.Members.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationMemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationMemberListParams">APIV1OrganizationMemberListParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationMemberListResponse">Apiv1OrganizationMemberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/organizations/{orgId}/members/{userId}">client.API.V1.Organizations.Members.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationMemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1OrganizationMemberDeleteParams">APIV1OrganizationMemberDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1OrganizationMemberDeleteResponse">Apiv1OrganizationMemberDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Invites

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteGetResponse">Apiv1InviteGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteListResponse">Apiv1InviteListResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteDeleteResponse">Apiv1InviteDeleteResponse</a>

Methods:

- <code title="get /api/v1/invites/{code}">client.API.V1.Invites.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1InviteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, code <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteGetResponse">Apiv1InviteGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/invites">client.API.V1.Invites.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1InviteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1InviteListParams">APIV1InviteListParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteListResponse">Apiv1InviteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/invites/{code}">client.API.V1.Invites.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1InviteService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, code <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1InviteDeleteResponse">Apiv1InviteDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Analytics

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetResponse">Apiv1AnalyticsGetResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetLeaderboardResponse">Apiv1AnalyticsGetLeaderboardResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetSystemResponse">Apiv1AnalyticsGetSystemResponse</a>

Methods:

- <code title="get /api/v1/analytics/organizations/{id}">client.API.V1.Analytics.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsGetParams">APIV1AnalyticsGetParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetResponse">Apiv1AnalyticsGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/analytics/leaderboard">client.API.V1.Analytics.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsService.GetLeaderboard">GetLeaderboard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsGetLeaderboardParams">APIV1AnalyticsGetLeaderboardParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetLeaderboardResponse">Apiv1AnalyticsGetLeaderboardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/analytics/system">client.API.V1.Analytics.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsService.GetSystem">GetSystem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsGetSystemParams">APIV1AnalyticsGetSystemParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsGetSystemResponse">Apiv1AnalyticsGetSystemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsUserGetMeResponse">Apiv1AnalyticsUserGetMeResponse</a>

Methods:

- <code title="get /api/v1/analytics/users/me">client.API.V1.Analytics.Users.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsUserService.GetMe">GetMe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#APIV1AnalyticsUserGetMeParams">APIV1AnalyticsUserGetMeParams</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#Apiv1AnalyticsUserGetMeResponse">Apiv1AnalyticsUserGetMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthCheckResponse">HealthCheckResponse</a>
- <a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthCheckDetailedResponse">HealthCheckDetailedResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/detailed">client.Health.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthService.CheckDetailed">CheckDetailed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev">openmev</a>.<a href="https://pkg.go.dev/github.com/manifoldfinance/openmev#HealthCheckDetailedResponse">HealthCheckDetailedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
