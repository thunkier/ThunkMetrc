package wrapper

import (
	"context"
	"strings"
	"sync"
	"time"
)

type RateLimiterConfig struct {
	Enabled                     bool
	MaxGetPerSecondPerFacility  float64
	MaxGetPerSecondIntegrator   float64
	MaxConcurrentGetPerFacility int
	MaxConcurrentGetIntegrator  int
}

func DefaultConfig() RateLimiterConfig {
	return RateLimiterConfig{
		Enabled:                     false,
		MaxGetPerSecondPerFacility:  50,
		MaxGetPerSecondIntegrator:   150,
		MaxConcurrentGetPerFacility: 10,
		MaxConcurrentGetIntegrator:  30,
	}
}

type TokenBucket struct {
	rate       float64
	capacity   float64
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(rate float64, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Wait(ctx context.Context) error {
	tb.mu.Lock()
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now

	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		tb.mu.Unlock()
		return nil
	}

	// Calculate wait time
	missing := 1.0 - tb.tokens
	waitSecs := missing / tb.rate
	wait := time.Duration(waitSecs * float64(time.Second))
	tb.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(wait):
		// Recurse to ensure we successfully take a token (race condition in simple impl, but good enough for sdk)
		// For strictness, we should re-acquire lock and check again.
		return tb.Wait(ctx)
	}
}

type FacilityLimiter struct {
	rate *TokenBucket
	sem  chan struct{}
}

type RateLimiter struct {
	config         RateLimiterConfig
	integratorRate *TokenBucket
	integratorSem  chan struct{}
	facilitiesMu   sync.Mutex
	facilities     map[string]*FacilityLimiter
}

func NewRateLimiter(config RateLimiterConfig) *RateLimiter {
	return &RateLimiter{
		config:         config,
		integratorRate: NewTokenBucket(config.MaxGetPerSecondIntegrator, config.MaxGetPerSecondIntegrator),
		integratorSem:  make(chan struct{}, config.MaxConcurrentGetIntegrator),
		facilities:     make(map[string]*FacilityLimiter),
	}
}

func (r *RateLimiter) Execute(ctx context.Context, facility string, isGet bool, op func() (interface{}, error)) (interface{}, error) {
	if !r.config.Enabled || !isGet {
		return op()
	}

	// 1. Acquire Integrator Concurrent Permit
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r.integratorSem <- struct{}{}:
		defer func() { <-r.integratorSem }()
	}

	// 2. Acquire Facility Concurrent Permit (if facility known)
	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case fl.sem <- struct{}{}:
			defer func() { <-fl.sem }()
		}
	}

	// 3. Rate Limit Wait (Global + Facility)
	if err := r.integratorRate.Wait(ctx); err != nil {
		return nil, err
	}
	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		if err := fl.rate.Wait(ctx); err != nil {
			return nil, err
		}
	}

	// 4. Retry Loop for 429
	for {
		res, err := op()
		if err != nil {
			// Check for 429
			// The Client returns generic error string currently: "API Error: %d"
			errMsg := err.Error()
			if strings.Contains(errMsg, "API Error: 429") {
				// Retry
				// Ideally parse Retry-After header, but client doesn't expose it in current error return
				// We fallback to exponential backoff or default wait
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(1 * time.Second): // Default backoff
					continue
				}
			}
			return nil, err
		}
		return res, nil
	}
}

func (r *RateLimiter) getFacilityLimiter(facility string) *FacilityLimiter {
	r.facilitiesMu.Lock()
	defer r.facilitiesMu.Unlock()

	if fl, ok := r.facilities[facility]; ok {
		return fl
	}

	fl := &FacilityLimiter{
		rate: NewTokenBucket(r.config.MaxGetPerSecondPerFacility, r.config.MaxGetPerSecondPerFacility),
		sem:  make(chan struct{}, r.config.MaxConcurrentGetPerFacility),
	}
	r.facilities[facility] = fl
	return fl
}
