package wrapper

import (
	"context"
	"net/http"
	"time"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)

// Factory creates properly configured Wrappers that share underlying resources
// (RateLimiter, HttpClient) to ensure safe high-throughput usage.
type Factory struct {
	limiter       services.RateLimiter
	clientOptions []client.ClientOption
}

// NewFactory creates a new Factory with a shared RateLimiter and HttpClient.
// All Wrappers created by this Factory will share the same Integrator limits
// and connection pool.
func NewFactory(config services.RateLimiterConfig, opts ...client.ClientOption) *Factory {
	dummyCfg := &client.ClientConfig{}
	for _, opt := range opts {
		opt(dummyCfg)
	}

	sharedClient := dummyCfg.HTTPClient
	if sharedClient == nil {
		sharedClient = &http.Client{Timeout: 100 * time.Second}
	}

	finalOpts := append([]client.ClientOption{}, opts...)
	finalOpts = append(finalOpts, client.WithHTTPClient(sharedClient))

	return &Factory{
		limiter:       services.NewRateLimiter(config, dummyCfg.Logger),
		clientOptions: finalOpts,
	}
}

// GetWrapper creates a new Wrapper for a specific Facility/License configuration.
// It reuses the Factory's RateLimiter and HttpClient.
func (f *Factory) GetWrapper(baseURL, vendorKey, userKey string) *MetrcWrapper {
	// Create a new client using the stored options (which include the shared httpClient)
	c := client.New(baseURL, vendorKey, userKey, f.clientOptions...)

	// Create wrapper using shared limiter
	return NewWithLimiter(c, f.limiter)
}

// Execute is a helper to run a function with the shared rate limiter context directly,
// if you need to manually wrap external calls.
func (f *Factory) Execute(ctx context.Context, facility string, isGet bool, op func() (interface{}, error)) (interface{}, error) {
	return f.limiter.Execute(ctx, facility, isGet, op)
}

