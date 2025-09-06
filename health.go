// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmev

import (
	"context"
	"net/http"

	"github.com/manifoldfinance/openmev/internal/apijson"
	"github.com/manifoldfinance/openmev/internal/requestconfig"
	"github.com/manifoldfinance/openmev/option"
	"github.com/manifoldfinance/openmev/packages/respjson"
)

// HealthService contains methods and other services that help with interacting
// with the openmev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Perform a basic health check of the system. Returns overall health status,
// uptime, and version information.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Perform a comprehensive health check of all system components including
// database, KV store, Durable Objects, and performance metrics.
func (r *HealthService) CheckDetailed(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckDetailedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/detailed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthCheckResponse struct {
	// Environment name
	Environment string `json:"environment,required"`
	// Health status message
	Message string `json:"message,required"`
	// Overall system health
	//
	// Any of "healthy", "degraded", "unhealthy".
	Status HealthCheckResponseStatus `json:"status,required"`
	// Any of true.
	Success bool `json:"success,required"`
	// Health check timestamp
	Timestamp float64 `json:"timestamp,required"`
	// System uptime in milliseconds
	Uptime float64 `json:"uptime,required"`
	// Application version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Environment respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Uptime      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Overall system health
type HealthCheckResponseStatus string

const (
	HealthCheckResponseStatusHealthy   HealthCheckResponseStatus = "healthy"
	HealthCheckResponseStatusDegraded  HealthCheckResponseStatus = "degraded"
	HealthCheckResponseStatusUnhealthy HealthCheckResponseStatus = "unhealthy"
)

type HealthCheckDetailedResponse struct {
	Components HealthCheckDetailedResponseComponents `json:"components,required"`
	// Environment name
	Environment string `json:"environment,required"`
	// Detailed health status message
	Message     string                                 `json:"message,required"`
	Performance HealthCheckDetailedResponsePerformance `json:"performance,required"`
	Statistics  HealthCheckDetailedResponseStatistics  `json:"statistics,required"`
	// Overall system health
	//
	// Any of "healthy", "degraded", "unhealthy".
	Status HealthCheckDetailedResponseStatus `json:"status,required"`
	// Any of true.
	Success bool `json:"success,required"`
	// Health check timestamp
	Timestamp float64 `json:"timestamp,required"`
	// System uptime in milliseconds
	Uptime float64 `json:"uptime,required"`
	// Application version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Components  respjson.Field
		Environment respjson.Field
		Message     respjson.Field
		Performance respjson.Field
		Statistics  respjson.Field
		Status      respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Uptime      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckDetailedResponseComponents struct {
	// Analytics Engine health
	Analytics HealthCheckDetailedResponseComponentsAnalytics `json:"analytics,required"`
	// Cache system health
	Cache HealthCheckDetailedResponseComponentsCache `json:"cache,required"`
	// Database health
	Database HealthCheckDetailedResponseComponentsDatabase `json:"database,required"`
	// Durable Objects health
	DurableObjects HealthCheckDetailedResponseComponentsDurableObjects `json:"durableObjects,required"`
	// External services health
	ExternalServices HealthCheckDetailedResponseComponentsExternalServices `json:"externalServices,required"`
	// KV Store health
	KvStore HealthCheckDetailedResponseComponentsKvStore `json:"kvStore,required"`
	// Rate Limiter health
	RateLimiter HealthCheckDetailedResponseComponentsRateLimiter `json:"rateLimiter,required"`
	// Auction Scheduler health
	Scheduler HealthCheckDetailedResponseComponentsScheduler `json:"scheduler,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analytics        respjson.Field
		Cache            respjson.Field
		Database         respjson.Field
		DurableObjects   respjson.Field
		ExternalServices respjson.Field
		KvStore          respjson.Field
		RateLimiter      respjson.Field
		Scheduler        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponents) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Analytics Engine health
type HealthCheckDetailedResponseComponentsAnalytics struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsAnalytics) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsAnalytics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache system health
type HealthCheckDetailedResponseComponentsCache struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsCache) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Database health
type HealthCheckDetailedResponseComponentsDatabase struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsDatabase) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Durable Objects health
type HealthCheckDetailedResponseComponentsDurableObjects struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsDurableObjects) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsDurableObjects) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External services health
type HealthCheckDetailedResponseComponentsExternalServices struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsExternalServices) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsExternalServices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KV Store health
type HealthCheckDetailedResponseComponentsKvStore struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsKvStore) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsKvStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate Limiter health
type HealthCheckDetailedResponseComponentsRateLimiter struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsRateLimiter) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsRateLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Auction Scheduler health
type HealthCheckDetailedResponseComponentsScheduler struct {
	// Last health check timestamp
	LastCheck float64 `json:"lastCheck,required"`
	// Response time in milliseconds
	ResponseTime float64 `json:"responseTime,required"`
	// Any of "healthy", "degraded", "unhealthy".
	Status string `json:"status,required"`
	// Additional health details
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastCheck    respjson.Field
		ResponseTime respjson.Field
		Status       respjson.Field
		Details      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseComponentsScheduler) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseComponentsScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckDetailedResponsePerformance struct {
	// Active connections count
	ActiveConnections int64 `json:"activeConnections,required"`
	// Average response time in milliseconds
	AverageResponseTime float64 `json:"averageResponseTime,required"`
	// CPU usage percentage
	CPUUsage float64 `json:"cpuUsage,required"`
	// Memory usage percentage
	MemoryUsage float64 `json:"memoryUsage,required"`
	// Queued background jobs
	QueuedJobs int64 `json:"queuedJobs,required"`
	// Current requests per second
	RequestsPerSecond float64 `json:"requestsPerSecond,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveConnections   respjson.Field
		AverageResponseTime respjson.Field
		CPUUsage            respjson.Field
		MemoryUsage         respjson.Field
		QueuedJobs          respjson.Field
		RequestsPerSecond   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponsePerformance) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponsePerformance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckDetailedResponseStatistics struct {
	// Active API keys
	ActiveAPIKeys int64 `json:"activeApiKeys,required"`
	// Total auctions created
	TotalAuctions int64 `json:"totalAuctions,required"`
	// Total bids placed
	TotalBids int64 `json:"totalBids,required"`
	// Total errors encountered
	TotalErrors int64 `json:"totalErrors,required"`
	// Total organizations
	TotalOrganizations int64 `json:"totalOrganizations,required"`
	// Total requests processed
	TotalRequests int64 `json:"totalRequests,required"`
	// Total registered users
	TotalUsers int64 `json:"totalUsers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveAPIKeys      respjson.Field
		TotalAuctions      respjson.Field
		TotalBids          respjson.Field
		TotalErrors        respjson.Field
		TotalOrganizations respjson.Field
		TotalRequests      respjson.Field
		TotalUsers         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckDetailedResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckDetailedResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Overall system health
type HealthCheckDetailedResponseStatus string

const (
	HealthCheckDetailedResponseStatusHealthy   HealthCheckDetailedResponseStatus = "healthy"
	HealthCheckDetailedResponseStatusDegraded  HealthCheckDetailedResponseStatus = "degraded"
	HealthCheckDetailedResponseStatusUnhealthy HealthCheckDetailedResponseStatus = "unhealthy"
)
